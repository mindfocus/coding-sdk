# DescribeGitFile200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitFile** | Pointer to [**GitFileItem**](GitFileItem.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitFile200ResponseResponse

`func NewDescribeGitFile200ResponseResponse() *DescribeGitFile200ResponseResponse`

NewDescribeGitFile200ResponseResponse instantiates a new DescribeGitFile200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitFile200ResponseResponseWithDefaults

`func NewDescribeGitFile200ResponseResponseWithDefaults() *DescribeGitFile200ResponseResponse`

NewDescribeGitFile200ResponseResponseWithDefaults instantiates a new DescribeGitFile200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitFile

`func (o *DescribeGitFile200ResponseResponse) GetGitFile() GitFileItem`

GetGitFile returns the GitFile field if non-nil, zero value otherwise.

### GetGitFileOk

`func (o *DescribeGitFile200ResponseResponse) GetGitFileOk() (*GitFileItem, bool)`

GetGitFileOk returns a tuple with the GitFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitFile

`func (o *DescribeGitFile200ResponseResponse) SetGitFile(v GitFileItem)`

SetGitFile sets GitFile field to given value.

### HasGitFile

`func (o *DescribeGitFile200ResponseResponse) HasGitFile() bool`

HasGitFile returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitFile200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitFile200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitFile200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitFile200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


