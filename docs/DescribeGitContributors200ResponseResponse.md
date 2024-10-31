# DescribeGitContributors200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contributors** | Pointer to [**[]Contributor**](Contributor.md) | 贡献者列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitContributors200ResponseResponse

`func NewDescribeGitContributors200ResponseResponse() *DescribeGitContributors200ResponseResponse`

NewDescribeGitContributors200ResponseResponse instantiates a new DescribeGitContributors200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitContributors200ResponseResponseWithDefaults

`func NewDescribeGitContributors200ResponseResponseWithDefaults() *DescribeGitContributors200ResponseResponse`

NewDescribeGitContributors200ResponseResponseWithDefaults instantiates a new DescribeGitContributors200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContributors

`func (o *DescribeGitContributors200ResponseResponse) GetContributors() []Contributor`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *DescribeGitContributors200ResponseResponse) GetContributorsOk() (*[]Contributor, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *DescribeGitContributors200ResponseResponse) SetContributors(v []Contributor)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *DescribeGitContributors200ResponseResponse) HasContributors() bool`

HasContributors returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitContributors200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitContributors200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitContributors200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitContributors200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


