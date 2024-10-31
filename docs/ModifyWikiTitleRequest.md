# ModifyWikiTitleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iid** | **int64** | wiki 编号 | 
**ProjectName** | **string** | 项目名称 | 
**Title** | **string** | wiki 标题 | 

## Methods

### NewModifyWikiTitleRequest

`func NewModifyWikiTitleRequest(iid int64, projectName string, title string, ) *ModifyWikiTitleRequest`

NewModifyWikiTitleRequest instantiates a new ModifyWikiTitleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyWikiTitleRequestWithDefaults

`func NewModifyWikiTitleRequestWithDefaults() *ModifyWikiTitleRequest`

NewModifyWikiTitleRequestWithDefaults instantiates a new ModifyWikiTitleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIid

`func (o *ModifyWikiTitleRequest) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *ModifyWikiTitleRequest) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *ModifyWikiTitleRequest) SetIid(v int64)`

SetIid sets Iid field to given value.


### GetProjectName

`func (o *ModifyWikiTitleRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyWikiTitleRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyWikiTitleRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetTitle

`func (o *ModifyWikiTitleRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModifyWikiTitleRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModifyWikiTitleRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


