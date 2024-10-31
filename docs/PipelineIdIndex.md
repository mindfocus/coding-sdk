# PipelineIdIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int64** | 排序索引 | 
**PipelineId** | **string** | 部署流程 ID | [default to ""]

## Methods

### NewPipelineIdIndex

`func NewPipelineIdIndex(index int64, pipelineId string, ) *PipelineIdIndex`

NewPipelineIdIndex instantiates a new PipelineIdIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineIdIndexWithDefaults

`func NewPipelineIdIndexWithDefaults() *PipelineIdIndex`

NewPipelineIdIndexWithDefaults instantiates a new PipelineIdIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *PipelineIdIndex) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *PipelineIdIndex) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *PipelineIdIndex) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetPipelineId

`func (o *PipelineIdIndex) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *PipelineIdIndex) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *PipelineIdIndex) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


