# GitMergeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | Pointer to **int32** | 代码仓库的唯一编号 | [optional] 
**DepotPath** | Pointer to **NullableString** | 仓库路径 | [optional] [default to ""]
**IId** | Pointer to **int32** | 定位一个项目的内的资源的 ID | [optional] 
**MergeId** | Pointer to **int32** | 定位一个项目的内的MR资源的 ID | [optional] 
**ProjectId** | Pointer to **int32** | 项目 ID | [optional] 

## Methods

### NewGitMergeRequest

`func NewGitMergeRequest() *GitMergeRequest`

NewGitMergeRequest instantiates a new GitMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitMergeRequestWithDefaults

`func NewGitMergeRequestWithDefaults() *GitMergeRequest`

NewGitMergeRequestWithDefaults instantiates a new GitMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *GitMergeRequest) GetDepotId() int32`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *GitMergeRequest) GetDepotIdOk() (*int32, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *GitMergeRequest) SetDepotId(v int32)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *GitMergeRequest) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetDepotPath

`func (o *GitMergeRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *GitMergeRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *GitMergeRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *GitMergeRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### SetDepotPathNil

`func (o *GitMergeRequest) SetDepotPathNil(b bool)`

 SetDepotPathNil sets the value for DepotPath to be an explicit nil

### UnsetDepotPath
`func (o *GitMergeRequest) UnsetDepotPath()`

UnsetDepotPath ensures that no value is present for DepotPath, not even an explicit nil
### GetIId

`func (o *GitMergeRequest) GetIId() int32`

GetIId returns the IId field if non-nil, zero value otherwise.

### GetIIdOk

`func (o *GitMergeRequest) GetIIdOk() (*int32, bool)`

GetIIdOk returns a tuple with the IId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIId

`func (o *GitMergeRequest) SetIId(v int32)`

SetIId sets IId field to given value.

### HasIId

`func (o *GitMergeRequest) HasIId() bool`

HasIId returns a boolean if a field has been set.

### GetMergeId

`func (o *GitMergeRequest) GetMergeId() int32`

GetMergeId returns the MergeId field if non-nil, zero value otherwise.

### GetMergeIdOk

`func (o *GitMergeRequest) GetMergeIdOk() (*int32, bool)`

GetMergeIdOk returns a tuple with the MergeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeId

`func (o *GitMergeRequest) SetMergeId(v int32)`

SetMergeId sets MergeId field to given value.

### HasMergeId

`func (o *GitMergeRequest) HasMergeId() bool`

HasMergeId returns a boolean if a field has been set.

### GetProjectId

`func (o *GitMergeRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GitMergeRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GitMergeRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GitMergeRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


