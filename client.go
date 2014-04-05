package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	Aditic = fetchAd("aditic")
	Adtech = fetchAd("adtech")
	Smaato = fetchAd("smaato")
)

type Result string

type Ad func(url string) Result

func fetchAd(dp string) Ad {
	return func(url string) Result {
		resp, err := http.Get(url)
		// fmt.Println("DP = ", dp)
		if err != nil {
			fmt.Println("Error fetching ad ", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		// fmt.Println("Body = ", string(body))
		return Result(fmt.Sprintf("ad from %s -> %s\n", dp, string(body)))
	}
}

func main() {
	fmt.Println("Fetching ads...")
	aditic := "http://generic.marketplace.aditic.net/?alid=10238&format=json&pid=15519b380232256&srcip=139.0.10.106&support=wap&ua=Mozilla%2F5.0+%28Linux%3B+U%3B+Android+2.1%3B+en-us%3B+Nexus+One+Build%2FERD62%29+AppleWebKit%2F530.17+%28KHTML%2C+like+Gecko%29+Version%2F4.0+Mobile+Safari%2F530.17&width=300"
	// buzzcity := "http://show.buzzcity.net/showads.php?get=rich&partnerid=95219&ip=112.215.64.252&ua=Mozilla/5.0 (iPad; U; CPU OS 3_2_1 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Mobile/7B405&limit=10&v=3"
	adtech := "http://a.adtech.de/adxml/2.0/23.3/0/0/-1/ADTECH;loc=100;rettype=js;alias=business_pub-top-5;random=09876543211234567890"
	smaato := "http://soma.smaato.net/oapi/reqAd.jsp?adspace=0&pub=0&beacon=true&devip=124.122.78.193&device=Nokia1680c-2%2F2.0+%2805.61%29+Profile%2FMIDP-2.1+Configuration%2FCLDC-1.1&format=IMG&response=xml"

	c := make(chan Result)

	var results [] Result

	go func() { c <- Aditic(aditic) } ()
	go func() { c <- Adtech(adtech) } ()
	go func() { c <- Smaato(smaato) } ()

	timeout := time.After(2000 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}

	}
	fmt.Println(results)
	return


	// go fetchContents(aditic, "aditic")
	// go fetchContents(buzzcity, "buzzcity")
	// go fetchContents(adtech, "adtech")
	// go fetchContents(smaato, "smaato")

	// fmt.Println("Got ad")
}
