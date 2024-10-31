/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Line type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Line{}

// Line git diff详情中每行的详细信息
type Line struct {
	// diff信息中的行数
	Index utils.NullableInt64 `json:"Index,omitempty"`
	// 修改前第几行
	LeftNo utils.NullableInt64 `json:"LeftNo,omitempty"`
	// 前缀 + -
	Prefix utils.NullableString `json:"Prefix,omitempty"`
	// 修改后第几行
	RightNo utils.NullableInt64 `json:"RightNo,omitempty"`
	// 文件每行的具体内容
	Text utils.NullableString `json:"Text,omitempty"`
}

// NewLine instantiates a new Line object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLine() *Line {
	this := Line{}
	var prefix string = ""
	this.Prefix = *utils.NewNullableString(&prefix)
	var text string = ""
	this.Text = *utils.NewNullableString(&text)
	return &this
}

// NewLineWithDefaults instantiates a new Line object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLineWithDefaults() *Line {
	this := Line{}
	var prefix string = ""
	this.Prefix = *utils.NewNullableString(&prefix)
	var text string = ""
	this.Text = *utils.NewNullableString(&text)
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Line) GetIndex() int64 {
	if o == nil || utils.IsNil(o.Index.Get()) {
		var ret int64
		return ret
	}
	return *o.Index.Get()
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Line) GetIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Index.Get(), o.Index.IsSet()
}

// HasIndex returns a boolean if a field has been set.
func (o *Line) HasIndex() bool {
	if o != nil && o.Index.IsSet() {
		return true
	}

	return false
}

// SetIndex gets a reference to the given utils.NullableInt64 and assigns it to the Index field.
func (o *Line) SetIndex(v int64) {
	o.Index.Set(&v)
}
// SetIndexNil sets the value for Index to be an explicit nil
func (o *Line) SetIndexNil() {
	o.Index.Set(nil)
}

// UnsetIndex ensures that no value is present for Index, not even an explicit nil
func (o *Line) UnsetIndex() {
	o.Index.Unset()
}

// GetLeftNo returns the LeftNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Line) GetLeftNo() int64 {
	if o == nil || utils.IsNil(o.LeftNo.Get()) {
		var ret int64
		return ret
	}
	return *o.LeftNo.Get()
}

// GetLeftNoOk returns a tuple with the LeftNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Line) GetLeftNoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeftNo.Get(), o.LeftNo.IsSet()
}

// HasLeftNo returns a boolean if a field has been set.
func (o *Line) HasLeftNo() bool {
	if o != nil && o.LeftNo.IsSet() {
		return true
	}

	return false
}

// SetLeftNo gets a reference to the given utils.NullableInt64 and assigns it to the LeftNo field.
func (o *Line) SetLeftNo(v int64) {
	o.LeftNo.Set(&v)
}
// SetLeftNoNil sets the value for LeftNo to be an explicit nil
func (o *Line) SetLeftNoNil() {
	o.LeftNo.Set(nil)
}

// UnsetLeftNo ensures that no value is present for LeftNo, not even an explicit nil
func (o *Line) UnsetLeftNo() {
	o.LeftNo.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Line) GetPrefix() string {
	if o == nil || utils.IsNil(o.Prefix.Get()) {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Line) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *Line) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given utils.NullableString and assigns it to the Prefix field.
func (o *Line) SetPrefix(v string) {
	o.Prefix.Set(&v)
}
// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *Line) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *Line) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetRightNo returns the RightNo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Line) GetRightNo() int64 {
	if o == nil || utils.IsNil(o.RightNo.Get()) {
		var ret int64
		return ret
	}
	return *o.RightNo.Get()
}

// GetRightNoOk returns a tuple with the RightNo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Line) GetRightNoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RightNo.Get(), o.RightNo.IsSet()
}

// HasRightNo returns a boolean if a field has been set.
func (o *Line) HasRightNo() bool {
	if o != nil && o.RightNo.IsSet() {
		return true
	}

	return false
}

// SetRightNo gets a reference to the given utils.NullableInt64 and assigns it to the RightNo field.
func (o *Line) SetRightNo(v int64) {
	o.RightNo.Set(&v)
}
// SetRightNoNil sets the value for RightNo to be an explicit nil
func (o *Line) SetRightNoNil() {
	o.RightNo.Set(nil)
}

// UnsetRightNo ensures that no value is present for RightNo, not even an explicit nil
func (o *Line) UnsetRightNo() {
	o.RightNo.Unset()
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Line) GetText() string {
	if o == nil || utils.IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Line) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *Line) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given utils.NullableString and assigns it to the Text field.
func (o *Line) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *Line) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *Line) UnsetText() {
	o.Text.Unset()
}

func (o Line) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Line) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Index.IsSet() {
		toSerialize["Index"] = o.Index.Get()
	}
	if o.LeftNo.IsSet() {
		toSerialize["LeftNo"] = o.LeftNo.Get()
	}
	if o.Prefix.IsSet() {
		toSerialize["Prefix"] = o.Prefix.Get()
	}
	if o.RightNo.IsSet() {
		toSerialize["RightNo"] = o.RightNo.Get()
	}
	if o.Text.IsSet() {
		toSerialize["Text"] = o.Text.Get()
	}
	return toSerialize, nil
}

type NullableLine struct {
	value *Line
	isSet bool
}

func (v NullableLine) Get() *Line {
	return v.value
}

func (v *NullableLine) Set(val *Line) {
	v.value = val
	v.isSet = true
}

func (v NullableLine) IsSet() bool {
	return v.isSet
}

func (v *NullableLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLine(val *Line) *NullableLine {
	return &NullableLine{value: val, isSet: true}
}

func (v NullableLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

