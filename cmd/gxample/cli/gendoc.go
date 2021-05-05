package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GendocCmd represents the gendoc command
var GendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "Generate project documentation",
	Run: func(cmd *cobra.Command, args []string) {
		genRoutesDoc()
	},
}

func genRoutesDoc() {
	fmt.Print("generating routes swagger file: ")
	// swagger := docgen.JSONRoutesDoc(router)
	// if err := ioutil.WriteFile("api/routes.json", []byte(swagger), 0644); err != nil {
	// logging.Logger.Fatalf("unable to generate documents, %s", err)
	//	panic(err)
	// }
	fmt.Println("Ok")
}
