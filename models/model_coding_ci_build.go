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

// checks if the CodingCIBuild type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CodingCIBuild{}

// CodingCIBuild CodingCiBuild 结构
type CodingCIBuild struct {
	// 分支
	Branch string `json:"Branch"`
	// 触发原因
	Cause string `json:"Cause"`
	// 构建唯一标识
	CodingCIId string `json:"CodingCIId"`
	// Git Commit ID
	CommitId string `json:"CommitId"`
	// 构建创建时间
	CreatedAt int32 `json:"CreatedAt"`
	// 构建执行时间 QUEUED  等待构建 INITIALIZING  初始化 NOT_BUILT  准备构建 RUNNING  运行中 SUCCEED  成功 FAILED  失败 ABORTED  被取消 TIMEOUT  超时
	Duration int32 `json:"Duration"`
	// 构建使用的环境变量
	EnvList []CIJobEnv `json:"EnvList,omitempty"`
	// 失败原因
	FailedMessage string `json:"FailedMessage"`
	// 构建 ID
	Id int32 `json:"Id"`
	// 本次构建的 Jenkinsfile
	JenkinsFileContent string `json:"JenkinsFileContent"`
	// 构建计划 ID
	JobId int32 `json:"JobId"`
	// 获取到执行机的时间，如果为负数表示还未获取到构建节点
	NodeObtainedAt utils.NullableInt64 `json:"NodeObtainedAt,omitempty"`
	// 构建序号
	Number int32 `json:"Number"`
	// 开始构建时间，如果为负数就是默认值表示还未开始
	StartedAt utils.NullableInt64 `json:"StartedAt,omitempty"`
	// 构建当前状态
	Status string `json:"Status"`
	// 构建进行到了哪个 stage/node
	StatusNode string `json:"StatusNode"`
	TestResult CIBuildTestResult `json:"TestResult"`
}

type _CodingCIBuild CodingCIBuild

// NewCodingCIBuild instantiates a new CodingCIBuild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodingCIBuild(branch string, cause string, codingCIId string, commitId string, createdAt int32, duration int32, failedMessage string, id int32, jenkinsFileContent string, jobId int32, number int32, status string, statusNode string, testResult CIBuildTestResult) *CodingCIBuild {
	this := CodingCIBuild{}
	this.Branch = branch
	this.Cause = cause
	this.CodingCIId = codingCIId
	this.CommitId = commitId
	this.CreatedAt = createdAt
	this.Duration = duration
	this.FailedMessage = failedMessage
	this.Id = id
	this.JenkinsFileContent = jenkinsFileContent
	this.JobId = jobId
	this.Number = number
	this.Status = status
	this.StatusNode = statusNode
	this.TestResult = testResult
	return &this
}

// NewCodingCIBuildWithDefaults instantiates a new CodingCIBuild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodingCIBuildWithDefaults() *CodingCIBuild {
	this := CodingCIBuild{}
	var branch string = ""
	this.Branch = branch
	var cause string = ""
	this.Cause = cause
	var codingCIId string = ""
	this.CodingCIId = codingCIId
	var commitId string = ""
	this.CommitId = commitId
	var failedMessage string = ""
	this.FailedMessage = failedMessage
	var jenkinsFileContent string = ""
	this.JenkinsFileContent = jenkinsFileContent
	var status string = ""
	this.Status = status
	var statusNode string = ""
	this.StatusNode = statusNode
	return &this
}

// GetBranch returns the Branch field value
func (o *CodingCIBuild) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *CodingCIBuild) SetBranch(v string) {
	o.Branch = v
}

// GetCause returns the Cause field value
func (o *CodingCIBuild) GetCause() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetCauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *CodingCIBuild) SetCause(v string) {
	o.Cause = v
}

// GetCodingCIId returns the CodingCIId field value
func (o *CodingCIBuild) GetCodingCIId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodingCIId
}

// GetCodingCIIdOk returns a tuple with the CodingCIId field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetCodingCIIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodingCIId, true
}

// SetCodingCIId sets field value
func (o *CodingCIBuild) SetCodingCIId(v string) {
	o.CodingCIId = v
}

// GetCommitId returns the CommitId field value
func (o *CodingCIBuild) GetCommitId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitId
}

// GetCommitIdOk returns a tuple with the CommitId field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetCommitIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitId, true
}

// SetCommitId sets field value
func (o *CodingCIBuild) SetCommitId(v string) {
	o.CommitId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CodingCIBuild) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CodingCIBuild) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetDuration returns the Duration field value
func (o *CodingCIBuild) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *CodingCIBuild) SetDuration(v int32) {
	o.Duration = v
}

// GetEnvList returns the EnvList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodingCIBuild) GetEnvList() []CIJobEnv {
	if o == nil {
		var ret []CIJobEnv
		return ret
	}
	return o.EnvList
}

// GetEnvListOk returns a tuple with the EnvList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodingCIBuild) GetEnvListOk() ([]CIJobEnv, bool) {
	if o == nil || utils.IsNil(o.EnvList) {
		return nil, false
	}
	return o.EnvList, true
}

// HasEnvList returns a boolean if a field has been set.
func (o *CodingCIBuild) HasEnvList() bool {
	if o != nil && !utils.IsNil(o.EnvList) {
		return true
	}

	return false
}

// SetEnvList gets a reference to the given []CIJobEnv and assigns it to the EnvList field.
func (o *CodingCIBuild) SetEnvList(v []CIJobEnv) {
	o.EnvList = v
}

// GetFailedMessage returns the FailedMessage field value
func (o *CodingCIBuild) GetFailedMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FailedMessage
}

// GetFailedMessageOk returns a tuple with the FailedMessage field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetFailedMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailedMessage, true
}

// SetFailedMessage sets field value
func (o *CodingCIBuild) SetFailedMessage(v string) {
	o.FailedMessage = v
}

// GetId returns the Id field value
func (o *CodingCIBuild) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CodingCIBuild) SetId(v int32) {
	o.Id = v
}

// GetJenkinsFileContent returns the JenkinsFileContent field value
func (o *CodingCIBuild) GetJenkinsFileContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JenkinsFileContent
}

// GetJenkinsFileContentOk returns a tuple with the JenkinsFileContent field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetJenkinsFileContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JenkinsFileContent, true
}

// SetJenkinsFileContent sets field value
func (o *CodingCIBuild) SetJenkinsFileContent(v string) {
	o.JenkinsFileContent = v
}

// GetJobId returns the JobId field value
func (o *CodingCIBuild) GetJobId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetJobIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *CodingCIBuild) SetJobId(v int32) {
	o.JobId = v
}

// GetNodeObtainedAt returns the NodeObtainedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodingCIBuild) GetNodeObtainedAt() int64 {
	if o == nil || utils.IsNil(o.NodeObtainedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.NodeObtainedAt.Get()
}

// GetNodeObtainedAtOk returns a tuple with the NodeObtainedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodingCIBuild) GetNodeObtainedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeObtainedAt.Get(), o.NodeObtainedAt.IsSet()
}

// HasNodeObtainedAt returns a boolean if a field has been set.
func (o *CodingCIBuild) HasNodeObtainedAt() bool {
	if o != nil && o.NodeObtainedAt.IsSet() {
		return true
	}

	return false
}

// SetNodeObtainedAt gets a reference to the given utils.NullableInt64 and assigns it to the NodeObtainedAt field.
func (o *CodingCIBuild) SetNodeObtainedAt(v int64) {
	o.NodeObtainedAt.Set(&v)
}
// SetNodeObtainedAtNil sets the value for NodeObtainedAt to be an explicit nil
func (o *CodingCIBuild) SetNodeObtainedAtNil() {
	o.NodeObtainedAt.Set(nil)
}

// UnsetNodeObtainedAt ensures that no value is present for NodeObtainedAt, not even an explicit nil
func (o *CodingCIBuild) UnsetNodeObtainedAt() {
	o.NodeObtainedAt.Unset()
}

// GetNumber returns the Number field value
func (o *CodingCIBuild) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *CodingCIBuild) SetNumber(v int32) {
	o.Number = v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CodingCIBuild) GetStartedAt() int64 {
	if o == nil || utils.IsNil(o.StartedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.StartedAt.Get()
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CodingCIBuild) GetStartedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartedAt.Get(), o.StartedAt.IsSet()
}

// HasStartedAt returns a boolean if a field has been set.
func (o *CodingCIBuild) HasStartedAt() bool {
	if o != nil && o.StartedAt.IsSet() {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given utils.NullableInt64 and assigns it to the StartedAt field.
func (o *CodingCIBuild) SetStartedAt(v int64) {
	o.StartedAt.Set(&v)
}
// SetStartedAtNil sets the value for StartedAt to be an explicit nil
func (o *CodingCIBuild) SetStartedAtNil() {
	o.StartedAt.Set(nil)
}

// UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
func (o *CodingCIBuild) UnsetStartedAt() {
	o.StartedAt.Unset()
}

// GetStatus returns the Status field value
func (o *CodingCIBuild) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CodingCIBuild) SetStatus(v string) {
	o.Status = v
}

// GetStatusNode returns the StatusNode field value
func (o *CodingCIBuild) GetStatusNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusNode
}

// GetStatusNodeOk returns a tuple with the StatusNode field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetStatusNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusNode, true
}

// SetStatusNode sets field value
func (o *CodingCIBuild) SetStatusNode(v string) {
	o.StatusNode = v
}

// GetTestResult returns the TestResult field value
func (o *CodingCIBuild) GetTestResult() CIBuildTestResult {
	if o == nil {
		var ret CIBuildTestResult
		return ret
	}

	return o.TestResult
}

// GetTestResultOk returns a tuple with the TestResult field value
// and a boolean to check if the value has been set.
func (o *CodingCIBuild) GetTestResultOk() (*CIBuildTestResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestResult, true
}

// SetTestResult sets field value
func (o *CodingCIBuild) SetTestResult(v CIBuildTestResult) {
	o.TestResult = v
}

func (o CodingCIBuild) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodingCIBuild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Branch"] = o.Branch
	toSerialize["Cause"] = o.Cause
	toSerialize["CodingCIId"] = o.CodingCIId
	toSerialize["CommitId"] = o.CommitId
	toSerialize["CreatedAt"] = o.CreatedAt
	toSerialize["Duration"] = o.Duration
	if o.EnvList != nil {
		toSerialize["EnvList"] = o.EnvList
	}
	toSerialize["FailedMessage"] = o.FailedMessage
	toSerialize["Id"] = o.Id
	toSerialize["JenkinsFileContent"] = o.JenkinsFileContent
	toSerialize["JobId"] = o.JobId
	if o.NodeObtainedAt.IsSet() {
		toSerialize["NodeObtainedAt"] = o.NodeObtainedAt.Get()
	}
	toSerialize["Number"] = o.Number
	if o.StartedAt.IsSet() {
		toSerialize["StartedAt"] = o.StartedAt.Get()
	}
	toSerialize["Status"] = o.Status
	toSerialize["StatusNode"] = o.StatusNode
	toSerialize["TestResult"] = o.TestResult
	return toSerialize, nil
}

func (o *CodingCIBuild) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Branch",
		"Cause",
		"CodingCIId",
		"CommitId",
		"CreatedAt",
		"Duration",
		"FailedMessage",
		"Id",
		"JenkinsFileContent",
		"JobId",
		"Number",
		"Status",
		"StatusNode",
		"TestResult",
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

	varCodingCIBuild := _CodingCIBuild{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCodingCIBuild)

	if err != nil {
		return err
	}

	*o = CodingCIBuild(varCodingCIBuild)

	return err
}

type NullableCodingCIBuild struct {
	value *CodingCIBuild
	isSet bool
}

func (v NullableCodingCIBuild) Get() *CodingCIBuild {
	return v.value
}

func (v *NullableCodingCIBuild) Set(val *CodingCIBuild) {
	v.value = val
	v.isSet = true
}

func (v NullableCodingCIBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableCodingCIBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodingCIBuild(val *CodingCIBuild) *NullableCodingCIBuild {
	return &NullableCodingCIBuild{value: val, isSet: true}
}

func (v NullableCodingCIBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodingCIBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


