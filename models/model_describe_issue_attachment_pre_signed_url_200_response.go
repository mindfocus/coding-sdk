/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeIssueAttachmentPreSignedUrl200Response type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueAttachmentPreSignedUrl200Response{}

// DescribeIssueAttachmentPreSignedUrl200Response struct for DescribeIssueAttachmentPreSignedUrl200Response
type DescribeIssueAttachmentPreSignedUrl200Response struct {
	Response *DescribeIssueAttachmentPreSignedUrl200ResponseResponse `json:"Response,omitempty"`
}

// NewDescribeIssueAttachmentPreSignedUrl200Response instantiates a new DescribeIssueAttachmentPreSignedUrl200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueAttachmentPreSignedUrl200Response() *DescribeIssueAttachmentPreSignedUrl200Response {
	this := DescribeIssueAttachmentPreSignedUrl200Response{}
	return &this
}

// NewDescribeIssueAttachmentPreSignedUrl200ResponseWithDefaults instantiates a new DescribeIssueAttachmentPreSignedUrl200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueAttachmentPreSignedUrl200ResponseWithDefaults() *DescribeIssueAttachmentPreSignedUrl200Response {
	this := DescribeIssueAttachmentPreSignedUrl200Response{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *DescribeIssueAttachmentPreSignedUrl200Response) GetResponse() DescribeIssueAttachmentPreSignedUrl200ResponseResponse {
	if o == nil || utils.IsNil(o.Response) {
		var ret DescribeIssueAttachmentPreSignedUrl200ResponseResponse
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeIssueAttachmentPreSignedUrl200Response) GetResponseOk() (*DescribeIssueAttachmentPreSignedUrl200ResponseResponse, bool) {
	if o == nil || utils.IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *DescribeIssueAttachmentPreSignedUrl200Response) HasResponse() bool {
	if o != nil && !utils.IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given DescribeIssueAttachmentPreSignedUrl200ResponseResponse and assigns it to the Response field.
func (o *DescribeIssueAttachmentPreSignedUrl200Response) SetResponse(v DescribeIssueAttachmentPreSignedUrl200ResponseResponse) {
	o.Response = &v
}

func (o DescribeIssueAttachmentPreSignedUrl200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueAttachmentPreSignedUrl200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Response) {
		toSerialize["Response"] = o.Response
	}
	return toSerialize, nil
}

type NullableDescribeIssueAttachmentPreSignedUrl200Response struct {
	value *DescribeIssueAttachmentPreSignedUrl200Response
	isSet bool
}

func (v NullableDescribeIssueAttachmentPreSignedUrl200Response) Get() *DescribeIssueAttachmentPreSignedUrl200Response {
	return v.value
}

func (v *NullableDescribeIssueAttachmentPreSignedUrl200Response) Set(val *DescribeIssueAttachmentPreSignedUrl200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueAttachmentPreSignedUrl200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueAttachmentPreSignedUrl200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueAttachmentPreSignedUrl200Response(val *DescribeIssueAttachmentPreSignedUrl200Response) *NullableDescribeIssueAttachmentPreSignedUrl200Response {
	return &NullableDescribeIssueAttachmentPreSignedUrl200Response{value: val, isSet: true}
}

func (v NullableDescribeIssueAttachmentPreSignedUrl200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueAttachmentPreSignedUrl200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


