package exec

import (
	"context"
	"fmt"
	"strings"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/exec/artifact"
	"github.com/concourse/concourse/atc/resource"
	"github.com/concourse/concourse/atc/worker"
)

type PutInputNotFoundError struct {
	Input string
}

func (e PutInputNotFoundError) Error() string {
	return fmt.Sprintf("input not found: %s", e.Input)
}

type PutInputs interface {
	FindAll(*artifact.Repository) ([]worker.InputSource, error)
}

type allInputs struct{}

func NewAllInputs() PutInputs {
	return &allInputs{}
}

func (i allInputs) FindAll(artifacts *artifact.Repository) ([]worker.InputSource, error) {
	inputs := []worker.InputSource{}

	for name, source := range artifacts.AsMap() {
		inputs = append(inputs, &putInputSource{
			name:   name,
			source: PutResourceSource{source},
		})
	}

	return inputs, nil
}

type specificInputs struct {
	inputs []string
}

func NewSpecificInputs(inputs []string) PutInputs {
	return &specificInputs{
		inputs: inputs,
	}
}

func (i specificInputs) FindAll(artifacts *artifact.Repository) ([]worker.InputSource, error) {
	artifactsMap := artifacts.AsMap()

	inputs := []worker.InputSource{}
	for _, i := range i.inputs {
		artifactSource, found := artifactsMap[artifact.Name(i)]
		if !found {
			return nil, PutInputNotFoundError{Input: i}
		}

		inputs = append(inputs, &putInputSource{
			name:   artifact.Name(i),
			source: PutResourceSource{artifactSource},
		})
	}

	return inputs, nil
}

type detectInputs struct {
	inputs []string
}

func detectInputsFromParam(value interface{}) []string {
	switch actual := value.(type) {
	case string:
		input := actual
		if idx := strings.IndexByte(actual, '/'); idx >= 0 {
			input = actual[:idx]
		}
		return []string{input}
	case map[string]interface{}:
		var inputs []string
		for _, value := range actual {
			inputs = append(inputs, detectInputsFromParam(value)...)
		}
		return inputs
	case []interface{}:
		var inputs []string
		for _, value := range actual {
			inputs = append(inputs, detectInputsFromParam(value)...)
		}
		return inputs
	default:
		return []string{}
	}
}

func NewDetectInputs(params map[string]interface{}) PutInputs {
	var inputs []string

	for _, value := range params {
		inputs = append(inputs, detectInputsFromParam(value)...)
	}
	return &detectInputs{
		inputs: inputs,
	}
}

func (i detectInputs) FindAll(artifacts *artifact.Repository) ([]worker.InputSource, error) {
	artifactsMap := artifacts.AsMap()

	inputs := []worker.InputSource{}
	for _, i := range i.inputs {
		artifactSource, found := artifactsMap[artifact.Name(i)]
		if found {
			inputs = append(inputs, &putInputSource{
				name:   artifact.Name(i),
				source: PutResourceSource{artifactSource},
			})
		}
	}

	return inputs, nil
}

type putInputSource struct {
	name   artifact.Name
	source worker.ArtifactSource
}

func (s *putInputSource) Source() worker.ArtifactSource { return s.source }

func (s *putInputSource) DestinationPath() string {
	return resource.ResourcesDir("put/" + string(s.name))
}

type PutResourceSource struct {
	worker.ArtifactSource
}

func (source PutResourceSource) StreamTo(ctx context.Context, logger lager.Logger, dest worker.ArtifactDestination) error {
	return source.ArtifactSource.StreamTo(ctx, logger, dest)
}
