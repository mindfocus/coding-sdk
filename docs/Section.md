# Section

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | 分组ID | [optional] 
**Name** | Pointer to **string** | 名称 | [optional] [default to ""]
**ParentId** | Pointer to **int32** | 父级 ID | [optional] 
**Sort** | Pointer to **int32** | 排序值 | [optional] 

## Methods

### NewSection

`func NewSection() *Section`

NewSection instantiates a new Section object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSectionWithDefaults

`func NewSectionWithDefaults() *Section`

NewSectionWithDefaults instantiates a new Section object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Section) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Section) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Section) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Section) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Section) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Section) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Section) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Section) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *Section) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Section) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Section) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Section) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetSort

`func (o *Section) GetSort() int32`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *Section) GetSortOk() (*int32, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *Section) SetSort(v int32)`

SetSort sets Sort field to given value.

### HasSort

`func (o *Section) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


