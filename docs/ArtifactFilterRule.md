# ArtifactFilterRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactType** | Pointer to **[]int64** | 制品类型筛选（1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan;10-cocoapods;11-rpm） | [optional] 
**Package** | Pointer to [**[]ArtifactFilterRuleDetail**](ArtifactFilterRuleDetail.md) | 包筛选 | [optional] 
**PackageVersion** | Pointer to [**[]ArtifactFilterRuleDetail**](ArtifactFilterRuleDetail.md) | 版本筛选 | [optional] 
**ProjectName** | Pointer to **[]string** | 项目筛选 | [optional] 
**Repository** | Pointer to **[]string** | 仓库筛选 | [optional] 

## Methods

### NewArtifactFilterRule

`func NewArtifactFilterRule() *ArtifactFilterRule`

NewArtifactFilterRule instantiates a new ArtifactFilterRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactFilterRuleWithDefaults

`func NewArtifactFilterRuleWithDefaults() *ArtifactFilterRule`

NewArtifactFilterRuleWithDefaults instantiates a new ArtifactFilterRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactType

`func (o *ArtifactFilterRule) GetArtifactType() []int64`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ArtifactFilterRule) GetArtifactTypeOk() (*[]int64, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ArtifactFilterRule) SetArtifactType(v []int64)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *ArtifactFilterRule) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetPackage

`func (o *ArtifactFilterRule) GetPackage() []ArtifactFilterRuleDetail`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *ArtifactFilterRule) GetPackageOk() (*[]ArtifactFilterRuleDetail, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *ArtifactFilterRule) SetPackage(v []ArtifactFilterRuleDetail)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *ArtifactFilterRule) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetPackageVersion

`func (o *ArtifactFilterRule) GetPackageVersion() []ArtifactFilterRuleDetail`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *ArtifactFilterRule) GetPackageVersionOk() (*[]ArtifactFilterRuleDetail, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *ArtifactFilterRule) SetPackageVersion(v []ArtifactFilterRuleDetail)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *ArtifactFilterRule) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetProjectName

`func (o *ArtifactFilterRule) GetProjectName() []string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ArtifactFilterRule) GetProjectNameOk() (*[]string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ArtifactFilterRule) SetProjectName(v []string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ArtifactFilterRule) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetRepository

`func (o *ArtifactFilterRule) GetRepository() []string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ArtifactFilterRule) GetRepositoryOk() (*[]string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ArtifactFilterRule) SetRepository(v []string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ArtifactFilterRule) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


