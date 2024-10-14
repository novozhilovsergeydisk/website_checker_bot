package checker

import (
    "net/http"
    "errors"
)

func CheckWebsite(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", errors.New("unable to reach the website")
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        return "available", nil
    }
    return "unavailable", nil
}
