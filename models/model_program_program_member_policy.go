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

// checks if the ProgramProgramMemberPolicy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProgramProgramMemberPolicy{}

// ProgramProgramMemberPolicy struct for ProgramProgramMemberPolicy
type ProgramProgramMemberPolicy struct {
	// 策略别名
	PolicyAlias string `json:"PolicyAlias"`
	// 策略 ID
	PolicyId string `json:"PolicyId"`
	// 策略名
	PolicyName string `json:"PolicyName"`
}

type _ProgramProgramMemberPolicy ProgramProgramMemberPolicy

// NewProgramProgramMemberPolicy instantiates a new ProgramProgramMemberPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramProgramMemberPolicy(policyAlias string, policyId string, policyName string) *ProgramProgramMemberPolicy {
	this := ProgramProgramMemberPolicy{}
	this.PolicyAlias = policyAlias
	this.PolicyId = policyId
	this.PolicyName = policyName
	return &this
}

// NewProgramProgramMemberPolicyWithDefaults instantiates a new ProgramProgramMemberPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramProgramMemberPolicyWithDefaults() *ProgramProgramMemberPolicy {
	this := ProgramProgramMemberPolicy{}
	return &this
}

// GetPolicyAlias returns the PolicyAlias field value
func (o *ProgramProgramMemberPolicy) GetPolicyAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyAlias
}

// GetPolicyAliasOk returns a tuple with the PolicyAlias field value
// and a boolean to check if the value has been set.
func (o *ProgramProgramMemberPolicy) GetPolicyAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyAlias, true
}

// SetPolicyAlias sets field value
func (o *ProgramProgramMemberPolicy) SetPolicyAlias(v string) {
	o.PolicyAlias = v
}

// GetPolicyId returns the PolicyId field value
func (o *ProgramProgramMemberPolicy) GetPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *ProgramProgramMemberPolicy) GetPolicyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *ProgramProgramMemberPolicy) SetPolicyId(v string) {
	o.PolicyId = v
}

// GetPolicyName returns the PolicyName field value
func (o *ProgramProgramMemberPolicy) GetPolicyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value
// and a boolean to check if the value has been set.
func (o *ProgramProgramMemberPolicy) GetPolicyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyName, true
}

// SetPolicyName sets field value
func (o *ProgramProgramMemberPolicy) SetPolicyName(v string) {
	o.PolicyName = v
}

func (o ProgramProgramMemberPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramProgramMemberPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PolicyAlias"] = o.PolicyAlias
	toSerialize["PolicyId"] = o.PolicyId
	toSerialize["PolicyName"] = o.PolicyName
	return toSerialize, nil
}

func (o *ProgramProgramMemberPolicy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PolicyAlias",
		"PolicyId",
		"PolicyName",
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

	varProgramProgramMemberPolicy := _ProgramProgramMemberPolicy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProgramProgramMemberPolicy)

	if err != nil {
		return err
	}

	*o = ProgramProgramMemberPolicy(varProgramProgramMemberPolicy)

	return err
}

type NullableProgramProgramMemberPolicy struct {
	value *ProgramProgramMemberPolicy
	isSet bool
}

func (v NullableProgramProgramMemberPolicy) Get() *ProgramProgramMemberPolicy {
	return v.value
}

func (v *NullableProgramProgramMemberPolicy) Set(val *ProgramProgramMemberPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramProgramMemberPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramProgramMemberPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramProgramMemberPolicy(val *ProgramProgramMemberPolicy) *NullableProgramProgramMemberPolicy {
	return &NullableProgramProgramMemberPolicy{value: val, isSet: true}
}

func (v NullableProgramProgramMemberPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramProgramMemberPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


