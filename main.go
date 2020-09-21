package main

import (
    "fmt"
    "log"
    "net/http"
)
// 루트 URL에 대한 모든 요청을 처리 할 홈페이지 
func homePage(w http.ResponseWriter, r *http.Request){ 
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
// URL경로와 일치하는 핸들 요청 함수 (나머지 함수가 정의되어 있는)
func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}
// 나머지 서버의 진입 역할을 하는 메인 함수
func main() { 
    handleRequests()
}