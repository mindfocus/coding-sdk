# DescribeCodingCIBuildMetricsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **string** | 结束时间 2020-11-20 | 
**ProjectId** | Pointer to **int32** | 项目 ID,Type 为 PROJECT 时必填 | [optional] 
**StartTime** | **string** | 开始时间 2020-11-10 | 
**Type** | Pointer to **string** | 查询级别 PROJECT 级别，TEAM 级别，目前只支持 PROJECT | [optional] 

## Methods

### NewDescribeCodingCIBuildMetricsRequest

`func NewDescribeCodingCIBuildMetricsRequest(endTime string, startTime string, ) *DescribeCodingCIBuildMetricsRequest`

NewDescribeCodingCIBuildMetricsRequest instantiates a new DescribeCodingCIBuildMetricsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildMetricsRequestWithDefaults

`func NewDescribeCodingCIBuildMetricsRequestWithDefaults() *DescribeCodingCIBuildMetricsRequest`

NewDescribeCodingCIBuildMetricsRequestWithDefaults instantiates a new DescribeCodingCIBuildMetricsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DescribeCodingCIBuildMetricsRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DescribeCodingCIBuildMetricsRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DescribeCodingCIBuildMetricsRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetProjectId

`func (o *DescribeCodingCIBuildMetricsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeCodingCIBuildMetricsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeCodingCIBuildMetricsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DescribeCodingCIBuildMetricsRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStartTime

`func (o *DescribeCodingCIBuildMetricsRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DescribeCodingCIBuildMetricsRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DescribeCodingCIBuildMetricsRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetType

`func (o *DescribeCodingCIBuildMetricsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DescribeCodingCIBuildMetricsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DescribeCodingCIBuildMetricsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DescribeCodingCIBuildMetricsRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


