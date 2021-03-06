package mts

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

// AddCategory invokes the mts.AddCategory API synchronously
// api document: https://help.aliyun.com/api/mts/addcategory.html
func (client *Client) AddCategory(request *AddCategoryRequest) (response *AddCategoryResponse, err error) {
	response = CreateAddCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// AddCategoryWithChan invokes the mts.AddCategory API asynchronously
// api document: https://help.aliyun.com/api/mts/addcategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCategoryWithChan(request *AddCategoryRequest) (<-chan *AddCategoryResponse, <-chan error) {
	responseChan := make(chan *AddCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddCategory(request)
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

// AddCategoryWithCallback invokes the mts.AddCategory API asynchronously
// api document: https://help.aliyun.com/api/mts/addcategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddCategoryWithCallback(request *AddCategoryRequest, callback func(response *AddCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddCategoryResponse
		var err error
		defer close(result)
		response, err = client.AddCategory(request)
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

// AddCategoryRequest is the request struct for api AddCategory
type AddCategoryRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CateName             string           `position:"Query" name:"CateName"`
	ParentId             requests.Integer `position:"Query" name:"ParentId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// AddCategoryResponse is the response struct for api AddCategory
type AddCategoryResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Category  Category `json:"Category" xml:"Category"`
}

// CreateAddCategoryRequest creates a request to invoke AddCategory API
func CreateAddCategoryRequest() (request *AddCategoryRequest) {
	request = &AddCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddCategory", "mts", "openAPI")
	return
}

// CreateAddCategoryResponse creates a response to parse from AddCategory response
func CreateAddCategoryResponse() (response *AddCategoryResponse) {
	response = &AddCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
