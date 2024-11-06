# Meta


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cloud_connector_ips** | **List[str]** | The IP addresses from which to allow inbound access to your network if you wish to connect to your devices through Security Cloud Control&#39;s Cloud Connector. | [optional] 

## Example

```python
from cdo_sdk_python.models.meta import Meta

# TODO update the JSON string below
json = "{}"
# create an instance of Meta from a JSON string
meta_instance = Meta.from_json(json)
# print the JSON string representation of the object
print(Meta.to_json())

# convert the object into a dict
meta_dict = meta_instance.to_dict()
# create an instance of Meta from a dict
meta_form_dict = meta.from_dict(meta_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


