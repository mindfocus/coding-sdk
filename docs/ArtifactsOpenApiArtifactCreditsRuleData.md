# ArtifactsOpenApiArtifactCreditsRuleData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | 制品版本 | [default to ""]
**ArtifactType** | **int32** | 制品的类型(1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan,10-cocoapods,11-rpm,12-Go) | 
**PkgNameAlgorithm** | **string** | 制品名称计算规则：1-EQUAL(等于)，2-REGEX(正则表达式) | [default to ""]
**PkgName** | **string** | 制品名称 | [default to ""]
**VersionAlgorithm** | **string** | 制品版本计算规则：1-EQUAL(等于)，2-REGEX(正则表达式) | [default to ""]

## Methods

### NewArtifactsOpenApiArtifactCreditsRuleData

`func NewArtifactsOpenApiArtifactCreditsRuleData(version string, artifactType int32, pkgNameAlgorithm string, pkgName string, versionAlgorithm string, ) *ArtifactsOpenApiArtifactCreditsRuleData`

NewArtifactsOpenApiArtifactCreditsRuleData instantiates a new ArtifactsOpenApiArtifactCreditsRuleData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiArtifactCreditsRuleDataWithDefaults

`func NewArtifactsOpenApiArtifactCreditsRuleDataWithDefaults() *ArtifactsOpenApiArtifactCreditsRuleData`

NewArtifactsOpenApiArtifactCreditsRuleDataWithDefaults instantiates a new ArtifactsOpenApiArtifactCreditsRuleData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetArtifactType

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetArtifactType() int32`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetArtifactTypeOk() (*int32, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetArtifactType(v int32)`

SetArtifactType sets ArtifactType field to given value.


### GetPkgNameAlgorithm

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgNameAlgorithm() string`

GetPkgNameAlgorithm returns the PkgNameAlgorithm field if non-nil, zero value otherwise.

### GetPkgNameAlgorithmOk

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgNameAlgorithmOk() (*string, bool)`

GetPkgNameAlgorithmOk returns a tuple with the PkgNameAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgNameAlgorithm

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetPkgNameAlgorithm(v string)`

SetPkgNameAlgorithm sets PkgNameAlgorithm field to given value.


### GetPkgName

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgName() string`

GetPkgName returns the PkgName field if non-nil, zero value otherwise.

### GetPkgNameOk

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgNameOk() (*string, bool)`

GetPkgNameOk returns a tuple with the PkgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgName

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetPkgName(v string)`

SetPkgName sets PkgName field to given value.


### GetVersionAlgorithm

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersionAlgorithm() string`

GetVersionAlgorithm returns the VersionAlgorithm field if non-nil, zero value otherwise.

### GetVersionAlgorithmOk

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersionAlgorithmOk() (*string, bool)`

GetVersionAlgorithmOk returns a tuple with the VersionAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionAlgorithm

`func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetVersionAlgorithm(v string)`

SetVersionAlgorithm sets VersionAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


