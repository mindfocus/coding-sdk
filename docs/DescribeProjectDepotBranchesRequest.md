# DescribeProjectDepotBranchesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotType** | **string** | 仓库类型 | 
**Id** | **int32** | 仓库 Id | 
**ProjectId** | **int32** | 项目 Id | 

## Methods

### NewDescribeProjectDepotBranchesRequest

`func NewDescribeProjectDepotBranchesRequest(depotType string, id int32, projectId int32, ) *DescribeProjectDepotBranchesRequest`

NewDescribeProjectDepotBranchesRequest instantiates a new DescribeProjectDepotBranchesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProjectDepotBranchesRequestWithDefaults

`func NewDescribeProjectDepotBranchesRequestWithDefaults() *DescribeProjectDepotBranchesRequest`

NewDescribeProjectDepotBranchesRequestWithDefaults instantiates a new DescribeProjectDepotBranchesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotType

`func (o *DescribeProjectDepotBranchesRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *DescribeProjectDepotBranchesRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *DescribeProjectDepotBranchesRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetId

`func (o *DescribeProjectDepotBranchesRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DescribeProjectDepotBranchesRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DescribeProjectDepotBranchesRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetProjectId

`func (o *DescribeProjectDepotBranchesRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribeProjectDepotBranchesRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribeProjectDepotBranchesRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


