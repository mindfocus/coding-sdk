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

// checks if the CreateWikiRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateWikiRequest{}

// CreateWikiRequest struct for CreateWikiRequest
type CreateWikiRequest struct {
	// Wiki内容
	Content string `json:"Content"`
	// 备注
	Msg *string `json:"Msg,omitempty"`
	// 父级iid
	ParentIid int64 `json:"ParentIid"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// Wiki标题
	Title string `json:"Title"`
}

type _CreateWikiRequest CreateWikiRequest

// NewCreateWikiRequest instantiates a new CreateWikiRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWikiRequest(content string, parentIid int64, projectName string, title string) *CreateWikiRequest {
	this := CreateWikiRequest{}
	this.Content = content
	this.ParentIid = parentIid
	this.ProjectName = projectName
	this.Title = title
	return &this
}

// NewCreateWikiRequestWithDefaults instantiates a new CreateWikiRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWikiRequestWithDefaults() *CreateWikiRequest {
	this := CreateWikiRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateWikiRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateWikiRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateWikiRequest) SetContent(v string) {
	o.Content = v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *CreateWikiRequest) GetMsg() string {
	if o == nil || utils.IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWikiRequest) GetMsgOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *CreateWikiRequest) HasMsg() bool {
	if o != nil && !utils.IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *CreateWikiRequest) SetMsg(v string) {
	o.Msg = &v
}

// GetParentIid returns the ParentIid field value
func (o *CreateWikiRequest) GetParentIid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ParentIid
}

// GetParentIidOk returns a tuple with the ParentIid field value
// and a boolean to check if the value has been set.
func (o *CreateWikiRequest) GetParentIidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentIid, true
}

// SetParentIid sets field value
func (o *CreateWikiRequest) SetParentIid(v int64) {
	o.ParentIid = v
}

// GetProjectName returns the ProjectName field value
func (o *CreateWikiRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *CreateWikiRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *CreateWikiRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetTitle returns the Title field value
func (o *CreateWikiRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateWikiRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateWikiRequest) SetTitle(v string) {
	o.Title = v
}

func (o CreateWikiRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWikiRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Content"] = o.Content
	if !utils.IsNil(o.Msg) {
		toSerialize["Msg"] = o.Msg
	}
	toSerialize["ParentIid"] = o.ParentIid
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["Title"] = o.Title
	return toSerialize, nil
}

func (o *CreateWikiRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Content",
		"ParentIid",
		"ProjectName",
		"Title",
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

	varCreateWikiRequest := _CreateWikiRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWikiRequest)

	if err != nil {
		return err
	}

	*o = CreateWikiRequest(varCreateWikiRequest)

	return err
}

type NullableCreateWikiRequest struct {
	value *CreateWikiRequest
	isSet bool
}

func (v NullableCreateWikiRequest) Get() *CreateWikiRequest {
	return v.value
}

func (v *NullableCreateWikiRequest) Set(val *CreateWikiRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWikiRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWikiRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWikiRequest(val *CreateWikiRequest) *NullableCreateWikiRequest {
	return &NullableCreateWikiRequest{value: val, isSet: true}
}

func (v NullableCreateWikiRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWikiRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

