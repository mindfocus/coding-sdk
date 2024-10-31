# CIBuildTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **int32** | 时长 | 
**Empty** | **bool** | 是否空 | [default to false]
**FailCount** | **int32** | 失败次数 | 
**PassCount** | **int32** | 通过次数 | 
**SkipCount** | **int32** | 跳过次数 | 
**TotalCount** | **int32** | 总次数 | 

## Methods

### NewCIBuildTestResult

`func NewCIBuildTestResult(duration int32, empty bool, failCount int32, passCount int32, skipCount int32, totalCount int32, ) *CIBuildTestResult`

NewCIBuildTestResult instantiates a new CIBuildTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCIBuildTestResultWithDefaults

`func NewCIBuildTestResultWithDefaults() *CIBuildTestResult`

NewCIBuildTestResultWithDefaults instantiates a new CIBuildTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *CIBuildTestResult) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CIBuildTestResult) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CIBuildTestResult) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetEmpty

`func (o *CIBuildTestResult) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *CIBuildTestResult) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *CIBuildTestResult) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.


### GetFailCount

`func (o *CIBuildTestResult) GetFailCount() int32`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *CIBuildTestResult) GetFailCountOk() (*int32, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *CIBuildTestResult) SetFailCount(v int32)`

SetFailCount sets FailCount field to given value.


### GetPassCount

`func (o *CIBuildTestResult) GetPassCount() int32`

GetPassCount returns the PassCount field if non-nil, zero value otherwise.

### GetPassCountOk

`func (o *CIBuildTestResult) GetPassCountOk() (*int32, bool)`

GetPassCountOk returns a tuple with the PassCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassCount

`func (o *CIBuildTestResult) SetPassCount(v int32)`

SetPassCount sets PassCount field to given value.


### GetSkipCount

`func (o *CIBuildTestResult) GetSkipCount() int32`

GetSkipCount returns the SkipCount field if non-nil, zero value otherwise.

### GetSkipCountOk

`func (o *CIBuildTestResult) GetSkipCountOk() (*int32, bool)`

GetSkipCountOk returns a tuple with the SkipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCount

`func (o *CIBuildTestResult) SetSkipCount(v int32)`

SetSkipCount sets SkipCount field to given value.


### GetTotalCount

`func (o *CIBuildTestResult) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CIBuildTestResult) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CIBuildTestResult) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


