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

// checks if the ArtifactsOpenApiArtifactCreditsRuleData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ArtifactsOpenApiArtifactCreditsRuleData{}

// ArtifactsOpenApiArtifactCreditsRuleData struct for ArtifactsOpenApiArtifactCreditsRuleData
type ArtifactsOpenApiArtifactCreditsRuleData struct {
	// 制品版本
	Version string `json:"Version"`
	// 制品的类型(1-generic;2-docker;3-maven;4-npm;5-pypi;6-helm;7-composer;8-nuget;9-conan,10-cocoapods,11-rpm,12-Go)
	ArtifactType int32 `json:"ArtifactType"`
	// 制品名称计算规则：1-EQUAL(等于)，2-REGEX(正则表达式)
	PkgNameAlgorithm string `json:"PkgNameAlgorithm"`
	// 制品名称
	PkgName string `json:"PkgName"`
	// 制品版本计算规则：1-EQUAL(等于)，2-REGEX(正则表达式)
	VersionAlgorithm string `json:"VersionAlgorithm"`
}

type _ArtifactsOpenApiArtifactCreditsRuleData ArtifactsOpenApiArtifactCreditsRuleData

// NewArtifactsOpenApiArtifactCreditsRuleData instantiates a new ArtifactsOpenApiArtifactCreditsRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactsOpenApiArtifactCreditsRuleData(version string, artifactType int32, pkgNameAlgorithm string, pkgName string, versionAlgorithm string) *ArtifactsOpenApiArtifactCreditsRuleData {
	this := ArtifactsOpenApiArtifactCreditsRuleData{}
	this.Version = version
	this.ArtifactType = artifactType
	this.PkgNameAlgorithm = pkgNameAlgorithm
	this.PkgName = pkgName
	this.VersionAlgorithm = versionAlgorithm
	return &this
}

// NewArtifactsOpenApiArtifactCreditsRuleDataWithDefaults instantiates a new ArtifactsOpenApiArtifactCreditsRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactsOpenApiArtifactCreditsRuleDataWithDefaults() *ArtifactsOpenApiArtifactCreditsRuleData {
	this := ArtifactsOpenApiArtifactCreditsRuleData{}
	var version string = ""
	this.Version = version
	var pkgNameAlgorithm string = ""
	this.PkgNameAlgorithm = pkgNameAlgorithm
	var pkgName string = ""
	this.PkgName = pkgName
	var versionAlgorithm string = ""
	this.VersionAlgorithm = versionAlgorithm
	return &this
}

// GetVersion returns the Version field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetVersion(v string) {
	o.Version = v
}

// GetArtifactType returns the ArtifactType field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetArtifactType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ArtifactType
}

// GetArtifactTypeOk returns a tuple with the ArtifactType field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetArtifactTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtifactType, true
}

// SetArtifactType sets field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetArtifactType(v int32) {
	o.ArtifactType = v
}

// GetPkgNameAlgorithm returns the PkgNameAlgorithm field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgNameAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkgNameAlgorithm
}

// GetPkgNameAlgorithmOk returns a tuple with the PkgNameAlgorithm field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgNameAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkgNameAlgorithm, true
}

// SetPkgNameAlgorithm sets field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetPkgNameAlgorithm(v string) {
	o.PkgNameAlgorithm = v
}

// GetPkgName returns the PkgName field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PkgName
}

// GetPkgNameOk returns a tuple with the PkgName field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetPkgNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkgName, true
}

// SetPkgName sets field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetPkgName(v string) {
	o.PkgName = v
}

// GetVersionAlgorithm returns the VersionAlgorithm field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersionAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionAlgorithm
}

// GetVersionAlgorithmOk returns a tuple with the VersionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiArtifactCreditsRuleData) GetVersionAlgorithmOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionAlgorithm, true
}

// SetVersionAlgorithm sets field value
func (o *ArtifactsOpenApiArtifactCreditsRuleData) SetVersionAlgorithm(v string) {
	o.VersionAlgorithm = v
}

func (o ArtifactsOpenApiArtifactCreditsRuleData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactsOpenApiArtifactCreditsRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Version"] = o.Version
	toSerialize["ArtifactType"] = o.ArtifactType
	toSerialize["PkgNameAlgorithm"] = o.PkgNameAlgorithm
	toSerialize["PkgName"] = o.PkgName
	toSerialize["VersionAlgorithm"] = o.VersionAlgorithm
	return toSerialize, nil
}

func (o *ArtifactsOpenApiArtifactCreditsRuleData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Version",
		"ArtifactType",
		"PkgNameAlgorithm",
		"PkgName",
		"VersionAlgorithm",
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

	varArtifactsOpenApiArtifactCreditsRuleData := _ArtifactsOpenApiArtifactCreditsRuleData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtifactsOpenApiArtifactCreditsRuleData)

	if err != nil {
		return err
	}

	*o = ArtifactsOpenApiArtifactCreditsRuleData(varArtifactsOpenApiArtifactCreditsRuleData)

	return err
}

type NullableArtifactsOpenApiArtifactCreditsRuleData struct {
	value *ArtifactsOpenApiArtifactCreditsRuleData
	isSet bool
}

func (v NullableArtifactsOpenApiArtifactCreditsRuleData) Get() *ArtifactsOpenApiArtifactCreditsRuleData {
	return v.value
}

func (v *NullableArtifactsOpenApiArtifactCreditsRuleData) Set(val *ArtifactsOpenApiArtifactCreditsRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactsOpenApiArtifactCreditsRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactsOpenApiArtifactCreditsRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactsOpenApiArtifactCreditsRuleData(val *ArtifactsOpenApiArtifactCreditsRuleData) *NullableArtifactsOpenApiArtifactCreditsRuleData {
	return &NullableArtifactsOpenApiArtifactCreditsRuleData{value: val, isSet: true}
}

func (v NullableArtifactsOpenApiArtifactCreditsRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactsOpenApiArtifactCreditsRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


