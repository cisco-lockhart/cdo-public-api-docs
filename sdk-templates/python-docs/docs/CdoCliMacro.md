# CdoCliMacro


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_type** | [**EntityType**](EntityType.md) |  | 
**name** | **str** | The name of the CLI macro. | [optional] 
**parameters** | **List[str]** | A set of parameters provided in the script | [optional] 
**script** | **str** | The script executed to generate the CLI result. | [optional] 
**system_defined** | **bool** | Indicating whether this CLI macro is system-defined. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the CLI macro. | 

## Example

```python
from scc_firewall_manager_sdk.models.cdo_cli_macro import CdoCliMacro

# TODO update the JSON string below
json = "{}"
# create an instance of CdoCliMacro from a JSON string
cdo_cli_macro_instance = CdoCliMacro.from_json(json)
# print the JSON string representation of the object
print(CdoCliMacro.to_json())

# convert the object into a dict
cdo_cli_macro_dict = cdo_cli_macro_instance.to_dict()
# create an instance of CdoCliMacro from a dict
cdo_cli_macro_form_dict = cdo_cli_macro.from_dict(cdo_cli_macro_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


