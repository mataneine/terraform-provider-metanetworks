package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"terraform-provider-metanetworks/metanetworks"
)

func die(message string) {
	log.Fatal(message)
}

func main() {
	csvFile, err := os.Open("users.csv")
	if err != nil {
		die(err.Error())
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	client, err := metanetworks.NewClientFromConfig()
	if err != nil {
		die(err.Error())
	}

	userList := []metanetworks.User{}
	userList, err = client.ListUsers()
	if err != nil {
		die(err.Error())
	}

	emailList := make(map[string]bool, 0)

	for _, v := range userList {
		emailList[v.Email] = true
	}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			die(err.Error())
		}

		user := metanetworks.User{
			GivenName:  line[0],
			FamilyName: line[1],
			Email:      line[2],
			Enabled:    true,
		}

		tags := make(map[string]string)
		tags["group"] = line[3]
		tags["location"] = line[4]

		_, ok := emailList[user.Email]
		if !ok {
			fmt.Println("Create " + user.Email)
			newUser, err := client.CreateUser(&user)
			if err != nil {
				die(err.Error())
			}

			err = client.SetUserTags(newUser.ID, tags)
			if err != nil {
				die(err.Error())
			}
		}
	}
}
