# DescribeAgentSecret200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**AgentMachineSecret**](AgentMachineSecret.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeAgentSecret200ResponseResponse

`func NewDescribeAgentSecret200ResponseResponse() *DescribeAgentSecret200ResponseResponse`

NewDescribeAgentSecret200ResponseResponse instantiates a new DescribeAgentSecret200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeAgentSecret200ResponseResponseWithDefaults

`func NewDescribeAgentSecret200ResponseResponseWithDefaults() *DescribeAgentSecret200ResponseResponse`

NewDescribeAgentSecret200ResponseResponseWithDefaults instantiates a new DescribeAgentSecret200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeAgentSecret200ResponseResponse) GetData() AgentMachineSecret`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeAgentSecret200ResponseResponse) GetDataOk() (*AgentMachineSecret, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeAgentSecret200ResponseResponse) SetData(v AgentMachineSecret)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeAgentSecret200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeAgentSecret200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeAgentSecret200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeAgentSecret200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeAgentSecret200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


