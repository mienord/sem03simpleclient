package main

import (
	"net"
	"log"
	"os"
	"github.com/mienord/is105sem03/mycrypt"
)

func main() {
    conn, err := net.Dial("tcp", "172.17.0.2:8001")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("os.Args[1] = ", os.Args[1])

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 2)
   	log.Println("Kryptert melding: ", string(kryptertMelding))

 	_, err = conn.Write([]byte(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	response := string(buf[:n])
	log.Printf("reply from proxy: %s", response)
}
