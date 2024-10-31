# DescribeGitDepotDeployKeys200ResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to [**[]DeployKeyInfo**](DeployKeyInfo.md) | 部署公钥列表 | [optional] 
**RequestId** | Pointer to **string** | 请求id | [optional] [default to "xxxxx"]

## Methods

### NewDescribeGitDepotDeployKeys200ResponseResponse

`func NewDescribeGitDepotDeployKeys200ResponseResponse() *DescribeGitDepotDeployKeys200ResponseResponse`

NewDescribeGitDepotDeployKeys200ResponseResponse instantiates a new DescribeGitDepotDeployKeys200ResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDescribeGitDepotDeployKeys200ResponseResponseWithDefaults

`func NewDescribeGitDepotDeployKeys200ResponseResponseWithDefaults() *DescribeGitDepotDeployKeys200ResponseResponse`

NewDescribeGitDepotDeployKeys200ResponseResponseWithDefaults instantiates a new DescribeGitDepotDeployKeys200ResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) GetKeys() []DeployKeyInfo`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) GetKeysOk() (*[]DeployKeyInfo, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) SetKeys(v []DeployKeyInfo)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetRequestId

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *DescribeGitDepotDeployKeys200ResponseResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


