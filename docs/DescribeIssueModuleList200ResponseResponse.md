# DescribeIssueModuleList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueModules** | Pointer to [**[]IssueModule**](IssueModule.md) | 模块列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeIssueModuleList200ResponseResponse

`func NewDescribeIssueModuleList200ResponseResponse() *DescribeIssueModuleList200ResponseResponse`

NewDescribeIssueModuleList200ResponseResponse instantiates a new DescribeIssueModuleList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeIssueModuleList200ResponseResponseWithDefaults

`func NewDescribeIssueModuleList200ResponseResponseWithDefaults() *DescribeIssueModuleList200ResponseResponse`

NewDescribeIssueModuleList200ResponseResponseWithDefaults instantiates a new DescribeIssueModuleList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueModules

`func (o *DescribeIssueModuleList200ResponseResponse) GetIssueModules() []IssueModule`

GetIssueModules returns the IssueModules field if non-nil, zero value otherwise.

### GetIssueModulesOk

`func (o *DescribeIssueModuleList200ResponseResponse) GetIssueModulesOk() (*[]IssueModule, bool)`

GetIssueModulesOk returns a tuple with the IssueModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueModules

`func (o *DescribeIssueModuleList200ResponseResponse) SetIssueModules(v []IssueModule)`

SetIssueModules sets IssueModules field to given value.

### HasIssueModules

`func (o *DescribeIssueModuleList200ResponseResponse) HasIssueModules() bool`

HasIssueModules returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeIssueModuleList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeIssueModuleList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeIssueModuleList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeIssueModuleList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


