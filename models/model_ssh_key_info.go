/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the SshKeyInfo type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SshKeyInfo{}

// SshKeyInfo SSH 公钥详细信息
type SshKeyInfo struct {
	// 创建时间
	CreatedAt *int32 `json:"CreatedAt,omitempty"`
	// 过期时间
	ExpirationDate *string `json:"ExpirationDate,omitempty"`
	// 指纹信息
	FingerPrint *string `json:"FingerPrint,omitempty"`
	// 是否过期
	HasExpired *bool `json:"HasExpired,omitempty"`
	// 公钥Id
	Id *int32 `json:"Id,omitempty"`
	// 公钥所属者Id
	OwnerId *int32 `json:"OwnerId,omitempty"`
	// 公钥标题
	Title *string `json:"Title,omitempty"`
}

// NewSshKeyInfo instantiates a new SshKeyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSshKeyInfo() *SshKeyInfo {
	this := SshKeyInfo{}
	var expirationDate string = ""
	this.ExpirationDate = &expirationDate
	var fingerPrint string = ""
	this.FingerPrint = &fingerPrint
	var hasExpired bool = false
	this.HasExpired = &hasExpired
	var title string = ""
	this.Title = &title
	return &this
}

// NewSshKeyInfoWithDefaults instantiates a new SshKeyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSshKeyInfoWithDefaults() *SshKeyInfo {
	this := SshKeyInfo{}
	var expirationDate string = ""
	this.ExpirationDate = &expirationDate
	var fingerPrint string = ""
	this.FingerPrint = &fingerPrint
	var hasExpired bool = false
	this.HasExpired = &hasExpired
	var title string = ""
	this.Title = &title
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SshKeyInfo) GetCreatedAt() int32 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetCreatedAtOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SshKeyInfo) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *SshKeyInfo) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SshKeyInfo) GetExpirationDate() string {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetExpirationDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SshKeyInfo) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *SshKeyInfo) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetFingerPrint returns the FingerPrint field value if set, zero value otherwise.
func (o *SshKeyInfo) GetFingerPrint() string {
	if o == nil || utils.IsNil(o.FingerPrint) {
		var ret string
		return ret
	}
	return *o.FingerPrint
}

// GetFingerPrintOk returns a tuple with the FingerPrint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetFingerPrintOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FingerPrint) {
		return nil, false
	}
	return o.FingerPrint, true
}

// HasFingerPrint returns a boolean if a field has been set.
func (o *SshKeyInfo) HasFingerPrint() bool {
	if o != nil && !utils.IsNil(o.FingerPrint) {
		return true
	}

	return false
}

// SetFingerPrint gets a reference to the given string and assigns it to the FingerPrint field.
func (o *SshKeyInfo) SetFingerPrint(v string) {
	o.FingerPrint = &v
}

// GetHasExpired returns the HasExpired field value if set, zero value otherwise.
func (o *SshKeyInfo) GetHasExpired() bool {
	if o == nil || utils.IsNil(o.HasExpired) {
		var ret bool
		return ret
	}
	return *o.HasExpired
}

// GetHasExpiredOk returns a tuple with the HasExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetHasExpiredOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasExpired) {
		return nil, false
	}
	return o.HasExpired, true
}

// HasHasExpired returns a boolean if a field has been set.
func (o *SshKeyInfo) HasHasExpired() bool {
	if o != nil && !utils.IsNil(o.HasExpired) {
		return true
	}

	return false
}

// SetHasExpired gets a reference to the given bool and assigns it to the HasExpired field.
func (o *SshKeyInfo) SetHasExpired(v bool) {
	o.HasExpired = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SshKeyInfo) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SshKeyInfo) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SshKeyInfo) SetId(v int32) {
	o.Id = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *SshKeyInfo) GetOwnerId() int32 {
	if o == nil || utils.IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetOwnerIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *SshKeyInfo) HasOwnerId() bool {
	if o != nil && !utils.IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *SshKeyInfo) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SshKeyInfo) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SshKeyInfo) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SshKeyInfo) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SshKeyInfo) SetTitle(v string) {
	o.Title = &v
}

func (o SshKeyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SshKeyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["ExpirationDate"] = o.ExpirationDate
	}
	if !utils.IsNil(o.FingerPrint) {
		toSerialize["FingerPrint"] = o.FingerPrint
	}
	if !utils.IsNil(o.HasExpired) {
		toSerialize["HasExpired"] = o.HasExpired
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.OwnerId) {
		toSerialize["OwnerId"] = o.OwnerId
	}
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	return toSerialize, nil
}

type NullableSshKeyInfo struct {
	value *SshKeyInfo
	isSet bool
}

func (v NullableSshKeyInfo) Get() *SshKeyInfo {
	return v.value
}

func (v *NullableSshKeyInfo) Set(val *SshKeyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSshKeyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSshKeyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSshKeyInfo(val *SshKeyInfo) *NullableSshKeyInfo {
	return &NullableSshKeyInfo{value: val, isSet: true}
}

func (v NullableSshKeyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSshKeyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


