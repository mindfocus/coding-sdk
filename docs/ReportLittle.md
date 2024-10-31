# ReportLittle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** | 创建时间 | [optional] [default to ""]
**Id** | Pointer to **NullableInt32** | ID 主键 | [optional] 
**Name** | Pointer to **NullableString** | 报告名称 | [optional] [default to ""]
**StatisticsEndTime** | Pointer to **NullableString** | 数据统计结束时间 | [optional] [default to ""]
**StatisticsStartTime** | Pointer to **NullableString** | 数据统计开始时间 | [optional] [default to ""]
**Status** | Pointer to **NullableString** | 报告状态：CREATING 创建中，AVAILABLE 可用，UNAVAILABLE 不可用 | [optional] [default to ""]
**Summary** | Pointer to **NullableString** | 报告总结 | [optional] [default to ""]

## Methods

### NewReportLittle

`func NewReportLittle() *ReportLittle`

NewReportLittle instantiates a new ReportLittle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportLittleWithDefaults

`func NewReportLittleWithDefaults() *ReportLittle`

NewReportLittleWithDefaults instantiates a new ReportLittle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ReportLittle) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportLittle) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportLittle) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReportLittle) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ReportLittle) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ReportLittle) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetId

`func (o *ReportLittle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportLittle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportLittle) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportLittle) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ReportLittle) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReportLittle) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ReportLittle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportLittle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportLittle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportLittle) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ReportLittle) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ReportLittle) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatisticsEndTime

`func (o *ReportLittle) GetStatisticsEndTime() string`

GetStatisticsEndTime returns the StatisticsEndTime field if non-nil, zero value otherwise.

### GetStatisticsEndTimeOk

`func (o *ReportLittle) GetStatisticsEndTimeOk() (*string, bool)`

GetStatisticsEndTimeOk returns a tuple with the StatisticsEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsEndTime

`func (o *ReportLittle) SetStatisticsEndTime(v string)`

SetStatisticsEndTime sets StatisticsEndTime field to given value.

### HasStatisticsEndTime

`func (o *ReportLittle) HasStatisticsEndTime() bool`

HasStatisticsEndTime returns a boolean if a field has been set.

### SetStatisticsEndTimeNil

`func (o *ReportLittle) SetStatisticsEndTimeNil(b bool)`

 SetStatisticsEndTimeNil sets the value for StatisticsEndTime to be an explicit nil

### UnsetStatisticsEndTime
`func (o *ReportLittle) UnsetStatisticsEndTime()`

UnsetStatisticsEndTime ensures that no value is present for StatisticsEndTime, not even an explicit nil
### GetStatisticsStartTime

`func (o *ReportLittle) GetStatisticsStartTime() string`

GetStatisticsStartTime returns the StatisticsStartTime field if non-nil, zero value otherwise.

### GetStatisticsStartTimeOk

`func (o *ReportLittle) GetStatisticsStartTimeOk() (*string, bool)`

GetStatisticsStartTimeOk returns a tuple with the StatisticsStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsStartTime

`func (o *ReportLittle) SetStatisticsStartTime(v string)`

SetStatisticsStartTime sets StatisticsStartTime field to given value.

### HasStatisticsStartTime

`func (o *ReportLittle) HasStatisticsStartTime() bool`

HasStatisticsStartTime returns a boolean if a field has been set.

### SetStatisticsStartTimeNil

`func (o *ReportLittle) SetStatisticsStartTimeNil(b bool)`

 SetStatisticsStartTimeNil sets the value for StatisticsStartTime to be an explicit nil

### UnsetStatisticsStartTime
`func (o *ReportLittle) UnsetStatisticsStartTime()`

UnsetStatisticsStartTime ensures that no value is present for StatisticsStartTime, not even an explicit nil
### GetStatus

`func (o *ReportLittle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReportLittle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReportLittle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReportLittle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ReportLittle) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ReportLittle) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSummary

`func (o *ReportLittle) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ReportLittle) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ReportLittle) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ReportLittle) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *ReportLittle) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *ReportLittle) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


