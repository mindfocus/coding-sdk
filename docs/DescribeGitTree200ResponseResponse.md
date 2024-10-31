# DescribeGitTree200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trees** | Pointer to [**[]GitTree**](GitTree.md) | 1 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitTree200ResponseResponse

`func NewDescribeGitTree200ResponseResponse() *DescribeGitTree200ResponseResponse`

NewDescribeGitTree200ResponseResponse instantiates a new DescribeGitTree200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitTree200ResponseResponseWithDefaults

`func NewDescribeGitTree200ResponseResponseWithDefaults() *DescribeGitTree200ResponseResponse`

NewDescribeGitTree200ResponseResponseWithDefaults instantiates a new DescribeGitTree200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrees

`func (o *DescribeGitTree200ResponseResponse) GetTrees() []GitTree`

GetTrees returns the Trees field if non-nil, zero value otherwise.

### GetTreesOk

`func (o *DescribeGitTree200ResponseResponse) GetTreesOk() (*[]GitTree, bool)`

GetTreesOk returns a tuple with the Trees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrees

`func (o *DescribeGitTree200ResponseResponse) SetTrees(v []GitTree)`

SetTrees sets Trees field to given value.

### HasTrees

`func (o *DescribeGitTree200ResponseResponse) HasTrees() bool`

HasTrees returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitTree200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitTree200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitTree200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitTree200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


