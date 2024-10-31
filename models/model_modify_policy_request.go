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

// checks if the ModifyPolicyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyPolicyRequest{}

// ModifyPolicyRequest struct for ModifyPolicyRequest
type ModifyPolicyRequest struct {
	// 显示名称
	Alias *string `json:"Alias,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 权限组 ID
	Id int64 `json:"Id"`
	// 名称
	Name *string `json:"Name,omitempty"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
	// 适用的资源类型
	ResourceType []string `json:"ResourceType"`
}

type _ModifyPolicyRequest ModifyPolicyRequest

// NewModifyPolicyRequest instantiates a new ModifyPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyPolicyRequest(id int64, policyDocument PolicyDocument, resourceType []string) *ModifyPolicyRequest {
	this := ModifyPolicyRequest{}
	this.Id = id
	this.PolicyDocument = policyDocument
	this.ResourceType = resourceType
	return &this
}

// NewModifyPolicyRequestWithDefaults instantiates a new ModifyPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyPolicyRequestWithDefaults() *ModifyPolicyRequest {
	this := ModifyPolicyRequest{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModifyPolicyRequest) GetAlias() string {
	if o == nil || utils.IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyPolicyRequest) GetAliasOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModifyPolicyRequest) HasAlias() bool {
	if o != nil && !utils.IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModifyPolicyRequest) SetAlias(v string) {
	o.Alias = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModifyPolicyRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyPolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModifyPolicyRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModifyPolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *ModifyPolicyRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModifyPolicyRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModifyPolicyRequest) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModifyPolicyRequest) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModifyPolicyRequest) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModifyPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetPolicyDocument returns the PolicyDocument field value
func (o *ModifyPolicyRequest) GetPolicyDocument() PolicyDocument {
	if o == nil {
		var ret PolicyDocument
		return ret
	}

	return o.PolicyDocument
}

// GetPolicyDocumentOk returns a tuple with the PolicyDocument field value
// and a boolean to check if the value has been set.
func (o *ModifyPolicyRequest) GetPolicyDocumentOk() (*PolicyDocument, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyDocument, true
}

// SetPolicyDocument sets field value
func (o *ModifyPolicyRequest) SetPolicyDocument(v PolicyDocument) {
	o.PolicyDocument = v
}

// GetResourceType returns the ResourceType field value
func (o *ModifyPolicyRequest) GetResourceType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ModifyPolicyRequest) GetResourceTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *ModifyPolicyRequest) SetResourceType(v []string) {
	o.ResourceType = v
}

func (o ModifyPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Alias) {
		toSerialize["Alias"] = o.Alias
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["Id"] = o.Id
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	toSerialize["PolicyDocument"] = o.PolicyDocument
	toSerialize["ResourceType"] = o.ResourceType
	return toSerialize, nil
}

func (o *ModifyPolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Id",
		"PolicyDocument",
		"ResourceType",
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

	varModifyPolicyRequest := _ModifyPolicyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyPolicyRequest)

	if err != nil {
		return err
	}

	*o = ModifyPolicyRequest(varModifyPolicyRequest)

	return err
}

type NullableModifyPolicyRequest struct {
	value *ModifyPolicyRequest
	isSet bool
}

func (v NullableModifyPolicyRequest) Get() *ModifyPolicyRequest {
	return v.value
}

func (v *NullableModifyPolicyRequest) Set(val *ModifyPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyPolicyRequest(val *ModifyPolicyRequest) *NullableModifyPolicyRequest {
	return &NullableModifyPolicyRequest{value: val, isSet: true}
}

func (v NullableModifyPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


