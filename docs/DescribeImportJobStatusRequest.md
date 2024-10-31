# DescribeImportJobStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **string** | 任务Id | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeImportJobStatusRequest

`func NewDescribeImportJobStatusRequest(jobId string, projectName string, ) *DescribeImportJobStatusRequest`

NewDescribeImportJobStatusRequest instantiates a new DescribeImportJobStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeImportJobStatusRequestWithDefaults

`func NewDescribeImportJobStatusRequestWithDefaults() *DescribeImportJobStatusRequest`

NewDescribeImportJobStatusRequestWithDefaults instantiates a new DescribeImportJobStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *DescribeImportJobStatusRequest) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DescribeImportJobStatusRequest) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DescribeImportJobStatusRequest) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetProjectName

`func (o *DescribeImportJobStatusRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeImportJobStatusRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeImportJobStatusRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


