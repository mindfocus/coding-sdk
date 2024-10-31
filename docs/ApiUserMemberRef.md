# ApiUserMemberRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignee** | Pointer to **bool** | 是否管理员 | [optional] [default to false]
**DepartmentId** | Pointer to **int64** | 部门ID | [optional] [default to 0]
**DepartmentMemberId** | Pointer to **int64** | 部门成员ID | [optional] [default to 0]
**DepartmentName** | Pointer to **string** | 部门名 | [optional] [default to ""]
**DescribeId** | Pointer to **string** | 描述ID | [optional] [default to ""]
**Pointer** | Pointer to **bool** | 是否是当前查询部门的直接成员 | [optional] [default to false]

## Methods

### NewApiUserMemberRef

`func NewApiUserMemberRef() *ApiUserMemberRef`

NewApiUserMemberRef instantiates a new ApiUserMemberRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserMemberRefWithDefaults

`func NewApiUserMemberRefWithDefaults() *ApiUserMemberRef`

NewApiUserMemberRefWithDefaults instantiates a new ApiUserMemberRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignee

`func (o *ApiUserMemberRef) GetAssignee() bool`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *ApiUserMemberRef) GetAssigneeOk() (*bool, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *ApiUserMemberRef) SetAssignee(v bool)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *ApiUserMemberRef) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### GetDepartmentId

`func (o *ApiUserMemberRef) GetDepartmentId() int64`

GetDepartmentId returns the DepartmentId field if non-nil, zero value otherwise.

### GetDepartmentIdOk

`func (o *ApiUserMemberRef) GetDepartmentIdOk() (*int64, bool)`

GetDepartmentIdOk returns a tuple with the DepartmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentId

`func (o *ApiUserMemberRef) SetDepartmentId(v int64)`

SetDepartmentId sets DepartmentId field to given value.

### HasDepartmentId

`func (o *ApiUserMemberRef) HasDepartmentId() bool`

HasDepartmentId returns a boolean if a field has been set.

### GetDepartmentMemberId

`func (o *ApiUserMemberRef) GetDepartmentMemberId() int64`

GetDepartmentMemberId returns the DepartmentMemberId field if non-nil, zero value otherwise.

### GetDepartmentMemberIdOk

`func (o *ApiUserMemberRef) GetDepartmentMemberIdOk() (*int64, bool)`

GetDepartmentMemberIdOk returns a tuple with the DepartmentMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentMemberId

`func (o *ApiUserMemberRef) SetDepartmentMemberId(v int64)`

SetDepartmentMemberId sets DepartmentMemberId field to given value.

### HasDepartmentMemberId

`func (o *ApiUserMemberRef) HasDepartmentMemberId() bool`

HasDepartmentMemberId returns a boolean if a field has been set.

### GetDepartmentName

`func (o *ApiUserMemberRef) GetDepartmentName() string`

GetDepartmentName returns the DepartmentName field if non-nil, zero value otherwise.

### GetDepartmentNameOk

`func (o *ApiUserMemberRef) GetDepartmentNameOk() (*string, bool)`

GetDepartmentNameOk returns a tuple with the DepartmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentName

`func (o *ApiUserMemberRef) SetDepartmentName(v string)`

SetDepartmentName sets DepartmentName field to given value.

### HasDepartmentName

`func (o *ApiUserMemberRef) HasDepartmentName() bool`

HasDepartmentName returns a boolean if a field has been set.

### GetDescribeId

`func (o *ApiUserMemberRef) GetDescribeId() string`

GetDescribeId returns the DescribeId field if non-nil, zero value otherwise.

### GetDescribeIdOk

`func (o *ApiUserMemberRef) GetDescribeIdOk() (*string, bool)`

GetDescribeIdOk returns a tuple with the DescribeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribeId

`func (o *ApiUserMemberRef) SetDescribeId(v string)`

SetDescribeId sets DescribeId field to given value.

### HasDescribeId

`func (o *ApiUserMemberRef) HasDescribeId() bool`

HasDescribeId returns a boolean if a field has been set.

### GetPointer

`func (o *ApiUserMemberRef) GetPointer() bool`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ApiUserMemberRef) GetPointerOk() (*bool, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ApiUserMemberRef) SetPointer(v bool)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ApiUserMemberRef) HasPointer() bool`

HasPointer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


