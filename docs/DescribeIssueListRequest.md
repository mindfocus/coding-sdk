# DescribeIssueListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]IssueCondition**](IssueCondition.md) | 筛选条件列表 | [optional] 
**IssueType** | **string** | 事项类型  ALL - 全部事项  DEFECT - 缺陷  REQUIREMENT - 需求  MISSION - 任务  EPIC - 史诗 | 
**Limit** | Pointer to **int64** | 限制数目，默认 20 | [optional] 
**Offset** | Pointer to **int64** | 偏移量，默认 0 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SortKey** | Pointer to **string** | 排序字段，默认：CODE  可选值：STATUS, CREATED_AT, PRIORITY, UPDATED_AT, DUE_DATE, CODE, JOIN_ITERATION_AT, STATUS_TYPE, ASSIGNEE, PROJECT_ID, ISSUE_STATUS_ID, ISSUE_ITERATION_SORT, ISSUE_ROADMAP_SORT, PARENT_ID, COMPLETED_AT | [optional] 
**SortValue** | Pointer to **string** | 排序方式  DESC - 倒序（默认值）  ASC - 正序 | [optional] 

## Methods

### NewDescribeIssueListRequest

`func NewDescribeIssueListRequest(issueType string, projectName string, ) *DescribeIssueListRequest`

NewDescribeIssueListRequest instantiates a new DescribeIssueListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueListRequestWithDefaults

`func NewDescribeIssueListRequestWithDefaults() *DescribeIssueListRequest`

NewDescribeIssueListRequestWithDefaults instantiates a new DescribeIssueListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DescribeIssueListRequest) GetConditions() []IssueCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DescribeIssueListRequest) GetConditionsOk() (*[]IssueCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DescribeIssueListRequest) SetConditions(v []IssueCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DescribeIssueListRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetIssueType

`func (o *DescribeIssueListRequest) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *DescribeIssueListRequest) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *DescribeIssueListRequest) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetLimit

`func (o *DescribeIssueListRequest) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DescribeIssueListRequest) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DescribeIssueListRequest) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DescribeIssueListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *DescribeIssueListRequest) GetOffset() int64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DescribeIssueListRequest) GetOffsetOk() (*int64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DescribeIssueListRequest) SetOffset(v int64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DescribeIssueListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeIssueListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSortKey

`func (o *DescribeIssueListRequest) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *DescribeIssueListRequest) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *DescribeIssueListRequest) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *DescribeIssueListRequest) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortValue

`func (o *DescribeIssueListRequest) GetSortValue() string`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *DescribeIssueListRequest) GetSortValueOk() (*string, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *DescribeIssueListRequest) SetSortValue(v string)`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *DescribeIssueListRequest) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


