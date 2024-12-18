/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the WikiListData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WikiListData{}

// WikiListData wiki列表信息
type WikiListData struct {
	// 子页面信息
	Children []WikiChildrenData `json:"Children,omitempty"`
	// 内容
	Content *string `json:"Content,omitempty"`
	// 创建时间
	CreatedAt *int32 `json:"CreatedAt,omitempty"`
	Creator *User `json:"Creator,omitempty"`
	Editor *User `json:"Editor,omitempty"`
	// wiki Id
	Id *int64 `json:"Id,omitempty"`
	// wiki编号
	Iid *int64 `json:"Iid,omitempty"`
	// 是否分享
	IsShared *bool `json:"IsShared,omitempty"`
	// 是否整树分享
	IsTreeShared *bool `json:"IsTreeShared,omitempty"`
	// 所处顺序位置
	Order *float32 `json:"Order,omitempty"`
	// 父级编号
	ParentIid *int64 `json:"ParentIid,omitempty"`
	// 标题
	Title *string `json:"Title,omitempty"`
	// 更新时间
	UpdatedAt *int32 `json:"UpdatedAt,omitempty"`
	// 可见范围
	VisibleRange *string `json:"VisibleRange,omitempty"`
}

// NewWikiListData instantiates a new WikiListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWikiListData() *WikiListData {
	this := WikiListData{}
	var content string = ""
	this.Content = &content
	var isShared bool = false
	this.IsShared = &isShared
	var isTreeShared bool = false
	this.IsTreeShared = &isTreeShared
	var title string = ""
	this.Title = &title
	var visibleRange string = ""
	this.VisibleRange = &visibleRange
	return &this
}

// NewWikiListDataWithDefaults instantiates a new WikiListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWikiListDataWithDefaults() *WikiListData {
	this := WikiListData{}
	var content string = ""
	this.Content = &content
	var isShared bool = false
	this.IsShared = &isShared
	var isTreeShared bool = false
	this.IsTreeShared = &isTreeShared
	var title string = ""
	this.Title = &title
	var visibleRange string = ""
	this.VisibleRange = &visibleRange
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *WikiListData) GetChildren() []WikiChildrenData {
	if o == nil || utils.IsNil(o.Children) {
		var ret []WikiChildrenData
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetChildrenOk() ([]WikiChildrenData, bool) {
	if o == nil || utils.IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *WikiListData) HasChildren() bool {
	if o != nil && !utils.IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []WikiChildrenData and assigns it to the Children field.
func (o *WikiListData) SetChildren(v []WikiChildrenData) {
	o.Children = v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *WikiListData) GetContent() string {
	if o == nil || utils.IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *WikiListData) HasContent() bool {
	if o != nil && !utils.IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *WikiListData) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WikiListData) GetCreatedAt() int32 {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetCreatedAtOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WikiListData) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *WikiListData) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *WikiListData) GetCreator() User {
	if o == nil || utils.IsNil(o.Creator) {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetCreatorOk() (*User, bool) {
	if o == nil || utils.IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *WikiListData) HasCreator() bool {
	if o != nil && !utils.IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *WikiListData) SetCreator(v User) {
	o.Creator = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *WikiListData) GetEditor() User {
	if o == nil || utils.IsNil(o.Editor) {
		var ret User
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetEditorOk() (*User, bool) {
	if o == nil || utils.IsNil(o.Editor) {
		return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *WikiListData) HasEditor() bool {
	if o != nil && !utils.IsNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given User and assigns it to the Editor field.
func (o *WikiListData) SetEditor(v User) {
	o.Editor = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WikiListData) GetId() int64 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WikiListData) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *WikiListData) SetId(v int64) {
	o.Id = &v
}

// GetIid returns the Iid field value if set, zero value otherwise.
func (o *WikiListData) GetIid() int64 {
	if o == nil || utils.IsNil(o.Iid) {
		var ret int64
		return ret
	}
	return *o.Iid
}

// GetIidOk returns a tuple with the Iid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetIidOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Iid) {
		return nil, false
	}
	return o.Iid, true
}

// HasIid returns a boolean if a field has been set.
func (o *WikiListData) HasIid() bool {
	if o != nil && !utils.IsNil(o.Iid) {
		return true
	}

	return false
}

// SetIid gets a reference to the given int64 and assigns it to the Iid field.
func (o *WikiListData) SetIid(v int64) {
	o.Iid = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *WikiListData) GetIsShared() bool {
	if o == nil || utils.IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetIsSharedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *WikiListData) HasIsShared() bool {
	if o != nil && !utils.IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *WikiListData) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetIsTreeShared returns the IsTreeShared field value if set, zero value otherwise.
func (o *WikiListData) GetIsTreeShared() bool {
	if o == nil || utils.IsNil(o.IsTreeShared) {
		var ret bool
		return ret
	}
	return *o.IsTreeShared
}

// GetIsTreeSharedOk returns a tuple with the IsTreeShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetIsTreeSharedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsTreeShared) {
		return nil, false
	}
	return o.IsTreeShared, true
}

// HasIsTreeShared returns a boolean if a field has been set.
func (o *WikiListData) HasIsTreeShared() bool {
	if o != nil && !utils.IsNil(o.IsTreeShared) {
		return true
	}

	return false
}

// SetIsTreeShared gets a reference to the given bool and assigns it to the IsTreeShared field.
func (o *WikiListData) SetIsTreeShared(v bool) {
	o.IsTreeShared = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *WikiListData) GetOrder() float32 {
	if o == nil || utils.IsNil(o.Order) {
		var ret float32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetOrderOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *WikiListData) HasOrder() bool {
	if o != nil && !utils.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given float32 and assigns it to the Order field.
func (o *WikiListData) SetOrder(v float32) {
	o.Order = &v
}

// GetParentIid returns the ParentIid field value if set, zero value otherwise.
func (o *WikiListData) GetParentIid() int64 {
	if o == nil || utils.IsNil(o.ParentIid) {
		var ret int64
		return ret
	}
	return *o.ParentIid
}

// GetParentIidOk returns a tuple with the ParentIid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetParentIidOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ParentIid) {
		return nil, false
	}
	return o.ParentIid, true
}

// HasParentIid returns a boolean if a field has been set.
func (o *WikiListData) HasParentIid() bool {
	if o != nil && !utils.IsNil(o.ParentIid) {
		return true
	}

	return false
}

// SetParentIid gets a reference to the given int64 and assigns it to the ParentIid field.
func (o *WikiListData) SetParentIid(v int64) {
	o.ParentIid = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WikiListData) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WikiListData) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WikiListData) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WikiListData) GetUpdatedAt() int32 {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WikiListData) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *WikiListData) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetVisibleRange returns the VisibleRange field value if set, zero value otherwise.
func (o *WikiListData) GetVisibleRange() string {
	if o == nil || utils.IsNil(o.VisibleRange) {
		var ret string
		return ret
	}
	return *o.VisibleRange
}

// GetVisibleRangeOk returns a tuple with the VisibleRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WikiListData) GetVisibleRangeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VisibleRange) {
		return nil, false
	}
	return o.VisibleRange, true
}

// HasVisibleRange returns a boolean if a field has been set.
func (o *WikiListData) HasVisibleRange() bool {
	if o != nil && !utils.IsNil(o.VisibleRange) {
		return true
	}

	return false
}

// SetVisibleRange gets a reference to the given string and assigns it to the VisibleRange field.
func (o *WikiListData) SetVisibleRange(v string) {
	o.VisibleRange = &v
}

func (o WikiListData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WikiListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Children) {
		toSerialize["Children"] = o.Children
	}
	if !utils.IsNil(o.Content) {
		toSerialize["Content"] = o.Content
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.Creator) {
		toSerialize["Creator"] = o.Creator
	}
	if !utils.IsNil(o.Editor) {
		toSerialize["Editor"] = o.Editor
	}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.Iid) {
		toSerialize["Iid"] = o.Iid
	}
	if !utils.IsNil(o.IsShared) {
		toSerialize["IsShared"] = o.IsShared
	}
	if !utils.IsNil(o.IsTreeShared) {
		toSerialize["IsTreeShared"] = o.IsTreeShared
	}
	if !utils.IsNil(o.Order) {
		toSerialize["Order"] = o.Order
	}
	if !utils.IsNil(o.ParentIid) {
		toSerialize["ParentIid"] = o.ParentIid
	}
	if !utils.IsNil(o.Title) {
		toSerialize["Title"] = o.Title
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["UpdatedAt"] = o.UpdatedAt
	}
	if !utils.IsNil(o.VisibleRange) {
		toSerialize["VisibleRange"] = o.VisibleRange
	}
	return toSerialize, nil
}

type NullableWikiListData struct {
	value *WikiListData
	isSet bool
}

func (v NullableWikiListData) Get() *WikiListData {
	return v.value
}

func (v *NullableWikiListData) Set(val *WikiListData) {
	v.value = val
	v.isSet = true
}

func (v NullableWikiListData) IsSet() bool {
	return v.isSet
}

func (v *NullableWikiListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWikiListData(val *WikiListData) *NullableWikiListData {
	return &NullableWikiListData{value: val, isSet: true}
}

func (v NullableWikiListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWikiListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


