# KubeConfigForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterContext** | **string** | 指定访问集群 KubeConfig 文件的上下文 | [default to ""]
**InsecureSkipTLSVerify** | **bool** | 是否接受非认证证书（是：true；否：false） | [default to false]
**KubeConfig** | Pointer to **NullableString** | 访问集群的 KubeConfig 文件（YAML 格式，Base64 编码），添加时必填，修改时可不填 | [optional] [default to ""]
**OnlySpinnakerManaged** | **bool** | 是否允许持续部署管理集群已有资源（是：false；否：true） | [default to false]

## Methods

### NewKubeConfigForm

`func NewKubeConfigForm(clusterContext string, insecureSkipTLSVerify bool, onlySpinnakerManaged bool, ) *KubeConfigForm`

NewKubeConfigForm instantiates a new KubeConfigForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubeConfigFormWithDefaults

`func NewKubeConfigFormWithDefaults() *KubeConfigForm`

NewKubeConfigFormWithDefaults instantiates a new KubeConfigForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterContext

`func (o *KubeConfigForm) GetClusterContext() string`

GetClusterContext returns the ClusterContext field if non-nil, zero value otherwise.

### GetClusterContextOk

`func (o *KubeConfigForm) GetClusterContextOk() (*string, bool)`

GetClusterContextOk returns a tuple with the ClusterContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterContext

`func (o *KubeConfigForm) SetClusterContext(v string)`

SetClusterContext sets ClusterContext field to given value.


### GetInsecureSkipTLSVerify

`func (o *KubeConfigForm) GetInsecureSkipTLSVerify() bool`

GetInsecureSkipTLSVerify returns the InsecureSkipTLSVerify field if non-nil, zero value otherwise.

### GetInsecureSkipTLSVerifyOk

`func (o *KubeConfigForm) GetInsecureSkipTLSVerifyOk() (*bool, bool)`

GetInsecureSkipTLSVerifyOk returns a tuple with the InsecureSkipTLSVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureSkipTLSVerify

`func (o *KubeConfigForm) SetInsecureSkipTLSVerify(v bool)`

SetInsecureSkipTLSVerify sets InsecureSkipTLSVerify field to given value.


### GetKubeConfig

`func (o *KubeConfigForm) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *KubeConfigForm) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *KubeConfigForm) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *KubeConfigForm) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### SetKubeConfigNil

`func (o *KubeConfigForm) SetKubeConfigNil(b bool)`

 SetKubeConfigNil sets the value for KubeConfig to be an explicit nil

### UnsetKubeConfig
`func (o *KubeConfigForm) UnsetKubeConfig()`

UnsetKubeConfig ensures that no value is present for KubeConfig, not even an explicit nil
### GetOnlySpinnakerManaged

`func (o *KubeConfigForm) GetOnlySpinnakerManaged() bool`

GetOnlySpinnakerManaged returns the OnlySpinnakerManaged field if non-nil, zero value otherwise.

### GetOnlySpinnakerManagedOk

`func (o *KubeConfigForm) GetOnlySpinnakerManagedOk() (*bool, bool)`

GetOnlySpinnakerManagedOk returns a tuple with the OnlySpinnakerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlySpinnakerManaged

`func (o *KubeConfigForm) SetOnlySpinnakerManaged(v bool)`

SetOnlySpinnakerManaged sets OnlySpinnakerManaged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


