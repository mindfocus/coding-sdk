# TriggerCdPipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** | CD 应用名 | 
**PipelineNameOrId** | **string** | 部署流程名称或 ID | 
**TriggerJsonContent** | **string** | 触发参数 JSON 配置 | 

## Methods

### NewTriggerCdPipelineRequest

`func NewTriggerCdPipelineRequest(application string, pipelineNameOrId string, triggerJsonContent string, ) *TriggerCdPipelineRequest`

NewTriggerCdPipelineRequest instantiates a new TriggerCdPipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCdPipelineRequestWithDefaults

`func NewTriggerCdPipelineRequestWithDefaults() *TriggerCdPipelineRequest`

NewTriggerCdPipelineRequestWithDefaults instantiates a new TriggerCdPipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *TriggerCdPipelineRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *TriggerCdPipelineRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *TriggerCdPipelineRequest) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetPipelineNameOrId

`func (o *TriggerCdPipelineRequest) GetPipelineNameOrId() string`

GetPipelineNameOrId returns the PipelineNameOrId field if non-nil, zero value otherwise.

### GetPipelineNameOrIdOk

`func (o *TriggerCdPipelineRequest) GetPipelineNameOrIdOk() (*string, bool)`

GetPipelineNameOrIdOk returns a tuple with the PipelineNameOrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineNameOrId

`func (o *TriggerCdPipelineRequest) SetPipelineNameOrId(v string)`

SetPipelineNameOrId sets PipelineNameOrId field to given value.


### GetTriggerJsonContent

`func (o *TriggerCdPipelineRequest) GetTriggerJsonContent() string`

GetTriggerJsonContent returns the TriggerJsonContent field if non-nil, zero value otherwise.

### GetTriggerJsonContentOk

`func (o *TriggerCdPipelineRequest) GetTriggerJsonContentOk() (*string, bool)`

GetTriggerJsonContentOk returns a tuple with the TriggerJsonContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerJsonContent

`func (o *TriggerCdPipelineRequest) SetTriggerJsonContent(v string)`

SetTriggerJsonContent sets TriggerJsonContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


