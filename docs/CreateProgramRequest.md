# CreateProgramRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | 描述信息 | [default to ""]
**DisplayName** | **string** | 展示名 | [default to ""]
**EndDate** | **string** | 截止时间 | [default to ""]
**Name** | **string** | 项目集名 | [default to ""]
**StartDate** | **string** | 开始时间 | [default to ""]
**WorkflowProgramId** | Pointer to **string** | 已有工作流项目集 Id | [optional] [default to ""]

## Methods

### NewCreateProgramRequest

`func NewCreateProgramRequest(description string, displayName string, endDate string, name string, startDate string, ) *CreateProgramRequest`

NewCreateProgramRequest instantiates a new CreateProgramRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProgramRequestWithDefaults

`func NewCreateProgramRequestWithDefaults() *CreateProgramRequest`

NewCreateProgramRequestWithDefaults instantiates a new CreateProgramRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateProgramRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProgramRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProgramRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *CreateProgramRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateProgramRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateProgramRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEndDate

`func (o *CreateProgramRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateProgramRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateProgramRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetName

`func (o *CreateProgramRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProgramRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProgramRequest) SetName(v string)`

SetName sets Name field to given value.


### GetStartDate

`func (o *CreateProgramRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateProgramRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateProgramRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetWorkflowProgramId

`func (o *CreateProgramRequest) GetWorkflowProgramId() string`

GetWorkflowProgramId returns the WorkflowProgramId field if non-nil, zero value otherwise.

### GetWorkflowProgramIdOk

`func (o *CreateProgramRequest) GetWorkflowProgramIdOk() (*string, bool)`

GetWorkflowProgramIdOk returns a tuple with the WorkflowProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowProgramId

`func (o *CreateProgramRequest) SetWorkflowProgramId(v string)`

SetWorkflowProgramId sets WorkflowProgramId field to given value.

### HasWorkflowProgramId

`func (o *CreateProgramRequest) HasWorkflowProgramId() bool`

HasWorkflowProgramId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


