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

from cdo_sdk_python.api.changelogs_api import ChangelogsApi


class TestChangelogsApi(unittest.TestCase):
    """ChangelogsApi unit test stubs"""

    def setUp(self) -> None:
        self.api = ChangelogsApi()

    def tearDown(self) -> None:
        pass

    def test_get_changelog(self) -> None:
        """Test case for get_changelog

        Get Change Log
        """
        pass

    def test_get_changelogs(self) -> None:
        """Test case for get_changelogs

        List Change Logs
        """
        pass


if __name__ == '__main__':
    unittest.main()