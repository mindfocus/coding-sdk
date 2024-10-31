# DescribeCdDeployTimeByProjectResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]CdDeployTimeDetail**](CdDeployTimeDetail.md) | 各应用发布时长详情 | [optional] 
**EndDate** | Pointer to **string** | 结束时间 | [optional] [default to ""]
**StartDate** | Pointer to **string** | 开始时间 | [optional] [default to ""]
**Total** | Pointer to [**CdDeployTime**](CdDeployTime.md) |  | [optional] 

## Methods

### NewDescribeCdDeployTimeByProjectResponseData

`func NewDescribeCdDeployTimeByProjectResponseData() *DescribeCdDeployTimeByProjectResponseData`

NewDescribeCdDeployTimeByProjectResponseData instantiates a new DescribeCdDeployTimeByProjectResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCdDeployTimeByProjectResponseDataWithDefaults

`func NewDescribeCdDeployTimeByProjectResponseDataWithDefaults() *DescribeCdDeployTimeByProjectResponseData`

NewDescribeCdDeployTimeByProjectResponseDataWithDefaults instantiates a new DescribeCdDeployTimeByProjectResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DescribeCdDeployTimeByProjectResponseData) GetDetails() []CdDeployTimeDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DescribeCdDeployTimeByProjectResponseData) GetDetailsOk() (*[]CdDeployTimeDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DescribeCdDeployTimeByProjectResponseData) SetDetails(v []CdDeployTimeDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DescribeCdDeployTimeByProjectResponseData) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetEndDate

`func (o *DescribeCdDeployTimeByProjectResponseData) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *DescribeCdDeployTimeByProjectResponseData) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *DescribeCdDeployTimeByProjectResponseData) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *DescribeCdDeployTimeByProjectResponseData) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetStartDate

`func (o *DescribeCdDeployTimeByProjectResponseData) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *DescribeCdDeployTimeByProjectResponseData) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *DescribeCdDeployTimeByProjectResponseData) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *DescribeCdDeployTimeByProjectResponseData) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetTotal

`func (o *DescribeCdDeployTimeByProjectResponseData) GetTotal() CdDeployTime`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DescribeCdDeployTimeByProjectResponseData) GetTotalOk() (*CdDeployTime, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DescribeCdDeployTimeByProjectResponseData) SetTotal(v CdDeployTime)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DescribeCdDeployTimeByProjectResponseData) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


