# DeleteGitDeployKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeployKeyId** | **int64** | SSH Key Id | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 

## Methods

### NewDeleteGitDeployKeyRequest

`func NewDeleteGitDeployKeyRequest(deployKeyId int64, depotId int64, ) *DeleteGitDeployKeyRequest`

NewDeleteGitDeployKeyRequest instantiates a new DeleteGitDeployKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGitDeployKeyRequestWithDefaults

`func NewDeleteGitDeployKeyRequestWithDefaults() *DeleteGitDeployKeyRequest`

NewDeleteGitDeployKeyRequestWithDefaults instantiates a new DeleteGitDeployKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployKeyId

`func (o *DeleteGitDeployKeyRequest) GetDeployKeyId() int64`

GetDeployKeyId returns the DeployKeyId field if non-nil, zero value otherwise.

### GetDeployKeyIdOk

`func (o *DeleteGitDeployKeyRequest) GetDeployKeyIdOk() (*int64, bool)`

GetDeployKeyIdOk returns a tuple with the DeployKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployKeyId

`func (o *DeleteGitDeployKeyRequest) SetDeployKeyId(v int64)`

SetDeployKeyId sets DeployKeyId field to given value.


### GetDepotId

`func (o *DeleteGitDeployKeyRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DeleteGitDeployKeyRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DeleteGitDeployKeyRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DeleteGitDeployKeyRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteGitDeployKeyRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteGitDeployKeyRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DeleteGitDeployKeyRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


