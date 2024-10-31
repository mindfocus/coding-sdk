# ModifyWikiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | 内容 | 
**Iid** | **int64** | wiki编号 | 
**Msg** | Pointer to **string** | 备注 | [optional] 
**ProjectName** | **string** | 项目名称 | 
**Title** | **string** | 标题 | 

## Methods

### NewModifyWikiRequest

`func NewModifyWikiRequest(content string, iid int64, projectName string, title string, ) *ModifyWikiRequest`

NewModifyWikiRequest instantiates a new ModifyWikiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWikiRequestWithDefaults

`func NewModifyWikiRequestWithDefaults() *ModifyWikiRequest`

NewModifyWikiRequestWithDefaults instantiates a new ModifyWikiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ModifyWikiRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModifyWikiRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModifyWikiRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetIid

`func (o *ModifyWikiRequest) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *ModifyWikiRequest) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *ModifyWikiRequest) SetIid(v int64)`

SetIid sets Iid field to given value.


### GetMsg

`func (o *ModifyWikiRequest) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ModifyWikiRequest) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ModifyWikiRequest) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ModifyWikiRequest) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetProjectName

`func (o *ModifyWikiRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyWikiRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyWikiRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTitle

`func (o *ModifyWikiRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModifyWikiRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModifyWikiRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


