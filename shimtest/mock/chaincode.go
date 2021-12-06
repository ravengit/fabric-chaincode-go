// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/ravengit/fabric-chaincode-go/shim"
)

type Chaincode struct {
	InitStub        func(shim.ChaincodeStubInterface) peer.Response
	initMutex       sync.RWMutex
	initArgsForCall []struct {
		arg1 shim.ChaincodeStubInterface
	}
	initReturns struct {
		result1 peer.Response
	}
	initReturnsOnCall map[int]struct {
		result1 peer.Response
	}
	InvokeStub        func(shim.ChaincodeStubInterface) peer.Response
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		arg1 shim.ChaincodeStubInterface
	}
	invokeReturns struct {
		result1 peer.Response
	}
	invokeReturnsOnCall map[int]struct {
		result1 peer.Response
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Chaincode) Init(arg1 shim.ChaincodeStubInterface) peer.Response {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
		arg1 shim.ChaincodeStubInterface
	}{arg1})
	fake.recordInvocation("Init", []interface{}{arg1})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initReturns
	return fakeReturns.result1
}

func (fake *Chaincode) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *Chaincode) InitCalls(stub func(shim.ChaincodeStubInterface) peer.Response) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *Chaincode) InitArgsForCall(i int) shim.ChaincodeStubInterface {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	argsForCall := fake.initArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Chaincode) InitReturns(result1 peer.Response) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 peer.Response
	}{result1}
}

func (fake *Chaincode) InitReturnsOnCall(i int, result1 peer.Response) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 peer.Response
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 peer.Response
	}{result1}
}

func (fake *Chaincode) Invoke(arg1 shim.ChaincodeStubInterface) peer.Response {
	fake.invokeMutex.Lock()
	ret, specificReturn := fake.invokeReturnsOnCall[len(fake.invokeArgsForCall)]
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		arg1 shim.ChaincodeStubInterface
	}{arg1})
	fake.recordInvocation("Invoke", []interface{}{arg1})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		return fake.InvokeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.invokeReturns
	return fakeReturns.result1
}

func (fake *Chaincode) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *Chaincode) InvokeCalls(stub func(shim.ChaincodeStubInterface) peer.Response) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = stub
}

func (fake *Chaincode) InvokeArgsForCall(i int) shim.ChaincodeStubInterface {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	argsForCall := fake.invokeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Chaincode) InvokeReturns(result1 peer.Response) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = nil
	fake.invokeReturns = struct {
		result1 peer.Response
	}{result1}
}

func (fake *Chaincode) InvokeReturnsOnCall(i int, result1 peer.Response) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = nil
	if fake.invokeReturnsOnCall == nil {
		fake.invokeReturnsOnCall = make(map[int]struct {
			result1 peer.Response
		})
	}
	fake.invokeReturnsOnCall[i] = struct {
		result1 peer.Response
	}{result1}
}

func (fake *Chaincode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Chaincode) recordInvocation(key string, args []interface{}) {
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
