# SecureAccessMetadata


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**commit_status** | **str** | The deployment status of the Secure Access configuration associated with the device. | [optional] 
**num_resources** | **int** | The number of private Cisco Secure Access resources associated with the device. | [optional] 
**num_rules** | **int** | The number of Cisco Secure Access rules associated with the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.secure_access_metadata import SecureAccessMetadata

# TODO update the JSON string below
json = "{}"
# create an instance of SecureAccessMetadata from a JSON string
secure_access_metadata_instance = SecureAccessMetadata.from_json(json)
# print the JSON string representation of the object
print(SecureAccessMetadata.to_json())

# convert the object into a dict
secure_access_metadata_dict = secure_access_metadata_instance.to_dict()
# create an instance of SecureAccessMetadata from a dict
secure_access_metadata_form_dict = secure_access_metadata.from_dict(secure_access_metadata_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


