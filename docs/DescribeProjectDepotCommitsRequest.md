# DescribeProjectDepotCommitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | 分支名称 | 
**DepotType** | **string** | 仓库类型 | 
**Id** | **int32** | 仓库 Id | 
**ProjectId** | **int32** | 项目 Id | 

## Methods

### NewDescribeProjectDepotCommitsRequest

`func NewDescribeProjectDepotCommitsRequest(branch string, depotType string, id int32, projectId int32, ) *DescribeProjectDepotCommitsRequest`

NewDescribeProjectDepotCommitsRequest instantiates a new DescribeProjectDepotCommitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectDepotCommitsRequestWithDefaults

`func NewDescribeProjectDepotCommitsRequestWithDefaults() *DescribeProjectDepotCommitsRequest`

NewDescribeProjectDepotCommitsRequestWithDefaults instantiates a new DescribeProjectDepotCommitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *DescribeProjectDepotCommitsRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DescribeProjectDepotCommitsRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DescribeProjectDepotCommitsRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetDepotType

`func (o *DescribeProjectDepotCommitsRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *DescribeProjectDepotCommitsRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *DescribeProjectDepotCommitsRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetId

`func (o *DescribeProjectDepotCommitsRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeProjectDepotCommitsRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeProjectDepotCommitsRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectId

`func (o *DescribeProjectDepotCommitsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeProjectDepotCommitsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeProjectDepotCommitsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


