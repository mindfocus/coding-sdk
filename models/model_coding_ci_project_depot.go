/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CodingCIProjectDepot type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CodingCIProjectDepot{}

// CodingCIProjectDepot Depot 数据结构
type CodingCIProjectDepot struct {
	// 外部仓库是否被关联
	Authorized utils.NullableBool `json:"Authorized,omitempty"`
	// 仓库 Https 地址
	DepotHttpsUrl utils.NullableString `json:"DepotHttpsUrl,omitempty"`
	// 仓库 Ssh 地址
	DepotSshUrl utils.NullableString `json:"DepotSshUrl,omitempty"`
	// 仓库类型
	DepotType *string `json:"DepotType,omitempty"`
	// 仓库 Id
	Id *int32 `json:"Id,omitempty"`
	// 是否是默认显示第一位的仓库
	IsDefault *bool `json:"IsDefault,omitempty"`
	// 仓库名称
	Name *string `json:"Name,omitempty"`
	// 无用字段
	OpenModule *string `json:"OpenModule,omitempty"`
}

// NewCodingCIProjectDepot instantiates a new CodingCIProjectDepot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodingCIProjectDepot() *CodingCIProjectDepot {
	this := CodingCIProjectDepot{}
	var authorized bool = false
	this.Authorized = *utils.NewNullableBool(&authorized)
	var depotHttpsUrl string = ""
	this.DepotHttpsUrl = *utils.NewNullableString(&depotHttpsUrl)
	var depotSshUrl string = ""
	this.DepotSshUrl = *utils.NewNullableString(&depotSshUrl)
	var depotType string = ""
	this.DepotType = &depotType
	var isDefault bool = false
	this.IsDefault = &isDefault
	var name string = ""
	this.Name = &name
	var openModule string = ""
	this.OpenModule = &openModule
	return &this
}

// NewCodingCIProjectDepotWithDefaults instantiates a new CodingCIProjectDepot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodingCIProjectDepotWithDefaults() *CodingCIProjectDepot {
	this := CodingCIProjectDepot{}
	var authorized bool = false
	this.Authorized = *utils.NewNullableBool(&authorized)
	var depotHttpsUrl string = ""
	this.DepotHttpsUrl = *utils.NewNullableString(&depotHttpsUrl)
	var depotSshUrl string = ""
	this.DepotSshUrl = *utils.NewNullableString(&depotSshUrl)
	var depotType string = ""
	this.DepotType = &depotType
	var isDefault bool = false
	this.IsDefault = &isDefault
	var name string = ""
	this.Name = &name
	var openModule string = ""
	this.OpenModule = &openModule
	return &this
}

// GetAuthorized returns the Authorized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodingCIProjectDepot) GetAuthorized() bool {
	if o == nil || utils.IsNil(o.Authorized.Get()) {
		var ret bool
		return ret
	}
	return *o.Authorized.Get()
}

// GetAuthorizedOk returns a tuple with the Authorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodingCIProjectDepot) GetAuthorizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Authorized.Get(), o.Authorized.IsSet()
}

// HasAuthorized returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasAuthorized() bool {
	if o != nil && o.Authorized.IsSet() {
		return true
	}

	return false
}

// SetAuthorized gets a reference to the given utils.NullableBool and assigns it to the Authorized field.
func (o *CodingCIProjectDepot) SetAuthorized(v bool) {
	o.Authorized.Set(&v)
}
// SetAuthorizedNil sets the value for Authorized to be an explicit nil
func (o *CodingCIProjectDepot) SetAuthorizedNil() {
	o.Authorized.Set(nil)
}

// UnsetAuthorized ensures that no value is present for Authorized, not even an explicit nil
func (o *CodingCIProjectDepot) UnsetAuthorized() {
	o.Authorized.Unset()
}

// GetDepotHttpsUrl returns the DepotHttpsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodingCIProjectDepot) GetDepotHttpsUrl() string {
	if o == nil || utils.IsNil(o.DepotHttpsUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DepotHttpsUrl.Get()
}

// GetDepotHttpsUrlOk returns a tuple with the DepotHttpsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodingCIProjectDepot) GetDepotHttpsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepotHttpsUrl.Get(), o.DepotHttpsUrl.IsSet()
}

// HasDepotHttpsUrl returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasDepotHttpsUrl() bool {
	if o != nil && o.DepotHttpsUrl.IsSet() {
		return true
	}

	return false
}

// SetDepotHttpsUrl gets a reference to the given utils.NullableString and assigns it to the DepotHttpsUrl field.
func (o *CodingCIProjectDepot) SetDepotHttpsUrl(v string) {
	o.DepotHttpsUrl.Set(&v)
}
// SetDepotHttpsUrlNil sets the value for DepotHttpsUrl to be an explicit nil
func (o *CodingCIProjectDepot) SetDepotHttpsUrlNil() {
	o.DepotHttpsUrl.Set(nil)
}

// UnsetDepotHttpsUrl ensures that no value is present for DepotHttpsUrl, not even an explicit nil
func (o *CodingCIProjectDepot) UnsetDepotHttpsUrl() {
	o.DepotHttpsUrl.Unset()
}

// GetDepotSshUrl returns the DepotSshUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodingCIProjectDepot) GetDepotSshUrl() string {
	if o == nil || utils.IsNil(o.DepotSshUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DepotSshUrl.Get()
}

// GetDepotSshUrlOk returns a tuple with the DepotSshUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodingCIProjectDepot) GetDepotSshUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepotSshUrl.Get(), o.DepotSshUrl.IsSet()
}

// HasDepotSshUrl returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasDepotSshUrl() bool {
	if o != nil && o.DepotSshUrl.IsSet() {
		return true
	}

	return false
}

// SetDepotSshUrl gets a reference to the given utils.NullableString and assigns it to the DepotSshUrl field.
func (o *CodingCIProjectDepot) SetDepotSshUrl(v string) {
	o.DepotSshUrl.Set(&v)
}
// SetDepotSshUrlNil sets the value for DepotSshUrl to be an explicit nil
func (o *CodingCIProjectDepot) SetDepotSshUrlNil() {
	o.DepotSshUrl.Set(nil)
}

// UnsetDepotSshUrl ensures that no value is present for DepotSshUrl, not even an explicit nil
func (o *CodingCIProjectDepot) UnsetDepotSshUrl() {
	o.DepotSshUrl.Unset()
}

// GetDepotType returns the DepotType field value if set, zero value otherwise.
func (o *CodingCIProjectDepot) GetDepotType() string {
	if o == nil || utils.IsNil(o.DepotType) {
		var ret string
		return ret
	}
	return *o.DepotType
}

// GetDepotTypeOk returns a tuple with the DepotType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIProjectDepot) GetDepotTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotType) {
		return nil, false
	}
	return o.DepotType, true
}

// HasDepotType returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasDepotType() bool {
	if o != nil && !utils.IsNil(o.DepotType) {
		return true
	}

	return false
}

// SetDepotType gets a reference to the given string and assigns it to the DepotType field.
func (o *CodingCIProjectDepot) SetDepotType(v string) {
	o.DepotType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CodingCIProjectDepot) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIProjectDepot) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CodingCIProjectDepot) SetId(v int32) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *CodingCIProjectDepot) GetIsDefault() bool {
	if o == nil || utils.IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIProjectDepot) GetIsDefaultOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasIsDefault() bool {
	if o != nil && !utils.IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *CodingCIProjectDepot) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CodingCIProjectDepot) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIProjectDepot) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CodingCIProjectDepot) SetName(v string) {
	o.Name = &v
}

// GetOpenModule returns the OpenModule field value if set, zero value otherwise.
func (o *CodingCIProjectDepot) GetOpenModule() string {
	if o == nil || utils.IsNil(o.OpenModule) {
		var ret string
		return ret
	}
	return *o.OpenModule
}

// GetOpenModuleOk returns a tuple with the OpenModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CodingCIProjectDepot) GetOpenModuleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OpenModule) {
		return nil, false
	}
	return o.OpenModule, true
}

// HasOpenModule returns a boolean if a field has been set.
func (o *CodingCIProjectDepot) HasOpenModule() bool {
	if o != nil && !utils.IsNil(o.OpenModule) {
		return true
	}

	return false
}

// SetOpenModule gets a reference to the given string and assigns it to the OpenModule field.
func (o *CodingCIProjectDepot) SetOpenModule(v string) {
	o.OpenModule = &v
}

func (o CodingCIProjectDepot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodingCIProjectDepot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Authorized.IsSet() {
		toSerialize["Authorized"] = o.Authorized.Get()
	}
	if o.DepotHttpsUrl.IsSet() {
		toSerialize["DepotHttpsUrl"] = o.DepotHttpsUrl.Get()
	}
	if o.DepotSshUrl.IsSet() {
		toSerialize["DepotSshUrl"] = o.DepotSshUrl.Get()
	}
	if !utils.IsNil(o.DepotType) {
		toSerialize["DepotType"] = o.DepotType
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.IsDefault) {
		toSerialize["IsDefault"] = o.IsDefault
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.OpenModule) {
		toSerialize["OpenModule"] = o.OpenModule
	}
	return toSerialize, nil
}

type NullableCodingCIProjectDepot struct {
	value *CodingCIProjectDepot
	isSet bool
}

func (v NullableCodingCIProjectDepot) Get() *CodingCIProjectDepot {
	return v.value
}

func (v *NullableCodingCIProjectDepot) Set(val *CodingCIProjectDepot) {
	v.value = val
	v.isSet = true
}

func (v NullableCodingCIProjectDepot) IsSet() bool {
	return v.isSet
}

func (v *NullableCodingCIProjectDepot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodingCIProjectDepot(val *CodingCIProjectDepot) *NullableCodingCIProjectDepot {
	return &NullableCodingCIProjectDepot{value: val, isSet: true}
}

func (v NullableCodingCIProjectDepot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodingCIProjectDepot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

