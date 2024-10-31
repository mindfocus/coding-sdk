/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the Iteration type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Iteration{}

// Iteration 迭代信息
type Iteration struct {
	// 处理人 ID ，为 0 代表没有设置
	Assignee *int64 `json:"Assignee,omitempty"`
	// 迭代编号，项目内唯一
	Code *int64 `json:"Code,omitempty"`
	// 迭代中完成事项总数
	CompletedCount *int64 `json:"CompletedCount,omitempty"`
	// 迭代中事项完成比率
	CompletedPercent *float32 `json:"CompletedPercent,omitempty"`
	// 完成人 ID
	Completer *int64 `json:"Completer,omitempty"`
	// 创建时间
	CreatedAt utils.NullableInt64 `json:"CreatedAt,omitempty"`
	// 创建人 ID
	Creator *int64 `json:"Creator,omitempty"`
	// 结束时间，时间戳，-28800000 代表没有设置
	EndAt utils.NullableInt64 `json:"EndAt,omitempty"`
	// 迭代目标
	Goal utils.NullableString `json:"Goal,omitempty"`
	// 标题
	Name *string `json:"Name,omitempty"`
	// 迭代中进行中事项总数
	ProcessingCount *int64 `json:"ProcessingCount,omitempty"`
	// 开始时间，时间戳，-28800000 代表没有设置
	StartAt utils.NullableInt64 `json:"StartAt,omitempty"`
	// 开始人 ID
	Starter *int64 `json:"Starter,omitempty"`
	// 迭代状态：WAIT_PROCESS,PROCESSING,COMPLETED
	Status *string `json:"Status,omitempty"`
	// 修改时间
	UpdatedAt utils.NullableInt64 `json:"UpdatedAt,omitempty"`
	// 迭代中待处理事项总数
	WaitProcessCount *int64 `json:"WaitProcessCount,omitempty"`
}

// NewIteration instantiates a new Iteration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIteration() *Iteration {
	this := Iteration{}
	var goal string = ""
	this.Goal = *utils.NewNullableString(&goal)
	var name string = ""
	this.Name = &name
	var status string = ""
	this.Status = &status
	return &this
}

// NewIterationWithDefaults instantiates a new Iteration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIterationWithDefaults() *Iteration {
	this := Iteration{}
	var goal string = ""
	this.Goal = *utils.NewNullableString(&goal)
	var name string = ""
	this.Name = &name
	var status string = ""
	this.Status = &status
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *Iteration) GetAssignee() int64 {
	if o == nil || utils.IsNil(o.Assignee) {
		var ret int64
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetAssigneeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *Iteration) HasAssignee() bool {
	if o != nil && !utils.IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given int64 and assigns it to the Assignee field.
func (o *Iteration) SetAssignee(v int64) {
	o.Assignee = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Iteration) GetCode() int64 {
	if o == nil || utils.IsNil(o.Code) {
		var ret int64
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetCodeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Iteration) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int64 and assigns it to the Code field.
func (o *Iteration) SetCode(v int64) {
	o.Code = &v
}

// GetCompletedCount returns the CompletedCount field value if set, zero value otherwise.
func (o *Iteration) GetCompletedCount() int64 {
	if o == nil || utils.IsNil(o.CompletedCount) {
		var ret int64
		return ret
	}
	return *o.CompletedCount
}

// GetCompletedCountOk returns a tuple with the CompletedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetCompletedCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.CompletedCount) {
		return nil, false
	}
	return o.CompletedCount, true
}

// HasCompletedCount returns a boolean if a field has been set.
func (o *Iteration) HasCompletedCount() bool {
	if o != nil && !utils.IsNil(o.CompletedCount) {
		return true
	}

	return false
}

// SetCompletedCount gets a reference to the given int64 and assigns it to the CompletedCount field.
func (o *Iteration) SetCompletedCount(v int64) {
	o.CompletedCount = &v
}

// GetCompletedPercent returns the CompletedPercent field value if set, zero value otherwise.
func (o *Iteration) GetCompletedPercent() float32 {
	if o == nil || utils.IsNil(o.CompletedPercent) {
		var ret float32
		return ret
	}
	return *o.CompletedPercent
}

// GetCompletedPercentOk returns a tuple with the CompletedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetCompletedPercentOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.CompletedPercent) {
		return nil, false
	}
	return o.CompletedPercent, true
}

// HasCompletedPercent returns a boolean if a field has been set.
func (o *Iteration) HasCompletedPercent() bool {
	if o != nil && !utils.IsNil(o.CompletedPercent) {
		return true
	}

	return false
}

// SetCompletedPercent gets a reference to the given float32 and assigns it to the CompletedPercent field.
func (o *Iteration) SetCompletedPercent(v float32) {
	o.CompletedPercent = &v
}

// GetCompleter returns the Completer field value if set, zero value otherwise.
func (o *Iteration) GetCompleter() int64 {
	if o == nil || utils.IsNil(o.Completer) {
		var ret int64
		return ret
	}
	return *o.Completer
}

// GetCompleterOk returns a tuple with the Completer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetCompleterOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Completer) {
		return nil, false
	}
	return o.Completer, true
}

// HasCompleter returns a boolean if a field has been set.
func (o *Iteration) HasCompleter() bool {
	if o != nil && !utils.IsNil(o.Completer) {
		return true
	}

	return false
}

// SetCompleter gets a reference to the given int64 and assigns it to the Completer field.
func (o *Iteration) SetCompleter(v int64) {
	o.Completer = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Iteration) GetCreatedAt() int64 {
	if o == nil || utils.IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Iteration) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Iteration) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given utils.NullableInt64 and assigns it to the CreatedAt field.
func (o *Iteration) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Iteration) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Iteration) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *Iteration) GetCreator() int64 {
	if o == nil || utils.IsNil(o.Creator) {
		var ret int64
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetCreatorOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *Iteration) HasCreator() bool {
	if o != nil && !utils.IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given int64 and assigns it to the Creator field.
func (o *Iteration) SetCreator(v int64) {
	o.Creator = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Iteration) GetEndAt() int64 {
	if o == nil || utils.IsNil(o.EndAt.Get()) {
		var ret int64
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Iteration) GetEndAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *Iteration) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given utils.NullableInt64 and assigns it to the EndAt field.
func (o *Iteration) SetEndAt(v int64) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *Iteration) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *Iteration) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetGoal returns the Goal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Iteration) GetGoal() string {
	if o == nil || utils.IsNil(o.Goal.Get()) {
		var ret string
		return ret
	}
	return *o.Goal.Get()
}

// GetGoalOk returns a tuple with the Goal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Iteration) GetGoalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Goal.Get(), o.Goal.IsSet()
}

// HasGoal returns a boolean if a field has been set.
func (o *Iteration) HasGoal() bool {
	if o != nil && o.Goal.IsSet() {
		return true
	}

	return false
}

// SetGoal gets a reference to the given utils.NullableString and assigns it to the Goal field.
func (o *Iteration) SetGoal(v string) {
	o.Goal.Set(&v)
}
// SetGoalNil sets the value for Goal to be an explicit nil
func (o *Iteration) SetGoalNil() {
	o.Goal.Set(nil)
}

// UnsetGoal ensures that no value is present for Goal, not even an explicit nil
func (o *Iteration) UnsetGoal() {
	o.Goal.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Iteration) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Iteration) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Iteration) SetName(v string) {
	o.Name = &v
}

// GetProcessingCount returns the ProcessingCount field value if set, zero value otherwise.
func (o *Iteration) GetProcessingCount() int64 {
	if o == nil || utils.IsNil(o.ProcessingCount) {
		var ret int64
		return ret
	}
	return *o.ProcessingCount
}

// GetProcessingCountOk returns a tuple with the ProcessingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetProcessingCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ProcessingCount) {
		return nil, false
	}
	return o.ProcessingCount, true
}

// HasProcessingCount returns a boolean if a field has been set.
func (o *Iteration) HasProcessingCount() bool {
	if o != nil && !utils.IsNil(o.ProcessingCount) {
		return true
	}

	return false
}

// SetProcessingCount gets a reference to the given int64 and assigns it to the ProcessingCount field.
func (o *Iteration) SetProcessingCount(v int64) {
	o.ProcessingCount = &v
}

// GetStartAt returns the StartAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Iteration) GetStartAt() int64 {
	if o == nil || utils.IsNil(o.StartAt.Get()) {
		var ret int64
		return ret
	}
	return *o.StartAt.Get()
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Iteration) GetStartAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartAt.Get(), o.StartAt.IsSet()
}

// HasStartAt returns a boolean if a field has been set.
func (o *Iteration) HasStartAt() bool {
	if o != nil && o.StartAt.IsSet() {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given utils.NullableInt64 and assigns it to the StartAt field.
func (o *Iteration) SetStartAt(v int64) {
	o.StartAt.Set(&v)
}
// SetStartAtNil sets the value for StartAt to be an explicit nil
func (o *Iteration) SetStartAtNil() {
	o.StartAt.Set(nil)
}

// UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
func (o *Iteration) UnsetStartAt() {
	o.StartAt.Unset()
}

// GetStarter returns the Starter field value if set, zero value otherwise.
func (o *Iteration) GetStarter() int64 {
	if o == nil || utils.IsNil(o.Starter) {
		var ret int64
		return ret
	}
	return *o.Starter
}

// GetStarterOk returns a tuple with the Starter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetStarterOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Starter) {
		return nil, false
	}
	return o.Starter, true
}

// HasStarter returns a boolean if a field has been set.
func (o *Iteration) HasStarter() bool {
	if o != nil && !utils.IsNil(o.Starter) {
		return true
	}

	return false
}

// SetStarter gets a reference to the given int64 and assigns it to the Starter field.
func (o *Iteration) SetStarter(v int64) {
	o.Starter = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Iteration) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Iteration) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Iteration) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Iteration) GetUpdatedAt() int64 {
	if o == nil || utils.IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Iteration) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Iteration) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given utils.NullableInt64 and assigns it to the UpdatedAt field.
func (o *Iteration) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Iteration) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Iteration) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetWaitProcessCount returns the WaitProcessCount field value if set, zero value otherwise.
func (o *Iteration) GetWaitProcessCount() int64 {
	if o == nil || utils.IsNil(o.WaitProcessCount) {
		var ret int64
		return ret
	}
	return *o.WaitProcessCount
}

// GetWaitProcessCountOk returns a tuple with the WaitProcessCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Iteration) GetWaitProcessCountOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.WaitProcessCount) {
		return nil, false
	}
	return o.WaitProcessCount, true
}

// HasWaitProcessCount returns a boolean if a field has been set.
func (o *Iteration) HasWaitProcessCount() bool {
	if o != nil && !utils.IsNil(o.WaitProcessCount) {
		return true
	}

	return false
}

// SetWaitProcessCount gets a reference to the given int64 and assigns it to the WaitProcessCount field.
func (o *Iteration) SetWaitProcessCount(v int64) {
	o.WaitProcessCount = &v
}

func (o Iteration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Iteration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Assignee) {
		toSerialize["Assignee"] = o.Assignee
	}
	if !utils.IsNil(o.Code) {
		toSerialize["Code"] = o.Code
	}
	if !utils.IsNil(o.CompletedCount) {
		toSerialize["CompletedCount"] = o.CompletedCount
	}
	if !utils.IsNil(o.CompletedPercent) {
		toSerialize["CompletedPercent"] = o.CompletedPercent
	}
	if !utils.IsNil(o.Completer) {
		toSerialize["Completer"] = o.Completer
	}
	if o.CreatedAt.IsSet() {
		toSerialize["CreatedAt"] = o.CreatedAt.Get()
	}
	if !utils.IsNil(o.Creator) {
		toSerialize["Creator"] = o.Creator
	}
	if o.EndAt.IsSet() {
		toSerialize["EndAt"] = o.EndAt.Get()
	}
	if o.Goal.IsSet() {
		toSerialize["Goal"] = o.Goal.Get()
	}
	if !utils.IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !utils.IsNil(o.ProcessingCount) {
		toSerialize["ProcessingCount"] = o.ProcessingCount
	}
	if o.StartAt.IsSet() {
		toSerialize["StartAt"] = o.StartAt.Get()
	}
	if !utils.IsNil(o.Starter) {
		toSerialize["Starter"] = o.Starter
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["UpdatedAt"] = o.UpdatedAt.Get()
	}
	if !utils.IsNil(o.WaitProcessCount) {
		toSerialize["WaitProcessCount"] = o.WaitProcessCount
	}
	return toSerialize, nil
}

type NullableIteration struct {
	value *Iteration
	isSet bool
}

func (v NullableIteration) Get() *Iteration {
	return v.value
}

func (v *NullableIteration) Set(val *Iteration) {
	v.value = val
	v.isSet = true
}

func (v NullableIteration) IsSet() bool {
	return v.isSet
}

func (v *NullableIteration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIteration(val *Iteration) *NullableIteration {
	return &NullableIteration{value: val, isSet: true}
}

func (v NullableIteration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIteration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


