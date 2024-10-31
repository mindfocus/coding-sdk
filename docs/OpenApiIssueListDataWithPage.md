# OpenApiIssueListDataWithPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]IssueListData**](IssueListData.md) | 所有事项数据 | [optional] 
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 分页的大小 | [optional] 
**TotalCount** | Pointer to **int64** | 所有行数 | [optional] 
**TotalPage** | Pointer to **int64** | 全部页 | [optional] 

## Methods

### NewOpenApiIssueListDataWithPage

`func NewOpenApiIssueListDataWithPage() *OpenApiIssueListDataWithPage`

NewOpenApiIssueListDataWithPage instantiates a new OpenApiIssueListDataWithPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiIssueListDataWithPageWithDefaults

`func NewOpenApiIssueListDataWithPageWithDefaults() *OpenApiIssueListDataWithPage`

NewOpenApiIssueListDataWithPageWithDefaults instantiates a new OpenApiIssueListDataWithPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *OpenApiIssueListDataWithPage) GetList() []IssueListData`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *OpenApiIssueListDataWithPage) GetListOk() (*[]IssueListData, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *OpenApiIssueListDataWithPage) SetList(v []IssueListData)`

SetList sets List field to given value.

### HasList

`func (o *OpenApiIssueListDataWithPage) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPageNumber

`func (o *OpenApiIssueListDataWithPage) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *OpenApiIssueListDataWithPage) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *OpenApiIssueListDataWithPage) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *OpenApiIssueListDataWithPage) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *OpenApiIssueListDataWithPage) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OpenApiIssueListDataWithPage) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OpenApiIssueListDataWithPage) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OpenApiIssueListDataWithPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *OpenApiIssueListDataWithPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpenApiIssueListDataWithPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpenApiIssueListDataWithPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OpenApiIssueListDataWithPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPage

`func (o *OpenApiIssueListDataWithPage) GetTotalPage() int64`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *OpenApiIssueListDataWithPage) GetTotalPageOk() (*int64, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *OpenApiIssueListDataWithPage) SetTotalPage(v int64)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *OpenApiIssueListDataWithPage) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


