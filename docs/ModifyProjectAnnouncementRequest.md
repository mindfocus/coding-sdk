# ModifyProjectAnnouncementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectName** | **string** | 项目名 | [default to ""]
**Content** | **string** | 公告内容 | [default to ""]
**Id** | **int64** | 公告id | [default to 0]

## Methods

### NewModifyProjectAnnouncementRequest

`func NewModifyProjectAnnouncementRequest(projectName string, content string, id int64, ) *ModifyProjectAnnouncementRequest`

NewModifyProjectAnnouncementRequest instantiates a new ModifyProjectAnnouncementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyProjectAnnouncementRequestWithDefaults

`func NewModifyProjectAnnouncementRequestWithDefaults() *ModifyProjectAnnouncementRequest`

NewModifyProjectAnnouncementRequestWithDefaults instantiates a new ModifyProjectAnnouncementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectName

`func (o *ModifyProjectAnnouncementRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ModifyProjectAnnouncementRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ModifyProjectAnnouncementRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.


### GetContent

`func (o *ModifyProjectAnnouncementRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ModifyProjectAnnouncementRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ModifyProjectAnnouncementRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetId

`func (o *ModifyProjectAnnouncementRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModifyProjectAnnouncementRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModifyProjectAnnouncementRequest) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


