# CreateDepotFilePushRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**IsDenyForAllUser** | **bool** | 拒绝所有人推送 | 
**Pattern** | **string** | 文件路径 | 

## Methods

### NewCreateDepotFilePushRuleRequest

`func NewCreateDepotFilePushRuleRequest(depotPath string, isDenyForAllUser bool, pattern string, ) *CreateDepotFilePushRuleRequest`

NewCreateDepotFilePushRuleRequest instantiates a new CreateDepotFilePushRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDepotFilePushRuleRequestWithDefaults

`func NewCreateDepotFilePushRuleRequestWithDefaults() *CreateDepotFilePushRuleRequest`

NewCreateDepotFilePushRuleRequestWithDefaults instantiates a new CreateDepotFilePushRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *CreateDepotFilePushRuleRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateDepotFilePushRuleRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateDepotFilePushRuleRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetIsDenyForAllUser

`func (o *CreateDepotFilePushRuleRequest) GetIsDenyForAllUser() bool`

GetIsDenyForAllUser returns the IsDenyForAllUser field if non-nil, zero value otherwise.

### GetIsDenyForAllUserOk

`func (o *CreateDepotFilePushRuleRequest) GetIsDenyForAllUserOk() (*bool, bool)`

GetIsDenyForAllUserOk returns a tuple with the IsDenyForAllUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDenyForAllUser

`func (o *CreateDepotFilePushRuleRequest) SetIsDenyForAllUser(v bool)`

SetIsDenyForAllUser sets IsDenyForAllUser field to given value.


### GetPattern

`func (o *CreateDepotFilePushRuleRequest) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateDepotFilePushRuleRequest) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateDepotFilePushRuleRequest) SetPattern(v string)`

SetPattern sets Pattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


