# DescribeTeamArtifactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int64** | 页码，默认：1 | [optional] 
**PageSize** | Pointer to **int64** | 每页展示数量，默认：10 | [optional] 
**Rule** | Pointer to [**ArtifactFilterRule**](ArtifactFilterRule.md) |  | [optional] 

## Methods

### NewDescribeTeamArtifactsRequest

`func NewDescribeTeamArtifactsRequest() *DescribeTeamArtifactsRequest`

NewDescribeTeamArtifactsRequest instantiates a new DescribeTeamArtifactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamArtifactsRequestWithDefaults

`func NewDescribeTeamArtifactsRequestWithDefaults() *DescribeTeamArtifactsRequest`

NewDescribeTeamArtifactsRequestWithDefaults instantiates a new DescribeTeamArtifactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeTeamArtifactsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeTeamArtifactsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeTeamArtifactsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeTeamArtifactsRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeTeamArtifactsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeTeamArtifactsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeTeamArtifactsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeTeamArtifactsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRule

`func (o *DescribeTeamArtifactsRequest) GetRule() ArtifactFilterRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *DescribeTeamArtifactsRequest) GetRuleOk() (*ArtifactFilterRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *DescribeTeamArtifactsRequest) SetRule(v ArtifactFilterRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *DescribeTeamArtifactsRequest) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


