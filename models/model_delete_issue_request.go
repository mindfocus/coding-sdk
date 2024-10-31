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

// checks if the DeleteIssueRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteIssueRequest{}

// DeleteIssueRequest struct for DeleteIssueRequest
type DeleteIssueRequest struct {
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 事项类型，可不填
	IssueType *string `json:"IssueType,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _DeleteIssueRequest DeleteIssueRequest

// NewDeleteIssueRequest instantiates a new DeleteIssueRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteIssueRequest(issueCode int64, projectName string) *DeleteIssueRequest {
	this := DeleteIssueRequest{}
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewDeleteIssueRequestWithDefaults instantiates a new DeleteIssueRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteIssueRequestWithDefaults() *DeleteIssueRequest {
	this := DeleteIssueRequest{}
	return &this
}

// GetIssueCode returns the IssueCode field value
func (o *DeleteIssueRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *DeleteIssueRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *DeleteIssueRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetIssueType returns the IssueType field value if set, zero value otherwise.
func (o *DeleteIssueRequest) GetIssueType() string {
	if o == nil || utils.IsNil(o.IssueType) {
		var ret string
		return ret
	}
	return *o.IssueType
}

// GetIssueTypeOk returns a tuple with the IssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteIssueRequest) GetIssueTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IssueType) {
		return nil, false
	}
	return o.IssueType, true
}

// HasIssueType returns a boolean if a field has been set.
func (o *DeleteIssueRequest) HasIssueType() bool {
	if o != nil && !utils.IsNil(o.IssueType) {
		return true
	}

	return false
}

// SetIssueType gets a reference to the given string and assigns it to the IssueType field.
func (o *DeleteIssueRequest) SetIssueType(v string) {
	o.IssueType = &v
}

// GetProjectName returns the ProjectName field value
func (o *DeleteIssueRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DeleteIssueRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DeleteIssueRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o DeleteIssueRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteIssueRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IssueCode"] = o.IssueCode
	if !utils.IsNil(o.IssueType) {
		toSerialize["IssueType"] = o.IssueType
	}
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *DeleteIssueRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"IssueCode",
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

	varDeleteIssueRequest := _DeleteIssueRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteIssueRequest)

	if err != nil {
		return err
	}

	*o = DeleteIssueRequest(varDeleteIssueRequest)

	return err
}

type NullableDeleteIssueRequest struct {
	value *DeleteIssueRequest
	isSet bool
}

func (v NullableDeleteIssueRequest) Get() *DeleteIssueRequest {
	return v.value
}

func (v *NullableDeleteIssueRequest) Set(val *DeleteIssueRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteIssueRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteIssueRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteIssueRequest(val *DeleteIssueRequest) *NullableDeleteIssueRequest {
	return &NullableDeleteIssueRequest{value: val, isSet: true}
}

func (v NullableDeleteIssueRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteIssueRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


