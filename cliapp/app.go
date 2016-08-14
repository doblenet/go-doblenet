package cliapp

import (
	"errors"
	"path/filepath"
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

type AppInfo struct {
	Name		string
	Version		string
	GitTag		string
	Copyright	string
	Email		string
}


type App struct {
	Name		string
	info		AppInfo
	cobra		*cobra.Command
	authors		[]Author
	metadata	map[string]interface{}
	init		bool
}





func New(i AppInfo) *App {
	
	var n string
	
	if n = i.Name; ""==n {
		n = filepath.Base(os.Args[0])
	}
	
	return &App{
		init:	true,
		Name:	n,
		info:	i,
	}
}

func (a *App) SetCommander(c *cobra.Command) error {
	if !a.init {
		return errors.New("Application object not initialized properly")
	}
	a.cobra = c
	return setupCobra(a,c)
}

func (a *App) SetMetadata(mm map[string]interface{}) error {
	if !a.init {
		return errors.New("Application object not initialized properly")
	}
	a.metadata = mm
	return nil
}

func (x *App) SetAuthors(aa ...Author) {
	
	if nil == x.authors {
		x.authors = make([]Author,0,1)
	}
	
	for _,a := range aa {
		x.authors = append(x.authors,a)
	}
}

func (a *App) Run() {	// NORETURN
	
	if !a.init {
		fmt.Println("FATAL: Application object not initialized properly")
		os.Exit(127)
	}
	
	r := os.Args
	c := a.cobra
	
	if 2 == len(r) {
		switch r[1] {
			case "-V":	fallthrough
			case "-version": fallthrough
			case "--version":
				printVersionInfo(a,os.Stdout)
				os.Exit(0)
			case "-h":
				c.HelpFunc()(a.cobra, []string{})
				os.Exit(0)
		}
	}
	
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

// Designed to integrate with spf13/cobra
func setupCobra(a *App, c *cobra.Command) error {
	
	if nil == c {
		return errors.New("Invalid 'commander' given)")
	}
		
	c.AddCommand(&cobra.Command{
		Use: "version",
		Short: "display version info",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
// 			c.HelpFunc()(c,[]string{})
			printExtendedInfo(a, os.Stdout)
			os.Exit(0)
		},
	})
	return nil
}
