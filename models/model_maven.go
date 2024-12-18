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

// checks if the Maven type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Maven{}

// Maven 查询制品版本可下载文件参数
type Maven struct {
	// 附属构件
	Classifier string `json:"Classifier"`
	// 打包方式：pom;jar;war
	Packaging string `json:"Packaging"`
}

type _Maven Maven

// NewMaven instantiates a new Maven object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaven(classifier string, packaging string) *Maven {
	this := Maven{}
	this.Classifier = classifier
	this.Packaging = packaging
	return &this
}

// NewMavenWithDefaults instantiates a new Maven object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMavenWithDefaults() *Maven {
	this := Maven{}
	var classifier string = ""
	this.Classifier = classifier
	var packaging string = ""
	this.Packaging = packaging
	return &this
}

// GetClassifier returns the Classifier field value
func (o *Maven) GetClassifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classifier
}

// GetClassifierOk returns a tuple with the Classifier field value
// and a boolean to check if the value has been set.
func (o *Maven) GetClassifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classifier, true
}

// SetClassifier sets field value
func (o *Maven) SetClassifier(v string) {
	o.Classifier = v
}

// GetPackaging returns the Packaging field value
func (o *Maven) GetPackaging() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Packaging
}

// GetPackagingOk returns a tuple with the Packaging field value
// and a boolean to check if the value has been set.
func (o *Maven) GetPackagingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Packaging, true
}

// SetPackaging sets field value
func (o *Maven) SetPackaging(v string) {
	o.Packaging = v
}

func (o Maven) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Maven) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Classifier"] = o.Classifier
	toSerialize["Packaging"] = o.Packaging
	return toSerialize, nil
}

func (o *Maven) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Classifier",
		"Packaging",
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

	varMaven := _Maven{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMaven)

	if err != nil {
		return err
	}

	*o = Maven(varMaven)

	return err
}

type NullableMaven struct {
	value *Maven
	isSet bool
}

func (v NullableMaven) Get() *Maven {
	return v.value
}

func (v *NullableMaven) Set(val *Maven) {
	v.value = val
	v.isSet = true
}

func (v NullableMaven) IsSet() bool {
	return v.isSet
}

func (v *NullableMaven) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaven(val *Maven) *NullableMaven {
	return &NullableMaven{value: val, isSet: true}
}

func (v NullableMaven) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaven) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


