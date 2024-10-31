# DescribePersonalExternalDepotsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotType** | **string** | 仓库类型 | 
**ProjectId** | **int32** | 项目 Id | 

## Methods

### NewDescribePersonalExternalDepotsRequest

`func NewDescribePersonalExternalDepotsRequest(depotType string, projectId int32, ) *DescribePersonalExternalDepotsRequest`

NewDescribePersonalExternalDepotsRequest instantiates a new DescribePersonalExternalDepotsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribePersonalExternalDepotsRequestWithDefaults

`func NewDescribePersonalExternalDepotsRequestWithDefaults() *DescribePersonalExternalDepotsRequest`

NewDescribePersonalExternalDepotsRequestWithDefaults instantiates a new DescribePersonalExternalDepotsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotType

`func (o *DescribePersonalExternalDepotsRequest) GetDepotType() string`

GetDepotType returns the DepotType field if non-nil, zero value otherwise.

### GetDepotTypeOk

`func (o *DescribePersonalExternalDepotsRequest) GetDepotTypeOk() (*string, bool)`

GetDepotTypeOk returns a tuple with the DepotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotType

`func (o *DescribePersonalExternalDepotsRequest) SetDepotType(v string)`

SetDepotType sets DepotType field to given value.


### GetProjectId

`func (o *DescribePersonalExternalDepotsRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *DescribePersonalExternalDepotsRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *DescribePersonalExternalDepotsRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


