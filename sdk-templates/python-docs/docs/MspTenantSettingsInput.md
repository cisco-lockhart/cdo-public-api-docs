# MspTenantSettingsInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**collect_asa_health_metrics** | **bool** | Specify whether to enable or disable the collection of health metrics for ASAs. | [optional] 
**tenant_uids** | **List[str]** | A set of unique identifiers, representing the tenants in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_tenant_settings_input import MspTenantSettingsInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspTenantSettingsInput from a JSON string
msp_tenant_settings_input_instance = MspTenantSettingsInput.from_json(json)
# print the JSON string representation of the object
print(MspTenantSettingsInput.to_json())

# convert the object into a dict
msp_tenant_settings_input_dict = msp_tenant_settings_input_instance.to_dict()
# create an instance of MspTenantSettingsInput from a dict
msp_tenant_settings_input_form_dict = msp_tenant_settings_input.from_dict(msp_tenant_settings_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


