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

// checks if the ArtifactsOpenApiCreditRemoteTeam type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ArtifactsOpenApiCreditRemoteTeam{}

// ArtifactsOpenApiCreditRemoteTeam struct for ArtifactsOpenApiCreditRemoteTeam
type ArtifactsOpenApiCreditRemoteTeam struct {
	// 远程团队地址
	RemoteTeamUrl string `json:"RemoteTeamUrl"`
	// 用户名
	UserName string `json:"UserName"`
	// 远程仓库列表
	RemoteRepos []ArtifactsOpenApiRemoteRepoData `json:"RemoteRepos"`
	// 个人令牌base64编码
	Password string `json:"Password"`
}

type _ArtifactsOpenApiCreditRemoteTeam ArtifactsOpenApiCreditRemoteTeam

// NewArtifactsOpenApiCreditRemoteTeam instantiates a new ArtifactsOpenApiCreditRemoteTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactsOpenApiCreditRemoteTeam(remoteTeamUrl string, userName string, remoteRepos []ArtifactsOpenApiRemoteRepoData, password string) *ArtifactsOpenApiCreditRemoteTeam {
	this := ArtifactsOpenApiCreditRemoteTeam{}
	this.RemoteTeamUrl = remoteTeamUrl
	this.UserName = userName
	this.RemoteRepos = remoteRepos
	this.Password = password
	return &this
}

// NewArtifactsOpenApiCreditRemoteTeamWithDefaults instantiates a new ArtifactsOpenApiCreditRemoteTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactsOpenApiCreditRemoteTeamWithDefaults() *ArtifactsOpenApiCreditRemoteTeam {
	this := ArtifactsOpenApiCreditRemoteTeam{}
	var remoteTeamUrl string = ""
	this.RemoteTeamUrl = remoteTeamUrl
	var userName string = ""
	this.UserName = userName
	var password string = ""
	this.Password = password
	return &this
}

// GetRemoteTeamUrl returns the RemoteTeamUrl field value
func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteTeamUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteTeamUrl
}

// GetRemoteTeamUrlOk returns a tuple with the RemoteTeamUrl field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteTeamUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteTeamUrl, true
}

// SetRemoteTeamUrl sets field value
func (o *ArtifactsOpenApiCreditRemoteTeam) SetRemoteTeamUrl(v string) {
	o.RemoteTeamUrl = v
}

// GetUserName returns the UserName field value
func (o *ArtifactsOpenApiCreditRemoteTeam) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreditRemoteTeam) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *ArtifactsOpenApiCreditRemoteTeam) SetUserName(v string) {
	o.UserName = v
}

// GetRemoteRepos returns the RemoteRepos field value
func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteRepos() []ArtifactsOpenApiRemoteRepoData {
	if o == nil {
		var ret []ArtifactsOpenApiRemoteRepoData
		return ret
	}

	return o.RemoteRepos
}

// GetRemoteReposOk returns a tuple with the RemoteRepos field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreditRemoteTeam) GetRemoteReposOk() ([]ArtifactsOpenApiRemoteRepoData, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteRepos, true
}

// SetRemoteRepos sets field value
func (o *ArtifactsOpenApiCreditRemoteTeam) SetRemoteRepos(v []ArtifactsOpenApiRemoteRepoData) {
	o.RemoteRepos = v
}

// GetPassword returns the Password field value
func (o *ArtifactsOpenApiCreditRemoteTeam) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ArtifactsOpenApiCreditRemoteTeam) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ArtifactsOpenApiCreditRemoteTeam) SetPassword(v string) {
	o.Password = v
}

func (o ArtifactsOpenApiCreditRemoteTeam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactsOpenApiCreditRemoteTeam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["RemoteTeamUrl"] = o.RemoteTeamUrl
	toSerialize["UserName"] = o.UserName
	toSerialize["RemoteRepos"] = o.RemoteRepos
	toSerialize["Password"] = o.Password
	return toSerialize, nil
}

func (o *ArtifactsOpenApiCreditRemoteTeam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"RemoteTeamUrl",
		"UserName",
		"RemoteRepos",
		"Password",
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

	varArtifactsOpenApiCreditRemoteTeam := _ArtifactsOpenApiCreditRemoteTeam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtifactsOpenApiCreditRemoteTeam)

	if err != nil {
		return err
	}

	*o = ArtifactsOpenApiCreditRemoteTeam(varArtifactsOpenApiCreditRemoteTeam)

	return err
}

type NullableArtifactsOpenApiCreditRemoteTeam struct {
	value *ArtifactsOpenApiCreditRemoteTeam
	isSet bool
}

func (v NullableArtifactsOpenApiCreditRemoteTeam) Get() *ArtifactsOpenApiCreditRemoteTeam {
	return v.value
}

func (v *NullableArtifactsOpenApiCreditRemoteTeam) Set(val *ArtifactsOpenApiCreditRemoteTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactsOpenApiCreditRemoteTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactsOpenApiCreditRemoteTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactsOpenApiCreditRemoteTeam(val *ArtifactsOpenApiCreditRemoteTeam) *NullableArtifactsOpenApiCreditRemoteTeam {
	return &NullableArtifactsOpenApiCreditRemoteTeam{value: val, isSet: true}
}

func (v NullableArtifactsOpenApiCreditRemoteTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactsOpenApiCreditRemoteTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


