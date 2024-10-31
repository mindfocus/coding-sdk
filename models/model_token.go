/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Token type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Token{}

// Token 用于wiki 压缩包导入时的验证
type Token struct {
	// 验证文件的Token（用于导入wiki zip import 使用）
	AuthToken *string `json:"AuthToken,omitempty"`
	// cos存储对象
	Provider *string `json:"Provider,omitempty"`
	// cos上传的Id
	SecretId *string `json:"SecretId,omitempty"`
	// cos上传的key （用于导入wiki zip import 使用）
	SecretKey *string `json:"SecretKey,omitempty"`
	// 获取token 的时间（用于导入wiki zip import使用）
	Time *int32 `json:"Time,omitempty"`
	// 上传文件的Token
	UpToken *string `json:"UpToken,omitempty"`
	// 上传地址
	UploadLink *string `json:"UploadLink,omitempty"`
}

// NewToken instantiates a new Token object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewToken() *Token {
	this := Token{}
	var authToken string = ""
	this.AuthToken = &authToken
	var provider string = ""
	this.Provider = &provider
	var secretId string = ""
	this.SecretId = &secretId
	var secretKey string = ""
	this.SecretKey = &secretKey
	var upToken string = ""
	this.UpToken = &upToken
	var uploadLink string = ""
	this.UploadLink = &uploadLink
	return &this
}

// NewTokenWithDefaults instantiates a new Token object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenWithDefaults() *Token {
	this := Token{}
	var authToken string = ""
	this.AuthToken = &authToken
	var provider string = ""
	this.Provider = &provider
	var secretId string = ""
	this.SecretId = &secretId
	var secretKey string = ""
	this.SecretKey = &secretKey
	var upToken string = ""
	this.UpToken = &upToken
	var uploadLink string = ""
	this.UploadLink = &uploadLink
	return &this
}

// GetAuthToken returns the AuthToken field value if set, zero value otherwise.
func (o *Token) GetAuthToken() string {
	if o == nil || utils.IsNil(o.AuthToken) {
		var ret string
		return ret
	}
	return *o.AuthToken
}

// GetAuthTokenOk returns a tuple with the AuthToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetAuthTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AuthToken) {
		return nil, false
	}
	return o.AuthToken, true
}

// HasAuthToken returns a boolean if a field has been set.
func (o *Token) HasAuthToken() bool {
	if o != nil && !utils.IsNil(o.AuthToken) {
		return true
	}

	return false
}

// SetAuthToken gets a reference to the given string and assigns it to the AuthToken field.
func (o *Token) SetAuthToken(v string) {
	o.AuthToken = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *Token) GetProvider() string {
	if o == nil || utils.IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetProviderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *Token) HasProvider() bool {
	if o != nil && !utils.IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *Token) SetProvider(v string) {
	o.Provider = &v
}

// GetSecretId returns the SecretId field value if set, zero value otherwise.
func (o *Token) GetSecretId() string {
	if o == nil || utils.IsNil(o.SecretId) {
		var ret string
		return ret
	}
	return *o.SecretId
}

// GetSecretIdOk returns a tuple with the SecretId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetSecretIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SecretId) {
		return nil, false
	}
	return o.SecretId, true
}

// HasSecretId returns a boolean if a field has been set.
func (o *Token) HasSecretId() bool {
	if o != nil && !utils.IsNil(o.SecretId) {
		return true
	}

	return false
}

// SetSecretId gets a reference to the given string and assigns it to the SecretId field.
func (o *Token) SetSecretId(v string) {
	o.SecretId = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *Token) GetSecretKey() string {
	if o == nil || utils.IsNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetSecretKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SecretKey) {
		return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *Token) HasSecretKey() bool {
	if o != nil && !utils.IsNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *Token) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Token) GetTime() int32 {
	if o == nil || utils.IsNil(o.Time) {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Token) HasTime() bool {
	if o != nil && !utils.IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *Token) SetTime(v int32) {
	o.Time = &v
}

// GetUpToken returns the UpToken field value if set, zero value otherwise.
func (o *Token) GetUpToken() string {
	if o == nil || utils.IsNil(o.UpToken) {
		var ret string
		return ret
	}
	return *o.UpToken
}

// GetUpTokenOk returns a tuple with the UpToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetUpTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UpToken) {
		return nil, false
	}
	return o.UpToken, true
}

// HasUpToken returns a boolean if a field has been set.
func (o *Token) HasUpToken() bool {
	if o != nil && !utils.IsNil(o.UpToken) {
		return true
	}

	return false
}

// SetUpToken gets a reference to the given string and assigns it to the UpToken field.
func (o *Token) SetUpToken(v string) {
	o.UpToken = &v
}

// GetUploadLink returns the UploadLink field value if set, zero value otherwise.
func (o *Token) GetUploadLink() string {
	if o == nil || utils.IsNil(o.UploadLink) {
		var ret string
		return ret
	}
	return *o.UploadLink
}

// GetUploadLinkOk returns a tuple with the UploadLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Token) GetUploadLinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.UploadLink) {
		return nil, false
	}
	return o.UploadLink, true
}

// HasUploadLink returns a boolean if a field has been set.
func (o *Token) HasUploadLink() bool {
	if o != nil && !utils.IsNil(o.UploadLink) {
		return true
	}

	return false
}

// SetUploadLink gets a reference to the given string and assigns it to the UploadLink field.
func (o *Token) SetUploadLink(v string) {
	o.UploadLink = &v
}

func (o Token) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Token) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AuthToken) {
		toSerialize["AuthToken"] = o.AuthToken
	}
	if !utils.IsNil(o.Provider) {
		toSerialize["Provider"] = o.Provider
	}
	if !utils.IsNil(o.SecretId) {
		toSerialize["SecretId"] = o.SecretId
	}
	if !utils.IsNil(o.SecretKey) {
		toSerialize["SecretKey"] = o.SecretKey
	}
	if !utils.IsNil(o.Time) {
		toSerialize["Time"] = o.Time
	}
	if !utils.IsNil(o.UpToken) {
		toSerialize["UpToken"] = o.UpToken
	}
	if !utils.IsNil(o.UploadLink) {
		toSerialize["UploadLink"] = o.UploadLink
	}
	return toSerialize, nil
}

type NullableToken struct {
	value *Token
	isSet bool
}

func (v NullableToken) Get() *Token {
	return v.value
}

func (v *NullableToken) Set(val *Token) {
	v.value = val
	v.isSet = true
}

func (v NullableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableToken(val *Token) *NullableToken {
	return &NullableToken{value: val, isSet: true}
}

func (v NullableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


