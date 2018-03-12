package main

import (
	"encoding/json"
)

var EventQueue = make(chan BatchEvent, 100)

// TODO Implement SQS Long Polling
func Poller() {
	str := `{
	"version": "0",
	"id": "c8f9c4b5-76e5-d76a-f980-7011e206042b",
	"detail-type": "Batch Job State Change",
	"source": "aws.batch",
	"account": "aws_account_id",
	"time": "2017-10-23T17:56:03Z",
	"region": "us-east-1",
	"resources": [
		"arn:aws:batch:us-east-1:aws_account_id:job/4c7599ae-0a82-49aa-ba5a-4727fcce14a8"
	],
	"detail": {
		"jobName": "event-test",
		"jobId": "4c7599ae-0a82-49aa-ba5a-4727fcce14a8",
		"jobQueue": "arn:aws:batch:us-east-1:aws_account_id:job-queue/HighPriority",
		"status": "RUNNABLE",
		"attempts": [],
		"createdAt": 1508781340401,
		"retryStrategy": {
			"attempts": 1
		},
		"dependsOn": [],
		"jobDefinition": "arn:aws:batch:us-east-1:aws_account_id:job-definition/first-run-job-definition:1",
		"parameters": {},
		"container": {
			"image": "busybox",
			"vcpus": 2,
			"memory": 2000,
			"command": [
				"echo",
				"'hello world'"
			],
			"volumes": [],
			"environment": [],
			"mountPoints": [],
			"ulimits": []
		}
	}}`
	event := BatchEvent{}
	json.Unmarshal([]byte(str), event)
	EventQueue <- event
	EventQueue <- event
	EventQueue <- event
}
