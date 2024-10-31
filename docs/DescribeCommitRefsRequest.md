# DescribeCommitRefsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径 | 
**Sha** | **string** | 提交sha | 
**Type** | **string** | ref类型,all 查询分支和标签，branch查询分支，tag查询标签 | 

## Methods

### NewDescribeCommitRefsRequest

`func NewDescribeCommitRefsRequest(depotPath string, sha string, type_ string, ) *DescribeCommitRefsRequest`

NewDescribeCommitRefsRequest instantiates a new DescribeCommitRefsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCommitRefsRequestWithDefaults

`func NewDescribeCommitRefsRequestWithDefaults() *DescribeCommitRefsRequest`

NewDescribeCommitRefsRequestWithDefaults instantiates a new DescribeCommitRefsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DescribeCommitRefsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeCommitRefsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeCommitRefsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetSha

`func (o *DescribeCommitRefsRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DescribeCommitRefsRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DescribeCommitRefsRequest) SetSha(v string)`

SetSha sets Sha field to given value.


### GetType

`func (o *DescribeCommitRefsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeCommitRefsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeCommitRefsRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


