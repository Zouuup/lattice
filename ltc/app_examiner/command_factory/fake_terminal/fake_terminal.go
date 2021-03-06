// This file was generated by counterfeiter
package fake_terminal

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/app_examiner/command_factory"
)

type FakeTerminal struct {
	GetWindowWidthStub        func() (uint16, error)
	getWindowWidthMutex       sync.RWMutex
	getWindowWidthArgsForCall []struct{}
	getWindowWidthReturns     struct {
		result1 uint16
		result2 error
	}
}

func (fake *FakeTerminal) GetWindowWidth() (uint16, error) {
	fake.getWindowWidthMutex.Lock()
	fake.getWindowWidthArgsForCall = append(fake.getWindowWidthArgsForCall, struct{}{})
	fake.getWindowWidthMutex.Unlock()
	if fake.GetWindowWidthStub != nil {
		return fake.GetWindowWidthStub()
	} else {
		return fake.getWindowWidthReturns.result1, fake.getWindowWidthReturns.result2
	}
}

func (fake *FakeTerminal) GetWindowWidthCallCount() int {
	fake.getWindowWidthMutex.RLock()
	defer fake.getWindowWidthMutex.RUnlock()
	return len(fake.getWindowWidthArgsForCall)
}

func (fake *FakeTerminal) GetWindowWidthReturns(result1 uint16, result2 error) {
	fake.GetWindowWidthStub = nil
	fake.getWindowWidthReturns = struct {
		result1 uint16
		result2 error
	}{result1, result2}
}

var _ command_factory.Terminal = new(FakeTerminal)
