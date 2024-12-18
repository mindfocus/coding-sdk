/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the PolicyInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PolicyInfo{}

// PolicyInfo 权限组信息
type PolicyInfo struct {
	// 显示名称
	Alias *string `json:"Alias,omitempty"`
	// 版本号
	CurrentVersion *string `json:"CurrentVersion,omitempty"`
	// 版本记录 ID
	CurrentVersionId *int64 `json:"CurrentVersionId,omitempty"`
	// 描述
	Description utils.NullableString `json:"Description,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty"`
	// 权限组 ID
	PolicyId *int64 `json:"PolicyId,omitempty"`
	// 类型：IDENTITY | RESOURCE
	PolicyType *string `json:"PolicyType,omitempty"`
	// 作用范围：SYSTEM | CUSTOM
	Scope *string `json:"Scope,omitempty"`
}

// NewPolicyInfo instantiates a new PolicyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyInfo() *PolicyInfo {
	this := PolicyInfo{}
	var alias string = ""
	this.Alias = &alias
	var currentVersion string = ""
	this.CurrentVersion = &currentVersion
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var name string = ""
	this.Name = &name
	var policyType string = ""
	this.PolicyType = &policyType
	var scope string = ""
	this.Scope = &scope
	return &this
}

// NewPolicyInfoWithDefaults instantiates a new PolicyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyInfoWithDefaults() *PolicyInfo {
	this := PolicyInfo{}
	var alias string = ""
	this.Alias = &alias
	var currentVersion string = ""
	this.CurrentVersion = &currentVersion
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var name string = ""
	this.Name = &name
	var policyType string = ""
	this.PolicyType = &policyType
	var scope string = ""
	this.Scope = &scope
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PolicyInfo) GetAlias() string {
	if o == nil || utils.IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetAliasOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PolicyInfo) HasAlias() bool {
	if o != nil && !utils.IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PolicyInfo) SetAlias(v string) {
	o.Alias = &v
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *PolicyInfo) GetCurrentVersion() string {
	if o == nil || utils.IsNil(o.CurrentVersion) {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetCurrentVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CurrentVersion) {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *PolicyInfo) HasCurrentVersion() bool {
	if o != nil && !utils.IsNil(o.CurrentVersion) {
		return true
	}

	return false
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *PolicyInfo) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetCurrentVersionId returns the CurrentVersionId field value if set, zero value otherwise.
func (o *PolicyInfo) GetCurrentVersionId() int64 {
	if o == nil || utils.IsNil(o.CurrentVersionId) {
		var ret int64
		return ret
	}
	return *o.CurrentVersionId
}

// GetCurrentVersionIdOk returns a tuple with the CurrentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetCurrentVersionIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CurrentVersionId) {
		return nil, false
	}
	return o.CurrentVersionId, true
}

// HasCurrentVersionId returns a boolean if a field has been set.
func (o *PolicyInfo) HasCurrentVersionId() bool {
	if o != nil && !utils.IsNil(o.CurrentVersionId) {
		return true
	}

	return false
}

// SetCurrentVersionId gets a reference to the given int64 and assigns it to the CurrentVersionId field.
func (o *PolicyInfo) SetCurrentVersionId(v int64) {
	o.CurrentVersionId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyInfo) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyInfo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given utils.NullableString and assigns it to the Description field.
func (o *PolicyInfo) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PolicyInfo) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PolicyInfo) UnsetDescription() {
	o.Description.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyInfo) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyInfo) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyInfo) SetName(v string) {
	o.Name = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *PolicyInfo) GetPolicyId() int64 {
	if o == nil || utils.IsNil(o.PolicyId) {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetPolicyIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *PolicyInfo) HasPolicyId() bool {
	if o != nil && !utils.IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *PolicyInfo) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetPolicyType returns the PolicyType field value if set, zero value otherwise.
func (o *PolicyInfo) GetPolicyType() string {
	if o == nil || utils.IsNil(o.PolicyType) {
		var ret string
		return ret
	}
	return *o.PolicyType
}

// GetPolicyTypeOk returns a tuple with the PolicyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetPolicyTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PolicyType) {
		return nil, false
	}
	return o.PolicyType, true
}

// HasPolicyType returns a boolean if a field has been set.
func (o *PolicyInfo) HasPolicyType() bool {
	if o != nil && !utils.IsNil(o.PolicyType) {
		return true
	}

	return false
}

// SetPolicyType gets a reference to the given string and assigns it to the PolicyType field.
func (o *PolicyInfo) SetPolicyType(v string) {
	o.PolicyType = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *PolicyInfo) GetScope() string {
	if o == nil || utils.IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyInfo) GetScopeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *PolicyInfo) HasScope() bool {
	if o != nil && !utils.IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *PolicyInfo) SetScope(v string) {
	o.Scope = &v
}

func (o PolicyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Alias) {
		toSerialize["Alias"] = o.Alias
	}
	if !utils.IsNil(o.CurrentVersion) {
		toSerialize["CurrentVersion"] = o.CurrentVersion
	}
	if !utils.IsNil(o.CurrentVersionId) {
		toSerialize["CurrentVersionId"] = o.CurrentVersionId
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.PolicyId) {
		toSerialize["PolicyId"] = o.PolicyId
	}
	if !utils.IsNil(o.PolicyType) {
		toSerialize["PolicyType"] = o.PolicyType
	}
	if !utils.IsNil(o.Scope) {
		toSerialize["Scope"] = o.Scope
	}
	return toSerialize, nil
}

type NullablePolicyInfo struct {
	value *PolicyInfo
	isSet bool
}

func (v NullablePolicyInfo) Get() *PolicyInfo {
	return v.value
}

func (v *NullablePolicyInfo) Set(val *PolicyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyInfo(val *PolicyInfo) *NullablePolicyInfo {
	return &NullablePolicyInfo{value: val, isSet: true}
}

func (v NullablePolicyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


