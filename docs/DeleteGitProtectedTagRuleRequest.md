# DeleteGitProtectedTagRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库 Id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**Rule** | **string** | 保护规则 | 

## Methods

### NewDeleteGitProtectedTagRuleRequest

`func NewDeleteGitProtectedTagRuleRequest(depotId int64, rule string, ) *DeleteGitProtectedTagRuleRequest`

NewDeleteGitProtectedTagRuleRequest instantiates a new DeleteGitProtectedTagRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGitProtectedTagRuleRequestWithDefaults

`func NewDeleteGitProtectedTagRuleRequestWithDefaults() *DeleteGitProtectedTagRuleRequest`

NewDeleteGitProtectedTagRuleRequestWithDefaults instantiates a new DeleteGitProtectedTagRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DeleteGitProtectedTagRuleRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DeleteGitProtectedTagRuleRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DeleteGitProtectedTagRuleRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DeleteGitProtectedTagRuleRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteGitProtectedTagRuleRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteGitProtectedTagRuleRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DeleteGitProtectedTagRuleRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetRule

`func (o *DeleteGitProtectedTagRuleRequest) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *DeleteGitProtectedTagRuleRequest) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *DeleteGitProtectedTagRuleRequest) SetRule(v string)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


