# DescribeTestListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | Pointer to **int32** | 优先级 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**RunId** | Pointer to **int32** | 测试计划 ID | [optional] 
**Status** | Pointer to **string** | 测试状态,UNTESTED:未测试,PASSED:通过,BLOCKED:阻塞,RETEST:重测,FAILED:失败 | [optional] 

## Methods

### NewDescribeTestListRequest

`func NewDescribeTestListRequest(projectName string, ) *DescribeTestListRequest`

NewDescribeTestListRequest instantiates a new DescribeTestListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTestListRequestWithDefaults

`func NewDescribeTestListRequestWithDefaults() *DescribeTestListRequest`

NewDescribeTestListRequestWithDefaults instantiates a new DescribeTestListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *DescribeTestListRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DescribeTestListRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DescribeTestListRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DescribeTestListRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeTestListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeTestListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeTestListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRunId

`func (o *DescribeTestListRequest) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *DescribeTestListRequest) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *DescribeTestListRequest) SetRunId(v int32)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *DescribeTestListRequest) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeTestListRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeTestListRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeTestListRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeTestListRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


