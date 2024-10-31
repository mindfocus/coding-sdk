# CodingCITeamDepot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **NullableBool** | 是否被关联 | [optional] [default to false]
**DepotHttpsUrl** | Pointer to **NullableString** | 仓库 Https 地址 | [optional] [default to ""]
**DepotSshUrl** | Pointer to **NullableString** | 仓库 Ssh 地址 | [optional] [default to ""]
**DepotType** | Pointer to **string** | 仓库类型 | [optional] [default to ""]
**Id** | Pointer to **int32** | 仓库 ID | [optional] 
**IsDefault** | Pointer to **bool** | 是否是默认显示第一位的仓库 | [optional] [default to false]
**Name** | Pointer to **string** | 仓库名称 | [optional] [default to ""]
**OpenModule** | Pointer to **string** | 该仓库是否 CI 可用，如果可用返回值为 continue_integration，如果仓库类型是 CODING 那么这个值永远是continue_integration | [optional] [default to ""]
**ProjectId** | Pointer to **int32** | 项目 ID | [optional] 

## Methods

### NewCodingCITeamDepot

`func NewCodingCITeamDepot() *CodingCITeamDepot`

NewCodingCITeamDepot instantiates a new CodingCITeamDepot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodingCITeamDepotWithDefaults

`func NewCodingCITeamDepotWithDefaults() *CodingCITeamDepot`

NewCodingCITeamDepotWithDefaults instantiates a new CodingCITeamDepot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *CodingCITeamDepot) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *CodingCITeamDepot) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *CodingCITeamDepot) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *CodingCITeamDepot) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### SetAuthorizedNil

`func (o *CodingCITeamDepot) SetAuthorizedNil(b bool)`

 SetAuthorizedNil sets the value for Authorized to be an explicit nil

### UnsetAuthorized
`func (o *CodingCITeamDepot) UnsetAuthorized()`

UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
### GetDepotHttpsUrl

`func (o *CodingCITeamDepot) GetDepotHttpsUrl() string`

GetDepotHttpsUrl returns the DepotHttpsUrl field if non-nil, zero value otherwise.

### GetDepotHttpsUrlOk

`func (o *CodingCITeamDepot) GetDepotHttpsUrlOk() (*string, bool)`

GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotHttpsUrl

`func (o *CodingCITeamDepot) SetDepotHttpsUrl(v string)`

SetDepotHttpsUrl sets DepotHttpsUrl field to given value.

### HasDepotHttpsUrl

`func (o *CodingCITeamDepot) HasDepotHttpsUrl() bool`

HasDepotHttpsUrl returns a boolean if a field has been set.

### SetDepotHttpsUrlNil

`func (o *CodingCITeamDepot) SetDepotHttpsUrlNil(b bool)`

 SetDepotHttpsUrlNil sets the value for DepotHttpsUrl to be an explicit nil

### UnsetDepotHttpsUrl
`func (o *CodingCITeamDepot) UnsetDepotHttpsUrl()`

UnsetDepotHttpsUrl ensures that no value is present for DepotHttpsUrl, not even an explicit nil
### GetDepotSshUrl

`func (o *CodingCITeamDepot) GetDepotSshUrl() string`

GetDepotSshUrl returns the DepotSshUrl field if non-nil, zero value otherwise.

### GetDepotSshUrlOk

`func (o *CodingCITeamDepot) GetDepotSshUrlOk() (*string, bool)`

GetDepotSshUrlOk returns a tuple with the DepotSshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotSshUrl

`func (o *CodingCITeamDepot) SetDepotSshUrl(v string)`

SetDepotSshUrl sets DepotSshUrl field to given value.

### HasDepotSshUrl

`func (o *CodingCITeamDepot) HasDepotSshUrl() bool`

HasDepotSshUrl returns a boolean if a field has been set.

### SetDepotSshUrlNil

`func (o *CodingCITeamDepot) SetDepotSshUrlNil(b bool)`

 SetDepotSshUrlNil sets the value for DepotSshUrl to be an explicit nil

### UnsetDepotSshUrl
`func (o *CodingCITeamDepot) UnsetDepotSshUrl()`

UnsetDepotSshUrl ensures that no value is present for DepotSshUrl, not even an explicit nil
### GetDepotType

`func (o *CodingCITeamDepot) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *CodingCITeamDepot) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *CodingCITeamDepot) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.

### HasDepotType

`func (o *CodingCITeamDepot) HasDepotType() bool`

HasDepotType returns a boolean if a field has been set.

### GetId

`func (o *CodingCITeamDepot) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodingCITeamDepot) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodingCITeamDepot) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CodingCITeamDepot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *CodingCITeamDepot) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CodingCITeamDepot) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CodingCITeamDepot) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CodingCITeamDepot) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *CodingCITeamDepot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodingCITeamDepot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodingCITeamDepot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CodingCITeamDepot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenModule

`func (o *CodingCITeamDepot) GetOpenModule() string`

GetOpenModule returns the OpenModule field if non-nil, zero value otherwise.

### GetOpenModuleOk

`func (o *CodingCITeamDepot) GetOpenModuleOk() (*string, bool)`

GetOpenModuleOk returns a tuple with the OpenModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenModule

`func (o *CodingCITeamDepot) SetOpenModule(v string)`

SetOpenModule sets OpenModule field to given value.

### HasOpenModule

`func (o *CodingCITeamDepot) HasOpenModule() bool`

HasOpenModule returns a boolean if a field has been set.

### GetProjectId

`func (o *CodingCITeamDepot) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CodingCITeamDepot) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CodingCITeamDepot) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *CodingCITeamDepot) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


