# DescribeUsersByGroupIdResponsePageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | **NullableInt64** | 页码 | 
**PageSize** | **NullableInt64** | 每页条数 | 
**TotalCount** | **NullableInt64** | 总条数 | 
**UserDataList** | Pointer to [**[]UserData**](UserData.md) | 用户组列表 | [optional] 

## Methods

### NewDescribeUsersByGroupIdResponsePageData

`func NewDescribeUsersByGroupIdResponsePageData(pageNumber NullableInt64, pageSize NullableInt64, totalCount NullableInt64, ) *DescribeUsersByGroupIdResponsePageData`

NewDescribeUsersByGroupIdResponsePageData instantiates a new DescribeUsersByGroupIdResponsePageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersByGroupIdResponsePageDataWithDefaults

`func NewDescribeUsersByGroupIdResponsePageDataWithDefaults() *DescribeUsersByGroupIdResponsePageData`

NewDescribeUsersByGroupIdResponsePageDataWithDefaults instantiates a new DescribeUsersByGroupIdResponsePageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeUsersByGroupIdResponsePageData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeUsersByGroupIdResponsePageData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeUsersByGroupIdResponsePageData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### SetPageNumberNil

`func (o *DescribeUsersByGroupIdResponsePageData) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *DescribeUsersByGroupIdResponsePageData) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageSize

`func (o *DescribeUsersByGroupIdResponsePageData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeUsersByGroupIdResponsePageData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeUsersByGroupIdResponsePageData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### SetPageSizeNil

`func (o *DescribeUsersByGroupIdResponsePageData) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *DescribeUsersByGroupIdResponsePageData) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetTotalCount

`func (o *DescribeUsersByGroupIdResponsePageData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeUsersByGroupIdResponsePageData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeUsersByGroupIdResponsePageData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### SetTotalCountNil

`func (o *DescribeUsersByGroupIdResponsePageData) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *DescribeUsersByGroupIdResponsePageData) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
### GetUserDataList

`func (o *DescribeUsersByGroupIdResponsePageData) GetUserDataList() []UserData`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *DescribeUsersByGroupIdResponsePageData) GetUserDataListOk() (*[]UserData, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *DescribeUsersByGroupIdResponsePageData) SetUserDataList(v []UserData)`

SetUserDataList sets UserDataList field to given value.

### HasUserDataList

`func (o *DescribeUsersByGroupIdResponsePageData) HasUserDataList() bool`

HasUserDataList returns a boolean if a field has been set.

### SetUserDataListNil

`func (o *DescribeUsersByGroupIdResponsePageData) SetUserDataListNil(b bool)`

 SetUserDataListNil sets the value for UserDataList to be an explicit nil

### UnsetUserDataList
`func (o *DescribeUsersByGroupIdResponsePageData) UnsetUserDataList()`

UnsetUserDataList ensures that no value is present for UserDataList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


