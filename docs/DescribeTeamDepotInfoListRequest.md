# DescribeTeamDepotInfoListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotName** | Pointer to **string** | 仓库名 | [optional] 
**PageNumber** | **int64** | 页码 | 
**PageSize** | **int64** | 页码大小 | 
**ProjectName** | Pointer to **string** | 项目名 | [optional] 

## Methods

### NewDescribeTeamDepotInfoListRequest

`func NewDescribeTeamDepotInfoListRequest(pageNumber int64, pageSize int64, ) *DescribeTeamDepotInfoListRequest`

NewDescribeTeamDepotInfoListRequest instantiates a new DescribeTeamDepotInfoListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamDepotInfoListRequestWithDefaults

`func NewDescribeTeamDepotInfoListRequestWithDefaults() *DescribeTeamDepotInfoListRequest`

NewDescribeTeamDepotInfoListRequestWithDefaults instantiates a new DescribeTeamDepotInfoListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotName

`func (o *DescribeTeamDepotInfoListRequest) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *DescribeTeamDepotInfoListRequest) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *DescribeTeamDepotInfoListRequest) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.

### HasDepotName

`func (o *DescribeTeamDepotInfoListRequest) HasDepotName() bool`

HasDepotName returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeTeamDepotInfoListRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeTeamDepotInfoListRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeTeamDepotInfoListRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeTeamDepotInfoListRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeTeamDepotInfoListRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeTeamDepotInfoListRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetProjectName

`func (o *DescribeTeamDepotInfoListRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeTeamDepotInfoListRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeTeamDepotInfoListRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *DescribeTeamDepotInfoListRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


