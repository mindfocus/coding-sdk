# PolicyDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statement** | [**[]PolicyStatement**](PolicyStatement.md) | 授权句柄 | 

## Methods

### NewPolicyDocument

`func NewPolicyDocument(statement []PolicyStatement, ) *PolicyDocument`

NewPolicyDocument instantiates a new PolicyDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyDocumentWithDefaults

`func NewPolicyDocumentWithDefaults() *PolicyDocument`

NewPolicyDocumentWithDefaults instantiates a new PolicyDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatement

`func (o *PolicyDocument) GetStatement() []PolicyStatement`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *PolicyDocument) GetStatementOk() (*[]PolicyStatement, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *PolicyDocument) SetStatement(v []PolicyStatement)`

SetStatement sets Statement field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


