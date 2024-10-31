# CreateWikiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Wiki内容 | 
**Msg** | Pointer to **string** | 备注 | [optional] 
**ParentIid** | **int64** | 父级iid | 
**ProjectName** | **string** | 项目名称 | 
**Title** | **string** | Wiki标题 | 

## Methods

### NewCreateWikiRequest

`func NewCreateWikiRequest(content string, parentIid int64, projectName string, title string, ) *CreateWikiRequest`

NewCreateWikiRequest instantiates a new CreateWikiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWikiRequestWithDefaults

`func NewCreateWikiRequestWithDefaults() *CreateWikiRequest`

NewCreateWikiRequestWithDefaults instantiates a new CreateWikiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *CreateWikiRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreateWikiRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreateWikiRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetMsg

`func (o *CreateWikiRequest) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CreateWikiRequest) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CreateWikiRequest) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CreateWikiRequest) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetParentIid

`func (o *CreateWikiRequest) GetParentIid() int64`

GetParentIid returns the ParentIid field if non-nil, zero value otherwise.

### GetParentIidOk

`func (o *CreateWikiRequest) GetParentIidOk() (*int64, bool)`

GetParentIidOk returns a tuple with the ParentIid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIid

`func (o *CreateWikiRequest) SetParentIid(v int64)`

SetParentIid sets ParentIid field to given value.


### GetProjectName

`func (o *CreateWikiRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateWikiRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateWikiRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTitle

`func (o *CreateWikiRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateWikiRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateWikiRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


