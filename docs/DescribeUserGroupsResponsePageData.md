# DescribeUserGroupsResponsePageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 每页条数 | [optional] 
**TotalCount** | Pointer to **int64** | 总条数 | [optional] 
**UserGroupList** | Pointer to [**[]UserGroup**](UserGroup.md) | 用户组列表 | [optional] 

## Methods

### NewDescribeUserGroupsResponsePageData

`func NewDescribeUserGroupsResponsePageData() *DescribeUserGroupsResponsePageData`

NewDescribeUserGroupsResponsePageData instantiates a new DescribeUserGroupsResponsePageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUserGroupsResponsePageDataWithDefaults

`func NewDescribeUserGroupsResponsePageDataWithDefaults() *DescribeUserGroupsResponsePageData`

NewDescribeUserGroupsResponsePageDataWithDefaults instantiates a new DescribeUserGroupsResponsePageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeUserGroupsResponsePageData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeUserGroupsResponsePageData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeUserGroupsResponsePageData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeUserGroupsResponsePageData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeUserGroupsResponsePageData) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeUserGroupsResponsePageData) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeUserGroupsResponsePageData) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeUserGroupsResponsePageData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *DescribeUserGroupsResponsePageData) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeUserGroupsResponsePageData) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeUserGroupsResponsePageData) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DescribeUserGroupsResponsePageData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetUserGroupList

`func (o *DescribeUserGroupsResponsePageData) GetUserGroupList() []UserGroup`

GetUserGroupList returns the UserGroupList field if non-nil, zero value otherwise.

### GetUserGroupListOk

`func (o *DescribeUserGroupsResponsePageData) GetUserGroupListOk() (*[]UserGroup, bool)`

GetUserGroupListOk returns a tuple with the UserGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupList

`func (o *DescribeUserGroupsResponsePageData) SetUserGroupList(v []UserGroup)`

SetUserGroupList sets UserGroupList field to given value.

### HasUserGroupList

`func (o *DescribeUserGroupsResponsePageData) HasUserGroupList() bool`

HasUserGroupList returns a boolean if a field has been set.

### SetUserGroupListNil

`func (o *DescribeUserGroupsResponsePageData) SetUserGroupListNil(b bool)`

 SetUserGroupListNil sets the value for UserGroupList to be an explicit nil

### UnsetUserGroupList
`func (o *DescribeUserGroupsResponsePageData) UnsetUserGroupList()`

UnsetUserGroupList ensures that no value is present for UserGroupList, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


