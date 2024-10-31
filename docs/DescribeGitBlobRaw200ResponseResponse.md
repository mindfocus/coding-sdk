# DescribeGitBlobRaw200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** | blob 文本内容 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitBlobRaw200ResponseResponse

`func NewDescribeGitBlobRaw200ResponseResponse() *DescribeGitBlobRaw200ResponseResponse`

NewDescribeGitBlobRaw200ResponseResponse instantiates a new DescribeGitBlobRaw200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBlobRaw200ResponseResponseWithDefaults

`func NewDescribeGitBlobRaw200ResponseResponseWithDefaults() *DescribeGitBlobRaw200ResponseResponse`

NewDescribeGitBlobRaw200ResponseResponseWithDefaults instantiates a new DescribeGitBlobRaw200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *DescribeGitBlobRaw200ResponseResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *DescribeGitBlobRaw200ResponseResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *DescribeGitBlobRaw200ResponseResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *DescribeGitBlobRaw200ResponseResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *DescribeGitBlobRaw200ResponseResponse) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *DescribeGitBlobRaw200ResponseResponse) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetRequestId

`func (o *DescribeGitBlobRaw200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitBlobRaw200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitBlobRaw200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitBlobRaw200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


