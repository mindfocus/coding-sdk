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

// checks if the DeleteGitDeployKeyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteGitDeployKeyRequest{}

// DeleteGitDeployKeyRequest struct for DeleteGitDeployKeyRequest
type DeleteGitDeployKeyRequest struct {
	// SSH Key Id
	DeployKeyId int64 `json:"DeployKeyId"`
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径
	DepotPath *string `json:"DepotPath,omitempty"`
}

type _DeleteGitDeployKeyRequest DeleteGitDeployKeyRequest

// NewDeleteGitDeployKeyRequest instantiates a new DeleteGitDeployKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteGitDeployKeyRequest(deployKeyId int64, depotId int64) *DeleteGitDeployKeyRequest {
	this := DeleteGitDeployKeyRequest{}
	this.DeployKeyId = deployKeyId
	this.DepotId = depotId
	return &this
}

// NewDeleteGitDeployKeyRequestWithDefaults instantiates a new DeleteGitDeployKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteGitDeployKeyRequestWithDefaults() *DeleteGitDeployKeyRequest {
	this := DeleteGitDeployKeyRequest{}
	return &this
}

// GetDeployKeyId returns the DeployKeyId field value
func (o *DeleteGitDeployKeyRequest) GetDeployKeyId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeployKeyId
}

// GetDeployKeyIdOk returns a tuple with the DeployKeyId field value
// and a boolean to check if the value has been set.
func (o *DeleteGitDeployKeyRequest) GetDeployKeyIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeployKeyId, true
}

// SetDeployKeyId sets field value
func (o *DeleteGitDeployKeyRequest) SetDeployKeyId(v int64) {
	o.DeployKeyId = v
}

// GetDepotId returns the DepotId field value
func (o *DeleteGitDeployKeyRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DeleteGitDeployKeyRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DeleteGitDeployKeyRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DeleteGitDeployKeyRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteGitDeployKeyRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DeleteGitDeployKeyRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DeleteGitDeployKeyRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

func (o DeleteGitDeployKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteGitDeployKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DeployKeyId"] = o.DeployKeyId
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	return toSerialize, nil
}

func (o *DeleteGitDeployKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DeployKeyId",
		"DepotId",
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

	varDeleteGitDeployKeyRequest := _DeleteGitDeployKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteGitDeployKeyRequest)

	if err != nil {
		return err
	}

	*o = DeleteGitDeployKeyRequest(varDeleteGitDeployKeyRequest)

	return err
}

type NullableDeleteGitDeployKeyRequest struct {
	value *DeleteGitDeployKeyRequest
	isSet bool
}

func (v NullableDeleteGitDeployKeyRequest) Get() *DeleteGitDeployKeyRequest {
	return v.value
}

func (v *NullableDeleteGitDeployKeyRequest) Set(val *DeleteGitDeployKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteGitDeployKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteGitDeployKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteGitDeployKeyRequest(val *DeleteGitDeployKeyRequest) *NullableDeleteGitDeployKeyRequest {
	return &NullableDeleteGitDeployKeyRequest{value: val, isSet: true}
}

func (v NullableDeleteGitDeployKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteGitDeployKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


