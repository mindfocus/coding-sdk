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

// checks if the CreateDepotByTemplateRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateDepotByTemplateRequest{}

// CreateDepotByTemplateRequest struct for CreateDepotByTemplateRequest
type CreateDepotByTemplateRequest struct {
	// 仓库名
	DepotName string `json:"DepotName"`
	// 仓库描述
	Description *string `json:"Description,omitempty"`
	// 项目Id
	ProjectId int64 `json:"ProjectId"`
	// 仓库模板。目前支持：Spring, Ruby on Rails，Ruby Sinatra，Node.js，Android，Python，Hexo，Jekyll。对应模板参数分别为：SPRING,ROR,SINATRA,NODEJS,ANDROID,FLASK,CLOUD_API_HEXO,CLOUD_API_JEKYLL。如果设置了自定义模版，可以传入自定义模版的仓库Id
	Template string `json:"Template"`
	// 用户Id
	UserId int64 `json:"UserId"`
}

type _CreateDepotByTemplateRequest CreateDepotByTemplateRequest

// NewCreateDepotByTemplateRequest instantiates a new CreateDepotByTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDepotByTemplateRequest(depotName string, projectId int64, template string, userId int64) *CreateDepotByTemplateRequest {
	this := CreateDepotByTemplateRequest{}
	this.DepotName = depotName
	this.ProjectId = projectId
	this.Template = template
	this.UserId = userId
	return &this
}

// NewCreateDepotByTemplateRequestWithDefaults instantiates a new CreateDepotByTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDepotByTemplateRequestWithDefaults() *CreateDepotByTemplateRequest {
	this := CreateDepotByTemplateRequest{}
	return &this
}

// GetDepotName returns the DepotName field value
func (o *CreateDepotByTemplateRequest) GetDepotName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotName
}

// GetDepotNameOk returns a tuple with the DepotName field value
// and a boolean to check if the value has been set.
func (o *CreateDepotByTemplateRequest) GetDepotNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotName, true
}

// SetDepotName sets field value
func (o *CreateDepotByTemplateRequest) SetDepotName(v string) {
	o.DepotName = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDepotByTemplateRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDepotByTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDepotByTemplateRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDepotByTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProjectId returns the ProjectId field value
func (o *CreateDepotByTemplateRequest) GetProjectId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateDepotByTemplateRequest) GetProjectIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateDepotByTemplateRequest) SetProjectId(v int64) {
	o.ProjectId = v
}

// GetTemplate returns the Template field value
func (o *CreateDepotByTemplateRequest) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CreateDepotByTemplateRequest) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CreateDepotByTemplateRequest) SetTemplate(v string) {
	o.Template = v
}

// GetUserId returns the UserId field value
func (o *CreateDepotByTemplateRequest) GetUserId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CreateDepotByTemplateRequest) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CreateDepotByTemplateRequest) SetUserId(v int64) {
	o.UserId = v
}

func (o CreateDepotByTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDepotByTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotName"] = o.DepotName
	if !utils.IsNil(o.Description) {
		toSerialize["Description"] = o.Description
	}
	toSerialize["ProjectId"] = o.ProjectId
	toSerialize["Template"] = o.Template
	toSerialize["UserId"] = o.UserId
	return toSerialize, nil
}

func (o *CreateDepotByTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotName",
		"ProjectId",
		"Template",
		"UserId",
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

	varCreateDepotByTemplateRequest := _CreateDepotByTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDepotByTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateDepotByTemplateRequest(varCreateDepotByTemplateRequest)

	return err
}

type NullableCreateDepotByTemplateRequest struct {
	value *CreateDepotByTemplateRequest
	isSet bool
}

func (v NullableCreateDepotByTemplateRequest) Get() *CreateDepotByTemplateRequest {
	return v.value
}

func (v *NullableCreateDepotByTemplateRequest) Set(val *CreateDepotByTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDepotByTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDepotByTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDepotByTemplateRequest(val *CreateDepotByTemplateRequest) *NullableCreateDepotByTemplateRequest {
	return &NullableCreateDepotByTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateDepotByTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDepotByTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

