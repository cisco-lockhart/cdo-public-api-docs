# Sdc


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**public_key** | [**PublicKey**](PublicKey.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the SDC in SCC. | [optional] 
**name** | **str** | The name of the SDC in CDO. SDC names are unique in SCC. | 
**software_version** | **str** | The software version running on the SDC. | [optional] 
**ip_address** | **str** | The IP address of the SDC. | [optional] 
**status** | [**Status**](Status.md) |  | [optional] 
**last_heartbeat** | **datetime** | The time (UTC; represented using the RFC-3339 standard) that a heartbeat was last received from the SDC. This serves as an indicator of the health of the SDC. | [optional] 
**bootstrap_data** | **str** | The bootstrap data is information used to automatically configure the SDC during its initial setup. This data is base64 encoded and includes essential details like the unique registration token and customer-specific settings that enable the SDC to communicate with and send data to SCC. This field is populated only if the SDC is not onboarded. | [optional] 

## Example

```python
from cdo_sdk_python.models.sdc import Sdc

# TODO update the JSON string below
json = "{}"
# create an instance of Sdc from a JSON string
sdc_instance = Sdc.from_json(json)
# print the JSON string representation of the object
print(Sdc.to_json())

# convert the object into a dict
sdc_dict = sdc_instance.to_dict()
# create an instance of Sdc from a dict
sdc_form_dict = sdc.from_dict(sdc_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


