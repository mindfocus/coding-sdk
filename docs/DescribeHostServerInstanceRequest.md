# DescribeHostServerInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | 部署账号 从 DescribeCdPipeline 获取 | [default to ""]
**ServerGroupName** | **string** | 主机组名称 从 DescribeCdPipeline 获取 | [default to ""]

## Methods

### NewDescribeHostServerInstanceRequest

`func NewDescribeHostServerInstanceRequest(account string, serverGroupName string, ) *DescribeHostServerInstanceRequest`

NewDescribeHostServerInstanceRequest instantiates a new DescribeHostServerInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeHostServerInstanceRequestWithDefaults

`func NewDescribeHostServerInstanceRequestWithDefaults() *DescribeHostServerInstanceRequest`

NewDescribeHostServerInstanceRequestWithDefaults instantiates a new DescribeHostServerInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DescribeHostServerInstanceRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DescribeHostServerInstanceRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DescribeHostServerInstanceRequest) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetServerGroupName

`func (o *DescribeHostServerInstanceRequest) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *DescribeHostServerInstanceRequest) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *DescribeHostServerInstanceRequest) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


