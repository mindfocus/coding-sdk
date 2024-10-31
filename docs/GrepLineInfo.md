# GrepLineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **NullableString** | 分支名称 | [optional] [default to ""]
**CommitId** | Pointer to **NullableString** | 提交Id | [optional] [default to ""]
**Contents** | Pointer to **[]string** | 文件内容 | [optional] 
**FileMaxLine** | Pointer to **NullableInt64** | 文件最大行数 | [optional] 
**LineNum** | Pointer to **NullableInt64** | 代码片段的行数 | [optional] 
**MaxNum** | Pointer to **NullableInt64** | 页面上展现的最大行数 | [optional] 
**MinNum** | Pointer to **NullableInt64** | 页面上展现的最小行数 | [optional] 
**Path** | Pointer to **NullableString** | 代码片段对应的文件路径 | [optional] [default to ""]
**Text** | Pointer to **NullableString** | 查询代码片段的文本 | [optional] [default to ""]

## Methods

### NewGrepLineInfo

`func NewGrepLineInfo() *GrepLineInfo`

NewGrepLineInfo instantiates a new GrepLineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrepLineInfoWithDefaults

`func NewGrepLineInfoWithDefaults() *GrepLineInfo`

NewGrepLineInfoWithDefaults instantiates a new GrepLineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *GrepLineInfo) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *GrepLineInfo) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *GrepLineInfo) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *GrepLineInfo) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### SetBranchNameNil

`func (o *GrepLineInfo) SetBranchNameNil(b bool)`

 SetBranchNameNil sets the value for BranchName to be an explicit nil

### UnsetBranchName
`func (o *GrepLineInfo) UnsetBranchName()`

UnsetBranchName ensures that no value is present for BranchName, not even an explicit nil
### GetCommitId

`func (o *GrepLineInfo) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *GrepLineInfo) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *GrepLineInfo) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *GrepLineInfo) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### SetCommitIdNil

`func (o *GrepLineInfo) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *GrepLineInfo) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetContents

`func (o *GrepLineInfo) GetContents() []string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *GrepLineInfo) GetContentsOk() (*[]string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *GrepLineInfo) SetContents(v []string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *GrepLineInfo) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *GrepLineInfo) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *GrepLineInfo) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetFileMaxLine

`func (o *GrepLineInfo) GetFileMaxLine() int64`

GetFileMaxLine returns the FileMaxLine field if non-nil, zero value otherwise.

### GetFileMaxLineOk

`func (o *GrepLineInfo) GetFileMaxLineOk() (*int64, bool)`

GetFileMaxLineOk returns a tuple with the FileMaxLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileMaxLine

`func (o *GrepLineInfo) SetFileMaxLine(v int64)`

SetFileMaxLine sets FileMaxLine field to given value.

### HasFileMaxLine

`func (o *GrepLineInfo) HasFileMaxLine() bool`

HasFileMaxLine returns a boolean if a field has been set.

### SetFileMaxLineNil

`func (o *GrepLineInfo) SetFileMaxLineNil(b bool)`

 SetFileMaxLineNil sets the value for FileMaxLine to be an explicit nil

### UnsetFileMaxLine
`func (o *GrepLineInfo) UnsetFileMaxLine()`

UnsetFileMaxLine ensures that no value is present for FileMaxLine, not even an explicit nil
### GetLineNum

`func (o *GrepLineInfo) GetLineNum() int64`

GetLineNum returns the LineNum field if non-nil, zero value otherwise.

### GetLineNumOk

`func (o *GrepLineInfo) GetLineNumOk() (*int64, bool)`

GetLineNumOk returns a tuple with the LineNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNum

`func (o *GrepLineInfo) SetLineNum(v int64)`

SetLineNum sets LineNum field to given value.

### HasLineNum

`func (o *GrepLineInfo) HasLineNum() bool`

HasLineNum returns a boolean if a field has been set.

### SetLineNumNil

`func (o *GrepLineInfo) SetLineNumNil(b bool)`

 SetLineNumNil sets the value for LineNum to be an explicit nil

### UnsetLineNum
`func (o *GrepLineInfo) UnsetLineNum()`

UnsetLineNum ensures that no value is present for LineNum, not even an explicit nil
### GetMaxNum

`func (o *GrepLineInfo) GetMaxNum() int64`

GetMaxNum returns the MaxNum field if non-nil, zero value otherwise.

### GetMaxNumOk

`func (o *GrepLineInfo) GetMaxNumOk() (*int64, bool)`

GetMaxNumOk returns a tuple with the MaxNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNum

`func (o *GrepLineInfo) SetMaxNum(v int64)`

SetMaxNum sets MaxNum field to given value.

### HasMaxNum

`func (o *GrepLineInfo) HasMaxNum() bool`

HasMaxNum returns a boolean if a field has been set.

### SetMaxNumNil

`func (o *GrepLineInfo) SetMaxNumNil(b bool)`

 SetMaxNumNil sets the value for MaxNum to be an explicit nil

### UnsetMaxNum
`func (o *GrepLineInfo) UnsetMaxNum()`

UnsetMaxNum ensures that no value is present for MaxNum, not even an explicit nil
### GetMinNum

`func (o *GrepLineInfo) GetMinNum() int64`

GetMinNum returns the MinNum field if non-nil, zero value otherwise.

### GetMinNumOk

`func (o *GrepLineInfo) GetMinNumOk() (*int64, bool)`

GetMinNumOk returns a tuple with the MinNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNum

`func (o *GrepLineInfo) SetMinNum(v int64)`

SetMinNum sets MinNum field to given value.

### HasMinNum

`func (o *GrepLineInfo) HasMinNum() bool`

HasMinNum returns a boolean if a field has been set.

### SetMinNumNil

`func (o *GrepLineInfo) SetMinNumNil(b bool)`

 SetMinNumNil sets the value for MinNum to be an explicit nil

### UnsetMinNum
`func (o *GrepLineInfo) UnsetMinNum()`

UnsetMinNum ensures that no value is present for MinNum, not even an explicit nil
### GetPath

`func (o *GrepLineInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GrepLineInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GrepLineInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GrepLineInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GrepLineInfo) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GrepLineInfo) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetText

`func (o *GrepLineInfo) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *GrepLineInfo) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *GrepLineInfo) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *GrepLineInfo) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *GrepLineInfo) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *GrepLineInfo) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


