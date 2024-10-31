# TKEConfigForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | TKE 集群 ID | [default to ""]
**Namespaces** | Pointer to **[]string** | 将为选择的每个命名空间自动生成用于访问 CODING Docker 仓库的凭据（ImagePullSecrets） | [optional] 
**OnlySpinnakerManaged** | **bool** | 是否允许持续部署管理集群已有资源（是：false；否：true） | [default to false]
**Region** | **string** | TKE 地域 | [default to ""]

## Methods

### NewTKEConfigForm

`func NewTKEConfigForm(clusterId string, onlySpinnakerManaged bool, region string, ) *TKEConfigForm`

NewTKEConfigForm instantiates a new TKEConfigForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTKEConfigFormWithDefaults

`func NewTKEConfigFormWithDefaults() *TKEConfigForm`

NewTKEConfigFormWithDefaults instantiates a new TKEConfigForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *TKEConfigForm) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *TKEConfigForm) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *TKEConfigForm) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetNamespaces

`func (o *TKEConfigForm) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *TKEConfigForm) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *TKEConfigForm) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *TKEConfigForm) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.

### SetNamespacesNil

`func (o *TKEConfigForm) SetNamespacesNil(b bool)`

 SetNamespacesNil sets the value for Namespaces to be an explicit nil

### UnsetNamespaces
`func (o *TKEConfigForm) UnsetNamespaces()`

UnsetNamespaces ensures that no value is present for Namespaces, not even an explicit nil
### GetOnlySpinnakerManaged

`func (o *TKEConfigForm) GetOnlySpinnakerManaged() bool`

GetOnlySpinnakerManaged returns the OnlySpinnakerManaged field if non-nil, zero value otherwise.

### GetOnlySpinnakerManagedOk

`func (o *TKEConfigForm) GetOnlySpinnakerManagedOk() (*bool, bool)`

GetOnlySpinnakerManagedOk returns a tuple with the OnlySpinnakerManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlySpinnakerManaged

`func (o *TKEConfigForm) SetOnlySpinnakerManaged(v bool)`

SetOnlySpinnakerManaged sets OnlySpinnakerManaged field to given value.


### GetRegion

`func (o *TKEConfigForm) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TKEConfigForm) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TKEConfigForm) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


