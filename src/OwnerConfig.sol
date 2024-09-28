// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract OwnerConfig is Ownable {
    constructor(address _owner) {
        transferOwnership(_owner);
    }

    function checkOwner(address sender) external view {
        require(owner() == sender, "Ownable: caller is not the owner");
    }
}
