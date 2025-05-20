# NotificationPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to [**OwnerDto**](OwnerDto.md) |  | [optional] 
**Service** | Pointer to [**ServiceDto**](ServiceDto.md) |  | [optional] 
**Repository** | Pointer to [**RepositoryDto**](RepositoryDto.md) |  | [optional] 

## Methods

### NewNotificationPayload

`func NewNotificationPayload() *NotificationPayload`

NewNotificationPayload instantiates a new NotificationPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationPayloadWithDefaults

`func NewNotificationPayloadWithDefaults() *NotificationPayload`

NewNotificationPayloadWithDefaults instantiates a new NotificationPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *NotificationPayload) GetOwner() OwnerDto`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NotificationPayload) GetOwnerOk() (*OwnerDto, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NotificationPayload) SetOwner(v OwnerDto)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NotificationPayload) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetService

`func (o *NotificationPayload) GetService() ServiceDto`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *NotificationPayload) GetServiceOk() (*ServiceDto, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *NotificationPayload) SetService(v ServiceDto)`

SetService sets Service field to given value.

### HasService

`func (o *NotificationPayload) HasService() bool`

HasService returns a boolean if a field has been set.

### GetRepository

`func (o *NotificationPayload) GetRepository() RepositoryDto`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *NotificationPayload) GetRepositoryOk() (*RepositoryDto, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *NotificationPayload) SetRepository(v RepositoryDto)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *NotificationPayload) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


