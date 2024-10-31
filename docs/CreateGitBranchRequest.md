# CreateGitBranchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | **string** | 分支名称 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**Note** | Pointer to **string** | 分支备注 | [optional] 
**StartPoint** | **string** | 创建来源的分支名称或者commitId | 

## Methods

### NewCreateGitBranchRequest

`func NewCreateGitBranchRequest(branchName string, depotId int64, startPoint string, ) *CreateGitBranchRequest`

NewCreateGitBranchRequest instantiates a new CreateGitBranchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGitBranchRequestWithDefaults

`func NewCreateGitBranchRequestWithDefaults() *CreateGitBranchRequest`

NewCreateGitBranchRequestWithDefaults instantiates a new CreateGitBranchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *CreateGitBranchRequest) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *CreateGitBranchRequest) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *CreateGitBranchRequest) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.


### GetDepotId

`func (o *CreateGitBranchRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *CreateGitBranchRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *CreateGitBranchRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *CreateGitBranchRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *CreateGitBranchRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *CreateGitBranchRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *CreateGitBranchRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetNote

`func (o *CreateGitBranchRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *CreateGitBranchRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *CreateGitBranchRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *CreateGitBranchRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetStartPoint

`func (o *CreateGitBranchRequest) GetStartPoint() string`

GetStartPoint returns the StartPoint field if non-nil, zero value otherwise.

### GetStartPointOk

`func (o *CreateGitBranchRequest) GetStartPointOk() (*string, bool)`

GetStartPointOk returns a tuple with the StartPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPoint

`func (o *CreateGitBranchRequest) SetStartPoint(v string)`

SetStartPoint sets StartPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


