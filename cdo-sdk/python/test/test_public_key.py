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

from cdo_sdk_python.models.public_key import PublicKey

class TestPublicKey(unittest.TestCase):
    """PublicKey unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> PublicKey:
        """Test PublicKey
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `PublicKey`
        """
        model = PublicKey()
        if include_optional:
            return PublicKey(
                base64_encoded_key = 'c3NoLXJzYSBBQUFBQjNOemFDMXljMkVBQUFBREFRQUJBQUFCZ1FEZGpxeDNNMjh3ZHE2Rk11R2NFZHJqaGl3SFp5NDE0WXZFUmRDcVdnRXgrSXloNjVteHlkaVh1OVpoMXUyVTRtNDlnL0VUL2JxTklFNGgraUViQURjcjlScVFGRGZ4RC9pK1ZPSkhiQzQ0UERhZHFJd21vUGNoazYvQ0RoZVlJSmZTM0xTNmlkaC9SRGVFVHpnZ3lTaUI2Mm5yMnRmTkJ3V0ZScjV1Sko5dkNvdUxKRVBZbDBVMkpZNnBjd0paNk1lRDg5dU4rTjlHWFN2Vlh4bEZKNXg0VThReGFCMzJuNHZoekNiUzlYSVg1bGJJQnVIUEZ1bmMrNThPaUFzS0dwTTZ1NzhVR2V2TndOVzU0eVZmU2c4Q01XQ09vM1hiNTIrMnU2VHZlcE1BT2ZFU290YUd5NHV1RTBnUHYwSnowano3ZnFGTCt4d1AzNnNVY2pPRlIzQ1VhbEVpUDEyT2tTeEhreVNMUXJ3a2lFTVUvQ3VPUjdOWkdjUFd0dkVBaUZOTEN0VHhzY25Ma0xNNENkUEt3WnZQV3l3cHAyVGdValU4MEFaWkJZaGZBWTd3SFFQdDhrTkROMURhMWpWaHlwMWlycE5VbTEvaDNrS2srOFdFOVV2RjNDeVgyL1NqWUpPR2lMUWdVaXFUL3E5UkNMZUdVWStzaUFwS3ZyNS91UkU9IHNpd2FycmllQFNJV0FSUklFLU0tS1JBRgo=',
                key_id = 'example-key-id'
            )
        else:
            return PublicKey(
        )
        """

    def testPublicKey(self):
        """Test PublicKey"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()