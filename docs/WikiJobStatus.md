# WikiJobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | 任务Id | [optional] [default to ""]
**Status** | Pointer to **string** | 任务状态 | [optional] [default to ""]

## Methods

### NewWikiJobStatus

`func NewWikiJobStatus() *WikiJobStatus`

NewWikiJobStatus instantiates a new WikiJobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiJobStatusWithDefaults

`func NewWikiJobStatusWithDefaults() *WikiJobStatus`

NewWikiJobStatusWithDefaults instantiates a new WikiJobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *WikiJobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *WikiJobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *WikiJobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *WikiJobStatus) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStatus

`func (o *WikiJobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WikiJobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WikiJobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WikiJobStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


