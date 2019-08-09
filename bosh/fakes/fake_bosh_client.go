// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/bosh"
	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/orchestrator"
)

type FakeBoshClient struct {
	FindInstancesStub        func(string) ([]orchestrator.Instance, error)
	findInstancesMutex       sync.RWMutex
	findInstancesArgsForCall []struct {
		arg1 string
	}
	findInstancesReturns struct {
		result1 []orchestrator.Instance
		result2 error
	}
	findInstancesReturnsOnCall map[int]struct {
		result1 []orchestrator.Instance
		result2 error
	}
	GetManifestStub        func(string) (string, error)
	getManifestMutex       sync.RWMutex
	getManifestArgsForCall []struct {
		arg1 string
	}
	getManifestReturns struct {
		result1 string
		result2 error
	}
	getManifestReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBoshClient) FindInstances(arg1 string) ([]orchestrator.Instance, error) {
	fake.findInstancesMutex.Lock()
	ret, specificReturn := fake.findInstancesReturnsOnCall[len(fake.findInstancesArgsForCall)]
	fake.findInstancesArgsForCall = append(fake.findInstancesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindInstances", []interface{}{arg1})
	fake.findInstancesMutex.Unlock()
	if fake.FindInstancesStub != nil {
		return fake.FindInstancesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.findInstancesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) FindInstancesCallCount() int {
	fake.findInstancesMutex.RLock()
	defer fake.findInstancesMutex.RUnlock()
	return len(fake.findInstancesArgsForCall)
}

func (fake *FakeBoshClient) FindInstancesCalls(stub func(string) ([]orchestrator.Instance, error)) {
	fake.findInstancesMutex.Lock()
	defer fake.findInstancesMutex.Unlock()
	fake.FindInstancesStub = stub
}

func (fake *FakeBoshClient) FindInstancesArgsForCall(i int) string {
	fake.findInstancesMutex.RLock()
	defer fake.findInstancesMutex.RUnlock()
	argsForCall := fake.findInstancesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBoshClient) FindInstancesReturns(result1 []orchestrator.Instance, result2 error) {
	fake.findInstancesMutex.Lock()
	defer fake.findInstancesMutex.Unlock()
	fake.FindInstancesStub = nil
	fake.findInstancesReturns = struct {
		result1 []orchestrator.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) FindInstancesReturnsOnCall(i int, result1 []orchestrator.Instance, result2 error) {
	fake.findInstancesMutex.Lock()
	defer fake.findInstancesMutex.Unlock()
	fake.FindInstancesStub = nil
	if fake.findInstancesReturnsOnCall == nil {
		fake.findInstancesReturnsOnCall = make(map[int]struct {
			result1 []orchestrator.Instance
			result2 error
		})
	}
	fake.findInstancesReturnsOnCall[i] = struct {
		result1 []orchestrator.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetManifest(arg1 string) (string, error) {
	fake.getManifestMutex.Lock()
	ret, specificReturn := fake.getManifestReturnsOnCall[len(fake.getManifestArgsForCall)]
	fake.getManifestArgsForCall = append(fake.getManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetManifest", []interface{}{arg1})
	fake.getManifestMutex.Unlock()
	if fake.GetManifestStub != nil {
		return fake.GetManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getManifestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBoshClient) GetManifestCallCount() int {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	return len(fake.getManifestArgsForCall)
}

func (fake *FakeBoshClient) GetManifestCalls(stub func(string) (string, error)) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = stub
}

func (fake *FakeBoshClient) GetManifestArgsForCall(i int) string {
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	argsForCall := fake.getManifestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBoshClient) GetManifestReturns(result1 string, result2 error) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = nil
	fake.getManifestReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) GetManifestReturnsOnCall(i int, result1 string, result2 error) {
	fake.getManifestMutex.Lock()
	defer fake.getManifestMutex.Unlock()
	fake.GetManifestStub = nil
	if fake.getManifestReturnsOnCall == nil {
		fake.getManifestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getManifestReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBoshClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findInstancesMutex.RLock()
	defer fake.findInstancesMutex.RUnlock()
	fake.getManifestMutex.RLock()
	defer fake.getManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBoshClient) recordInvocation(key string, args []interface{}) {
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

var _ bosh.BoshClient = new(FakeBoshClient)
