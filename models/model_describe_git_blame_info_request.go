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

// checks if the DescribeGitBlameInfoRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitBlameInfoRequest{}

// DescribeGitBlameInfoRequest struct for DescribeGitBlameInfoRequest
type DescribeGitBlameInfoRequest struct {
	// 提交sha
	CommitSha string `json:"CommitSha"`
	// 仓库路径
	DepotPath string `json:"DepotPath"`
	// 文件路径
	FilePath string `json:"FilePath"`
	// 结束行
	LineEnd *int64 `json:"LineEnd,omitempty"`
	// 开始行
	LineSnat *int64 `json:"LineSnat,omitempty"`
}

type _DescribeGitBlameInfoRequest DescribeGitBlameInfoRequest

// NewDescribeGitBlameInfoRequest instantiates a new DescribeGitBlameInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitBlameInfoRequest(commitSha string, depotPath string, filePath string) *DescribeGitBlameInfoRequest {
	this := DescribeGitBlameInfoRequest{}
	this.CommitSha = commitSha
	this.DepotPath = depotPath
	this.FilePath = filePath
	return &this
}

// NewDescribeGitBlameInfoRequestWithDefaults instantiates a new DescribeGitBlameInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitBlameInfoRequestWithDefaults() *DescribeGitBlameInfoRequest {
	this := DescribeGitBlameInfoRequest{}
	return &this
}

// GetCommitSha returns the CommitSha field value
func (o *DescribeGitBlameInfoRequest) GetCommitSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitSha
}

// GetCommitShaOk returns a tuple with the CommitSha field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBlameInfoRequest) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitSha, true
}

// SetCommitSha sets field value
func (o *DescribeGitBlameInfoRequest) SetCommitSha(v string) {
	o.CommitSha = v
}

// GetDepotPath returns the DepotPath field value
func (o *DescribeGitBlameInfoRequest) GetDepotPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBlameInfoRequest) GetDepotPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotPath, true
}

// SetDepotPath sets field value
func (o *DescribeGitBlameInfoRequest) SetDepotPath(v string) {
	o.DepotPath = v
}

// GetFilePath returns the FilePath field value
func (o *DescribeGitBlameInfoRequest) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBlameInfoRequest) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value
func (o *DescribeGitBlameInfoRequest) SetFilePath(v string) {
	o.FilePath = v
}

// GetLineEnd returns the LineEnd field value if set, zero value otherwise.
func (o *DescribeGitBlameInfoRequest) GetLineEnd() int64 {
	if o == nil || utils.IsNil(o.LineEnd) {
		var ret int64
		return ret
	}
	return *o.LineEnd
}

// GetLineEndOk returns a tuple with the LineEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitBlameInfoRequest) GetLineEndOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LineEnd) {
		return nil, false
	}
	return o.LineEnd, true
}

// HasLineEnd returns a boolean if a field has been set.
func (o *DescribeGitBlameInfoRequest) HasLineEnd() bool {
	if o != nil && !utils.IsNil(o.LineEnd) {
		return true
	}

	return false
}

// SetLineEnd gets a reference to the given int64 and assigns it to the LineEnd field.
func (o *DescribeGitBlameInfoRequest) SetLineEnd(v int64) {
	o.LineEnd = &v
}

// GetLineSnat returns the LineSnat field value if set, zero value otherwise.
func (o *DescribeGitBlameInfoRequest) GetLineSnat() int64 {
	if o == nil || utils.IsNil(o.LineSnat) {
		var ret int64
		return ret
	}
	return *o.LineSnat
}

// GetLineSnatOk returns a tuple with the LineSnat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitBlameInfoRequest) GetLineSnatOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.LineSnat) {
		return nil, false
	}
	return o.LineSnat, true
}

// HasLineSnat returns a boolean if a field has been set.
func (o *DescribeGitBlameInfoRequest) HasLineSnat() bool {
	if o != nil && !utils.IsNil(o.LineSnat) {
		return true
	}

	return false
}

// SetLineSnat gets a reference to the given int64 and assigns it to the LineSnat field.
func (o *DescribeGitBlameInfoRequest) SetLineSnat(v int64) {
	o.LineSnat = &v
}

func (o DescribeGitBlameInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitBlameInfoRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CommitSha"] = o.CommitSha
	toSerialize["DepotPath"] = o.DepotPath
	toSerialize["FilePath"] = o.FilePath
	if !utils.IsNil(o.LineEnd) {
		toSerialize["LineEnd"] = o.LineEnd
	}
	if !utils.IsNil(o.LineSnat) {
		toSerialize["LineSnat"] = o.LineSnat
	}
	return toSerialize, nil
}

func (o *DescribeGitBlameInfoRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"CommitSha",
		"DepotPath",
		"FilePath",
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

	varDescribeGitBlameInfoRequest := _DescribeGitBlameInfoRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitBlameInfoRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitBlameInfoRequest(varDescribeGitBlameInfoRequest)

	return err
}

type NullableDescribeGitBlameInfoRequest struct {
	value *DescribeGitBlameInfoRequest
	isSet bool
}

func (v NullableDescribeGitBlameInfoRequest) Get() *DescribeGitBlameInfoRequest {
	return v.value
}

func (v *NullableDescribeGitBlameInfoRequest) Set(val *DescribeGitBlameInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitBlameInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitBlameInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitBlameInfoRequest(val *DescribeGitBlameInfoRequest) *NullableDescribeGitBlameInfoRequest {
	return &NullableDescribeGitBlameInfoRequest{value: val, isSet: true}
}

func (v NullableDescribeGitBlameInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitBlameInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


