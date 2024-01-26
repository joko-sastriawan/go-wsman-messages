/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package wsman facilitates access to AMT, CIM, and IPS classes for communication with Intel® AMT devices.
package wsman

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/cim"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/ips"
)

// NewMessages instantiates a new Messages class with client connection parameters
func NewMessages(cp ClientParameters) Messages {
	client := client.NewWsman(cp.Target, cp.Username, cp.Password, cp.UseDigest, cp.UseTLS, cp.SelfSignedAllowed)
	m := Messages{
		client: client,
	}
	m.AMT = amt.NewMessages(client)
	m.CIM = cim.NewMessages(client)
	m.IPS = ips.NewMessages(client)
	return m
}