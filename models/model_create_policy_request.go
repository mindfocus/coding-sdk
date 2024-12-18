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

// checks if the CreatePolicyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreatePolicyRequest{}

// CreatePolicyRequest struct for CreatePolicyRequest
type CreatePolicyRequest struct {
	// 显示名称
	Alias string `json:"Alias"`
	// 描述
	Description *string `json:"Description,omitempty"`
	// 名称
	Name string `json:"Name"`
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
	// 权限组类型：IDENTITY | RESOURCE
	PolicyType string `json:"PolicyType"`
	// 适用的资源类型
	ResourceType []string `json:"ResourceType"`
}

type _CreatePolicyRequest CreatePolicyRequest

// NewCreatePolicyRequest instantiates a new CreatePolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePolicyRequest(alias string, name string, policyDocument PolicyDocument, policyType string, resourceType []string) *CreatePolicyRequest {
	this := CreatePolicyRequest{}
	this.Alias = alias
	this.Name = name
	this.PolicyDocument = policyDocument
	this.PolicyType = policyType
	this.ResourceType = resourceType
	return &this
}

// NewCreatePolicyRequestWithDefaults instantiates a new CreatePolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePolicyRequestWithDefaults() *CreatePolicyRequest {
	this := CreatePolicyRequest{}
	return &this
}

// GetAlias returns the Alias field value
func (o *CreatePolicyRequest) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *CreatePolicyRequest) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *CreatePolicyRequest) SetAlias(v string) {
	o.Alias = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePolicyRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePolicyRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePolicyRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePolicyRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *CreatePolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreatePolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreatePolicyRequest) SetName(v string) {
	o.Name = v
}

// GetPolicyDocument returns the PolicyDocument field value
func (o *CreatePolicyRequest) GetPolicyDocument() PolicyDocument {
	if o == nil {
		var ret PolicyDocument
		return ret
	}

	return o.PolicyDocument
}

// GetPolicyDocumentOk returns a tuple with the PolicyDocument field value
// and a boolean to check if the value has been set.
func (o *CreatePolicyRequest) GetPolicyDocumentOk() (*PolicyDocument, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyDocument, true
}

// SetPolicyDocument sets field value
func (o *CreatePolicyRequest) SetPolicyDocument(v PolicyDocument) {
	o.PolicyDocument = v
}

// GetPolicyType returns the PolicyType field value
func (o *CreatePolicyRequest) GetPolicyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value
// and a boolean to check if the value has been set.
func (o *CreatePolicyRequest) GetPolicyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyType, true
}

// SetPolicyType sets field value
func (o *CreatePolicyRequest) SetPolicyType(v string) {
	o.PolicyType = v
}

// GetResourceType returns the ResourceType field value
func (o *CreatePolicyRequest) GetResourceType() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CreatePolicyRequest) GetResourceTypeOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// SetResourceType sets field value
func (o *CreatePolicyRequest) SetResourceType(v []string) {
	o.ResourceType = v
}

func (o CreatePolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Alias"] = o.Alias
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["Name"] = o.Name
	toSerialize["PolicyDocument"] = o.PolicyDocument
	toSerialize["PolicyType"] = o.PolicyType
	toSerialize["ResourceType"] = o.ResourceType
	return toSerialize, nil
}

func (o *CreatePolicyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Alias",
		"Name",
		"PolicyDocument",
		"PolicyType",
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

	varCreatePolicyRequest := _CreatePolicyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatePolicyRequest)

	if err != nil {
		return err
	}

	*o = CreatePolicyRequest(varCreatePolicyRequest)

	return err
}

type NullableCreatePolicyRequest struct {
	value *CreatePolicyRequest
	isSet bool
}

func (v NullableCreatePolicyRequest) Get() *CreatePolicyRequest {
	return v.value
}

func (v *NullableCreatePolicyRequest) Set(val *CreatePolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePolicyRequest(val *CreatePolicyRequest) *NullableCreatePolicyRequest {
	return &NullableCreatePolicyRequest{value: val, isSet: true}
}

func (v NullableCreatePolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


