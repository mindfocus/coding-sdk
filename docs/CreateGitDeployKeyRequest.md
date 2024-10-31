# CreateGitDeployKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowWrite** | **bool** | 是否授予写入权限 | 
**Content** | **string** | SSH key | 
**DepotId** | **int64** | 仓库 Id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**ExpirationDate** | Pointer to **string** | 过期时间，不填则为永不过期 | [optional] 
**Title** | **string** | 部署公钥标题 | 

## Methods

### NewCreateGitDeployKeyRequest

`func NewCreateGitDeployKeyRequest(allowWrite bool, content string, depotId int64, title string, ) *CreateGitDeployKeyRequest`

NewCreateGitDeployKeyRequest instantiates a new CreateGitDeployKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitDeployKeyRequestWithDefaults

`func NewCreateGitDeployKeyRequestWithDefaults() *CreateGitDeployKeyRequest`

NewCreateGitDeployKeyRequestWithDefaults instantiates a new CreateGitDeployKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowWrite

`func (o *CreateGitDeployKeyRequest) GetAllowWrite() bool`

GetAllowWrite returns the AllowWrite field if non-nil, zero value otherwise.

### GetAllowWriteOk

`func (o *CreateGitDeployKeyRequest) GetAllowWriteOk() (*bool, bool)`

GetAllowWriteOk returns a tuple with the AllowWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWrite

`func (o *CreateGitDeployKeyRequest) SetAllowWrite(v bool)`

SetAllowWrite sets AllowWrite field to given value.


### GetContent

`func (o *CreateGitDeployKeyRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateGitDeployKeyRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateGitDeployKeyRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetDepotId

`func (o *CreateGitDeployKeyRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitDeployKeyRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitDeployKeyRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitDeployKeyRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitDeployKeyRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitDeployKeyRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitDeployKeyRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetExpirationDate

`func (o *CreateGitDeployKeyRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *CreateGitDeployKeyRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *CreateGitDeployKeyRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *CreateGitDeployKeyRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetTitle

`func (o *CreateGitDeployKeyRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateGitDeployKeyRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateGitDeployKeyRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


