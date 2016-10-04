//http://stackoverflow.com/questions/26723467/what-is-the-difference-between-form-data-x-www-form-urlencoded-and-raw-in-the-p

package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	myURI = "http://localhost:3000/api/v1/callbacks/palette_video_transcode_success"
)

var (
	TestClient Client
)

type Client struct {
	client *http.Client
}

func init() {
	TestClient = Client{&http.Client{}}
}

func (c Client) PostForm() (*http.Response, error) {
	data := url.Values{}
	data.Add("Name", "Vimlesh")
	data.Add("Password", "password")
	//Content-Type by default is application/x-www-form-urlencoded
	return c.client.PostForm(myURI, data)
}

//Post ...
func (c Client) Post() (*http.Response, error) {
	data := url.Values{}
	data.Add("Name", "Vimlesh")
	data.Add("Password", "password")
	return c.client.Post(myURI, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

//PostJSONData  http://stackoverflow.com/questions/24455147/how-do-i-send-a-json-string-in-a-post-request-in-go
func (c Client) PostJSONData() (*http.Response, error) {
	var jsonStr = []byte(`{
 	"Name": "Ed",
 	"Text": "Knock knock.",
 	"NAme": "Rocker"
    }`)
	return c.client.Post(myURI, "application/json", bytes.NewBuffer(jsonStr))
}

/*PostFileAndData : Reference
http://stackoverflow.com/questions/20205796/golang-post-data-using-the-content-type-multipart-form-data
Internally it usage form-Data content type look at mime/multipart/writer.go
*/
func (c Client) PostFileAndData() (*http.Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open("/home/vimlesh/Downloads/test.mp4")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fw, err := w.CreateFormFile("image", " /home/vimlesh/Downloads/test.mp4")
	if err != nil {
		panic(err)
	}
	if _, err = io.Copy(fw, f); err != nil {
		panic(err)
	}
	if fw, err = w.CreateFormField("Password"); err != nil {
		panic(err)
	}
	if _, err = fw.Write([]byte("myPassword#12345")); err != nil {
		panic(err)
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()
	return c.client.Post(myURI, w.FormDataContentType(), &b)
}
