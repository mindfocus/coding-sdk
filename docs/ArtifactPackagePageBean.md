# ArtifactPackagePageBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSet** | Pointer to [**[]ArtifactPackageBean**](ArtifactPackageBean.md) | 分页数据列表 | [optional] 
**PageNumber** | Pointer to **int32** | 页码 | [optional] 
**PageSize** | Pointer to **int32** | 每页展示数量 | [optional] 
**TotalCount** | Pointer to **int32** | 数据总数 | [optional] 

## Methods

### NewArtifactPackagePageBean

`func NewArtifactPackagePageBean() *ArtifactPackagePageBean`

NewArtifactPackagePageBean instantiates a new ArtifactPackagePageBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactPackagePageBeanWithDefaults

`func NewArtifactPackagePageBeanWithDefaults() *ArtifactPackagePageBean`

NewArtifactPackagePageBeanWithDefaults instantiates a new ArtifactPackagePageBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSet

`func (o *ArtifactPackagePageBean) GetInstanceSet() []ArtifactPackageBean`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *ArtifactPackagePageBean) GetInstanceSetOk() (*[]ArtifactPackageBean, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *ArtifactPackagePageBean) SetInstanceSet(v []ArtifactPackageBean)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *ArtifactPackagePageBean) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.

### GetPageNumber

`func (o *ArtifactPackagePageBean) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *ArtifactPackagePageBean) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *ArtifactPackagePageBean) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *ArtifactPackagePageBean) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *ArtifactPackagePageBean) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ArtifactPackagePageBean) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ArtifactPackagePageBean) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ArtifactPackagePageBean) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ArtifactPackagePageBean) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ArtifactPackagePageBean) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ArtifactPackagePageBean) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ArtifactPackagePageBean) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


