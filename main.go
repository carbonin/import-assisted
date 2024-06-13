package main

import (
	"fmt"
	"time"

	"github.com/openshift/assisted-service/pkg/thread"
	"github.com/sirupsen/logrus"
)

func main() {
	t := thread.New(logrus.New(), "thread-test", time.Second*10, func() {
		fmt.Println("test assisted thread package")
	})
	t.Start()
	time.Sleep(time.Minute)
}
