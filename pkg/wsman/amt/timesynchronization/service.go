/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/amt/methods"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/wsman/client"
)

func NewTimeSynchronizationServiceWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) Service {
	return Service{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMT_TimeSynchronizationService, client),
	}
}

// Get retrieves the representation of the instance
func (service Service) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Get(nil),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Enumerate(),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.Pull(enumerationContext),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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
func (service Service) SetHighAccuracyTimeSynch(ta0, tm1, tm2 int64) (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_TimeSynchronizationService, SetHighAccuracyTimeSynch), AMT_TimeSynchronizationService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(SetHighAccuracyTimeSynch), AMT_TimeSynchronizationService, &SetHighAccuracyTimeSynch_INPUT{
		H:   "http://intel.com/wbem/wscim/1/amt-schema/1/AMT_TimeSynchronizationService",
		Ta0: ta0,
		Tm1: tm1,
		Tm2: tm2,
	})
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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

func (service Service) GetLowAccuracyTimeSynch() (response Response, err error) {
	header := service.base.WSManMessageCreator.CreateHeader(methods.GenerateAction(AMT_TimeSynchronizationService, GetLowAccuracyTimeSynch), AMT_TimeSynchronizationService, nil, "", "")
	body := service.base.WSManMessageCreator.CreateBody(methods.GenerateInputMethod(GetLowAccuracyTimeSynch), AMT_TimeSynchronizationService, nil)
	response = Response{
		Message: &client.Message{
			XMLInput: service.base.WSManMessageCreator.CreateXML(header, body),
		},
	}
	// send the message to AMT
	err = service.base.Execute(response.Message)
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