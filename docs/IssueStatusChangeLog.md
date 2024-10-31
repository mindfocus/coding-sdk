# IssueStatusChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableInt64** | 创建时间 | [optional] 
**IssueCode** | Pointer to **NullableInt64** | 事项code | [optional] 
**IssueStatus** | Pointer to [**IssueStatus**](IssueStatus.md) |  | [optional] 
**StatusId** | Pointer to **NullableInt64** | 状态ID | [optional] 
**StatusName** | Pointer to **NullableString** | 事项章台名称 | [optional] [default to ""]

## Methods

### NewIssueStatusChangeLog

`func NewIssueStatusChangeLog() *IssueStatusChangeLog`

NewIssueStatusChangeLog instantiates a new IssueStatusChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueStatusChangeLogWithDefaults

`func NewIssueStatusChangeLogWithDefaults() *IssueStatusChangeLog`

NewIssueStatusChangeLogWithDefaults instantiates a new IssueStatusChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *IssueStatusChangeLog) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IssueStatusChangeLog) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IssueStatusChangeLog) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IssueStatusChangeLog) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IssueStatusChangeLog) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IssueStatusChangeLog) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIssueCode

`func (o *IssueStatusChangeLog) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *IssueStatusChangeLog) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *IssueStatusChangeLog) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.

### HasIssueCode

`func (o *IssueStatusChangeLog) HasIssueCode() bool`

HasIssueCode returns a boolean if a field has been set.

### SetIssueCodeNil

`func (o *IssueStatusChangeLog) SetIssueCodeNil(b bool)`

 SetIssueCodeNil sets the value for IssueCode to be an explicit nil

### UnsetIssueCode
`func (o *IssueStatusChangeLog) UnsetIssueCode()`

UnsetIssueCode ensures that no value is present for IssueCode, not even an explicit nil
### GetIssueStatus

`func (o *IssueStatusChangeLog) GetIssueStatus() IssueStatus`

GetIssueStatus returns the IssueStatus field if non-nil, zero value otherwise.

### GetIssueStatusOk

`func (o *IssueStatusChangeLog) GetIssueStatusOk() (*IssueStatus, bool)`

GetIssueStatusOk returns a tuple with the IssueStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueStatus

`func (o *IssueStatusChangeLog) SetIssueStatus(v IssueStatus)`

SetIssueStatus sets IssueStatus field to given value.

### HasIssueStatus

`func (o *IssueStatusChangeLog) HasIssueStatus() bool`

HasIssueStatus returns a boolean if a field has been set.

### GetStatusId

`func (o *IssueStatusChangeLog) GetStatusId() int64`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *IssueStatusChangeLog) GetStatusIdOk() (*int64, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *IssueStatusChangeLog) SetStatusId(v int64)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *IssueStatusChangeLog) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### SetStatusIdNil

`func (o *IssueStatusChangeLog) SetStatusIdNil(b bool)`

 SetStatusIdNil sets the value for StatusId to be an explicit nil

### UnsetStatusId
`func (o *IssueStatusChangeLog) UnsetStatusId()`

UnsetStatusId ensures that no value is present for StatusId, not even an explicit nil
### GetStatusName

`func (o *IssueStatusChangeLog) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *IssueStatusChangeLog) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *IssueStatusChangeLog) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *IssueStatusChangeLog) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### SetStatusNameNil

`func (o *IssueStatusChangeLog) SetStatusNameNil(b bool)`

 SetStatusNameNil sets the value for StatusName to be an explicit nil

### UnsetStatusName
`func (o *IssueStatusChangeLog) UnsetStatusName()`

UnsetStatusName ensures that no value is present for StatusName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


