package main

import (
  "fmt"
  "scitum.com.mx/goarcsight"
)

func main(){

f1 := &goarcsight.Field{"hashSum","String","string", true }
f2 := &goarcsight.Field{"campaign","String","string", false }
f3 := &goarcsight.Field{"note","String","string", false }
arr := [3]goarcsight.Field{*f1,*f2,*f3}
var sarr []goarcsight.Field = arr[:]
list:= goarcsight.NewList("AAName-2","AAName-2","Lista de Prueba", "0WgEuT4UBABCALT6Xc3IXQA==",sarr)


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
*/
token:=goarcsight.Login("Z311510","T.m2{OY3}?pJS")

fmt.Println(token)
//fmt.Println(goarcsight.CreateList(list,token))
//fmt.Println(goarcsight.DeleteList("Hg5GBp44BABDOkX+KHULZDw==",token))
fmt.Println(goarcsight.AddHashList("HyWRcq44BABCYt3Eysk0iIA==","asu....","campana2","nota2",token))
}
