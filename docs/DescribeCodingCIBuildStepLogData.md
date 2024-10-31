# DescribeCodingCIBuildStepLogData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Log** | Pointer to **string** | 日志 | [optional] [default to ""]
**MoreData** | Pointer to **bool** | 是否有更多数据 | [optional] [default to false]
**TextDelivered** | Pointer to **int32** | 当前展示总长度 | [optional] 
**TextSize** | Pointer to **int32** | 总长度 | [optional] 

## Methods

### NewDescribeCodingCIBuildStepLogData

`func NewDescribeCodingCIBuildStepLogData() *DescribeCodingCIBuildStepLogData`

NewDescribeCodingCIBuildStepLogData instantiates a new DescribeCodingCIBuildStepLogData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildStepLogDataWithDefaults

`func NewDescribeCodingCIBuildStepLogDataWithDefaults() *DescribeCodingCIBuildStepLogData`

NewDescribeCodingCIBuildStepLogDataWithDefaults instantiates a new DescribeCodingCIBuildStepLogData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLog

`func (o *DescribeCodingCIBuildStepLogData) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *DescribeCodingCIBuildStepLogData) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *DescribeCodingCIBuildStepLogData) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *DescribeCodingCIBuildStepLogData) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetMoreData

`func (o *DescribeCodingCIBuildStepLogData) GetMoreData() bool`

GetMoreData returns the MoreData field if non-nil, zero value otherwise.

### GetMoreDataOk

`func (o *DescribeCodingCIBuildStepLogData) GetMoreDataOk() (*bool, bool)`

GetMoreDataOk returns a tuple with the MoreData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreData

`func (o *DescribeCodingCIBuildStepLogData) SetMoreData(v bool)`

SetMoreData sets MoreData field to given value.

### HasMoreData

`func (o *DescribeCodingCIBuildStepLogData) HasMoreData() bool`

HasMoreData returns a boolean if a field has been set.

### GetTextDelivered

`func (o *DescribeCodingCIBuildStepLogData) GetTextDelivered() int32`

GetTextDelivered returns the TextDelivered field if non-nil, zero value otherwise.

### GetTextDeliveredOk

`func (o *DescribeCodingCIBuildStepLogData) GetTextDeliveredOk() (*int32, bool)`

GetTextDeliveredOk returns a tuple with the TextDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDelivered

`func (o *DescribeCodingCIBuildStepLogData) SetTextDelivered(v int32)`

SetTextDelivered sets TextDelivered field to given value.

### HasTextDelivered

`func (o *DescribeCodingCIBuildStepLogData) HasTextDelivered() bool`

HasTextDelivered returns a boolean if a field has been set.

### GetTextSize

`func (o *DescribeCodingCIBuildStepLogData) GetTextSize() int32`

GetTextSize returns the TextSize field if non-nil, zero value otherwise.

### GetTextSizeOk

`func (o *DescribeCodingCIBuildStepLogData) GetTextSizeOk() (*int32, bool)`

GetTextSizeOk returns a tuple with the TextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSize

`func (o *DescribeCodingCIBuildStepLogData) SetTextSize(v int32)`

SetTextSize sets TextSize field to given value.

### HasTextSize

`func (o *DescribeCodingCIBuildStepLogData) HasTextSize() bool`

HasTextSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


