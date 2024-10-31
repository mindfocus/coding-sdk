# Maven

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | **string** | 附属构件 | [default to ""]
**Packaging** | **string** | 打包方式：pom;jar;war | [default to ""]

## Methods

### NewMaven

`func NewMaven(classifier string, packaging string, ) *Maven`

NewMaven instantiates a new Maven object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMavenWithDefaults

`func NewMavenWithDefaults() *Maven`

NewMavenWithDefaults instantiates a new Maven object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *Maven) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *Maven) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *Maven) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.


### GetPackaging

`func (o *Maven) GetPackaging() string`

GetPackaging returns the Packaging field if non-nil, zero value otherwise.

### GetPackagingOk

`func (o *Maven) GetPackagingOk() (*string, bool)`

GetPackagingOk returns a tuple with the Packaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackaging

`func (o *Maven) SetPackaging(v string)`

SetPackaging sets Packaging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


