# PageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **NullableInt64** | 页码数量 | [optional] 
**PageSize** | Pointer to **NullableInt64** | 每页大小 | [optional] 
**TotalPage** | Pointer to **NullableInt64** | 页码总数量 | [optional] 
**TotalRow** | Pointer to **NullableInt64** | 数据总行数 | [optional] 

## Methods

### NewPageInfo

`func NewPageInfo() *PageInfo`

NewPageInfo instantiates a new PageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageInfoWithDefaults

`func NewPageInfoWithDefaults() *PageInfo`

NewPageInfoWithDefaults instantiates a new PageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *PageInfo) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *PageInfo) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *PageInfo) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *PageInfo) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### SetPageNumberNil

`func (o *PageInfo) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *PageInfo) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageSize

`func (o *PageInfo) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PageInfo) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PageInfo) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PageInfo) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *PageInfo) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *PageInfo) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetTotalPage

`func (o *PageInfo) GetTotalPage() int64`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *PageInfo) GetTotalPageOk() (*int64, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *PageInfo) SetTotalPage(v int64)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *PageInfo) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### SetTotalPageNil

`func (o *PageInfo) SetTotalPageNil(b bool)`

 SetTotalPageNil sets the value for TotalPage to be an explicit nil

### UnsetTotalPage
`func (o *PageInfo) UnsetTotalPage()`

UnsetTotalPage ensures that no value is present for TotalPage, not even an explicit nil
### GetTotalRow

`func (o *PageInfo) GetTotalRow() int64`

GetTotalRow returns the TotalRow field if non-nil, zero value otherwise.

### GetTotalRowOk

`func (o *PageInfo) GetTotalRowOk() (*int64, bool)`

GetTotalRowOk returns a tuple with the TotalRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRow

`func (o *PageInfo) SetTotalRow(v int64)`

SetTotalRow sets TotalRow field to given value.

### HasTotalRow

`func (o *PageInfo) HasTotalRow() bool`

HasTotalRow returns a boolean if a field has been set.

### SetTotalRowNil

`func (o *PageInfo) SetTotalRowNil(b bool)`

 SetTotalRowNil sets the value for TotalRow to be an explicit nil

### UnsetTotalRow
`func (o *PageInfo) UnsetTotalRow()`

UnsetTotalRow ensures that no value is present for TotalRow, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


