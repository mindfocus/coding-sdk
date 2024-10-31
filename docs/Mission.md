# Mission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **NullableInt64** | 返回码 | [optional] 
**Link** | Pointer to **NullableString** | 连接 | [optional] [default to ""]
**TargetId** | Pointer to **NullableInt64** | 目标id | [optional] 
**TargetProjectName** | Pointer to **NullableString** | 目标项目名称 | [optional] [default to ""]
**TargetType** | Pointer to **NullableString** | 类型 | [optional] [default to ""]
**Title** | Pointer to **NullableString** | 标题 | [optional] [default to ""]

## Methods

### NewMission

`func NewMission() *Mission`

NewMission instantiates a new Mission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMissionWithDefaults

`func NewMissionWithDefaults() *Mission`

NewMissionWithDefaults instantiates a new Mission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Mission) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Mission) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Mission) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *Mission) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Mission) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Mission) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetLink

`func (o *Mission) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Mission) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Mission) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Mission) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *Mission) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *Mission) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetTargetId

`func (o *Mission) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *Mission) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *Mission) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *Mission) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *Mission) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *Mission) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetProjectName

`func (o *Mission) GetTargetProjectName() string`

GetTargetProjectName returns the TargetProjectName field if non-nil, zero value otherwise.

### GetTargetProjectNameOk

`func (o *Mission) GetTargetProjectNameOk() (*string, bool)`

GetTargetProjectNameOk returns a tuple with the TargetProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetProjectName

`func (o *Mission) SetTargetProjectName(v string)`

SetTargetProjectName sets TargetProjectName field to given value.

### HasTargetProjectName

`func (o *Mission) HasTargetProjectName() bool`

HasTargetProjectName returns a boolean if a field has been set.

### SetTargetProjectNameNil

`func (o *Mission) SetTargetProjectNameNil(b bool)`

 SetTargetProjectNameNil sets the value for TargetProjectName to be an explicit nil

### UnsetTargetProjectName
`func (o *Mission) UnsetTargetProjectName()`

UnsetTargetProjectName ensures that no value is present for TargetProjectName, not even an explicit nil
### GetTargetType

`func (o *Mission) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *Mission) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *Mission) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *Mission) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *Mission) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *Mission) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTitle

`func (o *Mission) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Mission) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Mission) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Mission) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *Mission) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *Mission) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


