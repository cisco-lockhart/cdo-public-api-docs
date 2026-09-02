# scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi

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
[**generate_policy_analysis_insights_report**](PolicyAnalysisAndOptimizationApi.md#generate_policy_analysis_insights_report) | **POST** /v1/agenticops/policies/nat/report-gen/analysis | Generate NAT policy analysis report
[**get_latest_policy_analysis_data**](PolicyAnalysisAndOptimizationApi.md#get_latest_policy_analysis_data) | **GET** /v1/agenticops/policies/nat/analyses/latest | Get latest NAT policy analyses
[**get_nat_policy_analysis_summary**](PolicyAnalysisAndOptimizationApi.md#get_nat_policy_analysis_summary) | **GET** /v1/agenticops/policies/nat/analysis/insights/summary/{insightsUid} | Get NAT policy analysis summary
[**get_nat_policy_analysis_summary_by_data_source_uid**](PolicyAnalysisAndOptimizationApi.md#get_nat_policy_analysis_summary_by_data_source_uid) | **GET** /v1/agenticops/policies/nat/analyses/latest/summary | Get latest NAT policy analysis summary
[**get_nat_policy_rule_conflicts**](PolicyAnalysisAndOptimizationApi.md#get_nat_policy_rule_conflicts) | **GET** /v1/agenticops/policies/nat/analysis/insights/{insightsUid}/details | Get NAT policy analysis rule conflicts
[**get_policy_analysis_insights_report_url**](PolicyAnalysisAndOptimizationApi.md#get_policy_analysis_insights_report_url) | **GET** /v1/agenticops/policies/nat/report-gen/analysis/{insightsUid} | Get NAT policy analysis report URL
[**sync_nat_policies**](PolicyAnalysisAndOptimizationApi.md#sync_nat_policies) | **POST** /v1/agenticops/policies/nat/sync | Sync FMC NAT policies
[**trigger_nat_policy_analysis**](PolicyAnalysisAndOptimizationApi.md#trigger_nat_policy_analysis) | **POST** /v1/agenticops/policies/nat/analyses | Start NAT policy analysis


# **generate_policy_analysis_insights_report**
> CdoTransaction generate_policy_analysis_insights_report(insights_report_generation_request)

Generate NAT policy analysis report

Generates or reuses a PDF report for the requested NAT policy analysis. The operation returns a Security Cloud Control Transaction object that can be used to track report generation progress. When the transaction completes, `transactionDetails` contains `reportUrl`. The request requires the analysis and data source identifiers; an S3 key is not required.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.insights_report_generation_request import InsightsReportGenerationRequest
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    insights_report_generation_request = scc_firewall_manager_sdk.InsightsReportGenerationRequest() # InsightsReportGenerationRequest | 

    try:
        # Generate NAT policy analysis report
        api_response = api_instance.generate_policy_analysis_insights_report(insights_report_generation_request)
        print("The response of PolicyAnalysisAndOptimizationApi->generate_policy_analysis_insights_report:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->generate_policy_analysis_insights_report: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insights_report_generation_request** | [**InsightsReportGenerationRequest**](InsightsReportGenerationRequest.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track report generation progress |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_latest_policy_analysis_data**
> PolicyNodeDataPage get_latest_policy_analysis_data(limit=limit, offset=offset, q=q, sort=sort)

Get latest NAT policy analyses

Gets the latest NAT policy analyses across data sources. Results can be filtered, sorted, and paginated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.policy_node_data_page import PolicyNodeDataPage
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['analysisDetails.analysisStartTime:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get latest NAT policy analyses
        api_response = api_instance.get_latest_policy_analysis_data(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of PolicyAnalysisAndOptimizationApi->get_latest_policy_analysis_data:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->get_latest_policy_analysis_data: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**PolicyNodeDataPage**](PolicyNodeDataPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Paginated list of the latest NAT policy analyses |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_nat_policy_analysis_summary**
> PolicyAnalysisInsightsSummary get_nat_policy_analysis_summary(insights_uid)

Get NAT policy analysis summary

Gets the summary for the specified NAT policy analysis.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.policy_analysis_insights_summary import PolicyAnalysisInsightsSummary
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    insights_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.

    try:
        # Get NAT policy analysis summary
        api_response = api_instance.get_nat_policy_analysis_summary(insights_uid)
        print("The response of PolicyAnalysisAndOptimizationApi->get_nat_policy_analysis_summary:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->get_nat_policy_analysis_summary: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insights_uid** | **str**| The unique identifier of the NAT policy analysis. | 

### Return type

[**PolicyAnalysisInsightsSummary**](PolicyAnalysisInsightsSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | NAT policy analysis summary |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_nat_policy_analysis_summary_by_data_source_uid**
> PolicyAnalysisInsightsSummary get_nat_policy_analysis_summary_by_data_source_uid(q=q)

Get latest NAT policy analysis summary

Gets an aggregated summary of the latest NAT policy analyses for the specified data source.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.policy_analysis_insights_summary import PolicyAnalysisInsightsSummary
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get latest NAT policy analysis summary
        api_response = api_instance.get_nat_policy_analysis_summary_by_data_source_uid(q=q)
        print("The response of PolicyAnalysisAndOptimizationApi->get_nat_policy_analysis_summary_by_data_source_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->get_nat_policy_analysis_summary_by_data_source_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**PolicyAnalysisInsightsSummary**](PolicyAnalysisInsightsSummary.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Aggregated summary of the latest NAT policy analyses for the data source |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_nat_policy_rule_conflicts**
> RuleConflictsDataPage get_nat_policy_rule_conflicts(insights_uid, limit=limit, offset=offset, shadowed_only=shadowed_only, redundant_only=redundant_only)

Get NAT policy analysis rule conflicts

Gets rule-level conflict data for the specified NAT policy analysis. Each item includes the `hasShadowedRules` and `hasRedundantRules` fields, which indicate whether the conflicting rules are shadowed or redundant, respectively. If there are no conflicts, an empty page is returned.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.rule_conflicts_data_page import RuleConflictsDataPage
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    insights_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    shadowed_only = False # bool | Filter results to conflicts containing shadowed rules. (optional) (default to False)
    redundant_only = False # bool | Filter results to conflicts containing redundant rules. (optional) (default to False)

    try:
        # Get NAT policy analysis rule conflicts
        api_response = api_instance.get_nat_policy_rule_conflicts(insights_uid, limit=limit, offset=offset, shadowed_only=shadowed_only, redundant_only=redundant_only)
        print("The response of PolicyAnalysisAndOptimizationApi->get_nat_policy_rule_conflicts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->get_nat_policy_rule_conflicts: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insights_uid** | **str**| The unique identifier of the NAT policy analysis. | 
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **shadowed_only** | **bool**| Filter results to conflicts containing shadowed rules. | [optional] [default to False]
 **redundant_only** | **bool**| Filter results to conflicts containing redundant rules. | [optional] [default to False]

### Return type

[**RuleConflictsDataPage**](RuleConflictsDataPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Paginated list of NAT policy analysis rule conflicts |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_policy_analysis_insights_report_url**
> str get_policy_analysis_insights_report_url(insights_uid)

Get NAT policy analysis report URL

Gets a `reportUrl` field containing a presigned download URL that is valid for 30 minutes. A current PDF report must already exist for the requested analysis.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    insights_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis whose report URL is requested.

    try:
        # Get NAT policy analysis report URL
        api_response = api_instance.get_policy_analysis_insights_report_url(insights_uid)
        print("The response of PolicyAnalysisAndOptimizationApi->get_policy_analysis_insights_report_url:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->get_policy_analysis_insights_report_url: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insights_uid** | **str**| The unique identifier of the NAT policy analysis whose report URL is requested. | 

### Return type

**str**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Object containing a presigned download URL for the NAT policy analysis report |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sync_nat_policies**
> CdoTransactionListResponse sync_nat_policies(nat_policy_sync_request=nat_policy_sync_request)

Sync FMC NAT policies

Creates one NAT policy sync transaction for each requested data source. If the <code>dataSourceUids</code> field is omitted or empty, all eligible FMC data sources are synced. A data source is the device manager that owns the NAT policies on this tenant, so the data source UIDs are the unique identifiers of the FMC device managers onboarded to it. Use the <a href=\"https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-managers/\">Device Managers</a> endpoint to list them and pick the UIDs to sync.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction_list_response import CdoTransactionListResponse
from scc_firewall_manager_sdk.models.nat_policy_sync_request import NatPolicySyncRequest
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    nat_policy_sync_request = scc_firewall_manager_sdk.NatPolicySyncRequest() # NatPolicySyncRequest |  (optional)

    try:
        # Sync FMC NAT policies
        api_response = api_instance.sync_nat_policies(nat_policy_sync_request=nat_policy_sync_request)
        print("The response of PolicyAnalysisAndOptimizationApi->sync_nat_policies:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->sync_nat_policies: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nat_policy_sync_request** | [**NatPolicySyncRequest**](NatPolicySyncRequest.md)|  | [optional] 

### Return type

[**CdoTransactionListResponse**](CdoTransactionListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction objects that can be used to track NAT policy sync progress |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **trigger_nat_policy_analysis**
> CdoTransaction trigger_nat_policy_analysis(nat_policy_analysis_request)

Start NAT policy analysis

Starts an asynchronous analysis of one or more NAT policies for a single data source. A data source is the device manager that owns the NAT policies on this tenant, so each <code>dataSourceUid</code> is the unique identifier, represented as a UUID, of an FMC device manager onboarded to it. Use the <a href=\"https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-managers/\">Device Managers</a> endpoint to list them and pick the data source(s) to analyze. The operation returns a Security Cloud Control Transaction object that can be used to track progress.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.nat_policy_analysis_request import NatPolicyAnalysisRequest
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
    api_instance = scc_firewall_manager_sdk.PolicyAnalysisAndOptimizationApi(api_client)
    nat_policy_analysis_request = scc_firewall_manager_sdk.NatPolicyAnalysisRequest() # NatPolicyAnalysisRequest | 

    try:
        # Start NAT policy analysis
        api_response = api_instance.trigger_nat_policy_analysis(nat_policy_analysis_request)
        print("The response of PolicyAnalysisAndOptimizationApi->trigger_nat_policy_analysis:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling PolicyAnalysisAndOptimizationApi->trigger_nat_policy_analysis: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nat_policy_analysis_request** | [**NatPolicyAnalysisRequest**](NatPolicyAnalysisRequest.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track policy analysis progress |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**409** | No requested policy requires analysis or all are unavailable |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

