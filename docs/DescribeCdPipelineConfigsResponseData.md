# DescribeCdPipelineConfigsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineConfigs** | Pointer to [**[]PipelineConfig**](PipelineConfig.md) | 部署流程 JSON 配置列表 | [optional] 

## Methods

### NewDescribeCdPipelineConfigsResponseData

`func NewDescribeCdPipelineConfigsResponseData() *DescribeCdPipelineConfigsResponseData`

NewDescribeCdPipelineConfigsResponseData instantiates a new DescribeCdPipelineConfigsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdPipelineConfigsResponseDataWithDefaults

`func NewDescribeCdPipelineConfigsResponseDataWithDefaults() *DescribeCdPipelineConfigsResponseData`

NewDescribeCdPipelineConfigsResponseDataWithDefaults instantiates a new DescribeCdPipelineConfigsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineConfigs

`func (o *DescribeCdPipelineConfigsResponseData) GetPipelineConfigs() []PipelineConfig`

GetPipelineConfigs returns the PipelineConfigs field if non-nil, zero value otherwise.

### GetPipelineConfigsOk

`func (o *DescribeCdPipelineConfigsResponseData) GetPipelineConfigsOk() (*[]PipelineConfig, bool)`

GetPipelineConfigsOk returns a tuple with the PipelineConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineConfigs

`func (o *DescribeCdPipelineConfigsResponseData) SetPipelineConfigs(v []PipelineConfig)`

SetPipelineConfigs sets PipelineConfigs field to given value.

### HasPipelineConfigs

`func (o *DescribeCdPipelineConfigsResponseData) HasPipelineConfigs() bool`

HasPipelineConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


