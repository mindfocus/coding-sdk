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

// checks if the DepartmentDepartmentData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DepartmentDepartmentData{}

// DepartmentDepartmentData 部门信息
type DepartmentDepartmentData struct {
	// 创建人ID
	CreatorId int64 `json:"CreatorId"`
	// 部门ID
	Id int64 `json:"Id"`
	// 部门名
	Name string `json:"Name"`
	// 父级部门ID
	ParentId int64 `json:"ParentId"`
	// 排序值
	Sort int64 `json:"Sort"`
	// 团队ID
	TeamId int64 `json:"TeamId"`
}

type _DepartmentDepartmentData DepartmentDepartmentData

// NewDepartmentDepartmentData instantiates a new DepartmentDepartmentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentDepartmentData(creatorId int64, id int64, name string, parentId int64, sort int64, teamId int64) *DepartmentDepartmentData {
	this := DepartmentDepartmentData{}
	this.CreatorId = creatorId
	this.Id = id
	this.Name = name
	this.ParentId = parentId
	this.Sort = sort
	this.TeamId = teamId
	return &this
}

// NewDepartmentDepartmentDataWithDefaults instantiates a new DepartmentDepartmentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentDepartmentDataWithDefaults() *DepartmentDepartmentData {
	this := DepartmentDepartmentData{}
	var creatorId int64 = 0
	this.CreatorId = creatorId
	var id int64 = 0
	this.Id = id
	var name string = ""
	this.Name = name
	var parentId int64 = 0
	this.ParentId = parentId
	var sort int64 = 0
	this.Sort = sort
	var teamId int64 = 0
	this.TeamId = teamId
	return &this
}

// GetCreatorId returns the CreatorId field value
func (o *DepartmentDepartmentData) GetCreatorId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentData) GetCreatorIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *DepartmentDepartmentData) SetCreatorId(v int64) {
	o.CreatorId = v
}

// GetId returns the Id field value
func (o *DepartmentDepartmentData) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentData) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DepartmentDepartmentData) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DepartmentDepartmentData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DepartmentDepartmentData) SetName(v string) {
	o.Name = v
}

// GetParentId returns the ParentId field value
func (o *DepartmentDepartmentData) GetParentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentData) GetParentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *DepartmentDepartmentData) SetParentId(v int64) {
	o.ParentId = v
}

// GetSort returns the Sort field value
func (o *DepartmentDepartmentData) GetSort() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentData) GetSortOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sort, true
}

// SetSort sets field value
func (o *DepartmentDepartmentData) SetSort(v int64) {
	o.Sort = v
}

// GetTeamId returns the TeamId field value
func (o *DepartmentDepartmentData) GetTeamId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *DepartmentDepartmentData) GetTeamIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *DepartmentDepartmentData) SetTeamId(v int64) {
	o.TeamId = v
}

func (o DepartmentDepartmentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentDepartmentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CreatorId"] = o.CreatorId
	toSerialize["Id"] = o.Id
	toSerialize["Name"] = o.Name
	toSerialize["ParentId"] = o.ParentId
	toSerialize["Sort"] = o.Sort
	toSerialize["TeamId"] = o.TeamId
	return toSerialize, nil
}

func (o *DepartmentDepartmentData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CreatorId",
		"Id",
		"Name",
		"ParentId",
		"Sort",
		"TeamId",
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

	varDepartmentDepartmentData := _DepartmentDepartmentData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDepartmentDepartmentData)

	if err != nil {
		return err
	}

	*o = DepartmentDepartmentData(varDepartmentDepartmentData)

	return err
}

type NullableDepartmentDepartmentData struct {
	value *DepartmentDepartmentData
	isSet bool
}

func (v NullableDepartmentDepartmentData) Get() *DepartmentDepartmentData {
	return v.value
}

func (v *NullableDepartmentDepartmentData) Set(val *DepartmentDepartmentData) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentDepartmentData) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentDepartmentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentDepartmentData(val *DepartmentDepartmentData) *NullableDepartmentDepartmentData {
	return &NullableDepartmentDepartmentData{value: val, isSet: true}
}

func (v NullableDepartmentDepartmentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentDepartmentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

