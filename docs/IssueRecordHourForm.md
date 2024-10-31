# IssueRecordHourForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemainingHour** | **float32** | 剩余工时 | 
**SpendHour** | **float32** | 使用工时 | 
**StartAt** | **int64** | 开始时间 | 
**WorkingDesc** | **string** | 工作描述 | [default to ""]

## Methods

### NewIssueRecordHourForm

`func NewIssueRecordHourForm(remainingHour float32, spendHour float32, startAt int64, workingDesc string, ) *IssueRecordHourForm`

NewIssueRecordHourForm instantiates a new IssueRecordHourForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueRecordHourFormWithDefaults

`func NewIssueRecordHourFormWithDefaults() *IssueRecordHourForm`

NewIssueRecordHourFormWithDefaults instantiates a new IssueRecordHourForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemainingHour

`func (o *IssueRecordHourForm) GetRemainingHour() float32`

GetRemainingHour returns the RemainingHour field if non-nil, zero value otherwise.

### GetRemainingHourOk

`func (o *IssueRecordHourForm) GetRemainingHourOk() (*float32, bool)`

GetRemainingHourOk returns a tuple with the RemainingHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingHour

`func (o *IssueRecordHourForm) SetRemainingHour(v float32)`

SetRemainingHour sets RemainingHour field to given value.


### GetSpendHour

`func (o *IssueRecordHourForm) GetSpendHour() float32`

GetSpendHour returns the SpendHour field if non-nil, zero value otherwise.

### GetSpendHourOk

`func (o *IssueRecordHourForm) GetSpendHourOk() (*float32, bool)`

GetSpendHourOk returns a tuple with the SpendHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendHour

`func (o *IssueRecordHourForm) SetSpendHour(v float32)`

SetSpendHour sets SpendHour field to given value.


### GetStartAt

`func (o *IssueRecordHourForm) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *IssueRecordHourForm) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *IssueRecordHourForm) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.


### GetWorkingDesc

`func (o *IssueRecordHourForm) GetWorkingDesc() string`

GetWorkingDesc returns the WorkingDesc field if non-nil, zero value otherwise.

### GetWorkingDescOk

`func (o *IssueRecordHourForm) GetWorkingDescOk() (*string, bool)`

GetWorkingDescOk returns a tuple with the WorkingDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDesc

`func (o *IssueRecordHourForm) SetWorkingDesc(v string)`

SetWorkingDesc sets WorkingDesc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


