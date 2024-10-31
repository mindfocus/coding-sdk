# DefectType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IconUrl** | Pointer to **string** | 图标地址 | [optional] [default to ""]
**Id** | Pointer to **int64** | 缺陷类型 Id | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]

## Methods

### NewDefectType

`func NewDefectType() *DefectType`

NewDefectType instantiates a new DefectType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefectTypeWithDefaults

`func NewDefectTypeWithDefaults() *DefectType`

NewDefectTypeWithDefaults instantiates a new DefectType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIconUrl

`func (o *DefectType) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *DefectType) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *DefectType) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *DefectType) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *DefectType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DefectType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DefectType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DefectType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DefectType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefectType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefectType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefectType) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


