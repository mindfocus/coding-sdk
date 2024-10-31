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

// checks if the DeleteProgramMemberPolicyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteProgramMemberPolicyRequest{}

// DeleteProgramMemberPolicyRequest struct for DeleteProgramMemberPolicyRequest
type DeleteProgramMemberPolicyRequest struct {
	// 权限策略，默认策略名或者策略 ID
	Policies []string `json:"Policies"`
	// 身份 ID
	PrincipalId string `json:"PrincipalId"`
	// 身份类型
	PrincipalType string `json:"PrincipalType"`
	// 项目集 ID
	ProgramId int64 `json:"ProgramId"`
}

type _DeleteProgramMemberPolicyRequest DeleteProgramMemberPolicyRequest

// NewDeleteProgramMemberPolicyRequest instantiates a new DeleteProgramMemberPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProgramMemberPolicyRequest(policies []string, principalId string, principalType string, programId int64) *DeleteProgramMemberPolicyRequest {
	this := DeleteProgramMemberPolicyRequest{}
	this.Policies = policies
	this.PrincipalId = principalId
	this.PrincipalType = principalType
	this.ProgramId = programId
	return &this
}

// NewDeleteProgramMemberPolicyRequestWithDefaults instantiates a new DeleteProgramMemberPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProgramMemberPolicyRequestWithDefaults() *DeleteProgramMemberPolicyRequest {
	this := DeleteProgramMemberPolicyRequest{}
	return &this
}

// GetPolicies returns the Policies field value
func (o *DeleteProgramMemberPolicyRequest) GetPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *DeleteProgramMemberPolicyRequest) GetPoliciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Policies, true
}

// SetPolicies sets field value
func (o *DeleteProgramMemberPolicyRequest) SetPolicies(v []string) {
	o.Policies = v
}

// GetPrincipalId returns the PrincipalId field value
func (o *DeleteProgramMemberPolicyRequest) GetPrincipalId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalId
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value
// and a boolean to check if the value has been set.
func (o *DeleteProgramMemberPolicyRequest) GetPrincipalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalId, true
}

// SetPrincipalId sets field value
func (o *DeleteProgramMemberPolicyRequest) SetPrincipalId(v string) {
	o.PrincipalId = v
}

// GetPrincipalType returns the PrincipalType field value
func (o *DeleteProgramMemberPolicyRequest) GetPrincipalType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalType
}

// GetPrincipalTypeOk returns a tuple with the PrincipalType field value
// and a boolean to check if the value has been set.
func (o *DeleteProgramMemberPolicyRequest) GetPrincipalTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalType, true
}

// SetPrincipalType sets field value
func (o *DeleteProgramMemberPolicyRequest) SetPrincipalType(v string) {
	o.PrincipalType = v
}

// GetProgramId returns the ProgramId field value
func (o *DeleteProgramMemberPolicyRequest) GetProgramId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value
// and a boolean to check if the value has been set.
func (o *DeleteProgramMemberPolicyRequest) GetProgramIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramId, true
}

// SetProgramId sets field value
func (o *DeleteProgramMemberPolicyRequest) SetProgramId(v int64) {
	o.ProgramId = v
}

func (o DeleteProgramMemberPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProgramMemberPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Policies"] = o.Policies
	toSerialize["PrincipalId"] = o.PrincipalId
	toSerialize["PrincipalType"] = o.PrincipalType
	toSerialize["ProgramId"] = o.ProgramId
	return toSerialize, nil
}

func (o *DeleteProgramMemberPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Policies",
		"PrincipalId",
		"PrincipalType",
		"ProgramId",
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

	varDeleteProgramMemberPolicyRequest := _DeleteProgramMemberPolicyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteProgramMemberPolicyRequest)

	if err != nil {
		return err
	}

	*o = DeleteProgramMemberPolicyRequest(varDeleteProgramMemberPolicyRequest)

	return err
}

type NullableDeleteProgramMemberPolicyRequest struct {
	value *DeleteProgramMemberPolicyRequest
	isSet bool
}

func (v NullableDeleteProgramMemberPolicyRequest) Get() *DeleteProgramMemberPolicyRequest {
	return v.value
}

func (v *NullableDeleteProgramMemberPolicyRequest) Set(val *DeleteProgramMemberPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProgramMemberPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProgramMemberPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProgramMemberPolicyRequest(val *DeleteProgramMemberPolicyRequest) *NullableDeleteProgramMemberPolicyRequest {
	return &NullableDeleteProgramMemberPolicyRequest{value: val, isSet: true}
}

func (v NullableDeleteProgramMemberPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProgramMemberPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

