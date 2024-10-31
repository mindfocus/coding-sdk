# DepotSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 仓库规范描述信息 | [optional] [default to ""]
**Name** | Pointer to **string** | 仓库规范名字 | [optional] [default to ""]
**Type** | Pointer to **string** | 仓库规范类型 system:系统级别，team：团队级别 | [optional] [default to ""]

## Methods

### NewDepotSpec

`func NewDepotSpec() *DepotSpec`

NewDepotSpec instantiates a new DepotSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotSpecWithDefaults

`func NewDepotSpecWithDefaults() *DepotSpec`

NewDepotSpecWithDefaults instantiates a new DepotSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DepotSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DepotSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DepotSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DepotSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DepotSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepotSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepotSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DepotSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DepotSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DepotSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DepotSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DepotSpec) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


