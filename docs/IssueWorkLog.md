# IssueWorkLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间戳 | [optional] 
**Id** | Pointer to **int64** | 工时日志 Id | [optional] 
**IssueCode** | Pointer to **int64** | 事项编号 | [optional] 
**IssueId** | Pointer to **int64** | 事项 Id | [optional] 
**ProjectName** | Pointer to **string** | 项目名称 | [optional] [default to ""]
**RecordHours** | Pointer to **float32** | 使用工时 | [optional] 
**RemainingHours** | Pointer to **float32** | 剩余工时 | [optional] 
**StartAt** | Pointer to **int64** | 开始时间 | [optional] 
**UpdatedAt** | Pointer to **NullableInt64** | 更新时间 | [optional] 
**UserId** | Pointer to **int64** | 用户 Id | [optional] 
**WorkingDesc** | Pointer to **string** | 工作描述 | [optional] [default to ""]

## Methods

### NewIssueWorkLog

`func NewIssueWorkLog() *IssueWorkLog`

NewIssueWorkLog instantiates a new IssueWorkLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueWorkLogWithDefaults

`func NewIssueWorkLogWithDefaults() *IssueWorkLog`

NewIssueWorkLogWithDefaults instantiates a new IssueWorkLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IssueWorkLog) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueWorkLog) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueWorkLog) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueWorkLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *IssueWorkLog) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueWorkLog) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueWorkLog) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueWorkLog) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssueCode

`func (o *IssueWorkLog) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *IssueWorkLog) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *IssueWorkLog) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.

### HasIssueCode

`func (o *IssueWorkLog) HasIssueCode() bool`

HasIssueCode returns a boolean if a field has been set.

### GetIssueId

`func (o *IssueWorkLog) GetIssueId() int64`

GetIssueId returns the IssueId field if non-nil, zero value otherwise.

### GetIssueIdOk

`func (o *IssueWorkLog) GetIssueIdOk() (*int64, bool)`

GetIssueIdOk returns a tuple with the IssueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueId

`func (o *IssueWorkLog) SetIssueId(v int64)`

SetIssueId sets IssueId field to given value.

### HasIssueId

`func (o *IssueWorkLog) HasIssueId() bool`

HasIssueId returns a boolean if a field has been set.

### GetProjectName

`func (o *IssueWorkLog) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *IssueWorkLog) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *IssueWorkLog) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *IssueWorkLog) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetRecordHours

`func (o *IssueWorkLog) GetRecordHours() float32`

GetRecordHours returns the RecordHours field if non-nil, zero value otherwise.

### GetRecordHoursOk

`func (o *IssueWorkLog) GetRecordHoursOk() (*float32, bool)`

GetRecordHoursOk returns a tuple with the RecordHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordHours

`func (o *IssueWorkLog) SetRecordHours(v float32)`

SetRecordHours sets RecordHours field to given value.

### HasRecordHours

`func (o *IssueWorkLog) HasRecordHours() bool`

HasRecordHours returns a boolean if a field has been set.

### GetRemainingHours

`func (o *IssueWorkLog) GetRemainingHours() float32`

GetRemainingHours returns the RemainingHours field if non-nil, zero value otherwise.

### GetRemainingHoursOk

`func (o *IssueWorkLog) GetRemainingHoursOk() (*float32, bool)`

GetRemainingHoursOk returns a tuple with the RemainingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingHours

`func (o *IssueWorkLog) SetRemainingHours(v float32)`

SetRemainingHours sets RemainingHours field to given value.

### HasRemainingHours

`func (o *IssueWorkLog) HasRemainingHours() bool`

HasRemainingHours returns a boolean if a field has been set.

### GetStartAt

`func (o *IssueWorkLog) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *IssueWorkLog) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *IssueWorkLog) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *IssueWorkLog) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IssueWorkLog) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IssueWorkLog) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IssueWorkLog) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IssueWorkLog) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IssueWorkLog) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IssueWorkLog) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *IssueWorkLog) GetUserId() int64`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *IssueWorkLog) GetUserIdOk() (*int64, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *IssueWorkLog) SetUserId(v int64)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *IssueWorkLog) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWorkingDesc

`func (o *IssueWorkLog) GetWorkingDesc() string`

GetWorkingDesc returns the WorkingDesc field if non-nil, zero value otherwise.

### GetWorkingDescOk

`func (o *IssueWorkLog) GetWorkingDescOk() (*string, bool)`

GetWorkingDescOk returns a tuple with the WorkingDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDesc

`func (o *IssueWorkLog) SetWorkingDesc(v string)`

SetWorkingDesc sets WorkingDesc field to given value.

### HasWorkingDesc

`func (o *IssueWorkLog) HasWorkingDesc() bool`

HasWorkingDesc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


