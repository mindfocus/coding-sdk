# DescribeCdTaskResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to **string** | 应用名 | [optional] [default to ""]
**TaskExecutionJsonContent** | Pointer to **string** | CD 任务执行记录 | [optional] [default to ""]
**TaskExecutionStatus** | Pointer to **string** | 任务执行状态 | [optional] [default to ""]

## Methods

### NewDescribeCdTaskResponseData

`func NewDescribeCdTaskResponseData() *DescribeCdTaskResponseData`

NewDescribeCdTaskResponseData instantiates a new DescribeCdTaskResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdTaskResponseDataWithDefaults

`func NewDescribeCdTaskResponseDataWithDefaults() *DescribeCdTaskResponseData`

NewDescribeCdTaskResponseDataWithDefaults instantiates a new DescribeCdTaskResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *DescribeCdTaskResponseData) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DescribeCdTaskResponseData) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DescribeCdTaskResponseData) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DescribeCdTaskResponseData) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetTaskExecutionJsonContent

`func (o *DescribeCdTaskResponseData) GetTaskExecutionJsonContent() string`

GetTaskExecutionJsonContent returns the TaskExecutionJsonContent field if non-nil, zero value otherwise.

### GetTaskExecutionJsonContentOk

`func (o *DescribeCdTaskResponseData) GetTaskExecutionJsonContentOk() (*string, bool)`

GetTaskExecutionJsonContentOk returns a tuple with the TaskExecutionJsonContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecutionJsonContent

`func (o *DescribeCdTaskResponseData) SetTaskExecutionJsonContent(v string)`

SetTaskExecutionJsonContent sets TaskExecutionJsonContent field to given value.

### HasTaskExecutionJsonContent

`func (o *DescribeCdTaskResponseData) HasTaskExecutionJsonContent() bool`

HasTaskExecutionJsonContent returns a boolean if a field has been set.

### GetTaskExecutionStatus

`func (o *DescribeCdTaskResponseData) GetTaskExecutionStatus() string`

GetTaskExecutionStatus returns the TaskExecutionStatus field if non-nil, zero value otherwise.

### GetTaskExecutionStatusOk

`func (o *DescribeCdTaskResponseData) GetTaskExecutionStatusOk() (*string, bool)`

GetTaskExecutionStatusOk returns a tuple with the TaskExecutionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecutionStatus

`func (o *DescribeCdTaskResponseData) SetTaskExecutionStatus(v string)`

SetTaskExecutionStatus sets TaskExecutionStatus field to given value.

### HasTaskExecutionStatus

`func (o *DescribeCdTaskResponseData) HasTaskExecutionStatus() bool`

HasTaskExecutionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


