/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mps

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim/models"
)

const AMT_MPSUsernamePassword = "AMT_MPSUsernamePassword"

type MPSUsernamePassword struct {
	models.SharedCredential
}
type UsernamePassword struct {
	base message.Base
}

func NewMPSUsernamePassword(wsmanMessageCreator *message.WSManMessageCreator) UsernamePassword {
	return UsernamePassword{
		base: message.NewBase(wsmanMessageCreator, AMT_MPSUsernamePassword),
	}
}

// Get retrieves the representation of the instance
func (MPSUsernamePassword UsernamePassword) Get() string {
	return MPSUsernamePassword.base.Get(nil)
}

// Enumerates the instances of this class
func (MPSUsernamePassword UsernamePassword) Enumerate() string {
	return MPSUsernamePassword.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (MPSUsernamePassword UsernamePassword) Pull(enumerationContext string) string {
	return MPSUsernamePassword.base.Pull(enumerationContext)
}

// Put will change properties of the selected instance
func (MPSUsernamePassword UsernamePassword) Put(mpsUsernamePassword MPSUsernamePassword) string {
	return MPSUsernamePassword.base.Put(mpsUsernamePassword, false, nil)
}