# DescribeGitRefsBySha200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refs** | Pointer to [**[]GitReference**](GitReference.md) | ref集合 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitRefsBySha200ResponseResponse

`func NewDescribeGitRefsBySha200ResponseResponse() *DescribeGitRefsBySha200ResponseResponse`

NewDescribeGitRefsBySha200ResponseResponse instantiates a new DescribeGitRefsBySha200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitRefsBySha200ResponseResponseWithDefaults

`func NewDescribeGitRefsBySha200ResponseResponseWithDefaults() *DescribeGitRefsBySha200ResponseResponse`

NewDescribeGitRefsBySha200ResponseResponseWithDefaults instantiates a new DescribeGitRefsBySha200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefs

`func (o *DescribeGitRefsBySha200ResponseResponse) GetRefs() []GitReference`

GetRefs returns the Refs field if non-nil, zero value otherwise.

### GetRefsOk

`func (o *DescribeGitRefsBySha200ResponseResponse) GetRefsOk() (*[]GitReference, bool)`

GetRefsOk returns a tuple with the Refs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefs

`func (o *DescribeGitRefsBySha200ResponseResponse) SetRefs(v []GitReference)`

SetRefs sets Refs field to given value.

### HasRefs

`func (o *DescribeGitRefsBySha200ResponseResponse) HasRefs() bool`

HasRefs returns a boolean if a field has been set.

### SetRefsNil

`func (o *DescribeGitRefsBySha200ResponseResponse) SetRefsNil(b bool)`

 SetRefsNil sets the value for Refs to be an explicit nil

### UnsetRefs
`func (o *DescribeGitRefsBySha200ResponseResponse) UnsetRefs()`

UnsetRefs ensures that no value is present for Refs, not even an explicit nil
### GetRequestId

`func (o *DescribeGitRefsBySha200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitRefsBySha200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitRefsBySha200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitRefsBySha200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


