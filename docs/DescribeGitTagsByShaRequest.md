# DescribeGitTagsByShaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**Sha** | **string** | 提交id | 

## Methods

### NewDescribeGitTagsByShaRequest

`func NewDescribeGitTagsByShaRequest(depotId int64, sha string, ) *DescribeGitTagsByShaRequest`

NewDescribeGitTagsByShaRequest instantiates a new DescribeGitTagsByShaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitTagsByShaRequestWithDefaults

`func NewDescribeGitTagsByShaRequestWithDefaults() *DescribeGitTagsByShaRequest`

NewDescribeGitTagsByShaRequestWithDefaults instantiates a new DescribeGitTagsByShaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitTagsByShaRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitTagsByShaRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitTagsByShaRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitTagsByShaRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitTagsByShaRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitTagsByShaRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitTagsByShaRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetSha

`func (o *DescribeGitTagsByShaRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DescribeGitTagsByShaRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DescribeGitTagsByShaRequest) SetSha(v string)`

SetSha sets Sha field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


