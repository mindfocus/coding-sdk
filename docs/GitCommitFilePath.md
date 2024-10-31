# GitCommitFilePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPath** | Pointer to **NullableString** | 文件改动类型为 move 时，（移动、重命名），文件的新路径 | [optional] [default to ""]
**Path** | Pointer to **NullableString** | 改动文件的路径 | [optional] [default to ""]
**Type** | Pointer to **NullableString** | 文件改动类型 add update delete move | [optional] [default to ""]

## Methods

### NewGitCommitFilePath

`func NewGitCommitFilePath() *GitCommitFilePath`

NewGitCommitFilePath instantiates a new GitCommitFilePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCommitFilePathWithDefaults

`func NewGitCommitFilePathWithDefaults() *GitCommitFilePath`

NewGitCommitFilePathWithDefaults instantiates a new GitCommitFilePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPath

`func (o *GitCommitFilePath) GetNewPath() string`

GetNewPath returns the NewPath field if non-nil, zero value otherwise.

### GetNewPathOk

`func (o *GitCommitFilePath) GetNewPathOk() (*string, bool)`

GetNewPathOk returns a tuple with the NewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPath

`func (o *GitCommitFilePath) SetNewPath(v string)`

SetNewPath sets NewPath field to given value.

### HasNewPath

`func (o *GitCommitFilePath) HasNewPath() bool`

HasNewPath returns a boolean if a field has been set.

### SetNewPathNil

`func (o *GitCommitFilePath) SetNewPathNil(b bool)`

 SetNewPathNil sets the value for NewPath to be an explicit nil

### UnsetNewPath
`func (o *GitCommitFilePath) UnsetNewPath()`

UnsetNewPath ensures that no value is present for NewPath, not even an explicit nil
### GetPath

`func (o *GitCommitFilePath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitCommitFilePath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitCommitFilePath) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitCommitFilePath) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GitCommitFilePath) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GitCommitFilePath) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetType

`func (o *GitCommitFilePath) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitCommitFilePath) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitCommitFilePath) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GitCommitFilePath) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *GitCommitFilePath) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *GitCommitFilePath) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


