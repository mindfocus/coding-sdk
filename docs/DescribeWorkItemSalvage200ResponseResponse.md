# DescribeWorkItemSalvage200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpenApiIssueDetail** | Pointer to [**[]IssueDetail**](IssueDetail.md) | 事项关联的项目集工作项列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeWorkItemSalvage200ResponseResponse

`func NewDescribeWorkItemSalvage200ResponseResponse() *DescribeWorkItemSalvage200ResponseResponse`

NewDescribeWorkItemSalvage200ResponseResponse instantiates a new DescribeWorkItemSalvage200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeWorkItemSalvage200ResponseResponseWithDefaults

`func NewDescribeWorkItemSalvage200ResponseResponseWithDefaults() *DescribeWorkItemSalvage200ResponseResponse`

NewDescribeWorkItemSalvage200ResponseResponseWithDefaults instantiates a new DescribeWorkItemSalvage200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpenApiIssueDetail

`func (o *DescribeWorkItemSalvage200ResponseResponse) GetOpenApiIssueDetail() []IssueDetail`

GetOpenApiIssueDetail returns the OpenApiIssueDetail field if non-nil, zero value otherwise.

### GetOpenApiIssueDetailOk

`func (o *DescribeWorkItemSalvage200ResponseResponse) GetOpenApiIssueDetailOk() (*[]IssueDetail, bool)`

GetOpenApiIssueDetailOk returns a tuple with the OpenApiIssueDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenApiIssueDetail

`func (o *DescribeWorkItemSalvage200ResponseResponse) SetOpenApiIssueDetail(v []IssueDetail)`

SetOpenApiIssueDetail sets OpenApiIssueDetail field to given value.

### HasOpenApiIssueDetail

`func (o *DescribeWorkItemSalvage200ResponseResponse) HasOpenApiIssueDetail() bool`

HasOpenApiIssueDetail returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeWorkItemSalvage200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeWorkItemSalvage200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeWorkItemSalvage200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeWorkItemSalvage200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


