# DuoAdminPanelCreateOrUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A human-readable name for the device. | 
**host** | **str** | The address of the Duo Admin Panel to onboard. | [optional] 
**integration_key** | **str** | The integration key of the Admin API application used to authenticate with Duo Admin Panel. | 
**secret_key** | **str** | The secret key of the Admin API application used to authenticate with Duo Admin Panel. | 
**labels** | [**Labels**](Labels.md) |  | [optional] 

## Example

```python
from cdo_python_sdk.models.duo_admin_panel_create_or_update_input import DuoAdminPanelCreateOrUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of DuoAdminPanelCreateOrUpdateInput from a JSON string
duo_admin_panel_create_or_update_input_instance = DuoAdminPanelCreateOrUpdateInput.from_json(json)
# print the JSON string representation of the object
print(DuoAdminPanelCreateOrUpdateInput.to_json())

# convert the object into a dict
duo_admin_panel_create_or_update_input_dict = duo_admin_panel_create_or_update_input_instance.to_dict()
# create an instance of DuoAdminPanelCreateOrUpdateInput from a dict
duo_admin_panel_create_or_update_input_form_dict = duo_admin_panel_create_or_update_input.from_dict(duo_admin_panel_create_or_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


