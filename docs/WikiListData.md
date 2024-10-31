# WikiListData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]WikiChildrenData**](WikiChildrenData.md) | 子页面信息 | [optional] 
**Content** | Pointer to **string** | 内容 | [optional] [default to ""]
**CreatedAt** | Pointer to **int32** | 创建时间 | [optional] 
**Creator** | Pointer to [**User**](User.md) |  | [optional] 
**Editor** | Pointer to [**User**](User.md) |  | [optional] 
**Id** | Pointer to **int64** | wiki Id | [optional] 
**Iid** | Pointer to **int64** | wiki编号 | [optional] 
**IsShared** | Pointer to **bool** | 是否分享 | [optional] [default to false]
**IsTreeShared** | Pointer to **bool** | 是否整树分享 | [optional] [default to false]
**Order** | Pointer to **float32** | 所处顺序位置 | [optional] 
**ParentIid** | Pointer to **int64** | 父级编号 | [optional] 
**Title** | Pointer to **string** | 标题 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int32** | 更新时间 | [optional] 
**VisibleRange** | Pointer to **string** | 可见范围 | [optional] [default to ""]

## Methods

### NewWikiListData

`func NewWikiListData() *WikiListData`

NewWikiListData instantiates a new WikiListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiListDataWithDefaults

`func NewWikiListDataWithDefaults() *WikiListData`

NewWikiListDataWithDefaults instantiates a new WikiListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *WikiListData) GetChildren() []WikiChildrenData`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *WikiListData) GetChildrenOk() (*[]WikiChildrenData, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *WikiListData) SetChildren(v []WikiChildrenData)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *WikiListData) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetContent

`func (o *WikiListData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WikiListData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WikiListData) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WikiListData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WikiListData) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WikiListData) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WikiListData) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WikiListData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *WikiListData) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WikiListData) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WikiListData) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WikiListData) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEditor

`func (o *WikiListData) GetEditor() User`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *WikiListData) GetEditorOk() (*User, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *WikiListData) SetEditor(v User)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *WikiListData) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetId

`func (o *WikiListData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiListData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiListData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WikiListData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIid

`func (o *WikiListData) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *WikiListData) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *WikiListData) SetIid(v int64)`

SetIid sets Iid field to given value.

### HasIid

`func (o *WikiListData) HasIid() bool`

HasIid returns a boolean if a field has been set.

### GetIsShared

`func (o *WikiListData) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *WikiListData) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *WikiListData) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *WikiListData) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetIsTreeShared

`func (o *WikiListData) GetIsTreeShared() bool`

GetIsTreeShared returns the IsTreeShared field if non-nil, zero value otherwise.

### GetIsTreeSharedOk

`func (o *WikiListData) GetIsTreeSharedOk() (*bool, bool)`

GetIsTreeSharedOk returns a tuple with the IsTreeShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTreeShared

`func (o *WikiListData) SetIsTreeShared(v bool)`

SetIsTreeShared sets IsTreeShared field to given value.

### HasIsTreeShared

`func (o *WikiListData) HasIsTreeShared() bool`

HasIsTreeShared returns a boolean if a field has been set.

### GetOrder

`func (o *WikiListData) GetOrder() float32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WikiListData) GetOrderOk() (*float32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WikiListData) SetOrder(v float32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WikiListData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentIid

`func (o *WikiListData) GetParentIid() int64`

GetParentIid returns the ParentIid field if non-nil, zero value otherwise.

### GetParentIidOk

`func (o *WikiListData) GetParentIidOk() (*int64, bool)`

GetParentIidOk returns a tuple with the ParentIid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIid

`func (o *WikiListData) SetParentIid(v int64)`

SetParentIid sets ParentIid field to given value.

### HasParentIid

`func (o *WikiListData) HasParentIid() bool`

HasParentIid returns a boolean if a field has been set.

### GetTitle

`func (o *WikiListData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WikiListData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WikiListData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WikiListData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WikiListData) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WikiListData) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WikiListData) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WikiListData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibleRange

`func (o *WikiListData) GetVisibleRange() string`

GetVisibleRange returns the VisibleRange field if non-nil, zero value otherwise.

### GetVisibleRangeOk

`func (o *WikiListData) GetVisibleRangeOk() (*string, bool)`

GetVisibleRangeOk returns a tuple with the VisibleRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleRange

`func (o *WikiListData) SetVisibleRange(v string)`

SetVisibleRange sets VisibleRange field to given value.

### HasVisibleRange

`func (o *WikiListData) HasVisibleRange() bool`

HasVisibleRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


