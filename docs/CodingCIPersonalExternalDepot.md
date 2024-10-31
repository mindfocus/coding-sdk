# CodingCIPersonalExternalDepot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **NullableBool** | 外部仓库是否被关联 | [optional] [default to false]
**DepotHttpsUrl** | Pointer to **NullableString** | 仓库 Https 地址 | [optional] [default to ""]
**DepotSshUrl** | Pointer to **NullableString** | 仓库 Ssh 地址 | [optional] [default to ""]
**DepotType** | Pointer to **string** | 仓库类型 | [optional] [default to ""]
**Id** | Pointer to **int32** | 仓库 Id | [optional] 
**IsDefault** | Pointer to **bool** | 是否是默认显示第一位的仓库 | [optional] [default to false]
**Name** | Pointer to **string** | 仓库名称 | [optional] [default to ""]
**OpenModule** | Pointer to **string** | 请使用 Authorized 代替判断仓库是否关联，该仓库是否 CI 可用，如果可用返回值为 continue_integration，如果仓库类型是 CODING 那么这个值永远是continue_integration | [optional] [default to ""]
**OwnerName** | Pointer to **NullableString** | 所有者用户名 | [optional] [default to ""]
**SourceDepotId** | Pointer to **int64** | 源代码仓库 id （如果是 coding 跨项目代码仓库，此字段记录代码仓库的实际 id ） | [optional] 

## Methods

### NewCodingCIPersonalExternalDepot

`func NewCodingCIPersonalExternalDepot() *CodingCIPersonalExternalDepot`

NewCodingCIPersonalExternalDepot instantiates a new CodingCIPersonalExternalDepot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIPersonalExternalDepotWithDefaults

`func NewCodingCIPersonalExternalDepotWithDefaults() *CodingCIPersonalExternalDepot`

NewCodingCIPersonalExternalDepotWithDefaults instantiates a new CodingCIPersonalExternalDepot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *CodingCIPersonalExternalDepot) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *CodingCIPersonalExternalDepot) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *CodingCIPersonalExternalDepot) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *CodingCIPersonalExternalDepot) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### SetAuthorizedNil

`func (o *CodingCIPersonalExternalDepot) SetAuthorizedNil(b bool)`

 SetAuthorizedNil sets the value for Authorized to be an explicit nil

### UnsetAuthorized
`func (o *CodingCIPersonalExternalDepot) UnsetAuthorized()`

UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
### GetDepotHttpsUrl

`func (o *CodingCIPersonalExternalDepot) GetDepotHttpsUrl() string`

GetDepotHttpsUrl returns the DepotHttpsUrl field if non-nil, zero value otherwise.

### GetDepotHttpsUrlOk

`func (o *CodingCIPersonalExternalDepot) GetDepotHttpsUrlOk() (*string, bool)`

GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotHttpsUrl

`func (o *CodingCIPersonalExternalDepot) SetDepotHttpsUrl(v string)`

SetDepotHttpsUrl sets DepotHttpsUrl field to given value.

### HasDepotHttpsUrl

`func (o *CodingCIPersonalExternalDepot) HasDepotHttpsUrl() bool`

HasDepotHttpsUrl returns a boolean if a field has been set.

### SetDepotHttpsUrlNil

`func (o *CodingCIPersonalExternalDepot) SetDepotHttpsUrlNil(b bool)`

 SetDepotHttpsUrlNil sets the value for DepotHttpsUrl to be an explicit nil

### UnsetDepotHttpsUrl
`func (o *CodingCIPersonalExternalDepot) UnsetDepotHttpsUrl()`

UnsetDepotHttpsUrl ensures that no value is present for DepotHttpsUrl, not even an explicit nil
### GetDepotSshUrl

`func (o *CodingCIPersonalExternalDepot) GetDepotSshUrl() string`

GetDepotSshUrl returns the DepotSshUrl field if non-nil, zero value otherwise.

### GetDepotSshUrlOk

`func (o *CodingCIPersonalExternalDepot) GetDepotSshUrlOk() (*string, bool)`

GetDepotSshUrlOk returns a tuple with the DepotSshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSshUrl

`func (o *CodingCIPersonalExternalDepot) SetDepotSshUrl(v string)`

SetDepotSshUrl sets DepotSshUrl field to given value.

### HasDepotSshUrl

`func (o *CodingCIPersonalExternalDepot) HasDepotSshUrl() bool`

HasDepotSshUrl returns a boolean if a field has been set.

### SetDepotSshUrlNil

`func (o *CodingCIPersonalExternalDepot) SetDepotSshUrlNil(b bool)`

 SetDepotSshUrlNil sets the value for DepotSshUrl to be an explicit nil

### UnsetDepotSshUrl
`func (o *CodingCIPersonalExternalDepot) UnsetDepotSshUrl()`

UnsetDepotSshUrl ensures that no value is present for DepotSshUrl, not even an explicit nil
### GetDepotType

`func (o *CodingCIPersonalExternalDepot) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *CodingCIPersonalExternalDepot) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *CodingCIPersonalExternalDepot) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.

### HasDepotType

`func (o *CodingCIPersonalExternalDepot) HasDepotType() bool`

HasDepotType returns a boolean if a field has been set.

### GetId

`func (o *CodingCIPersonalExternalDepot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodingCIPersonalExternalDepot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodingCIPersonalExternalDepot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CodingCIPersonalExternalDepot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *CodingCIPersonalExternalDepot) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CodingCIPersonalExternalDepot) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CodingCIPersonalExternalDepot) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CodingCIPersonalExternalDepot) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *CodingCIPersonalExternalDepot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingCIPersonalExternalDepot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingCIPersonalExternalDepot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodingCIPersonalExternalDepot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenModule

`func (o *CodingCIPersonalExternalDepot) GetOpenModule() string`

GetOpenModule returns the OpenModule field if non-nil, zero value otherwise.

### GetOpenModuleOk

`func (o *CodingCIPersonalExternalDepot) GetOpenModuleOk() (*string, bool)`

GetOpenModuleOk returns a tuple with the OpenModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenModule

`func (o *CodingCIPersonalExternalDepot) SetOpenModule(v string)`

SetOpenModule sets OpenModule field to given value.

### HasOpenModule

`func (o *CodingCIPersonalExternalDepot) HasOpenModule() bool`

HasOpenModule returns a boolean if a field has been set.

### GetOwnerName

`func (o *CodingCIPersonalExternalDepot) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *CodingCIPersonalExternalDepot) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *CodingCIPersonalExternalDepot) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *CodingCIPersonalExternalDepot) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### SetOwnerNameNil

`func (o *CodingCIPersonalExternalDepot) SetOwnerNameNil(b bool)`

 SetOwnerNameNil sets the value for OwnerName to be an explicit nil

### UnsetOwnerName
`func (o *CodingCIPersonalExternalDepot) UnsetOwnerName()`

UnsetOwnerName ensures that no value is present for OwnerName, not even an explicit nil
### GetSourceDepotId

`func (o *CodingCIPersonalExternalDepot) GetSourceDepotId() int64`

GetSourceDepotId returns the SourceDepotId field if non-nil, zero value otherwise.

### GetSourceDepotIdOk

`func (o *CodingCIPersonalExternalDepot) GetSourceDepotIdOk() (*int64, bool)`

GetSourceDepotIdOk returns a tuple with the SourceDepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDepotId

`func (o *CodingCIPersonalExternalDepot) SetSourceDepotId(v int64)`

SetSourceDepotId sets SourceDepotId field to given value.

### HasSourceDepotId

`func (o *CodingCIPersonalExternalDepot) HasSourceDepotId() bool`

HasSourceDepotId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


