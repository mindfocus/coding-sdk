# ServiceHookEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupLabel** | Pointer to **string** | 分组名 | [optional] [default to ""]
**GroupName** | Pointer to **string** | 分组标识 | [optional] [default to ""]
**Label** | Pointer to **string** | 事件名 | [optional] [default to ""]
**Name** | Pointer to **string** | 事件标识 | [optional] [default to ""]

## Methods

### NewServiceHookEvent

`func NewServiceHookEvent() *ServiceHookEvent`

NewServiceHookEvent instantiates a new ServiceHookEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceHookEventWithDefaults

`func NewServiceHookEventWithDefaults() *ServiceHookEvent`

NewServiceHookEventWithDefaults instantiates a new ServiceHookEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupLabel

`func (o *ServiceHookEvent) GetGroupLabel() string`

GetGroupLabel returns the GroupLabel field if non-nil, zero value otherwise.

### GetGroupLabelOk

`func (o *ServiceHookEvent) GetGroupLabelOk() (*string, bool)`

GetGroupLabelOk returns a tuple with the GroupLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLabel

`func (o *ServiceHookEvent) SetGroupLabel(v string)`

SetGroupLabel sets GroupLabel field to given value.

### HasGroupLabel

`func (o *ServiceHookEvent) HasGroupLabel() bool`

HasGroupLabel returns a boolean if a field has been set.

### GetGroupName

`func (o *ServiceHookEvent) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ServiceHookEvent) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ServiceHookEvent) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ServiceHookEvent) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceHookEvent) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceHookEvent) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceHookEvent) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceHookEvent) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *ServiceHookEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceHookEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceHookEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceHookEvent) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


