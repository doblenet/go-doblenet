package main

import (
	cliapp "../"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {

	var RootCmd = &cobra.Command{
		Use:   "dncliapp-test",
		Short: "test short description",
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("run!")
			return nil
		},
	}

	app := cliapp.New(cliapp.AppInfo{Name: "dncliapp_test",
		Version: "0.1.0", GitTag: "abcd1234",
		Copyright: "(C)2016 DOBLENET Soluciones Tecnológicas", Email: "doblenet-go@xyz.labs.doblenet.com",
	})

	app.SetAuthors(
		cliapp.Author{Name: "Jane Smith", Email: "j.smith@nowhere.xyz"}, 
		cliapp.Author{Name: "John Doe", Email: "j.doe@nowhere.xyz"},
	)

	// app.SetMetadata(map[string]interface{}{"prop1": "one", "prop2": "two", "prop3": 42})

	app.SetCommander(RootCmd)

	app.Run()
}
