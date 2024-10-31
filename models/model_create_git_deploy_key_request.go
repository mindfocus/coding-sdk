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

// checks if the CreateGitDeployKeyRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateGitDeployKeyRequest{}

// CreateGitDeployKeyRequest struct for CreateGitDeployKeyRequest
type CreateGitDeployKeyRequest struct {
	// 是否授予写入权限
	AllowWrite bool `json:"AllowWrite"`
	// SSH key
	Content string `json:"Content"`
	// 仓库 Id
	DepotId int64 `json:"DepotId"`
	// 仓库路径
	DepotPath *string `json:"DepotPath,omitempty"`
	// 过期时间，不填则为永不过期
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
	// 部署公钥标题
	Title string `json:"Title"`
}

type _CreateGitDeployKeyRequest CreateGitDeployKeyRequest

// NewCreateGitDeployKeyRequest instantiates a new CreateGitDeployKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGitDeployKeyRequest(allowWrite bool, content string, depotId int64, title string) *CreateGitDeployKeyRequest {
	this := CreateGitDeployKeyRequest{}
	this.AllowWrite = allowWrite
	this.Content = content
	this.DepotId = depotId
	this.Title = title
	return &this
}

// NewCreateGitDeployKeyRequestWithDefaults instantiates a new CreateGitDeployKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGitDeployKeyRequestWithDefaults() *CreateGitDeployKeyRequest {
	this := CreateGitDeployKeyRequest{}
	return &this
}

// GetAllowWrite returns the AllowWrite field value
func (o *CreateGitDeployKeyRequest) GetAllowWrite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowWrite
}

// GetAllowWriteOk returns a tuple with the AllowWrite field value
// and a boolean to check if the value has been set.
func (o *CreateGitDeployKeyRequest) GetAllowWriteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowWrite, true
}

// SetAllowWrite sets field value
func (o *CreateGitDeployKeyRequest) SetAllowWrite(v bool) {
	o.AllowWrite = v
}

// GetContent returns the Content field value
func (o *CreateGitDeployKeyRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateGitDeployKeyRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateGitDeployKeyRequest) SetContent(v string) {
	o.Content = v
}

// GetDepotId returns the DepotId field value
func (o *CreateGitDeployKeyRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *CreateGitDeployKeyRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *CreateGitDeployKeyRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *CreateGitDeployKeyRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitDeployKeyRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *CreateGitDeployKeyRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *CreateGitDeployKeyRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CreateGitDeployKeyRequest) GetExpirationDate() string {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitDeployKeyRequest) GetExpirationDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CreateGitDeployKeyRequest) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *CreateGitDeployKeyRequest) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetTitle returns the Title field value
func (o *CreateGitDeployKeyRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateGitDeployKeyRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateGitDeployKeyRequest) SetTitle(v string) {
	o.Title = v
}

func (o CreateGitDeployKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGitDeployKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["AllowWrite"] = o.AllowWrite
	toSerialize["Content"] = o.Content
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	toSerialize["Title"] = o.Title
	return toSerialize, nil
}

func (o *CreateGitDeployKeyRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"AllowWrite",
		"Content",
		"DepotId",
		"Title",
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

	varCreateGitDeployKeyRequest := _CreateGitDeployKeyRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGitDeployKeyRequest)

	if err != nil {
		return err
	}

	*o = CreateGitDeployKeyRequest(varCreateGitDeployKeyRequest)

	return err
}

type NullableCreateGitDeployKeyRequest struct {
	value *CreateGitDeployKeyRequest
	isSet bool
}

func (v NullableCreateGitDeployKeyRequest) Get() *CreateGitDeployKeyRequest {
	return v.value
}

func (v *NullableCreateGitDeployKeyRequest) Set(val *CreateGitDeployKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGitDeployKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGitDeployKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGitDeployKeyRequest(val *CreateGitDeployKeyRequest) *NullableCreateGitDeployKeyRequest {
	return &NullableCreateGitDeployKeyRequest{value: val, isSet: true}
}

func (v NullableCreateGitDeployKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGitDeployKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


