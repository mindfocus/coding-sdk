# Report

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | 附件列表 | [optional] 
**CreatedAt** | Pointer to **NullableString** | 创建时间 | [optional] [default to ""]
**CreatedBy** | Pointer to **NullableInt32** | 创建人 | [optional] 
**Id** | Pointer to **NullableInt32** | ID 主键 | [optional] 
**IterationId** | Pointer to **NullableString** | 迭代 ID | [optional] [default to ""]
**IterationName** | Pointer to **NullableString** | 迭代名称 | [optional] [default to ""]
**Name** | Pointer to **NullableString** | 报告名称 | [optional] [default to ""]
**ProjectName** | Pointer to **NullableString** | 项目名称 | [optional] [default to ""]
**ReportOverview** | Pointer to [**ReportOverview**](ReportOverview.md) |  | [optional] 
**RunIds** | Pointer to **[]string** | 测试计划 ID | [optional] 
**RunNames** | Pointer to **[]string** | 测试计划名称 | [optional] 
**StatisticsEndTime** | Pointer to **NullableString** | 数据统计结束时间 | [optional] [default to ""]
**StatisticsStartTime** | Pointer to **NullableString** | 数据统计开始时间 | [optional] [default to ""]
**Status** | Pointer to **NullableString** | 报告状态：CREATING 创建中，AVAILABLE 可用，UNAVAILABLE 不可用 | [optional] [default to ""]
**Summary** | Pointer to **NullableString** | 报告总结 | [optional] [default to ""]
**TemplateId** | Pointer to **NullableInt32** | 模板 ID | [optional] 

## Methods

### NewReport

`func NewReport() *Report`

NewReport instantiates a new Report object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportWithDefaults

`func NewReportWithDefaults() *Report`

NewReportWithDefaults instantiates a new Report object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *Report) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Report) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Report) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Report) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *Report) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *Report) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetCreatedAt

`func (o *Report) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Report) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Report) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Report) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *Report) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Report) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *Report) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Report) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Report) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Report) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *Report) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Report) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetId

`func (o *Report) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Report) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Report) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Report) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *Report) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Report) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIterationId

`func (o *Report) GetIterationId() string`

GetIterationId returns the IterationId field if non-nil, zero value otherwise.

### GetIterationIdOk

`func (o *Report) GetIterationIdOk() (*string, bool)`

GetIterationIdOk returns a tuple with the IterationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationId

`func (o *Report) SetIterationId(v string)`

SetIterationId sets IterationId field to given value.

### HasIterationId

`func (o *Report) HasIterationId() bool`

HasIterationId returns a boolean if a field has been set.

### SetIterationIdNil

`func (o *Report) SetIterationIdNil(b bool)`

 SetIterationIdNil sets the value for IterationId to be an explicit nil

### UnsetIterationId
`func (o *Report) UnsetIterationId()`

UnsetIterationId ensures that no value is present for IterationId, not even an explicit nil
### GetIterationName

`func (o *Report) GetIterationName() string`

GetIterationName returns the IterationName field if non-nil, zero value otherwise.

### GetIterationNameOk

`func (o *Report) GetIterationNameOk() (*string, bool)`

GetIterationNameOk returns a tuple with the IterationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationName

`func (o *Report) SetIterationName(v string)`

SetIterationName sets IterationName field to given value.

### HasIterationName

`func (o *Report) HasIterationName() bool`

HasIterationName returns a boolean if a field has been set.

### SetIterationNameNil

`func (o *Report) SetIterationNameNil(b bool)`

 SetIterationNameNil sets the value for IterationName to be an explicit nil

### UnsetIterationName
`func (o *Report) UnsetIterationName()`

UnsetIterationName ensures that no value is present for IterationName, not even an explicit nil
### GetName

`func (o *Report) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Report) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Report) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Report) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Report) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Report) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectName

`func (o *Report) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *Report) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *Report) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *Report) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *Report) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *Report) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetReportOverview

`func (o *Report) GetReportOverview() ReportOverview`

GetReportOverview returns the ReportOverview field if non-nil, zero value otherwise.

### GetReportOverviewOk

`func (o *Report) GetReportOverviewOk() (*ReportOverview, bool)`

GetReportOverviewOk returns a tuple with the ReportOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportOverview

`func (o *Report) SetReportOverview(v ReportOverview)`

SetReportOverview sets ReportOverview field to given value.

### HasReportOverview

`func (o *Report) HasReportOverview() bool`

HasReportOverview returns a boolean if a field has been set.

### GetRunIds

`func (o *Report) GetRunIds() []string`

GetRunIds returns the RunIds field if non-nil, zero value otherwise.

### GetRunIdsOk

`func (o *Report) GetRunIdsOk() (*[]string, bool)`

GetRunIdsOk returns a tuple with the RunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunIds

`func (o *Report) SetRunIds(v []string)`

SetRunIds sets RunIds field to given value.

### HasRunIds

`func (o *Report) HasRunIds() bool`

HasRunIds returns a boolean if a field has been set.

### SetRunIdsNil

`func (o *Report) SetRunIdsNil(b bool)`

 SetRunIdsNil sets the value for RunIds to be an explicit nil

### UnsetRunIds
`func (o *Report) UnsetRunIds()`

UnsetRunIds ensures that no value is present for RunIds, not even an explicit nil
### GetRunNames

`func (o *Report) GetRunNames() []string`

GetRunNames returns the RunNames field if non-nil, zero value otherwise.

### GetRunNamesOk

`func (o *Report) GetRunNamesOk() (*[]string, bool)`

GetRunNamesOk returns a tuple with the RunNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNames

`func (o *Report) SetRunNames(v []string)`

SetRunNames sets RunNames field to given value.

### HasRunNames

`func (o *Report) HasRunNames() bool`

HasRunNames returns a boolean if a field has been set.

### SetRunNamesNil

`func (o *Report) SetRunNamesNil(b bool)`

 SetRunNamesNil sets the value for RunNames to be an explicit nil

### UnsetRunNames
`func (o *Report) UnsetRunNames()`

UnsetRunNames ensures that no value is present for RunNames, not even an explicit nil
### GetStatisticsEndTime

`func (o *Report) GetStatisticsEndTime() string`

GetStatisticsEndTime returns the StatisticsEndTime field if non-nil, zero value otherwise.

### GetStatisticsEndTimeOk

`func (o *Report) GetStatisticsEndTimeOk() (*string, bool)`

GetStatisticsEndTimeOk returns a tuple with the StatisticsEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsEndTime

`func (o *Report) SetStatisticsEndTime(v string)`

SetStatisticsEndTime sets StatisticsEndTime field to given value.

### HasStatisticsEndTime

`func (o *Report) HasStatisticsEndTime() bool`

HasStatisticsEndTime returns a boolean if a field has been set.

### SetStatisticsEndTimeNil

`func (o *Report) SetStatisticsEndTimeNil(b bool)`

 SetStatisticsEndTimeNil sets the value for StatisticsEndTime to be an explicit nil

### UnsetStatisticsEndTime
`func (o *Report) UnsetStatisticsEndTime()`

UnsetStatisticsEndTime ensures that no value is present for StatisticsEndTime, not even an explicit nil
### GetStatisticsStartTime

`func (o *Report) GetStatisticsStartTime() string`

GetStatisticsStartTime returns the StatisticsStartTime field if non-nil, zero value otherwise.

### GetStatisticsStartTimeOk

`func (o *Report) GetStatisticsStartTimeOk() (*string, bool)`

GetStatisticsStartTimeOk returns a tuple with the StatisticsStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsStartTime

`func (o *Report) SetStatisticsStartTime(v string)`

SetStatisticsStartTime sets StatisticsStartTime field to given value.

### HasStatisticsStartTime

`func (o *Report) HasStatisticsStartTime() bool`

HasStatisticsStartTime returns a boolean if a field has been set.

### SetStatisticsStartTimeNil

`func (o *Report) SetStatisticsStartTimeNil(b bool)`

 SetStatisticsStartTimeNil sets the value for StatisticsStartTime to be an explicit nil

### UnsetStatisticsStartTime
`func (o *Report) UnsetStatisticsStartTime()`

UnsetStatisticsStartTime ensures that no value is present for StatisticsStartTime, not even an explicit nil
### GetStatus

`func (o *Report) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Report) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Report) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Report) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Report) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Report) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetSummary

`func (o *Report) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Report) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Report) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Report) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *Report) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *Report) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetTemplateId

`func (o *Report) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *Report) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *Report) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *Report) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *Report) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *Report) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


