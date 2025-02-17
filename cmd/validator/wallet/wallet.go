package wallet

import (
	"github.com/prysmaticlabs/prysm/v3/cmd"
	"github.com/prysmaticlabs/prysm/v3/cmd/validator/flags"
	"github.com/prysmaticlabs/prysm/v3/config/features"
	"github.com/prysmaticlabs/prysm/v3/runtime/tos"
	"github.com/prysmaticlabs/prysm/v3/validator/accounts"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var log = logrus.WithField("prefix", "wallet")

// Commands for wallets for Prysm validators.
var Commands = &cli.Command{
	Name:     "wallet",
	Category: "wallet",
	Usage:    "defines commands for interacting with Ethereum validator wallets",
	Subcommands: []*cli.Command{
		{
			Name: "create",
			Usage: "creates a new wallet with a desired type of keymanager: " +
				"either on-disk (imported), derived, or using remote credentials",
			Flags: cmd.WrapFlags([]cli.Flag{
				flags.WalletDirFlag,
				flags.KeymanagerKindFlag,
				flags.GrpcRemoteAddressFlag,
				flags.DisableRemoteSignerTlsFlag,
				flags.RemoteSignerCertPathFlag,
				flags.RemoteSignerKeyPathFlag,
				flags.RemoteSignerCACertPathFlag,
				flags.WalletPasswordFileFlag,
				flags.Mnemonic25thWordFileFlag,
				flags.SkipMnemonic25thWordCheckFlag,
				features.Mainnet,
				features.PraterTestnet,
				features.RopstenTestnet,
				features.SepoliaTestnet,
				features.SeedTestnet,
				cmd.AcceptTosFlag,
			}),
			Before: func(cliCtx *cli.Context) error {
				if err := cmd.LoadFlagsFromConfig(cliCtx, cliCtx.Command.Flags); err != nil {
					return err
				}
				return tos.VerifyTosAcceptedOrPrompt(cliCtx)
			},
			Action: func(cliCtx *cli.Context) error {
				if err := features.ConfigureValidator(cliCtx); err != nil {
					return err
				}
				if _, err := accounts.CreateAndSaveWalletCli(cliCtx); err != nil {
					log.WithError(err).Fatal("Could not create a wallet")
				}
				return nil
			},
		},
		{
			Name:  "edit-config",
			Usage: "edits a wallet configuration options, such as gRPC connection credentials and TLS certificates",
			Flags: cmd.WrapFlags([]cli.Flag{
				flags.WalletDirFlag,
				flags.WalletPasswordFileFlag,
				flags.GrpcRemoteAddressFlag,
				flags.DisableRemoteSignerTlsFlag,
				flags.RemoteSignerCertPathFlag,
				flags.RemoteSignerKeyPathFlag,
				flags.RemoteSignerCACertPathFlag,
				features.Mainnet,
				features.PraterTestnet,
				features.RopstenTestnet,
				features.SepoliaTestnet,
				features.SeedTestnet,
				cmd.AcceptTosFlag,
			}),
			Before: func(cliCtx *cli.Context) error {
				if err := cmd.LoadFlagsFromConfig(cliCtx, cliCtx.Command.Flags); err != nil {
					return err
				}
				if err := tos.VerifyTosAcceptedOrPrompt(cliCtx); err != nil {
					return err
				}
				return features.ConfigureValidator(cliCtx)
			},
			Action: func(cliCtx *cli.Context) error {
				if err := remoteWalletEdit(cliCtx); err != nil {
					log.WithError(err).Fatal("Could not edit wallet configuration")
				}
				return nil
			},
		},
		{
			Name:  "recover",
			Usage: "uses a derived wallet seed recovery phase to recreate an existing HD wallet",
			Flags: cmd.WrapFlags([]cli.Flag{
				flags.WalletDirFlag,
				flags.MnemonicFileFlag,
				flags.WalletPasswordFileFlag,
				flags.NumAccountsFlag,
				flags.Mnemonic25thWordFileFlag,
				flags.SkipMnemonic25thWordCheckFlag,
				features.Mainnet,
				features.PraterTestnet,
				features.RopstenTestnet,
				features.SepoliaTestnet,
				features.SeedTestnet,
				cmd.AcceptTosFlag,
			}),
			Before: func(cliCtx *cli.Context) error {
				if err := cmd.LoadFlagsFromConfig(cliCtx, cliCtx.Command.Flags); err != nil {
					return err
				}
				if err := tos.VerifyTosAcceptedOrPrompt(cliCtx); err != nil {
					return err
				}
				return features.ConfigureBeaconChain(cliCtx)
			},
			Action: func(cliCtx *cli.Context) error {
				if err := walletRecover(cliCtx); err != nil {
					log.WithError(err).Fatal("Could not recover wallet")
				}
				return nil
			},
		},
	},
}
