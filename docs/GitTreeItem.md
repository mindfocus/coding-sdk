# GitTreeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | 文件类型 | [optional] [default to ""]
**Name** | Pointer to **string** | 文件名 | [optional] [default to ""]
**Path** | Pointer to **string** | 文件路径 | [optional] [default to ""]
**Sha** | Pointer to **string** | commitID | [optional] [default to ""]

## Methods

### NewGitTreeItem

`func NewGitTreeItem() *GitTreeItem`

NewGitTreeItem instantiates a new GitTreeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTreeItemWithDefaults

`func NewGitTreeItemWithDefaults() *GitTreeItem`

NewGitTreeItemWithDefaults instantiates a new GitTreeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GitTreeItem) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GitTreeItem) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GitTreeItem) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GitTreeItem) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *GitTreeItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitTreeItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitTreeItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitTreeItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *GitTreeItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GitTreeItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GitTreeItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GitTreeItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSha

`func (o *GitTreeItem) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitTreeItem) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitTreeItem) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *GitTreeItem) HasSha() bool`

HasSha returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


