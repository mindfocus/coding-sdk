# DescribeCdDeployTrendByProjectResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]CdDeployTrendDetail**](CdDeployTrendDetail.md) | 各应用发布趋势详情 | [optional] 
**EndDate** | Pointer to **string** | 结束时间 | [optional] [default to ""]
**StartDate** | Pointer to **string** | 开始时间 | [optional] [default to ""]
**Total** | Pointer to [**CdDeployTrendTotal**](CdDeployTrendTotal.md) |  | [optional] 

## Methods

### NewDescribeCdDeployTrendByProjectResponseData

`func NewDescribeCdDeployTrendByProjectResponseData() *DescribeCdDeployTrendByProjectResponseData`

NewDescribeCdDeployTrendByProjectResponseData instantiates a new DescribeCdDeployTrendByProjectResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdDeployTrendByProjectResponseDataWithDefaults

`func NewDescribeCdDeployTrendByProjectResponseDataWithDefaults() *DescribeCdDeployTrendByProjectResponseData`

NewDescribeCdDeployTrendByProjectResponseDataWithDefaults instantiates a new DescribeCdDeployTrendByProjectResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DescribeCdDeployTrendByProjectResponseData) GetDetails() []CdDeployTrendDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DescribeCdDeployTrendByProjectResponseData) GetDetailsOk() (*[]CdDeployTrendDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DescribeCdDeployTrendByProjectResponseData) SetDetails(v []CdDeployTrendDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DescribeCdDeployTrendByProjectResponseData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDate

`func (o *DescribeCdDeployTrendByProjectResponseData) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeCdDeployTrendByProjectResponseData) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeCdDeployTrendByProjectResponseData) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DescribeCdDeployTrendByProjectResponseData) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeCdDeployTrendByProjectResponseData) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeCdDeployTrendByProjectResponseData) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeCdDeployTrendByProjectResponseData) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribeCdDeployTrendByProjectResponseData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotal

`func (o *DescribeCdDeployTrendByProjectResponseData) GetTotal() CdDeployTrendTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DescribeCdDeployTrendByProjectResponseData) GetTotalOk() (*CdDeployTrendTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DescribeCdDeployTrendByProjectResponseData) SetTotal(v CdDeployTrendTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DescribeCdDeployTrendByProjectResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


