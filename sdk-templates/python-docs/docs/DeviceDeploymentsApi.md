# scc_firewall_manager_sdk.DeviceDeploymentsApi

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
[**deploy_changes_to_multiple_asa_devices**](DeviceDeploymentsApi.md#deploy_changes_to_multiple_asa_devices) | **POST** /v1/inventory/devices/asas/deploy | Deploy changes to multiple ASA devices
[**deploy_changes_to_multiple_ftd_devices**](DeviceDeploymentsApi.md#deploy_changes_to_multiple_ftd_devices) | **POST** /v1/inventory/devices/ftds/deploy | (cdFMC-managed FTDs only) Deploy changes to multiple FTD devices
[**get_changes_deployed_in_deployment_run**](DeviceDeploymentsApi.md#get_changes_deployed_in_deployment_run) | **GET** /v1/inventory/devices/deployments/runs/{deploymentRunUid}/devices/{deviceUid}/changes | Get Deployed Device Changes in Deployment Run
[**get_deployment_run_transcript**](DeviceDeploymentsApi.md#get_deployment_run_transcript) | **GET** /v1/inventory/devices/deployments/runs/{deploymentRunUid}/devices/{deviceUid}/transcript | (cdFMC-managed FTDs only) Get Deployment Run Transcript
[**get_device_deployment_run**](DeviceDeploymentsApi.md#get_device_deployment_run) | **GET** /v1/inventory/devices/deployments/runs/{deploymentRunUid} | Get Device Deployment Run
[**get_device_deployment_runs**](DeviceDeploymentsApi.md#get_device_deployment_runs) | **GET** /v1/inventory/devices/deployments/runs | Get Device Deployment Runs


# **deploy_changes_to_multiple_asa_devices**
> CdoTransaction deploy_changes_to_multiple_asa_devices(asa_multi_device_deployment_input)

Deploy changes to multiple ASA devices

This is an asynchronous operation to deploy changes made to the configurations of multiple ASA devices on Security Cloud Control. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.  Notes:  - This operation is for ASA devices only.  - This operation will only deploy changes to the device if there are pending changes to deploy.   

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_multi_device_deployment_input import AsaMultiDeviceDeploymentInput
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
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
    api_instance = scc_firewall_manager_sdk.DeviceDeploymentsApi(api_client)
    asa_multi_device_deployment_input = scc_firewall_manager_sdk.AsaMultiDeviceDeploymentInput() # AsaMultiDeviceDeploymentInput | 

    try:
        # Deploy changes to multiple ASA devices
        api_response = api_instance.deploy_changes_to_multiple_asa_devices(asa_multi_device_deployment_input)
        print("The response of DeviceDeploymentsApi->deploy_changes_to_multiple_asa_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceDeploymentsApi->deploy_changes_to_multiple_asa_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asa_multi_device_deployment_input** | [**AsaMultiDeviceDeploymentInput**](AsaMultiDeviceDeploymentInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the deployment operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deploy_changes_to_multiple_ftd_devices**
> CdoTransaction deploy_changes_to_multiple_ftd_devices(ftd_multi_device_deployment_input)

(cdFMC-managed FTDs only) Deploy changes to multiple FTD devices

This is an asynchronous operation to deploy changes made to the configurations of multiple cdFMC-managed FTD devices on Security Cloud Control. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.  Notes:  - This operation is only supported for cdFMC-managed FTD devices.  - This operation will only deploy changes to the device if there are pending changes to deploy.   

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ftd_multi_device_deployment_input import FtdMultiDeviceDeploymentInput
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
    api_instance = scc_firewall_manager_sdk.DeviceDeploymentsApi(api_client)
    ftd_multi_device_deployment_input = scc_firewall_manager_sdk.FtdMultiDeviceDeploymentInput() # FtdMultiDeviceDeploymentInput | 

    try:
        # (cdFMC-managed FTDs only) Deploy changes to multiple FTD devices
        api_response = api_instance.deploy_changes_to_multiple_ftd_devices(ftd_multi_device_deployment_input)
        print("The response of DeviceDeploymentsApi->deploy_changes_to_multiple_ftd_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceDeploymentsApi->deploy_changes_to_multiple_ftd_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ftd_multi_device_deployment_input** | [**FtdMultiDeviceDeploymentInput**](FtdMultiDeviceDeploymentInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_changes_deployed_in_deployment_run**
> List[GetChangesDeployedInDeploymentRun200ResponseInner] get_changes_deployed_in_deployment_run(deployment_run_uid, device_uid)

Get Deployed Device Changes in Deployment Run

Get the changes that were deployed to a device as part of a deployment run. Changes are represented differently for FTDs and ASAs. For FTDs, this endpoint returns one item per changed entity describing what changed. For ASAs, this endpoint returns the initial configuration, and the configuration that was deployed to the device. The `deploymentType` field on each item identifies which shape it is. If the deployment failed, the changes returned are the changes that were attempted.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.get_changes_deployed_in_deployment_run200_response_inner import GetChangesDeployedInDeploymentRun200ResponseInner
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
    api_instance = scc_firewall_manager_sdk.DeviceDeploymentsApi(api_client)
    deployment_run_uid = 'deployment_run_uid_example' # str | The unique identifier, represented as a UUID, of the device deployment Run in SCC Firewall Manager.
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.

    try:
        # Get Deployed Device Changes in Deployment Run
        api_response = api_instance.get_changes_deployed_in_deployment_run(deployment_run_uid, device_uid)
        print("The response of DeviceDeploymentsApi->get_changes_deployed_in_deployment_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceDeploymentsApi->get_changes_deployed_in_deployment_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_run_uid** | **str**| The unique identifier, represented as a UUID, of the device deployment Run in SCC Firewall Manager. | 
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 

### Return type

[**List[GetChangesDeployedInDeploymentRun200ResponseInner]**](GetChangesDeployedInDeploymentRun200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The change items captured for the device in this deployment run. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_deployment_run_transcript**
> DeploymentRunTranscriptDto get_deployment_run_transcript(deployment_run_uid, device_uid)

(cdFMC-managed FTDs only) Get Deployment Run Transcript

Get the deployment run transcript for a device that was deployed to as part of a deployment run. Transcripts are only available for cdFMC-managed FTDs.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.deployment_run_transcript_dto import DeploymentRunTranscriptDto
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
    api_instance = scc_firewall_manager_sdk.DeviceDeploymentsApi(api_client)
    deployment_run_uid = 'deployment_run_uid_example' # str | The unique identifier, represented as a UUID, of the device deployment Run in SCC Firewall Manager.
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the cdFMC-managed FTD device in Security Cloud Control.

    try:
        # (cdFMC-managed FTDs only) Get Deployment Run Transcript
        api_response = api_instance.get_deployment_run_transcript(deployment_run_uid, device_uid)
        print("The response of DeviceDeploymentsApi->get_deployment_run_transcript:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceDeploymentsApi->get_deployment_run_transcript: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_run_uid** | **str**| The unique identifier, represented as a UUID, of the device deployment Run in SCC Firewall Manager. | 
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the cdFMC-managed FTD device in Security Cloud Control. | 

### Return type

[**DeploymentRunTranscriptDto**](DeploymentRunTranscriptDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The deployment run transcript captured for the device. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_deployment_run**
> DeploymentRunDto get_device_deployment_run(deployment_run_uid)

Get Device Deployment Run

Get a Device Deployment Run by UID in the SCC Firewall Manager Tenant. Each deployment run represents a group of devices being deployed together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.deployment_run_dto import DeploymentRunDto
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
    api_instance = scc_firewall_manager_sdk.DeviceDeploymentsApi(api_client)
    deployment_run_uid = 'deployment_run_uid_example' # str | The unique identifier, represented as a UUID, of the device deployment Run in SCC Firewall Manager.

    try:
        # Get Device Deployment Run
        api_response = api_instance.get_device_deployment_run(deployment_run_uid)
        print("The response of DeviceDeploymentsApi->get_device_deployment_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceDeploymentsApi->get_device_deployment_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_run_uid** | **str**| The unique identifier, represented as a UUID, of the device deployment Run in SCC Firewall Manager. | 

### Return type

[**DeploymentRunDto**](DeploymentRunDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The Device Deployment Run object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_deployment_runs**
> DeploymentRunDtoPage get_device_deployment_runs(limit=limit, offset=offset, q=q, sort=sort)

Get Device Deployment Runs

Get a list of device deployment runs in the SCC Firewall Manager Tenant. Each deployment run represents a group of devices being deployed together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.deployment_run_dto_page import DeploymentRunDtoPage
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
    api_instance = scc_firewall_manager_sdk.DeviceDeploymentsApi(api_client)
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Device Deployment Runs
        api_response = api_instance.get_device_deployment_runs(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of DeviceDeploymentsApi->get_device_deployment_runs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceDeploymentsApi->get_device_deployment_runs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**DeploymentRunDtoPage**](DeploymentRunDtoPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Device Deployment Run objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

