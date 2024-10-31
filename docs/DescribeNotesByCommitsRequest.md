# DescribeNotesByCommitsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | **[]string** | 请求sha值列表 | 
**DepotId** | **int64** | 仓库id | 
**DepotPath** | Pointer to **string** | 仓库路径，DepotId与DepotPath二选一即可 | [optional] 
**NoteRef** | **string** | 查询的ref地址 | 

## Methods

### NewDescribeNotesByCommitsRequest

`func NewDescribeNotesByCommitsRequest(commits []string, depotId int64, noteRef string, ) *DescribeNotesByCommitsRequest`

NewDescribeNotesByCommitsRequest instantiates a new DescribeNotesByCommitsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeNotesByCommitsRequestWithDefaults

`func NewDescribeNotesByCommitsRequestWithDefaults() *DescribeNotesByCommitsRequest`

NewDescribeNotesByCommitsRequestWithDefaults instantiates a new DescribeNotesByCommitsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *DescribeNotesByCommitsRequest) GetCommits() []string`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *DescribeNotesByCommitsRequest) GetCommitsOk() (*[]string, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *DescribeNotesByCommitsRequest) SetCommits(v []string)`

SetCommits sets Commits field to given value.


### GetDepotId

`func (o *DescribeNotesByCommitsRequest) GetDepotId() int64`

GetDepotId returns the DepotId field if non-nil, zero value otherwise.

### GetDepotIdOk

`func (o *DescribeNotesByCommitsRequest) GetDepotIdOk() (*int64, bool)`

GetDepotIdOk returns a tuple with the DepotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotId

`func (o *DescribeNotesByCommitsRequest) SetDepotId(v int64)`

SetDepotId sets DepotId field to given value.


### GetDepotPath

`func (o *DescribeNotesByCommitsRequest) GetDepotPath() string`

GetDepotPath returns the DepotPath field if non-nil, zero value otherwise.

### GetDepotPathOk

`func (o *DescribeNotesByCommitsRequest) GetDepotPathOk() (*string, bool)`

GetDepotPathOk returns a tuple with the DepotPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepotPath

`func (o *DescribeNotesByCommitsRequest) SetDepotPath(v string)`

SetDepotPath sets DepotPath field to given value.

### HasDepotPath

`func (o *DescribeNotesByCommitsRequest) HasDepotPath() bool`

HasDepotPath returns a boolean if a field has been set.

### GetNoteRef

`func (o *DescribeNotesByCommitsRequest) GetNoteRef() string`

GetNoteRef returns the NoteRef field if non-nil, zero value otherwise.

### GetNoteRefOk

`func (o *DescribeNotesByCommitsRequest) GetNoteRefOk() (*string, bool)`

GetNoteRefOk returns a tuple with the NoteRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteRef

`func (o *DescribeNotesByCommitsRequest) SetNoteRef(v string)`

SetNoteRef sets NoteRef field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


