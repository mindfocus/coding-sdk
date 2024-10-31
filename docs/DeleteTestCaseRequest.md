# DeleteTestCaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseId** | **int32** | 用例 ID | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDeleteTestCaseRequest

`func NewDeleteTestCaseRequest(caseId int32, projectName string, ) *DeleteTestCaseRequest`

NewDeleteTestCaseRequest instantiates a new DeleteTestCaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteTestCaseRequestWithDefaults

`func NewDeleteTestCaseRequestWithDefaults() *DeleteTestCaseRequest`

NewDeleteTestCaseRequestWithDefaults instantiates a new DeleteTestCaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseId

`func (o *DeleteTestCaseRequest) GetCaseId() int32`

GetCaseId returns the CaseId field if non-nil, zero value otherwise.

### GetCaseIdOk

`func (o *DeleteTestCaseRequest) GetCaseIdOk() (*int32, bool)`

GetCaseIdOk returns a tuple with the CaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseId

`func (o *DeleteTestCaseRequest) SetCaseId(v int32)`

SetCaseId sets CaseId field to given value.


### GetProjectName

`func (o *DeleteTestCaseRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteTestCaseRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteTestCaseRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


