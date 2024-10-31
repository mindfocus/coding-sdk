# CreateBinaryFilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**DestRef** | Pointer to **string** | 目标分支名，如果想上传至源分支，值与源分支名一致 | [optional] 
**GitFiles** | [**[]GitFile**](GitFile.md) | 文件信息，其中Path需带上文件名，例如：/data/ReadMe.md。  Content为文件字节流进行Base64编码后的字符串。 | 
**LastCommitSha** | **string** | 源分支最后一次提交sha | 
**Message** | **string** | 提交信息 | 
**SrcRef** | **string** | 源分支名 | 
**UserId** | Pointer to **int64** | 用户id | [optional] 

## Methods

### NewCreateBinaryFilesRequest

`func NewCreateBinaryFilesRequest(depotId int64, gitFiles []GitFile, lastCommitSha string, message string, srcRef string, ) *CreateBinaryFilesRequest`

NewCreateBinaryFilesRequest instantiates a new CreateBinaryFilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBinaryFilesRequestWithDefaults

`func NewCreateBinaryFilesRequestWithDefaults() *CreateBinaryFilesRequest`

NewCreateBinaryFilesRequestWithDefaults instantiates a new CreateBinaryFilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateBinaryFilesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateBinaryFilesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateBinaryFilesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateBinaryFilesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateBinaryFilesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateBinaryFilesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateBinaryFilesRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetDestRef

`func (o *CreateBinaryFilesRequest) GetDestRef() string`

GetDestRef returns the DestRef field if non-nil, zero value otherwise.

### GetDestRefOk

`func (o *CreateBinaryFilesRequest) GetDestRefOk() (*string, bool)`

GetDestRefOk returns a tuple with the DestRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestRef

`func (o *CreateBinaryFilesRequest) SetDestRef(v string)`

SetDestRef sets DestRef field to given value.

### HasDestRef

`func (o *CreateBinaryFilesRequest) HasDestRef() bool`

HasDestRef returns a boolean if a field has been set.

### GetGitFiles

`func (o *CreateBinaryFilesRequest) GetGitFiles() []GitFile`

GetGitFiles returns the GitFiles field if non-nil, zero value otherwise.

### GetGitFilesOk

`func (o *CreateBinaryFilesRequest) GetGitFilesOk() (*[]GitFile, bool)`

GetGitFilesOk returns a tuple with the GitFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitFiles

`func (o *CreateBinaryFilesRequest) SetGitFiles(v []GitFile)`

SetGitFiles sets GitFiles field to given value.


### GetLastCommitSha

`func (o *CreateBinaryFilesRequest) GetLastCommitSha() string`

GetLastCommitSha returns the LastCommitSha field if non-nil, zero value otherwise.

### GetLastCommitShaOk

`func (o *CreateBinaryFilesRequest) GetLastCommitShaOk() (*string, bool)`

GetLastCommitShaOk returns a tuple with the LastCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommitSha

`func (o *CreateBinaryFilesRequest) SetLastCommitSha(v string)`

SetLastCommitSha sets LastCommitSha field to given value.


### GetMessage

`func (o *CreateBinaryFilesRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateBinaryFilesRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateBinaryFilesRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetSrcRef

`func (o *CreateBinaryFilesRequest) GetSrcRef() string`

GetSrcRef returns the SrcRef field if non-nil, zero value otherwise.

### GetSrcRefOk

`func (o *CreateBinaryFilesRequest) GetSrcRefOk() (*string, bool)`

GetSrcRefOk returns a tuple with the SrcRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcRef

`func (o *CreateBinaryFilesRequest) SetSrcRef(v string)`

SetSrcRef sets SrcRef field to given value.


### GetUserId

`func (o *CreateBinaryFilesRequest) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateBinaryFilesRequest) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateBinaryFilesRequest) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateBinaryFilesRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


