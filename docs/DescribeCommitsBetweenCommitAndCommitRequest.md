# DescribeCommitsBetweenCommitAndCommitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，选择其一即可 | [optional] 
**Source** | **string** | 查询起点sha值 | 
**Target** | **string** | 查询目标sha值 | 

## Methods

### NewDescribeCommitsBetweenCommitAndCommitRequest

`func NewDescribeCommitsBetweenCommitAndCommitRequest(depotId int64, source string, target string, ) *DescribeCommitsBetweenCommitAndCommitRequest`

NewDescribeCommitsBetweenCommitAndCommitRequest instantiates a new DescribeCommitsBetweenCommitAndCommitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCommitsBetweenCommitAndCommitRequestWithDefaults

`func NewDescribeCommitsBetweenCommitAndCommitRequestWithDefaults() *DescribeCommitsBetweenCommitAndCommitRequest`

NewDescribeCommitsBetweenCommitAndCommitRequestWithDefaults instantiates a new DescribeCommitsBetweenCommitAndCommitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetSource

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DescribeCommitsBetweenCommitAndCommitRequest) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


