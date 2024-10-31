# DeleteGitTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**TagName** | **string** | tag名称 | 

## Methods

### NewDeleteGitTagRequest

`func NewDeleteGitTagRequest(depotId int64, tagName string, ) *DeleteGitTagRequest`

NewDeleteGitTagRequest instantiates a new DeleteGitTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteGitTagRequestWithDefaults

`func NewDeleteGitTagRequestWithDefaults() *DeleteGitTagRequest`

NewDeleteGitTagRequestWithDefaults instantiates a new DeleteGitTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DeleteGitTagRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DeleteGitTagRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DeleteGitTagRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DeleteGitTagRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DeleteGitTagRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DeleteGitTagRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DeleteGitTagRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetTagName

`func (o *DeleteGitTagRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *DeleteGitTagRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *DeleteGitTagRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


