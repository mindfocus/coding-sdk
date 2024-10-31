# DeleteAPIDocRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | API 文档编号 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDeleteAPIDocRequest

`func NewDeleteAPIDocRequest(code int32, projectName string, ) *DeleteAPIDocRequest`

NewDeleteAPIDocRequest instantiates a new DeleteAPIDocRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteAPIDocRequestWithDefaults

`func NewDeleteAPIDocRequestWithDefaults() *DeleteAPIDocRequest`

NewDeleteAPIDocRequestWithDefaults instantiates a new DeleteAPIDocRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DeleteAPIDocRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DeleteAPIDocRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DeleteAPIDocRequest) SetCode(v int32)`

SetCode sets Code field to given value.


### GetProjectName

`func (o *DeleteAPIDocRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteAPIDocRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteAPIDocRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


