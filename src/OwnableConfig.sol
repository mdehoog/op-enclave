// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "./OwnerConfig.sol";
import "@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

abstract contract OwnableConfig is Initializable, ContextUpgradeable {
    OwnerConfig public immutable ownerConfig;

    constructor(OwnerConfig  _ownerConfig) {
        ownerConfig = _ownerConfig;
    }

    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    function owner() public view virtual returns (address) {
        return ownerConfig.owner();
    }

    function _checkOwner() internal view virtual {
        ownerConfig.checkOwner(_msgSender());
    }
}
