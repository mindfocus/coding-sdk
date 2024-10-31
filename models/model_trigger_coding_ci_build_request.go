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

// checks if the TriggerCodingCIBuildRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TriggerCodingCIBuildRequest{}

// TriggerCodingCIBuildRequest struct for TriggerCodingCIBuildRequest
type TriggerCodingCIBuildRequest struct {
	// 构建计划 Id
	JobId int32 `json:"JobId"`
	// 构建附加的环境变量
	ParamList []CodingCIJobEnv `json:"ParamList,omitempty"`
	// 可重入字符串
	Reentrant *string `json:"Reentrant,omitempty"`
	// 分支名或 CommitId，当为构建计划的 DepotType 为 NONE 是可不传
	Revision *string `json:"Revision,omitempty"`
}

type _TriggerCodingCIBuildRequest TriggerCodingCIBuildRequest

// NewTriggerCodingCIBuildRequest instantiates a new TriggerCodingCIBuildRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerCodingCIBuildRequest(jobId int32) *TriggerCodingCIBuildRequest {
	this := TriggerCodingCIBuildRequest{}
	this.JobId = jobId
	return &this
}

// NewTriggerCodingCIBuildRequestWithDefaults instantiates a new TriggerCodingCIBuildRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerCodingCIBuildRequestWithDefaults() *TriggerCodingCIBuildRequest {
	this := TriggerCodingCIBuildRequest{}
	return &this
}

// GetJobId returns the JobId field value
func (o *TriggerCodingCIBuildRequest) GetJobId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *TriggerCodingCIBuildRequest) GetJobIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *TriggerCodingCIBuildRequest) SetJobId(v int32) {
	o.JobId = v
}

// GetParamList returns the ParamList field value if set, zero value otherwise.
func (o *TriggerCodingCIBuildRequest) GetParamList() []CodingCIJobEnv {
	if o == nil || utils.IsNil(o.ParamList) {
		var ret []CodingCIJobEnv
		return ret
	}
	return o.ParamList
}

// GetParamListOk returns a tuple with the ParamList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerCodingCIBuildRequest) GetParamListOk() ([]CodingCIJobEnv, bool) {
	if o == nil || utils.IsNil(o.ParamList) {
		return nil, false
	}
	return o.ParamList, true
}

// HasParamList returns a boolean if a field has been set.
func (o *TriggerCodingCIBuildRequest) HasParamList() bool {
	if o != nil && !utils.IsNil(o.ParamList) {
		return true
	}

	return false
}

// SetParamList gets a reference to the given []CodingCIJobEnv and assigns it to the ParamList field.
func (o *TriggerCodingCIBuildRequest) SetParamList(v []CodingCIJobEnv) {
	o.ParamList = v
}

// GetReentrant returns the Reentrant field value if set, zero value otherwise.
func (o *TriggerCodingCIBuildRequest) GetReentrant() string {
	if o == nil || utils.IsNil(o.Reentrant) {
		var ret string
		return ret
	}
	return *o.Reentrant
}

// GetReentrantOk returns a tuple with the Reentrant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerCodingCIBuildRequest) GetReentrantOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Reentrant) {
		return nil, false
	}
	return o.Reentrant, true
}

// HasReentrant returns a boolean if a field has been set.
func (o *TriggerCodingCIBuildRequest) HasReentrant() bool {
	if o != nil && !utils.IsNil(o.Reentrant) {
		return true
	}

	return false
}

// SetReentrant gets a reference to the given string and assigns it to the Reentrant field.
func (o *TriggerCodingCIBuildRequest) SetReentrant(v string) {
	o.Reentrant = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *TriggerCodingCIBuildRequest) GetRevision() string {
	if o == nil || utils.IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerCodingCIBuildRequest) GetRevisionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *TriggerCodingCIBuildRequest) HasRevision() bool {
	if o != nil && !utils.IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *TriggerCodingCIBuildRequest) SetRevision(v string) {
	o.Revision = &v
}

func (o TriggerCodingCIBuildRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TriggerCodingCIBuildRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["JobId"] = o.JobId
	if !utils.IsNil(o.ParamList) {
		toSerialize["ParamList"] = o.ParamList
	}
	if !utils.IsNil(o.Reentrant) {
		toSerialize["Reentrant"] = o.Reentrant
	}
	if !utils.IsNil(o.Revision) {
		toSerialize["Revision"] = o.Revision
	}
	return toSerialize, nil
}

func (o *TriggerCodingCIBuildRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"JobId",
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

	varTriggerCodingCIBuildRequest := _TriggerCodingCIBuildRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTriggerCodingCIBuildRequest)

	if err != nil {
		return err
	}

	*o = TriggerCodingCIBuildRequest(varTriggerCodingCIBuildRequest)

	return err
}

type NullableTriggerCodingCIBuildRequest struct {
	value *TriggerCodingCIBuildRequest
	isSet bool
}

func (v NullableTriggerCodingCIBuildRequest) Get() *TriggerCodingCIBuildRequest {
	return v.value
}

func (v *NullableTriggerCodingCIBuildRequest) Set(val *TriggerCodingCIBuildRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCodingCIBuildRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCodingCIBuildRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCodingCIBuildRequest(val *TriggerCodingCIBuildRequest) *NullableTriggerCodingCIBuildRequest {
	return &NullableTriggerCodingCIBuildRequest{value: val, isSet: true}
}

func (v NullableTriggerCodingCIBuildRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCodingCIBuildRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


