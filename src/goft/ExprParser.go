package goft

import (
	"bytes"
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

const (
	VarPattern       = `[0-9a-zA-Z_]+`
	CompareSign      = ">|>=|<=|<|==|!="
	CompareSignToken = "gt|ge|le|lt|eq|ne"
	ComparePattern   = `^(` + VarPattern + `)\s*(` + CompareSign + `)\s*(` + VarPattern + `)\s*$`
)

//可比较表达式 解析类， 譬如a>3   b!=4 a!=n    a>3  [gt .a  3]
type ComparableExpr string
func (this ComparableExpr) filter() string {
	reg, err := regexp.Compile(ComparePattern) //Compile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本。
	if err != nil {
		return ""
	}
	ret := reg.FindStringSubmatch(string(this))
	if ret != nil && len(ret) == 4 {
		token := getCompareToken(ret[2])
		if token == "" {
			return ""
		}

		return fmt.Sprintf("%s %s %s", token, parseToken(ret[1]), parseToken(ret[3]))
	}

	return ""
}

//根据比较符，获取对应的gt|ge|le|lt|eq|ne
func getCompareToken(sign string) string {
	for index, item := range strings.Split(CompareSign, "|") {
		if item == sign {
			return strings.Split(CompareSignToken, "|")[index]
		}
	}

	return ""
}

func parseToken(token string) string {
	if IsNumric(token) {
		return token
	} else {
		return "." + token
	}
}

func IsComparableExpr(expr string) bool {
	reg, err := regexp.Compile(ComparePattern)
	if err != nil {
		return false
	}

	return reg.MatchString(expr)
}

type Expr string

func ExecExpr(expr Expr, data map[string]interface{}) (string, error) {
	tpl := template.New("expr").Funcs(map[string]interface{} {
		"echo": func(params ...interface{}) interface{} {
			return fmt.Sprintf("echo:%v", params[0])
		},
	})

	t, err := tpl.Parse(fmt.Sprintf("{{%s}}",expr))
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}