# DescribeTeamMemberByEmail200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamMember** | Pointer to [**ApiUserUserData**](ApiUserUserData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeTeamMemberByEmail200ResponseResponse

`func NewDescribeTeamMemberByEmail200ResponseResponse() *DescribeTeamMemberByEmail200ResponseResponse`

NewDescribeTeamMemberByEmail200ResponseResponse instantiates a new DescribeTeamMemberByEmail200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamMemberByEmail200ResponseResponseWithDefaults

`func NewDescribeTeamMemberByEmail200ResponseResponseWithDefaults() *DescribeTeamMemberByEmail200ResponseResponse`

NewDescribeTeamMemberByEmail200ResponseResponseWithDefaults instantiates a new DescribeTeamMemberByEmail200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamMember

`func (o *DescribeTeamMemberByEmail200ResponseResponse) GetTeamMember() ApiUserUserData`

GetTeamMember returns the TeamMember field if non-nil, zero value otherwise.

### GetTeamMemberOk

`func (o *DescribeTeamMemberByEmail200ResponseResponse) GetTeamMemberOk() (*ApiUserUserData, bool)`

GetTeamMemberOk returns a tuple with the TeamMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMember

`func (o *DescribeTeamMemberByEmail200ResponseResponse) SetTeamMember(v ApiUserUserData)`

SetTeamMember sets TeamMember field to given value.

### HasTeamMember

`func (o *DescribeTeamMemberByEmail200ResponseResponse) HasTeamMember() bool`

HasTeamMember returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeTeamMemberByEmail200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeTeamMemberByEmail200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeTeamMemberByEmail200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeTeamMemberByEmail200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


