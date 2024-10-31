# CreateAPIDocRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | API 文档内容 | 
**ContentType** | **string** | API 文档内容类型，SWAGGER：swagger 格式，APIDOCS：apidoc 格式，POSTMAN：postman 格式，TENCENTCLOUD_API_GATEWAY: 腾讯云 API 网关接口格式 | 
**PrefixUri** | **string** | API 文档域名前缀，以小写字母 a 到 z 或者 0 到 9 开头，中间可以出现中横线 。最长不能超过32位 | 
**ProjectName** | **string** | 项目名称 | 
**SharePassword** | Pointer to **string** | API 文档密码，6-8 位英文和字母 | [optional] 
**Title** | **string** | API 文档标题 | 

## Methods

### NewCreateAPIDocRequest

`func NewCreateAPIDocRequest(content string, contentType string, prefixUri string, projectName string, title string, ) *CreateAPIDocRequest`

NewCreateAPIDocRequest instantiates a new CreateAPIDocRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAPIDocRequestWithDefaults

`func NewCreateAPIDocRequestWithDefaults() *CreateAPIDocRequest`

NewCreateAPIDocRequestWithDefaults instantiates a new CreateAPIDocRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateAPIDocRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateAPIDocRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateAPIDocRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetContentType

`func (o *CreateAPIDocRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CreateAPIDocRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CreateAPIDocRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetPrefixUri

`func (o *CreateAPIDocRequest) GetPrefixUri() string`

GetPrefixUri returns the PrefixUri field if non-nil, zero value otherwise.

### GetPrefixUriOk

`func (o *CreateAPIDocRequest) GetPrefixUriOk() (*string, bool)`

GetPrefixUriOk returns a tuple with the PrefixUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixUri

`func (o *CreateAPIDocRequest) SetPrefixUri(v string)`

SetPrefixUri sets PrefixUri field to given value.


### GetProjectName

`func (o *CreateAPIDocRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateAPIDocRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateAPIDocRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSharePassword

`func (o *CreateAPIDocRequest) GetSharePassword() string`

GetSharePassword returns the SharePassword field if non-nil, zero value otherwise.

### GetSharePasswordOk

`func (o *CreateAPIDocRequest) GetSharePasswordOk() (*string, bool)`

GetSharePasswordOk returns a tuple with the SharePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePassword

`func (o *CreateAPIDocRequest) SetSharePassword(v string)`

SetSharePassword sets SharePassword field to given value.

### HasSharePassword

`func (o *CreateAPIDocRequest) HasSharePassword() bool`

HasSharePassword returns a boolean if a field has been set.

### GetTitle

`func (o *CreateAPIDocRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateAPIDocRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateAPIDocRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


