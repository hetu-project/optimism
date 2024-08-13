// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { CommonBase } from "forge-std/Base.sol";
import { console2 as console } from "forge-std/console2.sol";

/// @title ScriptExample
/// @notice ScriptExample is an example script. The Go forge script code tests that it can run this.
contract ScriptExample is CommonBase {
    /// @notice example function, runs through basic cheat-codes and console logs.
    function run() public {
        bool x = vm.envOr("EXAMPLE_BOOL", false);
        console.log("bool value from env", x);

        console.log("contract addr", address(this));
        console.log("contract nonce", vm.getNonce(address(this)));
        console.log("sender addr", address(msg.sender));
        console.log("sender nonce", vm.getNonce(address(msg.sender)));

        string memory json = '{"root_key": {"hello": {"a": 1, "b": 2}}}';
        string[] memory keys = vm.parseJsonKeys(json, ".root_key.hello");
        console.log("keys", keys[0], keys[1]);

        this.hello("from original");
        vm.startPrank(address(uint160(0x42)));
        this.hello("from prank 1");
        console.log("parent scope", address(msg.sender), address(this));
        this.hello("from prank 2");
        vm.stopPrank();
        this.hello("from original again");

        console.log("done!");
    }

    /// @notice example external function, to force a CALL, and test vm.startPrank with.
    function hello(string calldata _v) external view {
        console.log("hello addr", _v, address(msg.sender));
    }
}
