package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robdimsdale/wl"
	wl_logger "github.com/robdimsdale/wl/logger"
	wl_oauth "github.com/robdimsdale/wl/oauth"
)

func main() {
	client := wl_oauth.NewClient(
		os.Getenv("WL_ACCESS_TOKEN"),
		os.Getenv("WL_CLIENT_ID"),
		wl.APIURL,
		wl_logger.NewLogger(wl_logger.INFO),
	)

	root, err := client.Root()
	if err != nil {
		log.Fatal(err)
	}

	myID := root.UserID

	tasks, err := client.Tasks()
	if err != nil {
		log.Fatal(err)
	}

	filtered := []wl.Task{}

	for _, task := range tasks {
		// Remove completed tasks
		if task.Completed {
			continue
		}

		// Remove tasks assigned to someone else
		if task.AssigneeID != uint(0) && task.AssigneeID != myID {
			continue
		}

		// Include tasks assigned to me or starred
		if task.AssigneeID == myID || task.Starred {
			filtered = append(filtered, task)
			continue
		}

		// Include overdue tasks
		if !task.DueDate.IsZero() && time.Now().After(task.DueDate) {
			filtered = append(filtered, task)
		}

	}

	if len(filtered) == 0 {
		return
	}

	fmt.Printf("✅ %v\n", len(filtered))

	for _, task := range filtered {
		fmt.Println(task.Title)
	}
}
