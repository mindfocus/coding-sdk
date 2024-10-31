# DescribeTeamAdminMembers200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**TeamAdminMemberData**](TeamAdminMemberData.md) |  | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeTeamAdminMembers200ResponseResponse

`func NewDescribeTeamAdminMembers200ResponseResponse() *DescribeTeamAdminMembers200ResponseResponse`

NewDescribeTeamAdminMembers200ResponseResponse instantiates a new DescribeTeamAdminMembers200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeTeamAdminMembers200ResponseResponseWithDefaults

`func NewDescribeTeamAdminMembers200ResponseResponseWithDefaults() *DescribeTeamAdminMembers200ResponseResponse`

NewDescribeTeamAdminMembers200ResponseResponseWithDefaults instantiates a new DescribeTeamAdminMembers200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DescribeTeamAdminMembers200ResponseResponse) GetData() TeamAdminMemberData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DescribeTeamAdminMembers200ResponseResponse) GetDataOk() (*TeamAdminMemberData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DescribeTeamAdminMembers200ResponseResponse) SetData(v TeamAdminMemberData)`

SetData sets Data field to given value.

### HasData

`func (o *DescribeTeamAdminMembers200ResponseResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeTeamAdminMembers200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeTeamAdminMembers200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeTeamAdminMembers200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeTeamAdminMembers200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


