// https://www.elephdev.com/golang/285.html?ref=addtabs&lang=en
// Modified by: prr azul software
// Date 3/7/2023
// copyright 2023 prr, azul software
//


package main

import (
	"os"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	//"strconv"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func main() {

	numArg:= len(os.Args)
	portStr := ":10600"
	if numArg > 2 {log.Fatalf("too many cli arguments\n")}
	if numArg == 2 {portStr = ":" + os.Args[1]}

	log.Printf("starting Server, listening on port %s\n", portStr)

	http.ListenAndServe(portStr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Client connected")
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Printf("Error starting socket server: %v\n", err)
		}

		go func() {
			defer conn.Close()
			for i:=0; i< 20; i++ {
				msg, op, err := wsutil.ReadClientData(conn)
				if err != nil {
					log.Printf("Error receiving data: %v\n", err)
					log.Printf("Client disconnected")
					return
				}
				log.Printf("=> Client: %s\n", string(msg))
//				randomNumber := strconv.Itoa(rand.Intn(100))
				sendMsg := fmt.Sprintf("hello from server: %d\n", rand.Intn(100))
				err = wsutil.WriteServerMessage(conn, op, []byte(sendMsg))
				if err != nil {
					log.Printf("Error sending data: %v\n", err)
					log.Printf("Client disconnected")
					return
				}
				log.Printf("Server sent msg[%d]: %s\n", i+1, sendMsg)
			}

		}()
	}))
}
