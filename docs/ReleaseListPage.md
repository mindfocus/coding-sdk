# ReleaseListPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Releases** | Pointer to [**[]Release**](Release.md) | 版本分页列表 | [optional] 
**TotalCount** | Pointer to **int64** | 总共数量 | [optional] 

## Methods

### NewReleaseListPage

`func NewReleaseListPage() *ReleaseListPage`

NewReleaseListPage instantiates a new ReleaseListPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseListPageWithDefaults

`func NewReleaseListPageWithDefaults() *ReleaseListPage`

NewReleaseListPageWithDefaults instantiates a new ReleaseListPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleases

`func (o *ReleaseListPage) GetReleases() []Release`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *ReleaseListPage) GetReleasesOk() (*[]Release, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *ReleaseListPage) SetReleases(v []Release)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *ReleaseListPage) HasReleases() bool`

HasReleases returns a boolean if a field has been set.

### GetTotalCount

`func (o *ReleaseListPage) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ReleaseListPage) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ReleaseListPage) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ReleaseListPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


