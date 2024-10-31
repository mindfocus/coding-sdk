# CreateGitProtectedTagRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 Id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**Rule** | **[]string** | 保护规则列表 | 

## Methods

### NewCreateGitProtectedTagRulesRequest

`func NewCreateGitProtectedTagRulesRequest(depotId int64, rule []string, ) *CreateGitProtectedTagRulesRequest`

NewCreateGitProtectedTagRulesRequest instantiates a new CreateGitProtectedTagRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitProtectedTagRulesRequestWithDefaults

`func NewCreateGitProtectedTagRulesRequestWithDefaults() *CreateGitProtectedTagRulesRequest`

NewCreateGitProtectedTagRulesRequestWithDefaults instantiates a new CreateGitProtectedTagRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateGitProtectedTagRulesRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitProtectedTagRulesRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitProtectedTagRulesRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitProtectedTagRulesRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitProtectedTagRulesRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitProtectedTagRulesRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitProtectedTagRulesRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetRule

`func (o *CreateGitProtectedTagRulesRequest) GetRule() []string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateGitProtectedTagRulesRequest) GetRuleOk() (*[]string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateGitProtectedTagRulesRequest) SetRule(v []string)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


