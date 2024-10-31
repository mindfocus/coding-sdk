# DescribeGitFiles200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]GitTreeItem**](GitTreeItem.md) | 获取分支目录结构 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitFiles200ResponseResponse

`func NewDescribeGitFiles200ResponseResponse() *DescribeGitFiles200ResponseResponse`

NewDescribeGitFiles200ResponseResponse instantiates a new DescribeGitFiles200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitFiles200ResponseResponseWithDefaults

`func NewDescribeGitFiles200ResponseResponseWithDefaults() *DescribeGitFiles200ResponseResponse`

NewDescribeGitFiles200ResponseResponseWithDefaults instantiates a new DescribeGitFiles200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *DescribeGitFiles200ResponseResponse) GetItems() []GitTreeItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DescribeGitFiles200ResponseResponse) GetItemsOk() (*[]GitTreeItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DescribeGitFiles200ResponseResponse) SetItems(v []GitTreeItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *DescribeGitFiles200ResponseResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitFiles200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitFiles200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitFiles200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitFiles200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


