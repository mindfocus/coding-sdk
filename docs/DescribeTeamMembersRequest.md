# DescribeTeamMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 
**ShowDepartment** | Pointer to **bool** | 是否展示部门 | [optional] [default to false]

## Methods

### NewDescribeTeamMembersRequest

`func NewDescribeTeamMembersRequest(pageNumber int32, pageSize int32, ) *DescribeTeamMembersRequest`

NewDescribeTeamMembersRequest instantiates a new DescribeTeamMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamMembersRequestWithDefaults

`func NewDescribeTeamMembersRequestWithDefaults() *DescribeTeamMembersRequest`

NewDescribeTeamMembersRequestWithDefaults instantiates a new DescribeTeamMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeTeamMembersRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeTeamMembersRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeTeamMembersRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeTeamMembersRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeTeamMembersRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeTeamMembersRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetShowDepartment

`func (o *DescribeTeamMembersRequest) GetShowDepartment() bool`

GetShowDepartment returns the ShowDepartment field if non-nil, zero value otherwise.

### GetShowDepartmentOk

`func (o *DescribeTeamMembersRequest) GetShowDepartmentOk() (*bool, bool)`

GetShowDepartmentOk returns a tuple with the ShowDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDepartment

`func (o *DescribeTeamMembersRequest) SetShowDepartment(v bool)`

SetShowDepartment sets ShowDepartment field to given value.

### HasShowDepartment

`func (o *DescribeTeamMembersRequest) HasShowDepartment() bool`

HasShowDepartment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


