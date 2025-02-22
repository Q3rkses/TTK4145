package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"time"
)

var (
	backup_flag string = "--backup"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		MakeBackup()
		TakeOver(0)
	} else if args[1] == backup_flag {
		RunBackup()
	}
}

func Remember(memory *uint32, message []byte) {
	*memory = binary.LittleEndian.Uint32(message)
	fmt.Println("Remembered: ", *memory)
}

func MakeBackup() {
	exec.Command("gnome-terminal", "--", "go", "run", "main.go", backup_flag).Run()
}

func SendMessage(message uint32) {
	address := net.UDPAddr{
		IP:   nil,
		Port: 44273,
	}

	conn, err := net.DialUDP("udp", nil, &address)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, message)
	conn.Write(buffer)
}

func RunBackup() {
	addr := net.UDPAddr{
		IP:   nil,
		Port: 44273,
	}

	var conn *net.UDPConn
	var err error

	for {
		conn, err = net.ListenUDP("udp", &addr)
		if err == nil {
			break
		}

		fmt.Println("Couldn't listen ... trying again")
		time.Sleep(time.Second * 1)
	}

	var memory uint32
	buffer := make([]byte, 4)

	fmt.Println("Big brother is listening...")

	for {

		// Set a deadline for reading. Read operation will fail if no data arrives after deadline.
		err := conn.SetReadDeadline(time.Now().Add(time.Second * 3))
		if err != nil {
			log.Print(err)
			continue
		}

		// Read incoming data
		n, _, err := conn.ReadFromUDP(buffer) // n is length of incoming data
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				log.Print(err)
				continue
			}
			log.Print("It's oddly quiet, i need to check on Big Brother...\n")
			break
		}

		Remember(&memory, buffer[:n])
		time.Sleep(time.Second * 1)
	}

	conn.Close()
	MakeBackup()
	TakeOver(memory)
}

func TakeOver(memory uint32) {
	fmt.Println("I am the primary")
	var i uint32 = memory

	for {
		fmt.Printf("I hope Big Brother doesn't listen to my secret message: %v, this time\n", i)
		SendMessage(i)
		i++

		time.Sleep(time.Second * 1)
	}
}
