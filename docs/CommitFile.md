# CommitFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | 文件内容 | [optional] [default to ""]
**ContentEncoding** | **string** | ENCODING_RAW：不对文件编码；ENCODING_BASE64：传入的文件内容使用 base64 编码 | [default to ""]
**NewPath** | Pointer to **string** | 新建文件、重命名的文件路径 | [optional] [default to ""]
**Path** | **string** | 文件路径 | [default to ""]
**WantDelete** | **bool** | 是否删除 | [default to false]
**WantRename** | **bool** | 是否重命名 | [default to false]

## Methods

### NewCommitFile

`func NewCommitFile(contentEncoding string, path string, wantDelete bool, wantRename bool, ) *CommitFile`

NewCommitFile instantiates a new CommitFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitFileWithDefaults

`func NewCommitFileWithDefaults() *CommitFile`

NewCommitFileWithDefaults instantiates a new CommitFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CommitFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CommitFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CommitFile) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CommitFile) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentEncoding

`func (o *CommitFile) GetContentEncoding() string`

GetContentEncoding returns the ContentEncoding field if non-nil, zero value otherwise.

### GetContentEncodingOk

`func (o *CommitFile) GetContentEncodingOk() (*string, bool)`

GetContentEncodingOk returns a tuple with the ContentEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEncoding

`func (o *CommitFile) SetContentEncoding(v string)`

SetContentEncoding sets ContentEncoding field to given value.


### GetNewPath

`func (o *CommitFile) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *CommitFile) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *CommitFile) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *CommitFile) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### GetPath

`func (o *CommitFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CommitFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CommitFile) SetPath(v string)`

SetPath sets Path field to given value.


### GetWantDelete

`func (o *CommitFile) GetWantDelete() bool`

GetWantDelete returns the WantDelete field if non-nil, zero value otherwise.

### GetWantDeleteOk

`func (o *CommitFile) GetWantDeleteOk() (*bool, bool)`

GetWantDeleteOk returns a tuple with the WantDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantDelete

`func (o *CommitFile) SetWantDelete(v bool)`

SetWantDelete sets WantDelete field to given value.


### GetWantRename

`func (o *CommitFile) GetWantRename() bool`

GetWantRename returns the WantRename field if non-nil, zero value otherwise.

### GetWantRenameOk

`func (o *CommitFile) GetWantRenameOk() (*bool, bool)`

GetWantRenameOk returns a tuple with the WantRename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantRename

`func (o *CommitFile) SetWantRename(v bool)`

SetWantRename sets WantRename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


