# AttachmentPrepare

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentId** | Pointer to **NullableInt32** | 附件 ID | [optional] 
**PrepareSignUrl** | Pointer to **NullableString** | 附件上传地址 | [optional] [default to ""]

## Methods

### NewAttachmentPrepare

`func NewAttachmentPrepare() *AttachmentPrepare`

NewAttachmentPrepare instantiates a new AttachmentPrepare object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentPrepareWithDefaults

`func NewAttachmentPrepareWithDefaults() *AttachmentPrepare`

NewAttachmentPrepareWithDefaults instantiates a new AttachmentPrepare object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentId

`func (o *AttachmentPrepare) GetAttachmentId() int32`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *AttachmentPrepare) GetAttachmentIdOk() (*int32, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *AttachmentPrepare) SetAttachmentId(v int32)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *AttachmentPrepare) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### SetAttachmentIdNil

`func (o *AttachmentPrepare) SetAttachmentIdNil(b bool)`

 SetAttachmentIdNil sets the value for AttachmentId to be an explicit nil

### UnsetAttachmentId
`func (o *AttachmentPrepare) UnsetAttachmentId()`

UnsetAttachmentId ensures that no value is present for AttachmentId, not even an explicit nil
### GetPrepareSignUrl

`func (o *AttachmentPrepare) GetPrepareSignUrl() string`

GetPrepareSignUrl returns the PrepareSignUrl field if non-nil, zero value otherwise.

### GetPrepareSignUrlOk

`func (o *AttachmentPrepare) GetPrepareSignUrlOk() (*string, bool)`

GetPrepareSignUrlOk returns a tuple with the PrepareSignUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareSignUrl

`func (o *AttachmentPrepare) SetPrepareSignUrl(v string)`

SetPrepareSignUrl sets PrepareSignUrl field to given value.

### HasPrepareSignUrl

`func (o *AttachmentPrepare) HasPrepareSignUrl() bool`

HasPrepareSignUrl returns a boolean if a field has been set.

### SetPrepareSignUrlNil

`func (o *AttachmentPrepare) SetPrepareSignUrlNil(b bool)`

 SetPrepareSignUrlNil sets the value for PrepareSignUrl to be an explicit nil

### UnsetPrepareSignUrl
`func (o *AttachmentPrepare) UnsetPrepareSignUrl()`

UnsetPrepareSignUrl ensures that no value is present for PrepareSignUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


