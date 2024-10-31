# ModifyTestCaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentIds** | Pointer to **[]int32** | 附件 ID 数组：来自“生成附件预上传信息”接口 | [optional] 
**CaseId** | **int32** | 用例 ID | 
**CustomSteps** | Pointer to [**[]CustomStep**](CustomStep.md) | 自定义步骤（步骤用例必填） | [optional] 
**Expected** | Pointer to **string** | 预期结果（适用于文本用例） | [optional] 
**Preconds** | Pointer to **string** | 前置步骤 | [optional] 
**Priority** | Pointer to **int32** | 优先级，默认2（中），可选值：0（紧急）,1（高）,2（中）,3（低） | [optional] 
**ProjectName** | **string** | 项目名称 | 
**SectionId** | **int32** | 分组ID | 
**Steps** | Pointer to **string** | 文本描述（适用于文本用例） | [optional] 
**TemplateType** | **string** | 用例类型，可选值：STEPS(步骤用例)，TEXT(文本用例) | 
**Title** | **string** | 用例标题 | 

## Methods

### NewModifyTestCaseRequest

`func NewModifyTestCaseRequest(caseId int32, projectName string, sectionId int32, templateType string, title string, ) *ModifyTestCaseRequest`

NewModifyTestCaseRequest instantiates a new ModifyTestCaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyTestCaseRequestWithDefaults

`func NewModifyTestCaseRequestWithDefaults() *ModifyTestCaseRequest`

NewModifyTestCaseRequestWithDefaults instantiates a new ModifyTestCaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentIds

`func (o *ModifyTestCaseRequest) GetAttachmentIds() []int32`

GetAttachmentIds returns the AttachmentIds field if non-nil, zero value otherwise.

### GetAttachmentIdsOk

`func (o *ModifyTestCaseRequest) GetAttachmentIdsOk() (*[]int32, bool)`

GetAttachmentIdsOk returns a tuple with the AttachmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentIds

`func (o *ModifyTestCaseRequest) SetAttachmentIds(v []int32)`

SetAttachmentIds sets AttachmentIds field to given value.

### HasAttachmentIds

`func (o *ModifyTestCaseRequest) HasAttachmentIds() bool`

HasAttachmentIds returns a boolean if a field has been set.

### GetCaseId

`func (o *ModifyTestCaseRequest) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *ModifyTestCaseRequest) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *ModifyTestCaseRequest) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.


### GetCustomSteps

`func (o *ModifyTestCaseRequest) GetCustomSteps() []CustomStep`

GetCustomSteps returns the CustomSteps field if non-nil, zero value otherwise.

### GetCustomStepsOk

`func (o *ModifyTestCaseRequest) GetCustomStepsOk() (*[]CustomStep, bool)`

GetCustomStepsOk returns a tuple with the CustomSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSteps

`func (o *ModifyTestCaseRequest) SetCustomSteps(v []CustomStep)`

SetCustomSteps sets CustomSteps field to given value.

### HasCustomSteps

`func (o *ModifyTestCaseRequest) HasCustomSteps() bool`

HasCustomSteps returns a boolean if a field has been set.

### GetExpected

`func (o *ModifyTestCaseRequest) GetExpected() string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *ModifyTestCaseRequest) GetExpectedOk() (*string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *ModifyTestCaseRequest) SetExpected(v string)`

SetExpected sets Expected field to given value.

### HasExpected

`func (o *ModifyTestCaseRequest) HasExpected() bool`

HasExpected returns a boolean if a field has been set.

### GetPreconds

`func (o *ModifyTestCaseRequest) GetPreconds() string`

GetPreconds returns the Preconds field if non-nil, zero value otherwise.

### GetPrecondsOk

`func (o *ModifyTestCaseRequest) GetPrecondsOk() (*string, bool)`

GetPrecondsOk returns a tuple with the Preconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconds

`func (o *ModifyTestCaseRequest) SetPreconds(v string)`

SetPreconds sets Preconds field to given value.

### HasPreconds

`func (o *ModifyTestCaseRequest) HasPreconds() bool`

HasPreconds returns a boolean if a field has been set.

### GetPriority

`func (o *ModifyTestCaseRequest) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModifyTestCaseRequest) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModifyTestCaseRequest) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModifyTestCaseRequest) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyTestCaseRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyTestCaseRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyTestCaseRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSectionId

`func (o *ModifyTestCaseRequest) GetSectionId() int32`

GetSectionId returns the SectionId field if non-nil, zero value otherwise.

### GetSectionIdOk

`func (o *ModifyTestCaseRequest) GetSectionIdOk() (*int32, bool)`

GetSectionIdOk returns a tuple with the SectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionId

`func (o *ModifyTestCaseRequest) SetSectionId(v int32)`

SetSectionId sets SectionId field to given value.


### GetSteps

`func (o *ModifyTestCaseRequest) GetSteps() string`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *ModifyTestCaseRequest) GetStepsOk() (*string, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *ModifyTestCaseRequest) SetSteps(v string)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *ModifyTestCaseRequest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTemplateType

`func (o *ModifyTestCaseRequest) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *ModifyTestCaseRequest) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *ModifyTestCaseRequest) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.


### GetTitle

`func (o *ModifyTestCaseRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModifyTestCaseRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModifyTestCaseRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

