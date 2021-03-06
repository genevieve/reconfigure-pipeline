// Code generated by counterfeiter. DO NOT EDIT.
package actionsfakes

import (
	"sync"

	"github.com/genevieve/reconfigure-pipeline/actions"
)

type FakeReconfigurer struct {
	ReconfigureStub        func(string, string, string, string, bool) error
	reconfigureMutex       sync.RWMutex
	reconfigureArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}
	reconfigureReturns struct {
		result1 error
	}
	reconfigureReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReconfigurer) Reconfigure(arg1 string, arg2 string, arg3 string, arg4 string, arg5 bool) error {
	fake.reconfigureMutex.Lock()
	ret, specificReturn := fake.reconfigureReturnsOnCall[len(fake.reconfigureArgsForCall)]
	fake.reconfigureArgsForCall = append(fake.reconfigureArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 bool
	}{arg1, arg2, arg3, arg4, arg5})
	fake.recordInvocation("Reconfigure", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.reconfigureMutex.Unlock()
	if fake.ReconfigureStub != nil {
		return fake.ReconfigureStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.reconfigureReturns
	return fakeReturns.result1
}

func (fake *FakeReconfigurer) ReconfigureCallCount() int {
	fake.reconfigureMutex.RLock()
	defer fake.reconfigureMutex.RUnlock()
	return len(fake.reconfigureArgsForCall)
}

func (fake *FakeReconfigurer) ReconfigureCalls(stub func(string, string, string, string, bool) error) {
	fake.reconfigureMutex.Lock()
	defer fake.reconfigureMutex.Unlock()
	fake.ReconfigureStub = stub
}

func (fake *FakeReconfigurer) ReconfigureArgsForCall(i int) (string, string, string, string, bool) {
	fake.reconfigureMutex.RLock()
	defer fake.reconfigureMutex.RUnlock()
	argsForCall := fake.reconfigureArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeReconfigurer) ReconfigureReturns(result1 error) {
	fake.reconfigureMutex.Lock()
	defer fake.reconfigureMutex.Unlock()
	fake.ReconfigureStub = nil
	fake.reconfigureReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReconfigurer) ReconfigureReturnsOnCall(i int, result1 error) {
	fake.reconfigureMutex.Lock()
	defer fake.reconfigureMutex.Unlock()
	fake.ReconfigureStub = nil
	if fake.reconfigureReturnsOnCall == nil {
		fake.reconfigureReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.reconfigureReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReconfigurer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.reconfigureMutex.RLock()
	defer fake.reconfigureMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReconfigurer) recordInvocation(key string, args []interface{}) {
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

var _ actions.Reconfigurer = new(FakeReconfigurer)
