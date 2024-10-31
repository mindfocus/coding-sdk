# TriggerCodingScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncrScan** | Pointer to **bool** | 增量扫描 | [optional] 
**ScanId** | **int32** | 扫描任务 id | 

## Methods

### NewTriggerCodingScanRequest

`func NewTriggerCodingScanRequest(scanId int32, ) *TriggerCodingScanRequest`

NewTriggerCodingScanRequest instantiates a new TriggerCodingScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCodingScanRequestWithDefaults

`func NewTriggerCodingScanRequestWithDefaults() *TriggerCodingScanRequest`

NewTriggerCodingScanRequestWithDefaults instantiates a new TriggerCodingScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncrScan

`func (o *TriggerCodingScanRequest) GetIncrScan() bool`

GetIncrScan returns the IncrScan field if non-nil, zero value otherwise.

### GetIncrScanOk

`func (o *TriggerCodingScanRequest) GetIncrScanOk() (*bool, bool)`

GetIncrScanOk returns a tuple with the IncrScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrScan

`func (o *TriggerCodingScanRequest) SetIncrScan(v bool)`

SetIncrScan sets IncrScan field to given value.

### HasIncrScan

`func (o *TriggerCodingScanRequest) HasIncrScan() bool`

HasIncrScan returns a boolean if a field has been set.

### GetScanId

`func (o *TriggerCodingScanRequest) GetScanId() int32`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *TriggerCodingScanRequest) GetScanIdOk() (*int32, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *TriggerCodingScanRequest) SetScanId(v int32)`

SetScanId sets ScanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


