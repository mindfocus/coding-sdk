# DescribeArtifactRepositoryFileListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]SpecifiedArtifact**](SpecifiedArtifact.md) | 指定的制品 | [optional] 
**ContinuationToken** | Pointer to **string** | 翻页符。每次 list 操作会返回 ContinuationToken，在下一次 list 传入该值，即可接续上次 list 内容进行 list，最后一页该字段为空 | [optional] 
**PageSize** | Pointer to **int32** | 每页展示条数 | [optional] 
**Project** | **string** | 项目名称 | 
**Repository** | **string** | 仓库名 | 

## Methods

### NewDescribeArtifactRepositoryFileListRequest

`func NewDescribeArtifactRepositoryFileListRequest(project string, repository string, ) *DescribeArtifactRepositoryFileListRequest`

NewDescribeArtifactRepositoryFileListRequest instantiates a new DescribeArtifactRepositoryFileListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeArtifactRepositoryFileListRequestWithDefaults

`func NewDescribeArtifactRepositoryFileListRequestWithDefaults() *DescribeArtifactRepositoryFileListRequest`

NewDescribeArtifactRepositoryFileListRequestWithDefaults instantiates a new DescribeArtifactRepositoryFileListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *DescribeArtifactRepositoryFileListRequest) GetArtifacts() []SpecifiedArtifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *DescribeArtifactRepositoryFileListRequest) GetArtifactsOk() (*[]SpecifiedArtifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *DescribeArtifactRepositoryFileListRequest) SetArtifacts(v []SpecifiedArtifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *DescribeArtifactRepositoryFileListRequest) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetContinuationToken

`func (o *DescribeArtifactRepositoryFileListRequest) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *DescribeArtifactRepositoryFileListRequest) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *DescribeArtifactRepositoryFileListRequest) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *DescribeArtifactRepositoryFileListRequest) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeArtifactRepositoryFileListRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeArtifactRepositoryFileListRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeArtifactRepositoryFileListRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeArtifactRepositoryFileListRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetProject

`func (o *DescribeArtifactRepositoryFileListRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *DescribeArtifactRepositoryFileListRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *DescribeArtifactRepositoryFileListRequest) SetProject(v string)`

SetProject sets Project field to given value.


### GetRepository

`func (o *DescribeArtifactRepositoryFileListRequest) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DescribeArtifactRepositoryFileListRequest) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DescribeArtifactRepositoryFileListRequest) SetRepository(v string)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


