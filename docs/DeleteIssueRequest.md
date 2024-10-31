# DeleteIssueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueCode** | **int64** | 事项 Code | 
**IssueType** | Pointer to **string** | 事项类型，可不填 | [optional] 
**ProjectName** | **string** | 项目名称 | 

## Methods

### NewDeleteIssueRequest

`func NewDeleteIssueRequest(issueCode int64, projectName string, ) *DeleteIssueRequest`

NewDeleteIssueRequest instantiates a new DeleteIssueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteIssueRequestWithDefaults

`func NewDeleteIssueRequestWithDefaults() *DeleteIssueRequest`

NewDeleteIssueRequestWithDefaults instantiates a new DeleteIssueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueCode

`func (o *DeleteIssueRequest) GetIssueCode() int64`

GetIssueCode returns the IssueCode field if non-nil, zero value otherwise.

### GetIssueCodeOk

`func (o *DeleteIssueRequest) GetIssueCodeOk() (*int64, bool)`

GetIssueCodeOk returns a tuple with the IssueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueCode

`func (o *DeleteIssueRequest) SetIssueCode(v int64)`

SetIssueCode sets IssueCode field to given value.


### GetIssueType

`func (o *DeleteIssueRequest) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *DeleteIssueRequest) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *DeleteIssueRequest) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *DeleteIssueRequest) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetProjectName

`func (o *DeleteIssueRequest) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *DeleteIssueRequest) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *DeleteIssueRequest) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


