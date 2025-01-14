/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveCIMMediaAccessDevice(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/mediaaccess",
	}
	elementUnderTest := NewMediaAccessDeviceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_MediaAccessDevice Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{

			//ENUMERATES
			{
				"should create and parse a valid cim_MediaAccessDevice Enumerate call",
				CIM_MediaAccessDevice,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should create and parse a valid cim_MediaAccessDevice Pull call",
				CIM_MediaAccessDevice,
				wsmantesting.PULL,
				"",
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Pull"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						MediaAccessDevices: []MediaAccessDevice{
							{
								Capabilities:            4,
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 0",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            960197124,
								OperationalStatus:       0,
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
							},
							{
								Capabilities:            4,
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 1",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            500107862,
								OperationalStatus:       0,
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}

func TestNegativeCIMMediaAccessDevice(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "cim/mediaaccess",
	}
	elementUnderTest := NewMediaAccessDeviceWithClient(wsmanMessageCreator, &client)

	t.Run("cim_MediaAccessDevice Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			extraHeaders     string
			body             string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{

			//ENUMERATES
			{
				"should handle error when cim_MediaAccessDevice Enumerate call",
				CIM_MediaAccessDevice,
				wsmantesting.ENUMERATE,
				"",
				wsmantesting.ENUMERATE_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Enumerate()
				}, Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "CE020000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			{
				"should handle error when cim_MediaAccessDevice Pull call",
				CIM_MediaAccessDevice,
				wsmantesting.PULL,
				"",
				wsmantesting.PULL_BODY,
				func() (Response, error) {
					client.CurrentMessage = "Error"
					return elementUnderTest.Pull(wsmantesting.EnumerationContext)
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					PullResponse: PullResponse{
						XMLName: xml.Name{Space: "http://schemas.xmlsoap.org/ws/2004/09/enumeration", Local: "PullResponse"},
						MediaAccessDevices: []MediaAccessDevice{
							{
								Capabilities:            4,
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 0",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            960197124,
								OperationalStatus:       0,
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
							},
							{
								Capabilities:            4,
								CreationClassName:       "CIM_MediaAccessDevice",
								DeviceID:                "MEDIA DEV 1",
								ElementName:             "Managed System Media Access Device",
								EnabledDefault:          2,
								EnabledState:            0,
								MaxMediaSize:            500107862,
								OperationalStatus:       0,
								RequestedState:          12,
								Security:                2,
								SystemCreationClassName: "CIM_ComputerSystem",
								SystemName:              "ManagedSystem",
							},
						},
					},
				},
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, "", test.body)
				messageID++
				response, err := test.responseFunc()
				assert.Error(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.NotEqual(t, test.expectedResponse, response.Body)
			})
		}
	})
}
