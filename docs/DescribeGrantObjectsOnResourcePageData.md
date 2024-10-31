# DescribeGrantObjectsOnResourcePageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantObjectList** | Pointer to [**[]GrantObjectInfo**](GrantObjectInfo.md) | 授权资源列表 | [optional] 
**PageNumber** | **NullableInt64** | 请求页数 | 
**PageSize** | **NullableInt64** | 请求条数 | 
**TotalCount** | **NullableInt64** | 总条数 | 

## Methods

### NewDescribeGrantObjectsOnResourcePageData

`func NewDescribeGrantObjectsOnResourcePageData(pageNumber NullableInt64, pageSize NullableInt64, totalCount NullableInt64, ) *DescribeGrantObjectsOnResourcePageData`

NewDescribeGrantObjectsOnResourcePageData instantiates a new DescribeGrantObjectsOnResourcePageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGrantObjectsOnResourcePageDataWithDefaults

`func NewDescribeGrantObjectsOnResourcePageDataWithDefaults() *DescribeGrantObjectsOnResourcePageData`

NewDescribeGrantObjectsOnResourcePageDataWithDefaults instantiates a new DescribeGrantObjectsOnResourcePageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantObjectList

`func (o *DescribeGrantObjectsOnResourcePageData) GetGrantObjectList() []GrantObjectInfo`

GetGrantObjectList returns the GrantObjectList field if non-nil, zero value otherwise.

### GetGrantObjectListOk

`func (o *DescribeGrantObjectsOnResourcePageData) GetGrantObjectListOk() (*[]GrantObjectInfo, bool)`

GetGrantObjectListOk returns a tuple with the GrantObjectList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantObjectList

`func (o *DescribeGrantObjectsOnResourcePageData) SetGrantObjectList(v []GrantObjectInfo)`

SetGrantObjectList sets GrantObjectList field to given value.

### HasGrantObjectList

`func (o *DescribeGrantObjectsOnResourcePageData) HasGrantObjectList() bool`

HasGrantObjectList returns a boolean if a field has been set.

### SetGrantObjectListNil

`func (o *DescribeGrantObjectsOnResourcePageData) SetGrantObjectListNil(b bool)`

 SetGrantObjectListNil sets the value for GrantObjectList to be an explicit nil

### UnsetGrantObjectList
`func (o *DescribeGrantObjectsOnResourcePageData) UnsetGrantObjectList()`

UnsetGrantObjectList ensures that no value is present for GrantObjectList, not even an explicit nil
### GetPageNumber

`func (o *DescribeGrantObjectsOnResourcePageData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGrantObjectsOnResourcePageData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGrantObjectsOnResourcePageData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### SetPageNumberNil

`func (o *DescribeGrantObjectsOnResourcePageData) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *DescribeGrantObjectsOnResourcePageData) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageSize

`func (o *DescribeGrantObjectsOnResourcePageData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGrantObjectsOnResourcePageData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGrantObjectsOnResourcePageData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### SetPageSizeNil

`func (o *DescribeGrantObjectsOnResourcePageData) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *DescribeGrantObjectsOnResourcePageData) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetTotalCount

`func (o *DescribeGrantObjectsOnResourcePageData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeGrantObjectsOnResourcePageData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeGrantObjectsOnResourcePageData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### SetTotalCountNil

`func (o *DescribeGrantObjectsOnResourcePageData) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *DescribeGrantObjectsOnResourcePageData) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


