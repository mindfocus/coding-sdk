# DescribeTeamIssueTypeList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypes** | Pointer to [**[]IssueTypeDetail**](IssueTypeDetail.md) | 事项类型列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeTeamIssueTypeList200ResponseResponse

`func NewDescribeTeamIssueTypeList200ResponseResponse() *DescribeTeamIssueTypeList200ResponseResponse`

NewDescribeTeamIssueTypeList200ResponseResponse instantiates a new DescribeTeamIssueTypeList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamIssueTypeList200ResponseResponseWithDefaults

`func NewDescribeTeamIssueTypeList200ResponseResponseWithDefaults() *DescribeTeamIssueTypeList200ResponseResponse`

NewDescribeTeamIssueTypeList200ResponseResponseWithDefaults instantiates a new DescribeTeamIssueTypeList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypes

`func (o *DescribeTeamIssueTypeList200ResponseResponse) GetIssueTypes() []IssueTypeDetail`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *DescribeTeamIssueTypeList200ResponseResponse) GetIssueTypesOk() (*[]IssueTypeDetail, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *DescribeTeamIssueTypeList200ResponseResponse) SetIssueTypes(v []IssueTypeDetail)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *DescribeTeamIssueTypeList200ResponseResponse) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeTeamIssueTypeList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeTeamIssueTypeList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeTeamIssueTypeList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeTeamIssueTypeList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


