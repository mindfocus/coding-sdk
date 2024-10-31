# GitTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to [**GitAllTagCommit**](GitAllTagCommit.md) |  | [optional] 
**Message** | Pointer to **string** | tag信息 | [optional] [default to ""]
**TagName** | Pointer to **string** | tag的名称 | [optional] [default to ""]

## Methods

### NewGitTag

`func NewGitTag() *GitTag`

NewGitTag instantiates a new GitTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTagWithDefaults

`func NewGitTagWithDefaults() *GitTag`

NewGitTagWithDefaults instantiates a new GitTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *GitTag) GetCommit() GitAllTagCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *GitTag) GetCommitOk() (*GitAllTagCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *GitTag) SetCommit(v GitAllTagCommit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *GitTag) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetMessage

`func (o *GitTag) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitTag) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitTag) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GitTag) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTagName

`func (o *GitTag) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *GitTag) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *GitTag) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *GitTag) HasTagName() bool`

HasTagName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


