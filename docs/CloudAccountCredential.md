# CloudAccountCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubeConfig** | Pointer to [**KubeConfigForm**](KubeConfigForm.md) |  | [optional] 
**KubernetesType** | Pointer to **NullableString** | KUBERNETES 类型云账号认证方式（可选值：KUBE_CONFIG、SERVICE_ACCOUNT、TKE），TENCENT 类型云账号可不填 | [optional] [default to ""]
**ServiceAccount** | Pointer to [**ServiceAccountForm**](ServiceAccountForm.md) |  | [optional] 
**TKEConfig** | Pointer to [**TKEConfigForm**](TKEConfigForm.md) |  | [optional] 
**TencentCloudAccount** | Pointer to [**TencentCloudAccountForm**](TencentCloudAccountForm.md) |  | [optional] 

## Methods

### NewCloudAccountCredential

`func NewCloudAccountCredential() *CloudAccountCredential`

NewCloudAccountCredential instantiates a new CloudAccountCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountCredentialWithDefaults

`func NewCloudAccountCredentialWithDefaults() *CloudAccountCredential`

NewCloudAccountCredentialWithDefaults instantiates a new CloudAccountCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeConfig

`func (o *CloudAccountCredential) GetKubeConfig() KubeConfigForm`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *CloudAccountCredential) GetKubeConfigOk() (*KubeConfigForm, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *CloudAccountCredential) SetKubeConfig(v KubeConfigForm)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *CloudAccountCredential) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### GetKubernetesType

`func (o *CloudAccountCredential) GetKubernetesType() string`

GetKubernetesType returns the KubernetesType field if non-nil, zero value otherwise.

### GetKubernetesTypeOk

`func (o *CloudAccountCredential) GetKubernetesTypeOk() (*string, bool)`

GetKubernetesTypeOk returns a tuple with the KubernetesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesType

`func (o *CloudAccountCredential) SetKubernetesType(v string)`

SetKubernetesType sets KubernetesType field to given value.

### HasKubernetesType

`func (o *CloudAccountCredential) HasKubernetesType() bool`

HasKubernetesType returns a boolean if a field has been set.

### SetKubernetesTypeNil

`func (o *CloudAccountCredential) SetKubernetesTypeNil(b bool)`

 SetKubernetesTypeNil sets the value for KubernetesType to be an explicit nil

### UnsetKubernetesType
`func (o *CloudAccountCredential) UnsetKubernetesType()`

UnsetKubernetesType ensures that no value is present for KubernetesType, not even an explicit nil
### GetServiceAccount

`func (o *CloudAccountCredential) GetServiceAccount() ServiceAccountForm`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *CloudAccountCredential) GetServiceAccountOk() (*ServiceAccountForm, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *CloudAccountCredential) SetServiceAccount(v ServiceAccountForm)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *CloudAccountCredential) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetTKEConfig

`func (o *CloudAccountCredential) GetTKEConfig() TKEConfigForm`

GetTKEConfig returns the TKEConfig field if non-nil, zero value otherwise.

### GetTKEConfigOk

`func (o *CloudAccountCredential) GetTKEConfigOk() (*TKEConfigForm, bool)`

GetTKEConfigOk returns a tuple with the TKEConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTKEConfig

`func (o *CloudAccountCredential) SetTKEConfig(v TKEConfigForm)`

SetTKEConfig sets TKEConfig field to given value.

### HasTKEConfig

`func (o *CloudAccountCredential) HasTKEConfig() bool`

HasTKEConfig returns a boolean if a field has been set.

### GetTencentCloudAccount

`func (o *CloudAccountCredential) GetTencentCloudAccount() TencentCloudAccountForm`

GetTencentCloudAccount returns the TencentCloudAccount field if non-nil, zero value otherwise.

### GetTencentCloudAccountOk

`func (o *CloudAccountCredential) GetTencentCloudAccountOk() (*TencentCloudAccountForm, bool)`

GetTencentCloudAccountOk returns a tuple with the TencentCloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTencentCloudAccount

`func (o *CloudAccountCredential) SetTencentCloudAccount(v TencentCloudAccountForm)`

SetTencentCloudAccount sets TencentCloudAccount field to given value.

### HasTencentCloudAccount

`func (o *CloudAccountCredential) HasTencentCloudAccount() bool`

HasTencentCloudAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


