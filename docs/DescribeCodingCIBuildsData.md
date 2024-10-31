# DescribeCodingCIBuildsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildList** | Pointer to [**[]CodingCIBuild**](CodingCIBuild.md) | 构建列表 | [optional] 
**PageNumber** | Pointer to **int32** | 页码 | [optional] 
**PageSize** | Pointer to **int32** | 每页多少条 | [optional] 
**TotalCount** | Pointer to **int32** | 总条数 | [optional] 

## Methods

### NewDescribeCodingCIBuildsData

`func NewDescribeCodingCIBuildsData() *DescribeCodingCIBuildsData`

NewDescribeCodingCIBuildsData instantiates a new DescribeCodingCIBuildsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildsDataWithDefaults

`func NewDescribeCodingCIBuildsDataWithDefaults() *DescribeCodingCIBuildsData`

NewDescribeCodingCIBuildsDataWithDefaults instantiates a new DescribeCodingCIBuildsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildList

`func (o *DescribeCodingCIBuildsData) GetBuildList() []CodingCIBuild`

GetBuildList returns the BuildList field if non-nil, zero value otherwise.

### GetBuildListOk

`func (o *DescribeCodingCIBuildsData) GetBuildListOk() (*[]CodingCIBuild, bool)`

GetBuildListOk returns a tuple with the BuildList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildList

`func (o *DescribeCodingCIBuildsData) SetBuildList(v []CodingCIBuild)`

SetBuildList sets BuildList field to given value.

### HasBuildList

`func (o *DescribeCodingCIBuildsData) HasBuildList() bool`

HasBuildList returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeCodingCIBuildsData) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeCodingCIBuildsData) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeCodingCIBuildsData) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeCodingCIBuildsData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeCodingCIBuildsData) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeCodingCIBuildsData) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeCodingCIBuildsData) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeCodingCIBuildsData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *DescribeCodingCIBuildsData) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DescribeCodingCIBuildsData) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DescribeCodingCIBuildsData) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DescribeCodingCIBuildsData) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


