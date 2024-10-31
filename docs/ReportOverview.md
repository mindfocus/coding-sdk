# ReportOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomationPercent** | Pointer to **NullableInt32** | 自动化覆盖率：百分比 | [optional] 
**AvgClosedSeconds** | Pointer to **NullableInt64** | 平均关闭时长 | [optional] 
**CaseSum** | Pointer to **NullableInt32** | 用例总数 | [optional] 
**CompletedSum** | Pointer to **NullableInt32** | 已完成数量 | [optional] 
**DefectFixPercent** | Pointer to **NullableInt32** | 缺陷修复率：百分比 | [optional] 
**DefectReopenPercent** | Pointer to **NullableInt32** | 重新激活率：百分比 | [optional] 
**DefectSum** | Pointer to **NullableInt32** | 缺陷总数 | [optional] 
**DurationFixed** | Pointer to **NullableInt64** | 85%解决时长 | [optional] 
**ExecPercent** | Pointer to **NullableInt32** | 执行率：百分比 | [optional] 
**IssuesSum** | Pointer to **NullableInt32** | 需求总数 | [optional] 
**PassPercent** | Pointer to **NullableInt32** | 通过率：百分比 | [optional] 
**ProcessingSum** | Pointer to **NullableInt32** | 处理中数量 | [optional] 
**RequirementCoverPercent** | Pointer to **NullableInt32** | 需求覆盖率：百分比 | [optional] 
**TodoSum** | Pointer to **NullableInt32** | 未开始数量 | [optional] 

## Methods

### NewReportOverview

`func NewReportOverview() *ReportOverview`

NewReportOverview instantiates a new ReportOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportOverviewWithDefaults

`func NewReportOverviewWithDefaults() *ReportOverview`

NewReportOverviewWithDefaults instantiates a new ReportOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomationPercent

`func (o *ReportOverview) GetAutomationPercent() int32`

GetAutomationPercent returns the AutomationPercent field if non-nil, zero value otherwise.

### GetAutomationPercentOk

`func (o *ReportOverview) GetAutomationPercentOk() (*int32, bool)`

GetAutomationPercentOk returns a tuple with the AutomationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationPercent

`func (o *ReportOverview) SetAutomationPercent(v int32)`

SetAutomationPercent sets AutomationPercent field to given value.

### HasAutomationPercent

`func (o *ReportOverview) HasAutomationPercent() bool`

HasAutomationPercent returns a boolean if a field has been set.

### SetAutomationPercentNil

`func (o *ReportOverview) SetAutomationPercentNil(b bool)`

 SetAutomationPercentNil sets the value for AutomationPercent to be an explicit nil

### UnsetAutomationPercent
`func (o *ReportOverview) UnsetAutomationPercent()`

UnsetAutomationPercent ensures that no value is present for AutomationPercent, not even an explicit nil
### GetAvgClosedSeconds

`func (o *ReportOverview) GetAvgClosedSeconds() int64`

GetAvgClosedSeconds returns the AvgClosedSeconds field if non-nil, zero value otherwise.

### GetAvgClosedSecondsOk

`func (o *ReportOverview) GetAvgClosedSecondsOk() (*int64, bool)`

GetAvgClosedSecondsOk returns a tuple with the AvgClosedSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgClosedSeconds

`func (o *ReportOverview) SetAvgClosedSeconds(v int64)`

SetAvgClosedSeconds sets AvgClosedSeconds field to given value.

### HasAvgClosedSeconds

`func (o *ReportOverview) HasAvgClosedSeconds() bool`

HasAvgClosedSeconds returns a boolean if a field has been set.

### SetAvgClosedSecondsNil

`func (o *ReportOverview) SetAvgClosedSecondsNil(b bool)`

 SetAvgClosedSecondsNil sets the value for AvgClosedSeconds to be an explicit nil

### UnsetAvgClosedSeconds
`func (o *ReportOverview) UnsetAvgClosedSeconds()`

UnsetAvgClosedSeconds ensures that no value is present for AvgClosedSeconds, not even an explicit nil
### GetCaseSum

`func (o *ReportOverview) GetCaseSum() int32`

GetCaseSum returns the CaseSum field if non-nil, zero value otherwise.

### GetCaseSumOk

`func (o *ReportOverview) GetCaseSumOk() (*int32, bool)`

GetCaseSumOk returns a tuple with the CaseSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSum

`func (o *ReportOverview) SetCaseSum(v int32)`

SetCaseSum sets CaseSum field to given value.

### HasCaseSum

`func (o *ReportOverview) HasCaseSum() bool`

HasCaseSum returns a boolean if a field has been set.

### SetCaseSumNil

`func (o *ReportOverview) SetCaseSumNil(b bool)`

 SetCaseSumNil sets the value for CaseSum to be an explicit nil

### UnsetCaseSum
`func (o *ReportOverview) UnsetCaseSum()`

UnsetCaseSum ensures that no value is present for CaseSum, not even an explicit nil
### GetCompletedSum

`func (o *ReportOverview) GetCompletedSum() int32`

GetCompletedSum returns the CompletedSum field if non-nil, zero value otherwise.

### GetCompletedSumOk

`func (o *ReportOverview) GetCompletedSumOk() (*int32, bool)`

GetCompletedSumOk returns a tuple with the CompletedSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedSum

`func (o *ReportOverview) SetCompletedSum(v int32)`

SetCompletedSum sets CompletedSum field to given value.

### HasCompletedSum

`func (o *ReportOverview) HasCompletedSum() bool`

HasCompletedSum returns a boolean if a field has been set.

### SetCompletedSumNil

`func (o *ReportOverview) SetCompletedSumNil(b bool)`

 SetCompletedSumNil sets the value for CompletedSum to be an explicit nil

### UnsetCompletedSum
`func (o *ReportOverview) UnsetCompletedSum()`

UnsetCompletedSum ensures that no value is present for CompletedSum, not even an explicit nil
### GetDefectFixPercent

`func (o *ReportOverview) GetDefectFixPercent() int32`

GetDefectFixPercent returns the DefectFixPercent field if non-nil, zero value otherwise.

### GetDefectFixPercentOk

`func (o *ReportOverview) GetDefectFixPercentOk() (*int32, bool)`

GetDefectFixPercentOk returns a tuple with the DefectFixPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectFixPercent

`func (o *ReportOverview) SetDefectFixPercent(v int32)`

SetDefectFixPercent sets DefectFixPercent field to given value.

### HasDefectFixPercent

`func (o *ReportOverview) HasDefectFixPercent() bool`

HasDefectFixPercent returns a boolean if a field has been set.

### SetDefectFixPercentNil

`func (o *ReportOverview) SetDefectFixPercentNil(b bool)`

 SetDefectFixPercentNil sets the value for DefectFixPercent to be an explicit nil

### UnsetDefectFixPercent
`func (o *ReportOverview) UnsetDefectFixPercent()`

UnsetDefectFixPercent ensures that no value is present for DefectFixPercent, not even an explicit nil
### GetDefectReopenPercent

`func (o *ReportOverview) GetDefectReopenPercent() int32`

GetDefectReopenPercent returns the DefectReopenPercent field if non-nil, zero value otherwise.

### GetDefectReopenPercentOk

`func (o *ReportOverview) GetDefectReopenPercentOk() (*int32, bool)`

GetDefectReopenPercentOk returns a tuple with the DefectReopenPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectReopenPercent

`func (o *ReportOverview) SetDefectReopenPercent(v int32)`

SetDefectReopenPercent sets DefectReopenPercent field to given value.

### HasDefectReopenPercent

`func (o *ReportOverview) HasDefectReopenPercent() bool`

HasDefectReopenPercent returns a boolean if a field has been set.

### SetDefectReopenPercentNil

`func (o *ReportOverview) SetDefectReopenPercentNil(b bool)`

 SetDefectReopenPercentNil sets the value for DefectReopenPercent to be an explicit nil

### UnsetDefectReopenPercent
`func (o *ReportOverview) UnsetDefectReopenPercent()`

UnsetDefectReopenPercent ensures that no value is present for DefectReopenPercent, not even an explicit nil
### GetDefectSum

`func (o *ReportOverview) GetDefectSum() int32`

GetDefectSum returns the DefectSum field if non-nil, zero value otherwise.

### GetDefectSumOk

`func (o *ReportOverview) GetDefectSumOk() (*int32, bool)`

GetDefectSumOk returns a tuple with the DefectSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectSum

`func (o *ReportOverview) SetDefectSum(v int32)`

SetDefectSum sets DefectSum field to given value.

### HasDefectSum

`func (o *ReportOverview) HasDefectSum() bool`

HasDefectSum returns a boolean if a field has been set.

### SetDefectSumNil

`func (o *ReportOverview) SetDefectSumNil(b bool)`

 SetDefectSumNil sets the value for DefectSum to be an explicit nil

### UnsetDefectSum
`func (o *ReportOverview) UnsetDefectSum()`

UnsetDefectSum ensures that no value is present for DefectSum, not even an explicit nil
### GetDurationFixed

`func (o *ReportOverview) GetDurationFixed() int64`

GetDurationFixed returns the DurationFixed field if non-nil, zero value otherwise.

### GetDurationFixedOk

`func (o *ReportOverview) GetDurationFixedOk() (*int64, bool)`

GetDurationFixedOk returns a tuple with the DurationFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationFixed

`func (o *ReportOverview) SetDurationFixed(v int64)`

SetDurationFixed sets DurationFixed field to given value.

### HasDurationFixed

`func (o *ReportOverview) HasDurationFixed() bool`

HasDurationFixed returns a boolean if a field has been set.

### SetDurationFixedNil

`func (o *ReportOverview) SetDurationFixedNil(b bool)`

 SetDurationFixedNil sets the value for DurationFixed to be an explicit nil

### UnsetDurationFixed
`func (o *ReportOverview) UnsetDurationFixed()`

UnsetDurationFixed ensures that no value is present for DurationFixed, not even an explicit nil
### GetExecPercent

`func (o *ReportOverview) GetExecPercent() int32`

GetExecPercent returns the ExecPercent field if non-nil, zero value otherwise.

### GetExecPercentOk

`func (o *ReportOverview) GetExecPercentOk() (*int32, bool)`

GetExecPercentOk returns a tuple with the ExecPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecPercent

`func (o *ReportOverview) SetExecPercent(v int32)`

SetExecPercent sets ExecPercent field to given value.

### HasExecPercent

`func (o *ReportOverview) HasExecPercent() bool`

HasExecPercent returns a boolean if a field has been set.

### SetExecPercentNil

`func (o *ReportOverview) SetExecPercentNil(b bool)`

 SetExecPercentNil sets the value for ExecPercent to be an explicit nil

### UnsetExecPercent
`func (o *ReportOverview) UnsetExecPercent()`

UnsetExecPercent ensures that no value is present for ExecPercent, not even an explicit nil
### GetIssuesSum

`func (o *ReportOverview) GetIssuesSum() int32`

GetIssuesSum returns the IssuesSum field if non-nil, zero value otherwise.

### GetIssuesSumOk

`func (o *ReportOverview) GetIssuesSumOk() (*int32, bool)`

GetIssuesSumOk returns a tuple with the IssuesSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuesSum

`func (o *ReportOverview) SetIssuesSum(v int32)`

SetIssuesSum sets IssuesSum field to given value.

### HasIssuesSum

`func (o *ReportOverview) HasIssuesSum() bool`

HasIssuesSum returns a boolean if a field has been set.

### SetIssuesSumNil

`func (o *ReportOverview) SetIssuesSumNil(b bool)`

 SetIssuesSumNil sets the value for IssuesSum to be an explicit nil

### UnsetIssuesSum
`func (o *ReportOverview) UnsetIssuesSum()`

UnsetIssuesSum ensures that no value is present for IssuesSum, not even an explicit nil
### GetPassPercent

`func (o *ReportOverview) GetPassPercent() int32`

GetPassPercent returns the PassPercent field if non-nil, zero value otherwise.

### GetPassPercentOk

`func (o *ReportOverview) GetPassPercentOk() (*int32, bool)`

GetPassPercentOk returns a tuple with the PassPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassPercent

`func (o *ReportOverview) SetPassPercent(v int32)`

SetPassPercent sets PassPercent field to given value.

### HasPassPercent

`func (o *ReportOverview) HasPassPercent() bool`

HasPassPercent returns a boolean if a field has been set.

### SetPassPercentNil

`func (o *ReportOverview) SetPassPercentNil(b bool)`

 SetPassPercentNil sets the value for PassPercent to be an explicit nil

### UnsetPassPercent
`func (o *ReportOverview) UnsetPassPercent()`

UnsetPassPercent ensures that no value is present for PassPercent, not even an explicit nil
### GetProcessingSum

`func (o *ReportOverview) GetProcessingSum() int32`

GetProcessingSum returns the ProcessingSum field if non-nil, zero value otherwise.

### GetProcessingSumOk

`func (o *ReportOverview) GetProcessingSumOk() (*int32, bool)`

GetProcessingSumOk returns a tuple with the ProcessingSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingSum

`func (o *ReportOverview) SetProcessingSum(v int32)`

SetProcessingSum sets ProcessingSum field to given value.

### HasProcessingSum

`func (o *ReportOverview) HasProcessingSum() bool`

HasProcessingSum returns a boolean if a field has been set.

### SetProcessingSumNil

`func (o *ReportOverview) SetProcessingSumNil(b bool)`

 SetProcessingSumNil sets the value for ProcessingSum to be an explicit nil

### UnsetProcessingSum
`func (o *ReportOverview) UnsetProcessingSum()`

UnsetProcessingSum ensures that no value is present for ProcessingSum, not even an explicit nil
### GetRequirementCoverPercent

`func (o *ReportOverview) GetRequirementCoverPercent() int32`

GetRequirementCoverPercent returns the RequirementCoverPercent field if non-nil, zero value otherwise.

### GetRequirementCoverPercentOk

`func (o *ReportOverview) GetRequirementCoverPercentOk() (*int32, bool)`

GetRequirementCoverPercentOk returns a tuple with the RequirementCoverPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementCoverPercent

`func (o *ReportOverview) SetRequirementCoverPercent(v int32)`

SetRequirementCoverPercent sets RequirementCoverPercent field to given value.

### HasRequirementCoverPercent

`func (o *ReportOverview) HasRequirementCoverPercent() bool`

HasRequirementCoverPercent returns a boolean if a field has been set.

### SetRequirementCoverPercentNil

`func (o *ReportOverview) SetRequirementCoverPercentNil(b bool)`

 SetRequirementCoverPercentNil sets the value for RequirementCoverPercent to be an explicit nil

### UnsetRequirementCoverPercent
`func (o *ReportOverview) UnsetRequirementCoverPercent()`

UnsetRequirementCoverPercent ensures that no value is present for RequirementCoverPercent, not even an explicit nil
### GetTodoSum

`func (o *ReportOverview) GetTodoSum() int32`

GetTodoSum returns the TodoSum field if non-nil, zero value otherwise.

### GetTodoSumOk

`func (o *ReportOverview) GetTodoSumOk() (*int32, bool)`

GetTodoSumOk returns a tuple with the TodoSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTodoSum

`func (o *ReportOverview) SetTodoSum(v int32)`

SetTodoSum sets TodoSum field to given value.

### HasTodoSum

`func (o *ReportOverview) HasTodoSum() bool`

HasTodoSum returns a boolean if a field has been set.

### SetTodoSumNil

`func (o *ReportOverview) SetTodoSumNil(b bool)`

 SetTodoSumNil sets the value for TodoSum to be an explicit nil

### UnsetTodoSum
`func (o *ReportOverview) UnsetTodoSum()`

UnsetTodoSum ensures that no value is present for TodoSum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


