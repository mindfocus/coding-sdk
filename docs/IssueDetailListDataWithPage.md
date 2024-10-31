# IssueDetailListDataWithPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]IssueDetail**](IssueDetail.md) | 事项列表  | [optional] 
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 页数 | [optional] 
**TotalCount** | Pointer to **int64** | 总共几条事项  | [optional] 
**TotalPage** | Pointer to **int64** | 总共几页  | [optional] 

## Methods

### NewIssueDetailListDataWithPage

`func NewIssueDetailListDataWithPage() *IssueDetailListDataWithPage`

NewIssueDetailListDataWithPage instantiates a new IssueDetailListDataWithPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueDetailListDataWithPageWithDefaults

`func NewIssueDetailListDataWithPageWithDefaults() *IssueDetailListDataWithPage`

NewIssueDetailListDataWithPageWithDefaults instantiates a new IssueDetailListDataWithPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *IssueDetailListDataWithPage) GetList() []IssueDetail`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *IssueDetailListDataWithPage) GetListOk() (*[]IssueDetail, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *IssueDetailListDataWithPage) SetList(v []IssueDetail)`

SetList sets List field to given value.

### HasList

`func (o *IssueDetailListDataWithPage) HasList() bool`

HasList returns a boolean if a field has been set.

### GetPageNumber

`func (o *IssueDetailListDataWithPage) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *IssueDetailListDataWithPage) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *IssueDetailListDataWithPage) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *IssueDetailListDataWithPage) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *IssueDetailListDataWithPage) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *IssueDetailListDataWithPage) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *IssueDetailListDataWithPage) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *IssueDetailListDataWithPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *IssueDetailListDataWithPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *IssueDetailListDataWithPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *IssueDetailListDataWithPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *IssueDetailListDataWithPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetTotalPage

`func (o *IssueDetailListDataWithPage) GetTotalPage() int64`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *IssueDetailListDataWithPage) GetTotalPageOk() (*int64, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *IssueDetailListDataWithPage) SetTotalPage(v int64)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *IssueDetailListDataWithPage) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


