// Copyright 2017 The Zang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package examples

import (
	"fmt"

	zang "github.com/zang-cloud/zang-go"
)

func GetAccounts() error {
	client, err := zang.NewClient()
	response, err := client.GetAccounts(map[string]string{
		"Page":     "1",
		"PageSize": "10",
	})

	if err != nil {
		return fmt.Errorf("Unable to retrieve accounts due to: %s", err)
	}

	accounts := response.GetSubResource("accounts")

	for i := range accounts {
		account := accounts[i]
		fmt.Printf("\nFound account: %+v\n", account)
	}

	return nil
}

func GetAccount() error {
	client, err := zang.NewClient()
	response, err := client.GetAccount()

	if err != nil {
		return fmt.Errorf("Not able to discover account due to: %s", err)
	}

	fmt.Printf("Discovered account: %v", response.Resource)

	return nil
}

func UpdateAccount() error {
	client, err := zang.NewClient()
	response, err := client.UpdateAccount(map[string]string{
		"FriendlyName": "Hello from example",
	})

	if err != nil {
		return fmt.Errorf("Unable to retrieve accounts due to: %s", err)
	}

	fmt.Printf("Updated account: %v", response.Resource)
	return nil
}
