# OpenApiReleaseListDataWithPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**List** | Pointer to [**[]OpenApiRelease**](OpenApiRelease.md) | 版本信息列表 | [optional] 
**PageNumber** | Pointer to **NullableInt64** | 页码 | [optional] 
**PageSize** | Pointer to **NullableInt64** | 页数 | [optional] 
**TotalCount** | Pointer to **NullableInt64** | 数据总条数 | [optional] 
**TotalPage** | Pointer to **NullableInt64** | 总页数 | [optional] 

## Methods

### NewOpenApiReleaseListDataWithPage

`func NewOpenApiReleaseListDataWithPage() *OpenApiReleaseListDataWithPage`

NewOpenApiReleaseListDataWithPage instantiates a new OpenApiReleaseListDataWithPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenApiReleaseListDataWithPageWithDefaults

`func NewOpenApiReleaseListDataWithPageWithDefaults() *OpenApiReleaseListDataWithPage`

NewOpenApiReleaseListDataWithPageWithDefaults instantiates a new OpenApiReleaseListDataWithPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetList

`func (o *OpenApiReleaseListDataWithPage) GetList() []OpenApiRelease`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *OpenApiReleaseListDataWithPage) GetListOk() (*[]OpenApiRelease, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *OpenApiReleaseListDataWithPage) SetList(v []OpenApiRelease)`

SetList sets List field to given value.

### HasList

`func (o *OpenApiReleaseListDataWithPage) HasList() bool`

HasList returns a boolean if a field has been set.

### SetListNil

`func (o *OpenApiReleaseListDataWithPage) SetListNil(b bool)`

 SetListNil sets the value for List to be an explicit nil

### UnsetList
`func (o *OpenApiReleaseListDataWithPage) UnsetList()`

UnsetList ensures that no value is present for List, not even an explicit nil
### GetPageNumber

`func (o *OpenApiReleaseListDataWithPage) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *OpenApiReleaseListDataWithPage) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *OpenApiReleaseListDataWithPage) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *OpenApiReleaseListDataWithPage) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### SetPageNumberNil

`func (o *OpenApiReleaseListDataWithPage) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *OpenApiReleaseListDataWithPage) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageSize

`func (o *OpenApiReleaseListDataWithPage) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OpenApiReleaseListDataWithPage) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OpenApiReleaseListDataWithPage) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *OpenApiReleaseListDataWithPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### SetPageSizeNil

`func (o *OpenApiReleaseListDataWithPage) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *OpenApiReleaseListDataWithPage) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetTotalCount

`func (o *OpenApiReleaseListDataWithPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OpenApiReleaseListDataWithPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OpenApiReleaseListDataWithPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OpenApiReleaseListDataWithPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *OpenApiReleaseListDataWithPage) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *OpenApiReleaseListDataWithPage) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
### GetTotalPage

`func (o *OpenApiReleaseListDataWithPage) GetTotalPage() int64`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *OpenApiReleaseListDataWithPage) GetTotalPageOk() (*int64, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *OpenApiReleaseListDataWithPage) SetTotalPage(v int64)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *OpenApiReleaseListDataWithPage) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### SetTotalPageNil

`func (o *OpenApiReleaseListDataWithPage) SetTotalPageNil(b bool)`

 SetTotalPageNil sets the value for TotalPage to be an explicit nil

### UnsetTotalPage
`func (o *OpenApiReleaseListDataWithPage) UnsetTotalPage()`

UnsetTotalPage ensures that no value is present for TotalPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


