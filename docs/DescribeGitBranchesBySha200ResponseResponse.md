# DescribeGitBranchesBySha200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refs** | Pointer to [**[]RefInfo**](RefInfo.md) | ref信息列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitBranchesBySha200ResponseResponse

`func NewDescribeGitBranchesBySha200ResponseResponse() *DescribeGitBranchesBySha200ResponseResponse`

NewDescribeGitBranchesBySha200ResponseResponse instantiates a new DescribeGitBranchesBySha200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBranchesBySha200ResponseResponseWithDefaults

`func NewDescribeGitBranchesBySha200ResponseResponseWithDefaults() *DescribeGitBranchesBySha200ResponseResponse`

NewDescribeGitBranchesBySha200ResponseResponseWithDefaults instantiates a new DescribeGitBranchesBySha200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefs

`func (o *DescribeGitBranchesBySha200ResponseResponse) GetRefs() []RefInfo`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *DescribeGitBranchesBySha200ResponseResponse) GetRefsOk() (*[]RefInfo, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *DescribeGitBranchesBySha200ResponseResponse) SetRefs(v []RefInfo)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *DescribeGitBranchesBySha200ResponseResponse) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitBranchesBySha200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitBranchesBySha200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitBranchesBySha200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitBranchesBySha200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


