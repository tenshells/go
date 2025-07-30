package main 

import (
	"fmt"
	"strings"
	"time"
)

type ActivityCounter struct {
	UniqueUsers int
	UserActivity map[string]map[string]int
}


func main() {
	logs := []string{
		"2023-06-18T10:00:00Z user1 login",
		"2023-06-18T10:05:00Z user2 login",
		"2023-06-18T10:10:00Z user1 upload",
		"2023-06-18T10:15:00Z user2 upload",
		"2023-06-18T10:20:00Z user1 logout"}
	
	fmt.Println("logs are ")
	for _, log := range(logs) {
		fmt.Printf("\t%s\n", log)
	}
	
	// only count activities after below timestamp
    	cutoffTimeStr := "2023-06-18T10:05:00Z"
	cutoffTime, err := time.Parse(time.RFC3339, cutoffTimeStr)
	if err != nil {
		fmt.Println("Please make sure cutoffTimeStr has a valid, RFC3339 standard format")
		panic(err)
	}

	fmt.Printf("\nOnly counting logs after %s\n", cutoffTimeStr)

	var activityCounter ActivityCounter
	activityCounter.UserActivity = make(map[string]map[string]int, 0)
	
	for logNumber, userActivity := range(logs) {
		slice := strings.Split(userActivity, " ")
		timestamp, user, activity := slice[0], slice[1], slice[2]
		
		activityTime, err := time.Parse(time.RFC3339, timestamp)
		if err != nil {
			fmt.Printf("invalid timestamp %s at log number : %v", timestamp, logNumber)
			continue
		}
		if activityTime.Before(cutoffTime) {
			continue
		}

		fmt.Println("\t", timestamp, user, activity)
		_, ok := activityCounter.UserActivity[user]
		if !ok {
			activityCounter.UserActivity[user] = make(map[string]int)
		}
		activityCounter.UserActivity[user][activity] += 1
	}
	activityCounter.UniqueUsers = len(activityCounter.UserActivity)
	fmt.Printf("\nActivity Counter: \n%+v\n", activityCounter)
}
