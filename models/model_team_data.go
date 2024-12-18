/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the TeamData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TeamData{}

// TeamData 团队和团队所有者信息
type TeamData struct {
	// 团队域名
	TeamHost *string `json:"TeamHost,omitempty"`
	// 团队名-拼音
	NamePinYin *string `json:"NamePinYin,omitempty"`
	// 团队ID
	Id *int64 `json:"Id,omitempty"`
	TeamOwner *UserData `json:"TeamOwner,omitempty"`
	// 团队图标
	Avatar *string `json:"Avatar,omitempty"`
	// 团队名
	Name *string `json:"Name,omitempty"`
}

// NewTeamData instantiates a new TeamData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeamData() *TeamData {
	this := TeamData{}
	var teamHost string = ""
	this.TeamHost = &teamHost
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	var id int64 = 0
	this.Id = &id
	var avatar string = ""
	this.Avatar = &avatar
	var name string = ""
	this.Name = &name
	return &this
}

// NewTeamDataWithDefaults instantiates a new TeamData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamDataWithDefaults() *TeamData {
	this := TeamData{}
	var teamHost string = ""
	this.TeamHost = &teamHost
	var namePinYin string = ""
	this.NamePinYin = &namePinYin
	var id int64 = 0
	this.Id = &id
	var avatar string = ""
	this.Avatar = &avatar
	var name string = ""
	this.Name = &name
	return &this
}

// GetTeamHost returns the TeamHost field value if set, zero value otherwise.
func (o *TeamData) GetTeamHost() string {
	if o == nil || utils.IsNil(o.TeamHost) {
		var ret string
		return ret
	}
	return *o.TeamHost
}

// GetTeamHostOk returns a tuple with the TeamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamData) GetTeamHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TeamHost) {
		return nil, false
	}
	return o.TeamHost, true
}

// HasTeamHost returns a boolean if a field has been set.
func (o *TeamData) HasTeamHost() bool {
	if o != nil && !utils.IsNil(o.TeamHost) {
		return true
	}

	return false
}

// SetTeamHost gets a reference to the given string and assigns it to the TeamHost field.
func (o *TeamData) SetTeamHost(v string) {
	o.TeamHost = &v
}

// GetNamePinYin returns the NamePinYin field value if set, zero value otherwise.
func (o *TeamData) GetNamePinYin() string {
	if o == nil || utils.IsNil(o.NamePinYin) {
		var ret string
		return ret
	}
	return *o.NamePinYin
}

// GetNamePinYinOk returns a tuple with the NamePinYin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamData) GetNamePinYinOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NamePinYin) {
		return nil, false
	}
	return o.NamePinYin, true
}

// HasNamePinYin returns a boolean if a field has been set.
func (o *TeamData) HasNamePinYin() bool {
	if o != nil && !utils.IsNil(o.NamePinYin) {
		return true
	}

	return false
}

// SetNamePinYin gets a reference to the given string and assigns it to the NamePinYin field.
func (o *TeamData) SetNamePinYin(v string) {
	o.NamePinYin = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TeamData) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamData) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TeamData) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *TeamData) SetId(v int64) {
	o.Id = &v
}

// GetTeamOwner returns the TeamOwner field value if set, zero value otherwise.
func (o *TeamData) GetTeamOwner() UserData {
	if o == nil || utils.IsNil(o.TeamOwner) {
		var ret UserData
		return ret
	}
	return *o.TeamOwner
}

// GetTeamOwnerOk returns a tuple with the TeamOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamData) GetTeamOwnerOk() (*UserData, bool) {
	if o == nil || utils.IsNil(o.TeamOwner) {
		return nil, false
	}
	return o.TeamOwner, true
}

// HasTeamOwner returns a boolean if a field has been set.
func (o *TeamData) HasTeamOwner() bool {
	if o != nil && !utils.IsNil(o.TeamOwner) {
		return true
	}

	return false
}

// SetTeamOwner gets a reference to the given UserData and assigns it to the TeamOwner field.
func (o *TeamData) SetTeamOwner(v UserData) {
	o.TeamOwner = &v
}

// GetAvatar returns the Avatar field value if set, zero value otherwise.
func (o *TeamData) GetAvatar() string {
	if o == nil || utils.IsNil(o.Avatar) {
		var ret string
		return ret
	}
	return *o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamData) GetAvatarOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Avatar) {
		return nil, false
	}
	return o.Avatar, true
}

// HasAvatar returns a boolean if a field has been set.
func (o *TeamData) HasAvatar() bool {
	if o != nil && !utils.IsNil(o.Avatar) {
		return true
	}

	return false
}

// SetAvatar gets a reference to the given string and assigns it to the Avatar field.
func (o *TeamData) SetAvatar(v string) {
	o.Avatar = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TeamData) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamData) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TeamData) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TeamData) SetName(v string) {
	o.Name = &v
}

func (o TeamData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TeamData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TeamHost) {
		toSerialize["TeamHost"] = o.TeamHost
	}
	if !utils.IsNil(o.NamePinYin) {
		toSerialize["NamePinYin"] = o.NamePinYin
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.TeamOwner) {
		toSerialize["TeamOwner"] = o.TeamOwner
	}
	if !utils.IsNil(o.Avatar) {
		toSerialize["Avatar"] = o.Avatar
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableTeamData struct {
	value *TeamData
	isSet bool
}

func (v NullableTeamData) Get() *TeamData {
	return v.value
}

func (v *NullableTeamData) Set(val *TeamData) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamData) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamData(val *TeamData) *NullableTeamData {
	return &NullableTeamData{value: val, isSet: true}
}

func (v NullableTeamData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


