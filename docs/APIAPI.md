# \APIAPI

All URIs are relative to *https://e.coding.net/open-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAPIDoc**](APIAPI.md#CreateAPIDoc) | **Post** /?action&#x3D;CreateAPIDoc | API 文档创建
[**DeleteAPIDoc**](APIAPI.md#DeleteAPIDoc) | **Post** /?action&#x3D;DeleteAPIDoc | API 文档删除
[**DescribeAPIDoc**](APIAPI.md#DescribeAPIDoc) | **Post** /?action&#x3D;DescribeAPIDoc | API 文档详情获取
[**DescribeAPIDocList**](APIAPI.md#DescribeAPIDocList) | **Post** /?action&#x3D;DescribeAPIDocList | API 文档列表查询
[**ModifyAPIDocBaseInfo**](APIAPI.md#ModifyAPIDocBaseInfo) | **Post** /?action&#x3D;ModifyAPIDocBaseInfo | API 文档基础信息修改
[**ModifyAPIDocContent**](APIAPI.md#ModifyAPIDocContent) | **Post** /?action&#x3D;ModifyAPIDocContent | API 文档内容发布



## CreateAPIDoc

> ModifyAPIDocContent200Response CreateAPIDoc(ctx).Authorization(authorization).Action(action).CreateAPIDocRequest(createAPIDocRequest).Execute()

API 文档创建



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "CreateAPIDoc" // string | Action (default to "CreateAPIDoc")
	createAPIDocRequest := *openapiclient.NewCreateAPIDocRequest("Content_example", "ContentType_example", "PrefixUri_example", "ProjectName_example", "Title_example") // CreateAPIDocRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.CreateAPIDoc(context.Background()).Authorization(authorization).Action(action).CreateAPIDocRequest(createAPIDocRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.CreateAPIDoc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAPIDoc`: ModifyAPIDocContent200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.CreateAPIDoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAPIDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateAPIDoc&quot;]
 **createAPIDocRequest** | [**CreateAPIDocRequest**](CreateAPIDocRequest.md) |  | 

### Return type

[**ModifyAPIDocContent200Response**](ModifyAPIDocContent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAPIDoc

> DeleteAPIDoc200Response DeleteAPIDoc(ctx).Authorization(authorization).Action(action).DeleteAPIDocRequest(deleteAPIDocRequest).Execute()

API 文档删除



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DeleteAPIDoc" // string | Action (default to "DeleteAPIDoc")
	deleteAPIDocRequest := *openapiclient.NewDeleteAPIDocRequest(int32(123), "ProjectName_example") // DeleteAPIDocRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.DeleteAPIDoc(context.Background()).Authorization(authorization).Action(action).DeleteAPIDocRequest(deleteAPIDocRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.DeleteAPIDoc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAPIDoc`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.DeleteAPIDoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAPIDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteAPIDoc&quot;]
 **deleteAPIDocRequest** | [**DeleteAPIDocRequest**](DeleteAPIDocRequest.md) |  | 

### Return type

[**DeleteAPIDoc200Response**](DeleteAPIDoc200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAPIDoc

> ModifyAPIDocContent200Response DescribeAPIDoc(ctx).Authorization(authorization).Action(action).DescribeAPIDocRequest(describeAPIDocRequest).Execute()

API 文档详情获取



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAPIDoc" // string | Action (default to "DescribeAPIDoc")
	describeAPIDocRequest := *openapiclient.NewDescribeAPIDocRequest(int32(123), "ProjectName_example") // DescribeAPIDocRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.DescribeAPIDoc(context.Background()).Authorization(authorization).Action(action).DescribeAPIDocRequest(describeAPIDocRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.DescribeAPIDoc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAPIDoc`: ModifyAPIDocContent200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.DescribeAPIDoc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAPIDocRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAPIDoc&quot;]
 **describeAPIDocRequest** | [**DescribeAPIDocRequest**](DescribeAPIDocRequest.md) |  | 

### Return type

[**ModifyAPIDocContent200Response**](ModifyAPIDocContent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeAPIDocList

> DescribeAPIDocList200Response DescribeAPIDocList(ctx).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()

API 文档列表查询



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "DescribeAPIDocList" // string | Action (default to "DescribeAPIDocList")
	describeAPIDocListRequest := *openapiclient.NewDescribeAPIDocListRequest("ProjectName_example") // DescribeAPIDocListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.DescribeAPIDocList(context.Background()).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.DescribeAPIDocList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeAPIDocList`: DescribeAPIDocList200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.DescribeAPIDocList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeAPIDocListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeAPIDocList&quot;]
 **describeAPIDocListRequest** | [**DescribeAPIDocListRequest**](DescribeAPIDocListRequest.md) |  | 

### Return type

[**DescribeAPIDocList200Response**](DescribeAPIDocList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAPIDocBaseInfo

> ModifyAPIDocContent200Response ModifyAPIDocBaseInfo(ctx).Authorization(authorization).Action(action).ModifyAPIDocBaseInfoRequest(modifyAPIDocBaseInfoRequest).Execute()

API 文档基础信息修改



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyAPIDocBaseInfo" // string | Action (default to "ModifyAPIDocBaseInfo")
	modifyAPIDocBaseInfoRequest := *openapiclient.NewModifyAPIDocBaseInfoRequest(int32(123), "ProjectName_example") // ModifyAPIDocBaseInfoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.ModifyAPIDocBaseInfo(context.Background()).Authorization(authorization).Action(action).ModifyAPIDocBaseInfoRequest(modifyAPIDocBaseInfoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.ModifyAPIDocBaseInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAPIDocBaseInfo`: ModifyAPIDocContent200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.ModifyAPIDocBaseInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyAPIDocBaseInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyAPIDocBaseInfo&quot;]
 **modifyAPIDocBaseInfoRequest** | [**ModifyAPIDocBaseInfoRequest**](ModifyAPIDocBaseInfoRequest.md) |  | 

### Return type

[**ModifyAPIDocContent200Response**](ModifyAPIDocContent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyAPIDocContent

> ModifyAPIDocContent200Response ModifyAPIDocContent(ctx).Authorization(authorization).Action(action).ModifyAPIDocContentRequest(modifyAPIDocContentRequest).Execute()

API 文档内容发布



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authorization := "Bearer fHgnEkb9QQhE1818CUp7gh_pb2_WwSCYNYHSALIwympsIrqFIqmeRTe-Cl3mHGm_uTCjZg1P0uDJva2hLU4iTg28gzgTnxBVzCbgZcNSOz_cNcj68ZVosx9bZp4SX1ve" // string | 认证信息
	action := "ModifyAPIDocContent" // string | Action (default to "ModifyAPIDocContent")
	modifyAPIDocContentRequest := *openapiclient.NewModifyAPIDocContentRequest(int32(123), "Content_example", "ContentType_example", "ProjectName_example") // ModifyAPIDocContentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.APIAPI.ModifyAPIDocContent(context.Background()).Authorization(authorization).Action(action).ModifyAPIDocContentRequest(modifyAPIDocContentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `APIAPI.ModifyAPIDocContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyAPIDocContent`: ModifyAPIDocContent200Response
	fmt.Fprintf(os.Stdout, "Response from `APIAPI.ModifyAPIDocContent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyAPIDocContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyAPIDocContent&quot;]
 **modifyAPIDocContentRequest** | [**ModifyAPIDocContentRequest**](ModifyAPIDocContentRequest.md) |  | 

### Return type

[**ModifyAPIDocContent200Response**](ModifyAPIDocContent200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

