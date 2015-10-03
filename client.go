package main 
 
import ( 
	"fmt"
  "flag"
  "bytes"
  "strconv"
  "net/http"
  "github.com/gorilla/rpc/json"
) 

type Args struct {
    StockSymbolAndPercentage string 
    Budget float32                  
    TradeId int                       
}

type Reply struct {
    Message string
}

func ClientRequest(method string, args Args) (reply Reply, err error) {
    req, err := json.EncodeClientRequest(method, args)
    if err != nil {
        fmt.Println("ClientRequest Error")
        fmt.Println(err)
        return reply, err
    }
    res, err := http.Post("http://127.0.0.1:8080/rpc", "application/json", bytes.NewBuffer(req))
    if err != nil{
        fmt.Println("Post Error")
        fmt.Println(err)
        return reply, err
    }
    defer res.Body.Close()
    err = json.DecodeClientResponse(res.Body, &reply)
    return reply, err
}

func main() {
    var buyargs Args
    var checkargs Args
 	flag.Parse() // get the arguments from command line
    method := "Service."+flag.Arg(0) // get the source directory from 1st argument
    if flag.Arg(0) == "Buying" {
        fmt.Println("Buying Now...")
        fmt.Println(method)
        buyargs.StockSymbolAndPercentage = flag.Arg(1)
        budget,_ := strconv.ParseFloat(flag.Arg(2), 64)
        id,_ := strconv.Atoi(flag.Arg(3))
        buyargs.TradeId = id;
        buyargs.Budget = float32(budget)
        reply,_ := ClientRequest(method, buyargs)
        fmt.Println(reply.Message)
    }
    if flag.Arg(0) == "Checking" {
        fmt.Println("Checking Now...")
        fmt.Println(method)
        id,_ := strconv.Atoi(flag.Arg(1))
        checkargs.TradeId = id
        reply,_ := ClientRequest(method, checkargs)
        fmt.Println(reply.Message)
    } 
} 
