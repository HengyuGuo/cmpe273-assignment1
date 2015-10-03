# cmpe273-assignment1
MSSE-2015-FALL CMPE273 Assignment1

To Start this project.
First run the server side, get into the correct directory, using  go run server.go  in command-line. Now the server is running.

Secondly, run the client side, get into the correct directory, using  go run server.go  in command-line.
To use the correct arguments, you need to put the correct service method in the first command-line argument, when you buy a stock, you will use Buying method. There are three arguments with this method. And the first argument is the StockSymbolAndPercentage: like GOOG:50%,YHOO:50%. The second argument is the Budget like 3000 or 10000. The third argument is the transaction id like 1 or 1234.
So when you buy a stock, the correct command-line input to run the client side is like: 
      go run client.go Buying GOOG:50%,YHOO:50% 3600 1
      
When you check your portfolio, you will use Checking method, the only one argument is the transaction id. The command-line input like :           go run client.go Checking 1

Congratulations! So far, you are able to use this RPC.
