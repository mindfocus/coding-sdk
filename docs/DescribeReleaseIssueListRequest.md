# DescribeReleaseIssueListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignees** | Pointer to **[]int64** | 处理人ID数组 | [optional] 
**IssueTypeIds** | Pointer to **[]int64** | 事项类型ID数组  | [optional] 
**IssueTypes** | Pointer to **[]string** | 事项类型数组 | [optional] 
**Keywords** | Pointer to **string** | 关键字 | [optional] 
**Page** | Pointer to **int64** | 分页查询中的页数,page从1开始计数  | [optional] 
**PageSize** | Pointer to **int64** | 分页查询中每页的大小  | [optional] 
**ProjectName** | **string** | 项目名称 | 
**ReleaseCode** | **int64** | 页面上版本ID | 
**ShowImageOutUrl** | Pointer to **bool** | 是否展示描述中外部可访问的地址 | [optional] 
**ShowSubIssues** | Pointer to **bool** | 是否显示字事项，和页面开关对应  | [optional] 
**SortBy** | Pointer to **string** | 排序，取值如\&quot;ID:ASC\&quot;  | [optional] 
**StatusTypes** | Pointer to **[]int64** | 事项状态类型数组  | [optional] 
**Watchers** | Pointer to **[]int64** | 关注人ID数组  | [optional] 

## Methods

### NewDescribeReleaseIssueListRequest

`func NewDescribeReleaseIssueListRequest(projectName string, releaseCode int64, ) *DescribeReleaseIssueListRequest`

NewDescribeReleaseIssueListRequest instantiates a new DescribeReleaseIssueListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeReleaseIssueListRequestWithDefaults

`func NewDescribeReleaseIssueListRequestWithDefaults() *DescribeReleaseIssueListRequest`

NewDescribeReleaseIssueListRequestWithDefaults instantiates a new DescribeReleaseIssueListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignees

`func (o *DescribeReleaseIssueListRequest) GetAssignees() []int64`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *DescribeReleaseIssueListRequest) GetAssigneesOk() (*[]int64, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *DescribeReleaseIssueListRequest) SetAssignees(v []int64)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *DescribeReleaseIssueListRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.

### GetIssueTypeIds

`func (o *DescribeReleaseIssueListRequest) GetIssueTypeIds() []int64`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *DescribeReleaseIssueListRequest) GetIssueTypeIdsOk() (*[]int64, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *DescribeReleaseIssueListRequest) SetIssueTypeIds(v []int64)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *DescribeReleaseIssueListRequest) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetIssueTypes

`func (o *DescribeReleaseIssueListRequest) GetIssueTypes() []string`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *DescribeReleaseIssueListRequest) GetIssueTypesOk() (*[]string, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *DescribeReleaseIssueListRequest) SetIssueTypes(v []string)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *DescribeReleaseIssueListRequest) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetKeywords

`func (o *DescribeReleaseIssueListRequest) GetKeywords() string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *DescribeReleaseIssueListRequest) GetKeywordsOk() (*string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *DescribeReleaseIssueListRequest) SetKeywords(v string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *DescribeReleaseIssueListRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetPage

`func (o *DescribeReleaseIssueListRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DescribeReleaseIssueListRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DescribeReleaseIssueListRequest) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *DescribeReleaseIssueListRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeReleaseIssueListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeReleaseIssueListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeReleaseIssueListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeReleaseIssueListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProjectName

`func (o *DescribeReleaseIssueListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeReleaseIssueListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeReleaseIssueListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReleaseCode

`func (o *DescribeReleaseIssueListRequest) GetReleaseCode() int64`

GetReleaseCode returns the ReleaseCode field if non-nil, zero value otherwise.

### GetReleaseCodeOk

`func (o *DescribeReleaseIssueListRequest) GetReleaseCodeOk() (*int64, bool)`

GetReleaseCodeOk returns a tuple with the ReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCode

`func (o *DescribeReleaseIssueListRequest) SetReleaseCode(v int64)`

SetReleaseCode sets ReleaseCode field to given value.


### GetShowImageOutUrl

`func (o *DescribeReleaseIssueListRequest) GetShowImageOutUrl() bool`

GetShowImageOutUrl returns the ShowImageOutUrl field if non-nil, zero value otherwise.

### GetShowImageOutUrlOk

`func (o *DescribeReleaseIssueListRequest) GetShowImageOutUrlOk() (*bool, bool)`

GetShowImageOutUrlOk returns a tuple with the ShowImageOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowImageOutUrl

`func (o *DescribeReleaseIssueListRequest) SetShowImageOutUrl(v bool)`

SetShowImageOutUrl sets ShowImageOutUrl field to given value.

### HasShowImageOutUrl

`func (o *DescribeReleaseIssueListRequest) HasShowImageOutUrl() bool`

HasShowImageOutUrl returns a boolean if a field has been set.

### GetShowSubIssues

`func (o *DescribeReleaseIssueListRequest) GetShowSubIssues() bool`

GetShowSubIssues returns the ShowSubIssues field if non-nil, zero value otherwise.

### GetShowSubIssuesOk

`func (o *DescribeReleaseIssueListRequest) GetShowSubIssuesOk() (*bool, bool)`

GetShowSubIssuesOk returns a tuple with the ShowSubIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSubIssues

`func (o *DescribeReleaseIssueListRequest) SetShowSubIssues(v bool)`

SetShowSubIssues sets ShowSubIssues field to given value.

### HasShowSubIssues

`func (o *DescribeReleaseIssueListRequest) HasShowSubIssues() bool`

HasShowSubIssues returns a boolean if a field has been set.

### GetSortBy

`func (o *DescribeReleaseIssueListRequest) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *DescribeReleaseIssueListRequest) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *DescribeReleaseIssueListRequest) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *DescribeReleaseIssueListRequest) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetStatusTypes

`func (o *DescribeReleaseIssueListRequest) GetStatusTypes() []int64`

GetStatusTypes returns the StatusTypes field if non-nil, zero value otherwise.

### GetStatusTypesOk

`func (o *DescribeReleaseIssueListRequest) GetStatusTypesOk() (*[]int64, bool)`

GetStatusTypesOk returns a tuple with the StatusTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTypes

`func (o *DescribeReleaseIssueListRequest) SetStatusTypes(v []int64)`

SetStatusTypes sets StatusTypes field to given value.

### HasStatusTypes

`func (o *DescribeReleaseIssueListRequest) HasStatusTypes() bool`

HasStatusTypes returns a boolean if a field has been set.

### GetWatchers

`func (o *DescribeReleaseIssueListRequest) GetWatchers() []int64`

GetWatchers returns the Watchers field if non-nil, zero value otherwise.

### GetWatchersOk

`func (o *DescribeReleaseIssueListRequest) GetWatchersOk() (*[]int64, bool)`

GetWatchersOk returns a tuple with the Watchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchers

`func (o *DescribeReleaseIssueListRequest) SetWatchers(v []int64)`

SetWatchers sets Watchers field to given value.

### HasWatchers

`func (o *DescribeReleaseIssueListRequest) HasWatchers() bool`

HasWatchers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


