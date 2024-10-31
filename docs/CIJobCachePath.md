# CIJobCachePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **string** | 绝对路径 | [default to ""]
**Enabled** | **bool** | 是否启用 | [default to false]

## Methods

### NewCIJobCachePath

`func NewCIJobCachePath(absolutePath string, enabled bool, ) *CIJobCachePath`

NewCIJobCachePath instantiates a new CIJobCachePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCIJobCachePathWithDefaults

`func NewCIJobCachePathWithDefaults() *CIJobCachePath`

NewCIJobCachePathWithDefaults instantiates a new CIJobCachePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *CIJobCachePath) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *CIJobCachePath) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *CIJobCachePath) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### GetEnabled

`func (o *CIJobCachePath) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CIJobCachePath) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CIJobCachePath) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


