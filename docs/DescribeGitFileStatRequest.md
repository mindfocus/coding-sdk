# DescribeGitFileStatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotPath** | **string** | 仓库路径，格式：&lt;team&gt;/&lt;project&gt;/&lt;depot&gt; | 
**Path** | **string** | 文件路径 | 
**Ref** | **string** | tag | branch | commit | 

## Methods

### NewDescribeGitFileStatRequest

`func NewDescribeGitFileStatRequest(depotPath string, path string, ref string, ) *DescribeGitFileStatRequest`

NewDescribeGitFileStatRequest instantiates a new DescribeGitFileStatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitFileStatRequestWithDefaults

`func NewDescribeGitFileStatRequestWithDefaults() *DescribeGitFileStatRequest`

NewDescribeGitFileStatRequestWithDefaults instantiates a new DescribeGitFileStatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotPath

`func (o *DescribeGitFileStatRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitFileStatRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitFileStatRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.


### GetPath

`func (o *DescribeGitFileStatRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DescribeGitFileStatRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DescribeGitFileStatRequest) SetPath(v string)`

SetPath sets Path field to given value.


### GetRef

`func (o *DescribeGitFileStatRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DescribeGitFileStatRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DescribeGitFileStatRequest) SetRef(v string)`

SetRef sets Ref field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


