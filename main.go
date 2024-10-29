package main

import(
	"gerrit.o-ran-sc.org/r/aiml-fw/apm/monitoring-agent/pkg/api/commons/logger"	
)

func main(){
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")
}