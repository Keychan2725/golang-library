package main

import (
	"flag"
	"fmt"
)

func main() {
	var hostname *string = flag.String("hostname", "localhost", "Your database")
	var username *string = flag.String("username", "root", "Your database username")
	var password *string = flag.String("password", "root", "Your database password")
	var port *int = flag.Int("port", 3000, "Your PORT")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Hostname", *hostname)
	fmt.Println("Port", *port)
}
