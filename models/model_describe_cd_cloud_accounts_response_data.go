/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCdCloudAccountsResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdCloudAccountsResponseData{}

// DescribeCdCloudAccountsResponseData DescribeCdCloudAccountsResponseData 结构
type DescribeCdCloudAccountsResponseData struct {
	// 云账号列表
	CloudAccounts []CloudAccount `json:"CloudAccounts,omitempty"`
	// 页码
	PageNumber *int64 `json:"PageNumber,omitempty"`
	// 每页条数
	PageSize *string `json:"PageSize,omitempty"`
	// 总共页数
	TotalPage *string `json:"TotalPage,omitempty"`
	//  总共条数
	TotalRow *string `json:"TotalRow,omitempty"`
}

// NewDescribeCdCloudAccountsResponseData instantiates a new DescribeCdCloudAccountsResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdCloudAccountsResponseData() *DescribeCdCloudAccountsResponseData {
	this := DescribeCdCloudAccountsResponseData{}
	var pageSize string = ""
	this.PageSize = &pageSize
	var totalPage string = ""
	this.TotalPage = &totalPage
	var totalRow string = ""
	this.TotalRow = &totalRow
	return &this
}

// NewDescribeCdCloudAccountsResponseDataWithDefaults instantiates a new DescribeCdCloudAccountsResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdCloudAccountsResponseDataWithDefaults() *DescribeCdCloudAccountsResponseData {
	this := DescribeCdCloudAccountsResponseData{}
	var pageSize string = ""
	this.PageSize = &pageSize
	var totalPage string = ""
	this.TotalPage = &totalPage
	var totalRow string = ""
	this.TotalRow = &totalRow
	return &this
}

// GetCloudAccounts returns the CloudAccounts field value if set, zero value otherwise.
func (o *DescribeCdCloudAccountsResponseData) GetCloudAccounts() []CloudAccount {
	if o == nil || utils.IsNil(o.CloudAccounts) {
		var ret []CloudAccount
		return ret
	}
	return o.CloudAccounts
}

// GetCloudAccountsOk returns a tuple with the CloudAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdCloudAccountsResponseData) GetCloudAccountsOk() ([]CloudAccount, bool) {
	if o == nil || utils.IsNil(o.CloudAccounts) {
		return nil, false
	}
	return o.CloudAccounts, true
}

// HasCloudAccounts returns a boolean if a field has been set.
func (o *DescribeCdCloudAccountsResponseData) HasCloudAccounts() bool {
	if o != nil && !utils.IsNil(o.CloudAccounts) {
		return true
	}

	return false
}

// SetCloudAccounts gets a reference to the given []CloudAccount and assigns it to the CloudAccounts field.
func (o *DescribeCdCloudAccountsResponseData) SetCloudAccounts(v []CloudAccount) {
	o.CloudAccounts = v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *DescribeCdCloudAccountsResponseData) GetPageNumber() int64 {
	if o == nil || utils.IsNil(o.PageNumber) {
		var ret int64
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdCloudAccountsResponseData) GetPageNumberOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *DescribeCdCloudAccountsResponseData) HasPageNumber() bool {
	if o != nil && !utils.IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int64 and assigns it to the PageNumber field.
func (o *DescribeCdCloudAccountsResponseData) SetPageNumber(v int64) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeCdCloudAccountsResponseData) GetPageSize() string {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret string
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdCloudAccountsResponseData) GetPageSizeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeCdCloudAccountsResponseData) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given string and assigns it to the PageSize field.
func (o *DescribeCdCloudAccountsResponseData) SetPageSize(v string) {
	o.PageSize = &v
}

// GetTotalPage returns the TotalPage field value if set, zero value otherwise.
func (o *DescribeCdCloudAccountsResponseData) GetTotalPage() string {
	if o == nil || utils.IsNil(o.TotalPage) {
		var ret string
		return ret
	}
	return *o.TotalPage
}

// GetTotalPageOk returns a tuple with the TotalPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdCloudAccountsResponseData) GetTotalPageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TotalPage) {
		return nil, false
	}
	return o.TotalPage, true
}

// HasTotalPage returns a boolean if a field has been set.
func (o *DescribeCdCloudAccountsResponseData) HasTotalPage() bool {
	if o != nil && !utils.IsNil(o.TotalPage) {
		return true
	}

	return false
}

// SetTotalPage gets a reference to the given string and assigns it to the TotalPage field.
func (o *DescribeCdCloudAccountsResponseData) SetTotalPage(v string) {
	o.TotalPage = &v
}

// GetTotalRow returns the TotalRow field value if set, zero value otherwise.
func (o *DescribeCdCloudAccountsResponseData) GetTotalRow() string {
	if o == nil || utils.IsNil(o.TotalRow) {
		var ret string
		return ret
	}
	return *o.TotalRow
}

// GetTotalRowOk returns a tuple with the TotalRow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCdCloudAccountsResponseData) GetTotalRowOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TotalRow) {
		return nil, false
	}
	return o.TotalRow, true
}

// HasTotalRow returns a boolean if a field has been set.
func (o *DescribeCdCloudAccountsResponseData) HasTotalRow() bool {
	if o != nil && !utils.IsNil(o.TotalRow) {
		return true
	}

	return false
}

// SetTotalRow gets a reference to the given string and assigns it to the TotalRow field.
func (o *DescribeCdCloudAccountsResponseData) SetTotalRow(v string) {
	o.TotalRow = &v
}

func (o DescribeCdCloudAccountsResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdCloudAccountsResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CloudAccounts) {
		toSerialize["CloudAccounts"] = o.CloudAccounts
	}
	if !utils.IsNil(o.PageNumber) {
		toSerialize["PageNumber"] = o.PageNumber
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	if !utils.IsNil(o.TotalPage) {
		toSerialize["TotalPage"] = o.TotalPage
	}
	if !utils.IsNil(o.TotalRow) {
		toSerialize["TotalRow"] = o.TotalRow
	}
	return toSerialize, nil
}

type NullableDescribeCdCloudAccountsResponseData struct {
	value *DescribeCdCloudAccountsResponseData
	isSet bool
}

func (v NullableDescribeCdCloudAccountsResponseData) Get() *DescribeCdCloudAccountsResponseData {
	return v.value
}

func (v *NullableDescribeCdCloudAccountsResponseData) Set(val *DescribeCdCloudAccountsResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdCloudAccountsResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdCloudAccountsResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdCloudAccountsResponseData(val *DescribeCdCloudAccountsResponseData) *NullableDescribeCdCloudAccountsResponseData {
	return &NullableDescribeCdCloudAccountsResponseData{value: val, isSet: true}
}

func (v NullableDescribeCdCloudAccountsResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdCloudAccountsResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


