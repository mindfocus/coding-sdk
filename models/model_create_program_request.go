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

// checks if the CreateProgramRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateProgramRequest{}

// CreateProgramRequest struct for CreateProgramRequest
type CreateProgramRequest struct {
	// 描述信息
	Description string `json:"Description"`
	// 展示名
	DisplayName string `json:"DisplayName"`
	// 截止时间
	EndDate string `json:"EndDate"`
	// 项目集名
	Name string `json:"Name"`
	// 开始时间
	StartDate string `json:"StartDate"`
	// 已有工作流项目集 Id
	WorkflowProgramId *string `json:"WorkflowProgramId,omitempty"`
}

type _CreateProgramRequest CreateProgramRequest

// NewCreateProgramRequest instantiates a new CreateProgramRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProgramRequest(description string, displayName string, endDate string, name string, startDate string) *CreateProgramRequest {
	this := CreateProgramRequest{}
	this.Description = description
	this.DisplayName = displayName
	this.EndDate = endDate
	this.Name = name
	this.StartDate = startDate
	var workflowProgramId string = ""
	this.WorkflowProgramId = &workflowProgramId
	return &this
}

// NewCreateProgramRequestWithDefaults instantiates a new CreateProgramRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProgramRequestWithDefaults() *CreateProgramRequest {
	this := CreateProgramRequest{}
	var description string = ""
	this.Description = description
	var displayName string = ""
	this.DisplayName = displayName
	var endDate string = ""
	this.EndDate = endDate
	var name string = ""
	this.Name = name
	var startDate string = ""
	this.StartDate = startDate
	var workflowProgramId string = ""
	this.WorkflowProgramId = &workflowProgramId
	return &this
}

// GetDescription returns the Description field value
func (o *CreateProgramRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateProgramRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateProgramRequest) SetDescription(v string) {
	o.Description = v
}

// GetDisplayName returns the DisplayName field value
func (o *CreateProgramRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateProgramRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateProgramRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEndDate returns the EndDate field value
func (o *CreateProgramRequest) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *CreateProgramRequest) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *CreateProgramRequest) SetEndDate(v string) {
	o.EndDate = v
}

// GetName returns the Name field value
func (o *CreateProgramRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProgramRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProgramRequest) SetName(v string) {
	o.Name = v
}

// GetStartDate returns the StartDate field value
func (o *CreateProgramRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *CreateProgramRequest) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *CreateProgramRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetWorkflowProgramId returns the WorkflowProgramId field value if set, zero value otherwise.
func (o *CreateProgramRequest) GetWorkflowProgramId() string {
	if o == nil || utils.IsNil(o.WorkflowProgramId) {
		var ret string
		return ret
	}
	return *o.WorkflowProgramId
}

// GetWorkflowProgramIdOk returns a tuple with the WorkflowProgramId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProgramRequest) GetWorkflowProgramIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.WorkflowProgramId) {
		return nil, false
	}
	return o.WorkflowProgramId, true
}

// HasWorkflowProgramId returns a boolean if a field has been set.
func (o *CreateProgramRequest) HasWorkflowProgramId() bool {
	if o != nil && !utils.IsNil(o.WorkflowProgramId) {
		return true
	}

	return false
}

// SetWorkflowProgramId gets a reference to the given string and assigns it to the WorkflowProgramId field.
func (o *CreateProgramRequest) SetWorkflowProgramId(v string) {
	o.WorkflowProgramId = &v
}

func (o CreateProgramRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProgramRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Description"] = o.Description
	toSerialize["DisplayName"] = o.DisplayName
	toSerialize["EndDate"] = o.EndDate
	toSerialize["Name"] = o.Name
	toSerialize["StartDate"] = o.StartDate
	if !utils.IsNil(o.WorkflowProgramId) {
		toSerialize["WorkflowProgramId"] = o.WorkflowProgramId
	}
	return toSerialize, nil
}

func (o *CreateProgramRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Description",
		"DisplayName",
		"EndDate",
		"Name",
		"StartDate",
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

	varCreateProgramRequest := _CreateProgramRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateProgramRequest)

	if err != nil {
		return err
	}

	*o = CreateProgramRequest(varCreateProgramRequest)

	return err
}

type NullableCreateProgramRequest struct {
	value *CreateProgramRequest
	isSet bool
}

func (v NullableCreateProgramRequest) Get() *CreateProgramRequest {
	return v.value
}

func (v *NullableCreateProgramRequest) Set(val *CreateProgramRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProgramRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProgramRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProgramRequest(val *CreateProgramRequest) *NullableCreateProgramRequest {
	return &NullableCreateProgramRequest{value: val, isSet: true}
}

func (v NullableCreateProgramRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProgramRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


