# Tenant


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the tenant in Security Cloud Control. | [optional] 
**name** | **str** | The name of the tenant in Security Cloud Control. | [optional] 
**security_cloud_control_enterprise_id** | **str** | The unique identifier, represented as a UUID, of the Security Cloud Control Enterprise this tenant is associated with. | [optional] 
**display_name** | **str** | A human-readable display name for the tenant. This is the tenant name displayed in the Security Cloud Control Web UI. | [optional] 
**pay_type** | **str** | An enum that describes the payment type of the tenant in Security Cloud Control. | [optional] 

## Example

```python
from cdo_sdk_python.models.tenant import Tenant

# TODO update the JSON string below
json = "{}"
# create an instance of Tenant from a JSON string
tenant_instance = Tenant.from_json(json)
# print the JSON string representation of the object
print(Tenant.to_json())

# convert the object into a dict
tenant_dict = tenant_instance.to_dict()
# create an instance of Tenant from a dict
tenant_form_dict = tenant.from_dict(tenant_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


