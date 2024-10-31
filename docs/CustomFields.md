# CustomFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | 自定义属性 Id | [optional] 
**Name** | Pointer to **string** | 自定义属性名称 | [optional] [default to ""]
**ValueString** | Pointer to **string** | 自定义属性值 | [optional] [default to ""]

## Methods

### NewCustomFields

`func NewCustomFields() *CustomFields`

NewCustomFields instantiates a new CustomFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldsWithDefaults

`func NewCustomFieldsWithDefaults() *CustomFields`

NewCustomFieldsWithDefaults instantiates a new CustomFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFields) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFields) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFields) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueString

`func (o *CustomFields) GetValueString() string`

GetValueString returns the ValueString field if non-nil, zero value otherwise.

### GetValueStringOk

`func (o *CustomFields) GetValueStringOk() (*string, bool)`

GetValueStringOk returns a tuple with the ValueString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueString

`func (o *CustomFields) SetValueString(v string)`

SetValueString sets ValueString field to given value.

### HasValueString

`func (o *CustomFields) HasValueString() bool`

HasValueString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


