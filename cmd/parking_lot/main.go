package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rahulkushwaha12/simple_parking_lot/internal/controllers/command"
	"github.com/rahulkushwaha12/simple_parking_lot/internal/services/parking"
)

func main(){
	controller := command.NewController(parking.NewService())
	if len(os.Args)>1 && os.Args[1] != ""{
		cmdFile, err := os.Open(os.Args[1])
		if err != nil{
			log.Fatal(err)
		}
		defer cmdFile.Close()
		cmdScanner := bufio.NewScanner(cmdFile)
		for cmdScanner.Scan(){
			if err := cmdScanner.Err(); err != nil{
				log.Fatal(err)
			}
			cmdInput := cmdScanner.Text()
			cmdInput = strings.TrimRight(cmdInput,"\n")
			if cmdInput != ""{
				if cmdInput=="exit"{
					os.Exit(1)
				}
				fmt.Println(controller.Execute(cmdInput))
			}
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for{
		cmdInput,_:=reader.ReadString('\n')
		cmdInput = strings.TrimRight(cmdInput,"\n")
		if cmdInput != ""{
			if cmdInput=="exit"{
				os.Exit(1)
			}
			fmt.Println(controller.Execute(cmdInput))
		}
	}

}
