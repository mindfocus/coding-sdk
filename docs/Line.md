# Line

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **NullableInt64** | diff信息中的行数 | [optional] 
**LeftNo** | Pointer to **NullableInt64** | 修改前第几行 | [optional] 
**Prefix** | Pointer to **NullableString** | 前缀 + - | [optional] [default to ""]
**RightNo** | Pointer to **NullableInt64** | 修改后第几行 | [optional] 
**Text** | Pointer to **NullableString** | 文件每行的具体内容 | [optional] [default to ""]

## Methods

### NewLine

`func NewLine() *Line`

NewLine instantiates a new Line object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineWithDefaults

`func NewLineWithDefaults() *Line`

NewLineWithDefaults instantiates a new Line object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *Line) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Line) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Line) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *Line) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *Line) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *Line) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetLeftNo

`func (o *Line) GetLeftNo() int64`

GetLeftNo returns the LeftNo field if non-nil, zero value otherwise.

### GetLeftNoOk

`func (o *Line) GetLeftNoOk() (*int64, bool)`

GetLeftNoOk returns a tuple with the LeftNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftNo

`func (o *Line) SetLeftNo(v int64)`

SetLeftNo sets LeftNo field to given value.

### HasLeftNo

`func (o *Line) HasLeftNo() bool`

HasLeftNo returns a boolean if a field has been set.

### SetLeftNoNil

`func (o *Line) SetLeftNoNil(b bool)`

 SetLeftNoNil sets the value for LeftNo to be an explicit nil

### UnsetLeftNo
`func (o *Line) UnsetLeftNo()`

UnsetLeftNo ensures that no value is present for LeftNo, not even an explicit nil
### GetPrefix

`func (o *Line) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *Line) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *Line) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *Line) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *Line) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *Line) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetRightNo

`func (o *Line) GetRightNo() int64`

GetRightNo returns the RightNo field if non-nil, zero value otherwise.

### GetRightNoOk

`func (o *Line) GetRightNoOk() (*int64, bool)`

GetRightNoOk returns a tuple with the RightNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightNo

`func (o *Line) SetRightNo(v int64)`

SetRightNo sets RightNo field to given value.

### HasRightNo

`func (o *Line) HasRightNo() bool`

HasRightNo returns a boolean if a field has been set.

### SetRightNoNil

`func (o *Line) SetRightNoNil(b bool)`

 SetRightNoNil sets the value for RightNo to be an explicit nil

### UnsetRightNo
`func (o *Line) UnsetRightNo()`

UnsetRightNo ensures that no value is present for RightNo, not even an explicit nil
### GetText

`func (o *Line) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Line) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Line) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Line) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *Line) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *Line) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


