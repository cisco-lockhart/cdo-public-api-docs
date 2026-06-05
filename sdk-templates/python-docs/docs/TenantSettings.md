# TenantSettings


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ai_assistant** | **bool** | Indicates if the tenant has the AI assistant enabled. | [optional] 
**allow_cross_region_tenant_addition** | **bool** | Indicates whether adding tenants from different regions to this MSP Portal is allowed. | [optional] 
**asa_health_metrics** | **bool** | Indicates if the tenant collects ASA health metrics. | [optional] 
**auto_accept_device_changes** | **bool** | Indicates if changes made out-of-band on devices on the tenant are automatically accepted without manual approval. | [optional] 
**auto_discover_on_prem_fmcs** | **bool** | Indicates if the system automatically discovers on-premise FMCs. | [optional] 
**change_request_support** | **bool** | Indicates if the tenant supports change requests. | [optional] 
**conflict_detection_interval** | [**ConflictDetectionInterval**](ConflictDetectionInterval.md) |  | [optional] 
**deny_cisco_support_access_to_tenant** | **bool** | Indicates if Cisco support is denied access to the tenant. | [optional] 
**msp_portal_uid** | **str** | The unique identifier, represented as a UUID, of the MSP portal associated with the tenant. | [optional] 
**multicloud_defense** | **bool** | Indicates if the tenant has the multicloud defense enabled. | [optional] 
**object_sharing_enabled** | **bool** | Indicates if the tenant has Object Sharing enabled. | [optional] 
**policy_analyzer** | **bool** | Indicates if the tenant has the policy analyzer enabled. | [optional] 
**scheduled_deployments** | **bool** | Indicates if the tenant has scheduled deployments enabled. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the tenant in Security Cloud Control. | [optional] 
**web_analytics** | **bool** | Indicates if web analytics are enabled for the tenant. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.tenant_settings import TenantSettings

# TODO update the JSON string below
json = "{}"
# create an instance of TenantSettings from a JSON string
tenant_settings_instance = TenantSettings.from_json(json)
# print the JSON string representation of the object
print(TenantSettings.to_json())

# convert the object into a dict
tenant_settings_dict = tenant_settings_instance.to_dict()
# create an instance of TenantSettings from a dict
tenant_settings_form_dict = tenant_settings.from_dict(tenant_settings_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


