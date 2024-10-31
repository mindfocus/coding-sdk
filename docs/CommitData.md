# CommitData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commits** | Pointer to [**[]Commit**](Commit.md) | 提交信息 | [optional] 
**Page** | Pointer to [**PageInfo**](PageInfo.md) |  | [optional] 

## Methods

### NewCommitData

`func NewCommitData() *CommitData`

NewCommitData instantiates a new CommitData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitDataWithDefaults

`func NewCommitDataWithDefaults() *CommitData`

NewCommitDataWithDefaults instantiates a new CommitData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommits

`func (o *CommitData) GetCommits() []Commit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *CommitData) GetCommitsOk() (*[]Commit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *CommitData) SetCommits(v []Commit)`

SetCommits sets Commits field to given value.

### HasCommits

`func (o *CommitData) HasCommits() bool`

HasCommits returns a boolean if a field has been set.

### SetCommitsNil

`func (o *CommitData) SetCommitsNil(b bool)`

 SetCommitsNil sets the value for Commits to be an explicit nil

### UnsetCommits
`func (o *CommitData) UnsetCommits()`

UnsetCommits ensures that no value is present for Commits, not even an explicit nil
### GetPage

`func (o *CommitData) GetPage() PageInfo`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *CommitData) GetPageOk() (*PageInfo, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *CommitData) SetPage(v PageInfo)`

SetPage sets Page field to given value.

### HasPage

`func (o *CommitData) HasPage() bool`

HasPage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


