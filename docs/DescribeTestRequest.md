# DescribeTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**TestId** | Pointer to **int32** | 测试任务 ID | [optional] 

## Methods

### NewDescribeTestRequest

`func NewDescribeTestRequest(projectName string, ) *DescribeTestRequest`

NewDescribeTestRequest instantiates a new DescribeTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTestRequestWithDefaults

`func NewDescribeTestRequestWithDefaults() *DescribeTestRequest`

NewDescribeTestRequestWithDefaults instantiates a new DescribeTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DescribeTestRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeTestRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeTestRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTestId

`func (o *DescribeTestRequest) GetTestId() int32`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *DescribeTestRequest) GetTestIdOk() (*int32, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *DescribeTestRequest) SetTestId(v int32)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *DescribeTestRequest) HasTestId() bool`

HasTestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


