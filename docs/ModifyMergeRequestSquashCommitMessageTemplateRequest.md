# ModifyMergeRequestSquashCommitMessageTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitMessageTemplate** | **string** | 合并请求合并提交消息模版 | 
**DepotPath** | **string** | 仓库路径 | 

## Methods

### NewModifyMergeRequestSquashCommitMessageTemplateRequest

`func NewModifyMergeRequestSquashCommitMessageTemplateRequest(commitMessageTemplate string, depotPath string, ) *ModifyMergeRequestSquashCommitMessageTemplateRequest`

NewModifyMergeRequestSquashCommitMessageTemplateRequest instantiates a new ModifyMergeRequestSquashCommitMessageTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyMergeRequestSquashCommitMessageTemplateRequestWithDefaults

`func NewModifyMergeRequestSquashCommitMessageTemplateRequestWithDefaults() *ModifyMergeRequestSquashCommitMessageTemplateRequest`

NewModifyMergeRequestSquashCommitMessageTemplateRequestWithDefaults instantiates a new ModifyMergeRequestSquashCommitMessageTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitMessageTemplate

`func (o *ModifyMergeRequestSquashCommitMessageTemplateRequest) GetCommitMessageTemplate() string`

GetCommitMessageTemplate returns the CommitMessageTemplate field if non-nil, zero value otherwise.

### GetCommitMessageTemplateOk

`func (o *ModifyMergeRequestSquashCommitMessageTemplateRequest) GetCommitMessageTemplateOk() (*string, bool)`

GetCommitMessageTemplateOk returns a tuple with the CommitMessageTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessageTemplate

`func (o *ModifyMergeRequestSquashCommitMessageTemplateRequest) SetCommitMessageTemplate(v string)`

SetCommitMessageTemplate sets CommitMessageTemplate field to given value.


### GetDepotPath

`func (o *ModifyMergeRequestSquashCommitMessageTemplateRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyMergeRequestSquashCommitMessageTemplateRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyMergeRequestSquashCommitMessageTemplateRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


