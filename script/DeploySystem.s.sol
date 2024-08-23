// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import { Deploy } from "@eth-optimism-bedrock/scripts/deploy/Deploy.s.sol";
import { DeployConfig } from "@eth-optimism-bedrock/scripts/deploy/DeployConfig.s.sol";
import { Types } from "@eth-optimism-bedrock/scripts/Types.sol";
import { ChainAssertions } from "@eth-optimism-bedrock/scripts/ChainAssertions.sol";
import { SystemConfig } from "@eth-optimism-bedrock/src/L1/SystemConfig.sol";
import { SuperchainConfig } from "@eth-optimism-bedrock/src/L1/SuperchainConfig.sol";
import { OptimismPortal } from "@eth-optimism-bedrock/src/L1/OptimismPortal.sol";
import { L2OutputOracle } from "@eth-optimism-bedrock/src/L1/L2OutputOracle.sol";
import { Portal } from "src/Portal.sol";
import { OutputOracle } from "src/OutputOracle.sol";
import { SystemConfigOwnable } from "src/SystemConfigOwnable.sol";
import { Constants } from "@eth-optimism-bedrock/src/libraries/Constants.sol";

import { console2 as console } from "forge-std/console2.sol";

contract DeploySystem is Deploy {
    function deploy() public {
        console.log("start of L1 Deploy!");
        deploySafe("SystemOwnerSafe");
        console.log("deployed Safe!");
        setupSuperchain2();
        console.log("set up superchain!");
        setupOpChain2();
        console.log("set up op chain!");
    }

    function setupSuperchain2() public {
        console.log("Setting up Superchain");

        // Deploy a new ProxyAdmin and AddressManager
        // This proxy will be used on the SuperchainConfig and ProtocolVersions contracts, as well as the contracts
        // in the OP Chain system.
        deployAddressManager();
        deployProxyAdmin();
        transferProxyAdminOwnership();

        // Deploy the SuperchainConfigProxy
        deployERC1967Proxy("SuperchainConfigProxy");
        deploySuperchainConfig();
        initializeSuperchainConfig();

        // Deploy the ProtocolVersionsProxy
        deployERC1967Proxy("ProtocolVersionsProxy");
        deployProtocolVersions();
        initializeProtocolVersions();
    }

    function setupOpChain2() public {
        console.log("Deploying OP Chain");

        // Ensure that the requisite contracts are deployed
        mustGetAddress("SuperchainConfigProxy");
        mustGetAddress("SystemOwnerSafe");
        mustGetAddress("AddressManager");
        mustGetAddress("ProxyAdmin");

        deployProxies2();
        deployImplementations2();
        initializeImplementations2();
    }

    function deployProxies2() public {
        console.log("Deploying proxies");

        deployERC1967Proxy("OptimismPortalProxy");
        deployERC1967Proxy("SystemConfigProxy");
        deployL1StandardBridgeProxy();
        deployL1CrossDomainMessengerProxy();
        deployERC1967Proxy("OptimismMintableERC20FactoryProxy");
        deployERC1967Proxy("L1ERC721BridgeProxy");

        deployERC1967Proxy("L2OutputOracleProxy");

        transferAddressManagerOwnership(); // to the ProxyAdmin

        // fake these, required for calls to `_proxies()`
        save("DisputeGameFactoryProxy", address(1));
        save("DelayedWETHProxy", address(1));
        save("AnchorStateRegistryProxy", address(1));
    }

    function deployImplementations2() public {
        console.log("Deploying implementations");
        deployL1CrossDomainMessenger();
        deployOptimismMintableERC20Factory();
        deploySystemConfigOwnable();
        deployL1StandardBridge();
        deployL1ERC721Bridge();
        deployPortal();
        deployOutputOracle();
    }

    function initializeImplementations2() public {
        console.log("Initializing implementations");
        initializePortal();
        initializeSystemConfigOwnable();
        initializeL1StandardBridge();
        initializeL1ERC721Bridge();
        initializeOptimismMintableERC20Factory();
        initializeL1CrossDomainMessenger();
        initializeOutputOracle();
    }

    function deploySystemConfigOwnable() public broadcast returns (address addr_) {
        console.log("Deploying SystemConfig implementation");
        addr_ = address(new SystemConfigOwnable{ salt: _implSalt() }(cfg.finalSystemOwner()));
        save("SystemConfig", addr_);
        console.log("SystemConfig deployed at %s", addr_);

        // Override the `SystemConfig` contract to the deployed implementation. This is necessary
        // to check the `SystemConfig` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.SystemConfig = addr_;
        ChainAssertions.checkSystemConfig({ _contracts: contracts, _cfg: cfg, _isProxy: false });
    }

    function deployPortal() public broadcast returns (address addr_) {
        console.log("Deploying OptimismPortal implementation");
        addr_ = address(new Portal{ salt: _implSalt() }());
        save("OptimismPortal", addr_);
        console.log("OptimismPortal deployed at %s", addr_);

        // Override the `OptimismPortal` contract to the deployed implementation. This is necessary
        // to check the `OptimismPortal` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.OptimismPortal = addr_;
        ChainAssertions.checkOptimismPortal({ _contracts: contracts, _cfg: cfg, _isProxy: false });
    }

    function deployOutputOracle() public broadcast returns (address addr_) {
        console.log("Deploying L2OutputOracle implementation");
        OutputOracle oracle = new OutputOracle{ salt: _implSalt() }();

        save("L2OutputOracle", address(oracle));
        console.log("L2OutputOracle deployed at %s", address(oracle));

        // Override the `L2OutputOracle` contract to the deployed implementation. This is necessary
        // to check the `L2OutputOracle` implementation alongside dependent contracts, which
        // are always proxies.
        Types.ContractSet memory contracts = _proxiesUnstrict();
        contracts.L2OutputOracle = address(oracle);
        checkL2OutputOracle({
            _contracts: contracts,
            _cfg: cfg,
            _isProxy: false
        });

        addr_ = address(oracle);
    }

    function initializeSystemConfigOwnable() public broadcast {
        console.log("Upgrading and initializing SystemConfig proxy");
        address systemConfigProxy = mustGetAddress("SystemConfigProxy");
        address systemConfig = mustGetAddress("SystemConfig");

        bytes32 batcherHash = bytes32(uint256(uint160(cfg.batchSenderAddress())));

        address customGasTokenAddress = Constants.ETHER;
        if (cfg.useCustomGasToken()) {
            customGasTokenAddress = cfg.customGasTokenAddress();
        }

        _upgradeAndCallViaSafe({
            _proxy: payable(systemConfigProxy),
            _implementation: systemConfig,
            _innerCallData: abi.encodeCall(
                SystemConfigOwnable.initialize,
                (
                    cfg.basefeeScalar(),
                    cfg.blobbasefeeScalar(),
                    batcherHash,
                    uint64(cfg.l2GenesisBlockGasLimit()),
                    cfg.p2pSequencerAddress(),
                    Constants.DEFAULT_RESOURCE_CONFIG(),
                    cfg.batchInboxAddress(),
                    SystemConfig.Addresses({
                        l1CrossDomainMessenger: mustGetAddress("L1CrossDomainMessengerProxy"),
                        l1ERC721Bridge: mustGetAddress("L1ERC721BridgeProxy"),
                        l1StandardBridge: mustGetAddress("L1StandardBridgeProxy"),
                        disputeGameFactory: mustGetAddress("DisputeGameFactoryProxy"),
                        optimismPortal: mustGetAddress("OptimismPortalProxy"),
                        optimismMintableERC20Factory: mustGetAddress("OptimismMintableERC20FactoryProxy"),
                        gasPayingToken: customGasTokenAddress
                    })
                )
            )
        });

        SystemConfig config = SystemConfig(systemConfigProxy);
        string memory version = config.version();
        console.log("SystemConfig version: %s", version);

        ChainAssertions.checkSystemConfig({ _contracts: _proxies(), _cfg: cfg, _isProxy: true });
    }

    function initializePortal() public broadcast {
        console.log("Upgrading and initializing OptimismPortal proxy");
        address optimismPortalProxy = mustGetAddress("OptimismPortalProxy");
        address optimismPortal = mustGetAddress("OptimismPortal");
        address l2OutputOracleProxy = mustGetAddress("L2OutputOracleProxy");
        address systemConfigProxy = mustGetAddress("SystemConfigProxy");
        address superchainConfigProxy = mustGetAddress("SuperchainConfigProxy");

        _upgradeAndCallViaSafe({
            _proxy: payable(optimismPortalProxy),
            _implementation: optimismPortal,
            _innerCallData: abi.encodeCall(
                OptimismPortal.initialize,
                (
                    L2OutputOracle(l2OutputOracleProxy),
                    SystemConfig(systemConfigProxy),
                    SuperchainConfig(superchainConfigProxy)
                )
            )
        });

        Portal portal = Portal(payable(optimismPortalProxy));
        string memory version = portal.version();
        console.log("OptimismPortal version: %s", version);

        ChainAssertions.checkOptimismPortal({ _contracts: _proxies(), _cfg: cfg, _isProxy: true });
    }

    function initializeOutputOracle() public broadcast {
        console.log("Upgrading and initializing L2OutputOracle proxy");
        address l2OutputOracleProxy = mustGetAddress("L2OutputOracleProxy");
        address l2OutputOracle = mustGetAddress("L2OutputOracle");

        _upgradeAndCallViaSafe({
            _proxy: payable(l2OutputOracleProxy),
            _implementation: l2OutputOracle,
            _innerCallData: abi.encodeCall(
                OutputOracle.initialize,
                (
                    cfg.l2OutputOracleProposer(),
                    1000 // TODO move to config
                )
            )
        });

        OutputOracle oracle = OutputOracle(l2OutputOracleProxy);
        string memory version = oracle.version();
        console.log("L2OutputOracle version: %s", version);

        checkL2OutputOracle({
            _contracts: _proxies(),
            _cfg: cfg,
            _isProxy: true
        });
    }

    function checkL2OutputOracle(
        Types.ContractSet memory _contracts,
        DeployConfig _cfg,
        bool _isProxy
    )
    internal
    view
    {
        console.log("Running chain assertions on the L2OutputOracle");
        OutputOracle oracle = OutputOracle(_contracts.L2OutputOracle);

        // Check that the contract is initialized
        ChainAssertions.assertSlotValueIsOne({ _contractAddress: address(oracle), _slot: 0, _offset: 0 });

        if (_isProxy) {
            require(oracle.proposer() == _cfg.l2OutputOracleProposer());
        } else {
            require(oracle.proposer() == address(0));
        }
    }
}
