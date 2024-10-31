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

// checks if the CreateProjectMemberPrincipalRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateProjectMemberPrincipalRequest{}

// CreateProjectMemberPrincipalRequest struct for CreateProjectMemberPrincipalRequest
type CreateProjectMemberPrincipalRequest struct {
	// 成员主体信息
	Principals []PrincipalAdd `json:"Principals"`
	// 项目Id
	ProjectId int32 `json:"ProjectId"`
}

type _CreateProjectMemberPrincipalRequest CreateProjectMemberPrincipalRequest

// NewCreateProjectMemberPrincipalRequest instantiates a new CreateProjectMemberPrincipalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectMemberPrincipalRequest(principals []PrincipalAdd, projectId int32) *CreateProjectMemberPrincipalRequest {
	this := CreateProjectMemberPrincipalRequest{}
	this.Principals = principals
	this.ProjectId = projectId
	return &this
}

// NewCreateProjectMemberPrincipalRequestWithDefaults instantiates a new CreateProjectMemberPrincipalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectMemberPrincipalRequestWithDefaults() *CreateProjectMemberPrincipalRequest {
	this := CreateProjectMemberPrincipalRequest{}
	return &this
}

// GetPrincipals returns the Principals field value
func (o *CreateProjectMemberPrincipalRequest) GetPrincipals() []PrincipalAdd {
	if o == nil {
		var ret []PrincipalAdd
		return ret
	}

	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value
// and a boolean to check if the value has been set.
func (o *CreateProjectMemberPrincipalRequest) GetPrincipalsOk() ([]PrincipalAdd, bool) {
	if o == nil {
		return nil, false
	}
	return o.Principals, true
}

// SetPrincipals sets field value
func (o *CreateProjectMemberPrincipalRequest) SetPrincipals(v []PrincipalAdd) {
	o.Principals = v
}

// GetProjectId returns the ProjectId field value
func (o *CreateProjectMemberPrincipalRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateProjectMemberPrincipalRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateProjectMemberPrincipalRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o CreateProjectMemberPrincipalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectMemberPrincipalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Principals"] = o.Principals
	toSerialize["ProjectId"] = o.ProjectId
	return toSerialize, nil
}

func (o *CreateProjectMemberPrincipalRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Principals",
		"ProjectId",
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

	varCreateProjectMemberPrincipalRequest := _CreateProjectMemberPrincipalRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateProjectMemberPrincipalRequest)

	if err != nil {
		return err
	}

	*o = CreateProjectMemberPrincipalRequest(varCreateProjectMemberPrincipalRequest)

	return err
}

type NullableCreateProjectMemberPrincipalRequest struct {
	value *CreateProjectMemberPrincipalRequest
	isSet bool
}

func (v NullableCreateProjectMemberPrincipalRequest) Get() *CreateProjectMemberPrincipalRequest {
	return v.value
}

func (v *NullableCreateProjectMemberPrincipalRequest) Set(val *CreateProjectMemberPrincipalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectMemberPrincipalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectMemberPrincipalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectMemberPrincipalRequest(val *CreateProjectMemberPrincipalRequest) *NullableCreateProjectMemberPrincipalRequest {
	return &NullableCreateProjectMemberPrincipalRequest{value: val, isSet: true}
}

func (v NullableCreateProjectMemberPrincipalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectMemberPrincipalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


