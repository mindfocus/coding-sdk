# DescribeGitTags200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitTags** | Pointer to [**[]GitTag**](GitTag.md) | tag列表信息 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitTags200ResponseResponse

`func NewDescribeGitTags200ResponseResponse() *DescribeGitTags200ResponseResponse`

NewDescribeGitTags200ResponseResponse instantiates a new DescribeGitTags200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitTags200ResponseResponseWithDefaults

`func NewDescribeGitTags200ResponseResponseWithDefaults() *DescribeGitTags200ResponseResponse`

NewDescribeGitTags200ResponseResponseWithDefaults instantiates a new DescribeGitTags200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitTags

`func (o *DescribeGitTags200ResponseResponse) GetGitTags() []GitTag`

GetGitTags returns the GitTags field if non-nil, zero value otherwise.

### GetGitTagsOk

`func (o *DescribeGitTags200ResponseResponse) GetGitTagsOk() (*[]GitTag, bool)`

GetGitTagsOk returns a tuple with the GitTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTags

`func (o *DescribeGitTags200ResponseResponse) SetGitTags(v []GitTag)`

SetGitTags sets GitTags field to given value.

### HasGitTags

`func (o *DescribeGitTags200ResponseResponse) HasGitTags() bool`

HasGitTags returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitTags200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitTags200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitTags200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitTags200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


