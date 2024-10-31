# ArtifactProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**Id** | Pointer to **int32** | 属性 ID | [optional] 
**Immutable** | Pointer to **bool** | 是否不可变更 | [optional] [default to false]
**Name** | Pointer to **string** | 属性名称（以 ‘coding.’ 作为属性名称开头的属性，将不可变更及删除，即 Immutable &#x3D; false） | [optional] [default to ""]
**Value** | Pointer to **string** | 属性值 | [optional] [default to ""]
**Version** | Pointer to **string** | 制品版本 | [optional] [default to ""]

## Methods

### NewArtifactProperty

`func NewArtifactProperty() *ArtifactProperty`

NewArtifactProperty instantiates a new ArtifactProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactPropertyWithDefaults

`func NewArtifactPropertyWithDefaults() *ArtifactProperty`

NewArtifactPropertyWithDefaults instantiates a new ArtifactProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArtifactProperty) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactProperty) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactProperty) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactProperty) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ArtifactProperty) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactProperty) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactProperty) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactProperty) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImmutable

`func (o *ArtifactProperty) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *ArtifactProperty) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *ArtifactProperty) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *ArtifactProperty) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetName

`func (o *ArtifactProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *ArtifactProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArtifactProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArtifactProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ArtifactProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetVersion

`func (o *ArtifactProperty) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactProperty) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactProperty) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArtifactProperty) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


