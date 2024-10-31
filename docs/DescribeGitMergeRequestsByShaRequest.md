# DescribeGitMergeRequestsByShaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 与DepotId选择其一即可 | [optional] 
**Sha** | **string** | 提交 Id | 

## Methods

### NewDescribeGitMergeRequestsByShaRequest

`func NewDescribeGitMergeRequestsByShaRequest(depotId int64, sha string, ) *DescribeGitMergeRequestsByShaRequest`

NewDescribeGitMergeRequestsByShaRequest instantiates a new DescribeGitMergeRequestsByShaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestsByShaRequestWithDefaults

`func NewDescribeGitMergeRequestsByShaRequestWithDefaults() *DescribeGitMergeRequestsByShaRequest`

NewDescribeGitMergeRequestsByShaRequestWithDefaults instantiates a new DescribeGitMergeRequestsByShaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitMergeRequestsByShaRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitMergeRequestsByShaRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitMergeRequestsByShaRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitMergeRequestsByShaRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitMergeRequestsByShaRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitMergeRequestsByShaRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitMergeRequestsByShaRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetSha

`func (o *DescribeGitMergeRequestsByShaRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DescribeGitMergeRequestsByShaRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DescribeGitMergeRequestsByShaRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


