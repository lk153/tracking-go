package logger

import (
	"log"

	"github.com/getsentry/sentry-go"
)

//InitLogger ...
func InitLogger() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://4335485db4a144dba3e042f263034495@o532624.ingest.sentry.io/5651682",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
