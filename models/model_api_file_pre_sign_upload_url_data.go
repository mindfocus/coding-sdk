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

// checks if the ApiFilePreSignUploadUrlData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiFilePreSignUploadUrlData{}

// ApiFilePreSignUploadUrlData struct for ApiFilePreSignUploadUrlData
type ApiFilePreSignUploadUrlData struct {
	// token信息
	AuthToken string `json:"AuthToken"`
	// 请求头信息
	Headers string `json:"Headers"`
	// 文件存储的key信息
	StorageKey string `json:"StorageKey"`
	// 上传链接
	UploadLink string `json:"UploadLink"`
}

type _ApiFilePreSignUploadUrlData ApiFilePreSignUploadUrlData

// NewApiFilePreSignUploadUrlData instantiates a new ApiFilePreSignUploadUrlData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFilePreSignUploadUrlData(authToken string, headers string, storageKey string, uploadLink string) *ApiFilePreSignUploadUrlData {
	this := ApiFilePreSignUploadUrlData{}
	this.AuthToken = authToken
	this.Headers = headers
	this.StorageKey = storageKey
	this.UploadLink = uploadLink
	return &this
}

// NewApiFilePreSignUploadUrlDataWithDefaults instantiates a new ApiFilePreSignUploadUrlData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFilePreSignUploadUrlDataWithDefaults() *ApiFilePreSignUploadUrlData {
	this := ApiFilePreSignUploadUrlData{}
	var authToken string = ""
	this.AuthToken = authToken
	var headers string = ""
	this.Headers = headers
	var storageKey string = ""
	this.StorageKey = storageKey
	var uploadLink string = ""
	this.UploadLink = uploadLink
	return &this
}

// GetAuthToken returns the AuthToken field value
func (o *ApiFilePreSignUploadUrlData) GetAuthToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value
// and a boolean to check if the value has been set.
func (o *ApiFilePreSignUploadUrlData) GetAuthTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthToken, true
}

// SetAuthToken sets field value
func (o *ApiFilePreSignUploadUrlData) SetAuthToken(v string) {
	o.AuthToken = v
}

// GetHeaders returns the Headers field value
func (o *ApiFilePreSignUploadUrlData) GetHeaders() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value
// and a boolean to check if the value has been set.
func (o *ApiFilePreSignUploadUrlData) GetHeadersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headers, true
}

// SetHeaders sets field value
func (o *ApiFilePreSignUploadUrlData) SetHeaders(v string) {
	o.Headers = v
}

// GetStorageKey returns the StorageKey field value
func (o *ApiFilePreSignUploadUrlData) GetStorageKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorageKey
}

// GetStorageKeyOk returns a tuple with the StorageKey field value
// and a boolean to check if the value has been set.
func (o *ApiFilePreSignUploadUrlData) GetStorageKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageKey, true
}

// SetStorageKey sets field value
func (o *ApiFilePreSignUploadUrlData) SetStorageKey(v string) {
	o.StorageKey = v
}

// GetUploadLink returns the UploadLink field value
func (o *ApiFilePreSignUploadUrlData) GetUploadLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadLink
}

// GetUploadLinkOk returns a tuple with the UploadLink field value
// and a boolean to check if the value has been set.
func (o *ApiFilePreSignUploadUrlData) GetUploadLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadLink, true
}

// SetUploadLink sets field value
func (o *ApiFilePreSignUploadUrlData) SetUploadLink(v string) {
	o.UploadLink = v
}

func (o ApiFilePreSignUploadUrlData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFilePreSignUploadUrlData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AuthToken"] = o.AuthToken
	toSerialize["Headers"] = o.Headers
	toSerialize["StorageKey"] = o.StorageKey
	toSerialize["UploadLink"] = o.UploadLink
	return toSerialize, nil
}

func (o *ApiFilePreSignUploadUrlData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AuthToken",
		"Headers",
		"StorageKey",
		"UploadLink",
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

	varApiFilePreSignUploadUrlData := _ApiFilePreSignUploadUrlData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiFilePreSignUploadUrlData)

	if err != nil {
		return err
	}

	*o = ApiFilePreSignUploadUrlData(varApiFilePreSignUploadUrlData)

	return err
}

type NullableApiFilePreSignUploadUrlData struct {
	value *ApiFilePreSignUploadUrlData
	isSet bool
}

func (v NullableApiFilePreSignUploadUrlData) Get() *ApiFilePreSignUploadUrlData {
	return v.value
}

func (v *NullableApiFilePreSignUploadUrlData) Set(val *ApiFilePreSignUploadUrlData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFilePreSignUploadUrlData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFilePreSignUploadUrlData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFilePreSignUploadUrlData(val *ApiFilePreSignUploadUrlData) *NullableApiFilePreSignUploadUrlData {
	return &NullableApiFilePreSignUploadUrlData{value: val, isSet: true}
}

func (v NullableApiFilePreSignUploadUrlData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFilePreSignUploadUrlData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


