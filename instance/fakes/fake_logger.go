// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/instance"
)

type FakeLogger struct {
	DebugStub        func(string, string, ...interface{})
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}
	ErrorStub        func(string, string, ...interface{})
	errorMutex       sync.RWMutex
	errorArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}
	InfoStub        func(string, string, ...interface{})
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}
	WarnStub        func(string, string, ...interface{})
	warnMutex       sync.RWMutex
	warnArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) Debug(arg1 string, arg2 string, arg3 ...interface{}) {
	fake.debugMutex.Lock()
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("Debug", []interface{}{arg1, arg2, arg3})
	fake.debugMutex.Unlock()
	if fake.DebugStub != nil {
		fake.DebugStub(arg1, arg2, arg3...)
	}
}

func (fake *FakeLogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *FakeLogger) DebugCalls(stub func(string, string, ...interface{})) {
	fake.debugMutex.Lock()
	defer fake.debugMutex.Unlock()
	fake.DebugStub = stub
}

func (fake *FakeLogger) DebugArgsForCall(i int) (string, string, []interface{}) {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	argsForCall := fake.debugArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogger) Error(arg1 string, arg2 string, arg3 ...interface{}) {
	fake.errorMutex.Lock()
	fake.errorArgsForCall = append(fake.errorArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("Error", []interface{}{arg1, arg2, arg3})
	fake.errorMutex.Unlock()
	if fake.ErrorStub != nil {
		fake.ErrorStub(arg1, arg2, arg3...)
	}
}

func (fake *FakeLogger) ErrorCallCount() int {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	return len(fake.errorArgsForCall)
}

func (fake *FakeLogger) ErrorCalls(stub func(string, string, ...interface{})) {
	fake.errorMutex.Lock()
	defer fake.errorMutex.Unlock()
	fake.ErrorStub = stub
}

func (fake *FakeLogger) ErrorArgsForCall(i int) (string, string, []interface{}) {
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	argsForCall := fake.errorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogger) Info(arg1 string, arg2 string, arg3 ...interface{}) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("Info", []interface{}{arg1, arg2, arg3})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		fake.InfoStub(arg1, arg2, arg3...)
	}
}

func (fake *FakeLogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeLogger) InfoCalls(stub func(string, string, ...interface{})) {
	fake.infoMutex.Lock()
	defer fake.infoMutex.Unlock()
	fake.InfoStub = stub
}

func (fake *FakeLogger) InfoArgsForCall(i int) (string, string, []interface{}) {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	argsForCall := fake.infoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogger) Warn(arg1 string, arg2 string, arg3 ...interface{}) {
	fake.warnMutex.Lock()
	fake.warnArgsForCall = append(fake.warnArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("Warn", []interface{}{arg1, arg2, arg3})
	fake.warnMutex.Unlock()
	if fake.WarnStub != nil {
		fake.WarnStub(arg1, arg2, arg3...)
	}
}

func (fake *FakeLogger) WarnCallCount() int {
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	return len(fake.warnArgsForCall)
}

func (fake *FakeLogger) WarnCalls(stub func(string, string, ...interface{})) {
	fake.warnMutex.Lock()
	defer fake.warnMutex.Unlock()
	fake.WarnStub = stub
}

func (fake *FakeLogger) WarnArgsForCall(i int) (string, string, []interface{}) {
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	argsForCall := fake.warnArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	fake.errorMutex.RLock()
	defer fake.errorMutex.RUnlock()
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	fake.warnMutex.RLock()
	defer fake.warnMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ instance.Logger = new(FakeLogger)
