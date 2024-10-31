# ModifyAPIDocBaseInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | API 文档编号 | 
**PrefixUri** | Pointer to **string** | API 文档域名前缀，以小写字母 a 到 z 或者 0 到 9 开头，中间可以出现中横线 。最长不能超过32位 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SharePassword** | Pointer to **string** | API 文档密码，6-8 位英文和字母 | [optional] 
**Title** | Pointer to **string** | API 文档标题 | [optional] 

## Methods

### NewModifyAPIDocBaseInfoRequest

`func NewModifyAPIDocBaseInfoRequest(code int32, projectName string, ) *ModifyAPIDocBaseInfoRequest`

NewModifyAPIDocBaseInfoRequest instantiates a new ModifyAPIDocBaseInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyAPIDocBaseInfoRequestWithDefaults

`func NewModifyAPIDocBaseInfoRequestWithDefaults() *ModifyAPIDocBaseInfoRequest`

NewModifyAPIDocBaseInfoRequestWithDefaults instantiates a new ModifyAPIDocBaseInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ModifyAPIDocBaseInfoRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ModifyAPIDocBaseInfoRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ModifyAPIDocBaseInfoRequest) SetCode(v int32)`

SetCode sets Code field to given value.


### GetPrefixUri

`func (o *ModifyAPIDocBaseInfoRequest) GetPrefixUri() string`

GetPrefixUri returns the PrefixUri field if non-nil, zero value otherwise.

### GetPrefixUriOk

`func (o *ModifyAPIDocBaseInfoRequest) GetPrefixUriOk() (*string, bool)`

GetPrefixUriOk returns a tuple with the PrefixUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixUri

`func (o *ModifyAPIDocBaseInfoRequest) SetPrefixUri(v string)`

SetPrefixUri sets PrefixUri field to given value.

### HasPrefixUri

`func (o *ModifyAPIDocBaseInfoRequest) HasPrefixUri() bool`

HasPrefixUri returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyAPIDocBaseInfoRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyAPIDocBaseInfoRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyAPIDocBaseInfoRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSharePassword

`func (o *ModifyAPIDocBaseInfoRequest) GetSharePassword() string`

GetSharePassword returns the SharePassword field if non-nil, zero value otherwise.

### GetSharePasswordOk

`func (o *ModifyAPIDocBaseInfoRequest) GetSharePasswordOk() (*string, bool)`

GetSharePasswordOk returns a tuple with the SharePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePassword

`func (o *ModifyAPIDocBaseInfoRequest) SetSharePassword(v string)`

SetSharePassword sets SharePassword field to given value.

### HasSharePassword

`func (o *ModifyAPIDocBaseInfoRequest) HasSharePassword() bool`

HasSharePassword returns a boolean if a field has been set.

### GetTitle

`func (o *ModifyAPIDocBaseInfoRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModifyAPIDocBaseInfoRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModifyAPIDocBaseInfoRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModifyAPIDocBaseInfoRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


