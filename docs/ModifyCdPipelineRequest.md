# ModifyCdPipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineConfigId** | **string** | 部署流程 ID | 
**PipelineJsonContent** | **string** | 部署流程 JSON 配置 | 

## Methods

### NewModifyCdPipelineRequest

`func NewModifyCdPipelineRequest(pipelineConfigId string, pipelineJsonContent string, ) *ModifyCdPipelineRequest`

NewModifyCdPipelineRequest instantiates a new ModifyCdPipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCdPipelineRequestWithDefaults

`func NewModifyCdPipelineRequestWithDefaults() *ModifyCdPipelineRequest`

NewModifyCdPipelineRequestWithDefaults instantiates a new ModifyCdPipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineConfigId

`func (o *ModifyCdPipelineRequest) GetPipelineConfigId() string`

GetPipelineConfigId returns the PipelineConfigId field if non-nil, zero value otherwise.

### GetPipelineConfigIdOk

`func (o *ModifyCdPipelineRequest) GetPipelineConfigIdOk() (*string, bool)`

GetPipelineConfigIdOk returns a tuple with the PipelineConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineConfigId

`func (o *ModifyCdPipelineRequest) SetPipelineConfigId(v string)`

SetPipelineConfigId sets PipelineConfigId field to given value.


### GetPipelineJsonContent

`func (o *ModifyCdPipelineRequest) GetPipelineJsonContent() string`

GetPipelineJsonContent returns the PipelineJsonContent field if non-nil, zero value otherwise.

### GetPipelineJsonContentOk

`func (o *ModifyCdPipelineRequest) GetPipelineJsonContentOk() (*string, bool)`

GetPipelineJsonContentOk returns a tuple with the PipelineJsonContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineJsonContent

`func (o *ModifyCdPipelineRequest) SetPipelineJsonContent(v string)`

SetPipelineJsonContent sets PipelineJsonContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


