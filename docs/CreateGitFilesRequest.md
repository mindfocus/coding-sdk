# CreateGitFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**GitFiles** | [**[]GitFile**](GitFile.md) | 新增的文件 | 
**LastCommitSha** | **string** | 最后次提交 Sha | 
**Message** | **string** | 提交文本 | 
**NewRef** | Pointer to **string** | 新分支 | [optional] 
**Ref** | **string** | 基于改动的分支 | 

## Methods

### NewCreateGitFilesRequest

`func NewCreateGitFilesRequest(depotId int64, gitFiles []GitFile, lastCommitSha string, message string, ref string, ) *CreateGitFilesRequest`

NewCreateGitFilesRequest instantiates a new CreateGitFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitFilesRequestWithDefaults

`func NewCreateGitFilesRequestWithDefaults() *CreateGitFilesRequest`

NewCreateGitFilesRequestWithDefaults instantiates a new CreateGitFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateGitFilesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitFilesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitFilesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitFilesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitFilesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitFilesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitFilesRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetGitFiles

`func (o *CreateGitFilesRequest) GetGitFiles() []GitFile`

GetGitFiles returns the GitFiles field if non-nil, zero value otherwise.

### GetGitFilesOk

`func (o *CreateGitFilesRequest) GetGitFilesOk() (*[]GitFile, bool)`

GetGitFilesOk returns a tuple with the GitFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitFiles

`func (o *CreateGitFilesRequest) SetGitFiles(v []GitFile)`

SetGitFiles sets GitFiles field to given value.


### GetLastCommitSha

`func (o *CreateGitFilesRequest) GetLastCommitSha() string`

GetLastCommitSha returns the LastCommitSha field if non-nil, zero value otherwise.

### GetLastCommitShaOk

`func (o *CreateGitFilesRequest) GetLastCommitShaOk() (*string, bool)`

GetLastCommitShaOk returns a tuple with the LastCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitSha

`func (o *CreateGitFilesRequest) SetLastCommitSha(v string)`

SetLastCommitSha sets LastCommitSha field to given value.


### GetMessage

`func (o *CreateGitFilesRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateGitFilesRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateGitFilesRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNewRef

`func (o *CreateGitFilesRequest) GetNewRef() string`

GetNewRef returns the NewRef field if non-nil, zero value otherwise.

### GetNewRefOk

`func (o *CreateGitFilesRequest) GetNewRefOk() (*string, bool)`

GetNewRefOk returns a tuple with the NewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRef

`func (o *CreateGitFilesRequest) SetNewRef(v string)`

SetNewRef sets NewRef field to given value.

### HasNewRef

`func (o *CreateGitFilesRequest) HasNewRef() bool`

HasNewRef returns a boolean if a field has been set.

### GetRef

`func (o *CreateGitFilesRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CreateGitFilesRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CreateGitFilesRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


