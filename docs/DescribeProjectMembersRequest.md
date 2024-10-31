# DescribeProjectMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 
**ProjectId** | **int32** | 项目Id | 
**RoleId** | Pointer to **int32** | 用户组Id | [optional] 

## Methods

### NewDescribeProjectMembersRequest

`func NewDescribeProjectMembersRequest(pageNumber int32, pageSize int32, projectId int32, ) *DescribeProjectMembersRequest`

NewDescribeProjectMembersRequest instantiates a new DescribeProjectMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectMembersRequestWithDefaults

`func NewDescribeProjectMembersRequestWithDefaults() *DescribeProjectMembersRequest`

NewDescribeProjectMembersRequestWithDefaults instantiates a new DescribeProjectMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeProjectMembersRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProjectMembersRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProjectMembersRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeProjectMembersRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProjectMembersRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProjectMembersRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetProjectId

`func (o *DescribeProjectMembersRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeProjectMembersRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeProjectMembersRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRoleId

`func (o *DescribeProjectMembersRequest) GetRoleId() int32`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *DescribeProjectMembersRequest) GetRoleIdOk() (*int32, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *DescribeProjectMembersRequest) SetRoleId(v int32)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *DescribeProjectMembersRequest) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


