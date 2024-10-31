# CreateGitTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Message** | **string** | tag创建信息 | 
**StartPoint** | **string** | 创建来源的分支或commitId | 
**TagName** | **string** | tag名称 | 

## Methods

### NewCreateGitTagRequest

`func NewCreateGitTagRequest(depotId int64, message string, startPoint string, tagName string, ) *CreateGitTagRequest`

NewCreateGitTagRequest instantiates a new CreateGitTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitTagRequestWithDefaults

`func NewCreateGitTagRequestWithDefaults() *CreateGitTagRequest`

NewCreateGitTagRequestWithDefaults instantiates a new CreateGitTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *CreateGitTagRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitTagRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitTagRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitTagRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitTagRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitTagRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitTagRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetMessage

`func (o *CreateGitTagRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateGitTagRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateGitTagRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStartPoint

`func (o *CreateGitTagRequest) GetStartPoint() string`

GetStartPoint returns the StartPoint field if non-nil, zero value otherwise.

### GetStartPointOk

`func (o *CreateGitTagRequest) GetStartPointOk() (*string, bool)`

GetStartPointOk returns a tuple with the StartPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPoint

`func (o *CreateGitTagRequest) SetStartPoint(v string)`

SetStartPoint sets StartPoint field to given value.


### GetTagName

`func (o *CreateGitTagRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *CreateGitTagRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *CreateGitTagRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


