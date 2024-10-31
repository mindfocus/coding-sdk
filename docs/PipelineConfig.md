# PipelineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | 部署流程 ID | [optional] [default to ""]
**JsonContent** | Pointer to **string** | 部署流程 JSON 配置 | [optional] [default to ""]
**Name** | Pointer to **string** | 部署流程名称 | [optional] [default to ""]

## Methods

### NewPipelineConfig

`func NewPipelineConfig() *PipelineConfig`

NewPipelineConfig instantiates a new PipelineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineConfigWithDefaults

`func NewPipelineConfigWithDefaults() *PipelineConfig`

NewPipelineConfigWithDefaults instantiates a new PipelineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PipelineConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJsonContent

`func (o *PipelineConfig) GetJsonContent() string`

GetJsonContent returns the JsonContent field if non-nil, zero value otherwise.

### GetJsonContentOk

`func (o *PipelineConfig) GetJsonContentOk() (*string, bool)`

GetJsonContentOk returns a tuple with the JsonContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonContent

`func (o *PipelineConfig) SetJsonContent(v string)`

SetJsonContent sets JsonContent field to given value.

### HasJsonContent

`func (o *PipelineConfig) HasJsonContent() bool`

HasJsonContent returns a boolean if a field has been set.

### GetName

`func (o *PipelineConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineConfig) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


