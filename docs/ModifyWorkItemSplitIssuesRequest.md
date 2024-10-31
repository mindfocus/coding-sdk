# ModifyWorkItemSplitIssuesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **string** | 目标项目中的事项ID | 
**ProgramName** | **string** | 项目集名称 | 
**ProjectName** | **string** | 目标项目名称  | 
**Split** | **string** | true 表示分解， false 表示取消分解 | 
**WorkItemCode** | **int64** | 页面上工作项ID | 

## Methods

### NewModifyWorkItemSplitIssuesRequest

`func NewModifyWorkItemSplitIssuesRequest(issueCode string, programName string, projectName string, split string, workItemCode int64, ) *ModifyWorkItemSplitIssuesRequest`

NewModifyWorkItemSplitIssuesRequest instantiates a new ModifyWorkItemSplitIssuesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWorkItemSplitIssuesRequestWithDefaults

`func NewModifyWorkItemSplitIssuesRequestWithDefaults() *ModifyWorkItemSplitIssuesRequest`

NewModifyWorkItemSplitIssuesRequestWithDefaults instantiates a new ModifyWorkItemSplitIssuesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *ModifyWorkItemSplitIssuesRequest) GetIssueCode() string`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *ModifyWorkItemSplitIssuesRequest) GetIssueCodeOk() (*string, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *ModifyWorkItemSplitIssuesRequest) SetIssueCode(v string)`

SetIssueCode sets IssueCode field to given value.


### GetProgramName

`func (o *ModifyWorkItemSplitIssuesRequest) GetProgramName() string`

GetProgramName returns the ProgramName field if non-nil, zero value otherwise.

### GetProgramNameOk

`func (o *ModifyWorkItemSplitIssuesRequest) GetProgramNameOk() (*string, bool)`

GetProgramNameOk returns a tuple with the ProgramName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramName

`func (o *ModifyWorkItemSplitIssuesRequest) SetProgramName(v string)`

SetProgramName sets ProgramName field to given value.


### GetProjectName

`func (o *ModifyWorkItemSplitIssuesRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyWorkItemSplitIssuesRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyWorkItemSplitIssuesRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetSplit

`func (o *ModifyWorkItemSplitIssuesRequest) GetSplit() string`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *ModifyWorkItemSplitIssuesRequest) GetSplitOk() (*string, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *ModifyWorkItemSplitIssuesRequest) SetSplit(v string)`

SetSplit sets Split field to given value.


### GetWorkItemCode

`func (o *ModifyWorkItemSplitIssuesRequest) GetWorkItemCode() int64`

GetWorkItemCode returns the WorkItemCode field if non-nil, zero value otherwise.

### GetWorkItemCodeOk

`func (o *ModifyWorkItemSplitIssuesRequest) GetWorkItemCodeOk() (*int64, bool)`

GetWorkItemCodeOk returns a tuple with the WorkItemCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemCode

`func (o *ModifyWorkItemSplitIssuesRequest) SetWorkItemCode(v int64)`

SetWorkItemCode sets WorkItemCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


