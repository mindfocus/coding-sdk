# CreateMergeRequestReviewer200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reviewers** | Pointer to [**[]CodingUser**](CodingUser.md) | 评审者信息 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewCreateMergeRequestReviewer200ResponseResponse

`func NewCreateMergeRequestReviewer200ResponseResponse() *CreateMergeRequestReviewer200ResponseResponse`

NewCreateMergeRequestReviewer200ResponseResponse instantiates a new CreateMergeRequestReviewer200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMergeRequestReviewer200ResponseResponseWithDefaults

`func NewCreateMergeRequestReviewer200ResponseResponseWithDefaults() *CreateMergeRequestReviewer200ResponseResponse`

NewCreateMergeRequestReviewer200ResponseResponseWithDefaults instantiates a new CreateMergeRequestReviewer200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReviewers

`func (o *CreateMergeRequestReviewer200ResponseResponse) GetReviewers() []CodingUser`

GetReviewers returns the Reviewers field if non-nil, zero value otherwise.

### GetReviewersOk

`func (o *CreateMergeRequestReviewer200ResponseResponse) GetReviewersOk() (*[]CodingUser, bool)`

GetReviewersOk returns a tuple with the Reviewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewers

`func (o *CreateMergeRequestReviewer200ResponseResponse) SetReviewers(v []CodingUser)`

SetReviewers sets Reviewers field to given value.

### HasReviewers

`func (o *CreateMergeRequestReviewer200ResponseResponse) HasReviewers() bool`

HasReviewers returns a boolean if a field has been set.

### GetRequestId

`func (o *CreateMergeRequestReviewer200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateMergeRequestReviewer200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateMergeRequestReviewer200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateMergeRequestReviewer200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


