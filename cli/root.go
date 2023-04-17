// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/sourcenetwork/defradb/config"
	"github.com/sourcenetwork/defradb/errors"
)

var rootCmd = &cobra.Command{
	Use:   "defradb",
	Short: "DefraDB Edge Database",
	Long: `DefraDB is the edge database to power the user-centric future.

Start a database node, issue a request to a local or remote node, and much more.

DefraDB is released under the BSL license, (c) 2022 Democratized Data Foundation.
See https://docs.source.network/BSL.txt for more information.
`,
	// Runs on subcommands before their Run function, to handle configuration and top-level flags.
	// Loads the rootDir containing the configuration file, otherwise warn about it and load a default configuration.
	// This allows some subcommands (`init`, `start`) to override the PreRun to create a rootDir by default.
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		if cfg.ConfigFileExists() {
			if err := cfg.LoadWithRootdir(true); err != nil {
				return errors.Wrap("failed to load config", err)
			}
			log.FeedbackInfo(cmd.Context(), fmt.Sprintf("Configuration loaded from DefraDB directory %v", cfg.Rootdir))
		} else {
			if err := cfg.LoadWithRootdir(false); err != nil {
				return errors.Wrap("failed to load config", err)
			}
			log.FeedbackInfo(cmd.Context(), "Using default configuration")
		}
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(
		&cfg.Rootdir, "rootdir", config.DefaultRootDir(),
		"Directory for data and configuration to use",
	)

	rootCmd.PersistentFlags().String(
		"loglevel", cfg.Log.Level,
		"Log level to use. Options are debug, info, error, fatal",
	)
	err := cfg.BindFlag("log.level", rootCmd.PersistentFlags().Lookup("loglevel"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind log.loglevel", err)
	}

	rootCmd.PersistentFlags().StringArray(
		"logger", []string{},
		"Override logger parameters. Usage: --logger <name>,level=<level>,output=<output>,...",
	)
	err = cfg.BindFlag("log.logger", rootCmd.PersistentFlags().Lookup("logger"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind log.logger", err)
	}

	rootCmd.PersistentFlags().String(
		"logoutput", cfg.Log.Output,
		"Log output path",
	)
	err = cfg.BindFlag("log.output", rootCmd.PersistentFlags().Lookup("logoutput"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind log.output", err)
	}

	rootCmd.PersistentFlags().String(
		"logformat", cfg.Log.Format,
		"Log format to use. Options are csv, json",
	)
	err = cfg.BindFlag("log.format", rootCmd.PersistentFlags().Lookup("logformat"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind log.format", err)
	}

	rootCmd.PersistentFlags().Bool(
		"logtrace", cfg.Log.Stacktrace,
		"Include stacktrace in error and fatal logs",
	)
	err = cfg.BindFlag("log.stacktrace", rootCmd.PersistentFlags().Lookup("logtrace"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind log.stacktrace", err)
	}

	rootCmd.PersistentFlags().Bool(
		"lognocolor", cfg.Log.NoColor,
		"Disable colored log output",
	)
	err = cfg.BindFlag("log.nocolor", rootCmd.PersistentFlags().Lookup("lognocolor"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind log.nocolor", err)
	}

	rootCmd.PersistentFlags().String(
		"url", cfg.API.Address,
		"URL of HTTP endpoint to listen on or connect to",
	)
	err = cfg.BindFlag("api.address", rootCmd.PersistentFlags().Lookup("url"))
	if err != nil {
		log.FeedbackFatalE(context.Background(), "Could not bind api.address", err)
	}
}