# ModifyProjectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 项目描述 | [optional] 
**DisplayName** | **string** | 项目名称 | 
**EndDate** | Pointer to **string** | 结束时间 | [optional] 
**Name** | **string** | 项目标识 | 
**ProjectId** | **int64** | 项目编号 | 
**StartDate** | Pointer to **string** | 开始时间 | [optional] 

## Methods

### NewModifyProjectRequest

`func NewModifyProjectRequest(displayName string, name string, projectId int64, ) *ModifyProjectRequest`

NewModifyProjectRequest instantiates a new ModifyProjectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyProjectRequestWithDefaults

`func NewModifyProjectRequestWithDefaults() *ModifyProjectRequest`

NewModifyProjectRequestWithDefaults instantiates a new ModifyProjectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ModifyProjectRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyProjectRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyProjectRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyProjectRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModifyProjectRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModifyProjectRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModifyProjectRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEndDate

`func (o *ModifyProjectRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ModifyProjectRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ModifyProjectRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ModifyProjectRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetName

`func (o *ModifyProjectRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyProjectRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyProjectRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *ModifyProjectRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModifyProjectRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModifyProjectRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetStartDate

`func (o *ModifyProjectRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ModifyProjectRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ModifyProjectRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ModifyProjectRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


