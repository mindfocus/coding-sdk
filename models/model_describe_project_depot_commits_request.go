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

// checks if the DescribeProjectDepotCommitsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectDepotCommitsRequest{}

// DescribeProjectDepotCommitsRequest struct for DescribeProjectDepotCommitsRequest
type DescribeProjectDepotCommitsRequest struct {
	// 分支名称
	Branch string `json:"Branch"`
	// 仓库类型
	DepotType string `json:"DepotType"`
	// 仓库 Id
	Id int32 `json:"Id"`
	// 项目 Id
	ProjectId int32 `json:"ProjectId"`
}

type _DescribeProjectDepotCommitsRequest DescribeProjectDepotCommitsRequest

// NewDescribeProjectDepotCommitsRequest instantiates a new DescribeProjectDepotCommitsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectDepotCommitsRequest(branch string, depotType string, id int32, projectId int32) *DescribeProjectDepotCommitsRequest {
	this := DescribeProjectDepotCommitsRequest{}
	this.Branch = branch
	this.DepotType = depotType
	this.Id = id
	this.ProjectId = projectId
	return &this
}

// NewDescribeProjectDepotCommitsRequestWithDefaults instantiates a new DescribeProjectDepotCommitsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectDepotCommitsRequestWithDefaults() *DescribeProjectDepotCommitsRequest {
	this := DescribeProjectDepotCommitsRequest{}
	return &this
}

// GetBranch returns the Branch field value
func (o *DescribeProjectDepotCommitsRequest) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotCommitsRequest) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *DescribeProjectDepotCommitsRequest) SetBranch(v string) {
	o.Branch = v
}

// GetDepotType returns the DepotType field value
func (o *DescribeProjectDepotCommitsRequest) GetDepotType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotType
}

// GetDepotTypeOk returns a tuple with the DepotType field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotCommitsRequest) GetDepotTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotType, true
}

// SetDepotType sets field value
func (o *DescribeProjectDepotCommitsRequest) SetDepotType(v string) {
	o.DepotType = v
}

// GetId returns the Id field value
func (o *DescribeProjectDepotCommitsRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotCommitsRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DescribeProjectDepotCommitsRequest) SetId(v int32) {
	o.Id = v
}

// GetProjectId returns the ProjectId field value
func (o *DescribeProjectDepotCommitsRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectDepotCommitsRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeProjectDepotCommitsRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o DescribeProjectDepotCommitsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectDepotCommitsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Branch"] = o.Branch
	toSerialize["DepotType"] = o.DepotType
	toSerialize["Id"] = o.Id
	toSerialize["ProjectId"] = o.ProjectId
	return toSerialize, nil
}

func (o *DescribeProjectDepotCommitsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Branch",
		"DepotType",
		"Id",
		"ProjectId",
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

	varDescribeProjectDepotCommitsRequest := _DescribeProjectDepotCommitsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeProjectDepotCommitsRequest)

	if err != nil {
		return err
	}

	*o = DescribeProjectDepotCommitsRequest(varDescribeProjectDepotCommitsRequest)

	return err
}

type NullableDescribeProjectDepotCommitsRequest struct {
	value *DescribeProjectDepotCommitsRequest
	isSet bool
}

func (v NullableDescribeProjectDepotCommitsRequest) Get() *DescribeProjectDepotCommitsRequest {
	return v.value
}

func (v *NullableDescribeProjectDepotCommitsRequest) Set(val *DescribeProjectDepotCommitsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectDepotCommitsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectDepotCommitsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectDepotCommitsRequest(val *DescribeProjectDepotCommitsRequest) *NullableDescribeProjectDepotCommitsRequest {
	return &NullableDescribeProjectDepotCommitsRequest{value: val, isSet: true}
}

func (v NullableDescribeProjectDepotCommitsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectDepotCommitsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

