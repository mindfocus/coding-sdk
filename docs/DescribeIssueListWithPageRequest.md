# DescribeIssueListWithPageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**IssueCondition**](IssueCondition.md) |  | [optional] 
**ExcludeSubTask** | Pointer to **bool** | 是否展示子工作项 | [optional] 
**IssueType** | **string** | ALL - 全部事项 DEFECT - 缺陷 REQUIREMENT - 需求 MISSION - 任务 EPIC - 史诗 | 
**PageNumber** | **int64** | 页码 | 
**PageSize** | **int64** | 每页展示数，默认 20 ，最大值 500 | 
**ProjectName** | **string** | 项目名称 | 
**ShowSubIssues** | Pointer to **bool** | 是否展示子事项 | [optional] 
**SortKey** | Pointer to **string** | 排序字段，默认：CODE 可选值：STATUS, CREATED_AT, PRIORITY, UPDATED_AT, DUE_DATE, CODE, JOIN_ITERATION_AT, STATUS_TYPE, ASSIGNEE, PROJECT_ID, ISSUE_STATUS_ID, ISSUE_ITERATION_SORT, ISSUE_ROADMAP_SORT, PARENT_ID, COMPLETED_AT | [optional] 
**SortValue** | Pointer to **string** |  排序方式 DESC - 倒序（默认值） ASC - 正序 | [optional] 

## Methods

### NewDescribeIssueListWithPageRequest

`func NewDescribeIssueListWithPageRequest(issueType string, pageNumber int64, pageSize int64, projectName string, ) *DescribeIssueListWithPageRequest`

NewDescribeIssueListWithPageRequest instantiates a new DescribeIssueListWithPageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueListWithPageRequestWithDefaults

`func NewDescribeIssueListWithPageRequestWithDefaults() *DescribeIssueListWithPageRequest`

NewDescribeIssueListWithPageRequestWithDefaults instantiates a new DescribeIssueListWithPageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *DescribeIssueListWithPageRequest) GetConditions() IssueCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *DescribeIssueListWithPageRequest) GetConditionsOk() (*IssueCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *DescribeIssueListWithPageRequest) SetConditions(v IssueCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *DescribeIssueListWithPageRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetExcludeSubTask

`func (o *DescribeIssueListWithPageRequest) GetExcludeSubTask() bool`

GetExcludeSubTask returns the ExcludeSubTask field if non-nil, zero value otherwise.

### GetExcludeSubTaskOk

`func (o *DescribeIssueListWithPageRequest) GetExcludeSubTaskOk() (*bool, bool)`

GetExcludeSubTaskOk returns a tuple with the ExcludeSubTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSubTask

`func (o *DescribeIssueListWithPageRequest) SetExcludeSubTask(v bool)`

SetExcludeSubTask sets ExcludeSubTask field to given value.

### HasExcludeSubTask

`func (o *DescribeIssueListWithPageRequest) HasExcludeSubTask() bool`

HasExcludeSubTask returns a boolean if a field has been set.

### GetIssueType

`func (o *DescribeIssueListWithPageRequest) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *DescribeIssueListWithPageRequest) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *DescribeIssueListWithPageRequest) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetPageNumber

`func (o *DescribeIssueListWithPageRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeIssueListWithPageRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeIssueListWithPageRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeIssueListWithPageRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeIssueListWithPageRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeIssueListWithPageRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetProjectName

`func (o *DescribeIssueListWithPageRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueListWithPageRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueListWithPageRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetShowSubIssues

`func (o *DescribeIssueListWithPageRequest) GetShowSubIssues() bool`

GetShowSubIssues returns the ShowSubIssues field if non-nil, zero value otherwise.

### GetShowSubIssuesOk

`func (o *DescribeIssueListWithPageRequest) GetShowSubIssuesOk() (*bool, bool)`

GetShowSubIssuesOk returns a tuple with the ShowSubIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSubIssues

`func (o *DescribeIssueListWithPageRequest) SetShowSubIssues(v bool)`

SetShowSubIssues sets ShowSubIssues field to given value.

### HasShowSubIssues

`func (o *DescribeIssueListWithPageRequest) HasShowSubIssues() bool`

HasShowSubIssues returns a boolean if a field has been set.

### GetSortKey

`func (o *DescribeIssueListWithPageRequest) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *DescribeIssueListWithPageRequest) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *DescribeIssueListWithPageRequest) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *DescribeIssueListWithPageRequest) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.

### GetSortValue

`func (o *DescribeIssueListWithPageRequest) GetSortValue() string`

GetSortValue returns the SortValue field if non-nil, zero value otherwise.

### GetSortValueOk

`func (o *DescribeIssueListWithPageRequest) GetSortValueOk() (*string, bool)`

GetSortValueOk returns a tuple with the SortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortValue

`func (o *DescribeIssueListWithPageRequest) SetSortValue(v string)`

SetSortValue sets SortValue field to given value.

### HasSortValue

`func (o *DescribeIssueListWithPageRequest) HasSortValue() bool`

HasSortValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


