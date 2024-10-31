# DescribeGitBlobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobSha** | **string** | Blob ID | 
**DepotId** | **int64** | 仓库 ID | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 

## Methods

### NewDescribeGitBlobRequest

`func NewDescribeGitBlobRequest(blobSha string, depotId int64, ) *DescribeGitBlobRequest`

NewDescribeGitBlobRequest instantiates a new DescribeGitBlobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitBlobRequestWithDefaults

`func NewDescribeGitBlobRequestWithDefaults() *DescribeGitBlobRequest`

NewDescribeGitBlobRequestWithDefaults instantiates a new DescribeGitBlobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobSha

`func (o *DescribeGitBlobRequest) GetBlobSha() string`

GetBlobSha returns the BlobSha field if non-nil, zero value otherwise.

### GetBlobShaOk

`func (o *DescribeGitBlobRequest) GetBlobShaOk() (*string, bool)`

GetBlobShaOk returns a tuple with the BlobSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSha

`func (o *DescribeGitBlobRequest) SetBlobSha(v string)`

SetBlobSha sets BlobSha field to given value.


### GetDepotId

`func (o *DescribeGitBlobRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitBlobRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitBlobRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitBlobRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitBlobRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitBlobRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitBlobRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


