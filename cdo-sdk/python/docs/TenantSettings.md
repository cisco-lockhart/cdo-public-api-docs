# TenantSettings


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the tenant in SCC. | [optional] 
**change_request_support** | **bool** | Indicates if the tenant supports change requests. | [optional] 
**auto_accept_device_changes** | **bool** | Indicates if changes made out-of-band on devices on the tenant are automatically accepted without manual approval. | [optional] 
**web_analytics** | **bool** | Indicates if web analytics are enabled for the tenant. | [optional] 
**scheduled_deployments** | **bool** | Indicates if the tenant has scheduled deployments enabled. | [optional] 
**deny_cisco_support_access_to_tenant** | **bool** | Indicates if Cisco support is denied access to the tenant. | [optional] 
**multicloud_defense** | **bool** | Indicates if the tenant has the multicloud defense enabled. | [optional] 
**policy_analyzer** | **bool** | Indicates if the tenant has the policy analyzer enabled. | [optional] 
**ai_assistant** | **bool** | Indicates if the tenant has the AI assistant enabled. | [optional] 
**auto_discover_on_prem_fmcs** | **bool** | Indicates if the system automatically discovers on-premise FMCs. | [optional] 
**conflict_detection_interval** | [**ConflictDetectionInterval**](ConflictDetectionInterval.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.tenant_settings import TenantSettings

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


