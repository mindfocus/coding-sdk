# DifferentOfCommit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | Pointer to **NullableString** | 修改类型 | [optional] [default to ""]
**Deletions** | Pointer to **int64** | 删除的行数 | [optional] 
**Insertions** | Pointer to **int64** | 新增的行数 | [optional] 
**Name** | Pointer to **string** | 提交的名称 | [optional] [default to ""]
**Path** | Pointer to **string** | 文件路径 | [optional] [default to ""]

## Methods

### NewDifferentOfCommit

`func NewDifferentOfCommit() *DifferentOfCommit`

NewDifferentOfCommit instantiates a new DifferentOfCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDifferentOfCommitWithDefaults

`func NewDifferentOfCommitWithDefaults() *DifferentOfCommit`

NewDifferentOfCommitWithDefaults instantiates a new DifferentOfCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *DifferentOfCommit) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *DifferentOfCommit) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *DifferentOfCommit) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *DifferentOfCommit) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### SetChangeTypeNil

`func (o *DifferentOfCommit) SetChangeTypeNil(b bool)`

 SetChangeTypeNil sets the value for ChangeType to be an explicit nil

### UnsetChangeType
`func (o *DifferentOfCommit) UnsetChangeType()`

UnsetChangeType ensures that no value is present for ChangeType, not even an explicit nil
### GetDeletions

`func (o *DifferentOfCommit) GetDeletions() int64`

GetDeletions returns the Deletions field if non-nil, zero value otherwise.

### GetDeletionsOk

`func (o *DifferentOfCommit) GetDeletionsOk() (*int64, bool)`

GetDeletionsOk returns a tuple with the Deletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletions

`func (o *DifferentOfCommit) SetDeletions(v int64)`

SetDeletions sets Deletions field to given value.

### HasDeletions

`func (o *DifferentOfCommit) HasDeletions() bool`

HasDeletions returns a boolean if a field has been set.

### GetInsertions

`func (o *DifferentOfCommit) GetInsertions() int64`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *DifferentOfCommit) GetInsertionsOk() (*int64, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *DifferentOfCommit) SetInsertions(v int64)`

SetInsertions sets Insertions field to given value.

### HasInsertions

`func (o *DifferentOfCommit) HasInsertions() bool`

HasInsertions returns a boolean if a field has been set.

### GetName

`func (o *DifferentOfCommit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DifferentOfCommit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DifferentOfCommit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DifferentOfCommit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *DifferentOfCommit) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DifferentOfCommit) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DifferentOfCommit) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DifferentOfCommit) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


