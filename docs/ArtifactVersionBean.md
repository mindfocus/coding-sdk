# ArtifactVersionBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | 创建时间 | [optional] 
**Description** | Pointer to **string** | 版本描述 | [optional] [default to ""]
**DownloadCount** | Pointer to **int32** | 下载量 | [optional] 
**Hash** | Pointer to **string** | 版本哈希 | [optional] [default to ""]
**Id** | Pointer to **int32** | 版本 ID | [optional] 
**PkgId** | Pointer to **int32** | 制品包 ID | [optional] 
**ReleaseStatus** | Pointer to **int64** | 发布状态：1-未发布;2-已发布 | [optional] 
**Size** | Pointer to **NullableFloat32** | 版本大小（单位：MB） | [optional] 
**Version** | Pointer to **string** | 版本号 | [optional] [default to ""]

## Methods

### NewArtifactVersionBean

`func NewArtifactVersionBean() *ArtifactVersionBean`

NewArtifactVersionBean instantiates a new ArtifactVersionBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactVersionBeanWithDefaults

`func NewArtifactVersionBeanWithDefaults() *ArtifactVersionBean`

NewArtifactVersionBeanWithDefaults instantiates a new ArtifactVersionBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ArtifactVersionBean) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArtifactVersionBean) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArtifactVersionBean) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ArtifactVersionBean) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *ArtifactVersionBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArtifactVersionBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArtifactVersionBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArtifactVersionBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownloadCount

`func (o *ArtifactVersionBean) GetDownloadCount() int32`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ArtifactVersionBean) GetDownloadCountOk() (*int32, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ArtifactVersionBean) SetDownloadCount(v int32)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *ArtifactVersionBean) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetHash

`func (o *ArtifactVersionBean) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ArtifactVersionBean) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ArtifactVersionBean) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ArtifactVersionBean) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ArtifactVersionBean) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArtifactVersionBean) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArtifactVersionBean) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ArtifactVersionBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPkgId

`func (o *ArtifactVersionBean) GetPkgId() int32`

GetPkgId returns the PkgId field if non-nil, zero value otherwise.

### GetPkgIdOk

`func (o *ArtifactVersionBean) GetPkgIdOk() (*int32, bool)`

GetPkgIdOk returns a tuple with the PkgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkgId

`func (o *ArtifactVersionBean) SetPkgId(v int32)`

SetPkgId sets PkgId field to given value.

### HasPkgId

`func (o *ArtifactVersionBean) HasPkgId() bool`

HasPkgId returns a boolean if a field has been set.

### GetReleaseStatus

`func (o *ArtifactVersionBean) GetReleaseStatus() int64`

GetReleaseStatus returns the ReleaseStatus field if non-nil, zero value otherwise.

### GetReleaseStatusOk

`func (o *ArtifactVersionBean) GetReleaseStatusOk() (*int64, bool)`

GetReleaseStatusOk returns a tuple with the ReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatus

`func (o *ArtifactVersionBean) SetReleaseStatus(v int64)`

SetReleaseStatus sets ReleaseStatus field to given value.

### HasReleaseStatus

`func (o *ArtifactVersionBean) HasReleaseStatus() bool`

HasReleaseStatus returns a boolean if a field has been set.

### GetSize

`func (o *ArtifactVersionBean) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactVersionBean) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactVersionBean) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ArtifactVersionBean) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ArtifactVersionBean) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ArtifactVersionBean) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetVersion

`func (o *ArtifactVersionBean) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ArtifactVersionBean) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ArtifactVersionBean) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ArtifactVersionBean) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


