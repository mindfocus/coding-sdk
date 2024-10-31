# WikiData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanMaintain** | Pointer to **bool** | 是否为维护者 | [optional] [default to false]
**CanRead** | Pointer to **NullableBool** | 是否可以阅读 | [optional] [default to false]
**Content** | Pointer to **string** | 内容 | [optional] [default to ""]
**CreatedAt** | Pointer to **map[string]interface{}** | 创建时间 | [optional] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] 
**CurrentVersion** | Pointer to **int64** | 当前versionId | [optional] 
**Editor** | Pointer to [**User**](User.md) |  | [optional] 
**HistoriesCount** | Pointer to **int64** | 修改次数 | [optional] 
**HistoryId** | Pointer to **int64** | wiki历史Id | [optional] 
**Html** | Pointer to **string** | 内容转成的html | [optional] [default to ""]
**Id** | Pointer to **int64** | wikiId | [optional] 
**Iid** | Pointer to **int64** | wik的code | [optional] 
**LastVersion** | Pointer to **int64** | 最新versionId | [optional] 
**Msg** | Pointer to **NullableString** | 提交说明 | [optional] [default to ""]
**Order** | Pointer to **int64** | 所处顺序位置 | [optional] 
**ParentIid** | Pointer to **int64** | 父级 IiD | [optional] 
**ParentShared** | Pointer to **bool** | 是否父级分享 | [optional] [default to false]
**ParentVisibleRange** | Pointer to **NullableString** | 父级可见范围 | [optional] [default to ""]
**Path** | Pointer to **NullableString** | 路径 | [optional] [default to ""]
**Title** | Pointer to **string** | 标题 | [optional] [default to ""]
**UpdatedAt** | Pointer to **map[string]interface{}** | 修改时间 | [optional] 
**VisibleRange** | Pointer to **NullableString** | 可见范围 | [optional] [default to ""]

## Methods

### NewWikiData

`func NewWikiData() *WikiData`

NewWikiData instantiates a new WikiData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiDataWithDefaults

`func NewWikiDataWithDefaults() *WikiData`

NewWikiDataWithDefaults instantiates a new WikiData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanMaintain

`func (o *WikiData) GetCanMaintain() bool`

GetCanMaintain returns the CanMaintain field if non-nil, zero value otherwise.

### GetCanMaintainOk

`func (o *WikiData) GetCanMaintainOk() (*bool, bool)`

GetCanMaintainOk returns a tuple with the CanMaintain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanMaintain

`func (o *WikiData) SetCanMaintain(v bool)`

SetCanMaintain sets CanMaintain field to given value.

### HasCanMaintain

`func (o *WikiData) HasCanMaintain() bool`

HasCanMaintain returns a boolean if a field has been set.

### GetCanRead

`func (o *WikiData) GetCanRead() bool`

GetCanRead returns the CanRead field if non-nil, zero value otherwise.

### GetCanReadOk

`func (o *WikiData) GetCanReadOk() (*bool, bool)`

GetCanReadOk returns a tuple with the CanRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanRead

`func (o *WikiData) SetCanRead(v bool)`

SetCanRead sets CanRead field to given value.

### HasCanRead

`func (o *WikiData) HasCanRead() bool`

HasCanRead returns a boolean if a field has been set.

### SetCanReadNil

`func (o *WikiData) SetCanReadNil(b bool)`

 SetCanReadNil sets the value for CanRead to be an explicit nil

### UnsetCanRead
`func (o *WikiData) UnsetCanRead()`

UnsetCanRead ensures that no value is present for CanRead, not even an explicit nil
### GetContent

`func (o *WikiData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WikiData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WikiData) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WikiData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WikiData) GetCreatedAt() map[string]interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WikiData) GetCreatedAtOk() (*map[string]interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WikiData) SetCreatedAt(v map[string]interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WikiData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *WikiData) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WikiData) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WikiData) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WikiData) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *WikiData) GetCurrentVersion() int64`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *WikiData) GetCurrentVersionOk() (*int64, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *WikiData) SetCurrentVersion(v int64)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *WikiData) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetEditor

`func (o *WikiData) GetEditor() User`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *WikiData) GetEditorOk() (*User, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *WikiData) SetEditor(v User)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *WikiData) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetHistoriesCount

`func (o *WikiData) GetHistoriesCount() int64`

GetHistoriesCount returns the HistoriesCount field if non-nil, zero value otherwise.

### GetHistoriesCountOk

`func (o *WikiData) GetHistoriesCountOk() (*int64, bool)`

GetHistoriesCountOk returns a tuple with the HistoriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoriesCount

`func (o *WikiData) SetHistoriesCount(v int64)`

SetHistoriesCount sets HistoriesCount field to given value.

### HasHistoriesCount

`func (o *WikiData) HasHistoriesCount() bool`

HasHistoriesCount returns a boolean if a field has been set.

### GetHistoryId

`func (o *WikiData) GetHistoryId() int64`

GetHistoryId returns the HistoryId field if non-nil, zero value otherwise.

### GetHistoryIdOk

`func (o *WikiData) GetHistoryIdOk() (*int64, bool)`

GetHistoryIdOk returns a tuple with the HistoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryId

`func (o *WikiData) SetHistoryId(v int64)`

SetHistoryId sets HistoryId field to given value.

### HasHistoryId

`func (o *WikiData) HasHistoryId() bool`

HasHistoryId returns a boolean if a field has been set.

### GetHtml

`func (o *WikiData) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *WikiData) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *WikiData) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *WikiData) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetId

`func (o *WikiData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WikiData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIid

`func (o *WikiData) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *WikiData) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *WikiData) SetIid(v int64)`

SetIid sets Iid field to given value.

### HasIid

`func (o *WikiData) HasIid() bool`

HasIid returns a boolean if a field has been set.

### GetLastVersion

`func (o *WikiData) GetLastVersion() int64`

GetLastVersion returns the LastVersion field if non-nil, zero value otherwise.

### GetLastVersionOk

`func (o *WikiData) GetLastVersionOk() (*int64, bool)`

GetLastVersionOk returns a tuple with the LastVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVersion

`func (o *WikiData) SetLastVersion(v int64)`

SetLastVersion sets LastVersion field to given value.

### HasLastVersion

`func (o *WikiData) HasLastVersion() bool`

HasLastVersion returns a boolean if a field has been set.

### GetMsg

`func (o *WikiData) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *WikiData) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *WikiData) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *WikiData) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *WikiData) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *WikiData) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetOrder

`func (o *WikiData) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WikiData) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WikiData) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WikiData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentIid

`func (o *WikiData) GetParentIid() int64`

GetParentIid returns the ParentIid field if non-nil, zero value otherwise.

### GetParentIidOk

`func (o *WikiData) GetParentIidOk() (*int64, bool)`

GetParentIidOk returns a tuple with the ParentIid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIid

`func (o *WikiData) SetParentIid(v int64)`

SetParentIid sets ParentIid field to given value.

### HasParentIid

`func (o *WikiData) HasParentIid() bool`

HasParentIid returns a boolean if a field has been set.

### GetParentShared

`func (o *WikiData) GetParentShared() bool`

GetParentShared returns the ParentShared field if non-nil, zero value otherwise.

### GetParentSharedOk

`func (o *WikiData) GetParentSharedOk() (*bool, bool)`

GetParentSharedOk returns a tuple with the ParentShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentShared

`func (o *WikiData) SetParentShared(v bool)`

SetParentShared sets ParentShared field to given value.

### HasParentShared

`func (o *WikiData) HasParentShared() bool`

HasParentShared returns a boolean if a field has been set.

### GetParentVisibleRange

`func (o *WikiData) GetParentVisibleRange() string`

GetParentVisibleRange returns the ParentVisibleRange field if non-nil, zero value otherwise.

### GetParentVisibleRangeOk

`func (o *WikiData) GetParentVisibleRangeOk() (*string, bool)`

GetParentVisibleRangeOk returns a tuple with the ParentVisibleRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVisibleRange

`func (o *WikiData) SetParentVisibleRange(v string)`

SetParentVisibleRange sets ParentVisibleRange field to given value.

### HasParentVisibleRange

`func (o *WikiData) HasParentVisibleRange() bool`

HasParentVisibleRange returns a boolean if a field has been set.

### SetParentVisibleRangeNil

`func (o *WikiData) SetParentVisibleRangeNil(b bool)`

 SetParentVisibleRangeNil sets the value for ParentVisibleRange to be an explicit nil

### UnsetParentVisibleRange
`func (o *WikiData) UnsetParentVisibleRange()`

UnsetParentVisibleRange ensures that no value is present for ParentVisibleRange, not even an explicit nil
### GetPath

`func (o *WikiData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WikiData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WikiData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WikiData) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *WikiData) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *WikiData) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetTitle

`func (o *WikiData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WikiData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WikiData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WikiData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WikiData) GetUpdatedAt() map[string]interface{}`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WikiData) GetUpdatedAtOk() (*map[string]interface{}, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WikiData) SetUpdatedAt(v map[string]interface{})`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WikiData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibleRange

`func (o *WikiData) GetVisibleRange() string`

GetVisibleRange returns the VisibleRange field if non-nil, zero value otherwise.

### GetVisibleRangeOk

`func (o *WikiData) GetVisibleRangeOk() (*string, bool)`

GetVisibleRangeOk returns a tuple with the VisibleRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleRange

`func (o *WikiData) SetVisibleRange(v string)`

SetVisibleRange sets VisibleRange field to given value.

### HasVisibleRange

`func (o *WikiData) HasVisibleRange() bool`

HasVisibleRange returns a boolean if a field has been set.

### SetVisibleRangeNil

`func (o *WikiData) SetVisibleRangeNil(b bool)`

 SetVisibleRangeNil sets the value for VisibleRange to be an explicit nil

### UnsetVisibleRange
`func (o *WikiData) UnsetVisibleRange()`

UnsetVisibleRange ensures that no value is present for VisibleRange, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


