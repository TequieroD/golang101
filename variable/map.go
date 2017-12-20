package main

import "fmt"

func main() {
    elements := make(map[string]int)
    elements["aaa"] = 111
    fmt.Println(elements)
    fmt.Println(elements["aaa"])

    elements["bbb"] = 222
    elements["ccc"] = 333
    elements["ddd"] = 444
    fmt.Println(elements)

    delete(elements, "aaa")
    fmt.Println(elements)

    if name, ok := elements["denny"]; ok {
        fmt.Println(name)
    }else{
        fmt.Println(ok)
    }

    elements = map[string]int{
        "eee": 555,
        "fff": 666,
    }

    fmt.Println(elements)

    elements1 := map[string]map[string]string{
        "A": map[string]string{
            "name":  "Hydrogen",
            "state": "gas",
        },
        "B": map[string]string{
            "name":  "Helium",
            "state": "gas",
        },
        "C": map[string]string{
            "name":  "Lithium",
            "state": "solid",
        },
        "D": map[string]string{
            "name":  "Beryllium",
            "state": "solid",
        },
        "E": map[string]string{
            "name":  "Boron",
            "state": "solid",
        },
        "F": map[string]string{
            "name":  "Carbon",
            "state": "solid",
        },
        "G": map[string]string{
            "name":  "Nitrogen",
            "state": "gas",
        },
        "H": map[string]string{
            "name":  "Oxygen",
            "state": "gas",
        },
        "I": map[string]string{
            "name":  "Fluorine",
            "state": "gas",
        },
        "J": map[string]string{
            "name":  "Neon",
            "state": "gas",
        },
    }
    for key, values := range elements1 { 
        for k, v := range values {
            fmt.Println(key)
            fmt.Println(k,v)
        }
    }
}