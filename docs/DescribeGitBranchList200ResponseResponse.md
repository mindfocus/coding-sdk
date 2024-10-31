# DescribeGitBranchList200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitBranchesData** | Pointer to [**GitBranchesData**](GitBranchesData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitBranchList200ResponseResponse

`func NewDescribeGitBranchList200ResponseResponse() *DescribeGitBranchList200ResponseResponse`

NewDescribeGitBranchList200ResponseResponse instantiates a new DescribeGitBranchList200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBranchList200ResponseResponseWithDefaults

`func NewDescribeGitBranchList200ResponseResponseWithDefaults() *DescribeGitBranchList200ResponseResponse`

NewDescribeGitBranchList200ResponseResponseWithDefaults instantiates a new DescribeGitBranchList200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitBranchesData

`func (o *DescribeGitBranchList200ResponseResponse) GetGitBranchesData() GitBranchesData`

GetGitBranchesData returns the GitBranchesData field if non-nil, zero value otherwise.

### GetGitBranchesDataOk

`func (o *DescribeGitBranchList200ResponseResponse) GetGitBranchesDataOk() (*GitBranchesData, bool)`

GetGitBranchesDataOk returns a tuple with the GitBranchesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitBranchesData

`func (o *DescribeGitBranchList200ResponseResponse) SetGitBranchesData(v GitBranchesData)`

SetGitBranchesData sets GitBranchesData field to given value.

### HasGitBranchesData

`func (o *DescribeGitBranchList200ResponseResponse) HasGitBranchesData() bool`

HasGitBranchesData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitBranchList200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitBranchList200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitBranchList200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitBranchList200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


