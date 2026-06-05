# MspManagedTenantSettingsDto

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asa_health_metrics** | **bool** | Indicates whether ASA health metrics are enabled. | 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant these settings belong to. | 
**uid** | **str** | The unique identifier, represented as a UUID, of the tenant settings record. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_tenant_settings_dto import MspManagedTenantSettingsDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTenantSettingsDto from a JSON string
msp_managed_tenant_settings_dto_instance = MspManagedTenantSettingsDto.from_json(json)
# print the JSON string representation of the object
print(MspManagedTenantSettingsDto.to_json())

# convert the object into a dict
msp_managed_tenant_settings_dto_dict = msp_managed_tenant_settings_dto_instance.to_dict()
# create an instance of MspManagedTenantSettingsDto from a dict
msp_managed_tenant_settings_dto_form_dict = msp_managed_tenant_settings_dto.from_dict(msp_managed_tenant_settings_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


