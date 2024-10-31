# DescribeCodingCIBuildsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **int32** | 构建计划 ID | 
**PageNumber** | **int32** | 页码 | 
**PageSize** | **int32** | 每页条数 | 

## Methods

### NewDescribeCodingCIBuildsRequest

`func NewDescribeCodingCIBuildsRequest(jobId int32, pageNumber int32, pageSize int32, ) *DescribeCodingCIBuildsRequest`

NewDescribeCodingCIBuildsRequest instantiates a new DescribeCodingCIBuildsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildsRequestWithDefaults

`func NewDescribeCodingCIBuildsRequestWithDefaults() *DescribeCodingCIBuildsRequest`

NewDescribeCodingCIBuildsRequestWithDefaults instantiates a new DescribeCodingCIBuildsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *DescribeCodingCIBuildsRequest) GetJobId() int32`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *DescribeCodingCIBuildsRequest) GetJobIdOk() (*int32, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *DescribeCodingCIBuildsRequest) SetJobId(v int32)`

SetJobId sets JobId field to given value.


### GetPageNumber

`func (o *DescribeCodingCIBuildsRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeCodingCIBuildsRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeCodingCIBuildsRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeCodingCIBuildsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeCodingCIBuildsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeCodingCIBuildsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


