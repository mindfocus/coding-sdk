# DescribeArtifactRepositoryListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** | 页码，默认：1 | [optional] 
**PageSize** | Pointer to **int32** | 每页展示数量，默认：10 | [optional] 
**ProjectId** | **int32** | 项目 ID | 
**Type** | Pointer to **int64** | 仓库类型:1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan;10-cocoapods;11-rpm | [optional] 

## Methods

### NewDescribeArtifactRepositoryListRequest

`func NewDescribeArtifactRepositoryListRequest(projectId int32, ) *DescribeArtifactRepositoryListRequest`

NewDescribeArtifactRepositoryListRequest instantiates a new DescribeArtifactRepositoryListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactRepositoryListRequestWithDefaults

`func NewDescribeArtifactRepositoryListRequestWithDefaults() *DescribeArtifactRepositoryListRequest`

NewDescribeArtifactRepositoryListRequestWithDefaults instantiates a new DescribeArtifactRepositoryListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeArtifactRepositoryListRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeArtifactRepositoryListRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeArtifactRepositoryListRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeArtifactRepositoryListRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeArtifactRepositoryListRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeArtifactRepositoryListRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeArtifactRepositoryListRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeArtifactRepositoryListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeArtifactRepositoryListRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeArtifactRepositoryListRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeArtifactRepositoryListRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetType

`func (o *DescribeArtifactRepositoryListRequest) GetType() int64`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeArtifactRepositoryListRequest) GetTypeOk() (*int64, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeArtifactRepositoryListRequest) SetType(v int64)`

SetType sets Type field to given value.

### HasType

`func (o *DescribeArtifactRepositoryListRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


