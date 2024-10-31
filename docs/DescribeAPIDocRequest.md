# DescribeAPIDocRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | API 文档编号 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeAPIDocRequest

`func NewDescribeAPIDocRequest(code int32, projectName string, ) *DescribeAPIDocRequest`

NewDescribeAPIDocRequest instantiates a new DescribeAPIDocRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAPIDocRequestWithDefaults

`func NewDescribeAPIDocRequestWithDefaults() *DescribeAPIDocRequest`

NewDescribeAPIDocRequestWithDefaults instantiates a new DescribeAPIDocRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DescribeAPIDocRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DescribeAPIDocRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DescribeAPIDocRequest) SetCode(v int32)`

SetCode sets Code field to given value.


### GetProjectName

`func (o *DescribeAPIDocRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeAPIDocRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeAPIDocRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


