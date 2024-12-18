/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDepotByNameInfo200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotByNameInfo200ResponseResponse{}

// DescribeDepotByNameInfo200ResponseResponse 公共返回体
type DescribeDepotByNameInfo200ResponseResponse struct {
	Depot *DepotInfo `json:"Depot,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeDepotByNameInfo200ResponseResponse instantiates a new DescribeDepotByNameInfo200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotByNameInfo200ResponseResponse() *DescribeDepotByNameInfo200ResponseResponse {
	this := DescribeDepotByNameInfo200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeDepotByNameInfo200ResponseResponseWithDefaults instantiates a new DescribeDepotByNameInfo200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotByNameInfo200ResponseResponseWithDefaults() *DescribeDepotByNameInfo200ResponseResponse {
	this := DescribeDepotByNameInfo200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetDepot returns the Depot field value if set, zero value otherwise.
func (o *DescribeDepotByNameInfo200ResponseResponse) GetDepot() DepotInfo {
	if o == nil || utils.IsNil(o.Depot) {
		var ret DepotInfo
		return ret
	}
	return *o.Depot
}

// GetDepotOk returns a tuple with the Depot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotByNameInfo200ResponseResponse) GetDepotOk() (*DepotInfo, bool) {
	if o == nil || utils.IsNil(o.Depot) {
		return nil, false
	}
	return o.Depot, true
}

// HasDepot returns a boolean if a field has been set.
func (o *DescribeDepotByNameInfo200ResponseResponse) HasDepot() bool {
	if o != nil && !utils.IsNil(o.Depot) {
		return true
	}

	return false
}

// SetDepot gets a reference to the given DepotInfo and assigns it to the Depot field.
func (o *DescribeDepotByNameInfo200ResponseResponse) SetDepot(v DepotInfo) {
	o.Depot = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeDepotByNameInfo200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotByNameInfo200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeDepotByNameInfo200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeDepotByNameInfo200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeDepotByNameInfo200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotByNameInfo200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Depot) {
		toSerialize["Depot"] = o.Depot
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeDepotByNameInfo200ResponseResponse struct {
	value *DescribeDepotByNameInfo200ResponseResponse
	isSet bool
}

func (v NullableDescribeDepotByNameInfo200ResponseResponse) Get() *DescribeDepotByNameInfo200ResponseResponse {
	return v.value
}

func (v *NullableDescribeDepotByNameInfo200ResponseResponse) Set(val *DescribeDepotByNameInfo200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotByNameInfo200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotByNameInfo200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotByNameInfo200ResponseResponse(val *DescribeDepotByNameInfo200ResponseResponse) *NullableDescribeDepotByNameInfo200ResponseResponse {
	return &NullableDescribeDepotByNameInfo200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeDepotByNameInfo200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotByNameInfo200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


