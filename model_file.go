/*
OpenBuckets API

The OpenBuckets web-based tool is a powerful utility that allows users to quickly locate open buckets in cloud storage systems through a simple query. In addition, it provides a convenient way to search for various file types across these open buckets, making it an essential tool for security professionals, researchers, and anyone interested in discovering exposed data. This Postman collection aims to showcase the capabilities of OpenBuckets by providing a set of API requests that demonstrate how to leverage its features. By following this collection, you'll learn how to utilize OpenBuckets to identify open buckets and search for specific file types within them.

API version: 1.0.0
Contact: support@openbuckets.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openbuckets

import (
	"encoding/json"
)

// checks if the File type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &File{}

// File struct for File
type File struct {
	// Unique identifier for the file
	Id *string `json:"id,omitempty"`
	// Name of the bucket
	Bucket *string `json:"bucket,omitempty"`
	// Unique identifier for the bucket
	BucketId *int32 `json:"bucketId,omitempty"`
	// Name of the file
	Filename *string `json:"filename,omitempty"`
	// Full path to the file
	FullPath *string `json:"fullPath,omitempty"`
	// URL to access the file
	Url *string `json:"url,omitempty"`
	// Size of the file in bytes
	Size *int32 `json:"size,omitempty"`
	// Type of cloud storage
	Type *string `json:"type,omitempty"`
	// Timestamp when file was last modified
	LastModified *int32 `json:"lastModified,omitempty"`
	// Storage container name
	Container *string `json:"container,omitempty"`
}

// NewFile instantiates a new File object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFile() *File {
	this := File{}
	return &this
}

// NewFileWithDefaults instantiates a new File object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileWithDefaults() *File {
	this := File{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *File) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *File) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *File) SetId(v string) {
	o.Id = &v
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *File) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *File) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *File) SetBucket(v string) {
	o.Bucket = &v
}

// GetBucketId returns the BucketId field value if set, zero value otherwise.
func (o *File) GetBucketId() int32 {
	if o == nil || IsNil(o.BucketId) {
		var ret int32
		return ret
	}
	return *o.BucketId
}

// GetBucketIdOk returns a tuple with the BucketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetBucketIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BucketId) {
		return nil, false
	}
	return o.BucketId, true
}

// HasBucketId returns a boolean if a field has been set.
func (o *File) HasBucketId() bool {
	if o != nil && !IsNil(o.BucketId) {
		return true
	}

	return false
}

// SetBucketId gets a reference to the given int32 and assigns it to the BucketId field.
func (o *File) SetBucketId(v int32) {
	o.BucketId = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *File) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *File) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *File) SetFilename(v string) {
	o.Filename = &v
}

// GetFullPath returns the FullPath field value if set, zero value otherwise.
func (o *File) GetFullPath() string {
	if o == nil || IsNil(o.FullPath) {
		var ret string
		return ret
	}
	return *o.FullPath
}

// GetFullPathOk returns a tuple with the FullPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetFullPathOk() (*string, bool) {
	if o == nil || IsNil(o.FullPath) {
		return nil, false
	}
	return o.FullPath, true
}

// HasFullPath returns a boolean if a field has been set.
func (o *File) HasFullPath() bool {
	if o != nil && !IsNil(o.FullPath) {
		return true
	}

	return false
}

// SetFullPath gets a reference to the given string and assigns it to the FullPath field.
func (o *File) SetFullPath(v string) {
	o.FullPath = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *File) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *File) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *File) SetUrl(v string) {
	o.Url = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *File) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *File) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *File) SetSize(v int32) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *File) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *File) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *File) SetType(v string) {
	o.Type = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *File) GetLastModified() int32 {
	if o == nil || IsNil(o.LastModified) {
		var ret int32
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetLastModifiedOk() (*int32, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *File) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given int32 and assigns it to the LastModified field.
func (o *File) SetLastModified(v int32) {
	o.LastModified = &v
}

// GetContainer returns the Container field value if set, zero value otherwise.
func (o *File) GetContainer() string {
	if o == nil || IsNil(o.Container) {
		var ret string
		return ret
	}
	return *o.Container
}

// GetContainerOk returns a tuple with the Container field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *File) GetContainerOk() (*string, bool) {
	if o == nil || IsNil(o.Container) {
		return nil, false
	}
	return o.Container, true
}

// HasContainer returns a boolean if a field has been set.
func (o *File) HasContainer() bool {
	if o != nil && !IsNil(o.Container) {
		return true
	}

	return false
}

// SetContainer gets a reference to the given string and assigns it to the Container field.
func (o *File) SetContainer(v string) {
	o.Container = &v
}

func (o File) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o File) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Bucket) {
		toSerialize["bucket"] = o.Bucket
	}
	if !IsNil(o.BucketId) {
		toSerialize["bucketId"] = o.BucketId
	}
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.FullPath) {
		toSerialize["fullPath"] = o.FullPath
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.Container) {
		toSerialize["container"] = o.Container
	}
	return toSerialize, nil
}

type NullableFile struct {
	value *File
	isSet bool
}

func (v NullableFile) Get() *File {
	return v.value
}

func (v *NullableFile) Set(val *File) {
	v.value = val
	v.isSet = true
}

func (v NullableFile) IsSet() bool {
	return v.isSet
}

func (v *NullableFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFile(val *File) *NullableFile {
	return &NullableFile{value: val, isSet: true}
}

func (v NullableFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

