/*
Copyright © 2025 Tina Holly <info@tinaciousdesign.com>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tinacious/random/utils"
)

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Generate a random name",
	Long: `Generates a random name from a set of names (1000 female, 1000 male, 300 unisex).
When no gender is specified, the entire dataset will be sampled.`,
	Run: func(cmd *cobra.Command, args []string) {
		gender, err := cmd.Flags().GetString("gender")
		if err != nil {
			log.Fatal(err)
		}

		supportedGenders, err := utils.AllNameGenderKeys()
		if err != nil {
			log.Fatal(err)
		}
		if !slices.Contains(supportedGenders, gender) {
			log.Fatalf("supported genders include: %s", strings.Join(supportedGenders, ", "))
		}

		service := utils.NewPopulatedService()

		firstName, err := service.RandomFirstName(utils.NameGenderFromString(gender))
		if err != nil {
			log.Fatal(err)
		}
		lastName, err := service.RandomLastName()
		if err != nil {
			log.Fatal(err)
		}

		fullName := firstName + " " + lastName

		fmt.Println(fullName)
	},
}

func init() {

	// flags for gender
	docsGender, err := utils.AllNameGenderDescriptionsWithKeys()
	if err != nil {
		log.Fatal(err)
	}
	nameCmd.Flags().StringP("gender", "g", "a", fmt.Sprintf("Gender of the name. Omitting will sample all available names (optional) (options: %s)", strings.Join(docsGender, ", ")))

	rootCmd.AddCommand(nameCmd)
}
