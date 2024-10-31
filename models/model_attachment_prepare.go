/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the AttachmentPrepare type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AttachmentPrepare{}

// AttachmentPrepare 附件预上传信息
type AttachmentPrepare struct {
	// 附件 ID
	AttachmentId utils.NullableInt32 `json:"AttachmentId,omitempty"`
	// 附件上传地址
	PrepareSignUrl utils.NullableString `json:"PrepareSignUrl,omitempty"`
}

// NewAttachmentPrepare instantiates a new AttachmentPrepare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachmentPrepare() *AttachmentPrepare {
	this := AttachmentPrepare{}
	var prepareSignUrl string = ""
	this.PrepareSignUrl = *utils.NewNullableString(&prepareSignUrl)
	return &this
}

// NewAttachmentPrepareWithDefaults instantiates a new AttachmentPrepare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachmentPrepareWithDefaults() *AttachmentPrepare {
	this := AttachmentPrepare{}
	var prepareSignUrl string = ""
	this.PrepareSignUrl = *utils.NewNullableString(&prepareSignUrl)
	return &this
}

// GetAttachmentId returns the AttachmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentPrepare) GetAttachmentId() int32 {
	if o == nil || utils.IsNil(o.AttachmentId.Get()) {
		var ret int32
		return ret
	}
	return *o.AttachmentId.Get()
}

// GetAttachmentIdOk returns a tuple with the AttachmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentPrepare) GetAttachmentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttachmentId.Get(), o.AttachmentId.IsSet()
}

// HasAttachmentId returns a boolean if a field has been set.
func (o *AttachmentPrepare) HasAttachmentId() bool {
	if o != nil && o.AttachmentId.IsSet() {
		return true
	}

	return false
}

// SetAttachmentId gets a reference to the given utils.NullableInt32 and assigns it to the AttachmentId field.
func (o *AttachmentPrepare) SetAttachmentId(v int32) {
	o.AttachmentId.Set(&v)
}
// SetAttachmentIdNil sets the value for AttachmentId to be an explicit nil
func (o *AttachmentPrepare) SetAttachmentIdNil() {
	o.AttachmentId.Set(nil)
}

// UnsetAttachmentId ensures that no value is present for AttachmentId, not even an explicit nil
func (o *AttachmentPrepare) UnsetAttachmentId() {
	o.AttachmentId.Unset()
}

// GetPrepareSignUrl returns the PrepareSignUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AttachmentPrepare) GetPrepareSignUrl() string {
	if o == nil || utils.IsNil(o.PrepareSignUrl.Get()) {
		var ret string
		return ret
	}
	return *o.PrepareSignUrl.Get()
}

// GetPrepareSignUrlOk returns a tuple with the PrepareSignUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AttachmentPrepare) GetPrepareSignUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrepareSignUrl.Get(), o.PrepareSignUrl.IsSet()
}

// HasPrepareSignUrl returns a boolean if a field has been set.
func (o *AttachmentPrepare) HasPrepareSignUrl() bool {
	if o != nil && o.PrepareSignUrl.IsSet() {
		return true
	}

	return false
}

// SetPrepareSignUrl gets a reference to the given utils.NullableString and assigns it to the PrepareSignUrl field.
func (o *AttachmentPrepare) SetPrepareSignUrl(v string) {
	o.PrepareSignUrl.Set(&v)
}
// SetPrepareSignUrlNil sets the value for PrepareSignUrl to be an explicit nil
func (o *AttachmentPrepare) SetPrepareSignUrlNil() {
	o.PrepareSignUrl.Set(nil)
}

// UnsetPrepareSignUrl ensures that no value is present for PrepareSignUrl, not even an explicit nil
func (o *AttachmentPrepare) UnsetPrepareSignUrl() {
	o.PrepareSignUrl.Unset()
}

func (o AttachmentPrepare) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachmentPrepare) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachmentId.IsSet() {
		toSerialize["AttachmentId"] = o.AttachmentId.Get()
	}
	if o.PrepareSignUrl.IsSet() {
		toSerialize["PrepareSignUrl"] = o.PrepareSignUrl.Get()
	}
	return toSerialize, nil
}

type NullableAttachmentPrepare struct {
	value *AttachmentPrepare
	isSet bool
}

func (v NullableAttachmentPrepare) Get() *AttachmentPrepare {
	return v.value
}

func (v *NullableAttachmentPrepare) Set(val *AttachmentPrepare) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachmentPrepare) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachmentPrepare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachmentPrepare(val *AttachmentPrepare) *NullableAttachmentPrepare {
	return &NullableAttachmentPrepare{value: val, isSet: true}
}

func (v NullableAttachmentPrepare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachmentPrepare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


