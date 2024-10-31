# CreateGitProtectedTagRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 Id | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 
**Rule** | **string** | 保护规则 | 

## Methods

### NewCreateGitProtectedTagRuleRequest

`func NewCreateGitProtectedTagRuleRequest(depotId int64, rule string, ) *CreateGitProtectedTagRuleRequest`

NewCreateGitProtectedTagRuleRequest instantiates a new CreateGitProtectedTagRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitProtectedTagRuleRequestWithDefaults

`func NewCreateGitProtectedTagRuleRequestWithDefaults() *CreateGitProtectedTagRuleRequest`

NewCreateGitProtectedTagRuleRequestWithDefaults instantiates a new CreateGitProtectedTagRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateGitProtectedTagRuleRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitProtectedTagRuleRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitProtectedTagRuleRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitProtectedTagRuleRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitProtectedTagRuleRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitProtectedTagRuleRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitProtectedTagRuleRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetRule

`func (o *CreateGitProtectedTagRuleRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateGitProtectedTagRuleRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateGitProtectedTagRuleRequest) SetRule(v string)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


