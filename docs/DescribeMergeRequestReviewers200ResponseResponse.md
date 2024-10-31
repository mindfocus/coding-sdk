# DescribeMergeRequestReviewers200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviewers** | Pointer to [**[]DepotUser**](DepotUser.md) | 参与者列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeMergeRequestReviewers200ResponseResponse

`func NewDescribeMergeRequestReviewers200ResponseResponse() *DescribeMergeRequestReviewers200ResponseResponse`

NewDescribeMergeRequestReviewers200ResponseResponse instantiates a new DescribeMergeRequestReviewers200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeMergeRequestReviewers200ResponseResponseWithDefaults

`func NewDescribeMergeRequestReviewers200ResponseResponseWithDefaults() *DescribeMergeRequestReviewers200ResponseResponse`

NewDescribeMergeRequestReviewers200ResponseResponseWithDefaults instantiates a new DescribeMergeRequestReviewers200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewers

`func (o *DescribeMergeRequestReviewers200ResponseResponse) GetReviewers() []DepotUser`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *DescribeMergeRequestReviewers200ResponseResponse) GetReviewersOk() (*[]DepotUser, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *DescribeMergeRequestReviewers200ResponseResponse) SetReviewers(v []DepotUser)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *DescribeMergeRequestReviewers200ResponseResponse) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeMergeRequestReviewers200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeMergeRequestReviewers200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeMergeRequestReviewers200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeMergeRequestReviewers200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


