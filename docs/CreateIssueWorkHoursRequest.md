# CreateIssueWorkHoursRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项 Code | 
**ProjectName** | **string** | 项目名称 | 
**RemainingHour** | **float32** | 剩余工时 | 
**SpendHour** | **float32** | 使用工时 | 
**StartAt** | **int64** | 开始时间戳 | 
**WorkingDesc** | Pointer to **string** | 工作描述 | [optional] 

## Methods

### NewCreateIssueWorkHoursRequest

`func NewCreateIssueWorkHoursRequest(issueCode int64, projectName string, remainingHour float32, spendHour float32, startAt int64, ) *CreateIssueWorkHoursRequest`

NewCreateIssueWorkHoursRequest instantiates a new CreateIssueWorkHoursRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueWorkHoursRequestWithDefaults

`func NewCreateIssueWorkHoursRequestWithDefaults() *CreateIssueWorkHoursRequest`

NewCreateIssueWorkHoursRequestWithDefaults instantiates a new CreateIssueWorkHoursRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *CreateIssueWorkHoursRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *CreateIssueWorkHoursRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *CreateIssueWorkHoursRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *CreateIssueWorkHoursRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateIssueWorkHoursRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateIssueWorkHoursRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRemainingHour

`func (o *CreateIssueWorkHoursRequest) GetRemainingHour() float32`

GetRemainingHour returns the RemainingHour field if non-nil, zero value otherwise.

### GetRemainingHourOk

`func (o *CreateIssueWorkHoursRequest) GetRemainingHourOk() (*float32, bool)`

GetRemainingHourOk returns a tuple with the RemainingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingHour

`func (o *CreateIssueWorkHoursRequest) SetRemainingHour(v float32)`

SetRemainingHour sets RemainingHour field to given value.


### GetSpendHour

`func (o *CreateIssueWorkHoursRequest) GetSpendHour() float32`

GetSpendHour returns the SpendHour field if non-nil, zero value otherwise.

### GetSpendHourOk

`func (o *CreateIssueWorkHoursRequest) GetSpendHourOk() (*float32, bool)`

GetSpendHourOk returns a tuple with the SpendHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendHour

`func (o *CreateIssueWorkHoursRequest) SetSpendHour(v float32)`

SetSpendHour sets SpendHour field to given value.


### GetStartAt

`func (o *CreateIssueWorkHoursRequest) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *CreateIssueWorkHoursRequest) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *CreateIssueWorkHoursRequest) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.


### GetWorkingDesc

`func (o *CreateIssueWorkHoursRequest) GetWorkingDesc() string`

GetWorkingDesc returns the WorkingDesc field if non-nil, zero value otherwise.

### GetWorkingDescOk

`func (o *CreateIssueWorkHoursRequest) GetWorkingDescOk() (*string, bool)`

GetWorkingDescOk returns a tuple with the WorkingDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDesc

`func (o *CreateIssueWorkHoursRequest) SetWorkingDesc(v string)`

SetWorkingDesc sets WorkingDesc field to given value.

### HasWorkingDesc

`func (o *CreateIssueWorkHoursRequest) HasWorkingDesc() bool`

HasWorkingDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


