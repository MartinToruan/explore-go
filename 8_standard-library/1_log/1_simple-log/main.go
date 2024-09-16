package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("Sample Print Log")

	// Fatal Log will trigger a call to os.Exit(1)
	log.Fatalln("Sampe Fatal Log")

	// Panic Log will trigger a call to panic()
	log.Panicln("This is a panic Log")
}
