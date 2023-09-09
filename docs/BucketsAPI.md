# \BucketsAPI

All URIs are relative to *https://api.openbuckets.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchBuckets**](BucketsAPI.md#SearchBuckets) | **Get** /api/v2/buckets | Search Buckets



## SearchBuckets

> BucketSearchResults SearchBuckets(ctx).Keywords(keywords).Type_(type_).Exact(exact).Start(start).Limit(limit).Order(order).Direction(direction).Execute()

Search Buckets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "https://github.com/openbuckets/sdk-go"
)

func main() {
    keywords := "abg" // string | the search keywords to filter bucket names (e.g., \"abg\") (optional)
    type_ := "aws" // string | the type of bucket to filter (e.g., aws,dos,azure,gcp) (optional)
    exact := float32(0) // float32 | whether to perform an exact match for the keywords (0 for false, 1 for true) (optional)
    start := float32(0) // float32 | starting index for pagination (optional)
    limit := float32(1000) // float32 | number of search results to return per page (optional)
    order := "fileCount" // string | the sorting field for the search results (e.g., \"fileCount\" for sorting by file count) (optional)
    direction := "asc" // string | the sorting direction for the search results (e.g., \"asc\" for ascending) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketsAPI.SearchBuckets(context.Background()).Keywords(keywords).Type_(type_).Exact(exact).Start(start).Limit(limit).Order(order).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketsAPI.SearchBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchBuckets`: BucketSearchResults
    fmt.Fprintf(os.Stdout, "Response from `BucketsAPI.SearchBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywords** | **string** | the search keywords to filter bucket names (e.g., \&quot;abg\&quot;) | 
 **type_** | **string** | the type of bucket to filter (e.g., aws,dos,azure,gcp) | 
 **exact** | **float32** | whether to perform an exact match for the keywords (0 for false, 1 for true) | 
 **start** | **float32** | starting index for pagination | 
 **limit** | **float32** | number of search results to return per page | 
 **order** | **string** | the sorting field for the search results (e.g., \&quot;fileCount\&quot; for sorting by file count) | 
 **direction** | **string** | the sorting direction for the search results (e.g., \&quot;asc\&quot; for ascending) | 

### Return type

[**BucketSearchResults**](BucketSearchResults.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

