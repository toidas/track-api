package appcontroller

import (
	"fmt"
	"net/http"
	"../../helpers/database"
	"../../models"

	"github.com/gin-gonic/gin"
)

type Info struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Package     string `json:"package"`
	VersionName string `json:"version_name"`
	VersionCode string `json:"version_code"`
}
type Admob struct {
	Banner       string `json:"banner"`
	Interstitial string `json:"interstitial"`
	Video        string `json:"video"`
}

type Unity struct {
	VideoId string `json:"video_id"`
}

type Facebook struct {
	FacebookName  string `json:"facebook_name"`
	FacebookAppId string `json:"facebook_app_id"`
	AdsId         string `json:"ads_id"`
}

type StartApp struct {
	StartAppId string `json:"startapp_id"`
}

type Control struct {
	Interstitial string `json:"interstitial"`
	Video        string `json:"video"`
}

type Ads struct {
	Admob    Admob    `json:"admob"`
	Unity    Unity    `json:"unity"`
	Facebook Facebook `json:"facebook"`
	StartApp StartApp `json:"startapp"`
	Control  Control  `json:"control"`
}

func mapAppToResult(app models.Application) (info Info, ads Ads) {
	var myInfo Info
	var myAdmob Admob
	var myUnity Unity
	var myFacebook Facebook
	var myStartapp StartApp
	var control Control
	var myAds Ads

	//convert to control
	control.Video = app.Video
	control.Interstitial = app.Interstitial
	// convert to info
	myInfo.Name = app.Name
	myInfo.Type = app.Type
	myInfo.Package = app.Package
	myInfo.VersionName = app.VersionName
	myInfo.VersionCode = app.VersionCode

	//convert admob
	myAdmob.Banner = app.AdmobBanner
	myAdmob.Interstitial = app.AdmobInterstitial
	myAdmob.Video = app.AdmobVideo

	//convert unity
	myUnity.VideoId = app.UnityVideo

	//convert facebook
	myFacebook.FacebookName = app.FacebookName
	myFacebook.FacebookAppId = app.FacebookAppId
	myFacebook.AdsId = app.FacebookAds

	//convert startapp
	myStartapp.StartAppId = app.StartappId

	//convert ads
	myAds.Admob = myAdmob
	myAds.Facebook = myFacebook
	myAds.Unity = myUnity
	myAds.StartApp = myStartapp
	myAds.Control = control
	return myInfo, myAds
}
func GetAppWithToken(c *gin.Context) {
	token := c.Param("token")
	fmt.Println("Receive request with token: " + token)
	db := database.Connect()
	var app models.Application
	db.Where("token = ?", token).First(&app)
	info, ads := mapAppToResult(app)
	c.JSON(http.StatusOK, gin.H{"info": info, "ads": ads})
}
