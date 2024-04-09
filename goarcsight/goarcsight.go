package goarcsight

import (
  "bytes"
  "encoding/json"
  "time"
//  "errors"
  "fmt"
  "net/http"
  "net/url"
  "crypto/tls"
  "io"
)

type HashRecord struct{
  Hash string `json:"hashSum"`  
  Campaign string `json:"campaign"`  
  Note string `json:"note"`  
}

type Bear struct{
  Token string `json:"log.return"`
}

type Response  struct  {
  Auth Bear `json:"log.loginResponse"`
}

type Field struct{
  Name string `json:"name"`
  Type string `json:"type"`
  SubType string `json:"subType"`
  Key bool `json:"key"`
}

type List struct
{
  Name string `json:"name"`
  Alias string `json:"alias"`
  Description string `json:"description"`
  Capacity int `json:"capacity"`
  EntryTimeToLive int `json:"entryTimeToLive"`
  MultiMap  bool `json:"multiMap"`
  PartialCache bool `json:"partialCache"`
  TimePartitioned bool `json:"timePartitioned"`
  ActiveListType string `json:"activeListType"`
  CaseSensitiveType string `json:"caseSensitiveType"`
  Fields []Field `json:"fields"`
  GroupId string `json:"groupId"`
}

type JsonArray struct{
  Fields []string `json:"fields"`
}

type JsonArrayVal struct{
  Fields []string `json:"fields"`
  Entries []JsonArray `json:"entries"`
}

const url_base= "https://brux2806.claro.com.br:8443"
const url_activelist = "/detect-api/rest/v1/activelists"




func NewList(name string, alias string, desc string, father string, sarr []Field)*List{
list:= &List{
  Capacity: 7000,
  EntryTimeToLive: 0,
  MultiMap: false,
  PartialCache: true,
  TimePartitioned: false,
  ActiveListType: "FIELD_BASED",
  CaseSensitiveType: "KEY_CASE_INSENSITIVE",

  Name: name,
  Alias: alias,
  Description: desc,
  Fields: sarr,
  GroupId: father,
}

  return list

}


func Login(user string, passwd string)string{

  url_login := "/www/core-service/rest/LoginService/login?login=" + user +  "&password="
  url_encode_pass := passwd
  url_ok := url_base + url_login + url.QueryEscape(url_encode_pass)

  fmt.Println(url_ok)
  req, err := http.NewRequest("GET",url_ok,nil)
  
  if err != nil {
    fmt.Println("Build req error: %s",err)
  }

  req.Header.Set("Accept","application/json")

  tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  client:= http.Client{Timeout: 10 * time.Second, Transport: tr}
  res, err:= client.Do(req)
  if err != nil {
    fmt.Println("impossible to send request: %s",err)
  }
  resBody,err := io.ReadAll(res.Body)
  fmt.Println("Status code: ", res.StatusCode)
  //resJson,_:= json.Marshal(string(resBody))
  //fmt.Println("Response: ", string(resBody))
  //fmt.Println("Response: ", string(resJson))
  var resp Response
  json.Unmarshal(resBody, &resp)
  //fmt.Println("Response: ", resp.Auth.Token)
  return string(resp.Auth.Token)
}



func AddHashList(id string, hash string, campain string, note string, token string)(int,string){

  url_ok := url_base + url_activelist + "/" + url.QueryEscape(id) + "/entries"
 
  //jsondat := &JsonArrayVal{Fields: []string{"hashSum", "campaign", "note"},Entries: []string{hash,campain,note}}
  jsondat := &JsonArrayVal{Fields: []string{"hashSum", "campaign", "note"}}
  jsondat.Entries=[]JsonArray {JsonArray{[]string{hash,campain,note}}}
  encjson, _ := json.Marshal(jsondat)
  fmt.Println(string(encjson))

  fmt.Println(url_ok)
  req, err := http.NewRequest("POST",url_ok,bytes.NewReader(encjson))
  req.Header.Set("Content-Type","application/json")
  req.Header.Set("Accept","application/json")
  req.Header.Set("Authorization","Bearer " + token)

  tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  client:= http.Client{Timeout: 10 * time.Second, Transport: tr}
  res, err:= client.Do(req)
  if err != nil {
    fmt.Println("impossible to send request: ",err)
    return -1,"Error enviando request: " + err.Error()
  }
  resBody,err := io.ReadAll(res.Body)

  if err != nil {
    return res.StatusCode,err.Error()
  }else{
    return res.StatusCode,string(resBody)
  }
  //fmt.Println(string(resBody))
}

func DeleteList(id string, token string)(int,string){

  url_ok := url_base + url_activelist + "/" + url.QueryEscape(id)

  fmt.Println(url_ok)
  req, err := http.NewRequest("DELETE",url_ok,nil)
  req.Header.Set("Content-Type","application/json")
  req.Header.Set("Accept","application/json")
  req.Header.Set("Authorization","Bearer " + token)

  tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  client:= http.Client{Timeout: 10 * time.Second, Transport: tr}
  res, err:= client.Do(req)
  if err != nil {
    fmt.Println("impossible to send request: ",err)
    return -1,"Error enviando request: " + err.Error()
  }
  resBody,err := io.ReadAll(res.Body)

  if err != nil {
    return res.StatusCode,err.Error()
  }else{
    return res.StatusCode,string(resBody)
  }
  //fmt.Println(string(resBody))

}

func CreateList(list *List, token string )(int,string){

  url_ok := url_base + url_activelist

  fmt.Println(url_ok)

  rbody,_ := json.MarshalIndent(list,"","")
  req, err := http.NewRequest("POST",url_ok,bytes.NewReader(rbody))
  
  if err != nil {
    fmt.Println("Build req error: %s",err)
  }

  req.Header.Set("Content-Type","application/json")
  req.Header.Set("Accept","application/json")
  req.Header.Set("Authorization","Bearer " + token)

  tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  client:= http.Client{Timeout: 10 * time.Second, Transport: tr}
  res, err:= client.Do(req)
  if err != nil {
    fmt.Println("impossible to send request: ",err)
    return -1,err.Error()
  }
  resBody,err := io.ReadAll(res.Body)
  //fmt.Println(string(resBody))
  //fmt.Println(res.StatusCode)
  return res.StatusCode,string(resBody)
}

/*
func main(){

f1 := &Field{"hashSum","String","string", true }
f2 := &Field{"campaign","String","string", false }
f3 := &Field{"note","String","string", false }
arr := [3]Field{*f1,*f2,*f3}
var sarr []Field = arr[:]
list:= newList("AAName-2","AAName-2","Lista de Prueba", "0WgEuT4UBABCALT6Xc3IXQA==",sarr)


_=list


/*
lj,_ := json.MarshalIndent(list,"","  ")

fj1,_ := json.MarshalIndent(f1,"","  ")
fj2,_ := json.MarshalIndent(f2,"","  ")
fj3,_ := json.MarshalIndent(f3,"","  ")

fmt.Println(string(fj1))
fmt.Println(string(fj2))
fmt.Println(string(fj3))
fmt.Println(string(lj))
//--end comment
token:=login("Z311510","T.m2{OY3}?pJS")

fmt.Println(token)
//fmt.Println(createList(list,token))
//fmt.Println(deleteList("Hg5GBp44BABDOkX+KHULZDw==",token))
fmt.Println(addHashList("Hc5Ctp44BABDzX52QsHQ4WQ==","hashsumdfsfsfsfs2","campana2","nota2",token))
}

*/
