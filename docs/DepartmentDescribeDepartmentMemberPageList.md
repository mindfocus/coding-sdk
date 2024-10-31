# DepartmentDescribeDepartmentMemberPageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentMembers** | [**[]DepartmentDepartmentMembersData**](DepartmentDepartmentMembersData.md) | 部门成员 | 
**PageNumber** | **int64** | 页数 | [default to 0]
**PageSize** | **int64** | 每页数量 | [default to 0]
**TotalCount** | **int64** | 总数 | [default to 0]

## Methods

### NewDepartmentDescribeDepartmentMemberPageList

`func NewDepartmentDescribeDepartmentMemberPageList(departmentMembers []DepartmentDepartmentMembersData, pageNumber int64, pageSize int64, totalCount int64, ) *DepartmentDescribeDepartmentMemberPageList`

NewDepartmentDescribeDepartmentMemberPageList instantiates a new DepartmentDescribeDepartmentMemberPageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentDescribeDepartmentMemberPageListWithDefaults

`func NewDepartmentDescribeDepartmentMemberPageListWithDefaults() *DepartmentDescribeDepartmentMemberPageList`

NewDepartmentDescribeDepartmentMemberPageListWithDefaults instantiates a new DepartmentDescribeDepartmentMemberPageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentMembers

`func (o *DepartmentDescribeDepartmentMemberPageList) GetDepartmentMembers() []DepartmentDepartmentMembersData`

GetDepartmentMembers returns the DepartmentMembers field if non-nil, zero value otherwise.

### GetDepartmentMembersOk

`func (o *DepartmentDescribeDepartmentMemberPageList) GetDepartmentMembersOk() (*[]DepartmentDepartmentMembersData, bool)`

GetDepartmentMembersOk returns a tuple with the DepartmentMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentMembers

`func (o *DepartmentDescribeDepartmentMemberPageList) SetDepartmentMembers(v []DepartmentDepartmentMembersData)`

SetDepartmentMembers sets DepartmentMembers field to given value.


### GetPageNumber

`func (o *DepartmentDescribeDepartmentMemberPageList) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DepartmentDescribeDepartmentMemberPageList) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DepartmentDescribeDepartmentMemberPageList) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DepartmentDescribeDepartmentMemberPageList) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DepartmentDescribeDepartmentMemberPageList) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DepartmentDescribeDepartmentMemberPageList) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetTotalCount

`func (o *DepartmentDescribeDepartmentMemberPageList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DepartmentDescribeDepartmentMemberPageList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DepartmentDescribeDepartmentMemberPageList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


