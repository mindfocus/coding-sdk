# ModifyDepotFilePushRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**FilePushRuleId** | **int64** | 文件推送规则 ID | 
**IsDenyForAllUser** | **bool** | 拒绝所有人推送 | 

## Methods

### NewModifyDepotFilePushRuleRequest

`func NewModifyDepotFilePushRuleRequest(depotPath string, filePushRuleId int64, isDenyForAllUser bool, ) *ModifyDepotFilePushRuleRequest`

NewModifyDepotFilePushRuleRequest instantiates a new ModifyDepotFilePushRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyDepotFilePushRuleRequestWithDefaults

`func NewModifyDepotFilePushRuleRequestWithDefaults() *ModifyDepotFilePushRuleRequest`

NewModifyDepotFilePushRuleRequestWithDefaults instantiates a new ModifyDepotFilePushRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *ModifyDepotFilePushRuleRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyDepotFilePushRuleRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyDepotFilePushRuleRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetFilePushRuleId

`func (o *ModifyDepotFilePushRuleRequest) GetFilePushRuleId() int64`

GetFilePushRuleId returns the FilePushRuleId field if non-nil, zero value otherwise.

### GetFilePushRuleIdOk

`func (o *ModifyDepotFilePushRuleRequest) GetFilePushRuleIdOk() (*int64, bool)`

GetFilePushRuleIdOk returns a tuple with the FilePushRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePushRuleId

`func (o *ModifyDepotFilePushRuleRequest) SetFilePushRuleId(v int64)`

SetFilePushRuleId sets FilePushRuleId field to given value.


### GetIsDenyForAllUser

`func (o *ModifyDepotFilePushRuleRequest) GetIsDenyForAllUser() bool`

GetIsDenyForAllUser returns the IsDenyForAllUser field if non-nil, zero value otherwise.

### GetIsDenyForAllUserOk

`func (o *ModifyDepotFilePushRuleRequest) GetIsDenyForAllUserOk() (*bool, bool)`

GetIsDenyForAllUserOk returns a tuple with the IsDenyForAllUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDenyForAllUser

`func (o *ModifyDepotFilePushRuleRequest) SetIsDenyForAllUser(v bool)`

SetIsDenyForAllUser sets IsDenyForAllUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


