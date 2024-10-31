# DescribePinyinRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Heteronym** | **bool** | 多音字 | [default to false]
**Style** | **string** | 风格 | 
**Value** | **string** | 汉字 | [default to ""]

## Methods

### NewDescribePinyinRequest

`func NewDescribePinyinRequest(heteronym bool, style string, value string, ) *DescribePinyinRequest`

NewDescribePinyinRequest instantiates a new DescribePinyinRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePinyinRequestWithDefaults

`func NewDescribePinyinRequestWithDefaults() *DescribePinyinRequest`

NewDescribePinyinRequestWithDefaults instantiates a new DescribePinyinRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeteronym

`func (o *DescribePinyinRequest) GetHeteronym() bool`

GetHeteronym returns the Heteronym field if non-nil, zero value otherwise.

### GetHeteronymOk

`func (o *DescribePinyinRequest) GetHeteronymOk() (*bool, bool)`

GetHeteronymOk returns a tuple with the Heteronym field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeteronym

`func (o *DescribePinyinRequest) SetHeteronym(v bool)`

SetHeteronym sets Heteronym field to given value.


### GetStyle

`func (o *DescribePinyinRequest) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *DescribePinyinRequest) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *DescribePinyinRequest) SetStyle(v string)`

SetStyle sets Style field to given value.


### GetValue

`func (o *DescribePinyinRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DescribePinyinRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DescribePinyinRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


