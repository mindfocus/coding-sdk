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

// checks if the CreateAttachmentPrepareSignUrlRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateAttachmentPrepareSignUrlRequest{}

// CreateAttachmentPrepareSignUrlRequest struct for CreateAttachmentPrepareSignUrlRequest
type CreateAttachmentPrepareSignUrlRequest struct {
	// 文件名
	FileName string `json:"FileName"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _CreateAttachmentPrepareSignUrlRequest CreateAttachmentPrepareSignUrlRequest

// NewCreateAttachmentPrepareSignUrlRequest instantiates a new CreateAttachmentPrepareSignUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAttachmentPrepareSignUrlRequest(fileName string, projectName string) *CreateAttachmentPrepareSignUrlRequest {
	this := CreateAttachmentPrepareSignUrlRequest{}
	this.FileName = fileName
	this.ProjectName = projectName
	return &this
}

// NewCreateAttachmentPrepareSignUrlRequestWithDefaults instantiates a new CreateAttachmentPrepareSignUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAttachmentPrepareSignUrlRequestWithDefaults() *CreateAttachmentPrepareSignUrlRequest {
	this := CreateAttachmentPrepareSignUrlRequest{}
	return &this
}

// GetFileName returns the FileName field value
func (o *CreateAttachmentPrepareSignUrlRequest) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *CreateAttachmentPrepareSignUrlRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *CreateAttachmentPrepareSignUrlRequest) SetFileName(v string) {
	o.FileName = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateAttachmentPrepareSignUrlRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateAttachmentPrepareSignUrlRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateAttachmentPrepareSignUrlRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o CreateAttachmentPrepareSignUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAttachmentPrepareSignUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FileName"] = o.FileName
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *CreateAttachmentPrepareSignUrlRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"FileName",
		"ProjectName",
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

	varCreateAttachmentPrepareSignUrlRequest := _CreateAttachmentPrepareSignUrlRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateAttachmentPrepareSignUrlRequest)

	if err != nil {
		return err
	}

	*o = CreateAttachmentPrepareSignUrlRequest(varCreateAttachmentPrepareSignUrlRequest)

	return err
}

type NullableCreateAttachmentPrepareSignUrlRequest struct {
	value *CreateAttachmentPrepareSignUrlRequest
	isSet bool
}

func (v NullableCreateAttachmentPrepareSignUrlRequest) Get() *CreateAttachmentPrepareSignUrlRequest {
	return v.value
}

func (v *NullableCreateAttachmentPrepareSignUrlRequest) Set(val *CreateAttachmentPrepareSignUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAttachmentPrepareSignUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAttachmentPrepareSignUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAttachmentPrepareSignUrlRequest(val *CreateAttachmentPrepareSignUrlRequest) *NullableCreateAttachmentPrepareSignUrlRequest {
	return &NullableCreateAttachmentPrepareSignUrlRequest{value: val, isSet: true}
}

func (v NullableCreateAttachmentPrepareSignUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAttachmentPrepareSignUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

