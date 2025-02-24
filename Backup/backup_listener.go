package main

// import (
// 	"fmt"
// )

var (
	ipaddress string = "10.100.23.23"
	password  string = "Sanntid15"
	// customFlag        = "--Backup"
)

// type Backup struct {
// 	targetIP string
// 	password string
// 	AliveLock *sync.Mutex
// 	BackupView []elevalgo.Elevator
// }

// func checkExistence(targetIP string, password string) {

// 	// Check if the backup is already running, and if not start it up
// 	exitCodes, err := AlreadyRunning("backup", "localhost", "password")
// 	if err != nil {
// 		fmt.Printf("Failed to query: %v", err)
// 	} else if exitCodes == 0 {
// 		fmt.Println("Backup is already running, (remembering data ?)")
// 	} else {
// 		createBackupListener(targetIP, password)
// 	}

// 	fmt.Println("Creating backup")
// }

// func createBackupListener(targetIP string, password string) {
// 	// Create a backup listener
// 	// This will listen for a message from the primary server
// }

// func backupFunctionality() {
// 	// if the listener detects that main is ded, revive it
// }

func main() {
	Revive("localhost", password)
}
