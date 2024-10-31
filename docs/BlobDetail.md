# BlobDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobSha** | Pointer to **NullableString** | blob 的 Sha 值 | [optional] [default to ""]
**Content** | Pointer to **NullableString** | blob 内容使用 base64编码后的内容 | [optional] [default to ""]
**Encoding** | Pointer to **NullableString** | 编码方式(目前就base64一种) | [optional] [default to ""]
**Size** | Pointer to **NullableInt64** | blob大小(byte) | [optional] 

## Methods

### NewBlobDetail

`func NewBlobDetail() *BlobDetail`

NewBlobDetail instantiates a new BlobDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlobDetailWithDefaults

`func NewBlobDetailWithDefaults() *BlobDetail`

NewBlobDetailWithDefaults instantiates a new BlobDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobSha

`func (o *BlobDetail) GetBlobSha() string`

GetBlobSha returns the BlobSha field if non-nil, zero value otherwise.

### GetBlobShaOk

`func (o *BlobDetail) GetBlobShaOk() (*string, bool)`

GetBlobShaOk returns a tuple with the BlobSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSha

`func (o *BlobDetail) SetBlobSha(v string)`

SetBlobSha sets BlobSha field to given value.

### HasBlobSha

`func (o *BlobDetail) HasBlobSha() bool`

HasBlobSha returns a boolean if a field has been set.

### SetBlobShaNil

`func (o *BlobDetail) SetBlobShaNil(b bool)`

 SetBlobShaNil sets the value for BlobSha to be an explicit nil

### UnsetBlobSha
`func (o *BlobDetail) UnsetBlobSha()`

UnsetBlobSha ensures that no value is present for BlobSha, not even an explicit nil
### GetContent

`func (o *BlobDetail) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BlobDetail) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BlobDetail) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BlobDetail) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *BlobDetail) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *BlobDetail) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEncoding

`func (o *BlobDetail) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *BlobDetail) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *BlobDetail) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *BlobDetail) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### SetEncodingNil

`func (o *BlobDetail) SetEncodingNil(b bool)`

 SetEncodingNil sets the value for Encoding to be an explicit nil

### UnsetEncoding
`func (o *BlobDetail) UnsetEncoding()`

UnsetEncoding ensures that no value is present for Encoding, not even an explicit nil
### GetSize

`func (o *BlobDetail) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *BlobDetail) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *BlobDetail) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *BlobDetail) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *BlobDetail) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *BlobDetail) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


