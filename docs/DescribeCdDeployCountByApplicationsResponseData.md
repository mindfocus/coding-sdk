# DescribeCdDeployCountByApplicationsResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]CdDeployCountDetail**](CdDeployCountDetail.md) | 各应用发布次数详情 | [optional] 
**EndDate** | Pointer to **string** | 结束时间 | [optional] [default to ""]
**StartDate** | Pointer to **string** | 开始时间 | [optional] [default to ""]
**Total** | Pointer to [**CdDeployCount**](CdDeployCount.md) |  | [optional] 

## Methods

### NewDescribeCdDeployCountByApplicationsResponseData

`func NewDescribeCdDeployCountByApplicationsResponseData() *DescribeCdDeployCountByApplicationsResponseData`

NewDescribeCdDeployCountByApplicationsResponseData instantiates a new DescribeCdDeployCountByApplicationsResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdDeployCountByApplicationsResponseDataWithDefaults

`func NewDescribeCdDeployCountByApplicationsResponseDataWithDefaults() *DescribeCdDeployCountByApplicationsResponseData`

NewDescribeCdDeployCountByApplicationsResponseDataWithDefaults instantiates a new DescribeCdDeployCountByApplicationsResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetDetails() []CdDeployCountDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetDetailsOk() (*[]CdDeployCountDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DescribeCdDeployCountByApplicationsResponseData) SetDetails(v []CdDeployCountDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DescribeCdDeployCountByApplicationsResponseData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDate

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeCdDeployCountByApplicationsResponseData) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DescribeCdDeployCountByApplicationsResponseData) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeCdDeployCountByApplicationsResponseData) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribeCdDeployCountByApplicationsResponseData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotal

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetTotal() CdDeployCount`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DescribeCdDeployCountByApplicationsResponseData) GetTotalOk() (*CdDeployCount, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DescribeCdDeployCountByApplicationsResponseData) SetTotal(v CdDeployCount)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DescribeCdDeployCountByApplicationsResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


