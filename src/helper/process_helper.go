package helper

import "log"

// Process call action
func Process(description string, f func()) {
	log.Printf("%s start", description)
	f()
	log.Printf("%s end", description)
}
