package face
import (
	"github.com/noverguo/util"
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
)


type _gender struct {
	Confidence float32	`json:"confidence"`
	Value string		`json:"value"`
}

type _race struct {
	Confidence float32	`json:"confidence"`
	Value string		`json:"value"`
}

type _smiling struct {
	Value float32		`json:"value"`
}

type _attribute struct {
	Age map[string]int	`json:"age"`
	Gender _gender		`json:"gender"`
	Race _race			`json:"race"`
	Smiling _smiling	`json:"smiling"`
}

type _point struct {
	X float32			`json:"x"`
	Y float32			`json:"y"`
}

func (this _point) String() string {
	return fmt.Sprintf("(%f, %f)", this.X, this.Y)
}

type _position struct {
	Center _point			`json:"center"`
	Width float32			`json:"width"`
	Height float32			`json:"height"`
	Eye_left _point			`json:"eye_left"`
	Eye_right _point		`json:"eye_right"`
	Mouth_left _point		`json:"mouth_left"`
	Mouth_right _point		`json:"mouth_right"`
	Nose _point				`json:"nose"`
}

type _face struct {
	Attribute _attribute	`json:"attribute"`
	Face_id string			`json:"face_id"`
	Position _position		`json:"position"`
	Tag string				`json:"tag"`
}

type Faces struct {
	Faces[] _face			`json:"face"`
	Img_id string			`json:"img_id"`
	Img_height int			`json:"img_height"`
	Img_width int			`json:"img_width"`
	Session_id string		`json:"session_id"`
	Url string				`json:"url"`
}

// 通过face++ API进行人脸检测
func FaceDetect_faceplusplus(fileName string) *Faces {
	params := make(map[string]string)
	params["api_key"] = "7b73aadfe253b1e6b1db1f9037dd46de"
	params["api_secret"] = "hqm2yQthlRydFdvLBkHRfw59JDfcE8KX"
	req, err := util.NewFileUploadRequest("http://apicn.faceplusplus.com/v2/detection/detect", params, "img", fileName)
	util.CheckErr(err)

	resp, err := http.DefaultClient.Do(req)
	util.CheckErr(err)
	if resp.StatusCode != 200 {
		fmt.Println("\thttp status: ", resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	util.CheckErr(err)

	fs := &Faces{}
	err = json.Unmarshal(body, &fs)
	util.CheckErr(err)
	return fs
}