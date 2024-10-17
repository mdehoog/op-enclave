// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { CBORDecoding } from "marlinprotocol/solidity-cbor/CBORDecoding.sol";
import { CBOR } from "marlinprotocol/solidity-cbor/CBOREncoding.sol";
import { ByteParser } from "marlinprotocol/solidity-cbor/ByteParser.sol";
import { NitroProver } from "@marlinprotocol/NitroProver.sol";
import { INitroValidator } from "./INitroValidator.sol";

contract NitroValidator is NitroProver, INitroValidator {
    /// @notice based off of NitroProver's verifyAttestation. Will validate an attestation and return the public key and PRC0 used
    
    function validateAttestation(
        bytes memory attestation,
        uint256 maxAge
    ) external returns (bytes memory, bytes memory) {
        /* 
        https://github.com/aws/aws-nitro-enclaves-nsm-api/blob/main/docs/attestation_process.md#31-cose-and-cbor
        Attestation document is an array of 4 elements
        [
            protected:   Header,
            unprotected: Header,
            payload:     This field contains the serialized content to be signed,
            signature:   This field contains the computed signature value.
        ]
        */
        
        // GAS: Attestation decode gas ~62k
        bytes[] memory attestationDecoded = CBORDecoding.decodeArray(attestation);

        // TODO: confirm that the attestation is untagged CBOR structure
        // https://datatracker.ietf.org/doc/html/rfc8152#section-3.1
        // Protected header for COSE_Sign1
        bytes[2][] memory protectedHeader = CBORDecoding.decodeMapping(attestationDecoded[0]);
        // Protected header should have algorithm flag which is specified by 1
        require(ByteParser.bytesToUint64(protectedHeader[0][0]) == 1, "Not algo flag");
        // Algorithm should be ECDSA w/ SHA-384
        require(ByteParser.bytesToNegativeInt128(protectedHeader[0][1]) == -35, "Incorrect algorithm");
        // Protected header should just have sig algo flag
        require(protectedHeader.length == 1, "Only algo flag should be present");

        // Unprotected header for COSE_Sign1
        bytes[2][] memory unprotectedHeader = CBORDecoding.decodeMapping(attestationDecoded[1]);
        // Unprotected header should be empty
        require(unprotectedHeader.length == 0, "Unprotected header should be empty");

        bytes memory payload = attestationDecoded[2];
        (bytes memory certPubKey, bytes memory enclavePubKey, bytes memory pcr0) = _getAttestationDocKeysAndPCR0(payload, maxAge);

        // verify COSE signature as per https://www.rfc-editor.org/rfc/rfc9052.html#section-4.4
        bytes memory attestationSig = attestationDecoded[3];

        // create COSE structure
        // GAS: COSE structure creation gas ~42.7k
        // TODO: set CBOR length appropriately
        CBOR.CBORBuffer memory buf = CBOR.create(payload.length*2);
        CBOR.startFixedArray(buf, 4);
        // context to be written as Signature1 as COSE_Sign1 is used https://www.rfc-editor.org/rfc/rfc9052.html#section-4.4-2.1.1
        CBOR.writeString(buf, "Signature1");
        // Protected headers to be added https://www.rfc-editor.org/rfc/rfc9052.html#section-4.4-2.2
        CBOR.writeBytes(buf, attestationDecoded[0]);
        // externally supplied data is empty https://www.rfc-editor.org/rfc/rfc9052.html#section-4.4-2.4
        CBOR.writeBytes(buf, "");
        // Payload to be added https://www.rfc-editor.org/rfc/rfc9052.html#section-4.4-2.5
        CBOR.writeBytes(buf, payload);

        _processSignature(attestationSig, certPubKey, buf.buf.buf);
        return (enclavePubKey, pcr0);
    }

    /// @notice validates the attestation payload and returns the used public key, enclave public key and pcr0 used
    /// @return certPubKey certificate public key
    /// @return enclavePubKey enclave public key
    /// @return pcr0 platform configuration register 0 
    function _getAttestationDocKeysAndPCR0(bytes memory attestationPayload, uint256 maxAge) internal view returns(bytes memory certPubKey, bytes memory enclavePubKey, bytes memory pcr0) {
        // TODO: validate if this check is expected? https://github.com/aws/aws-nitro-enclaves-nsm-api/blob/main/docs/attestation_process.md?plain=1#L168
        require(attestationPayload.length <= 2**15, "Attestation too long");

        // validations as per https://github.com/aws/aws-nitro-enclaves-nsm-api/blob/main/docs/attestation_process.md#32-syntactical-validation
        // issuing Nitro hypervisor module ID
        // GAS: decoding takes ~173.5k gas
        bytes[2][] memory attestationStructure = CBORDecoding.decodeMapping(attestationPayload);
        bytes memory moduleId;
        bytes memory rawTimestamp;
        bytes memory digest;
        bytes memory rawPcrs;
        bytes memory certificate;
        bytes memory cabundle;
        bytes memory userData;

        for(uint256 i=0; i < attestationStructure.length; i++) {
            bytes32 keyHash = keccak256(attestationStructure[i][0]);
            if(keyHash == keccak256(bytes("module_id"))) {
                moduleId = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("timestamp"))) {
                rawTimestamp = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("digest"))) {
                digest = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("pcrs"))) {
                rawPcrs = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("certificate"))) {
                certificate = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("cabundle"))) {
                cabundle = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("public_key"))) {
                enclavePubKey = attestationStructure[i][1];
                continue;
            }
            if(keyHash == keccak256(bytes("user_data"))) {
                userData = attestationStructure[i][1];
                continue;
            }
        }

        require(moduleId.length != 0, "Invalid module id");

        uint64 timestamp = ByteParser.bytesToUint64(rawTimestamp);
        require(timestamp != 0, "invalid timestamp");
        require(timestamp + maxAge > block.timestamp, "attestation too old");

        require(bytes32(digest) == bytes32("SHA384"), "invalid digest algo");

        bytes[2][] memory pcrs = CBORDecoding.decodeMapping(rawPcrs);
        pcr0 = _getPCR0(pcrs);

        certPubKey = _verifyCerts(certificate, cabundle);

        return (certPubKey, enclavePubKey, pcr0);
    }

    /// @notice PCR0 should be first in the array but this ensures it will find it otherwise
    function _getPCR0(bytes[2][] memory pcrs) internal pure returns (bytes memory) {
        require(pcrs.length != 0, "no pcr specified");
        require(pcrs.length <= 32, "only 32 pcrs allowed");

        for(uint256 i=0; i < pcrs.length; i++) {
            if (uint8(bytes1(pcrs[i][0])) == uint8(0)) {
                return pcrs[i][0];
            }
        }

        require(false, "failed to find pcr0");
    }
}
