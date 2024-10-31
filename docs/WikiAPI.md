# \WikiAPI

All URIs are relative to *https://e.coding.net/open-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUploadToken**](WikiAPI.md#CreateUploadToken) | **Post** /?action&#x3D;CreateUploadToken | 上传文件的Token获取
[**CreateWiki**](WikiAPI.md#CreateWiki) | **Post** /?action&#x3D;CreateWiki | Wiki创建
[**CreateWikiByZip**](WikiAPI.md#CreateWikiByZip) | **Post** /?action&#x3D;CreateWikiByZip | Wiki 通过zip包上传
[**DeleteWiki**](WikiAPI.md#DeleteWiki) | **Post** /?action&#x3D;DeleteWiki | Wiki 移至回收站
[**DescribeImportJobStatus**](WikiAPI.md#DescribeImportJobStatus) | **Post** /?action&#x3D;DescribeImportJobStatus | zip包创建wiki的任务状态查询
[**DescribeUpdateJobStatus**](WikiAPI.md#DescribeUpdateJobStatus) | **Post** /?action&#x3D;DescribeUpdateJobStatus | zip包更新wiki的任务状态查询
[**DescribeWiki**](WikiAPI.md#DescribeWiki) | **Post** /?action&#x3D;DescribeWiki | Wiki 详情获取
[**DescribeWikiList**](WikiAPI.md#DescribeWikiList) | **Post** /?action&#x3D;DescribeWikiList | Wiki 列表详情获取
[**ModifyWiki**](WikiAPI.md#ModifyWiki) | **Post** /?action&#x3D;ModifyWiki | Wiki 更新
[**ModifyWikiByZip**](WikiAPI.md#ModifyWikiByZip) | **Post** /?action&#x3D;ModifyWikiByZip | 通过zip包更新wiki
[**ModifyWikiOrder**](WikiAPI.md#ModifyWikiOrder) | **Post** /?action&#x3D;ModifyWikiOrder | Wiki 父级修改
[**ModifyWikiTitle**](WikiAPI.md#ModifyWikiTitle) | **Post** /?action&#x3D;ModifyWikiTitle | Wiki 标题更新



## CreateUploadToken

> CreateUploadToken200Response CreateUploadToken(ctx).Authorization(authorization).Action(action).CreateUploadTokenRequest(createUploadTokenRequest).Execute()

上传文件的Token获取



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
	action := "CreateUploadToken" // string | Action (default to "CreateUploadToken")
	createUploadTokenRequest := *openapiclient.NewCreateUploadTokenRequest("FileName_example", "ProjectName_example") // CreateUploadTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.CreateUploadToken(context.Background()).Authorization(authorization).Action(action).CreateUploadTokenRequest(createUploadTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.CreateUploadToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadToken`: CreateUploadToken200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.CreateUploadToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateUploadToken&quot;]
 **createUploadTokenRequest** | [**CreateUploadTokenRequest**](CreateUploadTokenRequest.md) |  | 

### Return type

[**CreateUploadToken200Response**](CreateUploadToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWiki

> CreateWiki200Response CreateWiki(ctx).Authorization(authorization).Action(action).CreateWikiRequest(createWikiRequest).Execute()

Wiki创建



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
	action := "CreateWiki" // string | Action (default to "CreateWiki")
	createWikiRequest := *openapiclient.NewCreateWikiRequest("Content_example", int64(123), "ProjectName_example", "Title_example") // CreateWikiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.CreateWiki(context.Background()).Authorization(authorization).Action(action).CreateWikiRequest(createWikiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.CreateWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWiki`: CreateWiki200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.CreateWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateWiki&quot;]
 **createWikiRequest** | [**CreateWikiRequest**](CreateWikiRequest.md) |  | 

### Return type

[**CreateWiki200Response**](CreateWiki200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWikiByZip

> CreateWikiByZip200Response CreateWikiByZip(ctx).Authorization(authorization).Action(action).CreateWikiByZipRequest(createWikiByZipRequest).Execute()

Wiki 通过zip包上传



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
	action := "CreateWikiByZip" // string | Action (default to "CreateWikiByZip")
	createWikiByZipRequest := *openapiclient.NewCreateWikiByZipRequest("AuthToken_example", "FileName_example", "Key_example", int64(123), "ProjectName_example", int32(123)) // CreateWikiByZipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.CreateWikiByZip(context.Background()).Authorization(authorization).Action(action).CreateWikiByZipRequest(createWikiByZipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.CreateWikiByZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWikiByZip`: CreateWikiByZip200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.CreateWikiByZip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWikiByZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateWikiByZip&quot;]
 **createWikiByZipRequest** | [**CreateWikiByZipRequest**](CreateWikiByZipRequest.md) |  | 

### Return type

[**CreateWikiByZip200Response**](CreateWikiByZip200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWiki

> DeleteAPIDoc200Response DeleteWiki(ctx).Authorization(authorization).Action(action).DeleteWikiRequest(deleteWikiRequest).Execute()

Wiki 移至回收站



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
	action := "DeleteWiki" // string | Action (default to "DeleteWiki")
	deleteWikiRequest := *openapiclient.NewDeleteWikiRequest(int64(123), "ProjectName_example") // DeleteWikiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.DeleteWiki(context.Background()).Authorization(authorization).Action(action).DeleteWikiRequest(deleteWikiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.DeleteWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWiki`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.DeleteWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteWiki&quot;]
 **deleteWikiRequest** | [**DeleteWikiRequest**](DeleteWikiRequest.md) |  | 

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


## DescribeImportJobStatus

> DescribeImportJobStatus200Response DescribeImportJobStatus(ctx).Authorization(authorization).Action(action).DescribeImportJobStatusRequest(describeImportJobStatusRequest).Execute()

zip包创建wiki的任务状态查询



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
	action := "DescribeImportJobStatus" // string | Action (default to "DescribeImportJobStatus")
	describeImportJobStatusRequest := *openapiclient.NewDescribeImportJobStatusRequest("JobId_example", "ProjectName_example") // DescribeImportJobStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.DescribeImportJobStatus(context.Background()).Authorization(authorization).Action(action).DescribeImportJobStatusRequest(describeImportJobStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.DescribeImportJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeImportJobStatus`: DescribeImportJobStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.DescribeImportJobStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeImportJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeImportJobStatus&quot;]
 **describeImportJobStatusRequest** | [**DescribeImportJobStatusRequest**](DescribeImportJobStatusRequest.md) |  | 

### Return type

[**DescribeImportJobStatus200Response**](DescribeImportJobStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeUpdateJobStatus

> DescribeImportJobStatus200Response DescribeUpdateJobStatus(ctx).Authorization(authorization).Action(action).DescribeImportJobStatusRequest(describeImportJobStatusRequest).Execute()

zip包更新wiki的任务状态查询



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
	action := "DescribeUpdateJobStatus" // string | Action (default to "DescribeUpdateJobStatus")
	describeImportJobStatusRequest := *openapiclient.NewDescribeImportJobStatusRequest("JobId_example", "ProjectName_example") // DescribeImportJobStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.DescribeUpdateJobStatus(context.Background()).Authorization(authorization).Action(action).DescribeImportJobStatusRequest(describeImportJobStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.DescribeUpdateJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeUpdateJobStatus`: DescribeImportJobStatus200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.DescribeUpdateJobStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeUpdateJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeUpdateJobStatus&quot;]
 **describeImportJobStatusRequest** | [**DescribeImportJobStatusRequest**](DescribeImportJobStatusRequest.md) |  | 

### Return type

[**DescribeImportJobStatus200Response**](DescribeImportJobStatus200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWiki

> ModifyWiki200Response DescribeWiki(ctx).Authorization(authorization).Action(action).DescribeWikiRequest(describeWikiRequest).Execute()

Wiki 详情获取



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
	action := "DescribeWiki" // string | Action (default to "DescribeWiki")
	describeWikiRequest := *openapiclient.NewDescribeWikiRequest(int64(123), "ProjectName_example", int64(123)) // DescribeWikiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.DescribeWiki(context.Background()).Authorization(authorization).Action(action).DescribeWikiRequest(describeWikiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.DescribeWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWiki`: ModifyWiki200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.DescribeWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeWiki&quot;]
 **describeWikiRequest** | [**DescribeWikiRequest**](DescribeWikiRequest.md) |  | 

### Return type

[**ModifyWiki200Response**](ModifyWiki200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeWikiList

> DescribeWikiList200Response DescribeWikiList(ctx).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()

Wiki 列表详情获取



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
	action := "DescribeWikiList" // string | Action (default to "DescribeWikiList")
	describeAPIDocListRequest := *openapiclient.NewDescribeAPIDocListRequest("ProjectName_example") // DescribeAPIDocListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.DescribeWikiList(context.Background()).Authorization(authorization).Action(action).DescribeAPIDocListRequest(describeAPIDocListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.DescribeWikiList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeWikiList`: DescribeWikiList200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.DescribeWikiList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeWikiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeWikiList&quot;]
 **describeAPIDocListRequest** | [**DescribeAPIDocListRequest**](DescribeAPIDocListRequest.md) |  | 

### Return type

[**DescribeWikiList200Response**](DescribeWikiList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWiki

> ModifyWiki200Response ModifyWiki(ctx).Authorization(authorization).Action(action).ModifyWikiRequest(modifyWikiRequest).Execute()

Wiki 更新



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
	action := "ModifyWiki" // string | Action (default to "ModifyWiki")
	modifyWikiRequest := *openapiclient.NewModifyWikiRequest("Content_example", int64(123), "ProjectName_example", "Title_example") // ModifyWikiRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.ModifyWiki(context.Background()).Authorization(authorization).Action(action).ModifyWikiRequest(modifyWikiRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.ModifyWiki``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWiki`: ModifyWiki200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.ModifyWiki`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyWikiRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyWiki&quot;]
 **modifyWikiRequest** | [**ModifyWikiRequest**](ModifyWikiRequest.md) |  | 

### Return type

[**ModifyWiki200Response**](ModifyWiki200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWikiByZip

> CreateWikiByZip200Response ModifyWikiByZip(ctx).Authorization(authorization).Action(action).ModifyWikiByZipRequest(modifyWikiByZipRequest).Execute()

通过zip包更新wiki



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
	action := "ModifyWikiByZip" // string | Action (default to "ModifyWikiByZip")
	modifyWikiByZipRequest := *openapiclient.NewModifyWikiByZipRequest("AuthToken_example", "FileName_example", int64(123), "Key_example", "ProjectName_example", int32(123)) // ModifyWikiByZipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.ModifyWikiByZip(context.Background()).Authorization(authorization).Action(action).ModifyWikiByZipRequest(modifyWikiByZipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.ModifyWikiByZip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWikiByZip`: CreateWikiByZip200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.ModifyWikiByZip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyWikiByZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyWikiByZip&quot;]
 **modifyWikiByZipRequest** | [**ModifyWikiByZipRequest**](ModifyWikiByZipRequest.md) |  | 

### Return type

[**CreateWikiByZip200Response**](CreateWikiByZip200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyWikiOrder

> DeleteAPIDoc200Response ModifyWikiOrder(ctx).Authorization(authorization).Action(action).ModifyWikiOrderRequest(modifyWikiOrderRequest).Execute()

Wiki 父级修改



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
	action := "ModifyWikiOrder" // string | Action (default to "ModifyWikiOrder")
	modifyWikiOrderRequest := *openapiclient.NewModifyWikiOrderRequest(int64(123), int64(123), "ProjectName_example") // ModifyWikiOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.ModifyWikiOrder(context.Background()).Authorization(authorization).Action(action).ModifyWikiOrderRequest(modifyWikiOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.ModifyWikiOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWikiOrder`: DeleteAPIDoc200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.ModifyWikiOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyWikiOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyWikiOrder&quot;]
 **modifyWikiOrderRequest** | [**ModifyWikiOrderRequest**](ModifyWikiOrderRequest.md) |  | 

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


## ModifyWikiTitle

> ModifyWikiTitle200Response ModifyWikiTitle(ctx).Authorization(authorization).Action(action).ModifyWikiTitleRequest(modifyWikiTitleRequest).Execute()

Wiki 标题更新



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
	action := "ModifyWikiTitle" // string | Action (default to "ModifyWikiTitle")
	modifyWikiTitleRequest := *openapiclient.NewModifyWikiTitleRequest(int64(123), "ProjectName_example", "Title_example") // ModifyWikiTitleRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiAPI.ModifyWikiTitle(context.Background()).Authorization(authorization).Action(action).ModifyWikiTitleRequest(modifyWikiTitleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiAPI.ModifyWikiTitle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyWikiTitle`: ModifyWikiTitle200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiAPI.ModifyWikiTitle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyWikiTitleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyWikiTitle&quot;]
 **modifyWikiTitleRequest** | [**ModifyWikiTitleRequest**](ModifyWikiTitleRequest.md) |  | 

### Return type

[**ModifyWikiTitle200Response**](ModifyWikiTitle200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

