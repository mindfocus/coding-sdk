# CreateTestDefectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefectId** | **int32** | 缺陷 ID | 
**ProjectName** | **string** | 项目名称 | 
**TestId** | **int32** | 测试任务 ID | 

## Methods

### NewCreateTestDefectRequest

`func NewCreateTestDefectRequest(defectId int32, projectName string, testId int32, ) *CreateTestDefectRequest`

NewCreateTestDefectRequest instantiates a new CreateTestDefectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestDefectRequestWithDefaults

`func NewCreateTestDefectRequestWithDefaults() *CreateTestDefectRequest`

NewCreateTestDefectRequestWithDefaults instantiates a new CreateTestDefectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefectId

`func (o *CreateTestDefectRequest) GetDefectId() int32`

GetDefectId returns the DefectId field if non-nil, zero value otherwise.

### GetDefectIdOk

`func (o *CreateTestDefectRequest) GetDefectIdOk() (*int32, bool)`

GetDefectIdOk returns a tuple with the DefectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectId

`func (o *CreateTestDefectRequest) SetDefectId(v int32)`

SetDefectId sets DefectId field to given value.


### GetProjectName

`func (o *CreateTestDefectRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateTestDefectRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateTestDefectRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTestId

`func (o *CreateTestDefectRequest) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *CreateTestDefectRequest) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *CreateTestDefectRequest) SetTestId(v int32)`

SetTestId sets TestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


