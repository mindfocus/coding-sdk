# CreateFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FoldName** | **string** | 文件夹名 | [default to ""]
**ParentId** | **int64** | 父文件夹ID | [default to 0]
**ProjectName** | **string** | 项目名 | [default to ""]

## Methods

### NewCreateFolderRequest

`func NewCreateFolderRequest(foldName string, parentId int64, projectName string, ) *CreateFolderRequest`

NewCreateFolderRequest instantiates a new CreateFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFolderRequestWithDefaults

`func NewCreateFolderRequestWithDefaults() *CreateFolderRequest`

NewCreateFolderRequestWithDefaults instantiates a new CreateFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFoldName

`func (o *CreateFolderRequest) GetFoldName() string`

GetFoldName returns the FoldName field if non-nil, zero value otherwise.

### GetFoldNameOk

`func (o *CreateFolderRequest) GetFoldNameOk() (*string, bool)`

GetFoldNameOk returns a tuple with the FoldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldName

`func (o *CreateFolderRequest) SetFoldName(v string)`

SetFoldName sets FoldName field to given value.


### GetParentId

`func (o *CreateFolderRequest) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateFolderRequest) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateFolderRequest) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetProjectName

`func (o *CreateFolderRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *CreateFolderRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *CreateFolderRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


