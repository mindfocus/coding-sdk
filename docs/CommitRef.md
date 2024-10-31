# CommitRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | 分支明称/标签名称 | [optional] [default to ""]
**Type** | Pointer to **string** | 分支/标签 | [optional] [default to ""]

## Methods

### NewCommitRef

`func NewCommitRef() *CommitRef`

NewCommitRef instantiates a new CommitRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitRefWithDefaults

`func NewCommitRefWithDefaults() *CommitRef`

NewCommitRefWithDefaults instantiates a new CommitRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CommitRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CommitRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CommitRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CommitRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CommitRef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CommitRef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CommitRef) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CommitRef) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


