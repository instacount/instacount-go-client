# instacount-go-client
An Instacount client for GO

#Usage

```go get github.com/instacount/instacount-go-client```

#Example
```
import (
	instacount "github.com/instacount/instacount-go-client/client"
)

var client *instacount.Client

func initInstacount() {
	var appId string
	var apiKey string
	if (appengine.IsDevAppServer()) {
		appId = os.Getenv("INSTACOUNT_APP_ID__DEV")
		apiKey = os.Getenv("INSTACOUNT_API_KEY__DEV").
	} else {
		appId = os.Getenv("INSTACOUNT_APP_ID__PROD")
		apiKey = os.Getenv("INSTACOUNT_API_KEY__PROD").
	}
	client = instacount.NewClient(appId, apiKey)
}

func getCounter(counterName string) (*instacount.ShardedCounter) {
  if r, e, err := client.GetShardedCounter(counterName)); err != nil {
  		fmt.Printf("Unable to communicate with Instacount")
  		panic(err)
  	} else {
  		if e != nil {
  			fmt.Printf("Instacount responded with Errors: %v", e)
  		} else {
  			fmt.Printf("Instacount Response: %v", r)
  			return r
  		}
  	}
}
```
