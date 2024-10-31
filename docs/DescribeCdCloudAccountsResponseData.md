# DescribeCdCloudAccountsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAccounts** | Pointer to [**[]CloudAccount**](CloudAccount.md) | 云账号列表 | [optional] 
**PageNumber** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **string** | 每页条数 | [optional] [default to ""]
**TotalPage** | Pointer to **string** | 总共页数 | [optional] [default to ""]
**TotalRow** | Pointer to **string** |  总共条数 | [optional] [default to ""]

## Methods

### NewDescribeCdCloudAccountsResponseData

`func NewDescribeCdCloudAccountsResponseData() *DescribeCdCloudAccountsResponseData`

NewDescribeCdCloudAccountsResponseData instantiates a new DescribeCdCloudAccountsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdCloudAccountsResponseDataWithDefaults

`func NewDescribeCdCloudAccountsResponseDataWithDefaults() *DescribeCdCloudAccountsResponseData`

NewDescribeCdCloudAccountsResponseDataWithDefaults instantiates a new DescribeCdCloudAccountsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAccounts

`func (o *DescribeCdCloudAccountsResponseData) GetCloudAccounts() []CloudAccount`

GetCloudAccounts returns the CloudAccounts field if non-nil, zero value otherwise.

### GetCloudAccountsOk

`func (o *DescribeCdCloudAccountsResponseData) GetCloudAccountsOk() (*[]CloudAccount, bool)`

GetCloudAccountsOk returns a tuple with the CloudAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccounts

`func (o *DescribeCdCloudAccountsResponseData) SetCloudAccounts(v []CloudAccount)`

SetCloudAccounts sets CloudAccounts field to given value.

### HasCloudAccounts

`func (o *DescribeCdCloudAccountsResponseData) HasCloudAccounts() bool`

HasCloudAccounts returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeCdCloudAccountsResponseData) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeCdCloudAccountsResponseData) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeCdCloudAccountsResponseData) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeCdCloudAccountsResponseData) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeCdCloudAccountsResponseData) GetPageSize() string`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeCdCloudAccountsResponseData) GetPageSizeOk() (*string, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeCdCloudAccountsResponseData) SetPageSize(v string)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeCdCloudAccountsResponseData) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalPage

`func (o *DescribeCdCloudAccountsResponseData) GetTotalPage() string`

GetTotalPage returns the TotalPage field if non-nil, zero value otherwise.

### GetTotalPageOk

`func (o *DescribeCdCloudAccountsResponseData) GetTotalPageOk() (*string, bool)`

GetTotalPageOk returns a tuple with the TotalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPage

`func (o *DescribeCdCloudAccountsResponseData) SetTotalPage(v string)`

SetTotalPage sets TotalPage field to given value.

### HasTotalPage

`func (o *DescribeCdCloudAccountsResponseData) HasTotalPage() bool`

HasTotalPage returns a boolean if a field has been set.

### GetTotalRow

`func (o *DescribeCdCloudAccountsResponseData) GetTotalRow() string`

GetTotalRow returns the TotalRow field if non-nil, zero value otherwise.

### GetTotalRowOk

`func (o *DescribeCdCloudAccountsResponseData) GetTotalRowOk() (*string, bool)`

GetTotalRowOk returns a tuple with the TotalRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRow

`func (o *DescribeCdCloudAccountsResponseData) SetTotalRow(v string)`

SetTotalRow sets TotalRow field to given value.

### HasTotalRow

`func (o *DescribeCdCloudAccountsResponseData) HasTotalRow() bool`

HasTotalRow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


