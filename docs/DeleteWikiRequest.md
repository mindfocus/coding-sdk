# DeleteWikiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iid** | **int64** | wiki编号 | 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDeleteWikiRequest

`func NewDeleteWikiRequest(iid int64, projectName string, ) *DeleteWikiRequest`

NewDeleteWikiRequest instantiates a new DeleteWikiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteWikiRequestWithDefaults

`func NewDeleteWikiRequestWithDefaults() *DeleteWikiRequest`

NewDeleteWikiRequestWithDefaults instantiates a new DeleteWikiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIid

`func (o *DeleteWikiRequest) GetIid() int64`

GetIid returns the Iid field if non-nil, zero value otherwise.

### GetIidOk

`func (o *DeleteWikiRequest) GetIidOk() (*int64, bool)`

GetIidOk returns a tuple with the Iid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIid

`func (o *DeleteWikiRequest) SetIid(v int64)`

SetIid sets Iid field to given value.


### GetProjectName

`func (o *DeleteWikiRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteWikiRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteWikiRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


