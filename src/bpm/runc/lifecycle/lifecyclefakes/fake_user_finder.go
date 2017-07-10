// Code generated by counterfeiter. DO NOT EDIT.
package lifecyclefakes

import (
	"bpm/runc/lifecycle"
	"sync"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeUserFinder struct {
	LookupStub        func(username string) (specs.User, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		username string
	}
	lookupReturns struct {
		result1 specs.User
		result2 error
	}
	lookupReturnsOnCall map[int]struct {
		result1 specs.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserFinder) Lookup(username string) (specs.User, error) {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		username string
	}{username})
	fake.recordInvocation("Lookup", []interface{}{username})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(username)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lookupReturns.result1, fake.lookupReturns.result2
}

func (fake *FakeUserFinder) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeUserFinder) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].username
}

func (fake *FakeUserFinder) LookupReturns(result1 specs.User, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 specs.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserFinder) LookupReturnsOnCall(i int, result1 specs.User, result2 error) {
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 specs.User
			result2 error
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 specs.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserFinder) recordInvocation(key string, args []interface{}) {
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

var _ lifecycle.UserFinder = new(FakeUserFinder)