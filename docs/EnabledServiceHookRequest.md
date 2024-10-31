# EnabledServiceHookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | 是否开启 | 
**Id** | **[]string** | Service Hook 编号 | 
**ProjectId** | **int64** | 项目或者研发空间编号 | 
**TargetType** | **string** | 目标数据类型：PROJECT,SPACE_NODE,PROGRAM,默认PROJECT | 

## Methods

### NewEnabledServiceHookRequest

`func NewEnabledServiceHookRequest(enabled bool, id []string, projectId int64, targetType string, ) *EnabledServiceHookRequest`

NewEnabledServiceHookRequest instantiates a new EnabledServiceHookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnabledServiceHookRequestWithDefaults

`func NewEnabledServiceHookRequestWithDefaults() *EnabledServiceHookRequest`

NewEnabledServiceHookRequestWithDefaults instantiates a new EnabledServiceHookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EnabledServiceHookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EnabledServiceHookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EnabledServiceHookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *EnabledServiceHookRequest) GetId() []string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnabledServiceHookRequest) GetIdOk() (*[]string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnabledServiceHookRequest) SetId(v []string)`

SetId sets Id field to given value.


### GetProjectId

`func (o *EnabledServiceHookRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *EnabledServiceHookRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *EnabledServiceHookRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetTargetType

`func (o *EnabledServiceHookRequest) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *EnabledServiceHookRequest) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *EnabledServiceHookRequest) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


