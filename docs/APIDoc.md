# APIDoc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APICount** | Pointer to **int32** | 接口数量 | [optional] 
**Code** | Pointer to **int32** | API 文档编号 | [optional] 
**CreatedAt** | Pointer to **int32** | 创建时间 | [optional] 
**CreatedBy** | Pointer to **int32** | 创建者 | [optional] 
**Description** | Pointer to **string** | API 文档描述 | [optional] [default to ""]
**Introduction** | Pointer to **string** | API 文档简介 | [optional] [default to ""]
**LatestReleaseCode** | Pointer to **int32** | 上次发布编号，编号为 0 代表没有发布过 | [optional] 
**LatestReleaseStatus** | Pointer to **string** | 上次发布状态 | [optional] [default to ""]
**LatestSourceType** | Pointer to **string** | 上次发布的文档类型 | [optional] [default to ""]
**LatestUpdatedAt** | Pointer to **NullableInt32** | 上次发布时间 | [optional] 
**PrefixUri** | Pointer to **string** | API 文档路由前缀 | [optional] [default to ""]
**PrivateAccess** | Pointer to **bool** | API 文档是否私有访问 | [optional] [default to false]
**ReleaseCount** | Pointer to **int32** | 发布次数 | [optional] 
**SharePassword** | Pointer to **NullableString** | 分享密码 | [optional] [default to ""]
**Title** | Pointer to **string** | API 文档标题 | [optional] [default to ""]
**UpdatedAt** | Pointer to **int32** | 更新时间 | [optional] 
**Uri** | Pointer to **string** | API 文档路由 | [optional] [default to ""]
**UseFileDes** | Pointer to **bool** | API 文档是否使用文件中的描述 | [optional] [default to false]
**UseFileSetting** | Pointer to **bool** | API 文档是否使用文件中的配置 | [optional] [default to false]
**ViewCount** | Pointer to **int32** | 查看次数 | [optional] 

## Methods

### NewAPIDoc

`func NewAPIDoc() *APIDoc`

NewAPIDoc instantiates a new APIDoc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIDocWithDefaults

`func NewAPIDocWithDefaults() *APIDoc`

NewAPIDocWithDefaults instantiates a new APIDoc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPICount

`func (o *APIDoc) GetAPICount() int32`

GetAPICount returns the APICount field if non-nil, zero value otherwise.

### GetAPICountOk

`func (o *APIDoc) GetAPICountOk() (*int32, bool)`

GetAPICountOk returns a tuple with the APICount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPICount

`func (o *APIDoc) SetAPICount(v int32)`

SetAPICount sets APICount field to given value.

### HasAPICount

`func (o *APIDoc) HasAPICount() bool`

HasAPICount returns a boolean if a field has been set.

### GetCode

`func (o *APIDoc) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIDoc) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIDoc) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIDoc) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *APIDoc) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *APIDoc) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *APIDoc) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *APIDoc) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *APIDoc) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *APIDoc) GetCreatedByOk() (*int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *APIDoc) SetCreatedBy(v int32)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *APIDoc) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *APIDoc) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *APIDoc) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *APIDoc) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *APIDoc) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIntroduction

`func (o *APIDoc) GetIntroduction() string`

GetIntroduction returns the Introduction field if non-nil, zero value otherwise.

### GetIntroductionOk

`func (o *APIDoc) GetIntroductionOk() (*string, bool)`

GetIntroductionOk returns a tuple with the Introduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntroduction

`func (o *APIDoc) SetIntroduction(v string)`

SetIntroduction sets Introduction field to given value.

### HasIntroduction

`func (o *APIDoc) HasIntroduction() bool`

HasIntroduction returns a boolean if a field has been set.

### GetLatestReleaseCode

`func (o *APIDoc) GetLatestReleaseCode() int32`

GetLatestReleaseCode returns the LatestReleaseCode field if non-nil, zero value otherwise.

### GetLatestReleaseCodeOk

`func (o *APIDoc) GetLatestReleaseCodeOk() (*int32, bool)`

GetLatestReleaseCodeOk returns a tuple with the LatestReleaseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleaseCode

`func (o *APIDoc) SetLatestReleaseCode(v int32)`

SetLatestReleaseCode sets LatestReleaseCode field to given value.

### HasLatestReleaseCode

`func (o *APIDoc) HasLatestReleaseCode() bool`

HasLatestReleaseCode returns a boolean if a field has been set.

### GetLatestReleaseStatus

`func (o *APIDoc) GetLatestReleaseStatus() string`

GetLatestReleaseStatus returns the LatestReleaseStatus field if non-nil, zero value otherwise.

### GetLatestReleaseStatusOk

`func (o *APIDoc) GetLatestReleaseStatusOk() (*string, bool)`

GetLatestReleaseStatusOk returns a tuple with the LatestReleaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReleaseStatus

`func (o *APIDoc) SetLatestReleaseStatus(v string)`

SetLatestReleaseStatus sets LatestReleaseStatus field to given value.

### HasLatestReleaseStatus

`func (o *APIDoc) HasLatestReleaseStatus() bool`

HasLatestReleaseStatus returns a boolean if a field has been set.

### GetLatestSourceType

`func (o *APIDoc) GetLatestSourceType() string`

GetLatestSourceType returns the LatestSourceType field if non-nil, zero value otherwise.

### GetLatestSourceTypeOk

`func (o *APIDoc) GetLatestSourceTypeOk() (*string, bool)`

GetLatestSourceTypeOk returns a tuple with the LatestSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestSourceType

`func (o *APIDoc) SetLatestSourceType(v string)`

SetLatestSourceType sets LatestSourceType field to given value.

### HasLatestSourceType

`func (o *APIDoc) HasLatestSourceType() bool`

HasLatestSourceType returns a boolean if a field has been set.

### GetLatestUpdatedAt

`func (o *APIDoc) GetLatestUpdatedAt() int32`

GetLatestUpdatedAt returns the LatestUpdatedAt field if non-nil, zero value otherwise.

### GetLatestUpdatedAtOk

`func (o *APIDoc) GetLatestUpdatedAtOk() (*int32, bool)`

GetLatestUpdatedAtOk returns a tuple with the LatestUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestUpdatedAt

`func (o *APIDoc) SetLatestUpdatedAt(v int32)`

SetLatestUpdatedAt sets LatestUpdatedAt field to given value.

### HasLatestUpdatedAt

`func (o *APIDoc) HasLatestUpdatedAt() bool`

HasLatestUpdatedAt returns a boolean if a field has been set.

### SetLatestUpdatedAtNil

`func (o *APIDoc) SetLatestUpdatedAtNil(b bool)`

 SetLatestUpdatedAtNil sets the value for LatestUpdatedAt to be an explicit nil

### UnsetLatestUpdatedAt
`func (o *APIDoc) UnsetLatestUpdatedAt()`

UnsetLatestUpdatedAt ensures that no value is present for LatestUpdatedAt, not even an explicit nil
### GetPrefixUri

`func (o *APIDoc) GetPrefixUri() string`

GetPrefixUri returns the PrefixUri field if non-nil, zero value otherwise.

### GetPrefixUriOk

`func (o *APIDoc) GetPrefixUriOk() (*string, bool)`

GetPrefixUriOk returns a tuple with the PrefixUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixUri

`func (o *APIDoc) SetPrefixUri(v string)`

SetPrefixUri sets PrefixUri field to given value.

### HasPrefixUri

`func (o *APIDoc) HasPrefixUri() bool`

HasPrefixUri returns a boolean if a field has been set.

### GetPrivateAccess

`func (o *APIDoc) GetPrivateAccess() bool`

GetPrivateAccess returns the PrivateAccess field if non-nil, zero value otherwise.

### GetPrivateAccessOk

`func (o *APIDoc) GetPrivateAccessOk() (*bool, bool)`

GetPrivateAccessOk returns a tuple with the PrivateAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAccess

`func (o *APIDoc) SetPrivateAccess(v bool)`

SetPrivateAccess sets PrivateAccess field to given value.

### HasPrivateAccess

`func (o *APIDoc) HasPrivateAccess() bool`

HasPrivateAccess returns a boolean if a field has been set.

### GetReleaseCount

`func (o *APIDoc) GetReleaseCount() int32`

GetReleaseCount returns the ReleaseCount field if non-nil, zero value otherwise.

### GetReleaseCountOk

`func (o *APIDoc) GetReleaseCountOk() (*int32, bool)`

GetReleaseCountOk returns a tuple with the ReleaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseCount

`func (o *APIDoc) SetReleaseCount(v int32)`

SetReleaseCount sets ReleaseCount field to given value.

### HasReleaseCount

`func (o *APIDoc) HasReleaseCount() bool`

HasReleaseCount returns a boolean if a field has been set.

### GetSharePassword

`func (o *APIDoc) GetSharePassword() string`

GetSharePassword returns the SharePassword field if non-nil, zero value otherwise.

### GetSharePasswordOk

`func (o *APIDoc) GetSharePasswordOk() (*string, bool)`

GetSharePasswordOk returns a tuple with the SharePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePassword

`func (o *APIDoc) SetSharePassword(v string)`

SetSharePassword sets SharePassword field to given value.

### HasSharePassword

`func (o *APIDoc) HasSharePassword() bool`

HasSharePassword returns a boolean if a field has been set.

### SetSharePasswordNil

`func (o *APIDoc) SetSharePasswordNil(b bool)`

 SetSharePasswordNil sets the value for SharePassword to be an explicit nil

### UnsetSharePassword
`func (o *APIDoc) UnsetSharePassword()`

UnsetSharePassword ensures that no value is present for SharePassword, not even an explicit nil
### GetTitle

`func (o *APIDoc) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *APIDoc) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *APIDoc) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *APIDoc) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *APIDoc) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *APIDoc) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *APIDoc) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *APIDoc) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUri

`func (o *APIDoc) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *APIDoc) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *APIDoc) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *APIDoc) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetUseFileDes

`func (o *APIDoc) GetUseFileDes() bool`

GetUseFileDes returns the UseFileDes field if non-nil, zero value otherwise.

### GetUseFileDesOk

`func (o *APIDoc) GetUseFileDesOk() (*bool, bool)`

GetUseFileDesOk returns a tuple with the UseFileDes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFileDes

`func (o *APIDoc) SetUseFileDes(v bool)`

SetUseFileDes sets UseFileDes field to given value.

### HasUseFileDes

`func (o *APIDoc) HasUseFileDes() bool`

HasUseFileDes returns a boolean if a field has been set.

### GetUseFileSetting

`func (o *APIDoc) GetUseFileSetting() bool`

GetUseFileSetting returns the UseFileSetting field if non-nil, zero value otherwise.

### GetUseFileSettingOk

`func (o *APIDoc) GetUseFileSettingOk() (*bool, bool)`

GetUseFileSettingOk returns a tuple with the UseFileSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFileSetting

`func (o *APIDoc) SetUseFileSetting(v bool)`

SetUseFileSetting sets UseFileSetting field to given value.

### HasUseFileSetting

`func (o *APIDoc) HasUseFileSetting() bool`

HasUseFileSetting returns a boolean if a field has been set.

### GetViewCount

`func (o *APIDoc) GetViewCount() int32`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *APIDoc) GetViewCountOk() (*int32, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *APIDoc) SetViewCount(v int32)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *APIDoc) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


