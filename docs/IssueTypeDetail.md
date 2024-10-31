# IssueTypeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | 描述 | [optional] [default to ""]
**Id** | Pointer to **int64** | 事项类型 ID | [optional] 
**IsSystem** | Pointer to **bool** | 是否是系统类型 | [optional] [default to false]
**IssueType** | Pointer to **string** | 事项类型大类 | [optional] [default to ""]
**Name** | Pointer to **string** | 事项类型名称 | [optional] [default to ""]

## Methods

### NewIssueTypeDetail

`func NewIssueTypeDetail() *IssueTypeDetail`

NewIssueTypeDetail instantiates a new IssueTypeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeDetailWithDefaults

`func NewIssueTypeDetailWithDefaults() *IssueTypeDetail`

NewIssueTypeDetailWithDefaults instantiates a new IssueTypeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IssueTypeDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IssueTypeDetail) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueTypeDetail) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueTypeDetail) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IssueTypeDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsSystem

`func (o *IssueTypeDetail) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *IssueTypeDetail) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *IssueTypeDetail) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *IssueTypeDetail) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetIssueType

`func (o *IssueTypeDetail) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *IssueTypeDetail) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *IssueTypeDetail) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *IssueTypeDetail) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeDetail) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


