# CurrentTenantResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**display_name** | **str** | The human-readable tenant name displayed in the Security Cloud Control Web UI. | [optional] 
**feature_flags** | **List[str]** | The feature flags enabled for the tenant. | [optional] 
**name** | **str** | The name of the tenant in Security Cloud Control. | [optional] 
**security_cloud_control_enterprise_id** | **str** | The unique identifier of the Security Cloud Control enterprise associated with the tenant. | [optional] 
**ssx_tenant_uid** | **str** | The unique identifier of the tenant in Security Services Exchange. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the tenant in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.current_tenant_response import CurrentTenantResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CurrentTenantResponse from a JSON string
current_tenant_response_instance = CurrentTenantResponse.from_json(json)
# print the JSON string representation of the object
print(CurrentTenantResponse.to_json())

# convert the object into a dict
current_tenant_response_dict = current_tenant_response_instance.to_dict()
# create an instance of CurrentTenantResponse from a dict
current_tenant_response_form_dict = current_tenant_response.from_dict(current_tenant_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


