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

// checks if the DescribeArtifactFileDownloadUrlRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactFileDownloadUrlRequest{}

// DescribeArtifactFileDownloadUrlRequest struct for DescribeArtifactFileDownloadUrlRequest
type DescribeArtifactFileDownloadUrlRequest struct {
	// 文件名称
	FileName string `json:"FileName"`
	// 包名
	Package string `json:"Package"`
	// 版本号
	PackageVersion string `json:"PackageVersion"`
	// 项目 ID
	ProjectId int32 `json:"ProjectId"`
	// 仓库名
	Repository string `json:"Repository"`
	// 下载链接超时时间（单位：秒），默认：300
	Timeout *int32 `json:"Timeout,omitempty"`
}

type _DescribeArtifactFileDownloadUrlRequest DescribeArtifactFileDownloadUrlRequest

// NewDescribeArtifactFileDownloadUrlRequest instantiates a new DescribeArtifactFileDownloadUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactFileDownloadUrlRequest(fileName string, package_ string, packageVersion string, projectId int32, repository string) *DescribeArtifactFileDownloadUrlRequest {
	this := DescribeArtifactFileDownloadUrlRequest{}
	this.FileName = fileName
	this.Package = package_
	this.PackageVersion = packageVersion
	this.ProjectId = projectId
	this.Repository = repository
	return &this
}

// NewDescribeArtifactFileDownloadUrlRequestWithDefaults instantiates a new DescribeArtifactFileDownloadUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactFileDownloadUrlRequestWithDefaults() *DescribeArtifactFileDownloadUrlRequest {
	this := DescribeArtifactFileDownloadUrlRequest{}
	return &this
}

// GetFileName returns the FileName field value
func (o *DescribeArtifactFileDownloadUrlRequest) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *DescribeArtifactFileDownloadUrlRequest) SetFileName(v string) {
	o.FileName = v
}

// GetPackage returns the Package field value
func (o *DescribeArtifactFileDownloadUrlRequest) GetPackage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Package
}

// GetPackageOk returns a tuple with the Package field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) GetPackageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Package, true
}

// SetPackage sets field value
func (o *DescribeArtifactFileDownloadUrlRequest) SetPackage(v string) {
	o.Package = v
}

// GetPackageVersion returns the PackageVersion field value
func (o *DescribeArtifactFileDownloadUrlRequest) GetPackageVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageVersion
}

// GetPackageVersionOk returns a tuple with the PackageVersion field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) GetPackageVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageVersion, true
}

// SetPackageVersion sets field value
func (o *DescribeArtifactFileDownloadUrlRequest) SetPackageVersion(v string) {
	o.PackageVersion = v
}

// GetProjectId returns the ProjectId field value
func (o *DescribeArtifactFileDownloadUrlRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeArtifactFileDownloadUrlRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

// GetRepository returns the Repository field value
func (o *DescribeArtifactFileDownloadUrlRequest) GetRepository() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value
// and a boolean to check if the value has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) GetRepositoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Repository, true
}

// SetRepository sets field value
func (o *DescribeArtifactFileDownloadUrlRequest) SetRepository(v string) {
	o.Repository = v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *DescribeArtifactFileDownloadUrlRequest) GetTimeout() int32 {
	if o == nil || utils.IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *DescribeArtifactFileDownloadUrlRequest) HasTimeout() bool {
	if o != nil && !utils.IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *DescribeArtifactFileDownloadUrlRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o DescribeArtifactFileDownloadUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactFileDownloadUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FileName"] = o.FileName
	toSerialize["Package"] = o.Package
	toSerialize["PackageVersion"] = o.PackageVersion
	toSerialize["ProjectId"] = o.ProjectId
	toSerialize["Repository"] = o.Repository
	if !utils.IsNil(o.Timeout) {
		toSerialize["Timeout"] = o.Timeout
	}
	return toSerialize, nil
}

func (o *DescribeArtifactFileDownloadUrlRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"FileName",
		"Package",
		"PackageVersion",
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

	varDescribeArtifactFileDownloadUrlRequest := _DescribeArtifactFileDownloadUrlRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeArtifactFileDownloadUrlRequest)

	if err != nil {
		return err
	}

	*o = DescribeArtifactFileDownloadUrlRequest(varDescribeArtifactFileDownloadUrlRequest)

	return err
}

type NullableDescribeArtifactFileDownloadUrlRequest struct {
	value *DescribeArtifactFileDownloadUrlRequest
	isSet bool
}

func (v NullableDescribeArtifactFileDownloadUrlRequest) Get() *DescribeArtifactFileDownloadUrlRequest {
	return v.value
}

func (v *NullableDescribeArtifactFileDownloadUrlRequest) Set(val *DescribeArtifactFileDownloadUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactFileDownloadUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactFileDownloadUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactFileDownloadUrlRequest(val *DescribeArtifactFileDownloadUrlRequest) *NullableDescribeArtifactFileDownloadUrlRequest {
	return &NullableDescribeArtifactFileDownloadUrlRequest{value: val, isSet: true}
}

func (v NullableDescribeArtifactFileDownloadUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactFileDownloadUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


