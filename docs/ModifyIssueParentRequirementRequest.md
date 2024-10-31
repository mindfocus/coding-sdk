# ModifyIssueParentRequirementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项 Code | 
**ParentIssueCode** | **int64** | 父事项 Code | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewModifyIssueParentRequirementRequest

`func NewModifyIssueParentRequirementRequest(issueCode int64, parentIssueCode int64, projectName string, ) *ModifyIssueParentRequirementRequest`

NewModifyIssueParentRequirementRequest instantiates a new ModifyIssueParentRequirementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyIssueParentRequirementRequestWithDefaults

`func NewModifyIssueParentRequirementRequestWithDefaults() *ModifyIssueParentRequirementRequest`

NewModifyIssueParentRequirementRequestWithDefaults instantiates a new ModifyIssueParentRequirementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *ModifyIssueParentRequirementRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *ModifyIssueParentRequirementRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *ModifyIssueParentRequirementRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetParentIssueCode

`func (o *ModifyIssueParentRequirementRequest) GetParentIssueCode() int64`

GetParentIssueCode returns the ParentIssueCode field if non-nil, zero value otherwise.

### GetParentIssueCodeOk

`func (o *ModifyIssueParentRequirementRequest) GetParentIssueCodeOk() (*int64, bool)`

GetParentIssueCodeOk returns a tuple with the ParentIssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIssueCode

`func (o *ModifyIssueParentRequirementRequest) SetParentIssueCode(v int64)`

SetParentIssueCode sets ParentIssueCode field to given value.


### GetProjectName

`func (o *ModifyIssueParentRequirementRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyIssueParentRequirementRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyIssueParentRequirementRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


