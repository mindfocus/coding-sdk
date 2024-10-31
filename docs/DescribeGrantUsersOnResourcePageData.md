# DescribeGrantUsersOnResourcePageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int64** | 总条数 | [default to 0]
**PageSize** | **int64** | 请求条数 | [default to 0]
**PageNumber** | **int64** | 请求页数 | [default to 0]
**UserDataList** | [**[]UserData**](UserData.md) |  | 

## Methods

### NewDescribeGrantUsersOnResourcePageData

`func NewDescribeGrantUsersOnResourcePageData(totalCount int64, pageSize int64, pageNumber int64, userDataList []UserData, ) *DescribeGrantUsersOnResourcePageData`

NewDescribeGrantUsersOnResourcePageData instantiates a new DescribeGrantUsersOnResourcePageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGrantUsersOnResourcePageDataWithDefaults

`func NewDescribeGrantUsersOnResourcePageDataWithDefaults() *DescribeGrantUsersOnResourcePageData`

NewDescribeGrantUsersOnResourcePageDataWithDefaults instantiates a new DescribeGrantUsersOnResourcePageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *DescribeGrantUsersOnResourcePageData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeGrantUsersOnResourcePageData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeGrantUsersOnResourcePageData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetPageSize

`func (o *DescribeGrantUsersOnResourcePageData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGrantUsersOnResourcePageData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGrantUsersOnResourcePageData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetPageNumber

`func (o *DescribeGrantUsersOnResourcePageData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeGrantUsersOnResourcePageData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeGrantUsersOnResourcePageData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetUserDataList

`func (o *DescribeGrantUsersOnResourcePageData) GetUserDataList() []UserData`

GetUserDataList returns the UserDataList field if non-nil, zero value otherwise.

### GetUserDataListOk

`func (o *DescribeGrantUsersOnResourcePageData) GetUserDataListOk() (*[]UserData, bool)`

GetUserDataListOk returns a tuple with the UserDataList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataList

`func (o *DescribeGrantUsersOnResourcePageData) SetUserDataList(v []UserData)`

SetUserDataList sets UserDataList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


