# GitFileItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | 加密后文件的内容 | [optional] [default to ""]
**ContentSha256** | Pointer to **NullableString** | 文件内容的hash结果 | [optional] [default to ""]
**Encoding** | Pointer to **string** | 加密形式 | [optional] [default to ""]
**FileName** | Pointer to **string** | 文件名 | [optional] [default to ""]
**FilePath** | Pointer to **string** | 文件路径 | [optional] [default to ""]
**Sha** | Pointer to **string** | commitID | [optional] [default to ""]
**Size** | Pointer to **int64** | 文件大小 | [optional] 

## Methods

### NewGitFileItem

`func NewGitFileItem() *GitFileItem`

NewGitFileItem instantiates a new GitFileItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFileItemWithDefaults

`func NewGitFileItemWithDefaults() *GitFileItem`

NewGitFileItemWithDefaults instantiates a new GitFileItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *GitFileItem) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GitFileItem) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GitFileItem) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GitFileItem) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *GitFileItem) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *GitFileItem) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentSha256

`func (o *GitFileItem) GetContentSha256() string`

GetContentSha256 returns the ContentSha256 field if non-nil, zero value otherwise.

### GetContentSha256Ok

`func (o *GitFileItem) GetContentSha256Ok() (*string, bool)`

GetContentSha256Ok returns a tuple with the ContentSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSha256

`func (o *GitFileItem) SetContentSha256(v string)`

SetContentSha256 sets ContentSha256 field to given value.

### HasContentSha256

`func (o *GitFileItem) HasContentSha256() bool`

HasContentSha256 returns a boolean if a field has been set.

### SetContentSha256Nil

`func (o *GitFileItem) SetContentSha256Nil(b bool)`

 SetContentSha256Nil sets the value for ContentSha256 to be an explicit nil

### UnsetContentSha256
`func (o *GitFileItem) UnsetContentSha256()`

UnsetContentSha256 ensures that no value is present for ContentSha256, not even an explicit nil
### GetEncoding

`func (o *GitFileItem) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *GitFileItem) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *GitFileItem) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *GitFileItem) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetFileName

`func (o *GitFileItem) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GitFileItem) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GitFileItem) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *GitFileItem) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFilePath

`func (o *GitFileItem) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GitFileItem) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GitFileItem) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GitFileItem) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetSha

`func (o *GitFileItem) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitFileItem) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitFileItem) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitFileItem) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetSize

`func (o *GitFileItem) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GitFileItem) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GitFileItem) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GitFileItem) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


