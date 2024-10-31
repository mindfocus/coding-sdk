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

// checks if the ApiFileFolder type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiFileFolder{}

// ApiFileFolder struct for ApiFileFolder
type ApiFileFolder struct {
	// 文件夹ID
	Id int64 `json:"Id"`
	// 文件夹名
	Name string `json:"Name"`
}

type _ApiFileFolder ApiFileFolder

// NewApiFileFolder instantiates a new ApiFileFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFileFolder(id int64, name string) *ApiFileFolder {
	this := ApiFileFolder{}
	this.Id = id
	this.Name = name
	return &this
}

// NewApiFileFolderWithDefaults instantiates a new ApiFileFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFileFolderWithDefaults() *ApiFileFolder {
	this := ApiFileFolder{}
	var id int64 = 0
	this.Id = id
	var name string = ""
	this.Name = name
	return &this
}

// GetId returns the Id field value
func (o *ApiFileFolder) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiFileFolder) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiFileFolder) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApiFileFolder) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiFileFolder) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiFileFolder) SetName(v string) {
	o.Name = v
}

func (o ApiFileFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFileFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	return toSerialize, nil
}

func (o *ApiFileFolder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Id",
		"Name",
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

	varApiFileFolder := _ApiFileFolder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiFileFolder)

	if err != nil {
		return err
	}

	*o = ApiFileFolder(varApiFileFolder)

	return err
}

type NullableApiFileFolder struct {
	value *ApiFileFolder
	isSet bool
}

func (v NullableApiFileFolder) Get() *ApiFileFolder {
	return v.value
}

func (v *NullableApiFileFolder) Set(val *ApiFileFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFileFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFileFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFileFolder(val *ApiFileFolder) *NullableApiFileFolder {
	return &NullableApiFileFolder{value: val, isSet: true}
}

func (v NullableApiFileFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFileFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

