# IosCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connector_name** | **str** | The name of the Secure Device Connector (SDC) that will be used to communicate with the device. | 
**device_address** | **str** | The address of the device to onboard, specified in the format &#x60;host:port&#x60;. | [optional] 
**ignore_certificate** | **bool** | Set this attribute to true if you do not want Security Cloud Control to validate the certificate of this device before onboarding. | [optional] [default to False]
**labels** | [**Labels**](Labels.md) |  | [optional] 
**name** | **str** | A human-readable name for the device. | 
**password** | **str** | The password used to authenticate with the device. | 
**username** | **str** | The username used to authenticate with the device. | 

## Example

```python
from scc_firewall_manager_sdk.models.ios_create_or_update_input import IosCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of IosCreateOrUpdateInput from a JSON string
ios_create_or_update_input_instance = IosCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(IosCreateOrUpdateInput.to_json())

# convert the object into a dict
ios_create_or_update_input_dict = ios_create_or_update_input_instance.to_dict()
# create an instance of IosCreateOrUpdateInput from a dict
ios_create_or_update_input_form_dict = ios_create_or_update_input.from_dict(ios_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


