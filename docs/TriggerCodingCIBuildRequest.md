# TriggerCodingCIBuildRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **int32** | 构建计划 Id | 
**ParamList** | Pointer to [**[]CodingCIJobEnv**](CodingCIJobEnv.md) | 构建附加的环境变量 | [optional] 
**Reentrant** | Pointer to **string** | 可重入字符串 | [optional] 
**Revision** | Pointer to **string** | 分支名或 CommitId，当为构建计划的 DepotType 为 NONE 是可不传 | [optional] 

## Methods

### NewTriggerCodingCIBuildRequest

`func NewTriggerCodingCIBuildRequest(jobId int32, ) *TriggerCodingCIBuildRequest`

NewTriggerCodingCIBuildRequest instantiates a new TriggerCodingCIBuildRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCodingCIBuildRequestWithDefaults

`func NewTriggerCodingCIBuildRequestWithDefaults() *TriggerCodingCIBuildRequest`

NewTriggerCodingCIBuildRequestWithDefaults instantiates a new TriggerCodingCIBuildRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *TriggerCodingCIBuildRequest) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *TriggerCodingCIBuildRequest) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *TriggerCodingCIBuildRequest) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetParamList

`func (o *TriggerCodingCIBuildRequest) GetParamList() []CodingCIJobEnv`

GetParamList returns the ParamList field if non-nil, zero value otherwise.

### GetParamListOk

`func (o *TriggerCodingCIBuildRequest) GetParamListOk() (*[]CodingCIJobEnv, bool)`

GetParamListOk returns a tuple with the ParamList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamList

`func (o *TriggerCodingCIBuildRequest) SetParamList(v []CodingCIJobEnv)`

SetParamList sets ParamList field to given value.

### HasParamList

`func (o *TriggerCodingCIBuildRequest) HasParamList() bool`

HasParamList returns a boolean if a field has been set.

### GetReentrant

`func (o *TriggerCodingCIBuildRequest) GetReentrant() string`

GetReentrant returns the Reentrant field if non-nil, zero value otherwise.

### GetReentrantOk

`func (o *TriggerCodingCIBuildRequest) GetReentrantOk() (*string, bool)`

GetReentrantOk returns a tuple with the Reentrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReentrant

`func (o *TriggerCodingCIBuildRequest) SetReentrant(v string)`

SetReentrant sets Reentrant field to given value.

### HasReentrant

`func (o *TriggerCodingCIBuildRequest) HasReentrant() bool`

HasReentrant returns a boolean if a field has been set.

### GetRevision

`func (o *TriggerCodingCIBuildRequest) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *TriggerCodingCIBuildRequest) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *TriggerCodingCIBuildRequest) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *TriggerCodingCIBuildRequest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


