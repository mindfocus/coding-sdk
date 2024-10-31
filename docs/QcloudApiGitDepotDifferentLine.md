# QcloudApiGitDepotDifferentLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int64** | 排序号 | [default to 0]
**LeftNo** | **int64** | 操作起始行号 | [default to 0]
**Prefix** | **string** | 操作方式:\&quot;+\&quot;表示新增,\&quot;-\&quot;表示删除,\&quot; \&quot;表示不变 | [default to ""]
**RightNo** | **int64** | 操作结束行号 | [default to 0]
**Text** | **string** | 文本 | [default to "mcecjence"]

## Methods

### NewQcloudApiGitDepotDifferentLine

`func NewQcloudApiGitDepotDifferentLine(index int64, leftNo int64, prefix string, rightNo int64, text string, ) *QcloudApiGitDepotDifferentLine`

NewQcloudApiGitDepotDifferentLine instantiates a new QcloudApiGitDepotDifferentLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQcloudApiGitDepotDifferentLineWithDefaults

`func NewQcloudApiGitDepotDifferentLineWithDefaults() *QcloudApiGitDepotDifferentLine`

NewQcloudApiGitDepotDifferentLineWithDefaults instantiates a new QcloudApiGitDepotDifferentLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *QcloudApiGitDepotDifferentLine) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *QcloudApiGitDepotDifferentLine) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *QcloudApiGitDepotDifferentLine) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetLeftNo

`func (o *QcloudApiGitDepotDifferentLine) GetLeftNo() int64`

GetLeftNo returns the LeftNo field if non-nil, zero value otherwise.

### GetLeftNoOk

`func (o *QcloudApiGitDepotDifferentLine) GetLeftNoOk() (*int64, bool)`

GetLeftNoOk returns a tuple with the LeftNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftNo

`func (o *QcloudApiGitDepotDifferentLine) SetLeftNo(v int64)`

SetLeftNo sets LeftNo field to given value.


### GetPrefix

`func (o *QcloudApiGitDepotDifferentLine) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *QcloudApiGitDepotDifferentLine) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *QcloudApiGitDepotDifferentLine) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.


### GetRightNo

`func (o *QcloudApiGitDepotDifferentLine) GetRightNo() int64`

GetRightNo returns the RightNo field if non-nil, zero value otherwise.

### GetRightNoOk

`func (o *QcloudApiGitDepotDifferentLine) GetRightNoOk() (*int64, bool)`

GetRightNoOk returns a tuple with the RightNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightNo

`func (o *QcloudApiGitDepotDifferentLine) SetRightNo(v int64)`

SetRightNo sets RightNo field to given value.


### GetText

`func (o *QcloudApiGitDepotDifferentLine) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *QcloudApiGitDepotDifferentLine) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *QcloudApiGitDepotDifferentLine) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


