/*
Copyright © 2024 49pctber

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/49pctber/upgen"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "upgen",
	Short: "The Universal Password Generator",
	Long:  `The Universal Password Generator can generate passwords and passphrases based on the requirements specified by the given flags.`,
	Run: func(cmd *cobra.Command, args []string) {

		reqs := make(map[int]interface{}, 0)
		if include, err := cmd.Flags().GetBool("b32"); include && err == nil {
			reqs[upgen.CharacterBase32] = nil
		}
		if include, err := cmd.Flags().GetBool("b64"); include && err == nil {
			reqs[upgen.CharacterBase64] = nil
		}
		if include, err := cmd.Flags().GetBool("lower"); include && err == nil {
			reqs[upgen.CharacterLowercase] = nil
		}
		if include, err := cmd.Flags().GetBool("upper"); include && err == nil {
			reqs[upgen.CharacterUppercase] = nil
		}
		if include, err := cmd.Flags().GetBool("num"); include && err == nil {
			reqs[upgen.CharacterNumeric] = nil
		}
		if include, err := cmd.Flags().GetBool("alphanum"); include && err == nil {
			reqs[upgen.CharacterAlphanumeric] = nil
		}
		if include, err := cmd.Flags().GetBool("special"); include && err == nil {
			reqs[upgen.CharacterSpecial] = nil
		}
		if include, err := cmd.Flags().GetBool("letter"); include && err == nil {
			reqs[upgen.CharacterLetter] = nil
		}
		if include, err := cmd.Flags().GetBool("hex"); include && err == nil {
			reqs[upgen.CharacterHex] = nil
		}
		if include, err := cmd.Flags().GetBool("bip39"); include && err == nil {
			reqs[upgen.WordBIP39] = nil
		}

		if len(reqs) == 0 {
			fmt.Println("No character sets specified. Using base32 characters.")
			reqs[upgen.CharacterBase32] = nil
		}

		n, err := cmd.Flags().GetInt("number")
		if err != nil {
			panic(err)
		}
		if n < 1 {
			log.Fatal("cannot produce a nonpositive number of passwords")
		}

		min_entropy, err := cmd.Flags().GetInt("entropy")
		if err != nil {
			panic(err)
		}

		if n > 1 {
			fmt.Printf("Passwords with >=%d bits of entropy:\n", min_entropy)
		} else {
			fmt.Printf("Password with >=%d bits of entropy:\n", min_entropy)
		}

		// generate passwords
		for i := 0; i < n; i++ {
			password := upgen.GetPassword(reqs, min_entropy)
			fmt.Printf("%s\n", password)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("entropy", "e", 100, "Minimum password entropy")
	rootCmd.Flags().IntP("number", "n", 1, "Number of passwords to generate")
	rootCmd.Flags().Bool("b32", false, "Include base32 characters")
	rootCmd.Flags().Bool("b64", false, "Include base64 characters")
	rootCmd.Flags().Bool("lower", false, "Include lowercase characters")
	rootCmd.Flags().Bool("upper", false, "Include uppercase characters")
	rootCmd.Flags().Bool("num", false, "Include numeric characters")
	rootCmd.Flags().Bool("alphanum", false, "Include alphanumeric characters")
	rootCmd.Flags().Bool("special", false, "Include special characters")
	rootCmd.Flags().Bool("letter", false, "Include letter characters")
	rootCmd.Flags().Bool("hex", false, "Include hex characters")
	rootCmd.Flags().Bool("bip39", false, "Include BIP-39 words")
}
