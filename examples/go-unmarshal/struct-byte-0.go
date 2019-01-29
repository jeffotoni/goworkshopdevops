// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "encoding/json"
    "fmt"
    "log"
)

type crud struct {
    Create   bool
    Retrieve bool
    Update   bool
    Delete   bool
}

type crudWithDetail struct {
    crud
    Detail bool
}

type acceleration struct {
    crudWithDetail
    Participant crudWithDetail
    Challenge   crudWithDetail
}

type acl struct {
    User         crud
    Acceleration acceleration
}

var j = []byte(`{
        "user": {
            "create":true,
            "retrieve":false,
            "update": false,
            "delete": true
        },
        "acceleration": {
            "detail": true,
            "create": true,
            "retrieve": false,
            "update": false,
            "delete": true,
            "participant": {
                "detail": false,
                "create": true,
                "retrieve": false,
                "update": false,
                "delete": true
            },
            "challenge": {
                "detail": false,
                "create": true,
                "retrieve": false,
                "update": false,
                "delete": true
            }
        }        
    }
`)

func main() {

    var x acl
    if err := json.Unmarshal(j, &x); err != nil {
        log.Fatal(err)
    }

    m, err := json.Marshal(x)
    if err != nil {
        log.Println(err)
    }

    fmt.Println(string(m))

    //fmt.Println(x)
    //fmt.Println(x.Acceleration.Participant.Detail)
}
