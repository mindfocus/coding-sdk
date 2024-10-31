# GitRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotatedTag** | Pointer to **bool** | 是否是附注标签 | [optional] [default to false]
**DisplayName** | Pointer to **string** | 展示名 | [optional] [default to ""]
**FullMessage** | Pointer to **string** | 最后一次提交全部信息 | [optional] [default to ""]
**Name** | Pointer to **string** | 名字 | [optional] [default to ""]
**ObjectId** | Pointer to **string** | 对象id | [optional] [default to ""]
**RefObjectId** | Pointer to **string** | 分支objectId | [optional] [default to ""]
**ShortMessage** | Pointer to **string** | 最后一次提交简短信息 | [optional] [default to ""]

## Methods

### NewGitRef

`func NewGitRef() *GitRef`

NewGitRef instantiates a new GitRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRefWithDefaults

`func NewGitRefWithDefaults() *GitRef`

NewGitRefWithDefaults instantiates a new GitRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotatedTag

`func (o *GitRef) GetAnnotatedTag() bool`

GetAnnotatedTag returns the AnnotatedTag field if non-nil, zero value otherwise.

### GetAnnotatedTagOk

`func (o *GitRef) GetAnnotatedTagOk() (*bool, bool)`

GetAnnotatedTagOk returns a tuple with the AnnotatedTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotatedTag

`func (o *GitRef) SetAnnotatedTag(v bool)`

SetAnnotatedTag sets AnnotatedTag field to given value.

### HasAnnotatedTag

`func (o *GitRef) HasAnnotatedTag() bool`

HasAnnotatedTag returns a boolean if a field has been set.

### GetDisplayName

`func (o *GitRef) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GitRef) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GitRef) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GitRef) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFullMessage

`func (o *GitRef) GetFullMessage() string`

GetFullMessage returns the FullMessage field if non-nil, zero value otherwise.

### GetFullMessageOk

`func (o *GitRef) GetFullMessageOk() (*string, bool)`

GetFullMessageOk returns a tuple with the FullMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullMessage

`func (o *GitRef) SetFullMessage(v string)`

SetFullMessage sets FullMessage field to given value.

### HasFullMessage

`func (o *GitRef) HasFullMessage() bool`

HasFullMessage returns a boolean if a field has been set.

### GetName

`func (o *GitRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitRef) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GitRef) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjectId

`func (o *GitRef) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *GitRef) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *GitRef) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *GitRef) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetRefObjectId

`func (o *GitRef) GetRefObjectId() string`

GetRefObjectId returns the RefObjectId field if non-nil, zero value otherwise.

### GetRefObjectIdOk

`func (o *GitRef) GetRefObjectIdOk() (*string, bool)`

GetRefObjectIdOk returns a tuple with the RefObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefObjectId

`func (o *GitRef) SetRefObjectId(v string)`

SetRefObjectId sets RefObjectId field to given value.

### HasRefObjectId

`func (o *GitRef) HasRefObjectId() bool`

HasRefObjectId returns a boolean if a field has been set.

### GetShortMessage

`func (o *GitRef) GetShortMessage() string`

GetShortMessage returns the ShortMessage field if non-nil, zero value otherwise.

### GetShortMessageOk

`func (o *GitRef) GetShortMessageOk() (*string, bool)`

GetShortMessageOk returns a tuple with the ShortMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMessage

`func (o *GitRef) SetShortMessage(v string)`

SetShortMessage sets ShortMessage field to given value.

### HasShortMessage

`func (o *GitRef) HasShortMessage() bool`

HasShortMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


