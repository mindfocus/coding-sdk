# ModifyGitDepotUnarchiveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **string** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 

## Methods

### NewModifyGitDepotUnarchiveRequest

`func NewModifyGitDepotUnarchiveRequest(depotId string, ) *ModifyGitDepotUnarchiveRequest`

NewModifyGitDepotUnarchiveRequest instantiates a new ModifyGitDepotUnarchiveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyGitDepotUnarchiveRequestWithDefaults

`func NewModifyGitDepotUnarchiveRequestWithDefaults() *ModifyGitDepotUnarchiveRequest`

NewModifyGitDepotUnarchiveRequestWithDefaults instantiates a new ModifyGitDepotUnarchiveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *ModifyGitDepotUnarchiveRequest) GetDepotId() string`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ModifyGitDepotUnarchiveRequest) GetDepotIdOk() (*string, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ModifyGitDepotUnarchiveRequest) SetDepotId(v string)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *ModifyGitDepotUnarchiveRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *ModifyGitDepotUnarchiveRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *ModifyGitDepotUnarchiveRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *ModifyGitDepotUnarchiveRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


