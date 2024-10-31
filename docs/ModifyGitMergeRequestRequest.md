# ModifyGitMergeRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | 待修改的合并请求描述 | [optional] 
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**MergeId** | **int64** | 合并请求 IID | 
**Title** | Pointer to **string** | 待修改的合并请求标题 | [optional] 

## Methods

### NewModifyGitMergeRequestRequest

`func NewModifyGitMergeRequestRequest(depotId int64, mergeId int64, ) *ModifyGitMergeRequestRequest`

NewModifyGitMergeRequestRequest instantiates a new ModifyGitMergeRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitMergeRequestRequestWithDefaults

`func NewModifyGitMergeRequestRequestWithDefaults() *ModifyGitMergeRequestRequest`

NewModifyGitMergeRequestRequestWithDefaults instantiates a new ModifyGitMergeRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ModifyGitMergeRequestRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModifyGitMergeRequestRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModifyGitMergeRequestRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ModifyGitMergeRequestRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDepotId

`func (o *ModifyGitMergeRequestRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitMergeRequestRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitMergeRequestRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitMergeRequestRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitMergeRequestRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitMergeRequestRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitMergeRequestRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMergeId

`func (o *ModifyGitMergeRequestRequest) GetMergeId() int64`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *ModifyGitMergeRequestRequest) GetMergeIdOk() (*int64, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *ModifyGitMergeRequestRequest) SetMergeId(v int64)`

SetMergeId sets MergeId field to given value.


### GetTitle

`func (o *ModifyGitMergeRequestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModifyGitMergeRequestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModifyGitMergeRequestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModifyGitMergeRequestRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


