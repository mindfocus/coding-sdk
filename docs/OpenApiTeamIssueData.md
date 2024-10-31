# OpenApiTeamIssueData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueList** | Pointer to [**[]OpenApiWorkbenchIssue**](OpenApiWorkbenchIssue.md) | 工作台事项 | [optional] 
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 分页大小 | [optional] 
**TotalCount** | Pointer to **int64** | 所有数据条目 | [optional] 
**TotalPage** | Pointer to **int64** | 总页数 | [optional] 

## Methods

### NewOpenApiTeamIssueData

`func NewOpenApiTeamIssueData() *OpenApiTeamIssueData`

NewOpenApiTeamIssueData instantiates a new OpenApiTeamIssueData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiTeamIssueDataWithDefaults

`func NewOpenApiTeamIssueDataWithDefaults() *OpenApiTeamIssueData`

NewOpenApiTeamIssueDataWithDefaults instantiates a new OpenApiTeamIssueData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueList

`func (o *OpenApiTeamIssueData) GetIssueList() []OpenApiWorkbenchIssue`

GetIssueList returns the IssueList field if non-nil, zero value otherwise.

### GetIssueListOk

`func (o *OpenApiTeamIssueData) GetIssueListOk() (*[]OpenApiWorkbenchIssue, bool)`

GetIssueListOk returns a tuple with the IssueList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueList

`func (o *OpenApiTeamIssueData) SetIssueList(v []OpenApiWorkbenchIssue)`

SetIssueList sets IssueList field to given value.

### HasIssueList

`func (o *OpenApiTeamIssueData) HasIssueList() bool`

HasIssueList returns a boolean if a field has been set.

### GetPageNumber

`func (o *OpenApiTeamIssueData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *OpenApiTeamIssueData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *OpenApiTeamIssueData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *OpenApiTeamIssueData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *OpenApiTeamIssueData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OpenApiTeamIssueData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OpenApiTeamIssueData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OpenApiTeamIssueData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *OpenApiTeamIssueData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpenApiTeamIssueData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpenApiTeamIssueData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OpenApiTeamIssueData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPage

`func (o *OpenApiTeamIssueData) GetTotalPage() int64`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *OpenApiTeamIssueData) GetTotalPageOk() (*int64, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *OpenApiTeamIssueData) SetTotalPage(v int64)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *OpenApiTeamIssueData) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


