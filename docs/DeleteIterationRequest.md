# DeleteIterationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IterationCode** | **int64** | 迭代编号 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDeleteIterationRequest

`func NewDeleteIterationRequest(iterationCode int64, projectName string, ) *DeleteIterationRequest`

NewDeleteIterationRequest instantiates a new DeleteIterationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIterationRequestWithDefaults

`func NewDeleteIterationRequestWithDefaults() *DeleteIterationRequest`

NewDeleteIterationRequestWithDefaults instantiates a new DeleteIterationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIterationCode

`func (o *DeleteIterationRequest) GetIterationCode() int64`

GetIterationCode returns the IterationCode field if non-nil, zero value otherwise.

### GetIterationCodeOk

`func (o *DeleteIterationRequest) GetIterationCodeOk() (*int64, bool)`

GetIterationCodeOk returns a tuple with the IterationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterationCode

`func (o *DeleteIterationRequest) SetIterationCode(v int64)`

SetIterationCode sets IterationCode field to given value.


### GetProjectName

`func (o *DeleteIterationRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteIterationRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteIterationRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


