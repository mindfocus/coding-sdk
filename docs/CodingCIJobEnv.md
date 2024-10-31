# CodingCIJobEnv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 环境变量名称 | [default to ""]
**Sensitive** | **bool** | 是否保密 | [default to false]
**Value** | **string** | 环境变量值 | [default to ""]

## Methods

### NewCodingCIJobEnv

`func NewCodingCIJobEnv(name string, sensitive bool, value string, ) *CodingCIJobEnv`

NewCodingCIJobEnv instantiates a new CodingCIJobEnv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIJobEnvWithDefaults

`func NewCodingCIJobEnvWithDefaults() *CodingCIJobEnv`

NewCodingCIJobEnvWithDefaults instantiates a new CodingCIJobEnv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodingCIJobEnv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingCIJobEnv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingCIJobEnv) SetName(v string)`

SetName sets Name field to given value.


### GetSensitive

`func (o *CodingCIJobEnv) GetSensitive() bool`

GetSensitive returns the Sensitive field if non-nil, zero value otherwise.

### GetSensitiveOk

`func (o *CodingCIJobEnv) GetSensitiveOk() (*bool, bool)`

GetSensitiveOk returns a tuple with the Sensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitive

`func (o *CodingCIJobEnv) SetSensitive(v bool)`

SetSensitive sets Sensitive field to given value.


### GetValue

`func (o *CodingCIJobEnv) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CodingCIJobEnv) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CodingCIJobEnv) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


