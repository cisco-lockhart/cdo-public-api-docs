# scc_firewall_manager_sdk.NATRemediationReportAPIsApi

All URIs are relative to the base URL, which depends on the region your organization is deployed to.

Region | Base URL
------------- | -------------
US  | https://api.us.security.cisco.com/firewall
EU  | https://api.eu.security.cisco.com/firewall
APJ | https://api.apj.security.cisco.com/firewall
AU  | https://api.au.security.cisco.com/firewall
IN  | https://api.in.security.cisco.com/firewall
UAE | https://api.uae.security.cisco.com/firewall

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_latest_remediation_report**](NATRemediationReportAPIsApi.md#get_latest_remediation_report) | **GET** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-reports/latest | Get latest NAT policy remediation report
[**get_remediation_report_by_version**](NATRemediationReportAPIsApi.md#get_remediation_report_by_version) | **GET** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-reports/{version} | Get NAT policy remediation report by version
[**get_remediation_reports**](NATRemediationReportAPIsApi.md#get_remediation_reports) | **GET** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-reports | Get NAT policy remediation reports


# **get_latest_remediation_report**
> RemediationReportDto get_latest_remediation_report(insight_uid)

Get latest NAT policy remediation report

Gets the latest NAT policy remediation report for the specified NAT policy analysis. Only finished remediations are included, whether they succeeded or failed.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.remediation_report_dto import RemediationReportDto
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.NATRemediationReportAPIsApi(api_client)
    insight_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.

    try:
        # Get latest NAT policy remediation report
        api_response = api_instance.get_latest_remediation_report(insight_uid)
        print("The response of NATRemediationReportAPIsApi->get_latest_remediation_report:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationReportAPIsApi->get_latest_remediation_report: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**| The unique identifier of the NAT policy analysis. | 

### Return type

[**RemediationReportDto**](RemediationReportDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Latest NAT policy remediation report |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_remediation_report_by_version**
> RemediationReportDto get_remediation_report_by_version(insight_uid, version)

Get NAT policy remediation report by version

Gets the specified version of a NAT policy remediation report for the NAT policy analysis. The remediation must have finished, whether it succeeded or failed.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.remediation_report_dto import RemediationReportDto
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.NATRemediationReportAPIsApi(api_client)
    insight_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.
    version = 56 # int | The version of the NAT policy remediation report to retrieve.

    try:
        # Get NAT policy remediation report by version
        api_response = api_instance.get_remediation_report_by_version(insight_uid, version)
        print("The response of NATRemediationReportAPIsApi->get_remediation_report_by_version:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationReportAPIsApi->get_remediation_report_by_version: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**| The unique identifier of the NAT policy analysis. | 
 **version** | **int**| The version of the NAT policy remediation report to retrieve. | 

### Return type

[**RemediationReportDto**](RemediationReportDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | NAT policy remediation report |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_remediation_reports**
> RemediationReportPage get_remediation_reports(insight_uid, limit=limit, offset=offset)

Get NAT policy remediation reports

Gets NAT policy remediation reports for the specified NAT policy analysis. Only finished remediations are included, whether they succeeded or failed. Results are ordered by version from newest to oldest and can be paginated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.remediation_report_page import RemediationReportPage
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.NATRemediationReportAPIsApi(api_client)
    insight_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.
    limit = 10 # int | The maximum number of NAT policy remediation reports to retrieve. (optional) (default to 10)
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')

    try:
        # Get NAT policy remediation reports
        api_response = api_instance.get_remediation_reports(insight_uid, limit=limit, offset=offset)
        print("The response of NATRemediationReportAPIsApi->get_remediation_reports:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationReportAPIsApi->get_remediation_reports: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**| The unique identifier of the NAT policy analysis. | 
 **limit** | **int**| The maximum number of NAT policy remediation reports to retrieve. | [optional] [default to 10]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]

### Return type

[**RemediationReportPage**](RemediationReportPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Paginated list of NAT policy remediation reports |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

