# DescribeCodingCIBuildMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** | 日期 | [optional] [default to ""]
**SuccessBuildCount** | Pointer to **int32** | 构建成功总数 | [optional] 
**TotalBuildCount** | Pointer to **int32** | 构建总数 | [optional] 
**TotalDuration** | Pointer to **int32** | 构建总耗时，单位毫秒 | [optional] 

## Methods

### NewDescribeCodingCIBuildMetrics

`func NewDescribeCodingCIBuildMetrics() *DescribeCodingCIBuildMetrics`

NewDescribeCodingCIBuildMetrics instantiates a new DescribeCodingCIBuildMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildMetricsWithDefaults

`func NewDescribeCodingCIBuildMetricsWithDefaults() *DescribeCodingCIBuildMetrics`

NewDescribeCodingCIBuildMetricsWithDefaults instantiates a new DescribeCodingCIBuildMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DescribeCodingCIBuildMetrics) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DescribeCodingCIBuildMetrics) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DescribeCodingCIBuildMetrics) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DescribeCodingCIBuildMetrics) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *DescribeCodingCIBuildMetrics) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *DescribeCodingCIBuildMetrics) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetSuccessBuildCount

`func (o *DescribeCodingCIBuildMetrics) GetSuccessBuildCount() int32`

GetSuccessBuildCount returns the SuccessBuildCount field if non-nil, zero value otherwise.

### GetSuccessBuildCountOk

`func (o *DescribeCodingCIBuildMetrics) GetSuccessBuildCountOk() (*int32, bool)`

GetSuccessBuildCountOk returns a tuple with the SuccessBuildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessBuildCount

`func (o *DescribeCodingCIBuildMetrics) SetSuccessBuildCount(v int32)`

SetSuccessBuildCount sets SuccessBuildCount field to given value.

### HasSuccessBuildCount

`func (o *DescribeCodingCIBuildMetrics) HasSuccessBuildCount() bool`

HasSuccessBuildCount returns a boolean if a field has been set.

### GetTotalBuildCount

`func (o *DescribeCodingCIBuildMetrics) GetTotalBuildCount() int32`

GetTotalBuildCount returns the TotalBuildCount field if non-nil, zero value otherwise.

### GetTotalBuildCountOk

`func (o *DescribeCodingCIBuildMetrics) GetTotalBuildCountOk() (*int32, bool)`

GetTotalBuildCountOk returns a tuple with the TotalBuildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBuildCount

`func (o *DescribeCodingCIBuildMetrics) SetTotalBuildCount(v int32)`

SetTotalBuildCount sets TotalBuildCount field to given value.

### HasTotalBuildCount

`func (o *DescribeCodingCIBuildMetrics) HasTotalBuildCount() bool`

HasTotalBuildCount returns a boolean if a field has been set.

### GetTotalDuration

`func (o *DescribeCodingCIBuildMetrics) GetTotalDuration() int32`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *DescribeCodingCIBuildMetrics) GetTotalDurationOk() (*int32, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *DescribeCodingCIBuildMetrics) SetTotalDuration(v int32)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *DescribeCodingCIBuildMetrics) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


