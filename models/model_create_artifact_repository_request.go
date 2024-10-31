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

// checks if the CreateArtifactRepositoryRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateArtifactRepositoryRequest{}

// CreateArtifactRepositoryRequest struct for CreateArtifactRepositoryRequest
type CreateArtifactRepositoryRequest struct {
	// 仓库权限范围:1-项目内;2-团队内;3-公开，默认：1-项目内
	AccessLevel *int32 `json:"AccessLevel,omitempty"`
	// 是否开启代理，仅支持当 Type 为 3-maven;4-npm; 5-PyPI;7-composer;10-cocoapods 时可为 true，默认：false
	AllowProxy *bool `json:"AllowProxy,omitempty"`
	// 仓库描述信息
	Description *string `json:"Description,omitempty"`
	// 项目 ID
	ProjectId int32 `json:"ProjectId"`
	// 仓库名称
	RepositoryName string `json:"RepositoryName"`
	// 仓库类型:1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan;10-cocoapods;11-rpm
	Type int32 `json:"Type"`
}

type _CreateArtifactRepositoryRequest CreateArtifactRepositoryRequest

// NewCreateArtifactRepositoryRequest instantiates a new CreateArtifactRepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArtifactRepositoryRequest(projectId int32, repositoryName string, type_ int32) *CreateArtifactRepositoryRequest {
	this := CreateArtifactRepositoryRequest{}
	this.ProjectId = projectId
	this.RepositoryName = repositoryName
	this.Type = type_
	return &this
}

// NewCreateArtifactRepositoryRequestWithDefaults instantiates a new CreateArtifactRepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArtifactRepositoryRequestWithDefaults() *CreateArtifactRepositoryRequest {
	this := CreateArtifactRepositoryRequest{}
	return &this
}

// GetAccessLevel returns the AccessLevel field value if set, zero value otherwise.
func (o *CreateArtifactRepositoryRequest) GetAccessLevel() int32 {
	if o == nil || utils.IsNil(o.AccessLevel) {
		var ret int32
		return ret
	}
	return *o.AccessLevel
}

// GetAccessLevelOk returns a tuple with the AccessLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepositoryRequest) GetAccessLevelOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AccessLevel) {
		return nil, false
	}
	return o.AccessLevel, true
}

// HasAccessLevel returns a boolean if a field has been set.
func (o *CreateArtifactRepositoryRequest) HasAccessLevel() bool {
	if o != nil && !utils.IsNil(o.AccessLevel) {
		return true
	}

	return false
}

// SetAccessLevel gets a reference to the given int32 and assigns it to the AccessLevel field.
func (o *CreateArtifactRepositoryRequest) SetAccessLevel(v int32) {
	o.AccessLevel = &v
}

// GetAllowProxy returns the AllowProxy field value if set, zero value otherwise.
func (o *CreateArtifactRepositoryRequest) GetAllowProxy() bool {
	if o == nil || utils.IsNil(o.AllowProxy) {
		var ret bool
		return ret
	}
	return *o.AllowProxy
}

// GetAllowProxyOk returns a tuple with the AllowProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepositoryRequest) GetAllowProxyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AllowProxy) {
		return nil, false
	}
	return o.AllowProxy, true
}

// HasAllowProxy returns a boolean if a field has been set.
func (o *CreateArtifactRepositoryRequest) HasAllowProxy() bool {
	if o != nil && !utils.IsNil(o.AllowProxy) {
		return true
	}

	return false
}

// SetAllowProxy gets a reference to the given bool and assigns it to the AllowProxy field.
func (o *CreateArtifactRepositoryRequest) SetAllowProxy(v bool) {
	o.AllowProxy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateArtifactRepositoryRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepositoryRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateArtifactRepositoryRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateArtifactRepositoryRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProjectId returns the ProjectId field value
func (o *CreateArtifactRepositoryRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepositoryRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateArtifactRepositoryRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetRepositoryName returns the RepositoryName field value
func (o *CreateArtifactRepositoryRequest) GetRepositoryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RepositoryName
}

// GetRepositoryNameOk returns a tuple with the RepositoryName field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepositoryRequest) GetRepositoryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepositoryName, true
}

// SetRepositoryName sets field value
func (o *CreateArtifactRepositoryRequest) SetRepositoryName(v string) {
	o.RepositoryName = v
}

// GetType returns the Type field value
func (o *CreateArtifactRepositoryRequest) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepositoryRequest) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateArtifactRepositoryRequest) SetType(v int32) {
	o.Type = v
}

func (o CreateArtifactRepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateArtifactRepositoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccessLevel) {
		toSerialize["AccessLevel"] = o.AccessLevel
	}
	if !utils.IsNil(o.AllowProxy) {
		toSerialize["AllowProxy"] = o.AllowProxy
	}
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["ProjectId"] = o.ProjectId
	toSerialize["RepositoryName"] = o.RepositoryName
	toSerialize["Type"] = o.Type
	return toSerialize, nil
}

func (o *CreateArtifactRepositoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectId",
		"RepositoryName",
		"Type",
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

	varCreateArtifactRepositoryRequest := _CreateArtifactRepositoryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateArtifactRepositoryRequest)

	if err != nil {
		return err
	}

	*o = CreateArtifactRepositoryRequest(varCreateArtifactRepositoryRequest)

	return err
}

type NullableCreateArtifactRepositoryRequest struct {
	value *CreateArtifactRepositoryRequest
	isSet bool
}

func (v NullableCreateArtifactRepositoryRequest) Get() *CreateArtifactRepositoryRequest {
	return v.value
}

func (v *NullableCreateArtifactRepositoryRequest) Set(val *CreateArtifactRepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArtifactRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArtifactRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArtifactRepositoryRequest(val *CreateArtifactRepositoryRequest) *NullableCreateArtifactRepositoryRequest {
	return &NullableCreateArtifactRepositoryRequest{value: val, isSet: true}
}

func (v NullableCreateArtifactRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArtifactRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

