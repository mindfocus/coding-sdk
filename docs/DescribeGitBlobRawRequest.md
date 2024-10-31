# DescribeGitBlobRawRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobSha** | **string** | Blob ID | 
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，与仓库ID二选一 | [optional] 

## Methods

### NewDescribeGitBlobRawRequest

`func NewDescribeGitBlobRawRequest(blobSha string, depotId int64, ) *DescribeGitBlobRawRequest`

NewDescribeGitBlobRawRequest instantiates a new DescribeGitBlobRawRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBlobRawRequestWithDefaults

`func NewDescribeGitBlobRawRequestWithDefaults() *DescribeGitBlobRawRequest`

NewDescribeGitBlobRawRequestWithDefaults instantiates a new DescribeGitBlobRawRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobSha

`func (o *DescribeGitBlobRawRequest) GetBlobSha() string`

GetBlobSha returns the BlobSha field if non-nil, zero value otherwise.

### GetBlobShaOk

`func (o *DescribeGitBlobRawRequest) GetBlobShaOk() (*string, bool)`

GetBlobShaOk returns a tuple with the BlobSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSha

`func (o *DescribeGitBlobRawRequest) SetBlobSha(v string)`

SetBlobSha sets BlobSha field to given value.


### GetDepotId

`func (o *DescribeGitBlobRawRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitBlobRawRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitBlobRawRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitBlobRawRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitBlobRawRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitBlobRawRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitBlobRawRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


