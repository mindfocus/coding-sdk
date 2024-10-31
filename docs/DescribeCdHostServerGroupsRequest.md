# DescribeCdHostServerGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | 检索关键字 | [optional] 
**PageNumber** | Pointer to **int64** | 页码（默认为 1） | [optional] 
**PageSize** | Pointer to **int64** | 每页条数（默认为 10，最大值为 500） | [optional] 

## Methods

### NewDescribeCdHostServerGroupsRequest

`func NewDescribeCdHostServerGroupsRequest() *DescribeCdHostServerGroupsRequest`

NewDescribeCdHostServerGroupsRequest instantiates a new DescribeCdHostServerGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdHostServerGroupsRequestWithDefaults

`func NewDescribeCdHostServerGroupsRequestWithDefaults() *DescribeCdHostServerGroupsRequest`

NewDescribeCdHostServerGroupsRequestWithDefaults instantiates a new DescribeCdHostServerGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *DescribeCdHostServerGroupsRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeCdHostServerGroupsRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeCdHostServerGroupsRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeCdHostServerGroupsRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeCdHostServerGroupsRequest) GetPageNumber() int64`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeCdHostServerGroupsRequest) GetPageNumberOk() (*int64, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeCdHostServerGroupsRequest) SetPageNumber(v int64)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *DescribeCdHostServerGroupsRequest) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeCdHostServerGroupsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeCdHostServerGroupsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeCdHostServerGroupsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeCdHostServerGroupsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

