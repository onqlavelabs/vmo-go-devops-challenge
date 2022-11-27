package util

import (
    "encoding/json"
    "io/ioutil"
)

func ReadJsonFile(path string, data interface{}) error {
    content, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }

    err = json.Unmarshal(content, &data)
    if err != nil {
        return err
    }

    return nil
}
