# CreateCdPipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineJsonContent** | **string** | 部署流程 JSON 配置（注意：stages 的 refId 和 requisiteStageRefIds 字段必传） | 

## Methods

### NewCreateCdPipelineRequest

`func NewCreateCdPipelineRequest(pipelineJsonContent string, ) *CreateCdPipelineRequest`

NewCreateCdPipelineRequest instantiates a new CreateCdPipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCdPipelineRequestWithDefaults

`func NewCreateCdPipelineRequestWithDefaults() *CreateCdPipelineRequest`

NewCreateCdPipelineRequestWithDefaults instantiates a new CreateCdPipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineJsonContent

`func (o *CreateCdPipelineRequest) GetPipelineJsonContent() string`

GetPipelineJsonContent returns the PipelineJsonContent field if non-nil, zero value otherwise.

### GetPipelineJsonContentOk

`func (o *CreateCdPipelineRequest) GetPipelineJsonContentOk() (*string, bool)`

GetPipelineJsonContentOk returns a tuple with the PipelineJsonContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineJsonContent

`func (o *CreateCdPipelineRequest) SetPipelineJsonContent(v string)`

SetPipelineJsonContent sets PipelineJsonContent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


