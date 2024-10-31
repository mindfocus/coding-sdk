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

// checks if the ModifyWikiByZipRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyWikiByZipRequest{}

// ModifyWikiByZipRequest struct for ModifyWikiByZipRequest
type ModifyWikiByZipRequest struct {
	// 验证文件的token
	AuthToken string `json:"AuthToken"`
	// 文件名
	FileName string `json:"FileName"`
	// wiki的Iid
	Iid int64 `json:"Iid"`
	// 上传文件的uuid名称 b5d0d8e0-3aca-11eb-8673-a9b6d94ca755.png
	Key string `json:"Key"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 获取token的时间
	Time int32 `json:"Time"`
}

type _ModifyWikiByZipRequest ModifyWikiByZipRequest

// NewModifyWikiByZipRequest instantiates a new ModifyWikiByZipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWikiByZipRequest(authToken string, fileName string, iid int64, key string, projectName string, time int32) *ModifyWikiByZipRequest {
	this := ModifyWikiByZipRequest{}
	this.AuthToken = authToken
	this.FileName = fileName
	this.Iid = iid
	this.Key = key
	this.ProjectName = projectName
	this.Time = time
	return &this
}

// NewModifyWikiByZipRequestWithDefaults instantiates a new ModifyWikiByZipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWikiByZipRequestWithDefaults() *ModifyWikiByZipRequest {
	this := ModifyWikiByZipRequest{}
	return &this
}

// GetAuthToken returns the AuthToken field value
func (o *ModifyWikiByZipRequest) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiByZipRequest) GetAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *ModifyWikiByZipRequest) SetAuthToken(v string) {
	o.AuthToken = v
}

// GetFileName returns the FileName field value
func (o *ModifyWikiByZipRequest) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiByZipRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *ModifyWikiByZipRequest) SetFileName(v string) {
	o.FileName = v
}

// GetIid returns the Iid field value
func (o *ModifyWikiByZipRequest) GetIid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Iid
}

// GetIidOk returns a tuple with the Iid field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiByZipRequest) GetIidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iid, true
}

// SetIid sets field value
func (o *ModifyWikiByZipRequest) SetIid(v int64) {
	o.Iid = v
}

// GetKey returns the Key field value
func (o *ModifyWikiByZipRequest) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiByZipRequest) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ModifyWikiByZipRequest) SetKey(v string) {
	o.Key = v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyWikiByZipRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiByZipRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyWikiByZipRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetTime returns the Time field value
func (o *ModifyWikiByZipRequest) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiByZipRequest) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *ModifyWikiByZipRequest) SetTime(v int32) {
	o.Time = v
}

func (o ModifyWikiByZipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWikiByZipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AuthToken"] = o.AuthToken
	toSerialize["FileName"] = o.FileName
	toSerialize["Iid"] = o.Iid
	toSerialize["Key"] = o.Key
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["Time"] = o.Time
	return toSerialize, nil
}

func (o *ModifyWikiByZipRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AuthToken",
		"FileName",
		"Iid",
		"Key",
		"ProjectName",
		"Time",
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

	varModifyWikiByZipRequest := _ModifyWikiByZipRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyWikiByZipRequest)

	if err != nil {
		return err
	}

	*o = ModifyWikiByZipRequest(varModifyWikiByZipRequest)

	return err
}

type NullableModifyWikiByZipRequest struct {
	value *ModifyWikiByZipRequest
	isSet bool
}

func (v NullableModifyWikiByZipRequest) Get() *ModifyWikiByZipRequest {
	return v.value
}

func (v *NullableModifyWikiByZipRequest) Set(val *ModifyWikiByZipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWikiByZipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWikiByZipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWikiByZipRequest(val *ModifyWikiByZipRequest) *NullableModifyWikiByZipRequest {
	return &NullableModifyWikiByZipRequest{value: val, isSet: true}
}

func (v NullableModifyWikiByZipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWikiByZipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


