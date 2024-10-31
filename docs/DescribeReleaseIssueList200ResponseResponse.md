# DescribeReleaseIssueList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**IssueDetailListDataWithPage**](IssueDetailListDataWithPage.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeReleaseIssueList200ResponseResponse

`func NewDescribeReleaseIssueList200ResponseResponse() *DescribeReleaseIssueList200ResponseResponse`

NewDescribeReleaseIssueList200ResponseResponse instantiates a new DescribeReleaseIssueList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeReleaseIssueList200ResponseResponseWithDefaults

`func NewDescribeReleaseIssueList200ResponseResponseWithDefaults() *DescribeReleaseIssueList200ResponseResponse`

NewDescribeReleaseIssueList200ResponseResponseWithDefaults instantiates a new DescribeReleaseIssueList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeReleaseIssueList200ResponseResponse) GetData() IssueDetailListDataWithPage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeReleaseIssueList200ResponseResponse) GetDataOk() (*IssueDetailListDataWithPage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeReleaseIssueList200ResponseResponse) SetData(v IssueDetailListDataWithPage)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeReleaseIssueList200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeReleaseIssueList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeReleaseIssueList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeReleaseIssueList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeReleaseIssueList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


