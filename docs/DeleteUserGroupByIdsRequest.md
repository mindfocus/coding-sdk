# DeleteUserGroupByIdsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupIds** | **[]int64** | 用户组id列表 | 

## Methods

### NewDeleteUserGroupByIdsRequest

`func NewDeleteUserGroupByIdsRequest(groupIds []int64, ) *DeleteUserGroupByIdsRequest`

NewDeleteUserGroupByIdsRequest instantiates a new DeleteUserGroupByIdsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteUserGroupByIdsRequestWithDefaults

`func NewDeleteUserGroupByIdsRequestWithDefaults() *DeleteUserGroupByIdsRequest`

NewDeleteUserGroupByIdsRequestWithDefaults instantiates a new DeleteUserGroupByIdsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupIds

`func (o *DeleteUserGroupByIdsRequest) GetGroupIds() []int64`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *DeleteUserGroupByIdsRequest) GetGroupIdsOk() (*[]int64, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *DeleteUserGroupByIdsRequest) SetGroupIds(v []int64)`

SetGroupIds sets GroupIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


