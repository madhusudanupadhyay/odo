package cmd

import (
	"fmt"
	"github.com/redhat-developer/odo/pkg/catalog"
	"github.com/spf13/cobra"
)

var catalogCmd = &cobra.Command{
	Use:   "catalog [options]",
	Short: "Catalog related operations",
	Long:  "Catalog related operations",
	Example: fmt.Sprintf("%s\n%s",
		catalogListCmd.Example,
		catalogSearchCmd.Example),
}

var catalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available component types.",
	Long:  "List all available component types from OpenShift's Image Builder.",
	Example: `  # Get the supported components
  odo catalog list

  # Search for a supported component
  odo search nodejs
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		client := getOcClient()
		catalogList, err := catalog.List(client)
		checkError(err, "unable to list components")
		switch len(catalogList) {
		case 0:
			fmt.Printf("No deployable components found\n")
		default:
			fmt.Println("The following components can be deployed:")
			for _, component := range catalogList {
				fmt.Printf("- %v\n", component)
			}
		}
	},
}

var catalogSearchCmd = &cobra.Command{
	Use:   "search <component name>",
	Short: "Search component type in catalog",
	Long: `Search component type in catalog.

This searches for a partial match for the given search term in all the available
components.
`,
	Example: `  # Search for a component
  odo catalog search pyt
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := getOcClient()
		searchTerm := args[0]
		components, err := catalog.Search(client, searchTerm)
		checkError(err, "unable to search for components")

		switch len(components) {
		case 0:
			fmt.Printf("No component matched the query: %v\n", searchTerm)
		default:
			fmt.Println("The following components were found:")
			for _, component := range components {
				fmt.Printf("- %v\n", component)
			}
		}
	},
}

func init() {
	catalogCmd.AddCommand(catalogSearchCmd)
	catalogCmd.AddCommand(catalogListCmd)
	rootCmd.AddCommand(catalogCmd)
}
