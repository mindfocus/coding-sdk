# GitTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **NullableString** | 文件或目录权限 | [optional] [default to ""]
**Path** | Pointer to **NullableString** | 路径 | [optional] [default to ""]
**Sha** | Pointer to **NullableString** | 哈希值 | [optional] [default to ""]
**Type** | Pointer to **NullableString** | 类型 | [optional] [default to ""]

## Methods

### NewGitTree

`func NewGitTree() *GitTree`

NewGitTree instantiates a new GitTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTreeWithDefaults

`func NewGitTreeWithDefaults() *GitTree`

NewGitTreeWithDefaults instantiates a new GitTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GitTree) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GitTree) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GitTree) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GitTree) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *GitTree) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *GitTree) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetPath

`func (o *GitTree) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitTree) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitTree) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitTree) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GitTree) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GitTree) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetSha

`func (o *GitTree) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitTree) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitTree) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitTree) HasSha() bool`

HasSha returns a boolean if a field has been set.

### SetShaNil

`func (o *GitTree) SetShaNil(b bool)`

 SetShaNil sets the value for Sha to be an explicit nil

### UnsetSha
`func (o *GitTree) UnsetSha()`

UnsetSha ensures that no value is present for Sha, not even an explicit nil
### GetType

`func (o *GitTree) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitTree) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitTree) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitTree) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GitTree) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GitTree) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


