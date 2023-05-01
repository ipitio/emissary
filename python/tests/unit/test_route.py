import pytest

from ambassador.envoy.v3.v3route import V3Route
from tests.utils import econf_compile, econf_foreach_hcm, module_and_mapping_manifests


def _test_route(yaml, expectations={}):
    econf = econf_compile(yaml)

    def check(typed_config):
        # Find the one and virtual host in the route config
        vhosts = typed_config["route_config"]["virtual_hosts"]
        assert len(vhosts) == 1

        # Find the httpbin route. Run our expectations over that.
        routes = vhosts[0]["routes"]
        for r in routes:
            # Keep going until we find a real route
            if "route" not in r:
                continue

            # Keep going until we find a prefix match for /httpbin/
            match = r["match"]
            if "prefix" not in match or match["prefix"] != "/httpbin/":
                continue

            assert "route" in r
            route = r["route"]
            for key, expected in expectations.items():
                print("checking key %s" % key)
                assert key in route
                assert route[key] == expected
            break

    econf_foreach_hcm(econf, check)


@pytest.mark.compilertest
def test_timeout_ms():
    # If we do not set the config, we should get the default 3000ms.
    yaml = module_and_mapping_manifests(None, [])
    _test_route(yaml, expectations={"timeout": "3.000s"})


@pytest.mark.compilertest
def test_timeout_ms_module():
    # If we set a default on the Module, it should override the usual default of 3000ms.
    yaml = module_and_mapping_manifests(["cluster_request_timeout_ms: 4000"], [])
    _test_route(yaml, expectations={"timeout": "4.000s"})


@pytest.mark.compilertest
def test_timeout_ms_mapping():
    # If we set a default on the Module, it should override the usual default of 3000ms.
    yaml = module_and_mapping_manifests(None, ["timeout_ms: 1234"])
    _test_route(yaml, expectations={"timeout": "1.234s"})


@pytest.mark.compilertest
def test_timeout_ms_both():
    # If we set a default on the Module, it should override the usual default of 3000ms.
    yaml = module_and_mapping_manifests(["cluster_request_timeout_ms: 9000"], ["timeout_ms: 5001"])
    _test_route(yaml, expectations={"timeout": "5.001s"})


@pytest.mark.compilertest
def test_generate_headers_to_add():
    input_headers = {
        "x-test-proto": "%PROTOCOL%",
        "x-test-ip": "%DOWNSTREAM_REMOTE_ADDRESS_WITHOUT_PORT%",
        "x-test-static": "This is a test header",
        "x-test-append": {"append": False, "value": "this will not append header if already exist"},
        "x-test-no-append": {"append": True, "value": "this will allow appending if header exist"},
    }

    expected = [
        {"header": {"key": "x-test-proto", "value": "%PROTOCOL%"}, "append": True},
        {
            "header": {"key": "x-test-ip", "value": "%DOWNSTREAM_REMOTE_ADDRESS_WITHOUT_PORT%"},
            "append": True,
        },
        {"header": {"key": "x-test-static", "value": "This is a test header"}, "append": True},
        {
            "header": {
                "key": "x-test-append",
                "value": "this will not append header if already exist",
            },
            "append": False,
        },
        {
            "header": {
                "key": "x-test-no-append",
                "value": "this will allow appending if header exist",
            },
            "append": True,
        },
    ]

    result = V3Route.generate_headers_to_add(input_headers)
    assert expected == result
