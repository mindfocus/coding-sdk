# CreateCodingCIJobByTeamTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 描述信息 | [default to 0]
**DepotType** | **string** | 仓库类型 | 
**JobName** | **string** | 构建计划名称 | [default to ""]
**ProjectId** | **int64** | 项目 ID | [default to 0]
**TemplateId** | **int64** | 团队构建模版 ID | [default to 0]

## Methods

### NewCreateCodingCIJobByTeamTemplateRequest

`func NewCreateCodingCIJobByTeamTemplateRequest(depotId int64, depotType string, jobName string, projectId int64, templateId int64, ) *CreateCodingCIJobByTeamTemplateRequest`

NewCreateCodingCIJobByTeamTemplateRequest instantiates a new CreateCodingCIJobByTeamTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCodingCIJobByTeamTemplateRequestWithDefaults

`func NewCreateCodingCIJobByTeamTemplateRequestWithDefaults() *CreateCodingCIJobByTeamTemplateRequest`

NewCreateCodingCIJobByTeamTemplateRequestWithDefaults instantiates a new CreateCodingCIJobByTeamTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateCodingCIJobByTeamTemplateRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotType

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *CreateCodingCIJobByTeamTemplateRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetJobName

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetJobName() string`

GetJobName returns the JobName field if non-nil, zero value otherwise.

### GetJobNameOk

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetJobNameOk() (*string, bool)`

GetJobNameOk returns a tuple with the JobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobName

`func (o *CreateCodingCIJobByTeamTemplateRequest) SetJobName(v string)`

SetJobName sets JobName field to given value.


### GetProjectId

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateCodingCIJobByTeamTemplateRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetTemplateId

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *CreateCodingCIJobByTeamTemplateRequest) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *CreateCodingCIJobByTeamTemplateRequest) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


