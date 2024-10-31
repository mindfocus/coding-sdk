# IssueFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **int64** | 文件 Id | [optional] 
**Name** | Pointer to **string** | 文件名称 | [optional] [default to ""]
**Size** | Pointer to **int32** | 文件大小 | [optional] 
**Type** | Pointer to **int64** | 类型：  1-文本，  2-图片，  3-二进制文件，  4-SVG | [optional] 
**Url** | Pointer to **string** | 临时下载地址 | [optional] [default to ""]

## Methods

### NewIssueFile

`func NewIssueFile() *IssueFile`

NewIssueFile instantiates a new IssueFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFileWithDefaults

`func NewIssueFileWithDefaults() *IssueFile`

NewIssueFileWithDefaults instantiates a new IssueFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *IssueFile) GetFileId() int64`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *IssueFile) GetFileIdOk() (*int64, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *IssueFile) SetFileId(v int64)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *IssueFile) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetName

`func (o *IssueFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *IssueFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IssueFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IssueFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *IssueFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetType

`func (o *IssueFile) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueFile) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueFile) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *IssueFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *IssueFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssueFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssueFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *IssueFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


