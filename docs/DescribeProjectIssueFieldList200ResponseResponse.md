# DescribeProjectIssueFieldList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIssueFieldList** | Pointer to [**[]ProjectIssueField**](ProjectIssueField.md) | 项目的事项属性列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeProjectIssueFieldList200ResponseResponse

`func NewDescribeProjectIssueFieldList200ResponseResponse() *DescribeProjectIssueFieldList200ResponseResponse`

NewDescribeProjectIssueFieldList200ResponseResponse instantiates a new DescribeProjectIssueFieldList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectIssueFieldList200ResponseResponseWithDefaults

`func NewDescribeProjectIssueFieldList200ResponseResponseWithDefaults() *DescribeProjectIssueFieldList200ResponseResponse`

NewDescribeProjectIssueFieldList200ResponseResponseWithDefaults instantiates a new DescribeProjectIssueFieldList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIssueFieldList

`func (o *DescribeProjectIssueFieldList200ResponseResponse) GetProjectIssueFieldList() []ProjectIssueField`

GetProjectIssueFieldList returns the ProjectIssueFieldList field if non-nil, zero value otherwise.

### GetProjectIssueFieldListOk

`func (o *DescribeProjectIssueFieldList200ResponseResponse) GetProjectIssueFieldListOk() (*[]ProjectIssueField, bool)`

GetProjectIssueFieldListOk returns a tuple with the ProjectIssueFieldList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueFieldList

`func (o *DescribeProjectIssueFieldList200ResponseResponse) SetProjectIssueFieldList(v []ProjectIssueField)`

SetProjectIssueFieldList sets ProjectIssueFieldList field to given value.

### HasProjectIssueFieldList

`func (o *DescribeProjectIssueFieldList200ResponseResponse) HasProjectIssueFieldList() bool`

HasProjectIssueFieldList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeProjectIssueFieldList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeProjectIssueFieldList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeProjectIssueFieldList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeProjectIssueFieldList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


