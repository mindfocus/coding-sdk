# DescribeCanMergeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**Source** | **string** | 源分支 | 
**Target** | **string** | 目标分支 | 
**DepotPath** | Pointer to **string** | 仓库路径,与仓库Id二选一即可 | [optional] 

## Methods

### NewDescribeCanMergeRequest

`func NewDescribeCanMergeRequest(depotId int64, source string, target string, ) *DescribeCanMergeRequest`

NewDescribeCanMergeRequest instantiates a new DescribeCanMergeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCanMergeRequestWithDefaults

`func NewDescribeCanMergeRequestWithDefaults() *DescribeCanMergeRequest`

NewDescribeCanMergeRequestWithDefaults instantiates a new DescribeCanMergeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeCanMergeRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeCanMergeRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeCanMergeRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetSource

`func (o *DescribeCanMergeRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DescribeCanMergeRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DescribeCanMergeRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *DescribeCanMergeRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DescribeCanMergeRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DescribeCanMergeRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetDepotPath

`func (o *DescribeCanMergeRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeCanMergeRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeCanMergeRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeCanMergeRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


