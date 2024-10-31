# DescribeIterationListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **[]int64** | 处理人 ID 列表 | [optional] 
**EndDate** | Pointer to **map[string]interface{}** | 通过结束时间过滤，时间格式：2020-12-12 | [optional] 
**Keywords** | Pointer to **string** | 通过关键字搜索 | [optional] 
**Limit** | Pointer to **int64** | 偏移量，默认 0 | [optional] 
**Offset** | Pointer to **int64** | 每页数量，默认 20 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**StartDate** | Pointer to **map[string]interface{}** | 通过开始时间过滤，时间格式：2020-12-12 | [optional] 
**Status** | Pointer to **[]string** | 迭代状态,  WAIT_PROCESS,PROCESSING,COMPLETED | [optional] 

## Methods

### NewDescribeIterationListRequest

`func NewDescribeIterationListRequest(projectName string, ) *DescribeIterationListRequest`

NewDescribeIterationListRequest instantiates a new DescribeIterationListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIterationListRequestWithDefaults

`func NewDescribeIterationListRequestWithDefaults() *DescribeIterationListRequest`

NewDescribeIterationListRequestWithDefaults instantiates a new DescribeIterationListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *DescribeIterationListRequest) GetAssignee() []int64`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *DescribeIterationListRequest) GetAssigneeOk() (*[]int64, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *DescribeIterationListRequest) SetAssignee(v []int64)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *DescribeIterationListRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetEndDate

`func (o *DescribeIterationListRequest) GetEndDate() map[string]interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeIterationListRequest) GetEndDateOk() (*map[string]interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeIterationListRequest) SetEndDate(v map[string]interface{})`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DescribeIterationListRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetKeywords

`func (o *DescribeIterationListRequest) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *DescribeIterationListRequest) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *DescribeIterationListRequest) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *DescribeIterationListRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetLimit

`func (o *DescribeIterationListRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DescribeIterationListRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DescribeIterationListRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DescribeIterationListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DescribeIterationListRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DescribeIterationListRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DescribeIterationListRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DescribeIterationListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeIterationListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIterationListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIterationListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetStartDate

`func (o *DescribeIterationListRequest) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeIterationListRequest) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeIterationListRequest) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribeIterationListRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetStatus

`func (o *DescribeIterationListRequest) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DescribeIterationListRequest) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DescribeIterationListRequest) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DescribeIterationListRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


