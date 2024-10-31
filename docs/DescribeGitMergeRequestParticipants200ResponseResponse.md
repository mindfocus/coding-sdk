# DescribeGitMergeRequestParticipants200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Participants** | Pointer to [**[]DepotUser**](DepotUser.md) | 参与者列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitMergeRequestParticipants200ResponseResponse

`func NewDescribeGitMergeRequestParticipants200ResponseResponse() *DescribeGitMergeRequestParticipants200ResponseResponse`

NewDescribeGitMergeRequestParticipants200ResponseResponse instantiates a new DescribeGitMergeRequestParticipants200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitMergeRequestParticipants200ResponseResponseWithDefaults

`func NewDescribeGitMergeRequestParticipants200ResponseResponseWithDefaults() *DescribeGitMergeRequestParticipants200ResponseResponse`

NewDescribeGitMergeRequestParticipants200ResponseResponseWithDefaults instantiates a new DescribeGitMergeRequestParticipants200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParticipants

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) GetParticipants() []DepotUser`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) GetParticipantsOk() (*[]DepotUser, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) SetParticipants(v []DepotUser)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitMergeRequestParticipants200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


