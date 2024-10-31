# DescribeDifferentBetweenTwoCommitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Path** | **string** | 文件位置 | 
**Source** | **string** | 源请求 Sha 值,分支名称 | 
**Target** | **string** | 目标请求 Sha 值,分支名称 | 

## Methods

### NewDescribeDifferentBetweenTwoCommitsRequest

`func NewDescribeDifferentBetweenTwoCommitsRequest(depotId int64, path string, source string, target string, ) *DescribeDifferentBetweenTwoCommitsRequest`

NewDescribeDifferentBetweenTwoCommitsRequest instantiates a new DescribeDifferentBetweenTwoCommitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDifferentBetweenTwoCommitsRequestWithDefaults

`func NewDescribeDifferentBetweenTwoCommitsRequestWithDefaults() *DescribeDifferentBetweenTwoCommitsRequest`

NewDescribeDifferentBetweenTwoCommitsRequestWithDefaults instantiates a new DescribeDifferentBetweenTwoCommitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeDifferentBetweenTwoCommitsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeDifferentBetweenTwoCommitsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeDifferentBetweenTwoCommitsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetPath

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeDifferentBetweenTwoCommitsRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetSource

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DescribeDifferentBetweenTwoCommitsRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *DescribeDifferentBetweenTwoCommitsRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *DescribeDifferentBetweenTwoCommitsRequest) SetTarget(v string)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


