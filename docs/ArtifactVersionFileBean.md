# ArtifactVersionFileBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | 文件名称 | [optional] [default to ""]
**Size** | Pointer to **float32** | 文件大小，单位：MB | [optional] 

## Methods

### NewArtifactVersionFileBean

`func NewArtifactVersionFileBean() *ArtifactVersionFileBean`

NewArtifactVersionFileBean instantiates a new ArtifactVersionFileBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactVersionFileBeanWithDefaults

`func NewArtifactVersionFileBeanWithDefaults() *ArtifactVersionFileBean`

NewArtifactVersionFileBeanWithDefaults instantiates a new ArtifactVersionFileBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArtifactVersionFileBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArtifactVersionFileBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArtifactVersionFileBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArtifactVersionFileBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ArtifactVersionFileBean) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactVersionFileBean) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactVersionFileBean) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtifactVersionFileBean) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


