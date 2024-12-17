# FmcObjectReference

Reference to the network object on the FMC, defining the IPv4 address used for Network Address Translation (NAT) and/or Port Address Translation (PAT). Note: Required only if the device needs to support more than 65,000 simultaneous active connections.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the FMC Object. | [optional] 
**link** | **str** | The endpoint to access this resource from. | [optional] 

## Example

```python
from cdo_sdk_python.models.fmc_object_reference import FmcObjectReference

# TODO update the JSON string below
json = "{}"
# create an instance of FmcObjectReference from a JSON string
fmc_object_reference_instance = FmcObjectReference.from_json(json)
# print the JSON string representation of the object
print(FmcObjectReference.to_json())

# convert the object into a dict
fmc_object_reference_dict = fmc_object_reference_instance.to_dict()
# create an instance of FmcObjectReference from a dict
fmc_object_reference_form_dict = fmc_object_reference.from_dict(fmc_object_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


