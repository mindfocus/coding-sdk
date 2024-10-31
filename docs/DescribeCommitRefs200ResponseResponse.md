# DescribeCommitRefs200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitRefs** | Pointer to [**[]CommitRef**](CommitRef.md) | 仓库信息详情列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeCommitRefs200ResponseResponse

`func NewDescribeCommitRefs200ResponseResponse() *DescribeCommitRefs200ResponseResponse`

NewDescribeCommitRefs200ResponseResponse instantiates a new DescribeCommitRefs200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCommitRefs200ResponseResponseWithDefaults

`func NewDescribeCommitRefs200ResponseResponseWithDefaults() *DescribeCommitRefs200ResponseResponse`

NewDescribeCommitRefs200ResponseResponseWithDefaults instantiates a new DescribeCommitRefs200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitRefs

`func (o *DescribeCommitRefs200ResponseResponse) GetCommitRefs() []CommitRef`

GetCommitRefs returns the CommitRefs field if non-nil, zero value otherwise.

### GetCommitRefsOk

`func (o *DescribeCommitRefs200ResponseResponse) GetCommitRefsOk() (*[]CommitRef, bool)`

GetCommitRefsOk returns a tuple with the CommitRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitRefs

`func (o *DescribeCommitRefs200ResponseResponse) SetCommitRefs(v []CommitRef)`

SetCommitRefs sets CommitRefs field to given value.

### HasCommitRefs

`func (o *DescribeCommitRefs200ResponseResponse) HasCommitRefs() bool`

HasCommitRefs returns a boolean if a field has been set.

### SetCommitRefsNil

`func (o *DescribeCommitRefs200ResponseResponse) SetCommitRefsNil(b bool)`

 SetCommitRefsNil sets the value for CommitRefs to be an explicit nil

### UnsetCommitRefs
`func (o *DescribeCommitRefs200ResponseResponse) UnsetCommitRefs()`

UnsetCommitRefs ensures that no value is present for CommitRefs, not even an explicit nil
### GetRequestId

`func (o *DescribeCommitRefs200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeCommitRefs200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeCommitRefs200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeCommitRefs200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


