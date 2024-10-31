# DescribeCodingCIJobs200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobList** | Pointer to [**[]CodingCIJob**](CodingCIJob.md) | CI 任务列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCodingCIJobs200ResponseResponse

`func NewDescribeCodingCIJobs200ResponseResponse() *DescribeCodingCIJobs200ResponseResponse`

NewDescribeCodingCIJobs200ResponseResponse instantiates a new DescribeCodingCIJobs200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIJobs200ResponseResponseWithDefaults

`func NewDescribeCodingCIJobs200ResponseResponseWithDefaults() *DescribeCodingCIJobs200ResponseResponse`

NewDescribeCodingCIJobs200ResponseResponseWithDefaults instantiates a new DescribeCodingCIJobs200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobList

`func (o *DescribeCodingCIJobs200ResponseResponse) GetJobList() []CodingCIJob`

GetJobList returns the JobList field if non-nil, zero value otherwise.

### GetJobListOk

`func (o *DescribeCodingCIJobs200ResponseResponse) GetJobListOk() (*[]CodingCIJob, bool)`

GetJobListOk returns a tuple with the JobList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobList

`func (o *DescribeCodingCIJobs200ResponseResponse) SetJobList(v []CodingCIJob)`

SetJobList sets JobList field to given value.

### HasJobList

`func (o *DescribeCodingCIJobs200ResponseResponse) HasJobList() bool`

HasJobList returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeCodingCIJobs200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCodingCIJobs200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCodingCIJobs200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCodingCIJobs200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


