package alidns

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetDomainDnssecStatus invokes the alidns.SetDomainDnssecStatus API synchronously
// api document: https://help.aliyun.com/api/alidns/setdomaindnssecstatus.html
func (client *Client) SetDomainDnssecStatus(request *SetDomainDnssecStatusRequest) (response *SetDomainDnssecStatusResponse, err error) {
	response = CreateSetDomainDnssecStatusResponse()
	err = client.DoAction(request, response)
	return
}

// SetDomainDnssecStatusWithChan invokes the alidns.SetDomainDnssecStatus API asynchronously
// api document: https://help.aliyun.com/api/alidns/setdomaindnssecstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainDnssecStatusWithChan(request *SetDomainDnssecStatusRequest) (<-chan *SetDomainDnssecStatusResponse, <-chan error) {
	responseChan := make(chan *SetDomainDnssecStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetDomainDnssecStatus(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetDomainDnssecStatusWithCallback invokes the alidns.SetDomainDnssecStatus API asynchronously
// api document: https://help.aliyun.com/api/alidns/setdomaindnssecstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetDomainDnssecStatusWithCallback(request *SetDomainDnssecStatusRequest, callback func(response *SetDomainDnssecStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetDomainDnssecStatusResponse
		var err error
		defer close(result)
		response, err = client.SetDomainDnssecStatus(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetDomainDnssecStatusRequest is the request struct for api SetDomainDnssecStatus
type SetDomainDnssecStatusRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Query" name:"DomainName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
	Status       string `position:"Query" name:"Status"`
}

// SetDomainDnssecStatusResponse is the response struct for api SetDomainDnssecStatus
type SetDomainDnssecStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetDomainDnssecStatusRequest creates a request to invoke SetDomainDnssecStatus API
func CreateSetDomainDnssecStatusRequest() (request *SetDomainDnssecStatusRequest) {
	request = &SetDomainDnssecStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "SetDomainDnssecStatus", "alidns", "openAPI")
	return
}

// CreateSetDomainDnssecStatusResponse creates a response to parse from SetDomainDnssecStatus response
func CreateSetDomainDnssecStatusResponse() (response *SetDomainDnssecStatusResponse) {
	response = &SetDomainDnssecStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
