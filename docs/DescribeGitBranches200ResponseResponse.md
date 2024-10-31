# DescribeGitBranches200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branches** | Pointer to [**[]GitBranchInfo**](GitBranchInfo.md) | 分支详细信息 | [optional] 
**TotalCount** | Pointer to **NullableInt64** | 分支数量 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitBranches200ResponseResponse

`func NewDescribeGitBranches200ResponseResponse() *DescribeGitBranches200ResponseResponse`

NewDescribeGitBranches200ResponseResponse instantiates a new DescribeGitBranches200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBranches200ResponseResponseWithDefaults

`func NewDescribeGitBranches200ResponseResponseWithDefaults() *DescribeGitBranches200ResponseResponse`

NewDescribeGitBranches200ResponseResponseWithDefaults instantiates a new DescribeGitBranches200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranches

`func (o *DescribeGitBranches200ResponseResponse) GetBranches() []GitBranchInfo`

GetBranches returns the Branches field if non-nil, zero value otherwise.

### GetBranchesOk

`func (o *DescribeGitBranches200ResponseResponse) GetBranchesOk() (*[]GitBranchInfo, bool)`

GetBranchesOk returns a tuple with the Branches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranches

`func (o *DescribeGitBranches200ResponseResponse) SetBranches(v []GitBranchInfo)`

SetBranches sets Branches field to given value.

### HasBranches

`func (o *DescribeGitBranches200ResponseResponse) HasBranches() bool`

HasBranches returns a boolean if a field has been set.

### SetBranchesNil

`func (o *DescribeGitBranches200ResponseResponse) SetBranchesNil(b bool)`

 SetBranchesNil sets the value for Branches to be an explicit nil

### UnsetBranches
`func (o *DescribeGitBranches200ResponseResponse) UnsetBranches()`

UnsetBranches ensures that no value is present for Branches, not even an explicit nil
### GetTotalCount

`func (o *DescribeGitBranches200ResponseResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeGitBranches200ResponseResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeGitBranches200ResponseResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DescribeGitBranches200ResponseResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### SetTotalCountNil

`func (o *DescribeGitBranches200ResponseResponse) SetTotalCountNil(b bool)`

 SetTotalCountNil sets the value for TotalCount to be an explicit nil

### UnsetTotalCount
`func (o *DescribeGitBranches200ResponseResponse) UnsetTotalCount()`

UnsetTotalCount ensures that no value is present for TotalCount, not even an explicit nil
### GetRequestId

`func (o *DescribeGitBranches200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitBranches200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitBranches200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitBranches200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


