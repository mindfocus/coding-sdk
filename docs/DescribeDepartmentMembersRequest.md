# DescribeDepartmentMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentId** | **int64** | 部门名 | [default to 0]
**KeyWords** | Pointer to **string** | 关键词 | [optional] [default to ""]
**PageNumber** | **int64** | 页数 | [default to 0]
**PageSize** | **int64** | 每页数量 | [default to 0]
**Pointer** | Pointer to **bool** | 是否仅查询当前部门的直接成员 | [optional] [default to false]

## Methods

### NewDescribeDepartmentMembersRequest

`func NewDescribeDepartmentMembersRequest(departmentId int64, pageNumber int64, pageSize int64, ) *DescribeDepartmentMembersRequest`

NewDescribeDepartmentMembersRequest instantiates a new DescribeDepartmentMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepartmentMembersRequestWithDefaults

`func NewDescribeDepartmentMembersRequestWithDefaults() *DescribeDepartmentMembersRequest`

NewDescribeDepartmentMembersRequestWithDefaults instantiates a new DescribeDepartmentMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentId

`func (o *DescribeDepartmentMembersRequest) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *DescribeDepartmentMembersRequest) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *DescribeDepartmentMembersRequest) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.


### GetKeyWords

`func (o *DescribeDepartmentMembersRequest) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *DescribeDepartmentMembersRequest) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *DescribeDepartmentMembersRequest) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *DescribeDepartmentMembersRequest) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeDepartmentMembersRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeDepartmentMembersRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeDepartmentMembersRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeDepartmentMembersRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeDepartmentMembersRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeDepartmentMembersRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetPointer

`func (o *DescribeDepartmentMembersRequest) GetPointer() bool`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *DescribeDepartmentMembersRequest) GetPointerOk() (*bool, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *DescribeDepartmentMembersRequest) SetPointer(v bool)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *DescribeDepartmentMembersRequest) HasPointer() bool`

HasPointer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


