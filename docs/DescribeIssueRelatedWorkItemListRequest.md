# DescribeIssueRelatedWorkItemListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项code | 
**ProjectName** | **string** | 项目名称 | 
**ShowImageOutUrl** | Pointer to **string** | 是否展示描述中图片外部可访问的地址 | [optional] 

## Methods

### NewDescribeIssueRelatedWorkItemListRequest

`func NewDescribeIssueRelatedWorkItemListRequest(issueCode int64, projectName string, ) *DescribeIssueRelatedWorkItemListRequest`

NewDescribeIssueRelatedWorkItemListRequest instantiates a new DescribeIssueRelatedWorkItemListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueRelatedWorkItemListRequestWithDefaults

`func NewDescribeIssueRelatedWorkItemListRequestWithDefaults() *DescribeIssueRelatedWorkItemListRequest`

NewDescribeIssueRelatedWorkItemListRequestWithDefaults instantiates a new DescribeIssueRelatedWorkItemListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *DescribeIssueRelatedWorkItemListRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DescribeIssueRelatedWorkItemListRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DescribeIssueRelatedWorkItemListRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *DescribeIssueRelatedWorkItemListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueRelatedWorkItemListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueRelatedWorkItemListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetShowImageOutUrl

`func (o *DescribeIssueRelatedWorkItemListRequest) GetShowImageOutUrl() string`

GetShowImageOutUrl returns the ShowImageOutUrl field if non-nil, zero value otherwise.

### GetShowImageOutUrlOk

`func (o *DescribeIssueRelatedWorkItemListRequest) GetShowImageOutUrlOk() (*string, bool)`

GetShowImageOutUrlOk returns a tuple with the ShowImageOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowImageOutUrl

`func (o *DescribeIssueRelatedWorkItemListRequest) SetShowImageOutUrl(v string)`

SetShowImageOutUrl sets ShowImageOutUrl field to given value.

### HasShowImageOutUrl

`func (o *DescribeIssueRelatedWorkItemListRequest) HasShowImageOutUrl() bool`

HasShowImageOutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


