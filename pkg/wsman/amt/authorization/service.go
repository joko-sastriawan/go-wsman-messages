/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package authorization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

// Describes the Authorization Service, which is responsible for Access Control management in the Intel(R) AMT subsystem.
// Additional Notes:
// 1) Realms 'AuditLogRealm' (20) and 'ACLRealm' (21) are supported only in Intel AMT Release 4.0 and later releases.
// 2) Realm 'DTRealm' (23) is supported only in 'ME 5.1' and Intel AMT Release 5.1 and later releases.
// 3) All the methods of 'AMT_AuthorizationService' except for 'Get' are not supported in Remote Connectivity Service provisioning mode

func NewServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) AuthorizationService {
	return AuthorizationService{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_AuthorizationService, client),
	}
}

// Get retrieves the representation of the instance
func (as AuthorizationService) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.Get(nil),
		},
	}
	// send the message to AMT
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Enumerates the instances of this class
func (as AuthorizationService) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// Pulls instances of this class, following an Enumerate operation
func (as AuthorizationService) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}
	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// EnumerateUserAclEntries enumerates entries in the User Access Control List (ACL).
func (as AuthorizationService) EnumerateUserAclEntries(startIndex int) (response Response, err error) {
	if startIndex == 0 {
		startIndex = 1
	}
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, EnumerateUserAclEntries), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(EnumerateUserAclEntries), AMT_AuthorizationService, &EnumerateUserAclEntries_INPUT{StartIndex: startIndex})
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) GetAclEnabledState(handle int) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, GetAclEnabledState), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAclEnabledState), AMT_AuthorizationService, &GetAclEnabledState_INPUT{Handle: handle})
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) GetAdminAclEntry() (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, GetAdminAclEntry), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAdminAclEntry), AMT_AuthorizationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) GetAdminAclEntryStatus() (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, GetAdminAclEntryStatus), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAdminAclEntryStatus), AMT_AuthorizationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) GetAdminNetAclEntryStatus() (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, GetAdminNetAclEntryStatus), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetAdminNetAclEntryStatus), AMT_AuthorizationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

// GetUserAclEntryEx reads a user entry from the Intel(R) AMT device.
func (as AuthorizationService) GetUserAclEntryEx(handle int) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, GetUserAclEntryEx), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetUserAclEntryEx), AMT_AuthorizationService, &GetUserAclEntryEx_INPUT{Handle: handle})
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) RemoveUserAclEntry(handle int) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, RemoveUserAclEntry), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(RemoveUserAclEntry), AMT_AuthorizationService, &RemoveUserAclEntry_INPUT{Handle: handle})
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) SetAclEnabledState(handle int, enabled bool) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, SetAclEnabledState), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetAclEnabledState), AMT_AuthorizationService, &SetAclEnabledState_INPUT{Handle: handle, Enabled: enabled})
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}

func (as AuthorizationService) SetAdminACLEntryEx(username, digestPassword string) (response Response, err error) {
	header := as.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_AuthorizationService, SetAdminAclEntryEx), AMT_AuthorizationService, nil, "", "")
	body := as.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetAdminAclEntryEx), AMT_AuthorizationService, &SetAdminACLEntryEx_INPUT{Username: username, DigestPassword: digestPassword})
	response = Response{
		Message: &client.Message{
			XMLInput: as.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	err = as.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}
	return
}