# ProjectMemberMemberRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | **NullableBool** | 是否管理员 | [default to false]
**DepartmentId** | **NullableInt64** | 部门ID | 
**DepartmentMemberId** | **NullableInt64** | 部门成员ID | 
**DepartmentName** | **NullableString** | 部门名 | [default to ""]
**DescribeId** | **NullableString** | 描述ID | [default to ""]
**Pointer** | **NullableBool** | 是否是当前查询部门的直接成员 | [default to false]

## Methods

### NewProjectMemberMemberRef

`func NewProjectMemberMemberRef(assignee NullableBool, departmentId NullableInt64, departmentMemberId NullableInt64, departmentName NullableString, describeId NullableString, pointer NullableBool, ) *ProjectMemberMemberRef`

NewProjectMemberMemberRef instantiates a new ProjectMemberMemberRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectMemberMemberRefWithDefaults

`func NewProjectMemberMemberRefWithDefaults() *ProjectMemberMemberRef`

NewProjectMemberMemberRefWithDefaults instantiates a new ProjectMemberMemberRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ProjectMemberMemberRef) GetAssignee() bool`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ProjectMemberMemberRef) GetAssigneeOk() (*bool, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ProjectMemberMemberRef) SetAssignee(v bool)`

SetAssignee sets Assignee field to given value.


### SetAssigneeNil

`func (o *ProjectMemberMemberRef) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *ProjectMemberMemberRef) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetDepartmentId

`func (o *ProjectMemberMemberRef) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *ProjectMemberMemberRef) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *ProjectMemberMemberRef) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.


### SetDepartmentIdNil

`func (o *ProjectMemberMemberRef) SetDepartmentIdNil(b bool)`

 SetDepartmentIdNil sets the value for DepartmentId to be an explicit nil

### UnsetDepartmentId
`func (o *ProjectMemberMemberRef) UnsetDepartmentId()`

UnsetDepartmentId ensures that no value is present for DepartmentId, not even an explicit nil
### GetDepartmentMemberId

`func (o *ProjectMemberMemberRef) GetDepartmentMemberId() int64`

GetDepartmentMemberId returns the DepartmentMemberId field if non-nil, zero value otherwise.

### GetDepartmentMemberIdOk

`func (o *ProjectMemberMemberRef) GetDepartmentMemberIdOk() (*int64, bool)`

GetDepartmentMemberIdOk returns a tuple with the DepartmentMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentMemberId

`func (o *ProjectMemberMemberRef) SetDepartmentMemberId(v int64)`

SetDepartmentMemberId sets DepartmentMemberId field to given value.


### SetDepartmentMemberIdNil

`func (o *ProjectMemberMemberRef) SetDepartmentMemberIdNil(b bool)`

 SetDepartmentMemberIdNil sets the value for DepartmentMemberId to be an explicit nil

### UnsetDepartmentMemberId
`func (o *ProjectMemberMemberRef) UnsetDepartmentMemberId()`

UnsetDepartmentMemberId ensures that no value is present for DepartmentMemberId, not even an explicit nil
### GetDepartmentName

`func (o *ProjectMemberMemberRef) GetDepartmentName() string`

GetDepartmentName returns the DepartmentName field if non-nil, zero value otherwise.

### GetDepartmentNameOk

`func (o *ProjectMemberMemberRef) GetDepartmentNameOk() (*string, bool)`

GetDepartmentNameOk returns a tuple with the DepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentName

`func (o *ProjectMemberMemberRef) SetDepartmentName(v string)`

SetDepartmentName sets DepartmentName field to given value.


### SetDepartmentNameNil

`func (o *ProjectMemberMemberRef) SetDepartmentNameNil(b bool)`

 SetDepartmentNameNil sets the value for DepartmentName to be an explicit nil

### UnsetDepartmentName
`func (o *ProjectMemberMemberRef) UnsetDepartmentName()`

UnsetDepartmentName ensures that no value is present for DepartmentName, not even an explicit nil
### GetDescribeId

`func (o *ProjectMemberMemberRef) GetDescribeId() string`

GetDescribeId returns the DescribeId field if non-nil, zero value otherwise.

### GetDescribeIdOk

`func (o *ProjectMemberMemberRef) GetDescribeIdOk() (*string, bool)`

GetDescribeIdOk returns a tuple with the DescribeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribeId

`func (o *ProjectMemberMemberRef) SetDescribeId(v string)`

SetDescribeId sets DescribeId field to given value.


### SetDescribeIdNil

`func (o *ProjectMemberMemberRef) SetDescribeIdNil(b bool)`

 SetDescribeIdNil sets the value for DescribeId to be an explicit nil

### UnsetDescribeId
`func (o *ProjectMemberMemberRef) UnsetDescribeId()`

UnsetDescribeId ensures that no value is present for DescribeId, not even an explicit nil
### GetPointer

`func (o *ProjectMemberMemberRef) GetPointer() bool`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ProjectMemberMemberRef) GetPointerOk() (*bool, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ProjectMemberMemberRef) SetPointer(v bool)`

SetPointer sets Pointer field to given value.


### SetPointerNil

`func (o *ProjectMemberMemberRef) SetPointerNil(b bool)`

 SetPointerNil sets the value for Pointer to be an explicit nil

### UnsetPointer
`func (o *ProjectMemberMemberRef) UnsetPointer()`

UnsetPointer ensures that no value is present for Pointer, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


