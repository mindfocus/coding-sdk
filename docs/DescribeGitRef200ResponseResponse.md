# DescribeGitRef200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRef** | Pointer to [**GitRef**](GitRef.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitRef200ResponseResponse

`func NewDescribeGitRef200ResponseResponse() *DescribeGitRef200ResponseResponse`

NewDescribeGitRef200ResponseResponse instantiates a new DescribeGitRef200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitRef200ResponseResponseWithDefaults

`func NewDescribeGitRef200ResponseResponseWithDefaults() *DescribeGitRef200ResponseResponse`

NewDescribeGitRef200ResponseResponseWithDefaults instantiates a new DescribeGitRef200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRef

`func (o *DescribeGitRef200ResponseResponse) GetGitRef() GitRef`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *DescribeGitRef200ResponseResponse) GetGitRefOk() (*GitRef, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *DescribeGitRef200ResponseResponse) SetGitRef(v GitRef)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *DescribeGitRef200ResponseResponse) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitRef200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitRef200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitRef200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitRef200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


