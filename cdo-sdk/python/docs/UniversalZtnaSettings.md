# UniversalZtnaSettings

(FMC-managed FTDs only) Universal Zero-Trust Network Access (ZTNA) configuration.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_deploy_enabled** | **bool** | Specifies whether changes to ZTNA settings are automatically deployed to the device. Note: This applies only to ZTNA-specific changes and does not affect the deployment of other pending changes. | [optional] 
**domain_settings** | [**List[DomainSettings]**](DomainSettings.md) | Configuration that defines how Secure Client communicates with the device. | [optional] 
**source_nat_v4** | [**FmcObjectReference**](FmcObjectReference.md) |  | [optional] 
**source_nat_v6** | [**FmcObjectReference**](FmcObjectReference.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.universal_ztna_settings import UniversalZtnaSettings

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


