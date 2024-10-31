# ModifyGitReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**Description** | Pointer to **string** | 描述 | [optional] 
**Pre** | Pointer to **bool** | 是否预发布 | [optional] 
**ReleaseId** | **int64** | 项目下仓库版本唯一标识符 | 
**TagName** | **string** | 标签名称 | 
**Title** | Pointer to **string** | 标题 | [optional] 

## Methods

### NewModifyGitReleaseRequest

`func NewModifyGitReleaseRequest(depotId int64, releaseId int64, tagName string, ) *ModifyGitReleaseRequest`

NewModifyGitReleaseRequest instantiates a new ModifyGitReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitReleaseRequestWithDefaults

`func NewModifyGitReleaseRequestWithDefaults() *ModifyGitReleaseRequest`

NewModifyGitReleaseRequestWithDefaults instantiates a new ModifyGitReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyGitReleaseRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitReleaseRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitReleaseRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitReleaseRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitReleaseRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitReleaseRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitReleaseRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDescription

`func (o *ModifyGitReleaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyGitReleaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyGitReleaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyGitReleaseRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPre

`func (o *ModifyGitReleaseRequest) GetPre() bool`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *ModifyGitReleaseRequest) GetPreOk() (*bool, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *ModifyGitReleaseRequest) SetPre(v bool)`

SetPre sets Pre field to given value.

### HasPre

`func (o *ModifyGitReleaseRequest) HasPre() bool`

HasPre returns a boolean if a field has been set.

### GetReleaseId

`func (o *ModifyGitReleaseRequest) GetReleaseId() int64`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *ModifyGitReleaseRequest) GetReleaseIdOk() (*int64, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *ModifyGitReleaseRequest) SetReleaseId(v int64)`

SetReleaseId sets ReleaseId field to given value.


### GetTagName

`func (o *ModifyGitReleaseRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ModifyGitReleaseRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ModifyGitReleaseRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTitle

`func (o *ModifyGitReleaseRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModifyGitReleaseRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModifyGitReleaseRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModifyGitReleaseRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


