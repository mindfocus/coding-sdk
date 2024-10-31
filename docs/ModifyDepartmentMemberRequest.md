# ModifyDepartmentMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentIds** | **[]int64** | 部门ID信息 | 
**RefIds** | **[]int64** | 用户refIds | 

## Methods

### NewModifyDepartmentMemberRequest

`func NewModifyDepartmentMemberRequest(departmentIds []int64, refIds []int64, ) *ModifyDepartmentMemberRequest`

NewModifyDepartmentMemberRequest instantiates a new ModifyDepartmentMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepartmentMemberRequestWithDefaults

`func NewModifyDepartmentMemberRequestWithDefaults() *ModifyDepartmentMemberRequest`

NewModifyDepartmentMemberRequestWithDefaults instantiates a new ModifyDepartmentMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentIds

`func (o *ModifyDepartmentMemberRequest) GetDepartmentIds() []int64`

GetDepartmentIds returns the DepartmentIds field if non-nil, zero value otherwise.

### GetDepartmentIdsOk

`func (o *ModifyDepartmentMemberRequest) GetDepartmentIdsOk() (*[]int64, bool)`

GetDepartmentIdsOk returns a tuple with the DepartmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentIds

`func (o *ModifyDepartmentMemberRequest) SetDepartmentIds(v []int64)`

SetDepartmentIds sets DepartmentIds field to given value.


### GetRefIds

`func (o *ModifyDepartmentMemberRequest) GetRefIds() []int64`

GetRefIds returns the RefIds field if non-nil, zero value otherwise.

### GetRefIdsOk

`func (o *ModifyDepartmentMemberRequest) GetRefIdsOk() (*[]int64, bool)`

GetRefIdsOk returns a tuple with the RefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefIds

`func (o *ModifyDepartmentMemberRequest) SetRefIds(v []int64)`

SetRefIds sets RefIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


