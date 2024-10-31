# DescribeDepotByNameInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotName** | **string** | 仓库名 | 
**ProjectName** | **string** | 项目名称 | 
**TeamGk** | **string** | 团队GK | 

## Methods

### NewDescribeDepotByNameInfoRequest

`func NewDescribeDepotByNameInfoRequest(depotName string, projectName string, teamGk string, ) *DescribeDepotByNameInfoRequest`

NewDescribeDepotByNameInfoRequest instantiates a new DescribeDepotByNameInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeDepotByNameInfoRequestWithDefaults

`func NewDescribeDepotByNameInfoRequestWithDefaults() *DescribeDepotByNameInfoRequest`

NewDescribeDepotByNameInfoRequestWithDefaults instantiates a new DescribeDepotByNameInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotName

`func (o *DescribeDepotByNameInfoRequest) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *DescribeDepotByNameInfoRequest) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *DescribeDepotByNameInfoRequest) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.


### GetProjectName

`func (o *DescribeDepotByNameInfoRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeDepotByNameInfoRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeDepotByNameInfoRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTeamGk

`func (o *DescribeDepotByNameInfoRequest) GetTeamGk() string`

GetTeamGk returns the TeamGk field if non-nil, zero value otherwise.

### GetTeamGkOk

`func (o *DescribeDepotByNameInfoRequest) GetTeamGkOk() (*string, bool)`

GetTeamGkOk returns a tuple with the TeamGk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamGk

`func (o *DescribeDepotByNameInfoRequest) SetTeamGk(v string)`

SetTeamGk sets TeamGk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


