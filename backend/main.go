package main

import (
  // "bytes"
  "io"
  "log"
  "net/http"
  "os"
  "strconv"

  "github.com/joho/godotenv"
  "github.com/julienschmidt/httprouter"
  "github.com/minio/minio-go"
  "github.com/yuin/gopher-lua"
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

  // Routes
  router := httprouter.New()
  router.GET("/lua/:value", TestLua)
  router.POST("/parse", TestParse)

  // Start server
  log.Printf("starting server on %s", listen)
  log.Fatal(http.ListenAndServe(listen, router))
}

// func UploadModel(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//   var buf bytes.Buffer
//   file, header, err := r.FormFile("file")
//   if err != nil {
//     http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return
//   }
//   defer file.Close()
// }

func TestParse(w http.ResponseWriter, r * http.Request, p httprouter.Params) {
  file, _, err := r.FormFile("file")
  if err != nil {
    http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
  }
  defer file.Close()

  L := lua.NewState()
  defer L.Close()

  if err := L.DoFile("mnist_data_parser.lua"); err != nil {
    log.Printf("%s", err)
    http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    return
  }

  num := 0

  buf := make([]byte, 256)
  for {
    _, err := file.Read(buf)
    if err == io.EOF {
      break
    } else if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
  		return
    }

    if err := L.CallByParam(lua.P{
      Fn: L.GetGlobal("parse"),
      NRet: 1,
      Protect: true,
      }, lua.LString(buf)); err != nil {
      log.Printf("%s", err)
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
    }

    lv := L.Get(-1) // get the value at the top of the stack
    L.Pop(1)
    if table, ok := lv.(*lua.LTable); ok {
      num += table.Len()
    }
  }

  w.Write([]byte(strconv.Itoa(num)))
}

func TestLua(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  L := lua.NewState()
  defer L.Close()

  if err := L.DoString(`
    function test_function (a)
      test = {}
      for i = 1, 10 do
        test[i] = i * i
      end
      return test
    end
  `); err != nil {
      panic(err)
  }
  a := p.ByName("value");

  if err := L.CallByParam(lua.P{
    Fn: L.GetGlobal("test_function"),
    NRet: 1,
    Protect: true,
    }, lua.LString(a)); err != nil {
    log.Printf("%s", err)
  }

  lv := L.Get(-1) // get the value at the top of the stack
  L.Pop(1)
  returnVal := ""
  if table, ok := lv.(*lua.LTable); ok {
    table.ForEach(func (key lua.LValue, val lua.LValue) {
      returnVal = returnVal + key.(lua.LNumber).String() + ", " + val.(lua.LNumber).String() + "\n"
    })
    w.Write([]byte(returnVal))
  }

}
