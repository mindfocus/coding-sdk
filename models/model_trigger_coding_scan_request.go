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

// checks if the TriggerCodingScanRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TriggerCodingScanRequest{}

// TriggerCodingScanRequest struct for TriggerCodingScanRequest
type TriggerCodingScanRequest struct {
	// 增量扫描
	IncrScan *bool `json:"IncrScan,omitempty"`
	// 扫描任务 id
	ScanId int32 `json:"ScanId"`
}

type _TriggerCodingScanRequest TriggerCodingScanRequest

// NewTriggerCodingScanRequest instantiates a new TriggerCodingScanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerCodingScanRequest(scanId int32) *TriggerCodingScanRequest {
	this := TriggerCodingScanRequest{}
	this.ScanId = scanId
	return &this
}

// NewTriggerCodingScanRequestWithDefaults instantiates a new TriggerCodingScanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerCodingScanRequestWithDefaults() *TriggerCodingScanRequest {
	this := TriggerCodingScanRequest{}
	return &this
}

// GetIncrScan returns the IncrScan field value if set, zero value otherwise.
func (o *TriggerCodingScanRequest) GetIncrScan() bool {
	if o == nil || utils.IsNil(o.IncrScan) {
		var ret bool
		return ret
	}
	return *o.IncrScan
}

// GetIncrScanOk returns a tuple with the IncrScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerCodingScanRequest) GetIncrScanOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncrScan) {
		return nil, false
	}
	return o.IncrScan, true
}

// HasIncrScan returns a boolean if a field has been set.
func (o *TriggerCodingScanRequest) HasIncrScan() bool {
	if o != nil && !utils.IsNil(o.IncrScan) {
		return true
	}

	return false
}

// SetIncrScan gets a reference to the given bool and assigns it to the IncrScan field.
func (o *TriggerCodingScanRequest) SetIncrScan(v bool) {
	o.IncrScan = &v
}

// GetScanId returns the ScanId field value
func (o *TriggerCodingScanRequest) GetScanId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *TriggerCodingScanRequest) GetScanIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *TriggerCodingScanRequest) SetScanId(v int32) {
	o.ScanId = v
}

func (o TriggerCodingScanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerCodingScanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IncrScan) {
		toSerialize["IncrScan"] = o.IncrScan
	}
	toSerialize["ScanId"] = o.ScanId
	return toSerialize, nil
}

func (o *TriggerCodingScanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ScanId",
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

	varTriggerCodingScanRequest := _TriggerCodingScanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTriggerCodingScanRequest)

	if err != nil {
		return err
	}

	*o = TriggerCodingScanRequest(varTriggerCodingScanRequest)

	return err
}

type NullableTriggerCodingScanRequest struct {
	value *TriggerCodingScanRequest
	isSet bool
}

func (v NullableTriggerCodingScanRequest) Get() *TriggerCodingScanRequest {
	return v.value
}

func (v *NullableTriggerCodingScanRequest) Set(val *TriggerCodingScanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCodingScanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCodingScanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCodingScanRequest(val *TriggerCodingScanRequest) *NullableTriggerCodingScanRequest {
	return &NullableTriggerCodingScanRequest{value: val, isSet: true}
}

func (v NullableTriggerCodingScanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCodingScanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


