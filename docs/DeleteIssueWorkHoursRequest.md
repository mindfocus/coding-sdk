# DeleteIssueWorkHoursRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项 Code | 
**ProjectName** | **string** | 项目名称 | 
**RollbackRemainingHours** | **bool** | 是否将该工时日志的使用工时归还到剩余工时，默认：false | 
**WorkHourLogId** | **int32** | 工时日志 Id | 

## Methods

### NewDeleteIssueWorkHoursRequest

`func NewDeleteIssueWorkHoursRequest(issueCode int64, projectName string, rollbackRemainingHours bool, workHourLogId int32, ) *DeleteIssueWorkHoursRequest`

NewDeleteIssueWorkHoursRequest instantiates a new DeleteIssueWorkHoursRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIssueWorkHoursRequestWithDefaults

`func NewDeleteIssueWorkHoursRequestWithDefaults() *DeleteIssueWorkHoursRequest`

NewDeleteIssueWorkHoursRequestWithDefaults instantiates a new DeleteIssueWorkHoursRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *DeleteIssueWorkHoursRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DeleteIssueWorkHoursRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DeleteIssueWorkHoursRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *DeleteIssueWorkHoursRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteIssueWorkHoursRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteIssueWorkHoursRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRollbackRemainingHours

`func (o *DeleteIssueWorkHoursRequest) GetRollbackRemainingHours() bool`

GetRollbackRemainingHours returns the RollbackRemainingHours field if non-nil, zero value otherwise.

### GetRollbackRemainingHoursOk

`func (o *DeleteIssueWorkHoursRequest) GetRollbackRemainingHoursOk() (*bool, bool)`

GetRollbackRemainingHoursOk returns a tuple with the RollbackRemainingHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollbackRemainingHours

`func (o *DeleteIssueWorkHoursRequest) SetRollbackRemainingHours(v bool)`

SetRollbackRemainingHours sets RollbackRemainingHours field to given value.


### GetWorkHourLogId

`func (o *DeleteIssueWorkHoursRequest) GetWorkHourLogId() int32`

GetWorkHourLogId returns the WorkHourLogId field if non-nil, zero value otherwise.

### GetWorkHourLogIdOk

`func (o *DeleteIssueWorkHoursRequest) GetWorkHourLogIdOk() (*int32, bool)`

GetWorkHourLogIdOk returns a tuple with the WorkHourLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkHourLogId

`func (o *DeleteIssueWorkHoursRequest) SetWorkHourLogId(v int32)`

SetWorkHourLogId sets WorkHourLogId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


