# ArchiveTestRunRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**RunId** | **int32** | 计划 ID | 

## Methods

### NewArchiveTestRunRequest

`func NewArchiveTestRunRequest(projectName string, runId int32, ) *ArchiveTestRunRequest`

NewArchiveTestRunRequest instantiates a new ArchiveTestRunRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveTestRunRequestWithDefaults

`func NewArchiveTestRunRequestWithDefaults() *ArchiveTestRunRequest`

NewArchiveTestRunRequestWithDefaults instantiates a new ArchiveTestRunRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *ArchiveTestRunRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ArchiveTestRunRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ArchiveTestRunRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetRunId

`func (o *ArchiveTestRunRequest) GetRunId() int32`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ArchiveTestRunRequest) GetRunIdOk() (*int32, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ArchiveTestRunRequest) SetRunId(v int32)`

SetRunId sets RunId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


