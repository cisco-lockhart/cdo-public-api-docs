# DomainSettings

Configuration that defines how Secure Client communicates with the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fqdn** | **str** | The fully qualified domain name (FQDN) of the device. Secure Client uses this to communicate with the device. | [optional] 
**certificate** | [**FmcObjectReference**](FmcObjectReference.md) |  | [optional] 
**interfaces** | [**List[FmcObjectReference]**](FmcObjectReference.md) | List of references to interface objects on the FMC, which defines the interfaces on the device that are used for ZTNA. | [optional] 

## Example

```python
from cdo_sdk_python.models.domain_settings import DomainSettings

# TODO update the JSON string below
json = "{}"
# create an instance of DomainSettings from a JSON string
domain_settings_instance = DomainSettings.from_json(json)
# print the JSON string representation of the object
print(DomainSettings.to_json())

# convert the object into a dict
domain_settings_dict = domain_settings_instance.to_dict()
# create an instance of DomainSettings from a dict
domain_settings_form_dict = domain_settings.from_dict(domain_settings_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


