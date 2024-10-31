# ServiceAccountForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InsecureSkipTLSVerify** | **bool** | 是否接受非认证证书（是：true；否：false） | [default to false]
**OnlySpinnakerManaged** | **bool** | 是否允许持续部署管理集群已有资源（是：false；否：true） | [default to false]
**Secret** | Pointer to **NullableString** | ServiceAccount 关联的 Secret（YAML 格式，Base64 编码），添加时必填，修改时可不填 | [optional] [default to ""]
**Server** | **string** | API Server URL | [default to ""]

## Methods

### NewServiceAccountForm

`func NewServiceAccountForm(insecureSkipTLSVerify bool, onlySpinnakerManaged bool, server string, ) *ServiceAccountForm`

NewServiceAccountForm instantiates a new ServiceAccountForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountFormWithDefaults

`func NewServiceAccountFormWithDefaults() *ServiceAccountForm`

NewServiceAccountFormWithDefaults instantiates a new ServiceAccountForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsecureSkipTLSVerify

`func (o *ServiceAccountForm) GetInsecureSkipTLSVerify() bool`

GetInsecureSkipTLSVerify returns the InsecureSkipTLSVerify field if non-nil, zero value otherwise.

### GetInsecureSkipTLSVerifyOk

`func (o *ServiceAccountForm) GetInsecureSkipTLSVerifyOk() (*bool, bool)`

GetInsecureSkipTLSVerifyOk returns a tuple with the InsecureSkipTLSVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipTLSVerify

`func (o *ServiceAccountForm) SetInsecureSkipTLSVerify(v bool)`

SetInsecureSkipTLSVerify sets InsecureSkipTLSVerify field to given value.


### GetOnlySpinnakerManaged

`func (o *ServiceAccountForm) GetOnlySpinnakerManaged() bool`

GetOnlySpinnakerManaged returns the OnlySpinnakerManaged field if non-nil, zero value otherwise.

### GetOnlySpinnakerManagedOk

`func (o *ServiceAccountForm) GetOnlySpinnakerManagedOk() (*bool, bool)`

GetOnlySpinnakerManagedOk returns a tuple with the OnlySpinnakerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlySpinnakerManaged

`func (o *ServiceAccountForm) SetOnlySpinnakerManaged(v bool)`

SetOnlySpinnakerManaged sets OnlySpinnakerManaged field to given value.


### GetSecret

`func (o *ServiceAccountForm) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ServiceAccountForm) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ServiceAccountForm) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ServiceAccountForm) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### SetSecretNil

`func (o *ServiceAccountForm) SetSecretNil(b bool)`

 SetSecretNil sets the value for Secret to be an explicit nil

### UnsetSecret
`func (o *ServiceAccountForm) UnsetSecret()`

UnsetSecret ensures that no value is present for Secret, not even an explicit nil
### GetServer

`func (o *ServiceAccountForm) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ServiceAccountForm) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ServiceAccountForm) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


