# DescribeWorkItemSalvageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项code | 
**ProjectName** | **string** | 项目名称 | 
**ShowImageOutUrl** | Pointer to **bool** | 事项展示图片对外路径 | [optional] 

## Methods

### NewDescribeWorkItemSalvageRequest

`func NewDescribeWorkItemSalvageRequest(issueCode int64, projectName string, ) *DescribeWorkItemSalvageRequest`

NewDescribeWorkItemSalvageRequest instantiates a new DescribeWorkItemSalvageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWorkItemSalvageRequestWithDefaults

`func NewDescribeWorkItemSalvageRequestWithDefaults() *DescribeWorkItemSalvageRequest`

NewDescribeWorkItemSalvageRequestWithDefaults instantiates a new DescribeWorkItemSalvageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *DescribeWorkItemSalvageRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DescribeWorkItemSalvageRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DescribeWorkItemSalvageRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *DescribeWorkItemSalvageRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeWorkItemSalvageRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeWorkItemSalvageRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetShowImageOutUrl

`func (o *DescribeWorkItemSalvageRequest) GetShowImageOutUrl() bool`

GetShowImageOutUrl returns the ShowImageOutUrl field if non-nil, zero value otherwise.

### GetShowImageOutUrlOk

`func (o *DescribeWorkItemSalvageRequest) GetShowImageOutUrlOk() (*bool, bool)`

GetShowImageOutUrlOk returns a tuple with the ShowImageOutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowImageOutUrl

`func (o *DescribeWorkItemSalvageRequest) SetShowImageOutUrl(v bool)`

SetShowImageOutUrl sets ShowImageOutUrl field to given value.

### HasShowImageOutUrl

`func (o *DescribeWorkItemSalvageRequest) HasShowImageOutUrl() bool`

HasShowImageOutUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


