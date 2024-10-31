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

// checks if the ModifyWikiTitleRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyWikiTitleRequest{}

// ModifyWikiTitleRequest struct for ModifyWikiTitleRequest
type ModifyWikiTitleRequest struct {
	// wiki 编号
	Iid int64 `json:"Iid"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// wiki 标题
	Title string `json:"Title"`
}

type _ModifyWikiTitleRequest ModifyWikiTitleRequest

// NewModifyWikiTitleRequest instantiates a new ModifyWikiTitleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWikiTitleRequest(iid int64, projectName string, title string) *ModifyWikiTitleRequest {
	this := ModifyWikiTitleRequest{}
	this.Iid = iid
	this.ProjectName = projectName
	this.Title = title
	return &this
}

// NewModifyWikiTitleRequestWithDefaults instantiates a new ModifyWikiTitleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWikiTitleRequestWithDefaults() *ModifyWikiTitleRequest {
	this := ModifyWikiTitleRequest{}
	return &this
}

// GetIid returns the Iid field value
func (o *ModifyWikiTitleRequest) GetIid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Iid
}

// GetIidOk returns a tuple with the Iid field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiTitleRequest) GetIidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iid, true
}

// SetIid sets field value
func (o *ModifyWikiTitleRequest) SetIid(v int64) {
	o.Iid = v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyWikiTitleRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiTitleRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyWikiTitleRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetTitle returns the Title field value
func (o *ModifyWikiTitleRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiTitleRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ModifyWikiTitleRequest) SetTitle(v string) {
	o.Title = v
}

func (o ModifyWikiTitleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWikiTitleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Iid"] = o.Iid
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["Title"] = o.Title
	return toSerialize, nil
}

func (o *ModifyWikiTitleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Iid",
		"ProjectName",
		"Title",
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

	varModifyWikiTitleRequest := _ModifyWikiTitleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyWikiTitleRequest)

	if err != nil {
		return err
	}

	*o = ModifyWikiTitleRequest(varModifyWikiTitleRequest)

	return err
}

type NullableModifyWikiTitleRequest struct {
	value *ModifyWikiTitleRequest
	isSet bool
}

func (v NullableModifyWikiTitleRequest) Get() *ModifyWikiTitleRequest {
	return v.value
}

func (v *NullableModifyWikiTitleRequest) Set(val *ModifyWikiTitleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWikiTitleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWikiTitleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWikiTitleRequest(val *ModifyWikiTitleRequest) *NullableModifyWikiTitleRequest {
	return &NullableModifyWikiTitleRequest{value: val, isSet: true}
}

func (v NullableModifyWikiTitleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWikiTitleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


