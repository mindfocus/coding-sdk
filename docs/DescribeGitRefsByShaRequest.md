# DescribeGitRefsByShaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Sha** | **string** | 提交id | 
**Type** | **string** | all：列表中既包含分支也包含标签，tag：列表中只包含标签，branch：列表中只包含分支 | 

## Methods

### NewDescribeGitRefsByShaRequest

`func NewDescribeGitRefsByShaRequest(depotId int64, sha string, type_ string, ) *DescribeGitRefsByShaRequest`

NewDescribeGitRefsByShaRequest instantiates a new DescribeGitRefsByShaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitRefsByShaRequestWithDefaults

`func NewDescribeGitRefsByShaRequestWithDefaults() *DescribeGitRefsByShaRequest`

NewDescribeGitRefsByShaRequestWithDefaults instantiates a new DescribeGitRefsByShaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitRefsByShaRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitRefsByShaRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitRefsByShaRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitRefsByShaRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitRefsByShaRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitRefsByShaRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitRefsByShaRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetSha

`func (o *DescribeGitRefsByShaRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *DescribeGitRefsByShaRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *DescribeGitRefsByShaRequest) SetSha(v string)`

SetSha sets Sha field to given value.


### GetType

`func (o *DescribeGitRefsByShaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeGitRefsByShaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeGitRefsByShaRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


