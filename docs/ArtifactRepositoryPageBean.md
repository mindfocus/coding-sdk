# ArtifactRepositoryPageBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]ArtifactRepositoryBean**](ArtifactRepositoryBean.md) | 分页数据列表 | [optional] 
**PageNumber** | Pointer to **int32** | 页码 | [optional] 
**PageSize** | Pointer to **int32** | 每页展示数量 | [optional] 
**TotalCount** | Pointer to **int32** | 数据总数 | [optional] 

## Methods

### NewArtifactRepositoryPageBean

`func NewArtifactRepositoryPageBean() *ArtifactRepositoryPageBean`

NewArtifactRepositoryPageBean instantiates a new ArtifactRepositoryPageBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRepositoryPageBeanWithDefaults

`func NewArtifactRepositoryPageBeanWithDefaults() *ArtifactRepositoryPageBean`

NewArtifactRepositoryPageBeanWithDefaults instantiates a new ArtifactRepositoryPageBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *ArtifactRepositoryPageBean) GetInstanceSet() []ArtifactRepositoryBean`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *ArtifactRepositoryPageBean) GetInstanceSetOk() (*[]ArtifactRepositoryBean, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *ArtifactRepositoryPageBean) SetInstanceSet(v []ArtifactRepositoryBean)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *ArtifactRepositoryPageBean) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetPageNumber

`func (o *ArtifactRepositoryPageBean) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ArtifactRepositoryPageBean) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ArtifactRepositoryPageBean) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ArtifactRepositoryPageBean) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ArtifactRepositoryPageBean) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ArtifactRepositoryPageBean) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ArtifactRepositoryPageBean) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ArtifactRepositoryPageBean) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ArtifactRepositoryPageBean) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ArtifactRepositoryPageBean) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ArtifactRepositoryPageBean) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ArtifactRepositoryPageBean) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


