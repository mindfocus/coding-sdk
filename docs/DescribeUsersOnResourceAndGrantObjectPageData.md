# DescribeUsersOnResourceAndGrantObjectPageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | **NullableInt64** | 请求页数 | 
**PageSize** | **NullableInt64** | 请求条数 | 
**TotalCount** | **NullableInt64** | 总条数 | 
**UserDataList** | Pointer to [**[]UserData**](UserData.md) | 用户信息列表 | [optional] 

## Methods

### NewDescribeUsersOnResourceAndGrantObjectPageData

`func NewDescribeUsersOnResourceAndGrantObjectPageData(pageNumber NullableInt64, pageSize NullableInt64, totalCount NullableInt64, ) *DescribeUsersOnResourceAndGrantObjectPageData`

NewDescribeUsersOnResourceAndGrantObjectPageData instantiates a new DescribeUsersOnResourceAndGrantObjectPageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUsersOnResourceAndGrantObjectPageDataWithDefaults

`func NewDescribeUsersOnResourceAndGrantObjectPageDataWithDefaults() *DescribeUsersOnResourceAndGrantObjectPageData`

NewDescribeUsersOnResourceAndGrantObjectPageDataWithDefaults instantiates a new DescribeUsersOnResourceAndGrantObjectPageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### SetPageNumberNil

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetPageNumberNil(b bool)`

 SetPageNumberNil sets the value for PageNumber to be an explicit nil

### UnsetPageNumber
`func (o *DescribeUsersOnResourceAndGrantObjectPageData) UnsetPageNumber()`

UnsetPageNumber ensures that no value is present for PageNumber, not even an explicit nil
### GetPageSize

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### SetPageSizeNil

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *DescribeUsersOnResourceAndGrantObjectPageData) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil
### GetTotalCount

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### SetTotalCountNil

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *DescribeUsersOnResourceAndGrantObjectPageData) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
### GetUserDataList

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetUserDataList() []UserData`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) GetUserDataListOk() (*[]UserData, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetUserDataList(v []UserData)`

SetUserDataList sets UserDataList field to given value.

### HasUserDataList

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) HasUserDataList() bool`

HasUserDataList returns a boolean if a field has been set.

### SetUserDataListNil

`func (o *DescribeUsersOnResourceAndGrantObjectPageData) SetUserDataListNil(b bool)`

 SetUserDataListNil sets the value for UserDataList to be an explicit nil

### UnsetUserDataList
`func (o *DescribeUsersOnResourceAndGrantObjectPageData) UnsetUserDataList()`

UnsetUserDataList ensures that no value is present for UserDataList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


