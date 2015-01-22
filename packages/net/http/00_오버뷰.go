/*
http패키지는 http 클라이언트와 서버구현을 제공한다.
GET, HEAD, POST 와 PostForm  은 http(혹은 https)요청을 만든다.
*/
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
/*
클라이언트는 반ㄷ싀 응답몸체를 닫아야 한다. 그것과 함께 끝났을 때:"
*/
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
// http 클라이언트 헤더들, redirect방침, 다른 세팅들을 컨트롤하기 위해 Client를 생성한다.

client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
// 프록시들 TLS설정, keep-alives, compression, 그리고 다른 세팅들을 컨트롤하기 위해 Transport를 생성한다.

tr := &http.Transport{
	TLSClientConfig:    &tls.Config{RootCAs: pool},
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
/*
클라이언트와 Transports는 멀티고루틴에 의해 동시성사용에 안전하다. 
//Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.

ListenAndServe는 HTTP 서버를 주어진 주소와 핸들러와 함께 실행시킨다.
handler는 보통 DefaultServeMux를 사용하기위해 nil이다.
Handle과 HandleFunc는 DefaultServeMux에 handler를 추가한다.
*/
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
//커스텀서버를 만듦으로써 서버행동에 관한 더 많은 컨트롤이 가능하다.

s := &http.Server{
	Addr:           ":8080",
	Handler:        myHandler,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
log.Fatal(s.ListenAndServe())