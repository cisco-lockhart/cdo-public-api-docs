# InterfaceHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**buffer_overruns_avg** | **float** | Monitors the average number of times where incoming data exceeded buffer capacity, potentially leading to data loss. | [optional] 
**buffer_underruns_avg** | **float** | Tracks the average number of times the data buffer was insufficient to handle outgoing traffic, possibly causing transmission delays. | [optional] 
**drop_packets_avg** | **float** | Average number of packets dropped by the interface due to network congestion, buffer overflow, or errors. | [optional] 
**duplex_mode** | **str** | Configuration of the interface regarding data transmission, indicating whether it is set to full, half, or auto-duplex. | [optional] 
**input_bytes_avg** | **float** | Total amount of data received through the interface, providing insights into the volume of inbound traffic. | [optional] 
**input_errors_avg** | **float** | The average rate of erroneous packets received, indicative of issues like corruption or transmission errors. | [optional] 
**input_packet_size_avg** | **float** | Average size of packets received, useful for analysing the nature of inbound traffic. | [optional] 
**interface** | **str** | Identifier for a specific network interface on the FTD device, used for network traffic management and monitoring. | [optional] 
**interface_name** | **str** | The name assigned to the interface, facilitating easier identification and configuration. | [optional] 
**interface_type** | **str** | The physical or logical type of the interface (e.g., Ethernet, virtual, management). | [optional] 
**l2_decode_drops_avg** | **float** | The average number of packets that could not be processed due to issues at the Data Link layer, including protocol mismatches or corruption. | [optional] 
**link_status** | **str** | Indicates whether the physical link of the network interface is active (UP) or inactive (DOWN). The interface will be marked as DOWN if there is no traffic through the interface. | [optional] 
**operational_status** | **str** | Current state of the interface from a functional standpoint, influenced by both administrative settings and physical connectivity. | [optional] 
**output_bytes_avg** | **float** | Total data sent out through the interface, useful for tracking outbound traffic levels. | [optional] 
**output_errors_avg** | **float** | Average count of error-ridden packets sent from the device, pointing to problems in packet formation or hardware issues. | [optional] 
**output_packet_size_avg** | **float** | Average size of packets sent. Helps in understanding the traffic distribution and network load. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_health_metrics import InterfaceHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceHealthMetrics from a JSON string
interface_health_metrics_instance = InterfaceHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(InterfaceHealthMetrics.to_json())

# convert the object into a dict
interface_health_metrics_dict = interface_health_metrics_instance.to_dict()
# create an instance of InterfaceHealthMetrics from a dict
interface_health_metrics_form_dict = interface_health_metrics.from_dict(interface_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


