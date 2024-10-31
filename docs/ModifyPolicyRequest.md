# ModifyPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | 显示名称 | [optional] 
**Description** | Pointer to **string** | 描述 | [optional] 
**Id** | **int64** | 权限组 ID | 
**Name** | Pointer to **string** | 名称 | [optional] 
**PolicyDocument** | [**PolicyDocument**](PolicyDocument.md) |  | 
**ResourceType** | **[]string** | 适用的资源类型 | 

## Methods

### NewModifyPolicyRequest

`func NewModifyPolicyRequest(id int64, policyDocument PolicyDocument, resourceType []string, ) *ModifyPolicyRequest`

NewModifyPolicyRequest instantiates a new ModifyPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyPolicyRequestWithDefaults

`func NewModifyPolicyRequestWithDefaults() *ModifyPolicyRequest`

NewModifyPolicyRequestWithDefaults instantiates a new ModifyPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ModifyPolicyRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModifyPolicyRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModifyPolicyRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModifyPolicyRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDescription

`func (o *ModifyPolicyRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModifyPolicyRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModifyPolicyRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModifyPolicyRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModifyPolicyRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyPolicyRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyPolicyRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ModifyPolicyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModifyPolicyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModifyPolicyRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModifyPolicyRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyDocument

`func (o *ModifyPolicyRequest) GetPolicyDocument() PolicyDocument`

GetPolicyDocument returns the PolicyDocument field if non-nil, zero value otherwise.

### GetPolicyDocumentOk

`func (o *ModifyPolicyRequest) GetPolicyDocumentOk() (*PolicyDocument, bool)`

GetPolicyDocumentOk returns a tuple with the PolicyDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDocument

`func (o *ModifyPolicyRequest) SetPolicyDocument(v PolicyDocument)`

SetPolicyDocument sets PolicyDocument field to given value.


### GetResourceType

`func (o *ModifyPolicyRequest) GetResourceType() []string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModifyPolicyRequest) GetResourceTypeOk() (*[]string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModifyPolicyRequest) SetResourceType(v []string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


