package params

// UseSeedNetworkConfig uses the Seed beacon chain specific network config.
func UseSeedNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 29
	cfg.BootstrapNodes = []string{
		// boot node (admin.nodeInfo.enr)
		"enr:-J24QBSwf_3I8ihindmAN-UP4Zi0eqj3OLU4tGRKEnkzHlpgR4r0ODPLfpclUYwWQCgi1TxwS-igxLpPCxKFyJmHE-6GAYLqGB4wg2V0aMfGhJS9tcuAgmlkgnY0gmlwhAqWAEeJc2VjcDI1NmsxoQJ9UjwkZsqQpgTHeaMMoFyLUVM36QVdlrPYe2affIz_SoN0Y3CCdl-DdWRwgnZf",
	}
	OverrideBeaconNetworkConfig(cfg)
}

// SeedConfig defines the config for the Seed beacon chain testnet.
func SeedConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.MinGenesisTime = 1653318000
	cfg.GenesisDelay = 604800
	cfg.MinGenesisActiveValidatorCount = 1
	cfg.ConfigName = SeedName
	cfg.GenesisForkVersion = []byte{0x00, 0x00, 0x00, 0x00}
	cfg.SecondsPerETH1Block = 15
	cfg.DepositChainID = 10
	cfg.DepositNetworkID = 10
	cfg.AltairForkEpoch = 0
	cfg.AltairForkVersion = []byte{0x00, 0x00, 0x00, 0x00}
	cfg.BellatrixForkEpoch = 0
	cfg.BellatrixForkVersion = []byte{0x00, 0x00, 0x00, 0x00}
	cfg.TerminalTotalDifficulty = "10"
	cfg.DepositContractAddress = "0xa80791cfF2Dd5eEa1431359cd6Bb256e19f16E50"
	cfg.InitializeForkSchedule()
	return cfg
}
