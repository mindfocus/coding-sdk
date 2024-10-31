# DescribeDepartmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentId** | **int64** | 部门ID | [default to 0]
**GetTree** | Pointer to **bool** | 是否获取部门树 | [optional] [default to false]

## Methods

### NewDescribeDepartmentRequest

`func NewDescribeDepartmentRequest(departmentId int64, ) *DescribeDepartmentRequest`

NewDescribeDepartmentRequest instantiates a new DescribeDepartmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepartmentRequestWithDefaults

`func NewDescribeDepartmentRequestWithDefaults() *DescribeDepartmentRequest`

NewDescribeDepartmentRequestWithDefaults instantiates a new DescribeDepartmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentId

`func (o *DescribeDepartmentRequest) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *DescribeDepartmentRequest) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *DescribeDepartmentRequest) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.


### GetGetTree

`func (o *DescribeDepartmentRequest) GetGetTree() bool`

GetGetTree returns the GetTree field if non-nil, zero value otherwise.

### GetGetTreeOk

`func (o *DescribeDepartmentRequest) GetGetTreeOk() (*bool, bool)`

GetGetTreeOk returns a tuple with the GetTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetTree

`func (o *DescribeDepartmentRequest) SetGetTree(v bool)`

SetGetTree sets GetTree field to given value.

### HasGetTree

`func (o *DescribeDepartmentRequest) HasGetTree() bool`

HasGetTree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


