# CreateTestCaseSectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 分组名称 | 
**ParentId** | Pointer to **int32** | 父级 ID，默认 0 | [optional] 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewCreateTestCaseSectionRequest

`func NewCreateTestCaseSectionRequest(name string, projectName string, ) *CreateTestCaseSectionRequest`

NewCreateTestCaseSectionRequest instantiates a new CreateTestCaseSectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestCaseSectionRequestWithDefaults

`func NewCreateTestCaseSectionRequestWithDefaults() *CreateTestCaseSectionRequest`

NewCreateTestCaseSectionRequestWithDefaults instantiates a new CreateTestCaseSectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTestCaseSectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTestCaseSectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTestCaseSectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *CreateTestCaseSectionRequest) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateTestCaseSectionRequest) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateTestCaseSectionRequest) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateTestCaseSectionRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProjectName

`func (o *CreateTestCaseSectionRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateTestCaseSectionRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateTestCaseSectionRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


