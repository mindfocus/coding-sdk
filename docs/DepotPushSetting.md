# DepotPushSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckCommitAuthor** | Pointer to **NullableBool** | 检查 Git 提交的提交者 (Committer) 和提交作者 (Author) 必须是已验证的邮箱。 | [optional] [default to false]
**CommitMessageMustMatchRegex** | Pointer to **NullableString** | Git 提交信息的格式校验 | [optional] [default to ""]
**DenyForcePush** | Pointer to **NullableBool** | 禁止强制推送 (Force Push) | [optional] [default to false]
**PushDenyFile** | Pointer to **NullableString** | 禁止推送的文件（文件类型用换行符隔开） | [optional] [default to ""]
**PushFileSize** | Pointer to **NullableString** | 开启单次提交的文件总大小限制，Git LFS 文件除外（单位MB） | [optional] [default to ""]

## Methods

### NewDepotPushSetting

`func NewDepotPushSetting() *DepotPushSetting`

NewDepotPushSetting instantiates a new DepotPushSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepotPushSettingWithDefaults

`func NewDepotPushSettingWithDefaults() *DepotPushSetting`

NewDepotPushSettingWithDefaults instantiates a new DepotPushSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckCommitAuthor

`func (o *DepotPushSetting) GetCheckCommitAuthor() bool`

GetCheckCommitAuthor returns the CheckCommitAuthor field if non-nil, zero value otherwise.

### GetCheckCommitAuthorOk

`func (o *DepotPushSetting) GetCheckCommitAuthorOk() (*bool, bool)`

GetCheckCommitAuthorOk returns a tuple with the CheckCommitAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCommitAuthor

`func (o *DepotPushSetting) SetCheckCommitAuthor(v bool)`

SetCheckCommitAuthor sets CheckCommitAuthor field to given value.

### HasCheckCommitAuthor

`func (o *DepotPushSetting) HasCheckCommitAuthor() bool`

HasCheckCommitAuthor returns a boolean if a field has been set.

### SetCheckCommitAuthorNil

`func (o *DepotPushSetting) SetCheckCommitAuthorNil(b bool)`

 SetCheckCommitAuthorNil sets the value for CheckCommitAuthor to be an explicit nil

### UnsetCheckCommitAuthor
`func (o *DepotPushSetting) UnsetCheckCommitAuthor()`

UnsetCheckCommitAuthor ensures that no value is present for CheckCommitAuthor, not even an explicit nil
### GetCommitMessageMustMatchRegex

`func (o *DepotPushSetting) GetCommitMessageMustMatchRegex() string`

GetCommitMessageMustMatchRegex returns the CommitMessageMustMatchRegex field if non-nil, zero value otherwise.

### GetCommitMessageMustMatchRegexOk

`func (o *DepotPushSetting) GetCommitMessageMustMatchRegexOk() (*string, bool)`

GetCommitMessageMustMatchRegexOk returns a tuple with the CommitMessageMustMatchRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessageMustMatchRegex

`func (o *DepotPushSetting) SetCommitMessageMustMatchRegex(v string)`

SetCommitMessageMustMatchRegex sets CommitMessageMustMatchRegex field to given value.

### HasCommitMessageMustMatchRegex

`func (o *DepotPushSetting) HasCommitMessageMustMatchRegex() bool`

HasCommitMessageMustMatchRegex returns a boolean if a field has been set.

### SetCommitMessageMustMatchRegexNil

`func (o *DepotPushSetting) SetCommitMessageMustMatchRegexNil(b bool)`

 SetCommitMessageMustMatchRegexNil sets the value for CommitMessageMustMatchRegex to be an explicit nil

### UnsetCommitMessageMustMatchRegex
`func (o *DepotPushSetting) UnsetCommitMessageMustMatchRegex()`

UnsetCommitMessageMustMatchRegex ensures that no value is present for CommitMessageMustMatchRegex, not even an explicit nil
### GetDenyForcePush

`func (o *DepotPushSetting) GetDenyForcePush() bool`

GetDenyForcePush returns the DenyForcePush field if non-nil, zero value otherwise.

### GetDenyForcePushOk

`func (o *DepotPushSetting) GetDenyForcePushOk() (*bool, bool)`

GetDenyForcePushOk returns a tuple with the DenyForcePush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyForcePush

`func (o *DepotPushSetting) SetDenyForcePush(v bool)`

SetDenyForcePush sets DenyForcePush field to given value.

### HasDenyForcePush

`func (o *DepotPushSetting) HasDenyForcePush() bool`

HasDenyForcePush returns a boolean if a field has been set.

### SetDenyForcePushNil

`func (o *DepotPushSetting) SetDenyForcePushNil(b bool)`

 SetDenyForcePushNil sets the value for DenyForcePush to be an explicit nil

### UnsetDenyForcePush
`func (o *DepotPushSetting) UnsetDenyForcePush()`

UnsetDenyForcePush ensures that no value is present for DenyForcePush, not even an explicit nil
### GetPushDenyFile

`func (o *DepotPushSetting) GetPushDenyFile() string`

GetPushDenyFile returns the PushDenyFile field if non-nil, zero value otherwise.

### GetPushDenyFileOk

`func (o *DepotPushSetting) GetPushDenyFileOk() (*string, bool)`

GetPushDenyFileOk returns a tuple with the PushDenyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushDenyFile

`func (o *DepotPushSetting) SetPushDenyFile(v string)`

SetPushDenyFile sets PushDenyFile field to given value.

### HasPushDenyFile

`func (o *DepotPushSetting) HasPushDenyFile() bool`

HasPushDenyFile returns a boolean if a field has been set.

### SetPushDenyFileNil

`func (o *DepotPushSetting) SetPushDenyFileNil(b bool)`

 SetPushDenyFileNil sets the value for PushDenyFile to be an explicit nil

### UnsetPushDenyFile
`func (o *DepotPushSetting) UnsetPushDenyFile()`

UnsetPushDenyFile ensures that no value is present for PushDenyFile, not even an explicit nil
### GetPushFileSize

`func (o *DepotPushSetting) GetPushFileSize() string`

GetPushFileSize returns the PushFileSize field if non-nil, zero value otherwise.

### GetPushFileSizeOk

`func (o *DepotPushSetting) GetPushFileSizeOk() (*string, bool)`

GetPushFileSizeOk returns a tuple with the PushFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushFileSize

`func (o *DepotPushSetting) SetPushFileSize(v string)`

SetPushFileSize sets PushFileSize field to given value.

### HasPushFileSize

`func (o *DepotPushSetting) HasPushFileSize() bool`

HasPushFileSize returns a boolean if a field has been set.

### SetPushFileSizeNil

`func (o *DepotPushSetting) SetPushFileSizeNil(b bool)`

 SetPushFileSizeNil sets the value for PushFileSize to be an explicit nil

### UnsetPushFileSize
`func (o *DepotPushSetting) UnsetPushFileSize()`

UnsetPushFileSize ensures that no value is present for PushFileSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


