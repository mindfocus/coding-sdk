# DescribeRelatedCaseListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项ID | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeRelatedCaseListRequest

`func NewDescribeRelatedCaseListRequest(issueCode int64, projectName string, ) *DescribeRelatedCaseListRequest`

NewDescribeRelatedCaseListRequest instantiates a new DescribeRelatedCaseListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeRelatedCaseListRequestWithDefaults

`func NewDescribeRelatedCaseListRequestWithDefaults() *DescribeRelatedCaseListRequest`

NewDescribeRelatedCaseListRequestWithDefaults instantiates a new DescribeRelatedCaseListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *DescribeRelatedCaseListRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DescribeRelatedCaseListRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DescribeRelatedCaseListRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *DescribeRelatedCaseListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeRelatedCaseListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeRelatedCaseListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


