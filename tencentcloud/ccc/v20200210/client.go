// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20200210

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2020-02-10"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCreateSDKLoginTokenRequest() (request *CreateSDKLoginTokenRequest) {
    request = &CreateSDKLoginTokenRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateSDKLoginToken")
    return
}

func NewCreateSDKLoginTokenResponse() (response *CreateSDKLoginTokenResponse) {
    response = &CreateSDKLoginTokenResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建 SDK 登录 Token。
func (c *Client) CreateSDKLoginToken(request *CreateSDKLoginTokenRequest) (response *CreateSDKLoginTokenResponse, err error) {
    if request == nil {
        request = NewCreateSDKLoginTokenRequest()
    }
    response = NewCreateSDKLoginTokenResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaffRequest() (request *CreateStaffRequest) {
    request = &CreateStaffRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "CreateStaff")
    return
}

func NewCreateStaffResponse() (response *CreateStaffResponse) {
    response = &CreateStaffResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建客服账号。
func (c *Client) CreateStaff(request *CreateStaffRequest) (response *CreateStaffResponse, err error) {
    if request == nil {
        request = NewCreateStaffRequest()
    }
    response = NewCreateStaffResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeChatMessagesRequest() (request *DescribeChatMessagesRequest) {
    request = &DescribeChatMessagesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeChatMessages")
    return
}

func NewDescribeChatMessagesResponse() (response *DescribeChatMessagesResponse) {
    response = &DescribeChatMessagesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 包括具体聊天内容
func (c *Client) DescribeChatMessages(request *DescribeChatMessagesRequest) (response *DescribeChatMessagesResponse, err error) {
    if request == nil {
        request = NewDescribeChatMessagesRequest()
    }
    response = NewDescribeChatMessagesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeIMCdrsRequest() (request *DescribeIMCdrsRequest) {
    request = &DescribeIMCdrsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeIMCdrs")
    return
}

func NewDescribeIMCdrsResponse() (response *DescribeIMCdrsResponse) {
    response = &DescribeIMCdrsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 包括全媒体和文本两种类型
func (c *Client) DescribeIMCdrs(request *DescribeIMCdrsRequest) (response *DescribeIMCdrsResponse, err error) {
    if request == nil {
        request = NewDescribeIMCdrsRequest()
    }
    response = NewDescribeIMCdrsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTelCdrRequest() (request *DescribeTelCdrRequest) {
    request = &DescribeTelCdrRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("ccc", APIVersion, "DescribeTelCdr")
    return
}

func NewDescribeTelCdrResponse() (response *DescribeTelCdrResponse) {
    response = &DescribeTelCdrResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取电话服务记录与录音
func (c *Client) DescribeTelCdr(request *DescribeTelCdrRequest) (response *DescribeTelCdrResponse, err error) {
    if request == nil {
        request = NewDescribeTelCdrRequest()
    }
    response = NewDescribeTelCdrResponse()
    err = c.Send(request, response)
    return
}
