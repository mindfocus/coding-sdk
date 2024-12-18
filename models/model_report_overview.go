/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ReportOverview type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReportOverview{}

// ReportOverview 测试报告概览
type ReportOverview struct {
	// 自动化覆盖率：百分比
	AutomationPercent utils.NullableInt32 `json:"AutomationPercent,omitempty"`
	// 平均关闭时长
	AvgClosedSeconds utils.NullableInt64 `json:"AvgClosedSeconds,omitempty"`
	// 用例总数
	CaseSum utils.NullableInt32 `json:"CaseSum,omitempty"`
	// 已完成数量
	CompletedSum utils.NullableInt32 `json:"CompletedSum,omitempty"`
	// 缺陷修复率：百分比
	DefectFixPercent utils.NullableInt32 `json:"DefectFixPercent,omitempty"`
	// 重新激活率：百分比
	DefectReopenPercent utils.NullableInt32 `json:"DefectReopenPercent,omitempty"`
	// 缺陷总数
	DefectSum utils.NullableInt32 `json:"DefectSum,omitempty"`
	// 85%解决时长
	DurationFixed utils.NullableInt64 `json:"DurationFixed,omitempty"`
	// 执行率：百分比
	ExecPercent utils.NullableInt32 `json:"ExecPercent,omitempty"`
	// 需求总数
	IssuesSum utils.NullableInt32 `json:"IssuesSum,omitempty"`
	// 通过率：百分比
	PassPercent utils.NullableInt32 `json:"PassPercent,omitempty"`
	// 处理中数量
	ProcessingSum utils.NullableInt32 `json:"ProcessingSum,omitempty"`
	// 需求覆盖率：百分比
	RequirementCoverPercent utils.NullableInt32 `json:"RequirementCoverPercent,omitempty"`
	// 未开始数量
	TodoSum utils.NullableInt32 `json:"TodoSum,omitempty"`
}

// NewReportOverview instantiates a new ReportOverview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportOverview() *ReportOverview {
	this := ReportOverview{}
	return &this
}

// NewReportOverviewWithDefaults instantiates a new ReportOverview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportOverviewWithDefaults() *ReportOverview {
	this := ReportOverview{}
	return &this
}

// GetAutomationPercent returns the AutomationPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetAutomationPercent() int32 {
	if o == nil || utils.IsNil(o.AutomationPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.AutomationPercent.Get()
}

// GetAutomationPercentOk returns a tuple with the AutomationPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetAutomationPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AutomationPercent.Get(), o.AutomationPercent.IsSet()
}

// HasAutomationPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasAutomationPercent() bool {
	if o != nil && o.AutomationPercent.IsSet() {
		return true
	}

	return false
}

// SetAutomationPercent gets a reference to the given utils.NullableInt32 and assigns it to the AutomationPercent field.
func (o *ReportOverview) SetAutomationPercent(v int32) {
	o.AutomationPercent.Set(&v)
}
// SetAutomationPercentNil sets the value for AutomationPercent to be an explicit nil
func (o *ReportOverview) SetAutomationPercentNil() {
	o.AutomationPercent.Set(nil)
}

// UnsetAutomationPercent ensures that no value is present for AutomationPercent, not even an explicit nil
func (o *ReportOverview) UnsetAutomationPercent() {
	o.AutomationPercent.Unset()
}

// GetAvgClosedSeconds returns the AvgClosedSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetAvgClosedSeconds() int64 {
	if o == nil || utils.IsNil(o.AvgClosedSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.AvgClosedSeconds.Get()
}

// GetAvgClosedSecondsOk returns a tuple with the AvgClosedSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetAvgClosedSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvgClosedSeconds.Get(), o.AvgClosedSeconds.IsSet()
}

// HasAvgClosedSeconds returns a boolean if a field has been set.
func (o *ReportOverview) HasAvgClosedSeconds() bool {
	if o != nil && o.AvgClosedSeconds.IsSet() {
		return true
	}

	return false
}

// SetAvgClosedSeconds gets a reference to the given utils.NullableInt64 and assigns it to the AvgClosedSeconds field.
func (o *ReportOverview) SetAvgClosedSeconds(v int64) {
	o.AvgClosedSeconds.Set(&v)
}
// SetAvgClosedSecondsNil sets the value for AvgClosedSeconds to be an explicit nil
func (o *ReportOverview) SetAvgClosedSecondsNil() {
	o.AvgClosedSeconds.Set(nil)
}

// UnsetAvgClosedSeconds ensures that no value is present for AvgClosedSeconds, not even an explicit nil
func (o *ReportOverview) UnsetAvgClosedSeconds() {
	o.AvgClosedSeconds.Unset()
}

// GetCaseSum returns the CaseSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetCaseSum() int32 {
	if o == nil || utils.IsNil(o.CaseSum.Get()) {
		var ret int32
		return ret
	}
	return *o.CaseSum.Get()
}

// GetCaseSumOk returns a tuple with the CaseSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetCaseSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaseSum.Get(), o.CaseSum.IsSet()
}

// HasCaseSum returns a boolean if a field has been set.
func (o *ReportOverview) HasCaseSum() bool {
	if o != nil && o.CaseSum.IsSet() {
		return true
	}

	return false
}

// SetCaseSum gets a reference to the given utils.NullableInt32 and assigns it to the CaseSum field.
func (o *ReportOverview) SetCaseSum(v int32) {
	o.CaseSum.Set(&v)
}
// SetCaseSumNil sets the value for CaseSum to be an explicit nil
func (o *ReportOverview) SetCaseSumNil() {
	o.CaseSum.Set(nil)
}

// UnsetCaseSum ensures that no value is present for CaseSum, not even an explicit nil
func (o *ReportOverview) UnsetCaseSum() {
	o.CaseSum.Unset()
}

// GetCompletedSum returns the CompletedSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetCompletedSum() int32 {
	if o == nil || utils.IsNil(o.CompletedSum.Get()) {
		var ret int32
		return ret
	}
	return *o.CompletedSum.Get()
}

// GetCompletedSumOk returns a tuple with the CompletedSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetCompletedSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedSum.Get(), o.CompletedSum.IsSet()
}

// HasCompletedSum returns a boolean if a field has been set.
func (o *ReportOverview) HasCompletedSum() bool {
	if o != nil && o.CompletedSum.IsSet() {
		return true
	}

	return false
}

// SetCompletedSum gets a reference to the given utils.NullableInt32 and assigns it to the CompletedSum field.
func (o *ReportOverview) SetCompletedSum(v int32) {
	o.CompletedSum.Set(&v)
}
// SetCompletedSumNil sets the value for CompletedSum to be an explicit nil
func (o *ReportOverview) SetCompletedSumNil() {
	o.CompletedSum.Set(nil)
}

// UnsetCompletedSum ensures that no value is present for CompletedSum, not even an explicit nil
func (o *ReportOverview) UnsetCompletedSum() {
	o.CompletedSum.Unset()
}

// GetDefectFixPercent returns the DefectFixPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDefectFixPercent() int32 {
	if o == nil || utils.IsNil(o.DefectFixPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.DefectFixPercent.Get()
}

// GetDefectFixPercentOk returns a tuple with the DefectFixPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDefectFixPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefectFixPercent.Get(), o.DefectFixPercent.IsSet()
}

// HasDefectFixPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasDefectFixPercent() bool {
	if o != nil && o.DefectFixPercent.IsSet() {
		return true
	}

	return false
}

// SetDefectFixPercent gets a reference to the given utils.NullableInt32 and assigns it to the DefectFixPercent field.
func (o *ReportOverview) SetDefectFixPercent(v int32) {
	o.DefectFixPercent.Set(&v)
}
// SetDefectFixPercentNil sets the value for DefectFixPercent to be an explicit nil
func (o *ReportOverview) SetDefectFixPercentNil() {
	o.DefectFixPercent.Set(nil)
}

// UnsetDefectFixPercent ensures that no value is present for DefectFixPercent, not even an explicit nil
func (o *ReportOverview) UnsetDefectFixPercent() {
	o.DefectFixPercent.Unset()
}

// GetDefectReopenPercent returns the DefectReopenPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDefectReopenPercent() int32 {
	if o == nil || utils.IsNil(o.DefectReopenPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.DefectReopenPercent.Get()
}

// GetDefectReopenPercentOk returns a tuple with the DefectReopenPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDefectReopenPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefectReopenPercent.Get(), o.DefectReopenPercent.IsSet()
}

// HasDefectReopenPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasDefectReopenPercent() bool {
	if o != nil && o.DefectReopenPercent.IsSet() {
		return true
	}

	return false
}

// SetDefectReopenPercent gets a reference to the given utils.NullableInt32 and assigns it to the DefectReopenPercent field.
func (o *ReportOverview) SetDefectReopenPercent(v int32) {
	o.DefectReopenPercent.Set(&v)
}
// SetDefectReopenPercentNil sets the value for DefectReopenPercent to be an explicit nil
func (o *ReportOverview) SetDefectReopenPercentNil() {
	o.DefectReopenPercent.Set(nil)
}

// UnsetDefectReopenPercent ensures that no value is present for DefectReopenPercent, not even an explicit nil
func (o *ReportOverview) UnsetDefectReopenPercent() {
	o.DefectReopenPercent.Unset()
}

// GetDefectSum returns the DefectSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDefectSum() int32 {
	if o == nil || utils.IsNil(o.DefectSum.Get()) {
		var ret int32
		return ret
	}
	return *o.DefectSum.Get()
}

// GetDefectSumOk returns a tuple with the DefectSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDefectSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefectSum.Get(), o.DefectSum.IsSet()
}

// HasDefectSum returns a boolean if a field has been set.
func (o *ReportOverview) HasDefectSum() bool {
	if o != nil && o.DefectSum.IsSet() {
		return true
	}

	return false
}

// SetDefectSum gets a reference to the given utils.NullableInt32 and assigns it to the DefectSum field.
func (o *ReportOverview) SetDefectSum(v int32) {
	o.DefectSum.Set(&v)
}
// SetDefectSumNil sets the value for DefectSum to be an explicit nil
func (o *ReportOverview) SetDefectSumNil() {
	o.DefectSum.Set(nil)
}

// UnsetDefectSum ensures that no value is present for DefectSum, not even an explicit nil
func (o *ReportOverview) UnsetDefectSum() {
	o.DefectSum.Unset()
}

// GetDurationFixed returns the DurationFixed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetDurationFixed() int64 {
	if o == nil || utils.IsNil(o.DurationFixed.Get()) {
		var ret int64
		return ret
	}
	return *o.DurationFixed.Get()
}

// GetDurationFixedOk returns a tuple with the DurationFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetDurationFixedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationFixed.Get(), o.DurationFixed.IsSet()
}

// HasDurationFixed returns a boolean if a field has been set.
func (o *ReportOverview) HasDurationFixed() bool {
	if o != nil && o.DurationFixed.IsSet() {
		return true
	}

	return false
}

// SetDurationFixed gets a reference to the given utils.NullableInt64 and assigns it to the DurationFixed field.
func (o *ReportOverview) SetDurationFixed(v int64) {
	o.DurationFixed.Set(&v)
}
// SetDurationFixedNil sets the value for DurationFixed to be an explicit nil
func (o *ReportOverview) SetDurationFixedNil() {
	o.DurationFixed.Set(nil)
}

// UnsetDurationFixed ensures that no value is present for DurationFixed, not even an explicit nil
func (o *ReportOverview) UnsetDurationFixed() {
	o.DurationFixed.Unset()
}

// GetExecPercent returns the ExecPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetExecPercent() int32 {
	if o == nil || utils.IsNil(o.ExecPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.ExecPercent.Get()
}

// GetExecPercentOk returns a tuple with the ExecPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetExecPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExecPercent.Get(), o.ExecPercent.IsSet()
}

// HasExecPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasExecPercent() bool {
	if o != nil && o.ExecPercent.IsSet() {
		return true
	}

	return false
}

// SetExecPercent gets a reference to the given utils.NullableInt32 and assigns it to the ExecPercent field.
func (o *ReportOverview) SetExecPercent(v int32) {
	o.ExecPercent.Set(&v)
}
// SetExecPercentNil sets the value for ExecPercent to be an explicit nil
func (o *ReportOverview) SetExecPercentNil() {
	o.ExecPercent.Set(nil)
}

// UnsetExecPercent ensures that no value is present for ExecPercent, not even an explicit nil
func (o *ReportOverview) UnsetExecPercent() {
	o.ExecPercent.Unset()
}

// GetIssuesSum returns the IssuesSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetIssuesSum() int32 {
	if o == nil || utils.IsNil(o.IssuesSum.Get()) {
		var ret int32
		return ret
	}
	return *o.IssuesSum.Get()
}

// GetIssuesSumOk returns a tuple with the IssuesSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetIssuesSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuesSum.Get(), o.IssuesSum.IsSet()
}

// HasIssuesSum returns a boolean if a field has been set.
func (o *ReportOverview) HasIssuesSum() bool {
	if o != nil && o.IssuesSum.IsSet() {
		return true
	}

	return false
}

// SetIssuesSum gets a reference to the given utils.NullableInt32 and assigns it to the IssuesSum field.
func (o *ReportOverview) SetIssuesSum(v int32) {
	o.IssuesSum.Set(&v)
}
// SetIssuesSumNil sets the value for IssuesSum to be an explicit nil
func (o *ReportOverview) SetIssuesSumNil() {
	o.IssuesSum.Set(nil)
}

// UnsetIssuesSum ensures that no value is present for IssuesSum, not even an explicit nil
func (o *ReportOverview) UnsetIssuesSum() {
	o.IssuesSum.Unset()
}

// GetPassPercent returns the PassPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetPassPercent() int32 {
	if o == nil || utils.IsNil(o.PassPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.PassPercent.Get()
}

// GetPassPercentOk returns a tuple with the PassPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetPassPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PassPercent.Get(), o.PassPercent.IsSet()
}

// HasPassPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasPassPercent() bool {
	if o != nil && o.PassPercent.IsSet() {
		return true
	}

	return false
}

// SetPassPercent gets a reference to the given utils.NullableInt32 and assigns it to the PassPercent field.
func (o *ReportOverview) SetPassPercent(v int32) {
	o.PassPercent.Set(&v)
}
// SetPassPercentNil sets the value for PassPercent to be an explicit nil
func (o *ReportOverview) SetPassPercentNil() {
	o.PassPercent.Set(nil)
}

// UnsetPassPercent ensures that no value is present for PassPercent, not even an explicit nil
func (o *ReportOverview) UnsetPassPercent() {
	o.PassPercent.Unset()
}

// GetProcessingSum returns the ProcessingSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetProcessingSum() int32 {
	if o == nil || utils.IsNil(o.ProcessingSum.Get()) {
		var ret int32
		return ret
	}
	return *o.ProcessingSum.Get()
}

// GetProcessingSumOk returns a tuple with the ProcessingSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetProcessingSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProcessingSum.Get(), o.ProcessingSum.IsSet()
}

// HasProcessingSum returns a boolean if a field has been set.
func (o *ReportOverview) HasProcessingSum() bool {
	if o != nil && o.ProcessingSum.IsSet() {
		return true
	}

	return false
}

// SetProcessingSum gets a reference to the given utils.NullableInt32 and assigns it to the ProcessingSum field.
func (o *ReportOverview) SetProcessingSum(v int32) {
	o.ProcessingSum.Set(&v)
}
// SetProcessingSumNil sets the value for ProcessingSum to be an explicit nil
func (o *ReportOverview) SetProcessingSumNil() {
	o.ProcessingSum.Set(nil)
}

// UnsetProcessingSum ensures that no value is present for ProcessingSum, not even an explicit nil
func (o *ReportOverview) UnsetProcessingSum() {
	o.ProcessingSum.Unset()
}

// GetRequirementCoverPercent returns the RequirementCoverPercent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetRequirementCoverPercent() int32 {
	if o == nil || utils.IsNil(o.RequirementCoverPercent.Get()) {
		var ret int32
		return ret
	}
	return *o.RequirementCoverPercent.Get()
}

// GetRequirementCoverPercentOk returns a tuple with the RequirementCoverPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetRequirementCoverPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequirementCoverPercent.Get(), o.RequirementCoverPercent.IsSet()
}

// HasRequirementCoverPercent returns a boolean if a field has been set.
func (o *ReportOverview) HasRequirementCoverPercent() bool {
	if o != nil && o.RequirementCoverPercent.IsSet() {
		return true
	}

	return false
}

// SetRequirementCoverPercent gets a reference to the given utils.NullableInt32 and assigns it to the RequirementCoverPercent field.
func (o *ReportOverview) SetRequirementCoverPercent(v int32) {
	o.RequirementCoverPercent.Set(&v)
}
// SetRequirementCoverPercentNil sets the value for RequirementCoverPercent to be an explicit nil
func (o *ReportOverview) SetRequirementCoverPercentNil() {
	o.RequirementCoverPercent.Set(nil)
}

// UnsetRequirementCoverPercent ensures that no value is present for RequirementCoverPercent, not even an explicit nil
func (o *ReportOverview) UnsetRequirementCoverPercent() {
	o.RequirementCoverPercent.Unset()
}

// GetTodoSum returns the TodoSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportOverview) GetTodoSum() int32 {
	if o == nil || utils.IsNil(o.TodoSum.Get()) {
		var ret int32
		return ret
	}
	return *o.TodoSum.Get()
}

// GetTodoSumOk returns a tuple with the TodoSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportOverview) GetTodoSumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TodoSum.Get(), o.TodoSum.IsSet()
}

// HasTodoSum returns a boolean if a field has been set.
func (o *ReportOverview) HasTodoSum() bool {
	if o != nil && o.TodoSum.IsSet() {
		return true
	}

	return false
}

// SetTodoSum gets a reference to the given utils.NullableInt32 and assigns it to the TodoSum field.
func (o *ReportOverview) SetTodoSum(v int32) {
	o.TodoSum.Set(&v)
}
// SetTodoSumNil sets the value for TodoSum to be an explicit nil
func (o *ReportOverview) SetTodoSumNil() {
	o.TodoSum.Set(nil)
}

// UnsetTodoSum ensures that no value is present for TodoSum, not even an explicit nil
func (o *ReportOverview) UnsetTodoSum() {
	o.TodoSum.Unset()
}

func (o ReportOverview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportOverview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AutomationPercent.IsSet() {
		toSerialize["AutomationPercent"] = o.AutomationPercent.Get()
	}
	if o.AvgClosedSeconds.IsSet() {
		toSerialize["AvgClosedSeconds"] = o.AvgClosedSeconds.Get()
	}
	if o.CaseSum.IsSet() {
		toSerialize["CaseSum"] = o.CaseSum.Get()
	}
	if o.CompletedSum.IsSet() {
		toSerialize["CompletedSum"] = o.CompletedSum.Get()
	}
	if o.DefectFixPercent.IsSet() {
		toSerialize["DefectFixPercent"] = o.DefectFixPercent.Get()
	}
	if o.DefectReopenPercent.IsSet() {
		toSerialize["DefectReopenPercent"] = o.DefectReopenPercent.Get()
	}
	if o.DefectSum.IsSet() {
		toSerialize["DefectSum"] = o.DefectSum.Get()
	}
	if o.DurationFixed.IsSet() {
		toSerialize["DurationFixed"] = o.DurationFixed.Get()
	}
	if o.ExecPercent.IsSet() {
		toSerialize["ExecPercent"] = o.ExecPercent.Get()
	}
	if o.IssuesSum.IsSet() {
		toSerialize["IssuesSum"] = o.IssuesSum.Get()
	}
	if o.PassPercent.IsSet() {
		toSerialize["PassPercent"] = o.PassPercent.Get()
	}
	if o.ProcessingSum.IsSet() {
		toSerialize["ProcessingSum"] = o.ProcessingSum.Get()
	}
	if o.RequirementCoverPercent.IsSet() {
		toSerialize["RequirementCoverPercent"] = o.RequirementCoverPercent.Get()
	}
	if o.TodoSum.IsSet() {
		toSerialize["TodoSum"] = o.TodoSum.Get()
	}
	return toSerialize, nil
}

type NullableReportOverview struct {
	value *ReportOverview
	isSet bool
}

func (v NullableReportOverview) Get() *ReportOverview {
	return v.value
}

func (v *NullableReportOverview) Set(val *ReportOverview) {
	v.value = val
	v.isSet = true
}

func (v NullableReportOverview) IsSet() bool {
	return v.isSet
}

func (v *NullableReportOverview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportOverview(val *ReportOverview) *NullableReportOverview {
	return &NullableReportOverview{value: val, isSet: true}
}

func (v NullableReportOverview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportOverview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


