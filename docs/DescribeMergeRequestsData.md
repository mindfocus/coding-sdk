# DescribeMergeRequestsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]MergeRequestData**](MergeRequestData.md) | 合并请求列表 | [optional] 
**PageNumber** | Pointer to **int64** | 页数 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 | [optional] 
**TotalPage** | Pointer to **int64** | 总页数 | [optional] 
**TotalRow** | Pointer to **int64** | 总条数 | [optional] 

## Methods

### NewDescribeMergeRequestsData

`func NewDescribeMergeRequestsData() *DescribeMergeRequestsData`

NewDescribeMergeRequestsData instantiates a new DescribeMergeRequestsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeRequestsDataWithDefaults

`func NewDescribeMergeRequestsDataWithDefaults() *DescribeMergeRequestsData`

NewDescribeMergeRequestsDataWithDefaults instantiates a new DescribeMergeRequestsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *DescribeMergeRequestsData) GetList() []MergeRequestData`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *DescribeMergeRequestsData) GetListOk() (*[]MergeRequestData, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *DescribeMergeRequestsData) SetList(v []MergeRequestData)`

SetList sets List field to given value.

### HasList

`func (o *DescribeMergeRequestsData) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *DescribeMergeRequestsData) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *DescribeMergeRequestsData) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetPageNumber

`func (o *DescribeMergeRequestsData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeMergeRequestsData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeMergeRequestsData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeMergeRequestsData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeMergeRequestsData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeMergeRequestsData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeMergeRequestsData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeMergeRequestsData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPage

`func (o *DescribeMergeRequestsData) GetTotalPage() int64`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *DescribeMergeRequestsData) GetTotalPageOk() (*int64, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *DescribeMergeRequestsData) SetTotalPage(v int64)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *DescribeMergeRequestsData) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### GetTotalRow

`func (o *DescribeMergeRequestsData) GetTotalRow() int64`

GetTotalRow returns the TotalRow field if non-nil, zero value otherwise.

### GetTotalRowOk

`func (o *DescribeMergeRequestsData) GetTotalRowOk() (*int64, bool)`

GetTotalRowOk returns a tuple with the TotalRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRow

`func (o *DescribeMergeRequestsData) SetTotalRow(v int64)`

SetTotalRow sets TotalRow field to given value.

### HasTotalRow

`func (o *DescribeMergeRequestsData) HasTotalRow() bool`

HasTotalRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


