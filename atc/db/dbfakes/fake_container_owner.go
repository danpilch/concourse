// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc/db"
)

type FakeContainerOwner struct {
	CreateStub        func(db.Tx, string) (map[string]interface{}, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 db.Tx
		arg2 string
	}
	createReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	FindStub        func(db.Conn) (squirrel.Eq, bool, error)
	findMutex       sync.RWMutex
	findArgsForCall []struct {
		arg1 db.Conn
	}
	findReturns struct {
		result1 squirrel.Eq
		result2 bool
		result3 error
	}
	findReturnsOnCall map[int]struct {
		result1 squirrel.Eq
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeContainerOwner) Create(arg1 db.Tx, arg2 string) (map[string]interface{}, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 db.Tx
		arg2 string
	}{arg1, arg2})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeContainerOwner) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeContainerOwner) CreateCalls(stub func(db.Tx, string) (map[string]interface{}, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeContainerOwner) CreateArgsForCall(i int) (db.Tx, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeContainerOwner) CreateReturns(result1 map[string]interface{}, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerOwner) CreateReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeContainerOwner) Find(arg1 db.Conn) (squirrel.Eq, bool, error) {
	fake.findMutex.Lock()
	ret, specificReturn := fake.findReturnsOnCall[len(fake.findArgsForCall)]
	fake.findArgsForCall = append(fake.findArgsForCall, struct {
		arg1 db.Conn
	}{arg1})
	stub := fake.FindStub
	fakeReturns := fake.findReturns
	fake.recordInvocation("Find", []interface{}{arg1})
	fake.findMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeContainerOwner) FindCallCount() int {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	return len(fake.findArgsForCall)
}

func (fake *FakeContainerOwner) FindCalls(stub func(db.Conn) (squirrel.Eq, bool, error)) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = stub
}

func (fake *FakeContainerOwner) FindArgsForCall(i int) db.Conn {
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	argsForCall := fake.findArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeContainerOwner) FindReturns(result1 squirrel.Eq, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	fake.findReturns = struct {
		result1 squirrel.Eq
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerOwner) FindReturnsOnCall(i int, result1 squirrel.Eq, result2 bool, result3 error) {
	fake.findMutex.Lock()
	defer fake.findMutex.Unlock()
	fake.FindStub = nil
	if fake.findReturnsOnCall == nil {
		fake.findReturnsOnCall = make(map[int]struct {
			result1 squirrel.Eq
			result2 bool
			result3 error
		})
	}
	fake.findReturnsOnCall[i] = struct {
		result1 squirrel.Eq
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainerOwner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.findMutex.RLock()
	defer fake.findMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeContainerOwner) recordInvocation(key string, args []interface{}) {
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

var _ db.ContainerOwner = new(FakeContainerOwner)