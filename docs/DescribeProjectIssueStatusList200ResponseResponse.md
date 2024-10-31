# DescribeProjectIssueStatusList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIssueStatusList** | Pointer to [**[]ProjectIssueStatus**](ProjectIssueStatus.md) | 项目的事项状态列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProjectIssueStatusList200ResponseResponse

`func NewDescribeProjectIssueStatusList200ResponseResponse() *DescribeProjectIssueStatusList200ResponseResponse`

NewDescribeProjectIssueStatusList200ResponseResponse instantiates a new DescribeProjectIssueStatusList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectIssueStatusList200ResponseResponseWithDefaults

`func NewDescribeProjectIssueStatusList200ResponseResponseWithDefaults() *DescribeProjectIssueStatusList200ResponseResponse`

NewDescribeProjectIssueStatusList200ResponseResponseWithDefaults instantiates a new DescribeProjectIssueStatusList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIssueStatusList

`func (o *DescribeProjectIssueStatusList200ResponseResponse) GetProjectIssueStatusList() []ProjectIssueStatus`

GetProjectIssueStatusList returns the ProjectIssueStatusList field if non-nil, zero value otherwise.

### GetProjectIssueStatusListOk

`func (o *DescribeProjectIssueStatusList200ResponseResponse) GetProjectIssueStatusListOk() (*[]ProjectIssueStatus, bool)`

GetProjectIssueStatusListOk returns a tuple with the ProjectIssueStatusList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueStatusList

`func (o *DescribeProjectIssueStatusList200ResponseResponse) SetProjectIssueStatusList(v []ProjectIssueStatus)`

SetProjectIssueStatusList sets ProjectIssueStatusList field to given value.

### HasProjectIssueStatusList

`func (o *DescribeProjectIssueStatusList200ResponseResponse) HasProjectIssueStatusList() bool`

HasProjectIssueStatusList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProjectIssueStatusList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProjectIssueStatusList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProjectIssueStatusList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProjectIssueStatusList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


