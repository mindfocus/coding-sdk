# DescribeBranchProtectionMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchProtectionId** | **int64** | 保护分支规则id | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepptPath二选一即可 | [optional] 

## Methods

### NewDescribeBranchProtectionMembersRequest

`func NewDescribeBranchProtectionMembersRequest(branchProtectionId int64, depotId int64, ) *DescribeBranchProtectionMembersRequest`

NewDescribeBranchProtectionMembersRequest instantiates a new DescribeBranchProtectionMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeBranchProtectionMembersRequestWithDefaults

`func NewDescribeBranchProtectionMembersRequestWithDefaults() *DescribeBranchProtectionMembersRequest`

NewDescribeBranchProtectionMembersRequestWithDefaults instantiates a new DescribeBranchProtectionMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchProtectionId

`func (o *DescribeBranchProtectionMembersRequest) GetBranchProtectionId() int64`

GetBranchProtectionId returns the BranchProtectionId field if non-nil, zero value otherwise.

### GetBranchProtectionIdOk

`func (o *DescribeBranchProtectionMembersRequest) GetBranchProtectionIdOk() (*int64, bool)`

GetBranchProtectionIdOk returns a tuple with the BranchProtectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchProtectionId

`func (o *DescribeBranchProtectionMembersRequest) SetBranchProtectionId(v int64)`

SetBranchProtectionId sets BranchProtectionId field to given value.


### GetDepotId

`func (o *DescribeBranchProtectionMembersRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeBranchProtectionMembersRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeBranchProtectionMembersRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeBranchProtectionMembersRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeBranchProtectionMembersRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeBranchProtectionMembersRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeBranchProtectionMembersRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


