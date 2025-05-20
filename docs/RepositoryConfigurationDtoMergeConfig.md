# RepositoryConfigurationDtoMergeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultStrategy** | Pointer to [**MergeStrategy**](MergeStrategy.md) |  | [optional] 
**Strategies** | Pointer to [**[]MergeStrategy**](MergeStrategy.md) |  | [optional] 

## Methods

### NewRepositoryConfigurationDtoMergeConfig

`func NewRepositoryConfigurationDtoMergeConfig() *RepositoryConfigurationDtoMergeConfig`

NewRepositoryConfigurationDtoMergeConfig instantiates a new RepositoryConfigurationDtoMergeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryConfigurationDtoMergeConfigWithDefaults

`func NewRepositoryConfigurationDtoMergeConfigWithDefaults() *RepositoryConfigurationDtoMergeConfig`

NewRepositoryConfigurationDtoMergeConfigWithDefaults instantiates a new RepositoryConfigurationDtoMergeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultStrategy

`func (o *RepositoryConfigurationDtoMergeConfig) GetDefaultStrategy() MergeStrategy`

GetDefaultStrategy returns the DefaultStrategy field if non-nil, zero value otherwise.

### GetDefaultStrategyOk

`func (o *RepositoryConfigurationDtoMergeConfig) GetDefaultStrategyOk() (*MergeStrategy, bool)`

GetDefaultStrategyOk returns a tuple with the DefaultStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultStrategy

`func (o *RepositoryConfigurationDtoMergeConfig) SetDefaultStrategy(v MergeStrategy)`

SetDefaultStrategy sets DefaultStrategy field to given value.

### HasDefaultStrategy

`func (o *RepositoryConfigurationDtoMergeConfig) HasDefaultStrategy() bool`

HasDefaultStrategy returns a boolean if a field has been set.

### GetStrategies

`func (o *RepositoryConfigurationDtoMergeConfig) GetStrategies() []MergeStrategy`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *RepositoryConfigurationDtoMergeConfig) GetStrategiesOk() (*[]MergeStrategy, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *RepositoryConfigurationDtoMergeConfig) SetStrategies(v []MergeStrategy)`

SetStrategies sets Strategies field to given value.

### HasStrategies

`func (o *RepositoryConfigurationDtoMergeConfig) HasStrategies() bool`

HasStrategies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


