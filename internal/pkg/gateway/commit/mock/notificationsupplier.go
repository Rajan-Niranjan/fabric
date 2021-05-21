// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger"
)

type NotificationSupplier struct {
	CommitNotificationsStub        func(<-chan struct{}, string) (<-chan *ledger.CommitNotification, error)
	commitNotificationsMutex       sync.RWMutex
	commitNotificationsArgsForCall []struct {
		arg1 <-chan struct{}
		arg2 string
	}
	commitNotificationsReturns struct {
		result1 <-chan *ledger.CommitNotification
		result2 error
	}
	commitNotificationsReturnsOnCall map[int]struct {
		result1 <-chan *ledger.CommitNotification
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NotificationSupplier) CommitNotifications(arg1 <-chan struct{}, arg2 string) (<-chan *ledger.CommitNotification, error) {
	fake.commitNotificationsMutex.Lock()
	ret, specificReturn := fake.commitNotificationsReturnsOnCall[len(fake.commitNotificationsArgsForCall)]
	fake.commitNotificationsArgsForCall = append(fake.commitNotificationsArgsForCall, struct {
		arg1 <-chan struct{}
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CommitNotifications", []interface{}{arg1, arg2})
	fake.commitNotificationsMutex.Unlock()
	if fake.CommitNotificationsStub != nil {
		return fake.CommitNotificationsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.commitNotificationsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *NotificationSupplier) CommitNotificationsCallCount() int {
	fake.commitNotificationsMutex.RLock()
	defer fake.commitNotificationsMutex.RUnlock()
	return len(fake.commitNotificationsArgsForCall)
}

func (fake *NotificationSupplier) CommitNotificationsCalls(stub func(<-chan struct{}, string) (<-chan *ledger.CommitNotification, error)) {
	fake.commitNotificationsMutex.Lock()
	defer fake.commitNotificationsMutex.Unlock()
	fake.CommitNotificationsStub = stub
}

func (fake *NotificationSupplier) CommitNotificationsArgsForCall(i int) (<-chan struct{}, string) {
	fake.commitNotificationsMutex.RLock()
	defer fake.commitNotificationsMutex.RUnlock()
	argsForCall := fake.commitNotificationsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *NotificationSupplier) CommitNotificationsReturns(result1 <-chan *ledger.CommitNotification, result2 error) {
	fake.commitNotificationsMutex.Lock()
	defer fake.commitNotificationsMutex.Unlock()
	fake.CommitNotificationsStub = nil
	fake.commitNotificationsReturns = struct {
		result1 <-chan *ledger.CommitNotification
		result2 error
	}{result1, result2}
}

func (fake *NotificationSupplier) CommitNotificationsReturnsOnCall(i int, result1 <-chan *ledger.CommitNotification, result2 error) {
	fake.commitNotificationsMutex.Lock()
	defer fake.commitNotificationsMutex.Unlock()
	fake.CommitNotificationsStub = nil
	if fake.commitNotificationsReturnsOnCall == nil {
		fake.commitNotificationsReturnsOnCall = make(map[int]struct {
			result1 <-chan *ledger.CommitNotification
			result2 error
		})
	}
	fake.commitNotificationsReturnsOnCall[i] = struct {
		result1 <-chan *ledger.CommitNotification
		result2 error
	}{result1, result2}
}

func (fake *NotificationSupplier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.commitNotificationsMutex.RLock()
	defer fake.commitNotificationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *NotificationSupplier) recordInvocation(key string, args []interface{}) {
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
