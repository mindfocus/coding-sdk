# ScanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | 分支 | [optional] [default to ""]
**CIJob** | Pointer to **int32** | 构建 id | [optional] 
**DepotId** | Pointer to **int32** | 仓库 id | [optional] 
**DepotName** | Pointer to **string** | 仓库名称 | [optional] [default to ""]
**IncrScan** | Pointer to **bool** | 是否增量扫描 | [optional] [default to false]
**ProjectId** | Pointer to **int32** | 项目 id | [optional] 
**ProjectName** | Pointer to **string** | 项目名称 | [optional] [default to ""]
**ScanId** | Pointer to **int32** | 扫描任务id | [optional] 
**ScanName** | Pointer to **string** | 扫描任务名称 | [optional] [default to ""]
**ScanSchemeName** | Pointer to **string** | 扫描方案 | [optional] [default to ""]
**Sha** | Pointer to **string** | Sha 值 | [optional] [default to ""]
**Status** | Pointer to **string** | 状态 | [optional] [default to ""]

## Methods

### NewScanInfo

`func NewScanInfo() *ScanInfo`

NewScanInfo instantiates a new ScanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScanInfoWithDefaults

`func NewScanInfoWithDefaults() *ScanInfo`

NewScanInfoWithDefaults instantiates a new ScanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ScanInfo) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ScanInfo) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ScanInfo) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ScanInfo) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCIJob

`func (o *ScanInfo) GetCIJob() int32`

GetCIJob returns the CIJob field if non-nil, zero value otherwise.

### GetCIJobOk

`func (o *ScanInfo) GetCIJobOk() (*int32, bool)`

GetCIJobOk returns a tuple with the CIJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCIJob

`func (o *ScanInfo) SetCIJob(v int32)`

SetCIJob sets CIJob field to given value.

### HasCIJob

`func (o *ScanInfo) HasCIJob() bool`

HasCIJob returns a boolean if a field has been set.

### GetDepotId

`func (o *ScanInfo) GetDepotId() int32`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *ScanInfo) GetDepotIdOk() (*int32, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *ScanInfo) SetDepotId(v int32)`

SetDepotId sets DepotId field to given value.

### HasDepotId

`func (o *ScanInfo) HasDepotId() bool`

HasDepotId returns a boolean if a field has been set.

### GetDepotName

`func (o *ScanInfo) GetDepotName() string`

GetDepotName returns the DepotName field if non-nil, zero value otherwise.

### GetDepotNameOk

`func (o *ScanInfo) GetDepotNameOk() (*string, bool)`

GetDepotNameOk returns a tuple with the DepotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotName

`func (o *ScanInfo) SetDepotName(v string)`

SetDepotName sets DepotName field to given value.

### HasDepotName

`func (o *ScanInfo) HasDepotName() bool`

HasDepotName returns a boolean if a field has been set.

### GetIncrScan

`func (o *ScanInfo) GetIncrScan() bool`

GetIncrScan returns the IncrScan field if non-nil, zero value otherwise.

### GetIncrScanOk

`func (o *ScanInfo) GetIncrScanOk() (*bool, bool)`

GetIncrScanOk returns a tuple with the IncrScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrScan

`func (o *ScanInfo) SetIncrScan(v bool)`

SetIncrScan sets IncrScan field to given value.

### HasIncrScan

`func (o *ScanInfo) HasIncrScan() bool`

HasIncrScan returns a boolean if a field has been set.

### GetProjectId

`func (o *ScanInfo) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ScanInfo) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ScanInfo) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ScanInfo) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectName

`func (o *ScanInfo) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *ScanInfo) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *ScanInfo) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *ScanInfo) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### GetScanId

`func (o *ScanInfo) GetScanId() int32`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ScanInfo) GetScanIdOk() (*int32, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ScanInfo) SetScanId(v int32)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ScanInfo) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetScanName

`func (o *ScanInfo) GetScanName() string`

GetScanName returns the ScanName field if non-nil, zero value otherwise.

### GetScanNameOk

`func (o *ScanInfo) GetScanNameOk() (*string, bool)`

GetScanNameOk returns a tuple with the ScanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanName

`func (o *ScanInfo) SetScanName(v string)`

SetScanName sets ScanName field to given value.

### HasScanName

`func (o *ScanInfo) HasScanName() bool`

HasScanName returns a boolean if a field has been set.

### GetScanSchemeName

`func (o *ScanInfo) GetScanSchemeName() string`

GetScanSchemeName returns the ScanSchemeName field if non-nil, zero value otherwise.

### GetScanSchemeNameOk

`func (o *ScanInfo) GetScanSchemeNameOk() (*string, bool)`

GetScanSchemeNameOk returns a tuple with the ScanSchemeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanSchemeName

`func (o *ScanInfo) SetScanSchemeName(v string)`

SetScanSchemeName sets ScanSchemeName field to given value.

### HasScanSchemeName

`func (o *ScanInfo) HasScanSchemeName() bool`

HasScanSchemeName returns a boolean if a field has been set.

### GetSha

`func (o *ScanInfo) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *ScanInfo) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *ScanInfo) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *ScanInfo) HasSha() bool`

HasSha returns a boolean if a field has been set.

### GetStatus

`func (o *ScanInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScanInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScanInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScanInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


