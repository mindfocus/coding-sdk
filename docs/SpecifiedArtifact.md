# SpecifiedArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageName** | Pointer to **string** | 制品包名称（必须配合 VersionName 使用） | [optional] [default to ""]
**VersionName** | Pointer to **string** | 制品版本（必须配合 PackageName 使用） | [optional] [default to ""]

## Methods

### NewSpecifiedArtifact

`func NewSpecifiedArtifact() *SpecifiedArtifact`

NewSpecifiedArtifact instantiates a new SpecifiedArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecifiedArtifactWithDefaults

`func NewSpecifiedArtifactWithDefaults() *SpecifiedArtifact`

NewSpecifiedArtifactWithDefaults instantiates a new SpecifiedArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageName

`func (o *SpecifiedArtifact) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *SpecifiedArtifact) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *SpecifiedArtifact) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *SpecifiedArtifact) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetVersionName

`func (o *SpecifiedArtifact) GetVersionName() string`

GetVersionName returns the VersionName field if non-nil, zero value otherwise.

### GetVersionNameOk

`func (o *SpecifiedArtifact) GetVersionNameOk() (*string, bool)`

GetVersionNameOk returns a tuple with the VersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionName

`func (o *SpecifiedArtifact) SetVersionName(v string)`

SetVersionName sets VersionName field to given value.

### HasVersionName

`func (o *SpecifiedArtifact) HasVersionName() bool`

HasVersionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


