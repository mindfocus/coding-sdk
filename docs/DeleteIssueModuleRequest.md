# DeleteIssueModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModuleId** | **int64** | 模块ID | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDeleteIssueModuleRequest

`func NewDeleteIssueModuleRequest(moduleId int64, projectName string, ) *DeleteIssueModuleRequest`

NewDeleteIssueModuleRequest instantiates a new DeleteIssueModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIssueModuleRequestWithDefaults

`func NewDeleteIssueModuleRequestWithDefaults() *DeleteIssueModuleRequest`

NewDeleteIssueModuleRequestWithDefaults instantiates a new DeleteIssueModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModuleId

`func (o *DeleteIssueModuleRequest) GetModuleId() int64`

GetModuleId returns the ModuleId field if non-nil, zero value otherwise.

### GetModuleIdOk

`func (o *DeleteIssueModuleRequest) GetModuleIdOk() (*int64, bool)`

GetModuleIdOk returns a tuple with the ModuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleId

`func (o *DeleteIssueModuleRequest) SetModuleId(v int64)`

SetModuleId sets ModuleId field to given value.


### GetProjectName

`func (o *DeleteIssueModuleRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteIssueModuleRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteIssueModuleRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


