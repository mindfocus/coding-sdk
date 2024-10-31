# CodingCIProjectDepot

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
**OpenModule** | Pointer to **string** | 无用字段 | [optional] [default to ""]

## Methods

### NewCodingCIProjectDepot

`func NewCodingCIProjectDepot() *CodingCIProjectDepot`

NewCodingCIProjectDepot instantiates a new CodingCIProjectDepot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCIProjectDepotWithDefaults

`func NewCodingCIProjectDepotWithDefaults() *CodingCIProjectDepot`

NewCodingCIProjectDepotWithDefaults instantiates a new CodingCIProjectDepot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *CodingCIProjectDepot) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *CodingCIProjectDepot) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *CodingCIProjectDepot) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *CodingCIProjectDepot) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### SetAuthorizedNil

`func (o *CodingCIProjectDepot) SetAuthorizedNil(b bool)`

 SetAuthorizedNil sets the value for Authorized to be an explicit nil

### UnsetAuthorized
`func (o *CodingCIProjectDepot) UnsetAuthorized()`

UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
### GetDepotHttpsUrl

`func (o *CodingCIProjectDepot) GetDepotHttpsUrl() string`

GetDepotHttpsUrl returns the DepotHttpsUrl field if non-nil, zero value otherwise.

### GetDepotHttpsUrlOk

`func (o *CodingCIProjectDepot) GetDepotHttpsUrlOk() (*string, bool)`

GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotHttpsUrl

`func (o *CodingCIProjectDepot) SetDepotHttpsUrl(v string)`

SetDepotHttpsUrl sets DepotHttpsUrl field to given value.

### HasDepotHttpsUrl

`func (o *CodingCIProjectDepot) HasDepotHttpsUrl() bool`

HasDepotHttpsUrl returns a boolean if a field has been set.

### SetDepotHttpsUrlNil

`func (o *CodingCIProjectDepot) SetDepotHttpsUrlNil(b bool)`

 SetDepotHttpsUrlNil sets the value for DepotHttpsUrl to be an explicit nil

### UnsetDepotHttpsUrl
`func (o *CodingCIProjectDepot) UnsetDepotHttpsUrl()`

UnsetDepotHttpsUrl ensures that no value is present for DepotHttpsUrl, not even an explicit nil
### GetDepotSshUrl

`func (o *CodingCIProjectDepot) GetDepotSshUrl() string`

GetDepotSshUrl returns the DepotSshUrl field if non-nil, zero value otherwise.

### GetDepotSshUrlOk

`func (o *CodingCIProjectDepot) GetDepotSshUrlOk() (*string, bool)`

GetDepotSshUrlOk returns a tuple with the DepotSshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSshUrl

`func (o *CodingCIProjectDepot) SetDepotSshUrl(v string)`

SetDepotSshUrl sets DepotSshUrl field to given value.

### HasDepotSshUrl

`func (o *CodingCIProjectDepot) HasDepotSshUrl() bool`

HasDepotSshUrl returns a boolean if a field has been set.

### SetDepotSshUrlNil

`func (o *CodingCIProjectDepot) SetDepotSshUrlNil(b bool)`

 SetDepotSshUrlNil sets the value for DepotSshUrl to be an explicit nil

### UnsetDepotSshUrl
`func (o *CodingCIProjectDepot) UnsetDepotSshUrl()`

UnsetDepotSshUrl ensures that no value is present for DepotSshUrl, not even an explicit nil
### GetDepotType

`func (o *CodingCIProjectDepot) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *CodingCIProjectDepot) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *CodingCIProjectDepot) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.

### HasDepotType

`func (o *CodingCIProjectDepot) HasDepotType() bool`

HasDepotType returns a boolean if a field has been set.

### GetId

`func (o *CodingCIProjectDepot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodingCIProjectDepot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodingCIProjectDepot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CodingCIProjectDepot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *CodingCIProjectDepot) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CodingCIProjectDepot) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CodingCIProjectDepot) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CodingCIProjectDepot) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *CodingCIProjectDepot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingCIProjectDepot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingCIProjectDepot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodingCIProjectDepot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenModule

`func (o *CodingCIProjectDepot) GetOpenModule() string`

GetOpenModule returns the OpenModule field if non-nil, zero value otherwise.

### GetOpenModuleOk

`func (o *CodingCIProjectDepot) GetOpenModuleOk() (*string, bool)`

GetOpenModuleOk returns a tuple with the OpenModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenModule

`func (o *CodingCIProjectDepot) SetOpenModule(v string)`

SetOpenModule sets OpenModule field to given value.

### HasOpenModule

`func (o *CodingCIProjectDepot) HasOpenModule() bool`

HasOpenModule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


