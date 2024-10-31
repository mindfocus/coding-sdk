# WikiChildrenData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWikiChildrenData

`func NewWikiChildrenData() *WikiChildrenData`

NewWikiChildrenData instantiates a new WikiChildrenData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiChildrenDataWithDefaults

`func NewWikiChildrenDataWithDefaults() *WikiChildrenData`

NewWikiChildrenDataWithDefaults instantiates a new WikiChildrenData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *WikiChildrenData) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WikiChildrenData) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WikiChildrenData) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WikiChildrenData) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WikiChildrenData) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WikiChildrenData) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WikiChildrenData) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WikiChildrenData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreator

`func (o *WikiChildrenData) GetCreator() User`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *WikiChildrenData) GetCreatorOk() (*User, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *WikiChildrenData) SetCreator(v User)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *WikiChildrenData) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEditor

`func (o *WikiChildrenData) GetEditor() User`

GetEditor returns the Editor field if non-nil, zero value otherwise.

### GetEditorOk

`func (o *WikiChildrenData) GetEditorOk() (*User, bool)`

GetEditorOk returns a tuple with the Editor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditor

`func (o *WikiChildrenData) SetEditor(v User)`

SetEditor sets Editor field to given value.

### HasEditor

`func (o *WikiChildrenData) HasEditor() bool`

HasEditor returns a boolean if a field has been set.

### GetId

`func (o *WikiChildrenData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiChildrenData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiChildrenData) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *WikiChildrenData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIid

`func (o *WikiChildrenData) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *WikiChildrenData) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *WikiChildrenData) SetIid(v int64)`

SetIid sets Iid field to given value.

### HasIid

`func (o *WikiChildrenData) HasIid() bool`

HasIid returns a boolean if a field has been set.

### GetIsShared

`func (o *WikiChildrenData) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *WikiChildrenData) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *WikiChildrenData) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *WikiChildrenData) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetIsTreeShared

`func (o *WikiChildrenData) GetIsTreeShared() bool`

GetIsTreeShared returns the IsTreeShared field if non-nil, zero value otherwise.

### GetIsTreeSharedOk

`func (o *WikiChildrenData) GetIsTreeSharedOk() (*bool, bool)`

GetIsTreeSharedOk returns a tuple with the IsTreeShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTreeShared

`func (o *WikiChildrenData) SetIsTreeShared(v bool)`

SetIsTreeShared sets IsTreeShared field to given value.

### HasIsTreeShared

`func (o *WikiChildrenData) HasIsTreeShared() bool`

HasIsTreeShared returns a boolean if a field has been set.

### GetOrder

`func (o *WikiChildrenData) GetOrder() float32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WikiChildrenData) GetOrderOk() (*float32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WikiChildrenData) SetOrder(v float32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *WikiChildrenData) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentIid

`func (o *WikiChildrenData) GetParentIid() int64`

GetParentIid returns the ParentIid field if non-nil, zero value otherwise.

### GetParentIidOk

`func (o *WikiChildrenData) GetParentIidOk() (*int64, bool)`

GetParentIidOk returns a tuple with the ParentIid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentIid

`func (o *WikiChildrenData) SetParentIid(v int64)`

SetParentIid sets ParentIid field to given value.

### HasParentIid

`func (o *WikiChildrenData) HasParentIid() bool`

HasParentIid returns a boolean if a field has been set.

### GetTitle

`func (o *WikiChildrenData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WikiChildrenData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WikiChildrenData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WikiChildrenData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WikiChildrenData) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WikiChildrenData) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WikiChildrenData) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WikiChildrenData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVisibleRange

`func (o *WikiChildrenData) GetVisibleRange() string`

GetVisibleRange returns the VisibleRange field if non-nil, zero value otherwise.

### GetVisibleRangeOk

`func (o *WikiChildrenData) GetVisibleRangeOk() (*string, bool)`

GetVisibleRangeOk returns a tuple with the VisibleRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleRange

`func (o *WikiChildrenData) SetVisibleRange(v string)`

SetVisibleRange sets VisibleRange field to given value.

### HasVisibleRange

`func (o *WikiChildrenData) HasVisibleRange() bool`

HasVisibleRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


