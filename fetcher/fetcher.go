package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(500 * time.Millisecond)

func Fetcher(url string) ([]byte, error) {
	<-rateLimiter
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {

		return nil, err
	}
	ip := CreateIp()
	request.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	request.Header.Set("CLIENT-IP", ip);
	request.Header.Set("X-FORWARDED-FOR", ip);
	request.Header.Set("REMOTE_ADDR", ip);
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {

		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

		return nil, fmt.Errorf("Url:%s ,Wrong status code is %d", url, response.StatusCode)
	}
	reader := bufio.NewReader(response.Body)

	return ioutil.ReadAll(reader)
}

func CreateIp() string {
	rand.Seed(time.Now().Unix())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))

	return ip
}
