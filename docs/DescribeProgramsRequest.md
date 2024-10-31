# DescribeProgramsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keyword** | Pointer to **string** | 关键字搜索：项目名 | [optional] 
**PageNumber** | **int32** | 请求页数 | 
**PageSize** | **int32** | 请求条数 | 

## Methods

### NewDescribeProgramsRequest

`func NewDescribeProgramsRequest(pageNumber int32, pageSize int32, ) *DescribeProgramsRequest`

NewDescribeProgramsRequest instantiates a new DescribeProgramsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeProgramsRequestWithDefaults

`func NewDescribeProgramsRequestWithDefaults() *DescribeProgramsRequest`

NewDescribeProgramsRequestWithDefaults instantiates a new DescribeProgramsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyword

`func (o *DescribeProgramsRequest) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *DescribeProgramsRequest) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *DescribeProgramsRequest) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *DescribeProgramsRequest) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### GetPageNumber

`func (o *DescribeProgramsRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeProgramsRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeProgramsRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeProgramsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeProgramsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeProgramsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


