# ArtifactRepositoryFileListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuationToken** | Pointer to **string** | 翻页符，最后一页该字段为空 | [optional] [default to ""]
**InstanceSet** | Pointer to [**[]ArtifactRepositoryFile**](ArtifactRepositoryFile.md) | 文件列表 | [optional] 

## Methods

### NewArtifactRepositoryFileListData

`func NewArtifactRepositoryFileListData() *ArtifactRepositoryFileListData`

NewArtifactRepositoryFileListData instantiates a new ArtifactRepositoryFileListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactRepositoryFileListDataWithDefaults

`func NewArtifactRepositoryFileListDataWithDefaults() *ArtifactRepositoryFileListData`

NewArtifactRepositoryFileListDataWithDefaults instantiates a new ArtifactRepositoryFileListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuationToken

`func (o *ArtifactRepositoryFileListData) GetContinuationToken() string`

GetContinuationToken returns the ContinuationToken field if non-nil, zero value otherwise.

### GetContinuationTokenOk

`func (o *ArtifactRepositoryFileListData) GetContinuationTokenOk() (*string, bool)`

GetContinuationTokenOk returns a tuple with the ContinuationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuationToken

`func (o *ArtifactRepositoryFileListData) SetContinuationToken(v string)`

SetContinuationToken sets ContinuationToken field to given value.

### HasContinuationToken

`func (o *ArtifactRepositoryFileListData) HasContinuationToken() bool`

HasContinuationToken returns a boolean if a field has been set.

### GetInstanceSet

`func (o *ArtifactRepositoryFileListData) GetInstanceSet() []ArtifactRepositoryFile`

GetInstanceSet returns the InstanceSet field if non-nil, zero value otherwise.

### GetInstanceSetOk

`func (o *ArtifactRepositoryFileListData) GetInstanceSetOk() (*[]ArtifactRepositoryFile, bool)`

GetInstanceSetOk returns a tuple with the InstanceSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSet

`func (o *ArtifactRepositoryFileListData) SetInstanceSet(v []ArtifactRepositoryFile)`

SetInstanceSet sets InstanceSet field to given value.

### HasInstanceSet

`func (o *ArtifactRepositoryFileListData) HasInstanceSet() bool`

HasInstanceSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


