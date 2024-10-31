# DifferentLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int64** | 排序号，由小到大递增 | [optional] 
**LeftNo** | Pointer to **int64** | 操作起始行号 | [optional] 
**Prefix** | Pointer to **string** | 操作方式:”+”表示新增,”-“表示删除,” “表示不变 | [optional] [default to ""]
**RightNo** | Pointer to **int64** | 操作结束行号 | [optional] 
**Text** | Pointer to **string** | 文本 | [optional] [default to ""]

## Methods

### NewDifferentLine

`func NewDifferentLine() *DifferentLine`

NewDifferentLine instantiates a new DifferentLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDifferentLineWithDefaults

`func NewDifferentLineWithDefaults() *DifferentLine`

NewDifferentLineWithDefaults instantiates a new DifferentLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *DifferentLine) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DifferentLine) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DifferentLine) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DifferentLine) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLeftNo

`func (o *DifferentLine) GetLeftNo() int64`

GetLeftNo returns the LeftNo field if non-nil, zero value otherwise.

### GetLeftNoOk

`func (o *DifferentLine) GetLeftNoOk() (*int64, bool)`

GetLeftNoOk returns a tuple with the LeftNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftNo

`func (o *DifferentLine) SetLeftNo(v int64)`

SetLeftNo sets LeftNo field to given value.

### HasLeftNo

`func (o *DifferentLine) HasLeftNo() bool`

HasLeftNo returns a boolean if a field has been set.

### GetPrefix

`func (o *DifferentLine) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *DifferentLine) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *DifferentLine) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *DifferentLine) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRightNo

`func (o *DifferentLine) GetRightNo() int64`

GetRightNo returns the RightNo field if non-nil, zero value otherwise.

### GetRightNoOk

`func (o *DifferentLine) GetRightNoOk() (*int64, bool)`

GetRightNoOk returns a tuple with the RightNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightNo

`func (o *DifferentLine) SetRightNo(v int64)`

SetRightNo sets RightNo field to given value.

### HasRightNo

`func (o *DifferentLine) HasRightNo() bool`

HasRightNo returns a boolean if a field has been set.

### GetText

`func (o *DifferentLine) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *DifferentLine) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *DifferentLine) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *DifferentLine) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


