# DescribeArtifactPackageListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackagePrefix** | Pointer to **string** | 包名前缀 | [optional] 
**PageNumber** | Pointer to **int32** | 页码，默认：1 | [optional] 
**PageSize** | Pointer to **int32** | 每页展示数量，默认：10 | [optional] 
**ProjectId** | **int32** | 项目 ID | 
**Repository** | **string** | 仓库名称 | 

## Methods

### NewDescribeArtifactPackageListRequest

`func NewDescribeArtifactPackageListRequest(projectId int32, repository string, ) *DescribeArtifactPackageListRequest`

NewDescribeArtifactPackageListRequest instantiates a new DescribeArtifactPackageListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactPackageListRequestWithDefaults

`func NewDescribeArtifactPackageListRequestWithDefaults() *DescribeArtifactPackageListRequest`

NewDescribeArtifactPackageListRequestWithDefaults instantiates a new DescribeArtifactPackageListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackagePrefix

`func (o *DescribeArtifactPackageListRequest) GetPackagePrefix() string`

GetPackagePrefix returns the PackagePrefix field if non-nil, zero value otherwise.

### GetPackagePrefixOk

`func (o *DescribeArtifactPackageListRequest) GetPackagePrefixOk() (*string, bool)`

GetPackagePrefixOk returns a tuple with the PackagePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagePrefix

`func (o *DescribeArtifactPackageListRequest) SetPackagePrefix(v string)`

SetPackagePrefix sets PackagePrefix field to given value.

### HasPackagePrefix

`func (o *DescribeArtifactPackageListRequest) HasPackagePrefix() bool`

HasPackagePrefix returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeArtifactPackageListRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeArtifactPackageListRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeArtifactPackageListRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeArtifactPackageListRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeArtifactPackageListRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeArtifactPackageListRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeArtifactPackageListRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeArtifactPackageListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeArtifactPackageListRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeArtifactPackageListRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeArtifactPackageListRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepository

`func (o *DescribeArtifactPackageListRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DescribeArtifactPackageListRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DescribeArtifactPackageListRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


