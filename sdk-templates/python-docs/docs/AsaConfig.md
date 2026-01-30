# AsaConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_on_cloud** | **str** | The configuration of the ASA device in Security Cloud Control. This may include changes staged on Security Cloud Control that have not been deployed to the device. | [optional] 
**config_on_device** | **str** | The running configuration on the ASA device. Note: this may not include changes made out-of-band directly on the device since the last Out-of-Band check ran on Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_config import AsaConfig

# TODO update the JSON string below
json = "{}"
# create an instance of AsaConfig from a JSON string
asa_config_instance = AsaConfig.from_json(json)
# print the JSON string representation of the object
print(AsaConfig.to_json())

# convert the object into a dict
asa_config_dict = asa_config_instance.to_dict()
# create an instance of AsaConfig from a dict
asa_config_form_dict = asa_config.from_dict(asa_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


