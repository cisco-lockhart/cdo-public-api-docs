# FmcObjectReference


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **str** | The endpoint to access this resource from. | [optional] 
**name** | **str** | The name of the FMC Object. | [optional] 
**type** | **str** | The type of the FMC Object. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the FMC Object. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.fmc_object_reference import FmcObjectReference

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


