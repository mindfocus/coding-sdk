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

// checks if the DescribeIssueAttachmentPreSignedUrlRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeIssueAttachmentPreSignedUrlRequest{}

// DescribeIssueAttachmentPreSignedUrlRequest struct for DescribeIssueAttachmentPreSignedUrlRequest
type DescribeIssueAttachmentPreSignedUrlRequest struct {
	// 文件名
	FileName string `json:"FileName"`
	// 文件大小
	FileSize int64 `json:"FileSize"`
	// 项目名
	ProjectName string `json:"ProjectName"`
}

type _DescribeIssueAttachmentPreSignedUrlRequest DescribeIssueAttachmentPreSignedUrlRequest

// NewDescribeIssueAttachmentPreSignedUrlRequest instantiates a new DescribeIssueAttachmentPreSignedUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeIssueAttachmentPreSignedUrlRequest(fileName string, fileSize int64, projectName string) *DescribeIssueAttachmentPreSignedUrlRequest {
	this := DescribeIssueAttachmentPreSignedUrlRequest{}
	this.FileName = fileName
	this.FileSize = fileSize
	this.ProjectName = projectName
	return &this
}

// NewDescribeIssueAttachmentPreSignedUrlRequestWithDefaults instantiates a new DescribeIssueAttachmentPreSignedUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeIssueAttachmentPreSignedUrlRequestWithDefaults() *DescribeIssueAttachmentPreSignedUrlRequest {
	this := DescribeIssueAttachmentPreSignedUrlRequest{}
	var fileName string = ""
	this.FileName = fileName
	var fileSize int64 = 0
	this.FileSize = fileSize
	var projectName string = ""
	this.ProjectName = projectName
	return &this
}

// GetFileName returns the FileName field value
func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *DescribeIssueAttachmentPreSignedUrlRequest) SetFileName(v string) {
	o.FileName = v
}

// GetFileSize returns the FileSize field value
func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileSize, true
}

// SetFileSize sets field value
func (o *DescribeIssueAttachmentPreSignedUrlRequest) SetFileSize(v int64) {
	o.FileSize = v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeIssueAttachmentPreSignedUrlRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeIssueAttachmentPreSignedUrlRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o DescribeIssueAttachmentPreSignedUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeIssueAttachmentPreSignedUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["FileName"] = o.FileName
	toSerialize["FileSize"] = o.FileSize
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *DescribeIssueAttachmentPreSignedUrlRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"FileName",
		"FileSize",
		"ProjectName",
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

	varDescribeIssueAttachmentPreSignedUrlRequest := _DescribeIssueAttachmentPreSignedUrlRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeIssueAttachmentPreSignedUrlRequest)

	if err != nil {
		return err
	}

	*o = DescribeIssueAttachmentPreSignedUrlRequest(varDescribeIssueAttachmentPreSignedUrlRequest)

	return err
}

type NullableDescribeIssueAttachmentPreSignedUrlRequest struct {
	value *DescribeIssueAttachmentPreSignedUrlRequest
	isSet bool
}

func (v NullableDescribeIssueAttachmentPreSignedUrlRequest) Get() *DescribeIssueAttachmentPreSignedUrlRequest {
	return v.value
}

func (v *NullableDescribeIssueAttachmentPreSignedUrlRequest) Set(val *DescribeIssueAttachmentPreSignedUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeIssueAttachmentPreSignedUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeIssueAttachmentPreSignedUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeIssueAttachmentPreSignedUrlRequest(val *DescribeIssueAttachmentPreSignedUrlRequest) *NullableDescribeIssueAttachmentPreSignedUrlRequest {
	return &NullableDescribeIssueAttachmentPreSignedUrlRequest{value: val, isSet: true}
}

func (v NullableDescribeIssueAttachmentPreSignedUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeIssueAttachmentPreSignedUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


