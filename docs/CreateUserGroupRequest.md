# CreateUserGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | 用户组描述 | 
**Name** | **string** | 用户组名称 | 

## Methods

### NewCreateUserGroupRequest

`func NewCreateUserGroupRequest(description string, name string, ) *CreateUserGroupRequest`

NewCreateUserGroupRequest instantiates a new CreateUserGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserGroupRequestWithDefaults

`func NewCreateUserGroupRequestWithDefaults() *CreateUserGroupRequest`

NewCreateUserGroupRequestWithDefaults instantiates a new CreateUserGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateUserGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUserGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUserGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *CreateUserGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserGroupRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


