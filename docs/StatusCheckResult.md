# StatusCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | Pointer to **NullableString** | 提交sha | [optional] [default to ""]
**Context** | Pointer to **NullableString** | 流水线文本内容 | [optional] [default to ""]
**CreateDate** | Pointer to **NullableInt64** | 记录创建时间 | [optional] 
**DepotId** | Pointer to **NullableInt64** | 仓库路径 | [optional] 
**Description** | Pointer to **NullableString** | 流水线状态描述 | [optional] [default to ""]
**State** | Pointer to **NullableString** | 构建状态 | [optional] [default to ""]
**TargetURL** | Pointer to **NullableString** | 流水线链接地址 | [optional] [default to ""]

## Methods

### NewStatusCheckResult

`func NewStatusCheckResult() *StatusCheckResult`

NewStatusCheckResult instantiates a new StatusCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCheckResultWithDefaults

`func NewStatusCheckResultWithDefaults() *StatusCheckResult`

NewStatusCheckResultWithDefaults instantiates a new StatusCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *StatusCheckResult) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *StatusCheckResult) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *StatusCheckResult) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *StatusCheckResult) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### SetCommitShaNil

`func (o *StatusCheckResult) SetCommitShaNil(b bool)`

 SetCommitShaNil sets the value for CommitSha to be an explicit nil

### UnsetCommitSha
`func (o *StatusCheckResult) UnsetCommitSha()`

UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
### GetContext

`func (o *StatusCheckResult) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StatusCheckResult) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StatusCheckResult) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *StatusCheckResult) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *StatusCheckResult) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *StatusCheckResult) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetCreateDate

`func (o *StatusCheckResult) GetCreateDate() int64`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *StatusCheckResult) GetCreateDateOk() (*int64, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *StatusCheckResult) SetCreateDate(v int64)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *StatusCheckResult) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### SetCreateDateNil

`func (o *StatusCheckResult) SetCreateDateNil(b bool)`

 SetCreateDateNil sets the value for CreateDate to be an explicit nil

### UnsetCreateDate
`func (o *StatusCheckResult) UnsetCreateDate()`

UnsetCreateDate ensures that no value is present for CreateDate, not even an explicit nil
### GetDepotId

`func (o *StatusCheckResult) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *StatusCheckResult) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *StatusCheckResult) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *StatusCheckResult) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### SetDepotIdNil

`func (o *StatusCheckResult) SetDepotIdNil(b bool)`

 SetDepotIdNil sets the value for DepotId to be an explicit nil

### UnsetDepotId
`func (o *StatusCheckResult) UnsetDepotId()`

UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
### GetDescription

`func (o *StatusCheckResult) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StatusCheckResult) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StatusCheckResult) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StatusCheckResult) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StatusCheckResult) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StatusCheckResult) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetState

`func (o *StatusCheckResult) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StatusCheckResult) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StatusCheckResult) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StatusCheckResult) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *StatusCheckResult) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *StatusCheckResult) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTargetURL

`func (o *StatusCheckResult) GetTargetURL() string`

GetTargetURL returns the TargetURL field if non-nil, zero value otherwise.

### GetTargetURLOk

`func (o *StatusCheckResult) GetTargetURLOk() (*string, bool)`

GetTargetURLOk returns a tuple with the TargetURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetURL

`func (o *StatusCheckResult) SetTargetURL(v string)`

SetTargetURL sets TargetURL field to given value.

### HasTargetURL

`func (o *StatusCheckResult) HasTargetURL() bool`

HasTargetURL returns a boolean if a field has been set.

### SetTargetURLNil

`func (o *StatusCheckResult) SetTargetURLNil(b bool)`

 SetTargetURLNil sets the value for TargetURL to be an explicit nil

### UnsetTargetURL
`func (o *StatusCheckResult) UnsetTargetURL()`

UnsetTargetURL ensures that no value is present for TargetURL, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


