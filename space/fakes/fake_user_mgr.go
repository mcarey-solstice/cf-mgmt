// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotalservices/cf-mgmt/space"
)

type FakeUserMgr struct {
	RemoveSpaceAuditorByUsernameStub        func(spaceGUID, userName string) error
	removeSpaceAuditorByUsernameMutex       sync.RWMutex
	removeSpaceAuditorByUsernameArgsForCall []struct {
		spaceGUID string
		userName  string
	}
	removeSpaceAuditorByUsernameReturns struct {
		result1 error
	}
	RemoveSpaceDeveloperByUsernameStub        func(spaceGUID, userName string) error
	removeSpaceDeveloperByUsernameMutex       sync.RWMutex
	removeSpaceDeveloperByUsernameArgsForCall []struct {
		spaceGUID string
		userName  string
	}
	removeSpaceDeveloperByUsernameReturns struct {
		result1 error
	}
	RemoveSpaceManagerByUsernameStub        func(spaceGUID, userName string) error
	removeSpaceManagerByUsernameMutex       sync.RWMutex
	removeSpaceManagerByUsernameArgsForCall []struct {
		spaceGUID string
		userName  string
	}
	removeSpaceManagerByUsernameReturns struct {
		result1 error
	}
	ListSpaceAuditorsStub        func(spaceGUID string) (map[string]string, error)
	listSpaceAuditorsMutex       sync.RWMutex
	listSpaceAuditorsArgsForCall []struct {
		spaceGUID string
	}
	listSpaceAuditorsReturns struct {
		result1 map[string]string
		result2 error
	}
	ListSpaceDevelopersStub        func(spaceGUID string) (map[string]string, error)
	listSpaceDevelopersMutex       sync.RWMutex
	listSpaceDevelopersArgsForCall []struct {
		spaceGUID string
	}
	listSpaceDevelopersReturns struct {
		result1 map[string]string
		result2 error
	}
	ListSpaceManagersStub        func(spaceGUID string) (map[string]string, error)
	listSpaceManagersMutex       sync.RWMutex
	listSpaceManagersArgsForCall []struct {
		spaceGUID string
	}
	listSpaceManagersReturns struct {
		result1 map[string]string
		result2 error
	}
	AssociateSpaceAuditorByUsernameStub        func(orgGUID, spaceGUID, userName string) error
	associateSpaceAuditorByUsernameMutex       sync.RWMutex
	associateSpaceAuditorByUsernameArgsForCall []struct {
		orgGUID   string
		spaceGUID string
		userName  string
	}
	associateSpaceAuditorByUsernameReturns struct {
		result1 error
	}
	AssociateSpaceDeveloperByUsernameStub        func(orgGUID, spaceGUID, userName string) error
	associateSpaceDeveloperByUsernameMutex       sync.RWMutex
	associateSpaceDeveloperByUsernameArgsForCall []struct {
		orgGUID   string
		spaceGUID string
		userName  string
	}
	associateSpaceDeveloperByUsernameReturns struct {
		result1 error
	}
	AssociateSpaceManagerByUsernameStub        func(orgGUID, spaceGUID, userName string) error
	associateSpaceManagerByUsernameMutex       sync.RWMutex
	associateSpaceManagerByUsernameArgsForCall []struct {
		orgGUID   string
		spaceGUID string
		userName  string
	}
	associateSpaceManagerByUsernameReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserMgr) RemoveSpaceAuditorByUsername(spaceGUID string, userName string) error {
	fake.removeSpaceAuditorByUsernameMutex.Lock()
	fake.removeSpaceAuditorByUsernameArgsForCall = append(fake.removeSpaceAuditorByUsernameArgsForCall, struct {
		spaceGUID string
		userName  string
	}{spaceGUID, userName})
	fake.recordInvocation("RemoveSpaceAuditorByUsername", []interface{}{spaceGUID, userName})
	fake.removeSpaceAuditorByUsernameMutex.Unlock()
	if fake.RemoveSpaceAuditorByUsernameStub != nil {
		return fake.RemoveSpaceAuditorByUsernameStub(spaceGUID, userName)
	} else {
		return fake.removeSpaceAuditorByUsernameReturns.result1
	}
}

func (fake *FakeUserMgr) RemoveSpaceAuditorByUsernameCallCount() int {
	fake.removeSpaceAuditorByUsernameMutex.RLock()
	defer fake.removeSpaceAuditorByUsernameMutex.RUnlock()
	return len(fake.removeSpaceAuditorByUsernameArgsForCall)
}

func (fake *FakeUserMgr) RemoveSpaceAuditorByUsernameArgsForCall(i int) (string, string) {
	fake.removeSpaceAuditorByUsernameMutex.RLock()
	defer fake.removeSpaceAuditorByUsernameMutex.RUnlock()
	return fake.removeSpaceAuditorByUsernameArgsForCall[i].spaceGUID, fake.removeSpaceAuditorByUsernameArgsForCall[i].userName
}

func (fake *FakeUserMgr) RemoveSpaceAuditorByUsernameReturns(result1 error) {
	fake.RemoveSpaceAuditorByUsernameStub = nil
	fake.removeSpaceAuditorByUsernameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserMgr) RemoveSpaceDeveloperByUsername(spaceGUID string, userName string) error {
	fake.removeSpaceDeveloperByUsernameMutex.Lock()
	fake.removeSpaceDeveloperByUsernameArgsForCall = append(fake.removeSpaceDeveloperByUsernameArgsForCall, struct {
		spaceGUID string
		userName  string
	}{spaceGUID, userName})
	fake.recordInvocation("RemoveSpaceDeveloperByUsername", []interface{}{spaceGUID, userName})
	fake.removeSpaceDeveloperByUsernameMutex.Unlock()
	if fake.RemoveSpaceDeveloperByUsernameStub != nil {
		return fake.RemoveSpaceDeveloperByUsernameStub(spaceGUID, userName)
	} else {
		return fake.removeSpaceDeveloperByUsernameReturns.result1
	}
}

func (fake *FakeUserMgr) RemoveSpaceDeveloperByUsernameCallCount() int {
	fake.removeSpaceDeveloperByUsernameMutex.RLock()
	defer fake.removeSpaceDeveloperByUsernameMutex.RUnlock()
	return len(fake.removeSpaceDeveloperByUsernameArgsForCall)
}

func (fake *FakeUserMgr) RemoveSpaceDeveloperByUsernameArgsForCall(i int) (string, string) {
	fake.removeSpaceDeveloperByUsernameMutex.RLock()
	defer fake.removeSpaceDeveloperByUsernameMutex.RUnlock()
	return fake.removeSpaceDeveloperByUsernameArgsForCall[i].spaceGUID, fake.removeSpaceDeveloperByUsernameArgsForCall[i].userName
}

func (fake *FakeUserMgr) RemoveSpaceDeveloperByUsernameReturns(result1 error) {
	fake.RemoveSpaceDeveloperByUsernameStub = nil
	fake.removeSpaceDeveloperByUsernameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserMgr) RemoveSpaceManagerByUsername(spaceGUID string, userName string) error {
	fake.removeSpaceManagerByUsernameMutex.Lock()
	fake.removeSpaceManagerByUsernameArgsForCall = append(fake.removeSpaceManagerByUsernameArgsForCall, struct {
		spaceGUID string
		userName  string
	}{spaceGUID, userName})
	fake.recordInvocation("RemoveSpaceManagerByUsername", []interface{}{spaceGUID, userName})
	fake.removeSpaceManagerByUsernameMutex.Unlock()
	if fake.RemoveSpaceManagerByUsernameStub != nil {
		return fake.RemoveSpaceManagerByUsernameStub(spaceGUID, userName)
	} else {
		return fake.removeSpaceManagerByUsernameReturns.result1
	}
}

func (fake *FakeUserMgr) RemoveSpaceManagerByUsernameCallCount() int {
	fake.removeSpaceManagerByUsernameMutex.RLock()
	defer fake.removeSpaceManagerByUsernameMutex.RUnlock()
	return len(fake.removeSpaceManagerByUsernameArgsForCall)
}

func (fake *FakeUserMgr) RemoveSpaceManagerByUsernameArgsForCall(i int) (string, string) {
	fake.removeSpaceManagerByUsernameMutex.RLock()
	defer fake.removeSpaceManagerByUsernameMutex.RUnlock()
	return fake.removeSpaceManagerByUsernameArgsForCall[i].spaceGUID, fake.removeSpaceManagerByUsernameArgsForCall[i].userName
}

func (fake *FakeUserMgr) RemoveSpaceManagerByUsernameReturns(result1 error) {
	fake.RemoveSpaceManagerByUsernameStub = nil
	fake.removeSpaceManagerByUsernameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserMgr) ListSpaceAuditors(spaceGUID string) (map[string]string, error) {
	fake.listSpaceAuditorsMutex.Lock()
	fake.listSpaceAuditorsArgsForCall = append(fake.listSpaceAuditorsArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("ListSpaceAuditors", []interface{}{spaceGUID})
	fake.listSpaceAuditorsMutex.Unlock()
	if fake.ListSpaceAuditorsStub != nil {
		return fake.ListSpaceAuditorsStub(spaceGUID)
	} else {
		return fake.listSpaceAuditorsReturns.result1, fake.listSpaceAuditorsReturns.result2
	}
}

func (fake *FakeUserMgr) ListSpaceAuditorsCallCount() int {
	fake.listSpaceAuditorsMutex.RLock()
	defer fake.listSpaceAuditorsMutex.RUnlock()
	return len(fake.listSpaceAuditorsArgsForCall)
}

func (fake *FakeUserMgr) ListSpaceAuditorsArgsForCall(i int) string {
	fake.listSpaceAuditorsMutex.RLock()
	defer fake.listSpaceAuditorsMutex.RUnlock()
	return fake.listSpaceAuditorsArgsForCall[i].spaceGUID
}

func (fake *FakeUserMgr) ListSpaceAuditorsReturns(result1 map[string]string, result2 error) {
	fake.ListSpaceAuditorsStub = nil
	fake.listSpaceAuditorsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUserMgr) ListSpaceDevelopers(spaceGUID string) (map[string]string, error) {
	fake.listSpaceDevelopersMutex.Lock()
	fake.listSpaceDevelopersArgsForCall = append(fake.listSpaceDevelopersArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("ListSpaceDevelopers", []interface{}{spaceGUID})
	fake.listSpaceDevelopersMutex.Unlock()
	if fake.ListSpaceDevelopersStub != nil {
		return fake.ListSpaceDevelopersStub(spaceGUID)
	} else {
		return fake.listSpaceDevelopersReturns.result1, fake.listSpaceDevelopersReturns.result2
	}
}

func (fake *FakeUserMgr) ListSpaceDevelopersCallCount() int {
	fake.listSpaceDevelopersMutex.RLock()
	defer fake.listSpaceDevelopersMutex.RUnlock()
	return len(fake.listSpaceDevelopersArgsForCall)
}

func (fake *FakeUserMgr) ListSpaceDevelopersArgsForCall(i int) string {
	fake.listSpaceDevelopersMutex.RLock()
	defer fake.listSpaceDevelopersMutex.RUnlock()
	return fake.listSpaceDevelopersArgsForCall[i].spaceGUID
}

func (fake *FakeUserMgr) ListSpaceDevelopersReturns(result1 map[string]string, result2 error) {
	fake.ListSpaceDevelopersStub = nil
	fake.listSpaceDevelopersReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUserMgr) ListSpaceManagers(spaceGUID string) (map[string]string, error) {
	fake.listSpaceManagersMutex.Lock()
	fake.listSpaceManagersArgsForCall = append(fake.listSpaceManagersArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("ListSpaceManagers", []interface{}{spaceGUID})
	fake.listSpaceManagersMutex.Unlock()
	if fake.ListSpaceManagersStub != nil {
		return fake.ListSpaceManagersStub(spaceGUID)
	} else {
		return fake.listSpaceManagersReturns.result1, fake.listSpaceManagersReturns.result2
	}
}

func (fake *FakeUserMgr) ListSpaceManagersCallCount() int {
	fake.listSpaceManagersMutex.RLock()
	defer fake.listSpaceManagersMutex.RUnlock()
	return len(fake.listSpaceManagersArgsForCall)
}

func (fake *FakeUserMgr) ListSpaceManagersArgsForCall(i int) string {
	fake.listSpaceManagersMutex.RLock()
	defer fake.listSpaceManagersMutex.RUnlock()
	return fake.listSpaceManagersArgsForCall[i].spaceGUID
}

func (fake *FakeUserMgr) ListSpaceManagersReturns(result1 map[string]string, result2 error) {
	fake.ListSpaceManagersStub = nil
	fake.listSpaceManagersReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeUserMgr) AssociateSpaceAuditorByUsername(orgGUID string, spaceGUID string, userName string) error {
	fake.associateSpaceAuditorByUsernameMutex.Lock()
	fake.associateSpaceAuditorByUsernameArgsForCall = append(fake.associateSpaceAuditorByUsernameArgsForCall, struct {
		orgGUID   string
		spaceGUID string
		userName  string
	}{orgGUID, spaceGUID, userName})
	fake.recordInvocation("AssociateSpaceAuditorByUsername", []interface{}{orgGUID, spaceGUID, userName})
	fake.associateSpaceAuditorByUsernameMutex.Unlock()
	if fake.AssociateSpaceAuditorByUsernameStub != nil {
		return fake.AssociateSpaceAuditorByUsernameStub(orgGUID, spaceGUID, userName)
	} else {
		return fake.associateSpaceAuditorByUsernameReturns.result1
	}
}

func (fake *FakeUserMgr) AssociateSpaceAuditorByUsernameCallCount() int {
	fake.associateSpaceAuditorByUsernameMutex.RLock()
	defer fake.associateSpaceAuditorByUsernameMutex.RUnlock()
	return len(fake.associateSpaceAuditorByUsernameArgsForCall)
}

func (fake *FakeUserMgr) AssociateSpaceAuditorByUsernameArgsForCall(i int) (string, string, string) {
	fake.associateSpaceAuditorByUsernameMutex.RLock()
	defer fake.associateSpaceAuditorByUsernameMutex.RUnlock()
	return fake.associateSpaceAuditorByUsernameArgsForCall[i].orgGUID, fake.associateSpaceAuditorByUsernameArgsForCall[i].spaceGUID, fake.associateSpaceAuditorByUsernameArgsForCall[i].userName
}

func (fake *FakeUserMgr) AssociateSpaceAuditorByUsernameReturns(result1 error) {
	fake.AssociateSpaceAuditorByUsernameStub = nil
	fake.associateSpaceAuditorByUsernameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserMgr) AssociateSpaceDeveloperByUsername(orgGUID string, spaceGUID string, userName string) error {
	fake.associateSpaceDeveloperByUsernameMutex.Lock()
	fake.associateSpaceDeveloperByUsernameArgsForCall = append(fake.associateSpaceDeveloperByUsernameArgsForCall, struct {
		orgGUID   string
		spaceGUID string
		userName  string
	}{orgGUID, spaceGUID, userName})
	fake.recordInvocation("AssociateSpaceDeveloperByUsername", []interface{}{orgGUID, spaceGUID, userName})
	fake.associateSpaceDeveloperByUsernameMutex.Unlock()
	if fake.AssociateSpaceDeveloperByUsernameStub != nil {
		return fake.AssociateSpaceDeveloperByUsernameStub(orgGUID, spaceGUID, userName)
	} else {
		return fake.associateSpaceDeveloperByUsernameReturns.result1
	}
}

func (fake *FakeUserMgr) AssociateSpaceDeveloperByUsernameCallCount() int {
	fake.associateSpaceDeveloperByUsernameMutex.RLock()
	defer fake.associateSpaceDeveloperByUsernameMutex.RUnlock()
	return len(fake.associateSpaceDeveloperByUsernameArgsForCall)
}

func (fake *FakeUserMgr) AssociateSpaceDeveloperByUsernameArgsForCall(i int) (string, string, string) {
	fake.associateSpaceDeveloperByUsernameMutex.RLock()
	defer fake.associateSpaceDeveloperByUsernameMutex.RUnlock()
	return fake.associateSpaceDeveloperByUsernameArgsForCall[i].orgGUID, fake.associateSpaceDeveloperByUsernameArgsForCall[i].spaceGUID, fake.associateSpaceDeveloperByUsernameArgsForCall[i].userName
}

func (fake *FakeUserMgr) AssociateSpaceDeveloperByUsernameReturns(result1 error) {
	fake.AssociateSpaceDeveloperByUsernameStub = nil
	fake.associateSpaceDeveloperByUsernameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserMgr) AssociateSpaceManagerByUsername(orgGUID string, spaceGUID string, userName string) error {
	fake.associateSpaceManagerByUsernameMutex.Lock()
	fake.associateSpaceManagerByUsernameArgsForCall = append(fake.associateSpaceManagerByUsernameArgsForCall, struct {
		orgGUID   string
		spaceGUID string
		userName  string
	}{orgGUID, spaceGUID, userName})
	fake.recordInvocation("AssociateSpaceManagerByUsername", []interface{}{orgGUID, spaceGUID, userName})
	fake.associateSpaceManagerByUsernameMutex.Unlock()
	if fake.AssociateSpaceManagerByUsernameStub != nil {
		return fake.AssociateSpaceManagerByUsernameStub(orgGUID, spaceGUID, userName)
	} else {
		return fake.associateSpaceManagerByUsernameReturns.result1
	}
}

func (fake *FakeUserMgr) AssociateSpaceManagerByUsernameCallCount() int {
	fake.associateSpaceManagerByUsernameMutex.RLock()
	defer fake.associateSpaceManagerByUsernameMutex.RUnlock()
	return len(fake.associateSpaceManagerByUsernameArgsForCall)
}

func (fake *FakeUserMgr) AssociateSpaceManagerByUsernameArgsForCall(i int) (string, string, string) {
	fake.associateSpaceManagerByUsernameMutex.RLock()
	defer fake.associateSpaceManagerByUsernameMutex.RUnlock()
	return fake.associateSpaceManagerByUsernameArgsForCall[i].orgGUID, fake.associateSpaceManagerByUsernameArgsForCall[i].spaceGUID, fake.associateSpaceManagerByUsernameArgsForCall[i].userName
}

func (fake *FakeUserMgr) AssociateSpaceManagerByUsernameReturns(result1 error) {
	fake.AssociateSpaceManagerByUsernameStub = nil
	fake.associateSpaceManagerByUsernameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserMgr) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.removeSpaceAuditorByUsernameMutex.RLock()
	defer fake.removeSpaceAuditorByUsernameMutex.RUnlock()
	fake.removeSpaceDeveloperByUsernameMutex.RLock()
	defer fake.removeSpaceDeveloperByUsernameMutex.RUnlock()
	fake.removeSpaceManagerByUsernameMutex.RLock()
	defer fake.removeSpaceManagerByUsernameMutex.RUnlock()
	fake.listSpaceAuditorsMutex.RLock()
	defer fake.listSpaceAuditorsMutex.RUnlock()
	fake.listSpaceDevelopersMutex.RLock()
	defer fake.listSpaceDevelopersMutex.RUnlock()
	fake.listSpaceManagersMutex.RLock()
	defer fake.listSpaceManagersMutex.RUnlock()
	fake.associateSpaceAuditorByUsernameMutex.RLock()
	defer fake.associateSpaceAuditorByUsernameMutex.RUnlock()
	fake.associateSpaceDeveloperByUsernameMutex.RLock()
	defer fake.associateSpaceDeveloperByUsernameMutex.RUnlock()
	fake.associateSpaceManagerByUsernameMutex.RLock()
	defer fake.associateSpaceManagerByUsernameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUserMgr) recordInvocation(key string, args []interface{}) {
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

var _ space.UserMgr = new(FakeUserMgr)