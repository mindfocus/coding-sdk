/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the QcloudApiGitDepotDifferentLine type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QcloudApiGitDepotDifferentLine{}

// QcloudApiGitDepotDifferentLine struct for QcloudApiGitDepotDifferentLine
type QcloudApiGitDepotDifferentLine struct {
	// 排序号
	Index int64 `json:"Index"`
	// 操作起始行号
	LeftNo int64 `json:"LeftNo"`
	// 操作方式:\"+\"表示新增,\"-\"表示删除,\" \"表示不变
	Prefix string `json:"Prefix"`
	// 操作结束行号
	RightNo int64 `json:"RightNo"`
	// 文本
	Text string `json:"Text"`
}

type _QcloudApiGitDepotDifferentLine QcloudApiGitDepotDifferentLine

// NewQcloudApiGitDepotDifferentLine instantiates a new QcloudApiGitDepotDifferentLine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQcloudApiGitDepotDifferentLine(index int64, leftNo int64, prefix string, rightNo int64, text string) *QcloudApiGitDepotDifferentLine {
	this := QcloudApiGitDepotDifferentLine{}
	this.Index = index
	this.LeftNo = leftNo
	this.Prefix = prefix
	this.RightNo = rightNo
	this.Text = text
	return &this
}

// NewQcloudApiGitDepotDifferentLineWithDefaults instantiates a new QcloudApiGitDepotDifferentLine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQcloudApiGitDepotDifferentLineWithDefaults() *QcloudApiGitDepotDifferentLine {
	this := QcloudApiGitDepotDifferentLine{}
	var index int64 = 0
	this.Index = index
	var leftNo int64 = 0
	this.LeftNo = leftNo
	var prefix string = ""
	this.Prefix = prefix
	var rightNo int64 = 0
	this.RightNo = rightNo
	var text string = "mcecjence"
	this.Text = text
	return &this
}

// GetIndex returns the Index field value
func (o *QcloudApiGitDepotDifferentLine) GetIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *QcloudApiGitDepotDifferentLine) GetIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *QcloudApiGitDepotDifferentLine) SetIndex(v int64) {
	o.Index = v
}

// GetLeftNo returns the LeftNo field value
func (o *QcloudApiGitDepotDifferentLine) GetLeftNo() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LeftNo
}

// GetLeftNoOk returns a tuple with the LeftNo field value
// and a boolean to check if the value has been set.
func (o *QcloudApiGitDepotDifferentLine) GetLeftNoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LeftNo, true
}

// SetLeftNo sets field value
func (o *QcloudApiGitDepotDifferentLine) SetLeftNo(v int64) {
	o.LeftNo = v
}

// GetPrefix returns the Prefix field value
func (o *QcloudApiGitDepotDifferentLine) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *QcloudApiGitDepotDifferentLine) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *QcloudApiGitDepotDifferentLine) SetPrefix(v string) {
	o.Prefix = v
}

// GetRightNo returns the RightNo field value
func (o *QcloudApiGitDepotDifferentLine) GetRightNo() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RightNo
}

// GetRightNoOk returns a tuple with the RightNo field value
// and a boolean to check if the value has been set.
func (o *QcloudApiGitDepotDifferentLine) GetRightNoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RightNo, true
}

// SetRightNo sets field value
func (o *QcloudApiGitDepotDifferentLine) SetRightNo(v int64) {
	o.RightNo = v
}

// GetText returns the Text field value
func (o *QcloudApiGitDepotDifferentLine) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *QcloudApiGitDepotDifferentLine) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *QcloudApiGitDepotDifferentLine) SetText(v string) {
	o.Text = v
}

func (o QcloudApiGitDepotDifferentLine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QcloudApiGitDepotDifferentLine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Index"] = o.Index
	toSerialize["LeftNo"] = o.LeftNo
	toSerialize["Prefix"] = o.Prefix
	toSerialize["RightNo"] = o.RightNo
	toSerialize["Text"] = o.Text
	return toSerialize, nil
}

func (o *QcloudApiGitDepotDifferentLine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Index",
		"LeftNo",
		"Prefix",
		"RightNo",
		"Text",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQcloudApiGitDepotDifferentLine := _QcloudApiGitDepotDifferentLine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQcloudApiGitDepotDifferentLine)

	if err != nil {
		return err
	}

	*o = QcloudApiGitDepotDifferentLine(varQcloudApiGitDepotDifferentLine)

	return err
}

type NullableQcloudApiGitDepotDifferentLine struct {
	value *QcloudApiGitDepotDifferentLine
	isSet bool
}

func (v NullableQcloudApiGitDepotDifferentLine) Get() *QcloudApiGitDepotDifferentLine {
	return v.value
}

func (v *NullableQcloudApiGitDepotDifferentLine) Set(val *QcloudApiGitDepotDifferentLine) {
	v.value = val
	v.isSet = true
}

func (v NullableQcloudApiGitDepotDifferentLine) IsSet() bool {
	return v.isSet
}

func (v *NullableQcloudApiGitDepotDifferentLine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQcloudApiGitDepotDifferentLine(val *QcloudApiGitDepotDifferentLine) *NullableQcloudApiGitDepotDifferentLine {
	return &NullableQcloudApiGitDepotDifferentLine{value: val, isSet: true}
}

func (v NullableQcloudApiGitDepotDifferentLine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQcloudApiGitDepotDifferentLine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

