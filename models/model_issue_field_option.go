/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the IssueFieldOption type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &IssueFieldOption{}

// IssueFieldOption 事项属性选项
type IssueFieldOption struct {
	// 图标地址
	Icon utils.NullableString `json:"Icon,omitempty"`
	// 排序
	Sort *int64 `json:"Sort,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty"`
	// 选项值
	Value *string `json:"Value,omitempty"`
}

// NewIssueFieldOption instantiates a new IssueFieldOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueFieldOption() *IssueFieldOption {
	this := IssueFieldOption{}
	var icon string = ""
	this.Icon = *utils.NewNullableString(&icon)
	var title string = ""
	this.Title = &title
	var value string = ""
	this.Value = &value
	return &this
}

// NewIssueFieldOptionWithDefaults instantiates a new IssueFieldOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueFieldOptionWithDefaults() *IssueFieldOption {
	this := IssueFieldOption{}
	var icon string = ""
	this.Icon = *utils.NewNullableString(&icon)
	var title string = ""
	this.Title = &title
	var value string = ""
	this.Value = &value
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IssueFieldOption) GetIcon() string {
	if o == nil || utils.IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IssueFieldOption) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *IssueFieldOption) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given utils.NullableString and assigns it to the Icon field.
func (o *IssueFieldOption) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *IssueFieldOption) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *IssueFieldOption) UnsetIcon() {
	o.Icon.Unset()
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *IssueFieldOption) GetSort() int64 {
	if o == nil || utils.IsNil(o.Sort) {
		var ret int64
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFieldOption) GetSortOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *IssueFieldOption) HasSort() bool {
	if o != nil && !utils.IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given int64 and assigns it to the Sort field.
func (o *IssueFieldOption) SetSort(v int64) {
	o.Sort = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IssueFieldOption) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFieldOption) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IssueFieldOption) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IssueFieldOption) SetTitle(v string) {
	o.Title = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IssueFieldOption) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueFieldOption) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IssueFieldOption) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *IssueFieldOption) SetValue(v string) {
	o.Value = &v
}

func (o IssueFieldOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueFieldOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["Icon"] = o.Icon.Get()
	}
	if !utils.IsNil(o.Sort) {
		toSerialize["Sort"] = o.Sort
	}
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !utils.IsNil(o.Value) {
		toSerialize["Value"] = o.Value
	}
	return toSerialize, nil
}

type NullableIssueFieldOption struct {
	value *IssueFieldOption
	isSet bool
}

func (v NullableIssueFieldOption) Get() *IssueFieldOption {
	return v.value
}

func (v *NullableIssueFieldOption) Set(val *IssueFieldOption) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueFieldOption) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueFieldOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueFieldOption(val *IssueFieldOption) *NullableIssueFieldOption {
	return &NullableIssueFieldOption{value: val, isSet: true}
}

func (v NullableIssueFieldOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueFieldOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


