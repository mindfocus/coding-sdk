# DescribeCodingProjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | **int32** | 页数 | 
**PageSize** | **int32** | 每页条数 | 
**ProjectName** | Pointer to **string** | 项目名称 | [optional] 

## Methods

### NewDescribeCodingProjectsRequest

`func NewDescribeCodingProjectsRequest(pageNumber int32, pageSize int32, ) *DescribeCodingProjectsRequest`

NewDescribeCodingProjectsRequest instantiates a new DescribeCodingProjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingProjectsRequestWithDefaults

`func NewDescribeCodingProjectsRequestWithDefaults() *DescribeCodingProjectsRequest`

NewDescribeCodingProjectsRequestWithDefaults instantiates a new DescribeCodingProjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *DescribeCodingProjectsRequest) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *DescribeCodingProjectsRequest) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *DescribeCodingProjectsRequest) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.


### GetPageSize

`func (o *DescribeCodingProjectsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeCodingProjectsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeCodingProjectsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetProjectName

`func (o *DescribeCodingProjectsRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DescribeCodingProjectsRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DescribeCodingProjectsRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *DescribeCodingProjectsRequest) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


