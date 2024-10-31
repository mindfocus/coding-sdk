/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the StatusCheckResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StatusCheckResult{}

// StatusCheckResult 查询提交对应的流水线状态
type StatusCheckResult struct {
	// 提交sha
	CommitSha utils.NullableString `json:"CommitSha,omitempty"`
	// 流水线文本内容
	Context utils.NullableString `json:"Context,omitempty"`
	// 记录创建时间
	CreateDate utils.NullableInt64 `json:"CreateDate,omitempty"`
	// 仓库路径
	DepotId utils.NullableInt64 `json:"DepotId,omitempty"`
	// 流水线状态描述
	Description utils.NullableString `json:"Description,omitempty"`
	// 构建状态
	State utils.NullableString `json:"State,omitempty"`
	// 流水线链接地址
	TargetURL utils.NullableString `json:"TargetURL,omitempty"`
}

// NewStatusCheckResult instantiates a new StatusCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCheckResult() *StatusCheckResult {
	this := StatusCheckResult{}
	var commitSha string = ""
	this.CommitSha = *utils.NewNullableString(&commitSha)
	var context string = ""
	this.Context = *utils.NewNullableString(&context)
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var state string = ""
	this.State = *utils.NewNullableString(&state)
	var targetURL string = ""
	this.TargetURL = *utils.NewNullableString(&targetURL)
	return &this
}

// NewStatusCheckResultWithDefaults instantiates a new StatusCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCheckResultWithDefaults() *StatusCheckResult {
	this := StatusCheckResult{}
	var commitSha string = ""
	this.CommitSha = *utils.NewNullableString(&commitSha)
	var context string = ""
	this.Context = *utils.NewNullableString(&context)
	var description string = ""
	this.Description = *utils.NewNullableString(&description)
	var state string = ""
	this.State = *utils.NewNullableString(&state)
	var targetURL string = ""
	this.TargetURL = *utils.NewNullableString(&targetURL)
	return &this
}

// GetCommitSha returns the CommitSha field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetCommitSha() string {
	if o == nil || utils.IsNil(o.CommitSha.Get()) {
		var ret string
		return ret
	}
	return *o.CommitSha.Get()
}

// GetCommitShaOk returns a tuple with the CommitSha field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetCommitShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommitSha.Get(), o.CommitSha.IsSet()
}

// HasCommitSha returns a boolean if a field has been set.
func (o *StatusCheckResult) HasCommitSha() bool {
	if o != nil && o.CommitSha.IsSet() {
		return true
	}

	return false
}

// SetCommitSha gets a reference to the given utils.NullableString and assigns it to the CommitSha field.
func (o *StatusCheckResult) SetCommitSha(v string) {
	o.CommitSha.Set(&v)
}
// SetCommitShaNil sets the value for CommitSha to be an explicit nil
func (o *StatusCheckResult) SetCommitShaNil() {
	o.CommitSha.Set(nil)
}

// UnsetCommitSha ensures that no value is present for CommitSha, not even an explicit nil
func (o *StatusCheckResult) UnsetCommitSha() {
	o.CommitSha.Unset()
}

// GetContext returns the Context field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetContext() string {
	if o == nil || utils.IsNil(o.Context.Get()) {
		var ret string
		return ret
	}
	return *o.Context.Get()
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context.Get(), o.Context.IsSet()
}

// HasContext returns a boolean if a field has been set.
func (o *StatusCheckResult) HasContext() bool {
	if o != nil && o.Context.IsSet() {
		return true
	}

	return false
}

// SetContext gets a reference to the given utils.NullableString and assigns it to the Context field.
func (o *StatusCheckResult) SetContext(v string) {
	o.Context.Set(&v)
}
// SetContextNil sets the value for Context to be an explicit nil
func (o *StatusCheckResult) SetContextNil() {
	o.Context.Set(nil)
}

// UnsetContext ensures that no value is present for Context, not even an explicit nil
func (o *StatusCheckResult) UnsetContext() {
	o.Context.Unset()
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetCreateDate() int64 {
	if o == nil || utils.IsNil(o.CreateDate.Get()) {
		var ret int64
		return ret
	}
	return *o.CreateDate.Get()
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetCreateDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreateDate.Get(), o.CreateDate.IsSet()
}

// HasCreateDate returns a boolean if a field has been set.
func (o *StatusCheckResult) HasCreateDate() bool {
	if o != nil && o.CreateDate.IsSet() {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given utils.NullableInt64 and assigns it to the CreateDate field.
func (o *StatusCheckResult) SetCreateDate(v int64) {
	o.CreateDate.Set(&v)
}
// SetCreateDateNil sets the value for CreateDate to be an explicit nil
func (o *StatusCheckResult) SetCreateDateNil() {
	o.CreateDate.Set(nil)
}

// UnsetCreateDate ensures that no value is present for CreateDate, not even an explicit nil
func (o *StatusCheckResult) UnsetCreateDate() {
	o.CreateDate.Unset()
}

// GetDepotId returns the DepotId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetDepotId() int64 {
	if o == nil || utils.IsNil(o.DepotId.Get()) {
		var ret int64
		return ret
	}
	return *o.DepotId.Get()
}

// GetDepotIdOk returns a tuple with the DepotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DepotId.Get(), o.DepotId.IsSet()
}

// HasDepotId returns a boolean if a field has been set.
func (o *StatusCheckResult) HasDepotId() bool {
	if o != nil && o.DepotId.IsSet() {
		return true
	}

	return false
}

// SetDepotId gets a reference to the given utils.NullableInt64 and assigns it to the DepotId field.
func (o *StatusCheckResult) SetDepotId(v int64) {
	o.DepotId.Set(&v)
}
// SetDepotIdNil sets the value for DepotId to be an explicit nil
func (o *StatusCheckResult) SetDepotIdNil() {
	o.DepotId.Set(nil)
}

// UnsetDepotId ensures that no value is present for DepotId, not even an explicit nil
func (o *StatusCheckResult) UnsetDepotId() {
	o.DepotId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *StatusCheckResult) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given utils.NullableString and assigns it to the Description field.
func (o *StatusCheckResult) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *StatusCheckResult) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *StatusCheckResult) UnsetDescription() {
	o.Description.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetState() string {
	if o == nil || utils.IsNil(o.State.Get()) {
		var ret string
		return ret
	}
	return *o.State.Get()
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.State.Get(), o.State.IsSet()
}

// HasState returns a boolean if a field has been set.
func (o *StatusCheckResult) HasState() bool {
	if o != nil && o.State.IsSet() {
		return true
	}

	return false
}

// SetState gets a reference to the given utils.NullableString and assigns it to the State field.
func (o *StatusCheckResult) SetState(v string) {
	o.State.Set(&v)
}
// SetStateNil sets the value for State to be an explicit nil
func (o *StatusCheckResult) SetStateNil() {
	o.State.Set(nil)
}

// UnsetState ensures that no value is present for State, not even an explicit nil
func (o *StatusCheckResult) UnsetState() {
	o.State.Unset()
}

// GetTargetURL returns the TargetURL field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusCheckResult) GetTargetURL() string {
	if o == nil || utils.IsNil(o.TargetURL.Get()) {
		var ret string
		return ret
	}
	return *o.TargetURL.Get()
}

// GetTargetURLOk returns a tuple with the TargetURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StatusCheckResult) GetTargetURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetURL.Get(), o.TargetURL.IsSet()
}

// HasTargetURL returns a boolean if a field has been set.
func (o *StatusCheckResult) HasTargetURL() bool {
	if o != nil && o.TargetURL.IsSet() {
		return true
	}

	return false
}

// SetTargetURL gets a reference to the given utils.NullableString and assigns it to the TargetURL field.
func (o *StatusCheckResult) SetTargetURL(v string) {
	o.TargetURL.Set(&v)
}
// SetTargetURLNil sets the value for TargetURL to be an explicit nil
func (o *StatusCheckResult) SetTargetURLNil() {
	o.TargetURL.Set(nil)
}

// UnsetTargetURL ensures that no value is present for TargetURL, not even an explicit nil
func (o *StatusCheckResult) UnsetTargetURL() {
	o.TargetURL.Unset()
}

func (o StatusCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCheckResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CommitSha.IsSet() {
		toSerialize["CommitSha"] = o.CommitSha.Get()
	}
	if o.Context.IsSet() {
		toSerialize["Context"] = o.Context.Get()
	}
	if o.CreateDate.IsSet() {
		toSerialize["CreateDate"] = o.CreateDate.Get()
	}
	if o.DepotId.IsSet() {
		toSerialize["DepotId"] = o.DepotId.Get()
	}
	if o.Description.IsSet() {
		toSerialize["Description"] = o.Description.Get()
	}
	if o.State.IsSet() {
		toSerialize["State"] = o.State.Get()
	}
	if o.TargetURL.IsSet() {
		toSerialize["TargetURL"] = o.TargetURL.Get()
	}
	return toSerialize, nil
}

type NullableStatusCheckResult struct {
	value *StatusCheckResult
	isSet bool
}

func (v NullableStatusCheckResult) Get() *StatusCheckResult {
	return v.value
}

func (v *NullableStatusCheckResult) Set(val *StatusCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCheckResult(val *StatusCheckResult) *NullableStatusCheckResult {
	return &NullableStatusCheckResult{value: val, isSet: true}
}

func (v NullableStatusCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


