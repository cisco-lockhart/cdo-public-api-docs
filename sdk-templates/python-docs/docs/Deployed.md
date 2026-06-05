# Deployed


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_deploy_enabled** | **bool** | Specifies whether changes to ZTNA settings are automatically deployed to the device. Note: This applies only to ZTNA-specific changes and does not affect the deployment of other pending changes. | [optional] 
**domain_settings** | [**List[DomainSettings]**](DomainSettings.md) | Configuration that defines how Secure Client communicates with the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.deployed import Deployed

# TODO update the JSON string below
json = "{}"
# create an instance of Deployed from a JSON string
deployed_instance = Deployed.from_json(json)
# print the JSON string representation of the object
print(Deployed.to_json())

# convert the object into a dict
deployed_dict = deployed_instance.to_dict()
# create an instance of Deployed from a dict
deployed_form_dict = deployed.from_dict(deployed_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


