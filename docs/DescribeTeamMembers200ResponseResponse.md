# DescribeTeamMembers200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TeamMemberData**](TeamMemberData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeTeamMembers200ResponseResponse

`func NewDescribeTeamMembers200ResponseResponse() *DescribeTeamMembers200ResponseResponse`

NewDescribeTeamMembers200ResponseResponse instantiates a new DescribeTeamMembers200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamMembers200ResponseResponseWithDefaults

`func NewDescribeTeamMembers200ResponseResponseWithDefaults() *DescribeTeamMembers200ResponseResponse`

NewDescribeTeamMembers200ResponseResponseWithDefaults instantiates a new DescribeTeamMembers200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeTeamMembers200ResponseResponse) GetData() TeamMemberData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeTeamMembers200ResponseResponse) GetDataOk() (*TeamMemberData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeTeamMembers200ResponseResponse) SetData(v TeamMemberData)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeTeamMembers200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeTeamMembers200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeTeamMembers200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeTeamMembers200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeTeamMembers200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


