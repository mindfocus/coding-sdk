# \ServiceHookAPI

All URIs are relative to *https://e.coding.net/open-api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceHook**](ServiceHookAPI.md#CreateServiceHook) | **Post** /?action&#x3D;CreateServiceHook | Service Hook 创建
[**DeleteServiceHook**](ServiceHookAPI.md#DeleteServiceHook) | **Post** /?action&#x3D;DeleteServiceHook | Service Hook 删除
[**DescribeEvents**](ServiceHookAPI.md#DescribeEvents) | **Post** /?action&#x3D;DescribeEvents | Service Hook 事件列表查询
[**DescribeServiceHook**](ServiceHookAPI.md#DescribeServiceHook) | **Post** /?action&#x3D;DescribeServiceHook | Service Hook 查询单条
[**DescribeServiceHookLogs**](ServiceHookAPI.md#DescribeServiceHookLogs) | **Post** /?action&#x3D;DescribeServiceHookLogs | Service Hook 发送记录分页查询
[**DescribeServiceHooks**](ServiceHookAPI.md#DescribeServiceHooks) | **Post** /?action&#x3D;DescribeServiceHooks | Service Hook 列表分页查询
[**EnabledServiceHook**](ServiceHookAPI.md#EnabledServiceHook) | **Post** /?action&#x3D;EnabledServiceHook | Service Hook 事件开关
[**ModifyServiceHook**](ServiceHookAPI.md#ModifyServiceHook) | **Post** /?action&#x3D;ModifyServiceHook | Service Hook 修改
[**PingServiceHook**](ServiceHookAPI.md#PingServiceHook) | **Post** /?action&#x3D;PingServiceHook | Service Hook 测试
[**ResendServiceHookLog**](ServiceHookAPI.md#ResendServiceHookLog) | **Post** /?action&#x3D;ResendServiceHookLog | Service Hook 重发



## CreateServiceHook

> CreateServiceHook200Response CreateServiceHook(ctx).Authorization(authorization).Action(action).CreateServiceHookRequest(createServiceHookRequest).Execute()

Service Hook 创建



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
	action := "CreateServiceHook" // string | Action (default to "CreateServiceHook")
	createServiceHookRequest := *openapiclient.NewCreateServiceHookRequest("ActionProperty_example", false, []string{"Event_example"}, int64(123), "Service_example", "ServiceAction_example") // CreateServiceHookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.CreateServiceHook(context.Background()).Authorization(authorization).Action(action).CreateServiceHookRequest(createServiceHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.CreateServiceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceHook`: CreateServiceHook200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.CreateServiceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;CreateServiceHook&quot;]
 **createServiceHookRequest** | [**CreateServiceHookRequest**](CreateServiceHookRequest.md) |  | 

### Return type

[**CreateServiceHook200Response**](CreateServiceHook200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceHook

> EnabledServiceHook200Response DeleteServiceHook(ctx).Authorization(authorization).Action(action).DeleteServiceHookRequest(deleteServiceHookRequest).Execute()

Service Hook 删除



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
	action := "DeleteServiceHook" // string | Action (default to "DeleteServiceHook")
	deleteServiceHookRequest := *openapiclient.NewDeleteServiceHookRequest([]string{"Id_example"}, int64(123)) // DeleteServiceHookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.DeleteServiceHook(context.Background()).Authorization(authorization).Action(action).DeleteServiceHookRequest(deleteServiceHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.DeleteServiceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteServiceHook`: EnabledServiceHook200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.DeleteServiceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DeleteServiceHook&quot;]
 **deleteServiceHookRequest** | [**DeleteServiceHookRequest**](DeleteServiceHookRequest.md) |  | 

### Return type

[**EnabledServiceHook200Response**](EnabledServiceHook200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeEvents

> DescribeEvents200Response DescribeEvents(ctx).Authorization(authorization).Action(action).Body(body).Execute()

Service Hook 事件列表查询



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
	action := "DescribeEvents" // string | Action (default to "DescribeEvents")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.DescribeEvents(context.Background()).Authorization(authorization).Action(action).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.DescribeEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeEvents`: DescribeEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.DescribeEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeEvents&quot;]
 **body** | **map[string]interface{}** |  | 

### Return type

[**DescribeEvents200Response**](DescribeEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeServiceHook

> DescribeServiceHook200Response DescribeServiceHook(ctx).Authorization(authorization).Action(action).DescribeServiceHookRequest(describeServiceHookRequest).Execute()

Service Hook 查询单条



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
	action := "DescribeServiceHook" // string | Action (default to "DescribeServiceHook")
	describeServiceHookRequest := *openapiclient.NewDescribeServiceHookRequest("Id_example", int64(123)) // DescribeServiceHookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.DescribeServiceHook(context.Background()).Authorization(authorization).Action(action).DescribeServiceHookRequest(describeServiceHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.DescribeServiceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeServiceHook`: DescribeServiceHook200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.DescribeServiceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeServiceHook&quot;]
 **describeServiceHookRequest** | [**DescribeServiceHookRequest**](DescribeServiceHookRequest.md) |  | 

### Return type

[**DescribeServiceHook200Response**](DescribeServiceHook200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeServiceHookLogs

> DescribeServiceHookLogs200Response DescribeServiceHookLogs(ctx).Authorization(authorization).Action(action).DescribeServiceHookLogsRequest(describeServiceHookLogsRequest).Execute()

Service Hook 发送记录分页查询



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
	action := "DescribeServiceHookLogs" // string | Action (default to "DescribeServiceHookLogs")
	describeServiceHookLogsRequest := *openapiclient.NewDescribeServiceHookLogsRequest("Id_example", int64(123), int64(123), int64(123)) // DescribeServiceHookLogsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.DescribeServiceHookLogs(context.Background()).Authorization(authorization).Action(action).DescribeServiceHookLogsRequest(describeServiceHookLogsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.DescribeServiceHookLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeServiceHookLogs`: DescribeServiceHookLogs200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.DescribeServiceHookLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceHookLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeServiceHookLogs&quot;]
 **describeServiceHookLogsRequest** | [**DescribeServiceHookLogsRequest**](DescribeServiceHookLogsRequest.md) |  | 

### Return type

[**DescribeServiceHookLogs200Response**](DescribeServiceHookLogs200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DescribeServiceHooks

> DescribeServiceHooks200Response DescribeServiceHooks(ctx).Authorization(authorization).Action(action).DescribeServiceHooksRequest(describeServiceHooksRequest).Execute()

Service Hook 列表分页查询



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
	action := "DescribeServiceHooks" // string | Action (default to "DescribeServiceHooks")
	describeServiceHooksRequest := *openapiclient.NewDescribeServiceHooksRequest(int64(123), int64(123), int64(123), "TargetType_example") // DescribeServiceHooksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.DescribeServiceHooks(context.Background()).Authorization(authorization).Action(action).DescribeServiceHooksRequest(describeServiceHooksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.DescribeServiceHooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DescribeServiceHooks`: DescribeServiceHooks200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.DescribeServiceHooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDescribeServiceHooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;DescribeServiceHooks&quot;]
 **describeServiceHooksRequest** | [**DescribeServiceHooksRequest**](DescribeServiceHooksRequest.md) |  | 

### Return type

[**DescribeServiceHooks200Response**](DescribeServiceHooks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnabledServiceHook

> EnabledServiceHook200Response EnabledServiceHook(ctx).Authorization(authorization).Action(action).EnabledServiceHookRequest(enabledServiceHookRequest).Execute()

Service Hook 事件开关



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
	action := "EnabledServiceHook" // string | Action (default to "EnabledServiceHook")
	enabledServiceHookRequest := *openapiclient.NewEnabledServiceHookRequest(false, []string{"Id_example"}, int64(123), "TargetType_example") // EnabledServiceHookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.EnabledServiceHook(context.Background()).Authorization(authorization).Action(action).EnabledServiceHookRequest(enabledServiceHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.EnabledServiceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnabledServiceHook`: EnabledServiceHook200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.EnabledServiceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnabledServiceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;EnabledServiceHook&quot;]
 **enabledServiceHookRequest** | [**EnabledServiceHookRequest**](EnabledServiceHookRequest.md) |  | 

### Return type

[**EnabledServiceHook200Response**](EnabledServiceHook200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyServiceHook

> ModifyServiceHook200Response ModifyServiceHook(ctx).Authorization(authorization).Action(action).ModifyServiceHookRequest(modifyServiceHookRequest).Execute()

Service Hook 修改



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
	action := "ModifyServiceHook" // string | Action (default to "ModifyServiceHook")
	modifyServiceHookRequest := *openapiclient.NewModifyServiceHookRequest("ActionProperty_example", false, []string{"Event_example"}, "FilterProperty_example", "Id_example", int64(123), "Service_example", "ServiceAction_example") // ModifyServiceHookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.ModifyServiceHook(context.Background()).Authorization(authorization).Action(action).ModifyServiceHookRequest(modifyServiceHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.ModifyServiceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyServiceHook`: ModifyServiceHook200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.ModifyServiceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiModifyServiceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ModifyServiceHook&quot;]
 **modifyServiceHookRequest** | [**ModifyServiceHookRequest**](ModifyServiceHookRequest.md) |  | 

### Return type

[**ModifyServiceHook200Response**](ModifyServiceHook200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PingServiceHook

> EnabledServiceHook200Response PingServiceHook(ctx).Authorization(authorization).Action(action).PingServiceHookRequest(pingServiceHookRequest).Execute()

Service Hook 测试



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
	action := "PingServiceHook" // string | Action (default to "PingServiceHook")
	pingServiceHookRequest := *openapiclient.NewPingServiceHookRequest([]string{"Id_example"}, int64(123)) // PingServiceHookRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.PingServiceHook(context.Background()).Authorization(authorization).Action(action).PingServiceHookRequest(pingServiceHookRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.PingServiceHook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PingServiceHook`: EnabledServiceHook200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.PingServiceHook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPingServiceHookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;PingServiceHook&quot;]
 **pingServiceHookRequest** | [**PingServiceHookRequest**](PingServiceHookRequest.md) |  | 

### Return type

[**EnabledServiceHook200Response**](EnabledServiceHook200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendServiceHookLog

> ResendServiceHookLog200Response ResendServiceHookLog(ctx).Authorization(authorization).Action(action).ResendServiceHookLogRequest(resendServiceHookLogRequest).Execute()

Service Hook 重发



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
	action := "ResendServiceHookLog" // string | Action (default to "ResendServiceHookLog")
	resendServiceHookLogRequest := *openapiclient.NewResendServiceHookLogRequest("ServiceHookLogId_example") // ResendServiceHookLogRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceHookAPI.ResendServiceHookLog(context.Background()).Authorization(authorization).Action(action).ResendServiceHookLogRequest(resendServiceHookLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceHookAPI.ResendServiceHookLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResendServiceHookLog`: ResendServiceHookLog200Response
	fmt.Fprintf(os.Stdout, "Response from `ServiceHookAPI.ResendServiceHookLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResendServiceHookLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | 认证信息 | 
 **action** | **string** | Action | [default to &quot;ResendServiceHookLog&quot;]
 **resendServiceHookLogRequest** | [**ResendServiceHookLogRequest**](ResendServiceHookLogRequest.md) |  | 

### Return type

[**ResendServiceHookLog200Response**](ResendServiceHookLog200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

