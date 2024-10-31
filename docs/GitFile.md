# GitFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 文件内容 | [default to ""]
**NewPath** | Pointer to **string** | 是否创建一个新路径文件 | [optional] [default to ""]
**Path** | **string** | 文件路径 | [default to ""]

## Methods

### NewGitFile

`func NewGitFile(content string, path string, ) *GitFile`

NewGitFile instantiates a new GitFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFileWithDefaults

`func NewGitFileWithDefaults() *GitFile`

NewGitFileWithDefaults instantiates a new GitFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GitFile) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitFile) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitFile) SetContent(v string)`

SetContent sets Content field to given value.


### GetNewPath

`func (o *GitFile) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *GitFile) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *GitFile) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *GitFile) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### GetPath

`func (o *GitFile) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitFile) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitFile) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


