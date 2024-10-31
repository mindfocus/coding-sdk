# UpdateUserGroupByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | 用户组描述 | 
**GroupId** | **int64** | 用户组ID | 
**Name** | **string** | 用户组名称 | 

## Methods

### NewUpdateUserGroupByIdRequest

`func NewUpdateUserGroupByIdRequest(description string, groupId int64, name string, ) *UpdateUserGroupByIdRequest`

NewUpdateUserGroupByIdRequest instantiates a new UpdateUserGroupByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserGroupByIdRequestWithDefaults

`func NewUpdateUserGroupByIdRequestWithDefaults() *UpdateUserGroupByIdRequest`

NewUpdateUserGroupByIdRequestWithDefaults instantiates a new UpdateUserGroupByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateUserGroupByIdRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateUserGroupByIdRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateUserGroupByIdRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroupId

`func (o *UpdateUserGroupByIdRequest) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateUserGroupByIdRequest) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateUserGroupByIdRequest) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetName

`func (o *UpdateUserGroupByIdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserGroupByIdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserGroupByIdRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


