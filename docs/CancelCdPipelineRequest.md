# CancelCdPipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineExecutionId** | **string** | 部署流程执行记录 ID | 
**Reason** | Pointer to **string** | 取消原因 | [optional] 

## Methods

### NewCancelCdPipelineRequest

`func NewCancelCdPipelineRequest(pipelineExecutionId string, ) *CancelCdPipelineRequest`

NewCancelCdPipelineRequest instantiates a new CancelCdPipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelCdPipelineRequestWithDefaults

`func NewCancelCdPipelineRequestWithDefaults() *CancelCdPipelineRequest`

NewCancelCdPipelineRequestWithDefaults instantiates a new CancelCdPipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineExecutionId

`func (o *CancelCdPipelineRequest) GetPipelineExecutionId() string`

GetPipelineExecutionId returns the PipelineExecutionId field if non-nil, zero value otherwise.

### GetPipelineExecutionIdOk

`func (o *CancelCdPipelineRequest) GetPipelineExecutionIdOk() (*string, bool)`

GetPipelineExecutionIdOk returns a tuple with the PipelineExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineExecutionId

`func (o *CancelCdPipelineRequest) SetPipelineExecutionId(v string)`

SetPipelineExecutionId sets PipelineExecutionId field to given value.


### GetReason

`func (o *CancelCdPipelineRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CancelCdPipelineRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CancelCdPipelineRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CancelCdPipelineRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


