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

func main() {

    // on this boot we show exactly how it would be using non-embedded structs but all separated and making type of each one of them
    // all initialization is more friendly, although having nesting between structs is much better if we have built structs ie
    // structs being declared directly inside structs what they call embedded.

    // Example embedded strcut would
    // type my struct {
    //    M struct {
    //       Name string
    //   }
    // }

    j := acl{
        User: crud{Create: true, Retrieve: false, Update: false, Delete: true},
        Acceleration: acceleration{
            Participant:    crudWithDetail{Detail: true, crud: crud{Create: false, Retrieve: false, Update: true, Delete: false}},
            Challenge:      crudWithDetail{Detail: false, crud: crud{Create: true, Retrieve: true, Update: true, Delete: false}},
            crudWithDetail: crudWithDetail{Detail: false, crud: crud{Create: true, Retrieve: false, Update: true, Delete: false}}},
    }

    m, err := json.Marshal(j)
    if err != nil {
        log.Println(err)
    }

    fmt.Println(string(m))

    // improving output for json format viewing
    // var prettyJSON bytes.Buffer
    // err = json.Indent(&prettyJSON, m, "", "\t")
    // if err != nil {
    //     log.Println("JSON parse error: ", err)
    // }
    // fmt.Println(string(prettyJSON.Bytes()))

    // improving output for json format viewing
    json, err := json.MarshalIndent(j, "", "\t")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(json))
}
