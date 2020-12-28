// fetch all busca URLs em parelelo e informa os tempos gastos e os tamanhos
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go feth(url, ch) // inicia uma gorrotina
	}

	for range os.Args[1:]{
		fmt.Println(<-ch) // recebe do canal ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func feth(url string, ch chan string) {
	start := time.Now()
	resp , err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes , err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // evita vazamento de recursos
	if err != nil {
		ch <- fmt.Sprintf("wile reading %s: %v:", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}

// para executar
// go build /lpgo/ch1/fetchall.go
//  ./fetchall http://google.com.br http://youtube.com http://aws.com
