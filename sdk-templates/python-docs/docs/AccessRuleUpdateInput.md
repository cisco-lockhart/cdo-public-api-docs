# AccessRuleUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**destination_dynamic_object** | [**DestinationDynamicObjectContent**](DestinationDynamicObjectContent.md) |  | [optional] 
**destination_network** | [**DestinationNetworkContent**](DestinationNetworkContent.md) |  | [optional] 
**destination_port** | [**DestinationPortContent**](DestinationPortContent.md) |  | [optional] 
**index** | **int** | Access rule index position in Access Group ordered rule list. | [optional] 
**is_active_rule** | **bool** | Is active. True by default | [optional] 
**log_settings** | [**LogSettings**](LogSettings.md) |  | [optional] 
**protocol** | [**ProtocolContent**](ProtocolContent.md) |  | [optional] 
**remark** | **str** | A remark. | [optional] 
**rule_action** | **str** | The rule&#39;s action: PERMIT or DENY. | [optional] 
**rule_time_range** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**source_dynamic_object** | [**SourceDynamicObjectContent**](SourceDynamicObjectContent.md) |  | [optional] 
**source_network** | [**SourceNetworkContent**](SourceNetworkContent.md) |  | [optional] 
**source_port** | [**SourcePortContent**](SourcePortContent.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the Access Rule. | 

## Example

```python
from scc_firewall_manager_sdk.models.access_rule_update_input import AccessRuleUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AccessRuleUpdateInput from a JSON string
access_rule_update_input_instance = AccessRuleUpdateInput.from_json(json)
# print the JSON string representation of the object
print(AccessRuleUpdateInput.to_json())

# convert the object into a dict
access_rule_update_input_dict = access_rule_update_input_instance.to_dict()
# create an instance of AccessRuleUpdateInput from a dict
access_rule_update_input_form_dict = access_rule_update_input.from_dict(access_rule_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


