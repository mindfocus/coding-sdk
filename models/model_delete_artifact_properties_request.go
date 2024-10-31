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

// checks if the DeleteArtifactPropertiesRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteArtifactPropertiesRequest{}

// DeleteArtifactPropertiesRequest struct for DeleteArtifactPropertiesRequest
type DeleteArtifactPropertiesRequest struct {
	// 包名
	Package string `json:"Package"`
	// 版本号
	PackageVersion string `json:"PackageVersion"`
	// 项目 ID
	ProjectId int32 `json:"ProjectId"`
	// 属性名称列表（ 以 ‘coding.’ 作为属性名称开头的属性，将不可删除）
	PropertyNameSet []string `json:"PropertyNameSet"`
	// 仓库名
	Repository string `json:"Repository"`
}

type _DeleteArtifactPropertiesRequest DeleteArtifactPropertiesRequest

// NewDeleteArtifactPropertiesRequest instantiates a new DeleteArtifactPropertiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteArtifactPropertiesRequest(package_ string, packageVersion string, projectId int32, propertyNameSet []string, repository string) *DeleteArtifactPropertiesRequest {
	this := DeleteArtifactPropertiesRequest{}
	this.Package = package_
	this.PackageVersion = packageVersion
	this.ProjectId = projectId
	this.PropertyNameSet = propertyNameSet
	this.Repository = repository
	return &this
}

// NewDeleteArtifactPropertiesRequestWithDefaults instantiates a new DeleteArtifactPropertiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteArtifactPropertiesRequestWithDefaults() *DeleteArtifactPropertiesRequest {
	this := DeleteArtifactPropertiesRequest{}
	return &this
}

// GetPackage returns the Package field value
func (o *DeleteArtifactPropertiesRequest) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *DeleteArtifactPropertiesRequest) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *DeleteArtifactPropertiesRequest) SetPackage(v string) {
	o.Package = v
}

// GetPackageVersion returns the PackageVersion field value
func (o *DeleteArtifactPropertiesRequest) GetPackageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value
// and a boolean to check if the value has been set.
func (o *DeleteArtifactPropertiesRequest) GetPackageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageVersion, true
}

// SetPackageVersion sets field value
func (o *DeleteArtifactPropertiesRequest) SetPackageVersion(v string) {
	o.PackageVersion = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteArtifactPropertiesRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteArtifactPropertiesRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteArtifactPropertiesRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetPropertyNameSet returns the PropertyNameSet field value
func (o *DeleteArtifactPropertiesRequest) GetPropertyNameSet() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertyNameSet
}

// GetPropertyNameSetOk returns a tuple with the PropertyNameSet field value
// and a boolean to check if the value has been set.
func (o *DeleteArtifactPropertiesRequest) GetPropertyNameSetOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertyNameSet, true
}

// SetPropertyNameSet sets field value
func (o *DeleteArtifactPropertiesRequest) SetPropertyNameSet(v []string) {
	o.PropertyNameSet = v
}

// GetRepository returns the Repository field value
func (o *DeleteArtifactPropertiesRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *DeleteArtifactPropertiesRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *DeleteArtifactPropertiesRequest) SetRepository(v string) {
	o.Repository = v
}

func (o DeleteArtifactPropertiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteArtifactPropertiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Package"] = o.Package
	toSerialize["PackageVersion"] = o.PackageVersion
	toSerialize["ProjectId"] = o.ProjectId
	toSerialize["PropertyNameSet"] = o.PropertyNameSet
	toSerialize["Repository"] = o.Repository
	return toSerialize, nil
}

func (o *DeleteArtifactPropertiesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Package",
		"PackageVersion",
		"ProjectId",
		"PropertyNameSet",
		"Repository",
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

	varDeleteArtifactPropertiesRequest := _DeleteArtifactPropertiesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteArtifactPropertiesRequest)

	if err != nil {
		return err
	}

	*o = DeleteArtifactPropertiesRequest(varDeleteArtifactPropertiesRequest)

	return err
}

type NullableDeleteArtifactPropertiesRequest struct {
	value *DeleteArtifactPropertiesRequest
	isSet bool
}

func (v NullableDeleteArtifactPropertiesRequest) Get() *DeleteArtifactPropertiesRequest {
	return v.value
}

func (v *NullableDeleteArtifactPropertiesRequest) Set(val *DeleteArtifactPropertiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteArtifactPropertiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteArtifactPropertiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteArtifactPropertiesRequest(val *DeleteArtifactPropertiesRequest) *NullableDeleteArtifactPropertiesRequest {
	return &NullableDeleteArtifactPropertiesRequest{value: val, isSet: true}
}

func (v NullableDeleteArtifactPropertiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteArtifactPropertiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


