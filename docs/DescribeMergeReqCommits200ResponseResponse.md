# DescribeMergeReqCommits200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]GitCommit**](GitCommit.md) | commit集合 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeMergeReqCommits200ResponseResponse

`func NewDescribeMergeReqCommits200ResponseResponse() *DescribeMergeReqCommits200ResponseResponse`

NewDescribeMergeReqCommits200ResponseResponse instantiates a new DescribeMergeReqCommits200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeReqCommits200ResponseResponseWithDefaults

`func NewDescribeMergeReqCommits200ResponseResponseWithDefaults() *DescribeMergeReqCommits200ResponseResponse`

NewDescribeMergeReqCommits200ResponseResponseWithDefaults instantiates a new DescribeMergeReqCommits200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *DescribeMergeReqCommits200ResponseResponse) GetCommits() []GitCommit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DescribeMergeReqCommits200ResponseResponse) GetCommitsOk() (*[]GitCommit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DescribeMergeReqCommits200ResponseResponse) SetCommits(v []GitCommit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *DescribeMergeReqCommits200ResponseResponse) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeMergeReqCommits200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeMergeReqCommits200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeMergeReqCommits200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeMergeReqCommits200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


