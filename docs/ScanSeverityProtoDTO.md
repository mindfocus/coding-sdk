# ScanSeverityProtoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | Pointer to **int64** | 创建时间 | [optional] 
**Creator** | Pointer to **string** | 创建人 | [optional] [default to ""]
**EnableErrorThreshold** | Pointer to **bool** | 启用错误信息阈值 | [optional] [default to false]
**EnableFatalThreshold** | Pointer to **bool** | 启用致命信息阈值 | [optional] [default to false]
**EnableInfoThreshold** | Pointer to **bool** | 启用提示信息阈值 | [optional] [default to false]
**EnableWarningThreshold** | Pointer to **bool** | 启用警告信息阈值 | [optional] [default to false]
**Error** | Pointer to **int64** | 错误问题 | [optional] 
**ErrorThreshold** | Pointer to **int32** | 错误信息阈值 | [optional] 
**Fatal** | Pointer to **int64** | 致命问题 | [optional] 
**FatalThreshold** | Pointer to **int64** | 致命信息阈值 | [optional] 
**GlobalSwitch** | Pointer to **bool** | 全局开关 | [optional] [default to false]
**Increment** | Pointer to **bool** | 增量扫描 | [optional] [default to false]
**Info** | Pointer to **int64** | 提示问题 | [optional] 
**InfoThreshold** | Pointer to **int64** | 提示信息阈值 | [optional] 
**Status** | Pointer to **string** | 完成状态 | [optional] [default to ""]
**StatusCheck** | Pointer to **string** | 扫描状态 | [optional] [default to ""]
**Warning** | Pointer to **int64** | 警告问题 | [optional] 
**WarningThreshold** | Pointer to **int64** | 警告信息阈值 | [optional] 

## Methods

### NewScanSeverityProtoDTO

`func NewScanSeverityProtoDTO() *ScanSeverityProtoDTO`

NewScanSeverityProtoDTO instantiates a new ScanSeverityProtoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanSeverityProtoDTOWithDefaults

`func NewScanSeverityProtoDTOWithDefaults() *ScanSeverityProtoDTO`

NewScanSeverityProtoDTOWithDefaults instantiates a new ScanSeverityProtoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *ScanSeverityProtoDTO) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *ScanSeverityProtoDTO) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *ScanSeverityProtoDTO) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *ScanSeverityProtoDTO) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetCreator

`func (o *ScanSeverityProtoDTO) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ScanSeverityProtoDTO) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ScanSeverityProtoDTO) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *ScanSeverityProtoDTO) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetEnableErrorThreshold

`func (o *ScanSeverityProtoDTO) GetEnableErrorThreshold() bool`

GetEnableErrorThreshold returns the EnableErrorThreshold field if non-nil, zero value otherwise.

### GetEnableErrorThresholdOk

`func (o *ScanSeverityProtoDTO) GetEnableErrorThresholdOk() (*bool, bool)`

GetEnableErrorThresholdOk returns a tuple with the EnableErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableErrorThreshold

`func (o *ScanSeverityProtoDTO) SetEnableErrorThreshold(v bool)`

SetEnableErrorThreshold sets EnableErrorThreshold field to given value.

### HasEnableErrorThreshold

`func (o *ScanSeverityProtoDTO) HasEnableErrorThreshold() bool`

HasEnableErrorThreshold returns a boolean if a field has been set.

### GetEnableFatalThreshold

`func (o *ScanSeverityProtoDTO) GetEnableFatalThreshold() bool`

GetEnableFatalThreshold returns the EnableFatalThreshold field if non-nil, zero value otherwise.

### GetEnableFatalThresholdOk

`func (o *ScanSeverityProtoDTO) GetEnableFatalThresholdOk() (*bool, bool)`

GetEnableFatalThresholdOk returns a tuple with the EnableFatalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFatalThreshold

`func (o *ScanSeverityProtoDTO) SetEnableFatalThreshold(v bool)`

SetEnableFatalThreshold sets EnableFatalThreshold field to given value.

### HasEnableFatalThreshold

`func (o *ScanSeverityProtoDTO) HasEnableFatalThreshold() bool`

HasEnableFatalThreshold returns a boolean if a field has been set.

### GetEnableInfoThreshold

`func (o *ScanSeverityProtoDTO) GetEnableInfoThreshold() bool`

GetEnableInfoThreshold returns the EnableInfoThreshold field if non-nil, zero value otherwise.

### GetEnableInfoThresholdOk

`func (o *ScanSeverityProtoDTO) GetEnableInfoThresholdOk() (*bool, bool)`

GetEnableInfoThresholdOk returns a tuple with the EnableInfoThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableInfoThreshold

`func (o *ScanSeverityProtoDTO) SetEnableInfoThreshold(v bool)`

SetEnableInfoThreshold sets EnableInfoThreshold field to given value.

### HasEnableInfoThreshold

`func (o *ScanSeverityProtoDTO) HasEnableInfoThreshold() bool`

HasEnableInfoThreshold returns a boolean if a field has been set.

### GetEnableWarningThreshold

`func (o *ScanSeverityProtoDTO) GetEnableWarningThreshold() bool`

GetEnableWarningThreshold returns the EnableWarningThreshold field if non-nil, zero value otherwise.

### GetEnableWarningThresholdOk

`func (o *ScanSeverityProtoDTO) GetEnableWarningThresholdOk() (*bool, bool)`

GetEnableWarningThresholdOk returns a tuple with the EnableWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWarningThreshold

`func (o *ScanSeverityProtoDTO) SetEnableWarningThreshold(v bool)`

SetEnableWarningThreshold sets EnableWarningThreshold field to given value.

### HasEnableWarningThreshold

`func (o *ScanSeverityProtoDTO) HasEnableWarningThreshold() bool`

HasEnableWarningThreshold returns a boolean if a field has been set.

### GetError

`func (o *ScanSeverityProtoDTO) GetError() int64`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScanSeverityProtoDTO) GetErrorOk() (*int64, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScanSeverityProtoDTO) SetError(v int64)`

SetError sets Error field to given value.

### HasError

`func (o *ScanSeverityProtoDTO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorThreshold

`func (o *ScanSeverityProtoDTO) GetErrorThreshold() int32`

GetErrorThreshold returns the ErrorThreshold field if non-nil, zero value otherwise.

### GetErrorThresholdOk

`func (o *ScanSeverityProtoDTO) GetErrorThresholdOk() (*int32, bool)`

GetErrorThresholdOk returns a tuple with the ErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorThreshold

`func (o *ScanSeverityProtoDTO) SetErrorThreshold(v int32)`

SetErrorThreshold sets ErrorThreshold field to given value.

### HasErrorThreshold

`func (o *ScanSeverityProtoDTO) HasErrorThreshold() bool`

HasErrorThreshold returns a boolean if a field has been set.

### GetFatal

`func (o *ScanSeverityProtoDTO) GetFatal() int64`

GetFatal returns the Fatal field if non-nil, zero value otherwise.

### GetFatalOk

`func (o *ScanSeverityProtoDTO) GetFatalOk() (*int64, bool)`

GetFatalOk returns a tuple with the Fatal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatal

`func (o *ScanSeverityProtoDTO) SetFatal(v int64)`

SetFatal sets Fatal field to given value.

### HasFatal

`func (o *ScanSeverityProtoDTO) HasFatal() bool`

HasFatal returns a boolean if a field has been set.

### GetFatalThreshold

`func (o *ScanSeverityProtoDTO) GetFatalThreshold() int64`

GetFatalThreshold returns the FatalThreshold field if non-nil, zero value otherwise.

### GetFatalThresholdOk

`func (o *ScanSeverityProtoDTO) GetFatalThresholdOk() (*int64, bool)`

GetFatalThresholdOk returns a tuple with the FatalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFatalThreshold

`func (o *ScanSeverityProtoDTO) SetFatalThreshold(v int64)`

SetFatalThreshold sets FatalThreshold field to given value.

### HasFatalThreshold

`func (o *ScanSeverityProtoDTO) HasFatalThreshold() bool`

HasFatalThreshold returns a boolean if a field has been set.

### GetGlobalSwitch

`func (o *ScanSeverityProtoDTO) GetGlobalSwitch() bool`

GetGlobalSwitch returns the GlobalSwitch field if non-nil, zero value otherwise.

### GetGlobalSwitchOk

`func (o *ScanSeverityProtoDTO) GetGlobalSwitchOk() (*bool, bool)`

GetGlobalSwitchOk returns a tuple with the GlobalSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSwitch

`func (o *ScanSeverityProtoDTO) SetGlobalSwitch(v bool)`

SetGlobalSwitch sets GlobalSwitch field to given value.

### HasGlobalSwitch

`func (o *ScanSeverityProtoDTO) HasGlobalSwitch() bool`

HasGlobalSwitch returns a boolean if a field has been set.

### GetIncrement

`func (o *ScanSeverityProtoDTO) GetIncrement() bool`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *ScanSeverityProtoDTO) GetIncrementOk() (*bool, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *ScanSeverityProtoDTO) SetIncrement(v bool)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *ScanSeverityProtoDTO) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.

### GetInfo

`func (o *ScanSeverityProtoDTO) GetInfo() int64`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScanSeverityProtoDTO) GetInfoOk() (*int64, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScanSeverityProtoDTO) SetInfo(v int64)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScanSeverityProtoDTO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetInfoThreshold

`func (o *ScanSeverityProtoDTO) GetInfoThreshold() int64`

GetInfoThreshold returns the InfoThreshold field if non-nil, zero value otherwise.

### GetInfoThresholdOk

`func (o *ScanSeverityProtoDTO) GetInfoThresholdOk() (*int64, bool)`

GetInfoThresholdOk returns a tuple with the InfoThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoThreshold

`func (o *ScanSeverityProtoDTO) SetInfoThreshold(v int64)`

SetInfoThreshold sets InfoThreshold field to given value.

### HasInfoThreshold

`func (o *ScanSeverityProtoDTO) HasInfoThreshold() bool`

HasInfoThreshold returns a boolean if a field has been set.

### GetStatus

`func (o *ScanSeverityProtoDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScanSeverityProtoDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScanSeverityProtoDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScanSeverityProtoDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCheck

`func (o *ScanSeverityProtoDTO) GetStatusCheck() string`

GetStatusCheck returns the StatusCheck field if non-nil, zero value otherwise.

### GetStatusCheckOk

`func (o *ScanSeverityProtoDTO) GetStatusCheckOk() (*string, bool)`

GetStatusCheckOk returns a tuple with the StatusCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCheck

`func (o *ScanSeverityProtoDTO) SetStatusCheck(v string)`

SetStatusCheck sets StatusCheck field to given value.

### HasStatusCheck

`func (o *ScanSeverityProtoDTO) HasStatusCheck() bool`

HasStatusCheck returns a boolean if a field has been set.

### GetWarning

`func (o *ScanSeverityProtoDTO) GetWarning() int64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ScanSeverityProtoDTO) GetWarningOk() (*int64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ScanSeverityProtoDTO) SetWarning(v int64)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ScanSeverityProtoDTO) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetWarningThreshold

`func (o *ScanSeverityProtoDTO) GetWarningThreshold() int64`

GetWarningThreshold returns the WarningThreshold field if non-nil, zero value otherwise.

### GetWarningThresholdOk

`func (o *ScanSeverityProtoDTO) GetWarningThresholdOk() (*int64, bool)`

GetWarningThresholdOk returns a tuple with the WarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningThreshold

`func (o *ScanSeverityProtoDTO) SetWarningThreshold(v int64)`

SetWarningThreshold sets WarningThreshold field to given value.

### HasWarningThreshold

`func (o *ScanSeverityProtoDTO) HasWarningThreshold() bool`

HasWarningThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


