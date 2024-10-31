# ArtifactPropertyBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 属性名称（以 ‘coding.’ 作为属性名称开头的属性，将不可变更及删除，即 Immutable &#x3D; false） | [default to ""]
**Value** | **string** | 属性值 | [default to ""]

## Methods

### NewArtifactPropertyBean

`func NewArtifactPropertyBean(name string, value string, ) *ArtifactPropertyBean`

NewArtifactPropertyBean instantiates a new ArtifactPropertyBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactPropertyBeanWithDefaults

`func NewArtifactPropertyBeanWithDefaults() *ArtifactPropertyBean`

NewArtifactPropertyBeanWithDefaults instantiates a new ArtifactPropertyBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArtifactPropertyBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactPropertyBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactPropertyBean) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *ArtifactPropertyBean) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ArtifactPropertyBean) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ArtifactPropertyBean) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


