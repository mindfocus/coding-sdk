# DescribeCdDeployCountByProjectResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]CdDeployCountDetail**](CdDeployCountDetail.md) | 各应用发布次数详情 | [optional] 
**EndDate** | Pointer to **string** | 结束时间 | [optional] [default to ""]
**StartDate** | Pointer to **string** | 开始时间 | [optional] [default to ""]
**Total** | Pointer to [**CdDeployCount**](CdDeployCount.md) |  | [optional] 

## Methods

### NewDescribeCdDeployCountByProjectResponseData

`func NewDescribeCdDeployCountByProjectResponseData() *DescribeCdDeployCountByProjectResponseData`

NewDescribeCdDeployCountByProjectResponseData instantiates a new DescribeCdDeployCountByProjectResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdDeployCountByProjectResponseDataWithDefaults

`func NewDescribeCdDeployCountByProjectResponseDataWithDefaults() *DescribeCdDeployCountByProjectResponseData`

NewDescribeCdDeployCountByProjectResponseDataWithDefaults instantiates a new DescribeCdDeployCountByProjectResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DescribeCdDeployCountByProjectResponseData) GetDetails() []CdDeployCountDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DescribeCdDeployCountByProjectResponseData) GetDetailsOk() (*[]CdDeployCountDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DescribeCdDeployCountByProjectResponseData) SetDetails(v []CdDeployCountDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DescribeCdDeployCountByProjectResponseData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDate

`func (o *DescribeCdDeployCountByProjectResponseData) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeCdDeployCountByProjectResponseData) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeCdDeployCountByProjectResponseData) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DescribeCdDeployCountByProjectResponseData) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeCdDeployCountByProjectResponseData) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeCdDeployCountByProjectResponseData) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeCdDeployCountByProjectResponseData) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribeCdDeployCountByProjectResponseData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotal

`func (o *DescribeCdDeployCountByProjectResponseData) GetTotal() CdDeployCount`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DescribeCdDeployCountByProjectResponseData) GetTotalOk() (*CdDeployCount, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DescribeCdDeployCountByProjectResponseData) SetTotal(v CdDeployCount)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DescribeCdDeployCountByProjectResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


