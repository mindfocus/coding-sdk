# ModifyBranchProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchProtection** | [**BranchProtection**](BranchProtection.md) |  | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 

## Methods

### NewModifyBranchProtectionRequest

`func NewModifyBranchProtectionRequest(branchProtection BranchProtection, depotId int64, ) *ModifyBranchProtectionRequest`

NewModifyBranchProtectionRequest instantiates a new ModifyBranchProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyBranchProtectionRequestWithDefaults

`func NewModifyBranchProtectionRequestWithDefaults() *ModifyBranchProtectionRequest`

NewModifyBranchProtectionRequestWithDefaults instantiates a new ModifyBranchProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchProtection

`func (o *ModifyBranchProtectionRequest) GetBranchProtection() BranchProtection`

GetBranchProtection returns the BranchProtection field if non-nil, zero value otherwise.

### GetBranchProtectionOk

`func (o *ModifyBranchProtectionRequest) GetBranchProtectionOk() (*BranchProtection, bool)`

GetBranchProtectionOk returns a tuple with the BranchProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchProtection

`func (o *ModifyBranchProtectionRequest) SetBranchProtection(v BranchProtection)`

SetBranchProtection sets BranchProtection field to given value.


### GetDepotId

`func (o *ModifyBranchProtectionRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyBranchProtectionRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyBranchProtectionRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyBranchProtectionRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyBranchProtectionRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyBranchProtectionRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyBranchProtectionRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


