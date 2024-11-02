# AsaCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A human-readable name for the device. | 
**device_address** | **str** | The address of the device to onboard, specified in the format &#x60;host:port&#x60;. | [optional] 
**username** | **str** | The username used to authenticate with the device. | 
**password** | **str** | The password used to authenticate with the device. | 
**connector_type** | [**ConnectorType**](ConnectorType.md) |  | 
**ignore_certificate** | **bool** | Set this attribute to true if you do not want SCC to validate the certificate of this device before onboarding. | [optional] [default to False]
**connector_name** | **str** | The name of the Secure Device Connector (SDC) that will be used to communicate with the device. This value is not required if the connector type selected is Cloud Connector (CDG). | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.asa_create_or_update_input import AsaCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AsaCreateOrUpdateInput from a JSON string
asa_create_or_update_input_instance = AsaCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(AsaCreateOrUpdateInput.to_json())

# convert the object into a dict
asa_create_or_update_input_dict = asa_create_or_update_input_instance.to_dict()
# create an instance of AsaCreateOrUpdateInput from a dict
asa_create_or_update_input_form_dict = asa_create_or_update_input.from_dict(asa_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


