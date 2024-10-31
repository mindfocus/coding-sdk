# DescribeDifferentBetween2CommitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库Id | [default to 0]
**DepotPath** | Pointer to **string** | 仓库路径，与DepotID二选一即可 | [optional] [default to "codingcorp/test/depot"]
**Path** | **string** | 文件路径 | [default to "README.md"]
**Source** | **string** | 源分支 | [default to "master"]
**Straight** | Pointer to **bool** | commit对比方式,是否直接对比差异 | [optional] [default to false]
**Target** | **string** | 目标分支 | [default to "dev"]

## Methods

### NewDescribeDifferentBetween2CommitsRequest

`func NewDescribeDifferentBetween2CommitsRequest(depotId int64, path string, source string, target string, ) *DescribeDifferentBetween2CommitsRequest`

NewDescribeDifferentBetween2CommitsRequest instantiates a new DescribeDifferentBetween2CommitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDifferentBetween2CommitsRequestWithDefaults

`func NewDescribeDifferentBetween2CommitsRequestWithDefaults() *DescribeDifferentBetween2CommitsRequest`

NewDescribeDifferentBetween2CommitsRequestWithDefaults instantiates a new DescribeDifferentBetween2CommitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeDifferentBetween2CommitsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeDifferentBetween2CommitsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeDifferentBetween2CommitsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeDifferentBetween2CommitsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeDifferentBetween2CommitsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeDifferentBetween2CommitsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeDifferentBetween2CommitsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetPath

`func (o *DescribeDifferentBetween2CommitsRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeDifferentBetween2CommitsRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeDifferentBetween2CommitsRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetSource

`func (o *DescribeDifferentBetween2CommitsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DescribeDifferentBetween2CommitsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DescribeDifferentBetween2CommitsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetStraight

`func (o *DescribeDifferentBetween2CommitsRequest) GetStraight() bool`

GetStraight returns the Straight field if non-nil, zero value otherwise.

### GetStraightOk

`func (o *DescribeDifferentBetween2CommitsRequest) GetStraightOk() (*bool, bool)`

GetStraightOk returns a tuple with the Straight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraight

`func (o *DescribeDifferentBetween2CommitsRequest) SetStraight(v bool)`

SetStraight sets Straight field to given value.

### HasStraight

`func (o *DescribeDifferentBetween2CommitsRequest) HasStraight() bool`

HasStraight returns a boolean if a field has been set.

### GetTarget

`func (o *DescribeDifferentBetween2CommitsRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DescribeDifferentBetween2CommitsRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DescribeDifferentBetween2CommitsRequest) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


