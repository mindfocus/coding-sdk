# DescribeIssueReferenceResourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 查询的目标事项的code | 
**ProjectId** | **int64** | 查询的事项所在的项目ID | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeIssueReferenceResourcesRequest

`func NewDescribeIssueReferenceResourcesRequest(issueCode int64, projectId int64, projectName string, ) *DescribeIssueReferenceResourcesRequest`

NewDescribeIssueReferenceResourcesRequest instantiates a new DescribeIssueReferenceResourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueReferenceResourcesRequestWithDefaults

`func NewDescribeIssueReferenceResourcesRequestWithDefaults() *DescribeIssueReferenceResourcesRequest`

NewDescribeIssueReferenceResourcesRequestWithDefaults instantiates a new DescribeIssueReferenceResourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *DescribeIssueReferenceResourcesRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DescribeIssueReferenceResourcesRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DescribeIssueReferenceResourcesRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectId

`func (o *DescribeIssueReferenceResourcesRequest) GetProjectId() int64`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeIssueReferenceResourcesRequest) GetProjectIdOk() (*int64, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeIssueReferenceResourcesRequest) SetProjectId(v int64)`

SetProjectId sets ProjectId field to given value.


### GetProjectName

`func (o *DescribeIssueReferenceResourcesRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueReferenceResourcesRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueReferenceResourcesRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


