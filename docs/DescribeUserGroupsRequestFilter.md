# DescribeUserGroupsRequestFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | Pointer to **string** | 用户组名称 | [optional] [default to ""]

## Methods

### NewDescribeUserGroupsRequestFilter

`func NewDescribeUserGroupsRequestFilter() *DescribeUserGroupsRequestFilter`

NewDescribeUserGroupsRequestFilter instantiates a new DescribeUserGroupsRequestFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeUserGroupsRequestFilterWithDefaults

`func NewDescribeUserGroupsRequestFilterWithDefaults() *DescribeUserGroupsRequestFilter`

NewDescribeUserGroupsRequestFilterWithDefaults instantiates a new DescribeUserGroupsRequestFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *DescribeUserGroupsRequestFilter) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DescribeUserGroupsRequestFilter) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DescribeUserGroupsRequestFilter) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DescribeUserGroupsRequestFilter) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


