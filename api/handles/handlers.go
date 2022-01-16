package handles

import (
	"Video-Website/api/commen"
	"Video-Website/api/defs"
	"Video-Website/api/middleWare"
	"Video-Website/api/session"
	"Video-Website/api/utils"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	//"io"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &commen.UserPamas{}

	if err := json.Unmarshal(res, ubody); err != nil {
		commen.SendErrorResponse(w, commen.ErroRequestBodyParseFailed)
		return
	}

	if err := defs.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		commen.SendErrorResponse(w, commen.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &commen.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		commen.SendErrorResponse(w, commen.ErrorInternalFaults)
		return
	} else {
		commen.SendNormalResponse(w, string(resp), 201)
	}
}

// func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	uname := p.ByName("user_name")
// 	io.WriteString(w, uname)
// }

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	log.Printf("%s", res)
	ubody := &commen.UserPamas{}
	if err := json.Unmarshal(res, ubody); err != nil {
		log.Printf("%s", err)
		//io.WriteString(w, "wrong")
		commen.SendErrorResponse(w, commen.ErroRequestBodyParseFailed)
		return
	}

	// Validate the request body
	uname := p.ByName("username")
	log.Printf("Login url name: %s", uname)
	log.Printf("Login body name: %s", ubody.Username)
	if uname != ubody.Username {
		commen.SendErrorResponse(w, commen.ErrorNotAuthUser)
		return
	}

	log.Printf("%s", ubody.Username)
	pwd, err := defs.GetUserCredential(ubody.Username)
	log.Printf("Login pwd: %s", pwd)
	log.Printf("Login body pwd: %s", ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		commen.SendErrorResponse(w, commen.ErrorNotAuthUser)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	si := &commen.SignedIn{Success: true, SessionId: id}
	if resp, err := json.Marshal(si); err != nil {
		commen.SendErrorResponse(w, commen.ErrorInternalFaults)
	} else {
		commen.SendNormalResponse(w, string(resp), 200)
	}

	//io.WriteString(w, "signed in")
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleWare.ValidateUser(w, r) {
		log.Printf("Unathorized scheduler \n")
		return
	}

	uname := p.ByName("username")
	u, err := defs.GetUser(uname)
	if err != nil {
		log.Printf("Error in GetUserInfo: %s", err)
		commen.SendErrorResponse(w, commen.ErrorDBError)
		return
	}

	ui := &commen.UserInfo{Id: u.Id}
	if resp, err := json.Marshal(ui); err != nil {
		commen.SendErrorResponse(w, commen.ErrorInternalFaults)
	} else {
		commen.SendNormalResponse(w, string(resp), 200)
	}

}

func AddNewVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleWare.ValidateUser(w, r) {
		log.Printf("Unathorized scheduler \n")
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	nvbody := &commen.NewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		log.Printf("%s", err)
		commen.SendErrorResponse(w, commen.ErroRequestBodyParseFailed)
		return
	}

	vi, err := defs.AddNewVideo(nvbody.AuthorId, nvbody.Name)
	log.Printf("Author id : %d, name: %s \n", nvbody.AuthorId, nvbody.Name)
	if err != nil {
		log.Printf("Error in AddNewVideo: %s", err)
		commen.SendErrorResponse(w, commen.ErrorDBError)
		return
	}

	if resp, err := json.Marshal(vi); err != nil {
		commen.SendErrorResponse(w, commen.ErrorInternalFaults)
	} else {
		commen.SendNormalResponse(w, string(resp), 201)
	}

}

func ListAllVideos(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleWare.ValidateUser(w, r) {
		return
	}

	uname := p.ByName("username")
	vs, err := defs.ListVideoInfo(uname, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ListAllvideos: %s", err)
		commen.SendErrorResponse(w, commen.ErrorDBError)
		return
	}

	vsi := &commen.VideosInfo{Videos: vs}
	if resp, err := json.Marshal(vsi); err != nil {
		commen.SendErrorResponse(w, commen.ErrorInternalFaults)
	} else {
		commen.SendNormalResponse(w, string(resp), 200)
	}

}

func DeleteVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleWare.ValidateUser(w, r) {
		return
	}

	vid := p.ByName("vid-id")
	err := defs.DeleteVideoInfo(vid)
	if err != nil {
		log.Printf("Error in DeletVideo: %s", err)
		commen.SendErrorResponse(w, commen.ErrorDBError)
		return
	}
	// todo 后续待完善
	//go utils.SendDeleteVideoRequest(vid)
	commen.SendNormalResponse(w, "", 204)
}

func PostComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleWare.ValidateUser(w, r) {
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)

	cbody := &commen.NewComment{}
	if err := json.Unmarshal(reqBody, cbody); err != nil {
		log.Printf("%s", err)
		commen.SendErrorResponse(w, commen.ErroRequestBodyParseFailed)
		return
	}

	vid := p.ByName("vid-id")
	if err := defs.AddNewComments(vid, cbody.AuthorId, cbody.Content); err != nil {
		log.Printf("Error in PostComment: %s", err)
		commen.SendErrorResponse(w, commen.ErrorDBError)
	} else {
		commen.SendNormalResponse(w, "ok", 201)
	}

}

func ShowComments(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if !middleWare.ValidateUser(w, r) {
		return
	}

	vid := p.ByName("vid-id")
	cm, err := defs.ListComments(vid, 0, utils.GetCurrentTimestampSec())
	if err != nil {
		log.Printf("Error in ShowComments: %s", err)
		commen.SendErrorResponse(w, commen.ErrorDBError)
		return
	}

	cms := &commen.Comments{Comments: cm}
	if resp, err := json.Marshal(cms); err != nil {
		commen.SendErrorResponse(w, commen.ErrorInternalFaults)
	} else {
		commen.SendNormalResponse(w, string(resp), 200)
	}
}
