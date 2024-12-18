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

// checks if the SetPredicatePolicyOnResourceRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SetPredicatePolicyOnResourceRequest{}

// SetPredicatePolicyOnResourceRequest struct for SetPredicatePolicyOnResourceRequest
type SetPredicatePolicyOnResourceRequest struct {
	Resource ResourceInfo `json:"Resource"`
	//   SELF_PARENT // 同时使用父级资源+当前资源   SELF_NONE  // 只使用当前消息   NONE_PARENT  // 只使用父级资源
	ResourcePredicatePolicy string `json:"ResourcePredicatePolicy"`
}

type _SetPredicatePolicyOnResourceRequest SetPredicatePolicyOnResourceRequest

// NewSetPredicatePolicyOnResourceRequest instantiates a new SetPredicatePolicyOnResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetPredicatePolicyOnResourceRequest(resource ResourceInfo, resourcePredicatePolicy string) *SetPredicatePolicyOnResourceRequest {
	this := SetPredicatePolicyOnResourceRequest{}
	this.Resource = resource
	this.ResourcePredicatePolicy = resourcePredicatePolicy
	return &this
}

// NewSetPredicatePolicyOnResourceRequestWithDefaults instantiates a new SetPredicatePolicyOnResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetPredicatePolicyOnResourceRequestWithDefaults() *SetPredicatePolicyOnResourceRequest {
	this := SetPredicatePolicyOnResourceRequest{}
	return &this
}

// GetResource returns the Resource field value
func (o *SetPredicatePolicyOnResourceRequest) GetResource() ResourceInfo {
	if o == nil {
		var ret ResourceInfo
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *SetPredicatePolicyOnResourceRequest) GetResourceOk() (*ResourceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *SetPredicatePolicyOnResourceRequest) SetResource(v ResourceInfo) {
	o.Resource = v
}

// GetResourcePredicatePolicy returns the ResourcePredicatePolicy field value
func (o *SetPredicatePolicyOnResourceRequest) GetResourcePredicatePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourcePredicatePolicy
}

// GetResourcePredicatePolicyOk returns a tuple with the ResourcePredicatePolicy field value
// and a boolean to check if the value has been set.
func (o *SetPredicatePolicyOnResourceRequest) GetResourcePredicatePolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePredicatePolicy, true
}

// SetResourcePredicatePolicy sets field value
func (o *SetPredicatePolicyOnResourceRequest) SetResourcePredicatePolicy(v string) {
	o.ResourcePredicatePolicy = v
}

func (o SetPredicatePolicyOnResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetPredicatePolicyOnResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Resource"] = o.Resource
	toSerialize["ResourcePredicatePolicy"] = o.ResourcePredicatePolicy
	return toSerialize, nil
}

func (o *SetPredicatePolicyOnResourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Resource",
		"ResourcePredicatePolicy",
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

	varSetPredicatePolicyOnResourceRequest := _SetPredicatePolicyOnResourceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetPredicatePolicyOnResourceRequest)

	if err != nil {
		return err
	}

	*o = SetPredicatePolicyOnResourceRequest(varSetPredicatePolicyOnResourceRequest)

	return err
}

type NullableSetPredicatePolicyOnResourceRequest struct {
	value *SetPredicatePolicyOnResourceRequest
	isSet bool
}

func (v NullableSetPredicatePolicyOnResourceRequest) Get() *SetPredicatePolicyOnResourceRequest {
	return v.value
}

func (v *NullableSetPredicatePolicyOnResourceRequest) Set(val *SetPredicatePolicyOnResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSetPredicatePolicyOnResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSetPredicatePolicyOnResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetPredicatePolicyOnResourceRequest(val *SetPredicatePolicyOnResourceRequest) *NullableSetPredicatePolicyOnResourceRequest {
	return &NullableSetPredicatePolicyOnResourceRequest{value: val, isSet: true}
}

func (v NullableSetPredicatePolicyOnResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetPredicatePolicyOnResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


