# TextMagic Go SDK

This library provides you with an easy solution to send SMS and receive replies by integrating TextMagic SMS Gateway to your Go application.

## What is TextMagic?
TextMagic's application programming interface (API) provides the communication link between your application and TextMagic’s SMS Gateway, allowing you to send and receive text messages and to check the delivery status of text messages you’ve already sent.

For detailed documentation and more examples, please visit [http://docs.textmagictesting.com/](http://docs.textmagictesting.com/)

## Installation

```
go get -u github.com/imissyouso/textmagic-rest-go
```

## Usage Example

```go
package main

import (
    "context"
    "fmt"
    "github.com/antihax/optional"
    tm "github.com/imissyouso/textmagic-rest-go"
    "log"
)

func main() {

    cfg := tm.NewConfiguration()
    cfg.BasePath = "https://rest.textmagic.com"
    client := tm.NewAPIClient(cfg)

    auth := context.WithValue(context.Background(), tm.ContextBasicAuth, tm.BasicAuth{
        UserName: "YOUR_USERNAME",
        Password: "YOUR_PASSWORD",
    })

    // Simple ping request example
    pingResponse, _, err := client.TextMagicApi.Ping(auth)

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println(pingResponse.Ping)
    }

    // Send a new message request example
    sendMessageResponse, _, err := client.TextMagicApi.SendMessage(auth, tm.SendMessageInputObject{
        Text: "I love TextMagic!",
        Phones: "+19998887766",
    }, &tm.SendMessageOpts{})

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println(sendMessageResponse.Id)
    }

    // Get all outgoing messages request example
    getAllOutboundMessageResponse, _, err := client.TextMagicApi.GetAllOutboundMessages(auth, &tm.GetAllOutboundMessagesOpts{
        Page: optional.NewInt32(1),
        Limit: optional.NewInt32(250),
    })

    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println(getAllOutboundMessageResponse.Resources[0].Id)
    }
}
```

## Limitations
Due https://github.com/swagger-api/swagger-codegen/issues/7311 issue current version of Go SDK does not support any file uploading API calls.

## License
The library is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).