# ReorderCdPipelinesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** | CD 应用名 | 
**PipelineIdsIndices** | [**[]PipelineIdIndex**](PipelineIdIndex.md) | 部署流程排序列表 | 

## Methods

### NewReorderCdPipelinesRequest

`func NewReorderCdPipelinesRequest(application string, pipelineIdsIndices []PipelineIdIndex, ) *ReorderCdPipelinesRequest`

NewReorderCdPipelinesRequest instantiates a new ReorderCdPipelinesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReorderCdPipelinesRequestWithDefaults

`func NewReorderCdPipelinesRequestWithDefaults() *ReorderCdPipelinesRequest`

NewReorderCdPipelinesRequestWithDefaults instantiates a new ReorderCdPipelinesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ReorderCdPipelinesRequest) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ReorderCdPipelinesRequest) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ReorderCdPipelinesRequest) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetPipelineIdsIndices

`func (o *ReorderCdPipelinesRequest) GetPipelineIdsIndices() []PipelineIdIndex`

GetPipelineIdsIndices returns the PipelineIdsIndices field if non-nil, zero value otherwise.

### GetPipelineIdsIndicesOk

`func (o *ReorderCdPipelinesRequest) GetPipelineIdsIndicesOk() (*[]PipelineIdIndex, bool)`

GetPipelineIdsIndicesOk returns a tuple with the PipelineIdsIndices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineIdsIndices

`func (o *ReorderCdPipelinesRequest) SetPipelineIdsIndices(v []PipelineIdIndex)`

SetPipelineIdsIndices sets PipelineIdsIndices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


