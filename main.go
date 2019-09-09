package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
    "github.com/labstack/echo"
    "github.com/sirupsen/logrus"
    "github.com/google/uuid"
)

var log = logrus.New()

func init() {
    log.Formatter = new(logrus.JSONFormatter)
    log.Formatter = new(logrus.TextFormatter) // default
    log.Level = logrus.DebugLevel
}

func main() {
    fmt.Println("Main function :")
    
    e := echo.New()
    // Routes
    e.GET("/",  func(c echo.Context) error {
        html, err := ioutil.ReadAll(c.Request().Body)
        check(err)

        id, err := uuid.NewUUID()
        check(err)

        stringId := id.String()

        var htmlFilePath = "/tmp/" + stringId + ".html"
        var pdfFilePath = "/tmp/" + stringId + ".pdf"

        file, err := os.Create(htmlFilePath)
        check(err)

        _, err = file.Write(html)
        check(err)

        defer file.Close()

        if _, err := os.Stat(htmlFilePath); err == nil {
            fmt.Println("File exists")
        }

        args := []string{
            "--no-sandbox", 
            "--headless", 
            "--disable-gpu", 
            "--disable-software-rasterizer",
            "--disable-gpu",
            "--print-to-pdf=" + pdfFilePath,
            htmlFilePath,
        }

        cmd := exec.Command("chromium-browser", args...)

        var out bytes.Buffer
        var stderr bytes.Buffer
        cmd.Stdout = &out
        cmd.Stderr = &stderr
        err = cmd.Run()

        if err != nil {
            fmt.Println(fmt.Sprint(err) + ": " + stderr.String())

            return err
        }

        fmt.Println("Chromium : " + out.String())

        check(err)

        return c.File(pdfFilePath)
    })

    e.Logger.Fatal(e.Start(":8080"))
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}