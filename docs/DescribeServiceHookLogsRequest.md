# DescribeServiceHookLogsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Service Hook 编号 | 
**PageNumber** | **int64** | 分页页码 | 
**PageSize** | **int64** | 分页大小 | 
**ProjectId** | **int64** | 项目编号 | 
**TargetType** | Pointer to **string** | 目标数据类型：PROJECT,SPACE_NODE,PROGRAM,默认PROJECT | [optional] 

## Methods

### NewDescribeServiceHookLogsRequest

`func NewDescribeServiceHookLogsRequest(id string, pageNumber int64, pageSize int64, projectId int64, ) *DescribeServiceHookLogsRequest`

NewDescribeServiceHookLogsRequest instantiates a new DescribeServiceHookLogsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeServiceHookLogsRequestWithDefaults

`func NewDescribeServiceHookLogsRequestWithDefaults() *DescribeServiceHookLogsRequest`

NewDescribeServiceHookLogsRequestWithDefaults instantiates a new DescribeServiceHookLogsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DescribeServiceHookLogsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeServiceHookLogsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeServiceHookLogsRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPageNumber

`func (o *DescribeServiceHookLogsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeServiceHookLogsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeServiceHookLogsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeServiceHookLogsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeServiceHookLogsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeServiceHookLogsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetProjectId

`func (o *DescribeServiceHookLogsRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeServiceHookLogsRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeServiceHookLogsRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetTargetType

`func (o *DescribeServiceHookLogsRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *DescribeServiceHookLogsRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *DescribeServiceHookLogsRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *DescribeServiceHookLogsRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


