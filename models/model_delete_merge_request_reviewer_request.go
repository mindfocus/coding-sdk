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

// checks if the DeleteMergeRequestReviewerRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteMergeRequestReviewerRequest{}

// DeleteMergeRequestReviewerRequest struct for DeleteMergeRequestReviewerRequest
type DeleteMergeRequestReviewerRequest struct {
	// 仓库 ID
	DepotId int64 `json:"DepotId"`
	// 仓库路径，与仓库ID二选一
	DepotPath *string `json:"DepotPath,omitempty"`
	// iid
	MergeId int64 `json:"MergeId"`
	// 评审者的GK
	ReviewerGlobalKey string `json:"ReviewerGlobalKey"`
}

type _DeleteMergeRequestReviewerRequest DeleteMergeRequestReviewerRequest

// NewDeleteMergeRequestReviewerRequest instantiates a new DeleteMergeRequestReviewerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMergeRequestReviewerRequest(depotId int64, mergeId int64, reviewerGlobalKey string) *DeleteMergeRequestReviewerRequest {
	this := DeleteMergeRequestReviewerRequest{}
	this.DepotId = depotId
	this.MergeId = mergeId
	this.ReviewerGlobalKey = reviewerGlobalKey
	return &this
}

// NewDeleteMergeRequestReviewerRequestWithDefaults instantiates a new DeleteMergeRequestReviewerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMergeRequestReviewerRequestWithDefaults() *DeleteMergeRequestReviewerRequest {
	this := DeleteMergeRequestReviewerRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *DeleteMergeRequestReviewerRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DeleteMergeRequestReviewerRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DeleteMergeRequestReviewerRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DeleteMergeRequestReviewerRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMergeRequestReviewerRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DeleteMergeRequestReviewerRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DeleteMergeRequestReviewerRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetMergeId returns the MergeId field value
func (o *DeleteMergeRequestReviewerRequest) GetMergeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MergeId
}

// GetMergeIdOk returns a tuple with the MergeId field value
// and a boolean to check if the value has been set.
func (o *DeleteMergeRequestReviewerRequest) GetMergeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeId, true
}

// SetMergeId sets field value
func (o *DeleteMergeRequestReviewerRequest) SetMergeId(v int64) {
	o.MergeId = v
}

// GetReviewerGlobalKey returns the ReviewerGlobalKey field value
func (o *DeleteMergeRequestReviewerRequest) GetReviewerGlobalKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReviewerGlobalKey
}

// GetReviewerGlobalKeyOk returns a tuple with the ReviewerGlobalKey field value
// and a boolean to check if the value has been set.
func (o *DeleteMergeRequestReviewerRequest) GetReviewerGlobalKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReviewerGlobalKey, true
}

// SetReviewerGlobalKey sets field value
func (o *DeleteMergeRequestReviewerRequest) SetReviewerGlobalKey(v string) {
	o.ReviewerGlobalKey = v
}

func (o DeleteMergeRequestReviewerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMergeRequestReviewerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["MergeId"] = o.MergeId
	toSerialize["ReviewerGlobalKey"] = o.ReviewerGlobalKey
	return toSerialize, nil
}

func (o *DeleteMergeRequestReviewerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"MergeId",
		"ReviewerGlobalKey",
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

	varDeleteMergeRequestReviewerRequest := _DeleteMergeRequestReviewerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteMergeRequestReviewerRequest)

	if err != nil {
		return err
	}

	*o = DeleteMergeRequestReviewerRequest(varDeleteMergeRequestReviewerRequest)

	return err
}

type NullableDeleteMergeRequestReviewerRequest struct {
	value *DeleteMergeRequestReviewerRequest
	isSet bool
}

func (v NullableDeleteMergeRequestReviewerRequest) Get() *DeleteMergeRequestReviewerRequest {
	return v.value
}

func (v *NullableDeleteMergeRequestReviewerRequest) Set(val *DeleteMergeRequestReviewerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMergeRequestReviewerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMergeRequestReviewerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMergeRequestReviewerRequest(val *DeleteMergeRequestReviewerRequest) *NullableDeleteMergeRequestReviewerRequest {
	return &NullableDeleteMergeRequestReviewerRequest{value: val, isSet: true}
}

func (v NullableDeleteMergeRequestReviewerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMergeRequestReviewerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


