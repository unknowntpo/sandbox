package main

// mux parse args and call corresponding handler according to args[0]
func mux(args []string) {
	// How to register and add handler like http.DefaultServeMux ?
	if args[0] == "help" {
		help()
	}
}

func help() {
	fmt.Println("help handler called")
}
