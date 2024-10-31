# DescribeReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名称 | 
**ReleaseCode** | **int64** | 版本code | 

## Methods

### NewDescribeReleaseRequest

`func NewDescribeReleaseRequest(projectName string, releaseCode int64, ) *DescribeReleaseRequest`

NewDescribeReleaseRequest instantiates a new DescribeReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeReleaseRequestWithDefaults

`func NewDescribeReleaseRequestWithDefaults() *DescribeReleaseRequest`

NewDescribeReleaseRequestWithDefaults instantiates a new DescribeReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *DescribeReleaseRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeReleaseRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeReleaseRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetReleaseCode

`func (o *DescribeReleaseRequest) GetReleaseCode() int64`

GetReleaseCode returns the ReleaseCode field if non-nil, zero value otherwise.

### GetReleaseCodeOk

`func (o *DescribeReleaseRequest) GetReleaseCodeOk() (*int64, bool)`

GetReleaseCodeOk returns a tuple with the ReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCode

`func (o *DescribeReleaseRequest) SetReleaseCode(v int64)`

SetReleaseCode sets ReleaseCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


