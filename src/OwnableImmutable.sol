// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/utils/ContextUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

/**
 * @dev Modification of OpenZeppelin's Ownable contract, with an immutable owner.
 */
abstract contract OwnableImmutable is Initializable, ContextUpgradeable {
    address public immutable owner;

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor(address _owner) {
        owner = _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner == _msgSender(), "OwnableImmutable: caller is not the owner");
    }
}
