# DeleteServiceHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **[]string** | Service Hook 编号 | 
**ProjectId** | **int64** | 项目编号或者研发空间编号 | 
**TargetType** | Pointer to **string** | 目标数据类型：PROJECT,SPACE_NODE,PROGRAM,默认PROJECT | [optional] 

## Methods

### NewDeleteServiceHookRequest

`func NewDeleteServiceHookRequest(id []string, projectId int64, ) *DeleteServiceHookRequest`

NewDeleteServiceHookRequest instantiates a new DeleteServiceHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteServiceHookRequestWithDefaults

`func NewDeleteServiceHookRequestWithDefaults() *DeleteServiceHookRequest`

NewDeleteServiceHookRequestWithDefaults instantiates a new DeleteServiceHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeleteServiceHookRequest) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeleteServiceHookRequest) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeleteServiceHookRequest) SetId(v []string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *DeleteServiceHookRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DeleteServiceHookRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DeleteServiceHookRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetTargetType

`func (o *DeleteServiceHookRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *DeleteServiceHookRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *DeleteServiceHookRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *DeleteServiceHookRequest) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


