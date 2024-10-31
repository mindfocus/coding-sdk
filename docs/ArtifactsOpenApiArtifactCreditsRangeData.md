# ArtifactsOpenApiArtifactCreditsRangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | [**ArtifactsOpenApiProjectData**](ArtifactsOpenApiProjectData.md) |  | 
**RemoteTeam** | [**ArtifactsOpenApiRemoteTeamData**](ArtifactsOpenApiRemoteTeamData.md) |  | 
**RangeType** | **string** | 生效范围（TEAM:团队；REMOTE-TEAM:远程团队；PROJECT:项目;REPO:仓库） | [default to ""]
**Repository** | [**ArtifactsOpenApiRepositoryData**](ArtifactsOpenApiRepositoryData.md) |  | 
**RangeId** | **int64** | 适用范围ID | [default to 0]

## Methods

### NewArtifactsOpenApiArtifactCreditsRangeData

`func NewArtifactsOpenApiArtifactCreditsRangeData(project ArtifactsOpenApiProjectData, remoteTeam ArtifactsOpenApiRemoteTeamData, rangeType string, repository ArtifactsOpenApiRepositoryData, rangeId int64, ) *ArtifactsOpenApiArtifactCreditsRangeData`

NewArtifactsOpenApiArtifactCreditsRangeData instantiates a new ArtifactsOpenApiArtifactCreditsRangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactsOpenApiArtifactCreditsRangeDataWithDefaults

`func NewArtifactsOpenApiArtifactCreditsRangeDataWithDefaults() *ArtifactsOpenApiArtifactCreditsRangeData`

NewArtifactsOpenApiArtifactCreditsRangeDataWithDefaults instantiates a new ArtifactsOpenApiArtifactCreditsRangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetProject() ArtifactsOpenApiProjectData`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetProjectOk() (*ArtifactsOpenApiProjectData, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) SetProject(v ArtifactsOpenApiProjectData)`

SetProject sets Project field to given value.


### GetRemoteTeam

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRemoteTeam() ArtifactsOpenApiRemoteTeamData`

GetRemoteTeam returns the RemoteTeam field if non-nil, zero value otherwise.

### GetRemoteTeamOk

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRemoteTeamOk() (*ArtifactsOpenApiRemoteTeamData, bool)`

GetRemoteTeamOk returns a tuple with the RemoteTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTeam

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) SetRemoteTeam(v ArtifactsOpenApiRemoteTeamData)`

SetRemoteTeam sets RemoteTeam field to given value.


### GetRangeType

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.


### GetRepository

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRepository() ArtifactsOpenApiRepositoryData`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRepositoryOk() (*ArtifactsOpenApiRepositoryData, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) SetRepository(v ArtifactsOpenApiRepositoryData)`

SetRepository sets Repository field to given value.


### GetRangeId

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRangeId() int64`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) GetRangeIdOk() (*int64, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *ArtifactsOpenApiArtifactCreditsRangeData) SetRangeId(v int64)`

SetRangeId sets RangeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


