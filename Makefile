.PHONY: bindings
bindings:
	forge clean && forge build --via-ir --extra-output-files abi
	mkdir -p bindings
	abigen --abi out/OutputOracle.sol/OutputOracle.abi.json --pkg bindings --type OutputOracle --out bindings/output_oracle.go
	abigen --abi out/Portal.sol/Portal.abi.json --pkg bindings --type Portal --out bindings/portal.go
	abigen --abi out/SystemConfigOwnable.sol/SystemConfigOwnable.abi.json --pkg bindings --type SystemConfigOwnable --out bindings/system_config_ownable.go
