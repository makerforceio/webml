package main

import (
  "bytes"
  "crypto/rand"
  "encoding/hex"
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
  router.GET("/labels", GetLabels)
  router.POST("/labels", UploadLabels)
  router.GET("/data/batch", GetBatchData)
  router.GET("/labels/batch", GetBatchLabels)
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
  bucketName := r.FormValue("model")
  id := r.FormValue("id")

  exists, err := minioClient.BucketExists(bucketName)
  if !(err == nil && exists) {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedPutObject(bucketName, "data:" + id, expiry)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  w.Write([]byte(presignedURL.String()))
}

func UploadLabels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  bucketName := r.FormValue("model")
  id := r.FormValue("id")

  exists, err := minioClient.BucketExists(bucketName)
  if !(err == nil && exists) {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  expiry := time.Second * 120
  presignedURL, err := minioClient.PresignedPutObject(bucketName, "label:" + id, expiry)
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
  presignedURL, err := minioClient.PresignedGetObject(model, "data:" + id, expiry, reqParams)
  if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
  }

  w.Write([]byte(presignedURL.String()))
}

func GetLabels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
  presignedURL, err := minioClient.PresignedGetObject(model, "label:" + id, expiry, reqParams)
  if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
  }

  w.Write([]byte(presignedURL.String()))
}

func GetBatchData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
  presignedURL, err := minioClient.PresignedGetObject(model, "batch:data:" + id, expiry, reqParams)
  if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
  }

  w.Write([]byte(presignedURL.String()))
}

func GetBatchLabels(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
  presignedURL, err := minioClient.PresignedGetObject(model, "batch:label:" + id, expiry, reqParams)
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

  batchSize, err := strconv.Atoi(r.FormValue("batch_size"))
  if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  if dataParserId == "" || labelParserId == "" || modelId == "" || dataId == "" {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
    return
  }

  dataParserObject, err := minioClient.GetObject("parser", dataParserId, minio.GetObjectOptions{})
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  dataParserBytes := new(bytes.Buffer)
  dataParserBytes.ReadFrom(dataParserObject)
  dataParser := dataParserBytes.String()

  labelParserObject, err := minioClient.GetObject("parser", labelParserId, minio.GetObjectOptions{})
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }
  labelParserBytes := new(bytes.Buffer)
  labelParserBytes.ReadFrom(labelParserObject)
  labelParser := labelParserBytes.String()

  dataL := lua.NewState()
  defer dataL.Close()
  err = dataL.DoString(dataParser)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  dataObject, err := minioClient.GetObject(modelId, "data:" + dataId, minio.GetObjectOptions{})
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  batchIds := make([]string, 0)

  buf := make([]byte, 512)
  batch := make([][]byte, 0)
  for {
    n, err := dataObject.Read(buf)
    if err == io.EOF {
      break
    } else if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    err = dataL.CallByParam(lua.P {
      Fn: dataL.GetGlobal("parse"),
      NRet: 1,
      Protect: true,
      }, lua.LString(buf), lua.LNumber(n))
    if err != nil {
      log.Printf("%s", err)
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    lv := dataL.Get(-1)
    dataL.Pop(1)
    if table, ok := lv.(*lua.LTable); ok {
      table.ForEach(func(_ lua.LValue, v lua.LValue) {
        val := []byte(v.(lua.LString).String())
        batch = append(batch, val)
        if len(batch) >= batchSize {
          batchId := RandomHex()
          data := make([]byte, 0)
          for _, datum := range(batch) {
            data = append(data, datum...)
          }
          _, err := minioClient.PutObject(modelId, "batch:data:" + batchId, bytes.NewReader(data), -1, minio.PutObjectOptions{})
          if err != nil {
            log.Printf("%s", err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        		return
          }
          batchIds = append(batchIds, batchId)
        }
      })
    }
  }

  labelL := lua.NewState()
  defer labelL.Close()
  err = labelL.DoString(labelParser)
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  labelObject, err := minioClient.GetObject(modelId, "label:" + dataId, minio.GetObjectOptions{})
  if err != nil {
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
  }

  buf = make([]byte, 512)
  i := 0
  for {
    n, err := labelObject.Read(buf)
    if err == io.EOF {
      break
    } else if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    err = labelL.CallByParam(lua.P {
      Fn: labelL.GetGlobal("parse"),
      NRet: 1,
      Protect: true,
      }, lua.LString(buf), lua.LNumber(n))
    if err != nil {
      log.Printf("%s", err)
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    lv := labelL.Get(-1)
    labelL.Pop(1)
    if table, ok := lv.(*lua.LTable); ok {
      table.ForEach(func(_ lua.LValue, v lua.LValue) {
        val := []byte(v.(lua.LString).String())
        batch = append(batch, val)
        if len(batch) >= batchSize {
          batchId := batchIds[i]
          data := make([]byte, 0)
          for _, datum := range(batch) {
            data = append(data, datum...)
          }
          _, err := minioClient.PutObject(modelId, "batch:label:" + batchId, bytes.NewReader(data), -1, minio.PutObjectOptions{})
          if err != nil {
            log.Printf("%s", err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        		return
          }
          i += 1
          if i >= len(batchIds) {
            w.WriteHeader(200)
            return
          }
        }
      })
    }
  }

  w.WriteHeader(200)
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

func RandomHex() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic("unable to generate 16 bytes of randomness")
	}
	return hex.EncodeToString(b)
}
