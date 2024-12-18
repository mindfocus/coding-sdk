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

// checks if the ModifyWikiOrderRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ModifyWikiOrderRequest{}

// ModifyWikiOrderRequest struct for ModifyWikiOrderRequest
type ModifyWikiOrderRequest struct {
	// 在第几层级之后
	After *int64 `json:"After,omitempty"`
	// 在第几层级之前
	Before *int64 `json:"Before,omitempty"`
	// 是否检查权限,默认false
	Forced *bool `json:"Forced,omitempty"`
	// wiki Iid
	Iid int64 `json:"Iid"`
	// 父级 Iid
	ParentIid int64 `json:"ParentIid"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _ModifyWikiOrderRequest ModifyWikiOrderRequest

// NewModifyWikiOrderRequest instantiates a new ModifyWikiOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWikiOrderRequest(iid int64, parentIid int64, projectName string) *ModifyWikiOrderRequest {
	this := ModifyWikiOrderRequest{}
	this.Iid = iid
	this.ParentIid = parentIid
	this.ProjectName = projectName
	return &this
}

// NewModifyWikiOrderRequestWithDefaults instantiates a new ModifyWikiOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWikiOrderRequestWithDefaults() *ModifyWikiOrderRequest {
	this := ModifyWikiOrderRequest{}
	return &this
}

// GetAfter returns the After field value if set, zero value otherwise.
func (o *ModifyWikiOrderRequest) GetAfter() int64 {
	if o == nil || utils.IsNil(o.After) {
		var ret int64
		return ret
	}
	return *o.After
}

// GetAfterOk returns a tuple with the After field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyWikiOrderRequest) GetAfterOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.After) {
		return nil, false
	}
	return o.After, true
}

// HasAfter returns a boolean if a field has been set.
func (o *ModifyWikiOrderRequest) HasAfter() bool {
	if o != nil && !utils.IsNil(o.After) {
		return true
	}

	return false
}

// SetAfter gets a reference to the given int64 and assigns it to the After field.
func (o *ModifyWikiOrderRequest) SetAfter(v int64) {
	o.After = &v
}

// GetBefore returns the Before field value if set, zero value otherwise.
func (o *ModifyWikiOrderRequest) GetBefore() int64 {
	if o == nil || utils.IsNil(o.Before) {
		var ret int64
		return ret
	}
	return *o.Before
}

// GetBeforeOk returns a tuple with the Before field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyWikiOrderRequest) GetBeforeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Before) {
		return nil, false
	}
	return o.Before, true
}

// HasBefore returns a boolean if a field has been set.
func (o *ModifyWikiOrderRequest) HasBefore() bool {
	if o != nil && !utils.IsNil(o.Before) {
		return true
	}

	return false
}

// SetBefore gets a reference to the given int64 and assigns it to the Before field.
func (o *ModifyWikiOrderRequest) SetBefore(v int64) {
	o.Before = &v
}

// GetForced returns the Forced field value if set, zero value otherwise.
func (o *ModifyWikiOrderRequest) GetForced() bool {
	if o == nil || utils.IsNil(o.Forced) {
		var ret bool
		return ret
	}
	return *o.Forced
}

// GetForcedOk returns a tuple with the Forced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModifyWikiOrderRequest) GetForcedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Forced) {
		return nil, false
	}
	return o.Forced, true
}

// HasForced returns a boolean if a field has been set.
func (o *ModifyWikiOrderRequest) HasForced() bool {
	if o != nil && !utils.IsNil(o.Forced) {
		return true
	}

	return false
}

// SetForced gets a reference to the given bool and assigns it to the Forced field.
func (o *ModifyWikiOrderRequest) SetForced(v bool) {
	o.Forced = &v
}

// GetIid returns the Iid field value
func (o *ModifyWikiOrderRequest) GetIid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Iid
}

// GetIidOk returns a tuple with the Iid field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiOrderRequest) GetIidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iid, true
}

// SetIid sets field value
func (o *ModifyWikiOrderRequest) SetIid(v int64) {
	o.Iid = v
}

// GetParentIid returns the ParentIid field value
func (o *ModifyWikiOrderRequest) GetParentIid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentIid
}

// GetParentIidOk returns a tuple with the ParentIid field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiOrderRequest) GetParentIidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentIid, true
}

// SetParentIid sets field value
func (o *ModifyWikiOrderRequest) SetParentIid(v int64) {
	o.ParentIid = v
}

// GetProjectName returns the ProjectName field value
func (o *ModifyWikiOrderRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *ModifyWikiOrderRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *ModifyWikiOrderRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o ModifyWikiOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModifyWikiOrderRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.After) {
		toSerialize["After"] = o.After
	}
	if !utils.IsNil(o.Before) {
		toSerialize["Before"] = o.Before
	}
	if !utils.IsNil(o.Forced) {
		toSerialize["Forced"] = o.Forced
	}
	toSerialize["Iid"] = o.Iid
	toSerialize["ParentIid"] = o.ParentIid
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *ModifyWikiOrderRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Iid",
		"ParentIid",
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

	varModifyWikiOrderRequest := _ModifyWikiOrderRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModifyWikiOrderRequest)

	if err != nil {
		return err
	}

	*o = ModifyWikiOrderRequest(varModifyWikiOrderRequest)

	return err
}

type NullableModifyWikiOrderRequest struct {
	value *ModifyWikiOrderRequest
	isSet bool
}

func (v NullableModifyWikiOrderRequest) Get() *ModifyWikiOrderRequest {
	return v.value
}

func (v *NullableModifyWikiOrderRequest) Set(val *ModifyWikiOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWikiOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWikiOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWikiOrderRequest(val *ModifyWikiOrderRequest) *NullableModifyWikiOrderRequest {
	return &NullableModifyWikiOrderRequest{value: val, isSet: true}
}

func (v NullableModifyWikiOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWikiOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


