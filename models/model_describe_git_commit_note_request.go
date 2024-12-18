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

// checks if the DescribeGitCommitNoteRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitCommitNoteRequest{}

// DescribeGitCommitNoteRequest struct for DescribeGitCommitNoteRequest
type DescribeGitCommitNoteRequest struct {
	// 提交的 Sha
	CommitSha string `json:"CommitSha"`
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 注释 分支 Ref
	NotesRef string `json:"NotesRef"`
}

type _DescribeGitCommitNoteRequest DescribeGitCommitNoteRequest

// NewDescribeGitCommitNoteRequest instantiates a new DescribeGitCommitNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitCommitNoteRequest(commitSha string, depotId int64, notesRef string) *DescribeGitCommitNoteRequest {
	this := DescribeGitCommitNoteRequest{}
	this.CommitSha = commitSha
	this.DepotId = depotId
	this.NotesRef = notesRef
	return &this
}

// NewDescribeGitCommitNoteRequestWithDefaults instantiates a new DescribeGitCommitNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitCommitNoteRequestWithDefaults() *DescribeGitCommitNoteRequest {
	this := DescribeGitCommitNoteRequest{}
	return &this
}

// GetCommitSha returns the CommitSha field value
func (o *DescribeGitCommitNoteRequest) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitNoteRequest) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value
func (o *DescribeGitCommitNoteRequest) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetDepotId returns the DepotId field value
func (o *DescribeGitCommitNoteRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitNoteRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeGitCommitNoteRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeGitCommitNoteRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitNoteRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeGitCommitNoteRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeGitCommitNoteRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetNotesRef returns the NotesRef field value
func (o *DescribeGitCommitNoteRequest) GetNotesRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesRef
}

// GetNotesRefOk returns a tuple with the NotesRef field value
// and a boolean to check if the value has been set.
func (o *DescribeGitCommitNoteRequest) GetNotesRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesRef, true
}

// SetNotesRef sets field value
func (o *DescribeGitCommitNoteRequest) SetNotesRef(v string) {
	o.NotesRef = v
}

func (o DescribeGitCommitNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitCommitNoteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CommitSha"] = o.CommitSha
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["NotesRef"] = o.NotesRef
	return toSerialize, nil
}

func (o *DescribeGitCommitNoteRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CommitSha",
		"DepotId",
		"NotesRef",
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

	varDescribeGitCommitNoteRequest := _DescribeGitCommitNoteRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitCommitNoteRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitCommitNoteRequest(varDescribeGitCommitNoteRequest)

	return err
}

type NullableDescribeGitCommitNoteRequest struct {
	value *DescribeGitCommitNoteRequest
	isSet bool
}

func (v NullableDescribeGitCommitNoteRequest) Get() *DescribeGitCommitNoteRequest {
	return v.value
}

func (v *NullableDescribeGitCommitNoteRequest) Set(val *DescribeGitCommitNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitCommitNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitCommitNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitCommitNoteRequest(val *DescribeGitCommitNoteRequest) *NullableDescribeGitCommitNoteRequest {
	return &NullableDescribeGitCommitNoteRequest{value: val, isSet: true}
}

func (v NullableDescribeGitCommitNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitCommitNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


