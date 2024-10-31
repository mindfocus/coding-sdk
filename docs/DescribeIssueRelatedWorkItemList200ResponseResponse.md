# DescribeIssueRelatedWorkItemList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkItemList** | Pointer to [**[]IssueDetail**](IssueDetail.md) | 事项关联的项目集中的工作项列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueRelatedWorkItemList200ResponseResponse

`func NewDescribeIssueRelatedWorkItemList200ResponseResponse() *DescribeIssueRelatedWorkItemList200ResponseResponse`

NewDescribeIssueRelatedWorkItemList200ResponseResponse instantiates a new DescribeIssueRelatedWorkItemList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueRelatedWorkItemList200ResponseResponseWithDefaults

`func NewDescribeIssueRelatedWorkItemList200ResponseResponseWithDefaults() *DescribeIssueRelatedWorkItemList200ResponseResponse`

NewDescribeIssueRelatedWorkItemList200ResponseResponseWithDefaults instantiates a new DescribeIssueRelatedWorkItemList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkItemList

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) GetWorkItemList() []IssueDetail`

GetWorkItemList returns the WorkItemList field if non-nil, zero value otherwise.

### GetWorkItemListOk

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) GetWorkItemListOk() (*[]IssueDetail, bool)`

GetWorkItemListOk returns a tuple with the WorkItemList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkItemList

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) SetWorkItemList(v []IssueDetail)`

SetWorkItemList sets WorkItemList field to given value.

### HasWorkItemList

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) HasWorkItemList() bool`

HasWorkItemList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueRelatedWorkItemList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


