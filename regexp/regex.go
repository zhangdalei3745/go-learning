package main

import (
	"github.com/sirupsen/logrus"
	regexp "regexp"
)

func main()  {
	findStringSubmatch()
}

func findStringSubmatch()  {
	//jmsfServicePrefix := "jmesh/services/meshgroup/%s/namespace/"
	jmsfServicePattern := "jmesh/meshgroup/(.*)/namespace/(.*)/service/(.*)/application"
	k1 := "jmesh/meshgroup/mesh01/namespace/ns1/service/demo1"
	k2 := "jmesh/meshgroup/mesh02/namespace/ns2/service/demo2/test123"
	k3 := "jmesh/meshgroup/mesh03/namespace/ns3/service/demo3/application"
	serviceRegex, err := regexp.Compile(jmsfServicePattern)
	if err != nil {
		logrus.Fatal(err)
	}
	items := serviceRegex.FindStringSubmatch(k1)
	logrus.Info(items)
	items2 := serviceRegex.FindStringSubmatch(k2)
	logrus.Info(items2)
	items3 := serviceRegex.FindStringSubmatch(k3)
	logrus.Info(items3)
	r1, err := regexp.MatchString(jmsfServicePattern, k1)
	logrus.Info(r1)
	r2, err := regexp.MatchString(jmsfServicePattern, k2)
	logrus.Info(r2)
	r3, err := regexp.MatchString(jmsfServicePattern, k3)
	logrus.Info(r3)

}
