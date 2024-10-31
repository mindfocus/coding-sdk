# GitReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | 分支或标签的名字 | [optional] [default to ""]
**Type** | Pointer to **string** | all：全部，branch：分支，tag：标签 | [optional] [default to ""]

## Methods

### NewGitReference

`func NewGitReference() *GitReference`

NewGitReference instantiates a new GitReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitReferenceWithDefaults

`func NewGitReferenceWithDefaults() *GitReference`

NewGitReferenceWithDefaults instantiates a new GitReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GitReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitReference) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


