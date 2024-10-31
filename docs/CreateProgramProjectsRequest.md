# CreateProgramProjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | **int64** | 项目集 ID | [default to 0]
**ProjectId** | **[]int64** | 项目 ID | 

## Methods

### NewCreateProgramProjectsRequest

`func NewCreateProgramProjectsRequest(programId int64, projectId []int64, ) *CreateProgramProjectsRequest`

NewCreateProgramProjectsRequest instantiates a new CreateProgramProjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProgramProjectsRequestWithDefaults

`func NewCreateProgramProjectsRequestWithDefaults() *CreateProgramProjectsRequest`

NewCreateProgramProjectsRequestWithDefaults instantiates a new CreateProgramProjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *CreateProgramProjectsRequest) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *CreateProgramProjectsRequest) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *CreateProgramProjectsRequest) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetProjectId

`func (o *CreateProgramProjectsRequest) GetProjectId() []int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateProgramProjectsRequest) GetProjectIdOk() (*[]int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateProgramProjectsRequest) SetProjectId(v []int64)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


