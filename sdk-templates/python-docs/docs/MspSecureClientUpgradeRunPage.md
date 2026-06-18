# MspSecureClientUpgradeRunPage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[MspSecureClientUpgradeRunDto]**](MspSecureClientUpgradeRunDto.md) | The list of items retrieved. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_run_page import MspSecureClientUpgradeRunPage

# TODO update the JSON string below
json = "{}"
# create an instance of MspSecureClientUpgradeRunPage from a JSON string
msp_secure_client_upgrade_run_page_instance = MspSecureClientUpgradeRunPage.from_json(json)
# print the JSON string representation of the object
print(MspSecureClientUpgradeRunPage.to_json())

# convert the object into a dict
msp_secure_client_upgrade_run_page_dict = msp_secure_client_upgrade_run_page_instance.to_dict()
# create an instance of MspSecureClientUpgradeRunPage from a dict
msp_secure_client_upgrade_run_page_form_dict = msp_secure_client_upgrade_run_page.from_dict(msp_secure_client_upgrade_run_page_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


