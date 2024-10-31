# DepartmentDepartmentTreeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]DepartmentDepartmentTreeData**](DepartmentDepartmentTreeData.md) | 子部门信息 | 
**CreatorId** | **int64** | 创建人ID | [default to 0]
**DescribeId** | **string** | 部门描述ID | [default to ""]
**Id** | **int64** | 部门ID | [default to 0]
**Name** | **string** | 部门名 | [default to ""]
**ParentId** | **int64** | 父级部门ID | [default to 0]
**Sort** | **int64** | 排序值 | [default to 0]
**TeamId** | **int64** | 团队ID | [default to 0]

## Methods

### NewDepartmentDepartmentTreeData

`func NewDepartmentDepartmentTreeData(children []DepartmentDepartmentTreeData, creatorId int64, describeId string, id int64, name string, parentId int64, sort int64, teamId int64, ) *DepartmentDepartmentTreeData`

NewDepartmentDepartmentTreeData instantiates a new DepartmentDepartmentTreeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentDepartmentTreeDataWithDefaults

`func NewDepartmentDepartmentTreeDataWithDefaults() *DepartmentDepartmentTreeData`

NewDepartmentDepartmentTreeDataWithDefaults instantiates a new DepartmentDepartmentTreeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *DepartmentDepartmentTreeData) GetChildren() []DepartmentDepartmentTreeData`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *DepartmentDepartmentTreeData) GetChildrenOk() (*[]DepartmentDepartmentTreeData, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *DepartmentDepartmentTreeData) SetChildren(v []DepartmentDepartmentTreeData)`

SetChildren sets Children field to given value.


### GetCreatorId

`func (o *DepartmentDepartmentTreeData) GetCreatorId() int64`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *DepartmentDepartmentTreeData) GetCreatorIdOk() (*int64, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *DepartmentDepartmentTreeData) SetCreatorId(v int64)`

SetCreatorId sets CreatorId field to given value.


### GetDescribeId

`func (o *DepartmentDepartmentTreeData) GetDescribeId() string`

GetDescribeId returns the DescribeId field if non-nil, zero value otherwise.

### GetDescribeIdOk

`func (o *DepartmentDepartmentTreeData) GetDescribeIdOk() (*string, bool)`

GetDescribeIdOk returns a tuple with the DescribeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescribeId

`func (o *DepartmentDepartmentTreeData) SetDescribeId(v string)`

SetDescribeId sets DescribeId field to given value.


### GetId

`func (o *DepartmentDepartmentTreeData) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepartmentDepartmentTreeData) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepartmentDepartmentTreeData) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *DepartmentDepartmentTreeData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepartmentDepartmentTreeData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepartmentDepartmentTreeData) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *DepartmentDepartmentTreeData) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *DepartmentDepartmentTreeData) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *DepartmentDepartmentTreeData) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### GetSort

`func (o *DepartmentDepartmentTreeData) GetSort() int64`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *DepartmentDepartmentTreeData) GetSortOk() (*int64, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *DepartmentDepartmentTreeData) SetSort(v int64)`

SetSort sets Sort field to given value.


### GetTeamId

`func (o *DepartmentDepartmentTreeData) GetTeamId() int64`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DepartmentDepartmentTreeData) GetTeamIdOk() (*int64, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DepartmentDepartmentTreeData) SetTeamId(v int64)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


