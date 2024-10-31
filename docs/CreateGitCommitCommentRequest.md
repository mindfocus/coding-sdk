# CreateGitCommitCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 评论的内容 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Index** | **int64** | diff信息中的第几行 | 
**Path** | **string** | 文件路径 | 
**Sha** | **string** | 提交id | 

## Methods

### NewCreateGitCommitCommentRequest

`func NewCreateGitCommitCommentRequest(content string, depotId int64, index int64, path string, sha string, ) *CreateGitCommitCommentRequest`

NewCreateGitCommitCommentRequest instantiates a new CreateGitCommitCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitCommitCommentRequestWithDefaults

`func NewCreateGitCommitCommentRequestWithDefaults() *CreateGitCommitCommentRequest`

NewCreateGitCommitCommentRequestWithDefaults instantiates a new CreateGitCommitCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateGitCommitCommentRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateGitCommitCommentRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateGitCommitCommentRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetDepotId

`func (o *CreateGitCommitCommentRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitCommitCommentRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitCommitCommentRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitCommitCommentRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitCommitCommentRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitCommitCommentRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitCommitCommentRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetIndex

`func (o *CreateGitCommitCommentRequest) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CreateGitCommitCommentRequest) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CreateGitCommitCommentRequest) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetPath

`func (o *CreateGitCommitCommentRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CreateGitCommitCommentRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CreateGitCommitCommentRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetSha

`func (o *CreateGitCommitCommentRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CreateGitCommitCommentRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CreateGitCommitCommentRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


