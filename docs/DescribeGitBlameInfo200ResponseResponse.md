# DescribeGitBlameInfo200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**[]CommitInfo**](CommitInfo.md) | 提交信息 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitBlameInfo200ResponseResponse

`func NewDescribeGitBlameInfo200ResponseResponse() *DescribeGitBlameInfo200ResponseResponse`

NewDescribeGitBlameInfo200ResponseResponse instantiates a new DescribeGitBlameInfo200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBlameInfo200ResponseResponseWithDefaults

`func NewDescribeGitBlameInfo200ResponseResponseWithDefaults() *DescribeGitBlameInfo200ResponseResponse`

NewDescribeGitBlameInfo200ResponseResponseWithDefaults instantiates a new DescribeGitBlameInfo200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *DescribeGitBlameInfo200ResponseResponse) GetInfo() []CommitInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DescribeGitBlameInfo200ResponseResponse) GetInfoOk() (*[]CommitInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DescribeGitBlameInfo200ResponseResponse) SetInfo(v []CommitInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DescribeGitBlameInfo200ResponseResponse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### SetInfoNil

`func (o *DescribeGitBlameInfo200ResponseResponse) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *DescribeGitBlameInfo200ResponseResponse) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetRequestId

`func (o *DescribeGitBlameInfo200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitBlameInfo200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitBlameInfo200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitBlameInfo200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


