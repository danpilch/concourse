package exec

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagerctx"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db/lock"
	"github.com/concourse/concourse/tracing"
	"github.com/concourse/concourse/vars"
	gocache "github.com/patrickmn/go-cache"
)

type VarNotFoundError struct {
	Name string
}

func (e VarNotFoundError) Error() string {
	return fmt.Sprintf("var %s not found", e.Name)
}

type GetVarStep struct {
	planID            atc.PlanID // TODO: not being used, maybe drop it
	plan              atc.GetVarPlan
	metadata          StepMetadata
	delegateFactory   BuildStepDelegateFactory
	lockFactory       lock.LockFactory
	cache             *gocache.Cache
	varSourcePool     creds.VarSourcePool
	secretCacheConfig creds.SecretCacheConfig
}

func NewGetVarStep(
	planID atc.PlanID,
	plan atc.GetVarPlan,
	metadata StepMetadata,
	delegateFactory BuildStepDelegateFactory, // XXX: not needed yet b/c no image fetching but WHATEVER
	cache *gocache.Cache,
	lockFactory lock.LockFactory,
	varSourcePool creds.VarSourcePool,
	secretCacheConfig creds.SecretCacheConfig,
) Step {
	return &GetVarStep{
		planID:            planID,
		plan:              plan,
		metadata:          metadata,
		delegateFactory:   delegateFactory,
		lockFactory:       lockFactory,
		cache:             cache,
		varSourcePool:     varSourcePool,
		secretCacheConfig: secretCacheConfig,
	}
}

func (step *GetVarStep) Run(ctx context.Context, state RunState) (bool, error) {
	delegate := step.delegateFactory.BuildStepDelegate(state)
	ctx, span := delegate.StartSpan(ctx, "get_var", tracing.Attrs{
		"name": step.plan.Name,
	})

	ok, err := step.run(ctx, state, delegate)
	tracing.End(span, err)

	return ok, err
}

func (step *GetVarStep) run(ctx context.Context, state RunState, delegate BuildStepDelegate) (bool, error) {
	logger := lagerctx.FromContext(ctx)
	logger = logger.Session("get-var-step", lager.Data{
		"step-name": step.plan.Name,
		"job-id":    step.metadata.JobID,
	})

	delegate.Initializing(logger)

	stderr := delegate.Stderr()

	fmt.Fprintln(stderr, "\x1b[1;33mWARNING: the get_var step is experimental and subject to change!\x1b[0m")
	fmt.Fprintln(stderr, "")

	delegate.Starting(logger)

	hash, err := step.hashVarIdentifier(step.plan.Path, step.plan.Type, step.plan.Source)
	if err != nil {
		return false, fmt.Errorf("hash var identifier: %w", err)
	}

	for {
		var acquired bool
		lock, acquired, err := step.lockFactory.Acquire(logger, lock.NewGetVarStepLockID(step.metadata.BuildID, hash))
		if err != nil {
			return false, fmt.Errorf("acquire lock: %w", err)
		}

		if acquired {
			defer lock.Release()
			break
		}

		time.Sleep(time.Second)
	}

	varsRef := vars.Reference{
		Source: step.plan.Name,
		Path:   step.plan.Path,
		Fields: step.plan.Fields,
	}

	value, found, err := state.LocalVariables().Get(varsRef)
	if err != nil {
		return false, fmt.Errorf("get var from build vars: %w", err)
	}
	// If the var already exists in the builds vars, nothing needs to be done
	if found {
		result, err := vars.Traverse(value, varsRef.String(), step.plan.Fields)
		if err != nil {
			return false, err
		}

		state.StoreResult(step.planID, result)
		delegate.Finished(logger, true)
		return true, nil
	}

	value, found = step.cache.Get(hash)

	// If the var exists within the cache, use the value in the cache
	if found {
		result, err := vars.Traverse(value, varsRef.String(), step.plan.Fields)
		if err != nil {
			return false, err
		}

		state.LocalVariables().SetVar(step.plan.Name, step.plan.Path, value, !step.plan.Reveal)
		state.StoreResult(step.planID, result)

		delegate.Finished(logger, true)
		return true, nil
	}

	value, found, err = step.runGetVar(state, delegate, varsRef, ctx, logger)
	if err != nil {
		return false, err
	}

	if !found {
		return false, nil
	}

	result, err := vars.Traverse(value, varsRef.String(), step.plan.Fields)
	if err != nil {
		return false, err
	}

	step.cache.Add(hash, value, time.Second)

	state.LocalVariables().SetVar(step.plan.Name, step.plan.Path, value, !step.plan.Reveal)

	state.StoreResult(step.planID, result)

	delegate.Finished(logger, true)

	return true, nil
}

func (step *GetVarStep) runGetVar(state RunState, delegate BuildStepDelegate, ref vars.Reference, ctx context.Context, logger lager.Logger) (interface{}, bool, error) {
	return nil, false, nil
	// // Var is evaluated by global credential manager
	// if ref.Source == "" {
	// 	globalVars := creds.NewVariables(step.globalSecrets, step.metadata.TeamName, step.metadata.PipelineName, false)
	// 	return globalVars.Get(ref)
	// }

	// // Loop over each var source and try to match a var source to the source
	// // provided in the var
	// varSourceConfig, found := state.VarSourceConfigs().Lookup(ref.Source)
	// if !found {
	// 	return nil, nil, false, vars.MissingSourceError{Name: ref.String(), Source: ref.Source}
	// }

	// // Grab out the manager factory for th
	// factory := creds.ManagerFactories()[ref.Source]
	// if factory == nil {
	// 	return nil, nil, false, fmt.Errorf("unknown credential manager type: %s", ref.Source)
	// }

	// // Evaluate the var source's config. If the config of the var source has
	// // templated vars then it will end up recursing to evaluate the var
	// // source config's vars until it is able to evaluate a source that does
	// // not have any templated vars or is evaluated using the global
	// // credential manager.
	// source, ok := varSourceConfig.Config.(map[string]interface{})
	// if !ok {
	// 	return nil, nil, false, fmt.Errorf("invalid source for %s", varSourceConfig.Name)
	// }

	// // Pass in a list of var source configs that don't include the var source
	// // that we are currently trying to evaluate
	// evaluatedConfig, err := creds.NewSource(delegate.Variables(ctx, state.VarSourceConfigs().Without(ref.Source)), source).Evaluate()
	// if err != nil {
	// 	return nil, nil, false, fmt.Errorf("evaluate: %w", err)
	// }

	// secrets, err := step.varSourcePool.FindOrCreate(logger, evaluatedConfig, factory)
	// if err != nil {
	// 	return nil, nil, false, fmt.Errorf("find or create var source: %w", err)
	// }

	// 			result, _, found, err := secrets.Get(secretPath)
	// 			if err != nil {
	// 				return nil, false, err
	// 			}
	// 			if !found {
	// 				continue
	// 			}
	// 			return result, true, nil
	// 		}
	// 		return nil, false, nil
	// 	}
	// }

	// return nil, false, vars.MissingSourceError{Name: ref.String(), Source: ref.Source}
}

func (step *GetVarStep) hashVarIdentifier(path, type_ string, source atc.Source) (string, error) {
	varIdentifier, err := json.Marshal(struct {
		Path string `json:"path"`
		// TODO: Type might not be safe with prototypes, since the type is arbitrary
		Type   string     `json:"type"`
		Source atc.Source `json:"source"`
	}{path, type_, source})
	if err != nil {
		return "", err
	}

	hasher := md5.New()
	hasher.Write([]byte(varIdentifier))
	return hex.EncodeToString(hasher.Sum(nil)), nil
}