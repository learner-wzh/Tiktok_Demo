package basic

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"tiktok_Demo/config"
	"tiktok_Demo/models"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

// 截取视频第几帧图片
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (ImagePath string, err error) {
	snapshotPath = "/root/videoImage/" + snapshotPath
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	imgPath := snapshotPath + ".png"

	return imgPath, nil
}

// 将视频文件上传到阿里云oss
func UploadVideoToAliyunOss(file *multipart.FileHeader, token string) (string, error) {
	bucket := models.InitBucket
	oss := config.ReturnOss()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	user, err := models.QueryUserInfoByToken(token)
	var filename string
	if err == nil {
		filename = fmt.Sprintf("%d_%s", user.UserID, file.Filename)
	}

	// 将视频流上传至videos目录下
	path := "videos/" + filename
	err = bucket.PutObject(path, src)
	if err != nil {
		return "", err
	}

	// 获取对应视频流的URL
	VideoUrl := oss.VideoUrl + filename

	return VideoUrl, nil
}

// 将视频封面文件上传到阿里云oss
func UploadImageToAliyunOss(imagePath string, imageName string) (string, error) {
	bucket := models.InitBucket
	oss := config.ReturnOss()

	path := "images/" + imageName
	err := bucket.PutObjectFromFile(path, imagePath)
	if err != nil {
		return "", err
	}

	// 获取对应封面的URL
	ImageUrl := oss.ImageUrl + imageName
	return ImageUrl, nil
}

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	videoUrl, err := UploadVideoToAliyunOss(data, token)

	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if user, err := models.QueryUserInfoByToken(token); err == nil {
		finalName := fmt.Sprintf("%d_%s", user.UserID, filename)
		saveFile := filepath.Join("/root/video", finalName)

		strArray := strings.Split(finalName, ".")
		ImageName := strArray[0]

		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}

		imagePath, _ := GetSnapshot(saveFile, ImageName, 1)
		imageName := ImageName + ".png"

		imageUrl, err := UploadImageToAliyunOss(imagePath, imageName)

		var video models.Video
		video.CoverURL = imageUrl
		video.PlayURL = videoUrl
		video.Title = title
		video.UserID = user.UserID
		models.AddVideo(video)

		if err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  filename + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {

	// token := c.Query("token")
	userId := c.Query("user_id")

	UserID, _ := strconv.ParseInt(userId, 10, 64)

	videoList, err := models.QueryVideoListByUserId(UserID)

	var userInfo []models.UserInfo

	for _, video := range videoList {
		info, _ := models.QueryUserInfoByID(video.UserID)
		userInfo = append(userInfo, info)
	}

	if err == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
	}

}
