# DescribeTeamMemberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | 用户Id | 
**ShowDepartment** | Pointer to **bool** | 是否展示部门 | [optional] [default to false]

## Methods

### NewDescribeTeamMemberRequest

`func NewDescribeTeamMemberRequest(userId int32, ) *DescribeTeamMemberRequest`

NewDescribeTeamMemberRequest instantiates a new DescribeTeamMemberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamMemberRequestWithDefaults

`func NewDescribeTeamMemberRequestWithDefaults() *DescribeTeamMemberRequest`

NewDescribeTeamMemberRequestWithDefaults instantiates a new DescribeTeamMemberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *DescribeTeamMemberRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DescribeTeamMemberRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DescribeTeamMemberRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetShowDepartment

`func (o *DescribeTeamMemberRequest) GetShowDepartment() bool`

GetShowDepartment returns the ShowDepartment field if non-nil, zero value otherwise.

### GetShowDepartmentOk

`func (o *DescribeTeamMemberRequest) GetShowDepartmentOk() (*bool, bool)`

GetShowDepartmentOk returns a tuple with the ShowDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDepartment

`func (o *DescribeTeamMemberRequest) SetShowDepartment(v bool)`

SetShowDepartment sets ShowDepartment field to given value.

### HasShowDepartment

`func (o *DescribeTeamMemberRequest) HasShowDepartment() bool`

HasShowDepartment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


