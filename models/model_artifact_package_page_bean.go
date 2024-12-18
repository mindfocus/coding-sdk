/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ArtifactPackagePageBean type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ArtifactPackagePageBean{}

// ArtifactPackagePageBean 制品包列表分页返回体
type ArtifactPackagePageBean struct {
	// 分页数据列表
	InstanceSet []ArtifactPackageBean `json:"InstanceSet,omitempty"`
	// 页码
	PageNumber *int32 `json:"PageNumber,omitempty"`
	// 每页展示数量
	PageSize *int32 `json:"PageSize,omitempty"`
	// 数据总数
	TotalCount *int32 `json:"TotalCount,omitempty"`
}

// NewArtifactPackagePageBean instantiates a new ArtifactPackagePageBean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactPackagePageBean() *ArtifactPackagePageBean {
	this := ArtifactPackagePageBean{}
	return &this
}

// NewArtifactPackagePageBeanWithDefaults instantiates a new ArtifactPackagePageBean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactPackagePageBeanWithDefaults() *ArtifactPackagePageBean {
	this := ArtifactPackagePageBean{}
	return &this
}

// GetInstanceSet returns the InstanceSet field value if set, zero value otherwise.
func (o *ArtifactPackagePageBean) GetInstanceSet() []ArtifactPackageBean {
	if o == nil || utils.IsNil(o.InstanceSet) {
		var ret []ArtifactPackageBean
		return ret
	}
	return o.InstanceSet
}

// GetInstanceSetOk returns a tuple with the InstanceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactPackagePageBean) GetInstanceSetOk() ([]ArtifactPackageBean, bool) {
	if o == nil || utils.IsNil(o.InstanceSet) {
		return nil, false
	}
	return o.InstanceSet, true
}

// HasInstanceSet returns a boolean if a field has been set.
func (o *ArtifactPackagePageBean) HasInstanceSet() bool {
	if o != nil && !utils.IsNil(o.InstanceSet) {
		return true
	}

	return false
}

// SetInstanceSet gets a reference to the given []ArtifactPackageBean and assigns it to the InstanceSet field.
func (o *ArtifactPackagePageBean) SetInstanceSet(v []ArtifactPackageBean) {
	o.InstanceSet = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *ArtifactPackagePageBean) GetPageNumber() int32 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactPackagePageBean) GetPageNumberOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *ArtifactPackagePageBean) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *ArtifactPackagePageBean) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ArtifactPackagePageBean) GetPageSize() int32 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactPackagePageBean) GetPageSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ArtifactPackagePageBean) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ArtifactPackagePageBean) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ArtifactPackagePageBean) GetTotalCount() int32 {
	if o == nil || utils.IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactPackagePageBean) GetTotalCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ArtifactPackagePageBean) HasTotalCount() bool {
	if o != nil && !utils.IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ArtifactPackagePageBean) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ArtifactPackagePageBean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactPackagePageBean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.InstanceSet) {
		toSerialize["InstanceSet"] = o.InstanceSet
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.TotalCount) {
		toSerialize["TotalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableArtifactPackagePageBean struct {
	value *ArtifactPackagePageBean
	isSet bool
}

func (v NullableArtifactPackagePageBean) Get() *ArtifactPackagePageBean {
	return v.value
}

func (v *NullableArtifactPackagePageBean) Set(val *ArtifactPackagePageBean) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactPackagePageBean) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactPackagePageBean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactPackagePageBean(val *ArtifactPackagePageBean) *NullableArtifactPackagePageBean {
	return &NullableArtifactPackagePageBean{value: val, isSet: true}
}

func (v NullableArtifactPackagePageBean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactPackagePageBean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


