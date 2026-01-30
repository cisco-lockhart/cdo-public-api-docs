# Staged


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_deploy_enabled** | **bool** | Specifies whether changes to ZTNA settings are automatically deployed to the device. Note: This applies only to ZTNA-specific changes and does not affect the deployment of other pending changes. | [optional] 
**domain_settings** | [**List[DomainSettings]**](DomainSettings.md) | Configuration that defines how Secure Client communicates with the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.staged import Staged

# TODO update the JSON string below
json = "{}"
# create an instance of Staged from a JSON string
staged_instance = Staged.from_json(json)
# print the JSON string representation of the object
print(Staged.to_json())

# convert the object into a dict
staged_dict = staged_instance.to_dict()
# create an instance of Staged from a dict
staged_form_dict = staged.from_dict(staged_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


