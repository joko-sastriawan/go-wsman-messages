/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package tls

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/wsmantesting"
)

func TestPositiveAMT_TLSCredentialContext(t *testing.T) {
	messageID := 0
	resourceUriBase := "http://intel.com/wbem/wscim/1/amt-schema/1/"
	wsmanMessageCreator := message.NewWSManMessageCreator(resourceUriBase)
	client := wsmantesting.MockClient{
		PackageUnderTest: "amt/tls/credentialcontext",
	}
	elementUnderTest := NewTLSCredentialContextWithClient(wsmanMessageCreator, &client)

	t.Run("amt_TLSCredentialContext Tests", func(t *testing.T) {
		tests := []struct {
			name             string
			method           string
			action           string
			body             string
			extraHeader      string
			responseFunc     func() (Response, error)
			expectedResponse interface{}
		}{
			//GETS
			// {
			// 	"should create a valid AMT_TLSCredentialContext Get wsman message",
			// 	AMT_TLSCredentialContext,
			// 	wsmantesting.GET,
			// 	"",
			// 	"",
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "Get"
			// 		return elementUnderTest.Get()
			// 	},
			// 	Body{},
			// },
			//ENUMERATES
			{
				"should create a valid AMT_TLSCredentialContext Enumerate wsman message",
				AMT_TLSCredentialContext,
				wsmantesting.ENUMERATE,
				wsmantesting.ENUMERATE_BODY,
				"",
				func() (Response, error) {
					client.CurrentMessage = "Enumerate"
					return elementUnderTest.Enumerate()
				},
				Body{
					XMLName: xml.Name{Space: message.XMLBodySpace, Local: "Body"},
					EnumerateResponse: common.EnumerateResponse{
						EnumerationContext: "6B080000-0000-0000-0000-000000000000",
					},
				},
			},
			//PULLS
			// {
			// 	"should create a valid AMT_TLSCredentialContext Pull wsman message",
			// 	AMT_TLSCredentialContext,
			// 	wsmantesting.PULL,
			// 	wsmantesting.PULL_BODY,
			// 	"",
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "Pull"
			// 		return elementUnderTest.Pull(wsmantesting.EnumerationContext)
			// 	},
			// 	Body{
			// 		PullResponse: PullResponse{
			// 			XMLName: xml.Name{Space: message.XMLPullResponseSpace, Local: "PullResponse"},

			// 		},
			// 	},
			// },
			//DELETE
			// {
			// 	"should create a valid AMT_TLSCredentialContext Delete wsman message",
			// 	AMT_TLSCredentialContext,
			// 	wsmantesting.DELETE,
			// 	"",
			// 	"<w:SelectorSet><w:Selector Name=\"Name\">instanceID123</w:Selector></w:SelectorSet>",
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "Delete"
			// 		return elementUnderTest.Delete("instanceID123")
			// 	},
			// 	Body{},
			// },
			//Create
			// {
			// 	"should create a valid AMT_TLSCredentialContext Create wsman message",
			// 	AMT_TLSCredentialContext,
			// 	wsmantesting.CREATE,
			// 	"",
			// 	"", //fmt.Sprintf(`<Body><h:AMT_TLSCredentialContext xmlns:h="%sAMT_TLSCredentialContext"><h:ElementInContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_PublicKeyCertificate</w:ResourceURI><w:SelectorSet><w:Selector Name="InstanceID">%s</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementInContext><h:ElementProvidingContext><a:Address>/wsman</a:Address><a:ReferenceParameters><w:ResourceURI>%sAMT_TLSProtocolEndpointCollection</w:ResourceURI><w:SelectorSet><w:Selector Name="ElementName">TLSProtocolEndpointInstances Collection</w:Selector></w:SelectorSet></a:ReferenceParameters></h:ElementProvidingContext></h:AMT_TLSCredentialContext></Body>`, credentialContext.base.WSManMessageCreator.ResourceURIBase, credentialContext.base.WSManMessageCreator.ResourceURIBase, certHandle, credentialContext.base.WSManMessageCreator.ResourceURIBase),
			// 	func() (Response, error) {
			// 		client.CurrentMessage = "Create"
			// 		return elementUnderTest.Create("test")
			// 	},
			// 	Body{},
			// },
		}
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				expectedXMLInput := wsmantesting.ExpectedResponse(messageID, resourceUriBase, test.method, test.action, test.extraHeader, test.body)
				messageID++
				response, err := test.responseFunc()
				assert.NoError(t, err)
				assert.Equal(t, expectedXMLInput, response.XMLInput)
				assert.Equal(t, test.expectedResponse, response.Body)
			})
		}
	})
}
