# MspManagedTenant

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**display_name** | **str** | A human-readable display name for the tenant. This is the tenant name displayed in the CDO Web UI. | [optional] 
**uid** | **str** | The unique identifier of the tenant in CDO. | [optional] 
**name** | **str** | The name of the device in CDO. Device names are unique in CDO. | [optional] 
**region** | **str** | The CDO region the tenant exists in. | [optional] 

## Example

```python
from cdo_sdk_python.models.msp_managed_tenant import MspManagedTenant

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


