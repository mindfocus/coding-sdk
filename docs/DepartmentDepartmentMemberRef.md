# DepartmentDepartmentMemberRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | **bool** | 是否部门负责人 | [default to false]
**DepartmentId** | **int64** | 部门ID | [default to 0]
**DepartmentMemberId** | **int64** | 部门成员ID | [default to 0]
**DepartmentName** | **string** | 部门名 | [default to ""]
**DescribeId** | **string** | 部门描述ID | [default to ""]
**Pointer** | **bool** | 是否是当前查询部门的直接成员 | [default to false]
**RefId** | **int64** | 部门用户的refId | [default to 0]

## Methods

### NewDepartmentDepartmentMemberRef

`func NewDepartmentDepartmentMemberRef(assignee bool, departmentId int64, departmentMemberId int64, departmentName string, describeId string, pointer bool, refId int64, ) *DepartmentDepartmentMemberRef`

NewDepartmentDepartmentMemberRef instantiates a new DepartmentDepartmentMemberRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentDepartmentMemberRefWithDefaults

`func NewDepartmentDepartmentMemberRefWithDefaults() *DepartmentDepartmentMemberRef`

NewDepartmentDepartmentMemberRefWithDefaults instantiates a new DepartmentDepartmentMemberRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *DepartmentDepartmentMemberRef) GetAssignee() bool`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *DepartmentDepartmentMemberRef) GetAssigneeOk() (*bool, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *DepartmentDepartmentMemberRef) SetAssignee(v bool)`

SetAssignee sets Assignee field to given value.


### GetDepartmentId

`func (o *DepartmentDepartmentMemberRef) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *DepartmentDepartmentMemberRef) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *DepartmentDepartmentMemberRef) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.


### GetDepartmentMemberId

`func (o *DepartmentDepartmentMemberRef) GetDepartmentMemberId() int64`

GetDepartmentMemberId returns the DepartmentMemberId field if non-nil, zero value otherwise.

### GetDepartmentMemberIdOk

`func (o *DepartmentDepartmentMemberRef) GetDepartmentMemberIdOk() (*int64, bool)`

GetDepartmentMemberIdOk returns a tuple with the DepartmentMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentMemberId

`func (o *DepartmentDepartmentMemberRef) SetDepartmentMemberId(v int64)`

SetDepartmentMemberId sets DepartmentMemberId field to given value.


### GetDepartmentName

`func (o *DepartmentDepartmentMemberRef) GetDepartmentName() string`

GetDepartmentName returns the DepartmentName field if non-nil, zero value otherwise.

### GetDepartmentNameOk

`func (o *DepartmentDepartmentMemberRef) GetDepartmentNameOk() (*string, bool)`

GetDepartmentNameOk returns a tuple with the DepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentName

`func (o *DepartmentDepartmentMemberRef) SetDepartmentName(v string)`

SetDepartmentName sets DepartmentName field to given value.


### GetDescribeId

`func (o *DepartmentDepartmentMemberRef) GetDescribeId() string`

GetDescribeId returns the DescribeId field if non-nil, zero value otherwise.

### GetDescribeIdOk

`func (o *DepartmentDepartmentMemberRef) GetDescribeIdOk() (*string, bool)`

GetDescribeIdOk returns a tuple with the DescribeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribeId

`func (o *DepartmentDepartmentMemberRef) SetDescribeId(v string)`

SetDescribeId sets DescribeId field to given value.


### GetPointer

`func (o *DepartmentDepartmentMemberRef) GetPointer() bool`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *DepartmentDepartmentMemberRef) GetPointerOk() (*bool, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *DepartmentDepartmentMemberRef) SetPointer(v bool)`

SetPointer sets Pointer field to given value.


### GetRefId

`func (o *DepartmentDepartmentMemberRef) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *DepartmentDepartmentMemberRef) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *DepartmentDepartmentMemberRef) SetRefId(v int64)`

SetRefId sets RefId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


