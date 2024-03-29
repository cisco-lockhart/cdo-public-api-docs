# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from cdo_sdk_python.api.inventory_api import InventoryApi


class TestInventoryApi(unittest.TestCase):
    """InventoryApi unit test stubs"""

    def setUp(self) -> None:
        self.api = InventoryApi()

    def tearDown(self) -> None:
        pass

    def test_create_duo_admin_panel(self) -> None:
        """Test case for create_duo_admin_panel

        Onboard Duo Admin Panel
        """
        pass

    def test_create_ftd_device(self) -> None:
        """Test case for create_ftd_device

        Create FTD device.
        """
        pass

    def test_delete_cd_fmc_managed_ftd_device(self) -> None:
        """Test case for delete_cd_fmc_managed_ftd_device

        Delete cdFMC managed FTD device
        """
        pass

    def test_delete_cloud_service(self) -> None:
        """Test case for delete_cloud_service

        Delete Cloud Service
        """
        pass

    def test_delete_device(self) -> None:
        """Test case for delete_device

        Delete Device
        """
        pass

    def test_delete_device_manager(self) -> None:
        """Test case for delete_device_manager

        Delete Device Manager
        """
        pass

    def test_delete_template_device(self) -> None:
        """Test case for delete_template_device

        Delete Template Device
        """
        pass

    def test_deploy_asa_device_changes(self) -> None:
        """Test case for deploy_asa_device_changes

        Deploy ASA device changes
        """
        pass

    def test_finish_onboarding_ftd_device(self) -> None:
        """Test case for finish_onboarding_ftd_device

        Register FTD device.
        """
        pass

    def test_get_cloud_service(self) -> None:
        """Test case for get_cloud_service

        Get Cloud Service
        """
        pass

    def test_get_cloud_services(self) -> None:
        """Test case for get_cloud_services

        Get Cloud Services
        """
        pass

    def test_get_device(self) -> None:
        """Test case for get_device

        Get Device
        """
        pass

    def test_get_device_manager(self) -> None:
        """Test case for get_device_manager

        Get Device Manager
        """
        pass

    def test_get_device_managers(self) -> None:
        """Test case for get_device_managers

        Get Device Managers
        """
        pass

    def test_get_devices(self) -> None:
        """Test case for get_devices

        Get Devices
        """
        pass

    def test_get_template_device(self) -> None:
        """Test case for get_template_device

        Get Template Device
        """
        pass

    def test_get_template_devices(self) -> None:
        """Test case for get_template_devices

        Get Template Devices
        """
        pass

    def test_modify_cloud_service(self) -> None:
        """Test case for modify_cloud_service

        Modify Cloud Service
        """
        pass

    def test_modify_device(self) -> None:
        """Test case for modify_device

        Modify Device
        """
        pass

    def test_modify_device_manager(self) -> None:
        """Test case for modify_device_manager

        Modify Device Manager
        """
        pass

    def test_modify_template_device(self) -> None:
        """Test case for modify_template_device

        Modify Template Device.
        """
        pass

    def test_onboard_asa_device(self) -> None:
        """Test case for onboard_asa_device

        Onboard ASA device
        """
        pass

    def test_onboard_ios_device(self) -> None:
        """Test case for onboard_ios_device

        Onboard IOS Device
        """
        pass

    def test_read_asa_device_configuration(self) -> None:
        """Test case for read_asa_device_configuration

        Read ASA device configuration
        """
        pass

    def test_reconnect_asa_device(self) -> None:
        """Test case for reconnect_asa_device

        Reconnect ASA device
        """
        pass


if __name__ == '__main__':
    unittest.main()
