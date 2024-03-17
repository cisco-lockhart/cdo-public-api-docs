# cdo_python_sdk.TransactionsApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_transaction**](TransactionsApi.md#get_transaction) | **GET** /v1/transactions/{transactionUid} | Get information of an in-progress CDO transaction


# **get_transaction**
> CdoTransaction get_transaction(transaction_uid)

Get information of an in-progress CDO transaction

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_python_sdk
from cdo_python_sdk.models.cdo_transaction import CdoTransaction
from cdo_python_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_python_sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_python_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_python_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_python_sdk.TransactionsApi(api_client)
    transaction_uid = 'transaction_uid_example' # str | 

    try:
        # Get information of an in-progress CDO transaction
        api_response = api_instance.get_transaction(transaction_uid)
        print("The response of TransactionsApi->get_transaction:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TransactionsApi->get_transaction: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transaction_uid** | **str**|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | CDO Transaction that has completed, either successfully or unsuccessfully. Note: failed CDO Transactions do not roll back. |  -  |
**202** | CDO Transaction that is in progress |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

