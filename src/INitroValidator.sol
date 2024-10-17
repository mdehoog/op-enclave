// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

interface INitroValidator {
    /// @notice Verifies an AWS Nitro attestation
    /// @param attestation The attestation document
    /// @param maxAge Maximum age of the attestation in seconds
    /// @return enclavePubKey The enclave's public key
    /// @return pcr0 User data included in the attestation
    function validateAttestation(
        bytes memory attestation,
        uint256 maxAge
    ) external returns (bytes memory enclavePubKey, bytes memory pcr0);
}