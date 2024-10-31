# DescribeCodingCIJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**[]Filter**](Filter.md) | 过滤参数 | [optional] 
**ProjectId** | **int32** | 项目 ID | 

## Methods

### NewDescribeCodingCIJobsRequest

`func NewDescribeCodingCIJobsRequest(projectId int32, ) *DescribeCodingCIJobsRequest`

NewDescribeCodingCIJobsRequest instantiates a new DescribeCodingCIJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIJobsRequestWithDefaults

`func NewDescribeCodingCIJobsRequestWithDefaults() *DescribeCodingCIJobsRequest`

NewDescribeCodingCIJobsRequestWithDefaults instantiates a new DescribeCodingCIJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *DescribeCodingCIJobsRequest) GetFilter() []Filter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DescribeCodingCIJobsRequest) GetFilterOk() (*[]Filter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DescribeCodingCIJobsRequest) SetFilter(v []Filter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DescribeCodingCIJobsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetProjectId

`func (o *DescribeCodingCIJobsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeCodingCIJobsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeCodingCIJobsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


