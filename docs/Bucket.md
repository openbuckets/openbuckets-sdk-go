# Bucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier for the bucket | [optional] 
**Name** | Pointer to **string** | Name of the bucket | [optional] 
**FileCount** | Pointer to **int32** | Number of files in the bucket | [optional] 
**Type** | Pointer to **string** | Type of cloud storage (aws, azure, etc) | [optional] 

## Methods

### NewBucket

`func NewBucket() *Bucket`

NewBucket instantiates a new Bucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBucketWithDefaults

`func NewBucketWithDefaults() *Bucket`

NewBucketWithDefaults instantiates a new Bucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bucket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bucket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bucket) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Bucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Bucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Bucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFileCount

`func (o *Bucket) GetFileCount() int32`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *Bucket) GetFileCountOk() (*int32, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *Bucket) SetFileCount(v int32)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *Bucket) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetType

`func (o *Bucket) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Bucket) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Bucket) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Bucket) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


