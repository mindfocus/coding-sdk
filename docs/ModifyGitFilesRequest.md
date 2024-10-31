# ModifyGitFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**GitFiles** | [**[]GitFile**](GitFile.md) | 修改的文件 | 
**LastCommitSha** | **string** | 最后次提交 Sha | 
**Message** | **string** | 提交文本 | 
**NewRef** | Pointer to **string** | 要提交的新分支 | [optional] 
**Ref** | **string** | 基于改动的分支 | 

## Methods

### NewModifyGitFilesRequest

`func NewModifyGitFilesRequest(depotId int64, gitFiles []GitFile, lastCommitSha string, message string, ref string, ) *ModifyGitFilesRequest`

NewModifyGitFilesRequest instantiates a new ModifyGitFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitFilesRequestWithDefaults

`func NewModifyGitFilesRequestWithDefaults() *ModifyGitFilesRequest`

NewModifyGitFilesRequestWithDefaults instantiates a new ModifyGitFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyGitFilesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitFilesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitFilesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitFilesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitFilesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitFilesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitFilesRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetGitFiles

`func (o *ModifyGitFilesRequest) GetGitFiles() []GitFile`

GetGitFiles returns the GitFiles field if non-nil, zero value otherwise.

### GetGitFilesOk

`func (o *ModifyGitFilesRequest) GetGitFilesOk() (*[]GitFile, bool)`

GetGitFilesOk returns a tuple with the GitFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitFiles

`func (o *ModifyGitFilesRequest) SetGitFiles(v []GitFile)`

SetGitFiles sets GitFiles field to given value.


### GetLastCommitSha

`func (o *ModifyGitFilesRequest) GetLastCommitSha() string`

GetLastCommitSha returns the LastCommitSha field if non-nil, zero value otherwise.

### GetLastCommitShaOk

`func (o *ModifyGitFilesRequest) GetLastCommitShaOk() (*string, bool)`

GetLastCommitShaOk returns a tuple with the LastCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitSha

`func (o *ModifyGitFilesRequest) SetLastCommitSha(v string)`

SetLastCommitSha sets LastCommitSha field to given value.


### GetMessage

`func (o *ModifyGitFilesRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModifyGitFilesRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModifyGitFilesRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNewRef

`func (o *ModifyGitFilesRequest) GetNewRef() string`

GetNewRef returns the NewRef field if non-nil, zero value otherwise.

### GetNewRefOk

`func (o *ModifyGitFilesRequest) GetNewRefOk() (*string, bool)`

GetNewRefOk returns a tuple with the NewRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRef

`func (o *ModifyGitFilesRequest) SetNewRef(v string)`

SetNewRef sets NewRef field to given value.

### HasNewRef

`func (o *ModifyGitFilesRequest) HasNewRef() bool`

HasNewRef returns a boolean if a field has been set.

### GetRef

`func (o *ModifyGitFilesRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ModifyGitFilesRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ModifyGitFilesRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


