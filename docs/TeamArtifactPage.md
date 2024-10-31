# TeamArtifactPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]TeamArtifact**](TeamArtifact.md) | 当前页的版本列表 | [optional] 
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 每页展示数量 | [optional] 
**TotalCount** | Pointer to **int64** | 总数 | [optional] 

## Methods

### NewTeamArtifactPage

`func NewTeamArtifactPage() *TeamArtifactPage`

NewTeamArtifactPage instantiates a new TeamArtifactPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamArtifactPageWithDefaults

`func NewTeamArtifactPageWithDefaults() *TeamArtifactPage`

NewTeamArtifactPageWithDefaults instantiates a new TeamArtifactPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *TeamArtifactPage) GetInstanceSet() []TeamArtifact`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *TeamArtifactPage) GetInstanceSetOk() (*[]TeamArtifact, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *TeamArtifactPage) SetInstanceSet(v []TeamArtifact)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *TeamArtifactPage) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetPageNumber

`func (o *TeamArtifactPage) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *TeamArtifactPage) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *TeamArtifactPage) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *TeamArtifactPage) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *TeamArtifactPage) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *TeamArtifactPage) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *TeamArtifactPage) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *TeamArtifactPage) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *TeamArtifactPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *TeamArtifactPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *TeamArtifactPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *TeamArtifactPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


