# CreatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | 显示名称 | 
**Description** | Pointer to **string** | 描述 | [optional] 
**Name** | **string** | 名称 | 
**PolicyDocument** | [**PolicyDocument**](PolicyDocument.md) |  | 
**PolicyType** | **string** | 权限组类型：IDENTITY | RESOURCE | 
**ResourceType** | **[]string** | 适用的资源类型 | 

## Methods

### NewCreatePolicyRequest

`func NewCreatePolicyRequest(alias string, name string, policyDocument PolicyDocument, policyType string, resourceType []string, ) *CreatePolicyRequest`

NewCreatePolicyRequest instantiates a new CreatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePolicyRequestWithDefaults

`func NewCreatePolicyRequestWithDefaults() *CreatePolicyRequest`

NewCreatePolicyRequestWithDefaults instantiates a new CreatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *CreatePolicyRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CreatePolicyRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CreatePolicyRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetDescription

`func (o *CreatePolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreatePolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePolicyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPolicyDocument

`func (o *CreatePolicyRequest) GetPolicyDocument() PolicyDocument`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *CreatePolicyRequest) GetPolicyDocumentOk() (*PolicyDocument, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *CreatePolicyRequest) SetPolicyDocument(v PolicyDocument)`

SetPolicyDocument sets PolicyDocument field to given value.


### GetPolicyType

`func (o *CreatePolicyRequest) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CreatePolicyRequest) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *CreatePolicyRequest) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.


### GetResourceType

`func (o *CreatePolicyRequest) GetResourceType() []string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreatePolicyRequest) GetResourceTypeOk() (*[]string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreatePolicyRequest) SetResourceType(v []string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


