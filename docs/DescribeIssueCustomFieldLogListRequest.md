# DescribeIssueCustomFieldLogListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | 自定义属性名称 | 
**IssueCode** | **int64** | 事项编号 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDescribeIssueCustomFieldLogListRequest

`func NewDescribeIssueCustomFieldLogListRequest(fieldName string, issueCode int64, projectName string, ) *DescribeIssueCustomFieldLogListRequest`

NewDescribeIssueCustomFieldLogListRequest instantiates a new DescribeIssueCustomFieldLogListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueCustomFieldLogListRequestWithDefaults

`func NewDescribeIssueCustomFieldLogListRequestWithDefaults() *DescribeIssueCustomFieldLogListRequest`

NewDescribeIssueCustomFieldLogListRequestWithDefaults instantiates a new DescribeIssueCustomFieldLogListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *DescribeIssueCustomFieldLogListRequest) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *DescribeIssueCustomFieldLogListRequest) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *DescribeIssueCustomFieldLogListRequest) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetIssueCode

`func (o *DescribeIssueCustomFieldLogListRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DescribeIssueCustomFieldLogListRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DescribeIssueCustomFieldLogListRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetProjectName

`func (o *DescribeIssueCustomFieldLogListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeIssueCustomFieldLogListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeIssueCustomFieldLogListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


