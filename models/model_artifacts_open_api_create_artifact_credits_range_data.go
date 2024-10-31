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

// checks if the ArtifactsOpenApiCreateArtifactCreditsRangeData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ArtifactsOpenApiCreateArtifactCreditsRangeData{}

// ArtifactsOpenApiCreateArtifactCreditsRangeData struct for ArtifactsOpenApiCreateArtifactCreditsRangeData
type ArtifactsOpenApiCreateArtifactCreditsRangeData struct {
	RemoteTeam *ArtifactsOpenApiCreditRemoteTeam `json:"RemoteTeam,omitempty"`
	// 生效范围（TEAM:团队；REMOTE_TEAM:远程团队；PROJECT:项目;REPO:仓库）
	RangeType string `json:"RangeType"`
	// 适用范围ID
	RangeId int64 `json:"RangeId"`
}

type _ArtifactsOpenApiCreateArtifactCreditsRangeData ArtifactsOpenApiCreateArtifactCreditsRangeData

// NewArtifactsOpenApiCreateArtifactCreditsRangeData instantiates a new ArtifactsOpenApiCreateArtifactCreditsRangeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactsOpenApiCreateArtifactCreditsRangeData(rangeType string, rangeId int64) *ArtifactsOpenApiCreateArtifactCreditsRangeData {
	this := ArtifactsOpenApiCreateArtifactCreditsRangeData{}
	this.RangeType = rangeType
	this.RangeId = rangeId
	return &this
}

// NewArtifactsOpenApiCreateArtifactCreditsRangeDataWithDefaults instantiates a new ArtifactsOpenApiCreateArtifactCreditsRangeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactsOpenApiCreateArtifactCreditsRangeDataWithDefaults() *ArtifactsOpenApiCreateArtifactCreditsRangeData {
	this := ArtifactsOpenApiCreateArtifactCreditsRangeData{}
	var rangeType string = ""
	this.RangeType = rangeType
	var rangeId int64 = 0
	this.RangeId = rangeId
	return &this
}

// GetRemoteTeam returns the RemoteTeam field value if set, zero value otherwise.
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRemoteTeam() ArtifactsOpenApiCreditRemoteTeam {
	if o == nil || utils.IsNil(o.RemoteTeam) {
		var ret ArtifactsOpenApiCreditRemoteTeam
		return ret
	}
	return *o.RemoteTeam
}

// GetRemoteTeamOk returns a tuple with the RemoteTeam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRemoteTeamOk() (*ArtifactsOpenApiCreditRemoteTeam, bool) {
	if o == nil || utils.IsNil(o.RemoteTeam) {
		return nil, false
	}
	return o.RemoteTeam, true
}

// HasRemoteTeam returns a boolean if a field has been set.
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) HasRemoteTeam() bool {
	if o != nil && !utils.IsNil(o.RemoteTeam) {
		return true
	}

	return false
}

// SetRemoteTeam gets a reference to the given ArtifactsOpenApiCreditRemoteTeam and assigns it to the RemoteTeam field.
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) SetRemoteTeam(v ArtifactsOpenApiCreditRemoteTeam) {
	o.RemoteTeam = &v
}

// GetRangeType returns the RangeType field value
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RangeType
}

// GetRangeTypeOk returns a tuple with the RangeType field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeType, true
}

// SetRangeType sets field value
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) SetRangeType(v string) {
	o.RangeType = v
}

// GetRangeId returns the RangeId field value
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RangeId, true
}

// SetRangeId sets field value
func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) SetRangeId(v int64) {
	o.RangeId = v
}

func (o ArtifactsOpenApiCreateArtifactCreditsRangeData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactsOpenApiCreateArtifactCreditsRangeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.RemoteTeam) {
		toSerialize["RemoteTeam"] = o.RemoteTeam
	}
	toSerialize["RangeType"] = o.RangeType
	toSerialize["RangeId"] = o.RangeId
	return toSerialize, nil
}

func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"RangeType",
		"RangeId",
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

	varArtifactsOpenApiCreateArtifactCreditsRangeData := _ArtifactsOpenApiCreateArtifactCreditsRangeData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtifactsOpenApiCreateArtifactCreditsRangeData)

	if err != nil {
		return err
	}

	*o = ArtifactsOpenApiCreateArtifactCreditsRangeData(varArtifactsOpenApiCreateArtifactCreditsRangeData)

	return err
}

type NullableArtifactsOpenApiCreateArtifactCreditsRangeData struct {
	value *ArtifactsOpenApiCreateArtifactCreditsRangeData
	isSet bool
}

func (v NullableArtifactsOpenApiCreateArtifactCreditsRangeData) Get() *ArtifactsOpenApiCreateArtifactCreditsRangeData {
	return v.value
}

func (v *NullableArtifactsOpenApiCreateArtifactCreditsRangeData) Set(val *ArtifactsOpenApiCreateArtifactCreditsRangeData) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactsOpenApiCreateArtifactCreditsRangeData) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactsOpenApiCreateArtifactCreditsRangeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactsOpenApiCreateArtifactCreditsRangeData(val *ArtifactsOpenApiCreateArtifactCreditsRangeData) *NullableArtifactsOpenApiCreateArtifactCreditsRangeData {
	return &NullableArtifactsOpenApiCreateArtifactCreditsRangeData{value: val, isSet: true}
}

func (v NullableArtifactsOpenApiCreateArtifactCreditsRangeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactsOpenApiCreateArtifactCreditsRangeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

