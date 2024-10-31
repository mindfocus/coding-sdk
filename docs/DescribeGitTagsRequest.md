# DescribeGitTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径 | [optional] 
**KeyWord** | Pointer to **string** | TagName的模糊查询关键字 | [optional] 
**Page** | Pointer to **int64** | 页码 | [optional] 
**PageSize** | Pointer to **int64** | 每页数量 | [optional] 

## Methods

### NewDescribeGitTagsRequest

`func NewDescribeGitTagsRequest(depotId int64, ) *DescribeGitTagsRequest`

NewDescribeGitTagsRequest instantiates a new DescribeGitTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitTagsRequestWithDefaults

`func NewDescribeGitTagsRequestWithDefaults() *DescribeGitTagsRequest`

NewDescribeGitTagsRequestWithDefaults instantiates a new DescribeGitTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepotId

`func (o *DescribeGitTagsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeGitTagsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeGitTagsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeGitTagsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeGitTagsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeGitTagsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeGitTagsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetKeyWord

`func (o *DescribeGitTagsRequest) GetKeyWord() string`

GetKeyWord returns the KeyWord field if non-nil, zero value otherwise.

### GetKeyWordOk

`func (o *DescribeGitTagsRequest) GetKeyWordOk() (*string, bool)`

GetKeyWordOk returns a tuple with the KeyWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWord

`func (o *DescribeGitTagsRequest) SetKeyWord(v string)`

SetKeyWord sets KeyWord field to given value.

### HasKeyWord

`func (o *DescribeGitTagsRequest) HasKeyWord() bool`

HasKeyWord returns a boolean if a field has been set.

### GetPage

`func (o *DescribeGitTagsRequest) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DescribeGitTagsRequest) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DescribeGitTagsRequest) SetPage(v int64)`

SetPage sets Page field to given value.

### HasPage

`func (o *DescribeGitTagsRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *DescribeGitTagsRequest) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DescribeGitTagsRequest) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DescribeGitTagsRequest) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DescribeGitTagsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


