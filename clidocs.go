package main

import (
  "fmt"
  "os"
  "encoding/json"

  "github.com/urfave/cli"
)

type Response struct {
  Keyword string
  ArgList string
}

func main() {
  app := cli.NewApp()

  app.Flags = []cli.Flag {
    cli.StringFlag{
      Name: "lang",
      Value: "python",
      Usage: "Specify language for syntax lookup",
    },
  }

  app.Action = func(c *cli.Context) error {
    //query := "print"
    if c.NArg() > 0 {
      //query = c.Args().Get(0)
    }
    if c.String("lang") == "python" {
      var res Response
      b := getJSON("")
      err := json.Unmarshal(b, &res)
      if (err != nil) {
        // deal with error
      }
      fmt.Println(res)
    } else {
      //fmt.Println("Hello", name)
    }
    return nil
  }

  app.Run(os.Args)
}
