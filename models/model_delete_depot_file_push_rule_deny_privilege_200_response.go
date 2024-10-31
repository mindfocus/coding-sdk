/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DeleteDepotFilePushRuleDenyPrivilege200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteDepotFilePushRuleDenyPrivilege200Response{}

// DeleteDepotFilePushRuleDenyPrivilege200Response struct for DeleteDepotFilePushRuleDenyPrivilege200Response
type DeleteDepotFilePushRuleDenyPrivilege200Response struct {
	Response *ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse `json:"Response,omitempty"`
}

// NewDeleteDepotFilePushRuleDenyPrivilege200Response instantiates a new DeleteDepotFilePushRuleDenyPrivilege200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDepotFilePushRuleDenyPrivilege200Response() *DeleteDepotFilePushRuleDenyPrivilege200Response {
	this := DeleteDepotFilePushRuleDenyPrivilege200Response{}
	return &this
}

// NewDeleteDepotFilePushRuleDenyPrivilege200ResponseWithDefaults instantiates a new DeleteDepotFilePushRuleDenyPrivilege200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDepotFilePushRuleDenyPrivilege200ResponseWithDefaults() *DeleteDepotFilePushRuleDenyPrivilege200Response {
	this := DeleteDepotFilePushRuleDenyPrivilege200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DeleteDepotFilePushRuleDenyPrivilege200Response) GetResponse() ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDepotFilePushRuleDenyPrivilege200Response) GetResponseOk() (*ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DeleteDepotFilePushRuleDenyPrivilege200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse and assigns it to the Response field.
func (o *DeleteDepotFilePushRuleDenyPrivilege200Response) SetResponse(v ModifyDepotFilePushRuleDenyPrivilege200ResponseResponse) {
	o.Response = &v
}

func (o DeleteDepotFilePushRuleDenyPrivilege200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDepotFilePushRuleDenyPrivilege200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDeleteDepotFilePushRuleDenyPrivilege200Response struct {
	value *DeleteDepotFilePushRuleDenyPrivilege200Response
	isSet bool
}

func (v NullableDeleteDepotFilePushRuleDenyPrivilege200Response) Get() *DeleteDepotFilePushRuleDenyPrivilege200Response {
	return v.value
}

func (v *NullableDeleteDepotFilePushRuleDenyPrivilege200Response) Set(val *DeleteDepotFilePushRuleDenyPrivilege200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDepotFilePushRuleDenyPrivilege200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDepotFilePushRuleDenyPrivilege200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDepotFilePushRuleDenyPrivilege200Response(val *DeleteDepotFilePushRuleDenyPrivilege200Response) *NullableDeleteDepotFilePushRuleDenyPrivilege200Response {
	return &NullableDeleteDepotFilePushRuleDenyPrivilege200Response{value: val, isSet: true}
}

func (v NullableDeleteDepotFilePushRuleDenyPrivilege200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDepotFilePushRuleDenyPrivilege200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


