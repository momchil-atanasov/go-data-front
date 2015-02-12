// This file was generated by counterfeiter
package mtl_test_fake

import (
	"sync"

	"github.com/momchil-atanasov/go-data-front/mtl"
)

type FakeScannerHandler struct {
	OnCommentStub        func(comment string) error
	onCommentMutex       sync.RWMutex
	onCommentArgsForCall []struct {
		comment string
	}
	onCommentReturns struct {
		result1 error
	}
	OnMaterialStub        func(name string) error
	onMaterialMutex       sync.RWMutex
	onMaterialArgsForCall []struct {
		name string
	}
	onMaterialReturns struct {
		result1 error
	}
}

func (fake *FakeScannerHandler) OnComment(comment string) error {
	fake.onCommentMutex.Lock()
	fake.onCommentArgsForCall = append(fake.onCommentArgsForCall, struct {
		comment string
	}{comment})
	fake.onCommentMutex.Unlock()
	if fake.OnCommentStub != nil {
		return fake.OnCommentStub(comment)
	} else {
		return fake.onCommentReturns.result1
	}
}

func (fake *FakeScannerHandler) OnCommentCallCount() int {
	fake.onCommentMutex.RLock()
	defer fake.onCommentMutex.RUnlock()
	return len(fake.onCommentArgsForCall)
}

func (fake *FakeScannerHandler) OnCommentArgsForCall(i int) string {
	fake.onCommentMutex.RLock()
	defer fake.onCommentMutex.RUnlock()
	return fake.onCommentArgsForCall[i].comment
}

func (fake *FakeScannerHandler) OnCommentReturns(result1 error) {
	fake.OnCommentStub = nil
	fake.onCommentReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScannerHandler) OnMaterial(name string) error {
	fake.onMaterialMutex.Lock()
	fake.onMaterialArgsForCall = append(fake.onMaterialArgsForCall, struct {
		name string
	}{name})
	fake.onMaterialMutex.Unlock()
	if fake.OnMaterialStub != nil {
		return fake.OnMaterialStub(name)
	} else {
		return fake.onMaterialReturns.result1
	}
}

func (fake *FakeScannerHandler) OnMaterialCallCount() int {
	fake.onMaterialMutex.RLock()
	defer fake.onMaterialMutex.RUnlock()
	return len(fake.onMaterialArgsForCall)
}

func (fake *FakeScannerHandler) OnMaterialArgsForCall(i int) string {
	fake.onMaterialMutex.RLock()
	defer fake.onMaterialMutex.RUnlock()
	return fake.onMaterialArgsForCall[i].name
}

func (fake *FakeScannerHandler) OnMaterialReturns(result1 error) {
	fake.OnMaterialStub = nil
	fake.onMaterialReturns = struct {
		result1 error
	}{result1}
}

var _ mtl.ScannerHandler = new(FakeScannerHandler)
