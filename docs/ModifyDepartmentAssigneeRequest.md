# ModifyDepartmentAssigneeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartmentId** | **int64** | 部门ID | [default to 0]
**RefIds** | **[]int64** | 用户refIds 必填 | 

## Methods

### NewModifyDepartmentAssigneeRequest

`func NewModifyDepartmentAssigneeRequest(departmentId int64, refIds []int64, ) *ModifyDepartmentAssigneeRequest`

NewModifyDepartmentAssigneeRequest instantiates a new ModifyDepartmentAssigneeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepartmentAssigneeRequestWithDefaults

`func NewModifyDepartmentAssigneeRequestWithDefaults() *ModifyDepartmentAssigneeRequest`

NewModifyDepartmentAssigneeRequestWithDefaults instantiates a new ModifyDepartmentAssigneeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartmentId

`func (o *ModifyDepartmentAssigneeRequest) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *ModifyDepartmentAssigneeRequest) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *ModifyDepartmentAssigneeRequest) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.


### GetRefIds

`func (o *ModifyDepartmentAssigneeRequest) GetRefIds() []int64`

GetRefIds returns the RefIds field if non-nil, zero value otherwise.

### GetRefIdsOk

`func (o *ModifyDepartmentAssigneeRequest) GetRefIdsOk() (*[]int64, bool)`

GetRefIdsOk returns a tuple with the RefIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefIds

`func (o *ModifyDepartmentAssigneeRequest) SetRefIds(v []int64)`

SetRefIds sets RefIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


