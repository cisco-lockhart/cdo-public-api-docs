# UniversalZtnaSettings


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**deployed** | [**Deployed**](Deployed.md) |  | [optional] 
**staged** | [**Staged**](Staged.md) |  | [optional] 
**universal_ztna_configured** | **bool** | Indicates whether a device is configured for Zero Trust Network Access (ZTNA). | [optional] 
**universal_ztna_enabled** | **bool** | Indicates whether a device is enabled for Zero Trust Network Access (ZTNA). | [optional] 
**universal_ztna_supported** | **bool** | Indicates whether a device supports Zero Trust Network Access (ZTNA). | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.universal_ztna_settings import UniversalZtnaSettings

# TODO update the JSON string below
json = "{}"
# create an instance of UniversalZtnaSettings from a JSON string
universal_ztna_settings_instance = UniversalZtnaSettings.from_json(json)
# print the JSON string representation of the object
print(UniversalZtnaSettings.to_json())

# convert the object into a dict
universal_ztna_settings_dict = universal_ztna_settings_instance.to_dict()
# create an instance of UniversalZtnaSettings from a dict
universal_ztna_settings_form_dict = universal_ztna_settings.from_dict(universal_ztna_settings_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


