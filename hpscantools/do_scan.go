package main

import (
    "bytes"
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/xml"
)

type ScanSettings struct {
    XMLName         xml.Name    `xml:"http://www.hp.com/schemas/imaging/con/cnx/scan/2008/08/19 ScanSettings"`
    XResolution     int         `xml:"XResolution"` // 200
    YResolution     int         `xml:"YResolution"` // 200
    XStart          int         `xml:"XStart"` // 0
    YStart          int         `xml:"YStart"` // 0
    Width           int         `xml:"Width"`  // 2550
    Height          int         `xml:"Height"` // 3507
    Format          string      `xml:"Format"` // "Raw", "Jpeg"
    CompressionQFactor int      `xml:"CompressionQFactor"` // 0, 15
    ColorSpace      string      `xml:"ColorSpace"` // "Color"
    BitDepth        int         `xml:"BitDepth"` // 8
    InputSource     string      `xml:"InputSource"` // "Platen"
    GrayRendering   string      `xml:"GrayRendering"` // "NTSC"
    Gamma           int         `xml:"ToneMap>Gamma"` // 1000
    Brightness      int         `xml:"ToneMap>Brightness"` // 1000
    Contrast        int         `xml:"ToneMap>Contrast"` // 1000
    Highlite        int         `xml:"ToneMap>Highlite"` // 179
    Shadow          int         `xml:"ToneMap>Shadow"` // 25
    Threshold       int         `xml:"ToneMap>Threshold"` // 0
    SharpeningLevel int         `xml:"SharpeningLevel"` // 128
    NoiseRemoval    int         `xml:"NoiseRemoval"` // 0
    ContentType     string      `xml:"ContentType"` // "Photo", 
}

func DefaultSettings() *ScanSettings {
    return &ScanSettings{XResolution: 200, YResolution:200, XStart:0, YStart:0, Width: 2550, Height: 3507, Format: "Jpeg", CompressionQFactor: 15, ColorSpace: "Color", BitDepth: 8, InputSource: "Platen", GrayRendering:"NTSC", Gamma: 1000, Brightness: 1000, Contrast: 1000, Highlite: 179,Shadow: 25, Threshold: 0, SharpeningLevel: 128, NoiseRemoval: 0, ContentType: "Photo"}
}

func JpegScanSettings(xres int, yres int) *ScanSettings {
    s := DefaultSettings()
    s.XResolution = xres
    s.YResolution = yres
    return s
}

func RawScanSettings(xres int, yres int) *ScanSettings {
    s := DefaultSettings()
    s.XResolution = xres
    s.YResolution = yres
    s.Format = "Raw"
    s.CompressionQFactor = 0
    return s
}

type CancelScan struct {
    XMLName         xml.Name    `xml:"http://www.hp.com/schemas/imaging/con/ledm/jobs/2009/04/30 Job"`
    JobUrl          string                  // The job url from POST-ing SystemSettings
    JobState        string                  // "Canceled"
}

func main() {
    s := RawScanSettings(200, 200)
    xmlString, err := xml.MarshalIndent(s, "", "    ")

    if err != nil {
            fmt.Println(err)
    }
    payload := bytes.NewBuffer([]byte ( xml.Header+string(xmlString) ))
    //payload := bytes.NewBuffer(xmlString)

    fmt.Printf("%s \n", payload)

    resp, err := http.Post("http://httpbin.org/post", "test/xml", payload)

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    fmt.Printf("%s \n", body)

}
