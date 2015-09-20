
package main

import (
        "fmt"
        //"io/ioutil"
        "os"
        "encoding/xml"
)


type ScanSettings struct {

    XMLName         xml.Name    `xml:"ScanSettings"`
    XResolution     int         `xml:"XResolution"` // 200
    YResolution     int         `xml:"YResolution"` // 200
    XStart          int         `xml:"XStart"` // 0
    YStart          int         `xml:"YStart"` // 0
    Width           int         `xml:"Width"`  // 2550
    Height          int         `xml:"Height"` // 3507
    Format          string      `xml:"Format"` // "Raw", "Jpeg"
    CompressionQFactor int      `xml:"CompressionQFactor"` // 15
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



func main() {
    s := &ScanSettings{ XResolution:200, YResolution:201 }
    xmlString, err := xml.MarshalIndent(s, "", "    ")

    if err != nil {
            fmt.Println(err)
    }

    fmt.Printf("%s \n", string(xmlString))
    os.Stdout.Write(xmlString)
}
