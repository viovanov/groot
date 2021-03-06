// Code generated by counterfeiter. DO NOT EDIT.
package imagepullerfakes

import (
	"io"
	"sync"

	"code.cloudfoundry.org/groot/imagepuller"
	"code.cloudfoundry.org/lager"
)

type FakeVolumeDriver struct {
	UnpackStub        func(logger lager.Logger, layerID string, parentIDs []string, layerTar io.Reader) (int64, error)
	unpackMutex       sync.RWMutex
	unpackArgsForCall []struct {
		logger    lager.Logger
		layerID   string
		parentIDs []string
		layerTar  io.Reader
	}
	unpackReturns struct {
		result1 int64
		result2 error
	}
	unpackReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeDriver) Unpack(logger lager.Logger, layerID string, parentIDs []string, layerTar io.Reader) (int64, error) {
	var parentIDsCopy []string
	if parentIDs != nil {
		parentIDsCopy = make([]string, len(parentIDs))
		copy(parentIDsCopy, parentIDs)
	}
	fake.unpackMutex.Lock()
	ret, specificReturn := fake.unpackReturnsOnCall[len(fake.unpackArgsForCall)]
	fake.unpackArgsForCall = append(fake.unpackArgsForCall, struct {
		logger    lager.Logger
		layerID   string
		parentIDs []string
		layerTar  io.Reader
	}{logger, layerID, parentIDsCopy, layerTar})
	fake.recordInvocation("Unpack", []interface{}{logger, layerID, parentIDsCopy, layerTar})
	fake.unpackMutex.Unlock()
	if fake.UnpackStub != nil {
		return fake.UnpackStub(logger, layerID, parentIDs, layerTar)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unpackReturns.result1, fake.unpackReturns.result2
}

func (fake *FakeVolumeDriver) UnpackCallCount() int {
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	return len(fake.unpackArgsForCall)
}

func (fake *FakeVolumeDriver) UnpackArgsForCall(i int) (lager.Logger, string, []string, io.Reader) {
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	return fake.unpackArgsForCall[i].logger, fake.unpackArgsForCall[i].layerID, fake.unpackArgsForCall[i].parentIDs, fake.unpackArgsForCall[i].layerTar
}

func (fake *FakeVolumeDriver) UnpackReturns(result1 int64, result2 error) {
	fake.UnpackStub = nil
	fake.unpackReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeDriver) UnpackReturnsOnCall(i int, result1 int64, result2 error) {
	fake.UnpackStub = nil
	if fake.unpackReturnsOnCall == nil {
		fake.unpackReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.unpackReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unpackMutex.RLock()
	defer fake.unpackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumeDriver) recordInvocation(key string, args []interface{}) {
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

var _ imagepuller.VolumeDriver = new(FakeVolumeDriver)
