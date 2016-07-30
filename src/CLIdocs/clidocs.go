package main

import (
  "fmt"
  "os"
  "encoding/json"
  //"reflect"

  "github.com/urfave/cli"
)

type Response struct {
  Name string
  DataType string
  Declaration string
  Description string
  ExampleUsage string
}

// Given a Response object, cleanly prints it to the command line
func cleanPrint(r Response) {
  fmt.Printf("\n %s, %s, %s\n", r.Name, r.DataType, r.Declaration)
  fmt.Printf("\n %s\n", r.Description)
  fmt.Printf("\n %s\n", r.ExampleUsage)
  fmt.Printf("\n")
}

func main() {
  app := cli.NewApp()
  app.Name = "CLIdocs"
  app.Usage = "Command line documentation for your favorite languages!"
  app.Version = "1.0.0"

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang",
      Value: "python",
      Usage: "Specify language for syntax lookup",
    },
  }

  app.Action = func(c *cli.Context) error {
    query := "print"
    if c.NArg() > 0 {
      query = c.Args().Get(0)
    }
    var url string = c.String("lang") + "/" + query
    var res Response
    content := getJSON(url)
    err := json.Unmarshal(content, &res)
    if (err != nil) {
      // deal with error
    }
    //fmt.Printf("%+v\n ", res)
    cleanPrint(res)
    return nil
  }

  app.Run(os.Args)
}
