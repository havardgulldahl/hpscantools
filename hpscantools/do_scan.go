package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)


// XML Structs and their helpers //
// ============================= //


// The xml to send to start a job
type ScanSettings struct {
	XMLName            xml.Name `xml:"http://www.hp.com/schemas/imaging/con/cnx/scan/2008/08/19 ScanSettings"`
	SchemaLocation     string   `xml:"http://www.w3.org/2001/XMLSchema-instance schemaLocation,attr"`
	XResolution        int      `xml:"XResolution"`        // 200
	YResolution        int      `xml:"YResolution"`        // 200
	XStart             int      `xml:"XStart"`             // 0
	YStart             int      `xml:"YStart"`             // 0
	Width              int      `xml:"Width"`              // 2550
	Height             int      `xml:"Height"`             // 3507
	Format             string   `xml:"Format"`             // "Raw", "Jpeg"
	CompressionQFactor int      `xml:"CompressionQFactor"` // 0, 15
	ColorSpace         string   `xml:"ColorSpace"`         // "Color"
	BitDepth           int      `xml:"BitDepth"`           // 8
	InputSource        string   `xml:"InputSource"`        // "Platen"
	GrayRendering      string   `xml:"GrayRendering"`      // "NTSC"
	Gamma              int      `xml:"ToneMap>Gamma"`      // 1000
	Brightness         int      `xml:"ToneMap>Brightness"` // 1000
	Contrast           int      `xml:"ToneMap>Contrast"`   // 1000
	Highlite           int      `xml:"ToneMap>Highlite"`   // 179
	Shadow             int      `xml:"ToneMap>Shadow"`     // 25
	Threshold          int      `xml:"ToneMap>Threshold"`  // 0
	SharpeningLevel    int      `xml:"SharpeningLevel"`    // 128
	NoiseRemoval       int      `xml:"NoiseRemoval"`       // 0
	ContentType        string   `xml:"ContentType"`        // "Photo",
}

// call this to get default settings
func DefaultSettings() *ScanSettings {
	return &ScanSettings{SchemaLocation: "http://www.hp.com/schemas/imaging/con/cnx/scan/2008/08/19 Scan Schema - 0.26.xsd", XResolution: 200, YResolution: 200, XStart: 0, YStart: 0, Width: 2550, Height: 3507, Format: "Jpeg", CompressionQFactor: 15, ColorSpace: "Color", BitDepth: 8, InputSource: "Platen", GrayRendering: "NTSC", Gamma: 1000, Brightness: 1000, Contrast: 1000, Highlite: 179, Shadow: 25, Threshold: 0, SharpeningLevel: 128, NoiseRemoval: 0, ContentType: "Photo"}
}

// call this to get jpeg photo settings
func JpegScanSettings(xres int, yres int) *ScanSettings {
	s := DefaultSettings()
	s.XResolution = xres
	s.YResolution = yres
	return s
}


// call this to get raw settings
func RawScanSettings(xres int, yres int) *ScanSettings {
	s := DefaultSettings()
	s.XResolution = xres
	s.YResolution = yres
	s.Format = "Raw"
	s.CompressionQFactor = 0
	return s
}


// xml to cancel a print job
type CancelJob struct {
	XMLName  xml.Name `xml:"http://www.hp.com/schemas/imaging/con/ledm/jobs/2009/04/30 Job"`
	JobUrl   string   // The job url from POST-ing SystemSettings
	JobState string   // "Canceled"
}


func Cancel(jobId string) bool {

  return false
}



type Job struct {
//	XMLName    xml.Name   `xml:"http://www.hp.com/schemas/imaging/con/ledm/jobs/2009/04/30 Job"`
  JobState   string     `xml:"http://www.hp.com/schemas/imaging/con/ledm/jobs/2009/04/30 Job>JobState"`
  BinaryURL  string     `xml:"http://www.hp.com/schemas/imaging/con/cnx/scan/2008/08/19 ScanJob>PreScanPage>BinaryURL"`

}


func ParseJobURL(jobUrl string) string {
	resp, err := http.Get(jobUrl)
	if err != nil {
		panic(err)
	}
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

//   body, err := ioutil.ReadFile("doc/JobList.xml")
  fmt.Printf("%s \n", body)

  job := Job{}
	err = xml.Unmarshal(body, &job)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
  fmt.Printf("[DEBUG] url %s \n", job.BinaryURL)
  fmt.Printf("[DEBUG] state %s \n", job.JobState)
  return job.BinaryURL

}

// worker code //
// ============================= //



func StartJob(printer string) string {
  startJobURL := printer + "/Scan/Jobs"
	s := RawScanSettings(200, 200)
	xmlString, err := xml.MarshalIndent(s, "", "    ")

	if err != nil {
		fmt.Println(err)
	}
	payload := bytes.NewBuffer([]byte(xml.Header + string(xmlString)))

  fmt.Printf("%s \n", payload)
	resp, err := http.Post(startJobURL, "test/xml", payload)
	if err != nil {
		fmt.Println(err)
	}
	location := resp.Header.Get("Location")
	fmt.Printf("Location: --%s--", location)
	for k, v := range resp.Header {
		fmt.Printf("key:%s, value:%s \n", k, v)
	}
  return location
}

func SaveImage(imageURL string) bool {
	resp, err := http.Get(imageURL)
	if err != nil {
		fmt.Println(err)
	}
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  err = ioutil.WriteFile("image.png", body, 0644)
  return err != nil
}


func main() {
	printer := "http://192.168.0.196:8080"

  jobURL := StartJob(printer)

  fmt.Printf("jobUrl: %s", jobURL)

  imageURL := printer + ParseJobURL(jobURL)
  fmt.Println(imageURL)
  imageSaved := SaveImage(imageURL)
  fmt.Printf("IMage was saved ok: %s", imageSaved)


}
