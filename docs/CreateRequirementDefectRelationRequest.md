# CreateRequirementDefectRelationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefectCode** | **int64** | 缺陷 Code | 
**ProjectName** | **string** | 项目名称 | 
**RequirementCode** | **int64** | 需求 Code | 

## Methods

### NewCreateRequirementDefectRelationRequest

`func NewCreateRequirementDefectRelationRequest(defectCode int64, projectName string, requirementCode int64, ) *CreateRequirementDefectRelationRequest`

NewCreateRequirementDefectRelationRequest instantiates a new CreateRequirementDefectRelationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequirementDefectRelationRequestWithDefaults

`func NewCreateRequirementDefectRelationRequestWithDefaults() *CreateRequirementDefectRelationRequest`

NewCreateRequirementDefectRelationRequestWithDefaults instantiates a new CreateRequirementDefectRelationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefectCode

`func (o *CreateRequirementDefectRelationRequest) GetDefectCode() int64`

GetDefectCode returns the DefectCode field if non-nil, zero value otherwise.

### GetDefectCodeOk

`func (o *CreateRequirementDefectRelationRequest) GetDefectCodeOk() (*int64, bool)`

GetDefectCodeOk returns a tuple with the DefectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCode

`func (o *CreateRequirementDefectRelationRequest) SetDefectCode(v int64)`

SetDefectCode sets DefectCode field to given value.


### GetProjectName

`func (o *CreateRequirementDefectRelationRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateRequirementDefectRelationRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateRequirementDefectRelationRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRequirementCode

`func (o *CreateRequirementDefectRelationRequest) GetRequirementCode() int64`

GetRequirementCode returns the RequirementCode field if non-nil, zero value otherwise.

### GetRequirementCodeOk

`func (o *CreateRequirementDefectRelationRequest) GetRequirementCodeOk() (*int64, bool)`

GetRequirementCodeOk returns a tuple with the RequirementCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementCode

`func (o *CreateRequirementDefectRelationRequest) SetRequirementCode(v int64)`

SetRequirementCode sets RequirementCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


