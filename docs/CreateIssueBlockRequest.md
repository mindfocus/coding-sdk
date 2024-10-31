# CreateIssueBlockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockIssueCode** | **int64** | 前置事项 Code | 
**IssueCode** | **int64** | 事项 Code | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewCreateIssueBlockRequest

`func NewCreateIssueBlockRequest(blockIssueCode int64, issueCode int64, projectName string, ) *CreateIssueBlockRequest`

NewCreateIssueBlockRequest instantiates a new CreateIssueBlockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIssueBlockRequestWithDefaults

`func NewCreateIssueBlockRequestWithDefaults() *CreateIssueBlockRequest`

NewCreateIssueBlockRequestWithDefaults instantiates a new CreateIssueBlockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockIssueCode

`func (o *CreateIssueBlockRequest) GetBlockIssueCode() int64`

GetBlockIssueCode returns the BlockIssueCode field if non-nil, zero value otherwise.

### GetBlockIssueCodeOk

`func (o *CreateIssueBlockRequest) GetBlockIssueCodeOk() (*int64, bool)`

GetBlockIssueCodeOk returns a tuple with the BlockIssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockIssueCode

`func (o *CreateIssueBlockRequest) SetBlockIssueCode(v int64)`

SetBlockIssueCode sets BlockIssueCode field to given value.


### GetIssueCode

`func (o *CreateIssueBlockRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *CreateIssueBlockRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *CreateIssueBlockRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *CreateIssueBlockRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateIssueBlockRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateIssueBlockRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


