# MspManagedTemplate


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_state** | [**ConfigState**](ConfigState.md) |  | [optional] 
**managed_tenant_display_name** | **str** | The display name of the managed tenant in CDO. | [optional] 
**managed_tenant_name** | **str** | The name of the managed tenant in CDO. | [optional] 
**managed_tenant_region** | **str** | The region of the managed tenant in CDO. This is the region where the template is located. | [optional] 
**managed_tenant_uid** | **str** | The unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control that this template belongs to. | [optional] 
**name** | **str** | The name of the template in CDO. Template names are unique in Security Cloud Control. | 
**software_version** | **str** | The software version that the template is applicable to. | [optional] 
**template_type** | [**EntityType**](EntityType.md) |  | [optional] 
**tenant_uid** | **str** |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the template in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_template import MspManagedTemplate

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTemplate from a JSON string
msp_managed_template_instance = MspManagedTemplate.from_json(json)
# print the JSON string representation of the object
print(MspManagedTemplate.to_json())

# convert the object into a dict
msp_managed_template_dict = msp_managed_template_instance.to_dict()
# create an instance of MspManagedTemplate from a dict
msp_managed_template_form_dict = msp_managed_template.from_dict(msp_managed_template_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


