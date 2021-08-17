// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"context"
	"io"
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/worker"
)

type FakeArtifactDestination struct {
	GetStreamInP2pUrlStub        func(context.Context, string) (string, error)
	getStreamInP2pUrlMutex       sync.RWMutex
	getStreamInP2pUrlArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStreamInP2pUrlReturns struct {
		result1 string
		result2 error
	}
	getStreamInP2pUrlReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	InitializeStreamedResourceCacheStub        func(db.UsedResourceCache, string) error
	initializeStreamedResourceCacheMutex       sync.RWMutex
	initializeStreamedResourceCacheArgsForCall []struct {
		arg1 db.UsedResourceCache
		arg2 string
	}
	initializeStreamedResourceCacheReturns struct {
		result1 error
	}
	initializeStreamedResourceCacheReturnsOnCall map[int]struct {
		result1 error
	}
	SetPrivilegedStub        func(bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		arg1 bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(context.Context, string, baggageclaim.Encoding, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArtifactDestination) GetStreamInP2pUrl(arg1 context.Context, arg2 string) (string, error) {
	fake.getStreamInP2pUrlMutex.Lock()
	ret, specificReturn := fake.getStreamInP2pUrlReturnsOnCall[len(fake.getStreamInP2pUrlArgsForCall)]
	fake.getStreamInP2pUrlArgsForCall = append(fake.getStreamInP2pUrlArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStreamInP2pUrlStub
	fakeReturns := fake.getStreamInP2pUrlReturns
	fake.recordInvocation("GetStreamInP2pUrl", []interface{}{arg1, arg2})
	fake.getStreamInP2pUrlMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeArtifactDestination) GetStreamInP2pUrlCallCount() int {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	return len(fake.getStreamInP2pUrlArgsForCall)
}

func (fake *FakeArtifactDestination) GetStreamInP2pUrlCalls(stub func(context.Context, string) (string, error)) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = stub
}

func (fake *FakeArtifactDestination) GetStreamInP2pUrlArgsForCall(i int) (context.Context, string) {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	argsForCall := fake.getStreamInP2pUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArtifactDestination) GetStreamInP2pUrlReturns(result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	fake.getStreamInP2pUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactDestination) GetStreamInP2pUrlReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	if fake.getStreamInP2pUrlReturnsOnCall == nil {
		fake.getStreamInP2pUrlReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStreamInP2pUrlReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeArtifactDestination) InitializeStreamedResourceCache(arg1 db.UsedResourceCache, arg2 string) error {
	fake.initializeStreamedResourceCacheMutex.Lock()
	ret, specificReturn := fake.initializeStreamedResourceCacheReturnsOnCall[len(fake.initializeStreamedResourceCacheArgsForCall)]
	fake.initializeStreamedResourceCacheArgsForCall = append(fake.initializeStreamedResourceCacheArgsForCall, struct {
		arg1 db.UsedResourceCache
		arg2 string
	}{arg1, arg2})
	stub := fake.InitializeStreamedResourceCacheStub
	fakeReturns := fake.initializeStreamedResourceCacheReturns
	fake.recordInvocation("InitializeStreamedResourceCache", []interface{}{arg1, arg2})
	fake.initializeStreamedResourceCacheMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArtifactDestination) InitializeStreamedResourceCacheCallCount() int {
	fake.initializeStreamedResourceCacheMutex.RLock()
	defer fake.initializeStreamedResourceCacheMutex.RUnlock()
	return len(fake.initializeStreamedResourceCacheArgsForCall)
}

func (fake *FakeArtifactDestination) InitializeStreamedResourceCacheCalls(stub func(db.UsedResourceCache, string) error) {
	fake.initializeStreamedResourceCacheMutex.Lock()
	defer fake.initializeStreamedResourceCacheMutex.Unlock()
	fake.InitializeStreamedResourceCacheStub = stub
}

func (fake *FakeArtifactDestination) InitializeStreamedResourceCacheArgsForCall(i int) (db.UsedResourceCache, string) {
	fake.initializeStreamedResourceCacheMutex.RLock()
	defer fake.initializeStreamedResourceCacheMutex.RUnlock()
	argsForCall := fake.initializeStreamedResourceCacheArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeArtifactDestination) InitializeStreamedResourceCacheReturns(result1 error) {
	fake.initializeStreamedResourceCacheMutex.Lock()
	defer fake.initializeStreamedResourceCacheMutex.Unlock()
	fake.InitializeStreamedResourceCacheStub = nil
	fake.initializeStreamedResourceCacheReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactDestination) InitializeStreamedResourceCacheReturnsOnCall(i int, result1 error) {
	fake.initializeStreamedResourceCacheMutex.Lock()
	defer fake.initializeStreamedResourceCacheMutex.Unlock()
	fake.InitializeStreamedResourceCacheStub = nil
	if fake.initializeStreamedResourceCacheReturnsOnCall == nil {
		fake.initializeStreamedResourceCacheReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initializeStreamedResourceCacheReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactDestination) SetPrivileged(arg1 bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetPrivilegedStub
	fakeReturns := fake.setPrivilegedReturns
	fake.recordInvocation("SetPrivileged", []interface{}{arg1})
	fake.setPrivilegedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArtifactDestination) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeArtifactDestination) SetPrivilegedCalls(stub func(bool) error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = stub
}

func (fake *FakeArtifactDestination) SetPrivilegedArgsForCall(i int) bool {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	argsForCall := fake.setPrivilegedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArtifactDestination) SetPrivilegedReturns(result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactDestination) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactDestination) StreamIn(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding, arg4 io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 io.Reader
	}{arg1, arg2, arg3, arg4})
	stub := fake.StreamInStub
	fakeReturns := fake.streamInReturns
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamInMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeArtifactDestination) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeArtifactDestination) StreamInCalls(stub func(context.Context, string, baggageclaim.Encoding, io.Reader) error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeArtifactDestination) StreamInArgsForCall(i int) (context.Context, string, baggageclaim.Encoding, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeArtifactDestination) StreamInReturns(result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactDestination) StreamInReturnsOnCall(i int, result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeArtifactDestination) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	fake.initializeStreamedResourceCacheMutex.RLock()
	defer fake.initializeStreamedResourceCacheMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArtifactDestination) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ worker.ArtifactDestination = new(FakeArtifactDestination)
