// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"time"

	"github.com/cprior/slmbg/sunlightmap"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tickCmd represents the tick command
var tickCmd = &cobra.Command{
	Use:   "tick",
	Short: "generates a fresh png file at current system time",
	Long: `With ./slmbg tick a png file is created depicting
the world map in daylight and nighttime
at the current system time (which is hopefully "now" ;):

For the math behind this and the Android app visit http://slm.prdv.de/ .`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("tick called")

		slm := sunlightmap.NewStatic(viper.GetInt("width"), time.Now().Local())
		slm.DaylightImageFilename = viper.GetString("DaylightImageFilename")
		slm.NighttimeImageFilename = viper.GetString("NighttimeImageFilename")
		_ = sunlightmap.WriteStaticPng(&slm, viper.GetString("OutputImageFilename"))
	},
}

func init() {
	rootCmd.AddCommand(tickCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tickCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tickCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}