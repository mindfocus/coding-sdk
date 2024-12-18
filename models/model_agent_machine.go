/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/mindfocus/coding-sdk/utils"
)

// checks if the AgentMachine type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentMachine{}

// AgentMachine AgentMachine 结构
type AgentMachine struct {
	// 堡垒机 id
	Id *int64 `json:"Id,omitempty"`
	// 堡垒机名称
	Name *string `json:"Name,omitempty"`
	// 堡垒机状态
	Status *string `json:"Status,omitempty"`
}

// NewAgentMachine instantiates a new AgentMachine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentMachine() *AgentMachine {
	this := AgentMachine{}
	var name string = ""
	this.Name = &name
	var status string = ""
	this.Status = &status
	return &this
}

// NewAgentMachineWithDefaults instantiates a new AgentMachine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentMachineWithDefaults() *AgentMachine {
	this := AgentMachine{}
	var name string = ""
	this.Name = &name
	var status string = ""
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentMachine) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentMachine) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentMachine) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *AgentMachine) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentMachine) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentMachine) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentMachine) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentMachine) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentMachine) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentMachine) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentMachine) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AgentMachine) SetStatus(v string) {
	o.Status = &v
}

func (o AgentMachine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentMachine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAgentMachine struct {
	value *AgentMachine
	isSet bool
}

func (v NullableAgentMachine) Get() *AgentMachine {
	return v.value
}

func (v *NullableAgentMachine) Set(val *AgentMachine) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentMachine) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentMachine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentMachine(val *AgentMachine) *NullableAgentMachine {
	return &NullableAgentMachine{value: val, isSet: true}
}

func (v NullableAgentMachine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentMachine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


