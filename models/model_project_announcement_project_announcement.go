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

// checks if the ProjectAnnouncementProjectAnnouncement type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProjectAnnouncementProjectAnnouncement{}

// ProjectAnnouncementProjectAnnouncement struct for ProjectAnnouncementProjectAnnouncement
type ProjectAnnouncementProjectAnnouncement struct {
	// 公告内容
	ContentHtml string `json:"ContentHtml"`
	// markdown 源文本
	ContentOrigin string `json:"ContentOrigin"`
	// 公告id
	Id int64 `json:"Id"`
}

type _ProjectAnnouncementProjectAnnouncement ProjectAnnouncementProjectAnnouncement

// NewProjectAnnouncementProjectAnnouncement instantiates a new ProjectAnnouncementProjectAnnouncement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectAnnouncementProjectAnnouncement(contentHtml string, contentOrigin string, id int64) *ProjectAnnouncementProjectAnnouncement {
	this := ProjectAnnouncementProjectAnnouncement{}
	this.ContentHtml = contentHtml
	this.ContentOrigin = contentOrigin
	this.Id = id
	return &this
}

// NewProjectAnnouncementProjectAnnouncementWithDefaults instantiates a new ProjectAnnouncementProjectAnnouncement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectAnnouncementProjectAnnouncementWithDefaults() *ProjectAnnouncementProjectAnnouncement {
	this := ProjectAnnouncementProjectAnnouncement{}
	var contentHtml string = ""
	this.ContentHtml = contentHtml
	var contentOrigin string = ""
	this.ContentOrigin = contentOrigin
	var id int64 = 0
	this.Id = id
	return &this
}

// GetContentHtml returns the ContentHtml field value
func (o *ProjectAnnouncementProjectAnnouncement) GetContentHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentHtml
}

// GetContentHtmlOk returns a tuple with the ContentHtml field value
// and a boolean to check if the value has been set.
func (o *ProjectAnnouncementProjectAnnouncement) GetContentHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentHtml, true
}

// SetContentHtml sets field value
func (o *ProjectAnnouncementProjectAnnouncement) SetContentHtml(v string) {
	o.ContentHtml = v
}

// GetContentOrigin returns the ContentOrigin field value
func (o *ProjectAnnouncementProjectAnnouncement) GetContentOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContentOrigin
}

// GetContentOriginOk returns a tuple with the ContentOrigin field value
// and a boolean to check if the value has been set.
func (o *ProjectAnnouncementProjectAnnouncement) GetContentOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentOrigin, true
}

// SetContentOrigin sets field value
func (o *ProjectAnnouncementProjectAnnouncement) SetContentOrigin(v string) {
	o.ContentOrigin = v
}

// GetId returns the Id field value
func (o *ProjectAnnouncementProjectAnnouncement) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectAnnouncementProjectAnnouncement) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectAnnouncementProjectAnnouncement) SetId(v int64) {
	o.Id = v
}

func (o ProjectAnnouncementProjectAnnouncement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectAnnouncementProjectAnnouncement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ContentHtml"] = o.ContentHtml
	toSerialize["ContentOrigin"] = o.ContentOrigin
	toSerialize["Id"] = o.Id
	return toSerialize, nil
}

func (o *ProjectAnnouncementProjectAnnouncement) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ContentHtml",
		"ContentOrigin",
		"Id",
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

	varProjectAnnouncementProjectAnnouncement := _ProjectAnnouncementProjectAnnouncement{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProjectAnnouncementProjectAnnouncement)

	if err != nil {
		return err
	}

	*o = ProjectAnnouncementProjectAnnouncement(varProjectAnnouncementProjectAnnouncement)

	return err
}

type NullableProjectAnnouncementProjectAnnouncement struct {
	value *ProjectAnnouncementProjectAnnouncement
	isSet bool
}

func (v NullableProjectAnnouncementProjectAnnouncement) Get() *ProjectAnnouncementProjectAnnouncement {
	return v.value
}

func (v *NullableProjectAnnouncementProjectAnnouncement) Set(val *ProjectAnnouncementProjectAnnouncement) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectAnnouncementProjectAnnouncement) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectAnnouncementProjectAnnouncement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectAnnouncementProjectAnnouncement(val *ProjectAnnouncementProjectAnnouncement) *NullableProjectAnnouncementProjectAnnouncement {
	return &NullableProjectAnnouncementProjectAnnouncement{value: val, isSet: true}
}

func (v NullableProjectAnnouncementProjectAnnouncement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectAnnouncementProjectAnnouncement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


