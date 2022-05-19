package main

import (
	"fmt"
	"strings"

	"ldap/config"

	"github.com/go-ldap/ldap"
)

func main() {
	fmt.Println(strings.Split(config.MyDB, ",")[0])
	l, err := ldap.Dial("tcp", config.MyPath)
	if err != nil {
		fmt.Println("连接失败", err)
	}
	err = l.Bind(config.MyDB, config.MyPass)
	if err != nil {
		fmt.Println("管理员认证失败", err)
	}

	//  搜索所有用户
	searchRequest := ldap.NewSearchRequest(config.MyBaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",
		[]string{"dn", "cn", "uid", "mail"}, nil)
	search, err := l.Search(searchRequest)
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range search.Entries {
		fmt.Printf("dn：%s:  cn：%v uid： %v Email: %v\n", entry.DN, entry.GetAttributeValue("cn"), entry.GetAttributeValue("uid"), entry.GetAttributeValue("mail"))
	}

	// 创建新用户
	// addResponse := ldap.NewAddRequest("uid=java1,ou=people,dc=mingxinsk,dc=com", []ldap.Control{})
	// addResponse.Attribute("cn", []string{"java1"})
	// addResponse.Attribute("sn", []string{"java1"})
	// addResponse.Attribute("uid", []string{"java1"})
	// addResponse.Attribute("homeDirectory", []string{"/home/java1"})
	// addResponse.Attribute("loginShell", []string{"java1"})
	// addResponse.Attribute("gidNumber", []string{"501"})
	// addResponse.Attribute("uidNumber", []string{"8001"})
	// addResponse.Attribute("objectClass", []string{"shadowAccount", "posixAccount", "top", "inetOrgPerson"})
	// err = l.Add(addResponse)
	// if err != nil {
	// fmt.Println("创建用户失败")
	// }
	//
	// 随机给用户生成密码，并将新密码输出
	// https://juejin.cn/post/7030968139924013087#heading-9
	// passwordModifyRequest2 := ldap.NewPasswordModifyRequest("uid=java1,ou=people,dc=mingxinsk,dc=com", "", "")
	// passwordModifyResponse2, err := l.PasswordModify(passwordModifyRequest2)
	// if err != nil {
	// fmt.Println(err)
	// }
	// generatedPassword := passwordModifyResponse2.GeneratedPassword
	// fmt.Println("生成的密码: ", generatedPassword)
	// var ld func()
	// ld = LoadConfig
	fmt.Println(config.MyPort)
	// fmt.Printf("host: %v \n", MyHost)
	// fmt.Println(MyPort, MyDB, MyUser, MyPass, MyConf, MyPath)

}
