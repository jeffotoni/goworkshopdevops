// Go in action
// @jeffotoni
// 2019-01-29

package main

import (
    "fmt"
    "reflect"
)

type VehicleInfo struct {
    // ID         bson.ObjectId `bson:"_id,omitempty"`
    VehicleId string `bson:"编号"`
    Date      string `bson:"日期"`
    Type      string `bson:"类型"`
    Brand     string `bson:"型号"`
    Color     string `bson:"颜色"`
}

func main() {
    vinfo := VehicleInfo{
        VehicleId: "123456",
        Date:      "20140101",
        Type:      "Truck",
        Brand:     "Ford",
        Color:     "White",
    }

    vt := reflect.TypeOf(vinfo)
    vv := reflect.ValueOf(vinfo)
    for i := 0; i < vt.NumField(); i++ {
        f := vt.Field(i)
        chKey := f.Tag.Get("bson")
        fmt.Printf("%q => %q, ", chKey, vv.FieldByName(f.Name).String())
    }
}
