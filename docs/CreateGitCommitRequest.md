# CreateGitCommitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitFiles** | [**[]CommitFile**](CommitFile.md) | 文件列表 | 
**DepotPath** | **string** | 仓库路径 | 
**LastCommitSha** | **string** | 最后次提交 Sha | 
**Message** | **string** | 提交文本 | 
**NewRef** | Pointer to **string** | 新分支 | [optional] 
**Ref** | **string** | 基于改动的分支 | 

## Methods

### NewCreateGitCommitRequest

`func NewCreateGitCommitRequest(commitFiles []CommitFile, depotPath string, lastCommitSha string, message string, ref string, ) *CreateGitCommitRequest`

NewCreateGitCommitRequest instantiates a new CreateGitCommitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitCommitRequestWithDefaults

`func NewCreateGitCommitRequestWithDefaults() *CreateGitCommitRequest`

NewCreateGitCommitRequestWithDefaults instantiates a new CreateGitCommitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitFiles

`func (o *CreateGitCommitRequest) GetCommitFiles() []CommitFile`

GetCommitFiles returns the CommitFiles field if non-nil, zero value otherwise.

### GetCommitFilesOk

`func (o *CreateGitCommitRequest) GetCommitFilesOk() (*[]CommitFile, bool)`

GetCommitFilesOk returns a tuple with the CommitFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitFiles

`func (o *CreateGitCommitRequest) SetCommitFiles(v []CommitFile)`

SetCommitFiles sets CommitFiles field to given value.


### GetDepotPath

`func (o *CreateGitCommitRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitCommitRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitCommitRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetLastCommitSha

`func (o *CreateGitCommitRequest) GetLastCommitSha() string`

GetLastCommitSha returns the LastCommitSha field if non-nil, zero value otherwise.

### GetLastCommitShaOk

`func (o *CreateGitCommitRequest) GetLastCommitShaOk() (*string, bool)`

GetLastCommitShaOk returns a tuple with the LastCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitSha

`func (o *CreateGitCommitRequest) SetLastCommitSha(v string)`

SetLastCommitSha sets LastCommitSha field to given value.


### GetMessage

`func (o *CreateGitCommitRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateGitCommitRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateGitCommitRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNewRef

`func (o *CreateGitCommitRequest) GetNewRef() string`

GetNewRef returns the NewRef field if non-nil, zero value otherwise.

### GetNewRefOk

`func (o *CreateGitCommitRequest) GetNewRefOk() (*string, bool)`

GetNewRefOk returns a tuple with the NewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRef

`func (o *CreateGitCommitRequest) SetNewRef(v string)`

SetNewRef sets NewRef field to given value.

### HasNewRef

`func (o *CreateGitCommitRequest) HasNewRef() bool`

HasNewRef returns a boolean if a field has been set.

### GetRef

`func (o *CreateGitCommitRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CreateGitCommitRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CreateGitCommitRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


