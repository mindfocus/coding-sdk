# CdDeployOpenApiHostServerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | 部署状态 succeed:成功 failed:失败 running:运行中 | [default to ""]
**HumanReadableName** | **string** | 主机IP | [default to ""]
**Account** | **string** | 部署账号 | [default to ""]
**DeployedTime** | **int64** | 部署时间 | [default to 0]
**Zone** | **string** | 主机组名称 | [default to ""]
**CloudProvider** | **string** | 云主机类型 主机组部署默认为 hostserver | [default to "hostserver"]
**Ip** | **string** | 主机IP | [default to ""]
**Region** | **string** | 主机组名称 | [default to ""]
**ProviderType** | **string** | 云主机类型 主机组部署默认为 hostserver | [default to "hostserver"]
**Name** | **string** | 该主机部署主键 | [default to ""]

## Methods

### NewCdDeployOpenApiHostServerInstance

`func NewCdDeployOpenApiHostServerInstance(status string, humanReadableName string, account string, deployedTime int64, zone string, cloudProvider string, ip string, region string, providerType string, name string, ) *CdDeployOpenApiHostServerInstance`

NewCdDeployOpenApiHostServerInstance instantiates a new CdDeployOpenApiHostServerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdDeployOpenApiHostServerInstanceWithDefaults

`func NewCdDeployOpenApiHostServerInstanceWithDefaults() *CdDeployOpenApiHostServerInstance`

NewCdDeployOpenApiHostServerInstanceWithDefaults instantiates a new CdDeployOpenApiHostServerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CdDeployOpenApiHostServerInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CdDeployOpenApiHostServerInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CdDeployOpenApiHostServerInstance) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetHumanReadableName

`func (o *CdDeployOpenApiHostServerInstance) GetHumanReadableName() string`

GetHumanReadableName returns the HumanReadableName field if non-nil, zero value otherwise.

### GetHumanReadableNameOk

`func (o *CdDeployOpenApiHostServerInstance) GetHumanReadableNameOk() (*string, bool)`

GetHumanReadableNameOk returns a tuple with the HumanReadableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumanReadableName

`func (o *CdDeployOpenApiHostServerInstance) SetHumanReadableName(v string)`

SetHumanReadableName sets HumanReadableName field to given value.


### GetAccount

`func (o *CdDeployOpenApiHostServerInstance) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CdDeployOpenApiHostServerInstance) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CdDeployOpenApiHostServerInstance) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetDeployedTime

`func (o *CdDeployOpenApiHostServerInstance) GetDeployedTime() int64`

GetDeployedTime returns the DeployedTime field if non-nil, zero value otherwise.

### GetDeployedTimeOk

`func (o *CdDeployOpenApiHostServerInstance) GetDeployedTimeOk() (*int64, bool)`

GetDeployedTimeOk returns a tuple with the DeployedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedTime

`func (o *CdDeployOpenApiHostServerInstance) SetDeployedTime(v int64)`

SetDeployedTime sets DeployedTime field to given value.


### GetZone

`func (o *CdDeployOpenApiHostServerInstance) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *CdDeployOpenApiHostServerInstance) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *CdDeployOpenApiHostServerInstance) SetZone(v string)`

SetZone sets Zone field to given value.


### GetCloudProvider

`func (o *CdDeployOpenApiHostServerInstance) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *CdDeployOpenApiHostServerInstance) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *CdDeployOpenApiHostServerInstance) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetIp

`func (o *CdDeployOpenApiHostServerInstance) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CdDeployOpenApiHostServerInstance) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CdDeployOpenApiHostServerInstance) SetIp(v string)`

SetIp sets Ip field to given value.


### GetRegion

`func (o *CdDeployOpenApiHostServerInstance) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CdDeployOpenApiHostServerInstance) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CdDeployOpenApiHostServerInstance) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetProviderType

`func (o *CdDeployOpenApiHostServerInstance) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *CdDeployOpenApiHostServerInstance) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *CdDeployOpenApiHostServerInstance) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetName

`func (o *CdDeployOpenApiHostServerInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CdDeployOpenApiHostServerInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CdDeployOpenApiHostServerInstance) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


