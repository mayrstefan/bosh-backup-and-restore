// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/bosh-backup-and-restore/ssh"
	ssha "golang.org/x/crypto/ssh"
)

type FakeRemoteRunnerFactory struct {
	Stub        func(string, string, string, ssha.HostKeyCallback, []string, ssh.Logger) (ssh.RemoteRunner, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 ssha.HostKeyCallback
		arg5 []string
		arg6 ssh.Logger
	}
	returns struct {
		result1 ssh.RemoteRunner
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 ssh.RemoteRunner
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRemoteRunnerFactory) Spy(arg1 string, arg2 string, arg3 string, arg4 ssha.HostKeyCallback, arg5 []string, arg6 ssh.Logger) (ssh.RemoteRunner, error) {
	var arg5Copy []string
	if arg5 != nil {
		arg5Copy = make([]string, len(arg5))
		copy(arg5Copy, arg5)
	}
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 ssha.HostKeyCallback
		arg5 []string
		arg6 ssh.Logger
	}{arg1, arg2, arg3, arg4, arg5Copy, arg6})
	fake.recordInvocation("RemoteRunnerFactory", []interface{}{arg1, arg2, arg3, arg4, arg5Copy, arg6})
	fake.mutex.Unlock()
	if fake.Stub != nil {
		return fake.Stub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.returns.result1, fake.returns.result2
}

func (fake *FakeRemoteRunnerFactory) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *FakeRemoteRunnerFactory) Calls(stub func(string, string, string, ssha.HostKeyCallback, []string, ssh.Logger) (ssh.RemoteRunner, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *FakeRemoteRunnerFactory) ArgsForCall(i int) (string, string, string, ssha.HostKeyCallback, []string, ssh.Logger) {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1, fake.argsForCall[i].arg2, fake.argsForCall[i].arg3, fake.argsForCall[i].arg4, fake.argsForCall[i].arg5, fake.argsForCall[i].arg6
}

func (fake *FakeRemoteRunnerFactory) Returns(result1 ssh.RemoteRunner, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 ssh.RemoteRunner
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoteRunnerFactory) ReturnsOnCall(i int, result1 ssh.RemoteRunner, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 ssh.RemoteRunner
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 ssh.RemoteRunner
		result2 error
	}{result1, result2}
}

func (fake *FakeRemoteRunnerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRemoteRunnerFactory) recordInvocation(key string, args []interface{}) {
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

var _ ssh.RemoteRunnerFactory = new(FakeRemoteRunnerFactory).Spy
