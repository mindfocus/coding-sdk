# DescribeGitTagsBySha200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitTags** | Pointer to [**[]GitTag**](GitTag.md) | tag集合 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitTagsBySha200ResponseResponse

`func NewDescribeGitTagsBySha200ResponseResponse() *DescribeGitTagsBySha200ResponseResponse`

NewDescribeGitTagsBySha200ResponseResponse instantiates a new DescribeGitTagsBySha200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitTagsBySha200ResponseResponseWithDefaults

`func NewDescribeGitTagsBySha200ResponseResponseWithDefaults() *DescribeGitTagsBySha200ResponseResponse`

NewDescribeGitTagsBySha200ResponseResponseWithDefaults instantiates a new DescribeGitTagsBySha200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitTags

`func (o *DescribeGitTagsBySha200ResponseResponse) GetGitTags() []GitTag`

GetGitTags returns the GitTags field if non-nil, zero value otherwise.

### GetGitTagsOk

`func (o *DescribeGitTagsBySha200ResponseResponse) GetGitTagsOk() (*[]GitTag, bool)`

GetGitTagsOk returns a tuple with the GitTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTags

`func (o *DescribeGitTagsBySha200ResponseResponse) SetGitTags(v []GitTag)`

SetGitTags sets GitTags field to given value.

### HasGitTags

`func (o *DescribeGitTagsBySha200ResponseResponse) HasGitTags() bool`

HasGitTags returns a boolean if a field has been set.

### SetGitTagsNil

`func (o *DescribeGitTagsBySha200ResponseResponse) SetGitTagsNil(b bool)`

 SetGitTagsNil sets the value for GitTags to be an explicit nil

### UnsetGitTags
`func (o *DescribeGitTagsBySha200ResponseResponse) UnsetGitTags()`

UnsetGitTags ensures that no value is present for GitTags, not even an explicit nil
### GetRequestId

`func (o *DescribeGitTagsBySha200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitTagsBySha200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitTagsBySha200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitTagsBySha200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


