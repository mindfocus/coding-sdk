# BoundExternalDepotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialUUID** | Pointer to **string** | 当授权类型为USERNAME_PASSWORD时，填入和当前仓库用户名密码关联的凭证UUID | [optional] 
**DepotId** | Pointer to **int64** | 如果是跨项目代码仓库，则此字段必填，为源代码仓库id | [optional] 
**DepotType** | **string** | 仓库类型 | 
**ExternalDepotAddress** | **string** | 外部仓库标识 | 
**GrantType** | Pointer to **string** | 如果是跨项目代码仓库，则需要填入“CODING_PERSONAL_CREDENTIAL”，其他外部仓库填入“OAUTH” | [optional] 
**ProjectId** | **int32** | 项目 Id | 
**RepoUrl** | Pointer to **string** | 当授权类型为USERNAME_PASSWORD时，填入外部仓库http格式的仓库地址 | [optional] 
**WebHook** | **bool** | 是否开启 WebHook 一般都填写 true | 

## Methods

### NewBoundExternalDepotRequest

`func NewBoundExternalDepotRequest(depotType string, externalDepotAddress string, projectId int32, webHook bool, ) *BoundExternalDepotRequest`

NewBoundExternalDepotRequest instantiates a new BoundExternalDepotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoundExternalDepotRequestWithDefaults

`func NewBoundExternalDepotRequestWithDefaults() *BoundExternalDepotRequest`

NewBoundExternalDepotRequestWithDefaults instantiates a new BoundExternalDepotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialUUID

`func (o *BoundExternalDepotRequest) GetCredentialUUID() string`

GetCredentialUUID returns the CredentialUUID field if non-nil, zero value otherwise.

### GetCredentialUUIDOk

`func (o *BoundExternalDepotRequest) GetCredentialUUIDOk() (*string, bool)`

GetCredentialUUIDOk returns a tuple with the CredentialUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialUUID

`func (o *BoundExternalDepotRequest) SetCredentialUUID(v string)`

SetCredentialUUID sets CredentialUUID field to given value.

### HasCredentialUUID

`func (o *BoundExternalDepotRequest) HasCredentialUUID() bool`

HasCredentialUUID returns a boolean if a field has been set.

### GetDepotId

`func (o *BoundExternalDepotRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *BoundExternalDepotRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *BoundExternalDepotRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *BoundExternalDepotRequest) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetDepotType

`func (o *BoundExternalDepotRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *BoundExternalDepotRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *BoundExternalDepotRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetExternalDepotAddress

`func (o *BoundExternalDepotRequest) GetExternalDepotAddress() string`

GetExternalDepotAddress returns the ExternalDepotAddress field if non-nil, zero value otherwise.

### GetExternalDepotAddressOk

`func (o *BoundExternalDepotRequest) GetExternalDepotAddressOk() (*string, bool)`

GetExternalDepotAddressOk returns a tuple with the ExternalDepotAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalDepotAddress

`func (o *BoundExternalDepotRequest) SetExternalDepotAddress(v string)`

SetExternalDepotAddress sets ExternalDepotAddress field to given value.


### GetGrantType

`func (o *BoundExternalDepotRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *BoundExternalDepotRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *BoundExternalDepotRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *BoundExternalDepotRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetProjectId

`func (o *BoundExternalDepotRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BoundExternalDepotRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BoundExternalDepotRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetRepoUrl

`func (o *BoundExternalDepotRequest) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *BoundExternalDepotRequest) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *BoundExternalDepotRequest) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *BoundExternalDepotRequest) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetWebHook

`func (o *BoundExternalDepotRequest) GetWebHook() bool`

GetWebHook returns the WebHook field if non-nil, zero value otherwise.

### GetWebHookOk

`func (o *BoundExternalDepotRequest) GetWebHookOk() (*bool, bool)`

GetWebHookOk returns a tuple with the WebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHook

`func (o *BoundExternalDepotRequest) SetWebHook(v bool)`

SetWebHook sets WebHook field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


