# FmcObjectOverride


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Specify the name of the object to override. | 
**type** | **str** | Specify the type of the object to override. The type should match the type of the object used in the template. | 
**uid** | **str** | Specify the unique identifier, represented as a UUID, of the object to override. | 
**value** | **str** | Specify the value to override the object with. | 

## Example

```python
from scc_firewall_manager_sdk.models.fmc_object_override import FmcObjectOverride

# TODO update the JSON string below
json = "{}"
# create an instance of FmcObjectOverride from a JSON string
fmc_object_override_instance = FmcObjectOverride.from_json(json)
# print the JSON string representation of the object
print(FmcObjectOverride.to_json())

# convert the object into a dict
fmc_object_override_dict = fmc_object_override_instance.to_dict()
# create an instance of FmcObjectOverride from a dict
fmc_object_override_form_dict = fmc_object_override.from_dict(fmc_object_override_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


