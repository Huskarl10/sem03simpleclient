package main

import (
	"log"
	"net"
	"os"

	"github.com/Huskarl10/is105sem03/mycrypt"
)

func main() {
    conn, err := net.Dial("tcp", "172.17.0.2:8000")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("os.Args[1] = ", os.Args[1])

    // Encrypt the message
    kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALFSEM03, 4)

    , err = conn.Write([]byte(string(kryptertMelding)))
    if err != nil {
        log.Fatal(err)
    }
    buf := make([]byte, 1024)
    n, err := conn.Read(buf)
    if err != nil {
        log.Fatal(err)
    }

    // Decrypt the response from the server
    deCrypt := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
    log.Println("Dekrypter melding: ", string(deCrypt))

    response := string(deCrypt)
    log.Printf("reply from proxy: %s", response)

}
