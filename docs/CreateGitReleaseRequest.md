# CreateGitReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | 提交 Sha 值 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**Description** | **string** | 版本描述信息 | 
**Pre** | **bool** | 是否为预发布版本 | 
**TagName** | **string** | 创建来源的分支名称或者commitId | 
**TargetCommitBranch** | **string** | 目标提交 Sha 值 | 
**Title** | **string** | 版本标题 | 
**UserId** | **int64** | 分支名称 | 

## Methods

### NewCreateGitReleaseRequest

`func NewCreateGitReleaseRequest(commitSha string, depotId int64, description string, pre bool, tagName string, targetCommitBranch string, title string, userId int64, ) *CreateGitReleaseRequest`

NewCreateGitReleaseRequest instantiates a new CreateGitReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitReleaseRequestWithDefaults

`func NewCreateGitReleaseRequestWithDefaults() *CreateGitReleaseRequest`

NewCreateGitReleaseRequestWithDefaults instantiates a new CreateGitReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *CreateGitReleaseRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CreateGitReleaseRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CreateGitReleaseRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetDepotId

`func (o *CreateGitReleaseRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitReleaseRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitReleaseRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitReleaseRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitReleaseRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitReleaseRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitReleaseRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDescription

`func (o *CreateGitReleaseRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGitReleaseRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGitReleaseRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPre

`func (o *CreateGitReleaseRequest) GetPre() bool`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *CreateGitReleaseRequest) GetPreOk() (*bool, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *CreateGitReleaseRequest) SetPre(v bool)`

SetPre sets Pre field to given value.


### GetTagName

`func (o *CreateGitReleaseRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *CreateGitReleaseRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *CreateGitReleaseRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTargetCommitBranch

`func (o *CreateGitReleaseRequest) GetTargetCommitBranch() string`

GetTargetCommitBranch returns the TargetCommitBranch field if non-nil, zero value otherwise.

### GetTargetCommitBranchOk

`func (o *CreateGitReleaseRequest) GetTargetCommitBranchOk() (*string, bool)`

GetTargetCommitBranchOk returns a tuple with the TargetCommitBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCommitBranch

`func (o *CreateGitReleaseRequest) SetTargetCommitBranch(v string)`

SetTargetCommitBranch sets TargetCommitBranch field to given value.


### GetTitle

`func (o *CreateGitReleaseRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGitReleaseRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGitReleaseRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUserId

`func (o *CreateGitReleaseRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateGitReleaseRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateGitReleaseRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


