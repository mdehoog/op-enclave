// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "@eth-optimism-bedrock/src/universal/interfaces/ISemver.sol";

contract SystemConfigGlobal is OwnableUpgradeable, ISemver {
    /// @notice The address of the proposer.
    address public proposer; // TODO allow this to be overridden in the per-chain system config

    /// @notice Mapping of valid PCR0s attested from AWS Nitro.
    mapping(bytes32 => bool) public validPCR0s;

    /// @notice Mapping of valid signers attested from AWS Nitro.
    mapping(address => bool) public validSigners;

    /// @notice Semantic version.
    /// @custom:semver 0.0.1
    function version() public pure virtual returns (string memory) {
        return "0.0.1";
    }

    function initialize(address _owner) public initializer {
        __Ownable_init();
        transferOwnership(_owner);
    }

    function setProposer(address _proposer) external onlyOwner {
        proposer = _proposer;
    }

    function registerPCR0(bytes calldata pcr0) external onlyOwner {
        validPCR0s[keccak256(pcr0)] = true;
    }

    function deregisterPCR0(bytes calldata pcr0) external onlyOwner {
        delete validPCR0s[keccak256(pcr0)];
    }

    function registerSigner(bytes calldata attestation) external onlyOwner {
        // TODO validate AWS attestation, check PCR0, then add public key to mapping of valid signers
        // https://github.com/marlinprotocol/NitroProver
        revert("Not implemented");
    }

    // TODO remove this method once the above method is implemented
    function registerSignerAddress(address signer) external onlyOwner {
        validSigners[signer] = true;
    }

    function deregisterSigner(address signer) external onlyOwner {
        delete validSigners[signer];
    }
}
