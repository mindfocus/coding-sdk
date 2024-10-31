# GitFileContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | 内容 | [optional] [default to ""]
**IsLargeFileStorage** | Pointer to **bool** | 是否为lfs文件 | [optional] [default to false]
**IsLfs** | Pointer to **bool** | 是否lfs文件 | [optional] [default to false]
**IsText** | Pointer to **bool** | 是否文本 | [optional] [default to false]

## Methods

### NewGitFileContent

`func NewGitFileContent() *GitFileContent`

NewGitFileContent instantiates a new GitFileContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFileContentWithDefaults

`func NewGitFileContentWithDefaults() *GitFileContent`

NewGitFileContentWithDefaults instantiates a new GitFileContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GitFileContent) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitFileContent) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitFileContent) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitFileContent) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetIsLargeFileStorage

`func (o *GitFileContent) GetIsLargeFileStorage() bool`

GetIsLargeFileStorage returns the IsLargeFileStorage field if non-nil, zero value otherwise.

### GetIsLargeFileStorageOk

`func (o *GitFileContent) GetIsLargeFileStorageOk() (*bool, bool)`

GetIsLargeFileStorageOk returns a tuple with the IsLargeFileStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLargeFileStorage

`func (o *GitFileContent) SetIsLargeFileStorage(v bool)`

SetIsLargeFileStorage sets IsLargeFileStorage field to given value.

### HasIsLargeFileStorage

`func (o *GitFileContent) HasIsLargeFileStorage() bool`

HasIsLargeFileStorage returns a boolean if a field has been set.

### GetIsLfs

`func (o *GitFileContent) GetIsLfs() bool`

GetIsLfs returns the IsLfs field if non-nil, zero value otherwise.

### GetIsLfsOk

`func (o *GitFileContent) GetIsLfsOk() (*bool, bool)`

GetIsLfsOk returns a tuple with the IsLfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLfs

`func (o *GitFileContent) SetIsLfs(v bool)`

SetIsLfs sets IsLfs field to given value.

### HasIsLfs

`func (o *GitFileContent) HasIsLfs() bool`

HasIsLfs returns a boolean if a field has been set.

### GetIsText

`func (o *GitFileContent) GetIsText() bool`

GetIsText returns the IsText field if non-nil, zero value otherwise.

### GetIsTextOk

`func (o *GitFileContent) GetIsTextOk() (*bool, bool)`

GetIsTextOk returns a tuple with the IsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsText

`func (o *GitFileContent) SetIsText(v bool)`

SetIsText sets IsText field to given value.

### HasIsText

`func (o *GitFileContent) HasIsText() bool`

HasIsText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


