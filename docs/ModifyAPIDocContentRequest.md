# ModifyAPIDocContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | API 文档编号 | 
**Content** | **string** | API 文档内容 | 
**ContentType** | **string** | API 文档内容类型，SWAGGER：swagger 格式，APIDOCS：apidoc 格式，POSTMAN：postman 格式，TENCENTCLOUD_API_GATEWAY：腾讯云 API 网关接口格式 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewModifyAPIDocContentRequest

`func NewModifyAPIDocContentRequest(code int32, content string, contentType string, projectName string, ) *ModifyAPIDocContentRequest`

NewModifyAPIDocContentRequest instantiates a new ModifyAPIDocContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyAPIDocContentRequestWithDefaults

`func NewModifyAPIDocContentRequestWithDefaults() *ModifyAPIDocContentRequest`

NewModifyAPIDocContentRequestWithDefaults instantiates a new ModifyAPIDocContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModifyAPIDocContentRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModifyAPIDocContentRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModifyAPIDocContentRequest) SetCode(v int32)`

SetCode sets Code field to given value.


### GetContent

`func (o *ModifyAPIDocContentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModifyAPIDocContentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModifyAPIDocContentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *ModifyAPIDocContentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ModifyAPIDocContentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ModifyAPIDocContentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetProjectName

`func (o *ModifyAPIDocContentRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyAPIDocContentRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyAPIDocContentRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


