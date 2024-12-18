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

// checks if the DescribeArtifactVersionListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactVersionListRequest{}

// DescribeArtifactVersionListRequest struct for DescribeArtifactVersionListRequest
type DescribeArtifactVersionListRequest struct {
	// 包名称
	Package string `json:"Package"`
	// 页码，默认：1
	PageNumber *int32 `json:"PageNumber,omitempty"`
	// 每页展示数量，默认：10
	PageSize *int32 `json:"PageSize,omitempty"`
	// 项目 ID
	ProjectId int32 `json:"ProjectId"`
	// 仓库名称
	Repository string `json:"Repository"`
}

type _DescribeArtifactVersionListRequest DescribeArtifactVersionListRequest

// NewDescribeArtifactVersionListRequest instantiates a new DescribeArtifactVersionListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactVersionListRequest(package_ string, projectId int32, repository string) *DescribeArtifactVersionListRequest {
	this := DescribeArtifactVersionListRequest{}
	this.Package = package_
	this.ProjectId = projectId
	this.Repository = repository
	return &this
}

// NewDescribeArtifactVersionListRequestWithDefaults instantiates a new DescribeArtifactVersionListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactVersionListRequestWithDefaults() *DescribeArtifactVersionListRequest {
	this := DescribeArtifactVersionListRequest{}
	return &this
}

// GetPackage returns the Package field value
func (o *DescribeArtifactVersionListRequest) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionListRequest) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *DescribeArtifactVersionListRequest) SetPackage(v string) {
	o.Package = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeArtifactVersionListRequest) GetPageNumber() int32 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionListRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeArtifactVersionListRequest) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *DescribeArtifactVersionListRequest) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeArtifactVersionListRequest) GetPageSize() int32 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionListRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeArtifactVersionListRequest) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *DescribeArtifactVersionListRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetProjectId returns the ProjectId field value
func (o *DescribeArtifactVersionListRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionListRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeArtifactVersionListRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetRepository returns the Repository field value
func (o *DescribeArtifactVersionListRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionListRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *DescribeArtifactVersionListRequest) SetRepository(v string) {
	o.Repository = v
}

func (o DescribeArtifactVersionListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactVersionListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Package"] = o.Package
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	toSerialize["ProjectId"] = o.ProjectId
	toSerialize["Repository"] = o.Repository
	return toSerialize, nil
}

func (o *DescribeArtifactVersionListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Package",
		"ProjectId",
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

	varDescribeArtifactVersionListRequest := _DescribeArtifactVersionListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeArtifactVersionListRequest)

	if err != nil {
		return err
	}

	*o = DescribeArtifactVersionListRequest(varDescribeArtifactVersionListRequest)

	return err
}

type NullableDescribeArtifactVersionListRequest struct {
	value *DescribeArtifactVersionListRequest
	isSet bool
}

func (v NullableDescribeArtifactVersionListRequest) Get() *DescribeArtifactVersionListRequest {
	return v.value
}

func (v *NullableDescribeArtifactVersionListRequest) Set(val *DescribeArtifactVersionListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactVersionListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactVersionListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactVersionListRequest(val *DescribeArtifactVersionListRequest) *NullableDescribeArtifactVersionListRequest {
	return &NullableDescribeArtifactVersionListRequest{value: val, isSet: true}
}

func (v NullableDescribeArtifactVersionListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactVersionListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


