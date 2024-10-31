# TencentCloudAccountForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | **[]string** | CVM 地域 | 
**SecretId** | **NullableString** | 腾讯云 API 密钥 ID | [default to ""]
**SecretKey** | Pointer to **NullableString** | 腾讯云 API 密钥 Key | [optional] [default to ""]

## Methods

### NewTencentCloudAccountForm

`func NewTencentCloudAccountForm(regions []string, secretId NullableString, ) *TencentCloudAccountForm`

NewTencentCloudAccountForm instantiates a new TencentCloudAccountForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTencentCloudAccountFormWithDefaults

`func NewTencentCloudAccountFormWithDefaults() *TencentCloudAccountForm`

NewTencentCloudAccountFormWithDefaults instantiates a new TencentCloudAccountForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *TencentCloudAccountForm) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *TencentCloudAccountForm) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *TencentCloudAccountForm) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### SetRegionsNil

`func (o *TencentCloudAccountForm) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *TencentCloudAccountForm) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetSecretId

`func (o *TencentCloudAccountForm) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *TencentCloudAccountForm) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *TencentCloudAccountForm) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### SetSecretIdNil

`func (o *TencentCloudAccountForm) SetSecretIdNil(b bool)`

 SetSecretIdNil sets the value for SecretId to be an explicit nil

### UnsetSecretId
`func (o *TencentCloudAccountForm) UnsetSecretId()`

UnsetSecretId ensures that no value is present for SecretId, not even an explicit nil
### GetSecretKey

`func (o *TencentCloudAccountForm) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *TencentCloudAccountForm) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *TencentCloudAccountForm) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *TencentCloudAccountForm) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *TencentCloudAccountForm) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *TencentCloudAccountForm) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


