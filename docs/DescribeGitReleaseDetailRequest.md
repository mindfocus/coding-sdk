# DescribeGitReleaseDetailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**ReleaseId** | **int64** | 版本序号Id | 
**TagName** | **string** | 标签名称 | 

## Methods

### NewDescribeGitReleaseDetailRequest

`func NewDescribeGitReleaseDetailRequest(depotId int64, releaseId int64, tagName string, ) *DescribeGitReleaseDetailRequest`

NewDescribeGitReleaseDetailRequest instantiates a new DescribeGitReleaseDetailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitReleaseDetailRequestWithDefaults

`func NewDescribeGitReleaseDetailRequestWithDefaults() *DescribeGitReleaseDetailRequest`

NewDescribeGitReleaseDetailRequestWithDefaults instantiates a new DescribeGitReleaseDetailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitReleaseDetailRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitReleaseDetailRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitReleaseDetailRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitReleaseDetailRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitReleaseDetailRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitReleaseDetailRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitReleaseDetailRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetReleaseId

`func (o *DescribeGitReleaseDetailRequest) GetReleaseId() int64`

GetReleaseId returns the ReleaseId field if non-nil, zero value otherwise.

### GetReleaseIdOk

`func (o *DescribeGitReleaseDetailRequest) GetReleaseIdOk() (*int64, bool)`

GetReleaseIdOk returns a tuple with the ReleaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseId

`func (o *DescribeGitReleaseDetailRequest) SetReleaseId(v int64)`

SetReleaseId sets ReleaseId field to given value.


### GetTagName

`func (o *DescribeGitReleaseDetailRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *DescribeGitReleaseDetailRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *DescribeGitReleaseDetailRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


