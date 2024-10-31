# CreateGitMergeReqRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 合并请求内容 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径与depotId选择其一就可以 | [optional] 
**DestBranch** | **string** | 目标分支 | 
**Reviewers** | Pointer to **string** | 评审者 | [optional] 
**SrcBranch** | **string** | 源分支 | 
**Title** | **string** | 合并请求title | 

## Methods

### NewCreateGitMergeReqRequest

`func NewCreateGitMergeReqRequest(content string, depotId int64, destBranch string, srcBranch string, title string, ) *CreateGitMergeReqRequest`

NewCreateGitMergeReqRequest instantiates a new CreateGitMergeReqRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitMergeReqRequestWithDefaults

`func NewCreateGitMergeReqRequestWithDefaults() *CreateGitMergeReqRequest`

NewCreateGitMergeReqRequestWithDefaults instantiates a new CreateGitMergeReqRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateGitMergeReqRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateGitMergeReqRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateGitMergeReqRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetDepotId

`func (o *CreateGitMergeReqRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitMergeReqRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitMergeReqRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitMergeReqRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitMergeReqRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitMergeReqRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitMergeReqRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDestBranch

`func (o *CreateGitMergeReqRequest) GetDestBranch() string`

GetDestBranch returns the DestBranch field if non-nil, zero value otherwise.

### GetDestBranchOk

`func (o *CreateGitMergeReqRequest) GetDestBranchOk() (*string, bool)`

GetDestBranchOk returns a tuple with the DestBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestBranch

`func (o *CreateGitMergeReqRequest) SetDestBranch(v string)`

SetDestBranch sets DestBranch field to given value.


### GetReviewers

`func (o *CreateGitMergeReqRequest) GetReviewers() string`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *CreateGitMergeReqRequest) GetReviewersOk() (*string, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *CreateGitMergeReqRequest) SetReviewers(v string)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *CreateGitMergeReqRequest) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetSrcBranch

`func (o *CreateGitMergeReqRequest) GetSrcBranch() string`

GetSrcBranch returns the SrcBranch field if non-nil, zero value otherwise.

### GetSrcBranchOk

`func (o *CreateGitMergeReqRequest) GetSrcBranchOk() (*string, bool)`

GetSrcBranchOk returns a tuple with the SrcBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcBranch

`func (o *CreateGitMergeReqRequest) SetSrcBranch(v string)`

SetSrcBranch sets SrcBranch field to given value.


### GetTitle

`func (o *CreateGitMergeReqRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGitMergeReqRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGitMergeReqRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


