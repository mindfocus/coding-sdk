# CreateGitMergeRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 描述内容 | 
**DepotId** | **int32** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一   | [optional] 
**DestBranch** | **string** | 目标分支名 | 
**SrcBranch** | **string** | 源分支名 | 
**Title** | **string** | 标题 | 

## Methods

### NewCreateGitMergeRequestRequest

`func NewCreateGitMergeRequestRequest(content string, depotId int32, destBranch string, srcBranch string, title string, ) *CreateGitMergeRequestRequest`

NewCreateGitMergeRequestRequest instantiates a new CreateGitMergeRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitMergeRequestRequestWithDefaults

`func NewCreateGitMergeRequestRequestWithDefaults() *CreateGitMergeRequestRequest`

NewCreateGitMergeRequestRequestWithDefaults instantiates a new CreateGitMergeRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateGitMergeRequestRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateGitMergeRequestRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateGitMergeRequestRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetDepotId

`func (o *CreateGitMergeRequestRequest) GetDepotId() int32`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitMergeRequestRequest) GetDepotIdOk() (*int32, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitMergeRequestRequest) SetDepotId(v int32)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitMergeRequestRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitMergeRequestRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitMergeRequestRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitMergeRequestRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDestBranch

`func (o *CreateGitMergeRequestRequest) GetDestBranch() string`

GetDestBranch returns the DestBranch field if non-nil, zero value otherwise.

### GetDestBranchOk

`func (o *CreateGitMergeRequestRequest) GetDestBranchOk() (*string, bool)`

GetDestBranchOk returns a tuple with the DestBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestBranch

`func (o *CreateGitMergeRequestRequest) SetDestBranch(v string)`

SetDestBranch sets DestBranch field to given value.


### GetSrcBranch

`func (o *CreateGitMergeRequestRequest) GetSrcBranch() string`

GetSrcBranch returns the SrcBranch field if non-nil, zero value otherwise.

### GetSrcBranchOk

`func (o *CreateGitMergeRequestRequest) GetSrcBranchOk() (*string, bool)`

GetSrcBranchOk returns a tuple with the SrcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBranch

`func (o *CreateGitMergeRequestRequest) SetSrcBranch(v string)`

SetSrcBranch sets SrcBranch field to given value.


### GetTitle

`func (o *CreateGitMergeRequestRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGitMergeRequestRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGitMergeRequestRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


