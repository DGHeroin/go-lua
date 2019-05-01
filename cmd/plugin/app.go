package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "plugin"
)

func openPlugin(filePath string, startFunc string, arg string) error {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println(e)
        }
    }()
    p, err := plugin.Open(filePath)
    if err != nil {
        return err
    }
    sym, err := p.Lookup(startFunc)
    if err != nil {
        return err
    }

    if f, ok := sym.(func(string) error); ok {
        return f(arg)
    } else {
        return fmt.Errorf("{func %s() error} not dound", startFunc)
    }
}

type PluginInfo struct {
    FileName  string `json:"path"`
    EntryFunc string `json:"func"`
    Arg       string `json:"arg"`
}

func readConfig(path string) (result []PluginInfo, err error) {
    var data []byte
    data, err = ioutil.ReadFile(path)
    if err != nil { return }
    json.Unmarshal(data, &result)
    return
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("config file not found")
        return
    }
    plugins, err := readConfig(os.Args[1])
    if err != nil { return }
    for _, plugin := range plugins {
        err := openPlugin(plugin.FileName, plugin.EntryFunc, plugin.Arg)
        if err != nil {
            fmt.Println(err)
        }
    }
}
