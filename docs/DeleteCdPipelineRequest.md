# DeleteCdPipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** | CD 应用名 | 
**PipelineName** | **string** | 部署流程名称 | 

## Methods

### NewDeleteCdPipelineRequest

`func NewDeleteCdPipelineRequest(application string, pipelineName string, ) *DeleteCdPipelineRequest`

NewDeleteCdPipelineRequest instantiates a new DeleteCdPipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCdPipelineRequestWithDefaults

`func NewDeleteCdPipelineRequestWithDefaults() *DeleteCdPipelineRequest`

NewDeleteCdPipelineRequestWithDefaults instantiates a new DeleteCdPipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *DeleteCdPipelineRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DeleteCdPipelineRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DeleteCdPipelineRequest) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetPipelineName

`func (o *DeleteCdPipelineRequest) GetPipelineName() string`

GetPipelineName returns the PipelineName field if non-nil, zero value otherwise.

### GetPipelineNameOk

`func (o *DeleteCdPipelineRequest) GetPipelineNameOk() (*string, bool)`

GetPipelineNameOk returns a tuple with the PipelineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineName

`func (o *DeleteCdPipelineRequest) SetPipelineName(v string)`

SetPipelineName sets PipelineName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


