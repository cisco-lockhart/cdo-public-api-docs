# MspManagedTenant


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**display_name** | **str** | A human-readable display name for the tenant. This is the tenant name displayed in the Security Cloud Control Web UI. | [optional] 
**name** | **str** | The name of the device in CDO. Device names are unique in Security Cloud Control. | [optional] 
**region** | **str** | The Security Cloud Control region the tenant exists in. | [optional] 
**uid** | **str** | The unique identifier of the tenant in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_tenant import MspManagedTenant

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTenant from a JSON string
msp_managed_tenant_instance = MspManagedTenant.from_json(json)
# print the JSON string representation of the object
print(MspManagedTenant.to_json())

# convert the object into a dict
msp_managed_tenant_dict = msp_managed_tenant_instance.to_dict()
# create an instance of MspManagedTenant from a dict
msp_managed_tenant_form_dict = msp_managed_tenant.from_dict(msp_managed_tenant_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


