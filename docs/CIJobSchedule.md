# CIJobSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | 要触发的分支 | [default to ""]
**EndTime** | **string** | 结束时间。如果是单次触发，结束时间为空 | [default to ""]
**Interval** | **string** | 间隔 | [default to ""]
**RefChangeTrigger** | **bool** | 代码无变化时是否触发 | [default to false]
**Repeat** | **bool** | 是否周期触发（周期触发/单次触发） | [default to false]
**StartTime** | **string** | 开始时间。如果是周期触发，精确到小时（ 8 ）如果是单次触发，精确到分钟数（ 8:20 ） | [default to ""]
**Weekend** | **string** | 星期几 | [default to ""]

## Methods

### NewCIJobSchedule

`func NewCIJobSchedule(branch string, endTime string, interval string, refChangeTrigger bool, repeat bool, startTime string, weekend string, ) *CIJobSchedule`

NewCIJobSchedule instantiates a new CIJobSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCIJobScheduleWithDefaults

`func NewCIJobScheduleWithDefaults() *CIJobSchedule`

NewCIJobScheduleWithDefaults instantiates a new CIJobSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *CIJobSchedule) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CIJobSchedule) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CIJobSchedule) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetEndTime

`func (o *CIJobSchedule) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CIJobSchedule) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CIJobSchedule) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetInterval

`func (o *CIJobSchedule) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CIJobSchedule) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CIJobSchedule) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetRefChangeTrigger

`func (o *CIJobSchedule) GetRefChangeTrigger() bool`

GetRefChangeTrigger returns the RefChangeTrigger field if non-nil, zero value otherwise.

### GetRefChangeTriggerOk

`func (o *CIJobSchedule) GetRefChangeTriggerOk() (*bool, bool)`

GetRefChangeTriggerOk returns a tuple with the RefChangeTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefChangeTrigger

`func (o *CIJobSchedule) SetRefChangeTrigger(v bool)`

SetRefChangeTrigger sets RefChangeTrigger field to given value.


### GetRepeat

`func (o *CIJobSchedule) GetRepeat() bool`

GetRepeat returns the Repeat field if non-nil, zero value otherwise.

### GetRepeatOk

`func (o *CIJobSchedule) GetRepeatOk() (*bool, bool)`

GetRepeatOk returns a tuple with the Repeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeat

`func (o *CIJobSchedule) SetRepeat(v bool)`

SetRepeat sets Repeat field to given value.


### GetStartTime

`func (o *CIJobSchedule) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CIJobSchedule) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CIJobSchedule) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetWeekend

`func (o *CIJobSchedule) GetWeekend() string`

GetWeekend returns the Weekend field if non-nil, zero value otherwise.

### GetWeekendOk

`func (o *CIJobSchedule) GetWeekendOk() (*string, bool)`

GetWeekendOk returns a tuple with the Weekend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekend

`func (o *CIJobSchedule) SetWeekend(v string)`

SetWeekend sets Weekend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


