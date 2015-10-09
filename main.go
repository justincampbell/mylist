package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/robdimsdale/wl"
	wl_logger "github.com/robdimsdale/wl/logger"
	wl_oauth "github.com/robdimsdale/wl/oauth"
)

var maxCacheAge, _ = time.ParseDuration("15m")

func main() {
	tmpDir := os.TempDir()
	cacheFile := path.Join(tmpDir, "mylist")

	if !isCacheStale(cacheFile) {
		cached, err := ioutil.ReadFile(cacheFile)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(cached))
		return
	}

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
		if task.DueDate != "" {
			dueDate, err := time.Parse("2006-01-02", task.DueDate)
			if err != nil {
				log.Fatal(err)
			}

			if time.Now().After(dueDate) {
				filtered = append(filtered, task)
			}
		}

	}

	var output bytes.Buffer

	if len(filtered) != 0 {
		output.WriteString(fmt.Sprintf("✅ %v\n", len(filtered)))

		for _, task := range filtered {
			output.WriteString(task.Title)
			output.WriteString("\n")
		}
	}

	err = ioutil.WriteFile(cacheFile, []byte(output.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(output.String())
}

func isCacheStale(cacheFile string) bool {
	stat, err := os.Stat(cacheFile)

	return os.IsNotExist(err) || time.Since(stat.ModTime()) > maxCacheAge
}
