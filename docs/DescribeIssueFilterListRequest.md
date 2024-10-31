# DescribeIssueFilterListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueType** | **string** | 事项类型  ALL - 全部事项  DEFECT - 缺陷  REQUIREMENT - 需求  MISSION - 任务  EPIC - 史诗 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeIssueFilterListRequest

`func NewDescribeIssueFilterListRequest(issueType string, projectName string, ) *DescribeIssueFilterListRequest`

NewDescribeIssueFilterListRequest instantiates a new DescribeIssueFilterListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueFilterListRequestWithDefaults

`func NewDescribeIssueFilterListRequestWithDefaults() *DescribeIssueFilterListRequest`

NewDescribeIssueFilterListRequestWithDefaults instantiates a new DescribeIssueFilterListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueType

`func (o *DescribeIssueFilterListRequest) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *DescribeIssueFilterListRequest) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *DescribeIssueFilterListRequest) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.


### GetProjectName

`func (o *DescribeIssueFilterListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueFilterListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueFilterListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


