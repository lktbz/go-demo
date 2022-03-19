package main

import (
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
)
//https://www.cnblogs.com/nickchen121/p/11517428.html
func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

}
