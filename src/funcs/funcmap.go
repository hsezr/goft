package funcs
import (
 
     "html/template"
 
 )
var FuncMap=map[string]interface{}{
 
   "Abc" :func () string {
	return "abc"
},
  
   "Strong" :func (txt string) template.HTML {
	return template.HTML("<strong>" + txt + "</strong>")
},
  
   "Test" :func () string {
	return "test"
},
  
}
			