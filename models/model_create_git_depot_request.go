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

// checks if the CreateGitDepotRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateGitDepotRequest{}

// CreateGitDepotRequest struct for CreateGitDepotRequest
type CreateGitDepotRequest struct {
	// 仓库名称
	DepotName string `json:"DepotName"`
	// 仓库的描述信息
	Description string `json:"Description,omitempty"`
	// 项目id
	ProjectId int64 `json:"ProjectId"`
	// 仓库是否允许公开访问
	Shared bool `json:"Shared,omitempty"`
}

type _CreateGitDepotRequest CreateGitDepotRequest

// NewCreateGitDepotRequest instantiates a new CreateGitDepotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGitDepotRequest(depotName string, projectId int64) *CreateGitDepotRequest {
	this := CreateGitDepotRequest{}
	this.DepotName = depotName
	this.ProjectId = projectId
	return &this
}

// NewCreateGitDepotRequestWithDefaults instantiates a new CreateGitDepotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGitDepotRequestWithDefaults() *CreateGitDepotRequest {
	this := CreateGitDepotRequest{}
	return &this
}

// GetDepotName returns the DepotName field value
func (o *CreateGitDepotRequest) GetDepotName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotName
}

// GetDepotNameOk returns a tuple with the DepotName field value
// and a boolean to check if the value has been set.
func (o *CreateGitDepotRequest) GetDepotNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotName, true
}

// SetDepotName sets field value
func (o *CreateGitDepotRequest) SetDepotName(v string) {
	o.DepotName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateGitDepotRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitDepotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return &o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateGitDepotRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateGitDepotRequest) SetDescription(v string) {
	o.Description = v
}

// GetProjectId returns the ProjectId field value
func (o *CreateGitDepotRequest) GetProjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateGitDepotRequest) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateGitDepotRequest) SetProjectId(v int64) {
	o.ProjectId = v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *CreateGitDepotRequest) GetShared() bool {
	if o == nil || utils.IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGitDepotRequest) GetSharedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Shared) {
		return nil, false
	}
	return &o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *CreateGitDepotRequest) HasShared() bool {
	if o != nil && !utils.IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *CreateGitDepotRequest) SetShared(v bool) {
	o.Shared = v
}

func (o CreateGitDepotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGitDepotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotName"] = o.DepotName
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["ProjectId"] = o.ProjectId
	if !utils.IsNil(o.Shared) {
		toSerialize["Shared"] = o.Shared
	}
	return toSerialize, nil
}

func (o *CreateGitDepotRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotName",
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

	varCreateGitDepotRequest := _CreateGitDepotRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateGitDepotRequest)

	if err != nil {
		return err
	}

	*o = CreateGitDepotRequest(varCreateGitDepotRequest)

	return err
}

type NullableCreateGitDepotRequest struct {
	value *CreateGitDepotRequest
	isSet bool
}

func (v NullableCreateGitDepotRequest) Get() *CreateGitDepotRequest {
	return v.value
}

func (v *NullableCreateGitDepotRequest) Set(val *CreateGitDepotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGitDepotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGitDepotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGitDepotRequest(val *CreateGitDepotRequest) *NullableCreateGitDepotRequest {
	return &NullableCreateGitDepotRequest{value: val, isSet: true}
}

func (v NullableCreateGitDepotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGitDepotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


