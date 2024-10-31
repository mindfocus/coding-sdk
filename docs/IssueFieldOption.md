# IssueFieldOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **NullableString** | 图标地址 | [optional] [default to ""]
**Sort** | Pointer to **int64** | 排序 | [optional] 
**Title** | Pointer to **string** | 标题 | [optional] [default to ""]
**Value** | Pointer to **string** | 选项值 | [optional] [default to ""]

## Methods

### NewIssueFieldOption

`func NewIssueFieldOption() *IssueFieldOption`

NewIssueFieldOption instantiates a new IssueFieldOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFieldOptionWithDefaults

`func NewIssueFieldOptionWithDefaults() *IssueFieldOption`

NewIssueFieldOptionWithDefaults instantiates a new IssueFieldOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *IssueFieldOption) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IssueFieldOption) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IssueFieldOption) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IssueFieldOption) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *IssueFieldOption) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *IssueFieldOption) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetSort

`func (o *IssueFieldOption) GetSort() int64`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *IssueFieldOption) GetSortOk() (*int64, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *IssueFieldOption) SetSort(v int64)`

SetSort sets Sort field to given value.

### HasSort

`func (o *IssueFieldOption) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetTitle

`func (o *IssueFieldOption) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssueFieldOption) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssueFieldOption) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssueFieldOption) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetValue

`func (o *IssueFieldOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IssueFieldOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IssueFieldOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *IssueFieldOption) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


