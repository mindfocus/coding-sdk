# DescribeCodingCIBuildLogData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Log** | Pointer to **string** | 日志 | [optional] [default to ""]
**MoreData** | Pointer to **bool** | 是否有更多的日志 | [optional] [default to false]
**TextDelivered** | Pointer to **int32** | 当前展示日志长度 | [optional] 
**TextSize** | Pointer to **int32** | 总日志长度 | [optional] 

## Methods

### NewDescribeCodingCIBuildLogData

`func NewDescribeCodingCIBuildLogData() *DescribeCodingCIBuildLogData`

NewDescribeCodingCIBuildLogData instantiates a new DescribeCodingCIBuildLogData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeCodingCIBuildLogDataWithDefaults

`func NewDescribeCodingCIBuildLogDataWithDefaults() *DescribeCodingCIBuildLogData`

NewDescribeCodingCIBuildLogDataWithDefaults instantiates a new DescribeCodingCIBuildLogData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLog

`func (o *DescribeCodingCIBuildLogData) GetLog() string`

GetLog returns the Log field if non-nil, zero value otherwise.

### GetLogOk

`func (o *DescribeCodingCIBuildLogData) GetLogOk() (*string, bool)`

GetLogOk returns a tuple with the Log field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog

`func (o *DescribeCodingCIBuildLogData) SetLog(v string)`

SetLog sets Log field to given value.

### HasLog

`func (o *DescribeCodingCIBuildLogData) HasLog() bool`

HasLog returns a boolean if a field has been set.

### GetMoreData

`func (o *DescribeCodingCIBuildLogData) GetMoreData() bool`

GetMoreData returns the MoreData field if non-nil, zero value otherwise.

### GetMoreDataOk

`func (o *DescribeCodingCIBuildLogData) GetMoreDataOk() (*bool, bool)`

GetMoreDataOk returns a tuple with the MoreData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreData

`func (o *DescribeCodingCIBuildLogData) SetMoreData(v bool)`

SetMoreData sets MoreData field to given value.

### HasMoreData

`func (o *DescribeCodingCIBuildLogData) HasMoreData() bool`

HasMoreData returns a boolean if a field has been set.

### GetTextDelivered

`func (o *DescribeCodingCIBuildLogData) GetTextDelivered() int32`

GetTextDelivered returns the TextDelivered field if non-nil, zero value otherwise.

### GetTextDeliveredOk

`func (o *DescribeCodingCIBuildLogData) GetTextDeliveredOk() (*int32, bool)`

GetTextDeliveredOk returns a tuple with the TextDelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextDelivered

`func (o *DescribeCodingCIBuildLogData) SetTextDelivered(v int32)`

SetTextDelivered sets TextDelivered field to given value.

### HasTextDelivered

`func (o *DescribeCodingCIBuildLogData) HasTextDelivered() bool`

HasTextDelivered returns a boolean if a field has been set.

### GetTextSize

`func (o *DescribeCodingCIBuildLogData) GetTextSize() int32`

GetTextSize returns the TextSize field if non-nil, zero value otherwise.

### GetTextSizeOk

`func (o *DescribeCodingCIBuildLogData) GetTextSizeOk() (*int32, bool)`

GetTextSizeOk returns a tuple with the TextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSize

`func (o *DescribeCodingCIBuildLogData) SetTextSize(v int32)`

SetTextSize sets TextSize field to given value.

### HasTextSize

`func (o *DescribeCodingCIBuildLogData) HasTextSize() bool`

HasTextSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


