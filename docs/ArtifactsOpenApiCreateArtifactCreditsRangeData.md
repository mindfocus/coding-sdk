# ArtifactsOpenApiCreateArtifactCreditsRangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteTeam** | Pointer to [**ArtifactsOpenApiCreditRemoteTeam**](ArtifactsOpenApiCreditRemoteTeam.md) |  | [optional] 
**RangeType** | **string** | 生效范围（TEAM:团队；REMOTE_TEAM:远程团队；PROJECT:项目;REPO:仓库） | [default to ""]
**RangeId** | **int64** | 适用范围ID | [default to 0]

## Methods

### NewArtifactsOpenApiCreateArtifactCreditsRangeData

`func NewArtifactsOpenApiCreateArtifactCreditsRangeData(rangeType string, rangeId int64, ) *ArtifactsOpenApiCreateArtifactCreditsRangeData`

NewArtifactsOpenApiCreateArtifactCreditsRangeData instantiates a new ArtifactsOpenApiCreateArtifactCreditsRangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiCreateArtifactCreditsRangeDataWithDefaults

`func NewArtifactsOpenApiCreateArtifactCreditsRangeDataWithDefaults() *ArtifactsOpenApiCreateArtifactCreditsRangeData`

NewArtifactsOpenApiCreateArtifactCreditsRangeDataWithDefaults instantiates a new ArtifactsOpenApiCreateArtifactCreditsRangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteTeam

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRemoteTeam() ArtifactsOpenApiCreditRemoteTeam`

GetRemoteTeam returns the RemoteTeam field if non-nil, zero value otherwise.

### GetRemoteTeamOk

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRemoteTeamOk() (*ArtifactsOpenApiCreditRemoteTeam, bool)`

GetRemoteTeamOk returns a tuple with the RemoteTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTeam

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) SetRemoteTeam(v ArtifactsOpenApiCreditRemoteTeam)`

SetRemoteTeam sets RemoteTeam field to given value.

### HasRemoteTeam

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) HasRemoteTeam() bool`

HasRemoteTeam returns a boolean if a field has been set.

### GetRangeType

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.


### GetRangeId

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeId() int64`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) GetRangeIdOk() (*int64, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *ArtifactsOpenApiCreateArtifactCreditsRangeData) SetRangeId(v int64)`

SetRangeId sets RangeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


