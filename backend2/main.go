package main

import (
  "io"
  "log"
  "net/http"
  "net/url"
  "os"
  "strconv"
  "time"

  "github.com/joho/godotenv"
  "github.com/julienschmidt/httprouter"
  "github.com/yuin/gopher-lua"
  "github.com/minio/minio-go"
)

var listen string
var minioClient *minio.Client

func main() {
  // Load .env
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
  listen = os.Getenv("LISTEN")
  minioEndpoint := os.Getenv("MINIO_ENDPOINT")
  minioID := os.Getenv("MINIO_ACCESS_KEY")
  minioKey := os.Getenv("MINIO_SECRET_KEY")

  // Minio client
  minioClient, err = minio.New(minioEndpoint, minioID, minioKey, false)
  if err != nil {
    log.Fatal("Error loading minio")
  }

  // Create bucket if it doesn't exist
  err = minioClient.MakeBucket("parser", "us-east-1")
  if err != nil {
    exists, err := minioClient.BucketExists("parser")
    if err == nil && exists {
      log.Printf("Bucket %s already exists", "parser")
    } else {
      log.Printf("%s", err)
      log.Fatal("Error creating bucket")
    }
  } else {
    log.Printf("Created bucket %s", "parser")
  }

  // Routes
  router := httprouter.New()
  // Return minio presigned URLs
  router.GET("/model", GetModel)
  router.POST("/model", UploadModel)
  router.GET("/data", GetData)
  router.POST("/data", UploadData)
  router.GET("/data_parser", GetDataParser)
  router.POST("/data_parser", UploadDataParser)
  router.POST("/batch", BatchData)

  router.POST("/parse", TestParse)

  // Start server
  log.Printf("starting server on %s", listen)
  log.Fatal(http.ListenAndServe(listen, router))
}

func UploadModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  bucketName := r.FormValue("id")

  // Create bucket if it doesn't exist
  err := minioClient.MakeBucket(bucketName, "us-east-1")
  if err != nil {
    exists, err := minioClient.BucketExists(bucketName)
    if !(err == nil && exists) {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }
  }

  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedPutObject(bucketName, "model", expiry)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  w.Write([]byte(presignedURL.String()))
}

func UploadData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  bucketName := r.FormValue("id")
  hash := r.FormValue("hash")

  exists, err := minioClient.BucketExists(bucketName)
  if !(err == nil && exists) {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedPutObject(bucketName, hash, expiry)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  w.Write([]byte(presignedURL.String()))
}

func UploadDataParser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := r.FormValue("id")

  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedPutObject("parser", id, expiry)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  w.Write([]byte(presignedURL.String()))
}

func GetModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := r.FormValue("id")
  if id == "" {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  reqParams := make(url.Values)
  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedGetObject(id, "model", expiry, reqParams)
  if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
  }

  w.Write([]byte(presignedURL.String()))
}

func GetData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  model := r.FormValue("model")
  if model == "" {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  id := r.FormValue("id")
  if id == "" {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  reqParams := make(url.Values)
  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedGetObject(model, id, expiry, reqParams)
  if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
  }

  w.Write([]byte(presignedURL.String()))
}

func GetDataParser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  id := r.FormValue("id")
  if id == "" {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  reqParams := make(url.Values)
  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedGetObject("parser", id, expiry, reqParams)
  if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
  }

  w.Write([]byte(presignedURL.String()))
}

func BatchData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  dataParserId := r.FormValue("data_parser")
  labelParserId := r.FormValue("label_parser")

  modelId := r.FormValue("model_id")
  dataId := r.FormValue("data_id")

  if dataParserId == "" || labelParserId == "" || modelId == "" || dataId == "" {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  
}

func TestParse(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  file, _, err := r.FormFile("file")
  if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
  }
  defer file.Close()

  L := lua.NewState()
  defer L.Close()
  err = L.DoFile("../mnist_data_parser.lua")
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  num := 0

  buf := make([]byte, 512)
  for {
    n, err := file.Read(buf)
    if err == io.EOF {
      break
    } else if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    err = L.CallByParam(lua.P {
      Fn: L.GetGlobal("parse"),
      NRet: 1,
      Protect: true,
      }, lua.LString(buf), lua.LNumber(n))
    if err != nil {
      log.Printf("%s", err)
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    lv := L.Get(-1)
    L.Pop(1)
    if table, ok := lv.(*lua.LTable); ok {
      num += table.Len()
    }
  }

  w.Write([]byte(strconv.Itoa(num)))
}
