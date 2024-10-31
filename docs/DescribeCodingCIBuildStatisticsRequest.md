# DescribeCodingCIBuildStatisticsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **string** | 结束时间 | 
**JobId** | Pointer to **int32** | MetricType 为 JOB 的时候使用该值，此时 ProjectId 可不传 | [optional] 
**MetricType** | **string** | 统计维度 Project 还是 JOB 目前只有 PROJECT | 
**Period** | **int32** | 统计间隔单位秒 | 
**ProjectId** | Pointer to **int32** | MetricType 为 PROJECT 的时候使用该值，此时 JobId 可不传 | [optional] 
**StartTime** | **string** | 开始时间 | 

## Methods

### NewDescribeCodingCIBuildStatisticsRequest

`func NewDescribeCodingCIBuildStatisticsRequest(endTime string, metricType string, period int32, startTime string, ) *DescribeCodingCIBuildStatisticsRequest`

NewDescribeCodingCIBuildStatisticsRequest instantiates a new DescribeCodingCIBuildStatisticsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildStatisticsRequestWithDefaults

`func NewDescribeCodingCIBuildStatisticsRequestWithDefaults() *DescribeCodingCIBuildStatisticsRequest`

NewDescribeCodingCIBuildStatisticsRequestWithDefaults instantiates a new DescribeCodingCIBuildStatisticsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *DescribeCodingCIBuildStatisticsRequest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DescribeCodingCIBuildStatisticsRequest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DescribeCodingCIBuildStatisticsRequest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetJobId

`func (o *DescribeCodingCIBuildStatisticsRequest) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DescribeCodingCIBuildStatisticsRequest) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DescribeCodingCIBuildStatisticsRequest) SetJobId(v int32)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *DescribeCodingCIBuildStatisticsRequest) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetMetricType

`func (o *DescribeCodingCIBuildStatisticsRequest) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *DescribeCodingCIBuildStatisticsRequest) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *DescribeCodingCIBuildStatisticsRequest) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.


### GetPeriod

`func (o *DescribeCodingCIBuildStatisticsRequest) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DescribeCodingCIBuildStatisticsRequest) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DescribeCodingCIBuildStatisticsRequest) SetPeriod(v int32)`

SetPeriod sets Period field to given value.


### GetProjectId

`func (o *DescribeCodingCIBuildStatisticsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeCodingCIBuildStatisticsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeCodingCIBuildStatisticsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *DescribeCodingCIBuildStatisticsRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetStartTime

`func (o *DescribeCodingCIBuildStatisticsRequest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DescribeCodingCIBuildStatisticsRequest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DescribeCodingCIBuildStatisticsRequest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


