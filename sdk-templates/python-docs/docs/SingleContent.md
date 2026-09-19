# SingleContent

A single, non-group content literal. Network literals carry a literal, URL literals a url, and service literals a protocol and serviceValue.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**dns_resolution** | **str** | For FQDN network objects sourced from FTD/cdFMC, the DNS resolution family that governs which DNS record types FTD resolves at policy-enforcement time. Populated only for FQDN objects; absent otherwise. | [optional] 
**literal** | **str** | The literal content of the network object | [optional] 
**protocol** | **str** | The service object protocol. Populated only for service objects; absent otherwise. | [optional] 
**service_value** | [**ServiceObjectValueContent**](ServiceObjectValueContent.md) |  | [optional] 
**url** | **str** | The URL literal | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.single_content import SingleContent

# TODO update the JSON string below
json = "{}"
# create an instance of SingleContent from a JSON string
single_content_instance = SingleContent.from_json(json)
# print the JSON string representation of the object
print(SingleContent.to_json())

# convert the object into a dict
single_content_dict = single_content_instance.to_dict()
# create an instance of SingleContent from a dict
single_content_form_dict = single_content.from_dict(single_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


