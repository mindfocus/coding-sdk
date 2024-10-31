# CreateGitDepotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotName** | **string** | 仓库名称 | 
**Description** | Pointer to **string** | 仓库的描述信息 | [optional] 
**ProjectId** | **int64** | 项目id | 
**Shared** | Pointer to **bool** | 仓库是否允许公开访问 | [optional] 

## Methods

### NewCreateGitDepotRequest

`func NewCreateGitDepotRequest(depotName string, projectId int64, ) *CreateGitDepotRequest`

NewCreateGitDepotRequest instantiates a new CreateGitDepotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitDepotRequestWithDefaults

`func NewCreateGitDepotRequestWithDefaults() *CreateGitDepotRequest`

NewCreateGitDepotRequestWithDefaults instantiates a new CreateGitDepotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotName

`func (o *CreateGitDepotRequest) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *CreateGitDepotRequest) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *CreateGitDepotRequest) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.


### GetDescription

`func (o *CreateGitDepotRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGitDepotRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGitDepotRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGitDepotRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProjectId

`func (o *CreateGitDepotRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateGitDepotRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateGitDepotRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetShared

`func (o *CreateGitDepotRequest) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *CreateGitDepotRequest) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *CreateGitDepotRequest) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *CreateGitDepotRequest) HasShared() bool`

HasShared returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


