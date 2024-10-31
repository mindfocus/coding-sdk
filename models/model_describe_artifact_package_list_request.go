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

// checks if the DescribeArtifactPackageListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactPackageListRequest{}

// DescribeArtifactPackageListRequest struct for DescribeArtifactPackageListRequest
type DescribeArtifactPackageListRequest struct {
	// 包名前缀
	PackagePrefix *string `json:"PackagePrefix,omitempty"`
	// 页码，默认：1
	PageNumber *int32 `json:"PageNumber,omitempty"`
	// 每页展示数量，默认：10
	PageSize *int32 `json:"PageSize,omitempty"`
	// 项目 ID
	ProjectId int32 `json:"ProjectId"`
	// 仓库名称
	Repository string `json:"Repository"`
}

type _DescribeArtifactPackageListRequest DescribeArtifactPackageListRequest

// NewDescribeArtifactPackageListRequest instantiates a new DescribeArtifactPackageListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactPackageListRequest(projectId int32, repository string) *DescribeArtifactPackageListRequest {
	this := DescribeArtifactPackageListRequest{}
	this.ProjectId = projectId
	this.Repository = repository
	return &this
}

// NewDescribeArtifactPackageListRequestWithDefaults instantiates a new DescribeArtifactPackageListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactPackageListRequestWithDefaults() *DescribeArtifactPackageListRequest {
	this := DescribeArtifactPackageListRequest{}
	return &this
}

// GetPackagePrefix returns the PackagePrefix field value if set, zero value otherwise.
func (o *DescribeArtifactPackageListRequest) GetPackagePrefix() string {
	if o == nil || utils.IsNil(o.PackagePrefix) {
		var ret string
		return ret
	}
	return *o.PackagePrefix
}

// GetPackagePrefixOk returns a tuple with the PackagePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactPackageListRequest) GetPackagePrefixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PackagePrefix) {
		return nil, false
	}
	return o.PackagePrefix, true
}

// HasPackagePrefix returns a boolean if a field has been set.
func (o *DescribeArtifactPackageListRequest) HasPackagePrefix() bool {
	if o != nil && !utils.IsNil(o.PackagePrefix) {
		return true
	}

	return false
}

// SetPackagePrefix gets a reference to the given string and assigns it to the PackagePrefix field.
func (o *DescribeArtifactPackageListRequest) SetPackagePrefix(v string) {
	o.PackagePrefix = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeArtifactPackageListRequest) GetPageNumber() int32 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactPackageListRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeArtifactPackageListRequest) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *DescribeArtifactPackageListRequest) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeArtifactPackageListRequest) GetPageSize() int32 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactPackageListRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeArtifactPackageListRequest) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *DescribeArtifactPackageListRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetProjectId returns the ProjectId field value
func (o *DescribeArtifactPackageListRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactPackageListRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeArtifactPackageListRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetRepository returns the Repository field value
func (o *DescribeArtifactPackageListRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactPackageListRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *DescribeArtifactPackageListRequest) SetRepository(v string) {
	o.Repository = v
}

func (o DescribeArtifactPackageListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactPackageListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PackagePrefix) {
		toSerialize["PackagePrefix"] = o.PackagePrefix
	}
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

func (o *DescribeArtifactPackageListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDescribeArtifactPackageListRequest := _DescribeArtifactPackageListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeArtifactPackageListRequest)

	if err != nil {
		return err
	}

	*o = DescribeArtifactPackageListRequest(varDescribeArtifactPackageListRequest)

	return err
}

type NullableDescribeArtifactPackageListRequest struct {
	value *DescribeArtifactPackageListRequest
	isSet bool
}

func (v NullableDescribeArtifactPackageListRequest) Get() *DescribeArtifactPackageListRequest {
	return v.value
}

func (v *NullableDescribeArtifactPackageListRequest) Set(val *DescribeArtifactPackageListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactPackageListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactPackageListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactPackageListRequest(val *DescribeArtifactPackageListRequest) *NullableDescribeArtifactPackageListRequest {
	return &NullableDescribeArtifactPackageListRequest{value: val, isSet: true}
}

func (v NullableDescribeArtifactPackageListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactPackageListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


