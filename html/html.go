package html

import (
	"strings"
	"fmt"
)

type _ElementBuilder interface {
	String() string
}

type _Element struct {
	id string
	name string
	text string
	element string
	attributes map[string]string
	styles map[string]string
	classes []string
	childs []interface{}
}

func (this *_Element) _init(el string) {
	this.attributes = make(map[string]string)
	this.childs = []interface{}{}
	this.classes = []string{}
	this.styles = make(map[string]string)
	this.element = el
}

func Element(custom string) *_Element {
	el := new(_Element)
	el.element = custom
	return el
}

func (this *_Element) Name(n string) *_Element {
	this.name = n
	return this
}

func (this *_Element) Id(n string) *_Element {
	this.id = n
	return this
}

func (this *_Element) Text(n string, args ...interface{}) *_Element {
	this.text = fmt.Sprintf(n, args...)
	return this
}

func (this *_Element) Styles(n string) *_Element {

	vs := strings.Split(n, ";")

	for _, v := range vs {
		if len(strings.TrimSpace(v)) > 0 && strings.Contains(v, ":") {
			vss := strings.Split(v, ":")

			if len(vss) == 2 && len(strings.TrimSpace(vss[0])) > 0 && len(strings.TrimSpace(vss[1])) > 0 {
				this.styles[vss[0]] = vss[1] 
			}
		}
	}

	
	return this
}

func (this *_Element) Style(k string, v string) *_Element {
	this.styles[k] = v
	return this
}


func (this *_Element) Class(n string) *_Element {

	vs := strings.Split(n, " ")

	for _, v := range vs {
		if len(strings.TrimSpace(v)) > 0  {
			this.classes = append(this.classes, v)
		}
	}

	return this
}

func (this *_Element) Child(vs ...interface{}) *_Element {
	for _, v := range vs {
		this.childs = append(this.childs, v)
	}
	return this
}

func (this *_Element) Attr(n string, v string) *_Element {
	this.attributes[n] = v
	return this
}

func (this *_Element) Val(val string) *_Element {
	this.attributes["value"] = val
	return this
}

func (this *_Element) Type(val string) *_Element {
	this.attributes["type"] = val
	return this
}

func (this *_Element) String() string {
	return this.elementAsHtlm(this)	
}

func (this *_Element) elementAsHtlm(child *_Element) string{

	var b strings.Builder

  fmt.Fprintf(&b, "<%v", child.element)

  if len(strings.TrimSpace(this.id)) > 0{
  	fmt.Fprintf(&b, " id='%v'", this.id)
  }

  if len(strings.TrimSpace(this.name)) > 0 {
  	fmt.Fprintf(&b, " name='%v'", this.name)
  }

  if len(child.classes) > 0 {
  	fmt.Fprintf(&b, " class='")
	  for _, v := range child.classes {
			fmt.Fprintf(&b, "%v ", v)	  	
	  } 
	  
	  fmt.Fprintf(&b, "'")
	}


  if len(child.styles) > 0 { 
  	fmt.Fprintf(&b, " style='")
	  for k, v := range child.styles {
			fmt.Fprintf(&b, "%v:%v;", k, v)	  	
	  } 

	  fmt.Fprintf(&b, "'")
  }


  for k, v := range child.attributes {
		fmt.Fprintf(&b, " %v='%v'", k, v)	  	
  } 

	fmt.Fprintf(&b, ">")

	if len(child.text) > 0 {
		fmt.Fprintf(&b, "\n\t%v", child.text)
	}

	for _, v := range child.childs {

		if builder, ok := v.(_ElementBuilder); ok {
			fmt.Fprintf(&b, "\n\t%v", builder.String())	
		} 
	}

	fmt.Fprintf(&b, "\n</%v>", child.element)	 	

	return b.String()
}

type _Html struct {
	_Element
}

func Html() *_Html{
	v := new(_Html)
	v._init("html")
	return v
}

type _Head struct {
	_Element
}

func Head() *_Html{
	v := new(_Html)
	v._init("head")
	return v
}

type _Body struct {
	_Element
}

func Body() *_Body{
	v := new(_Body)
	v._init("body")
	return v
}

type _Script struct {
	_Element
}

func Script() *_Body{
	v := new(_Body)
	v._init("script")
	return v
}

func (this *_Script) Src(v string) *_Script{
	this.attributes["src"] = v
	return this
}

func (this *_Script) Type(v string) *_Script{
	this.attributes["type"] = v
	return this
}

func (this *_Script) JS() *_Script{
	this.attributes["type"] = "text/javascript"
	return this
}

type _Link struct {
	_Element
}

func Link() *_Link{
	v := new(_Link)
	v._init("link")
	return v
}

func (this *_Link) Rel(v string) *_Link{
	this.attributes["rel"] = v
	return this
}

func (this *_Link) Href(v string) *_Link{
	this.attributes["href"] = v
	return this
}

func (this *_Link) CSS() *_Link{
	this.attributes["type"] = "stylesheet"
	return this
}


type _Div struct {
	_Element
}

func Div() *_Div{
	v := new(_Div)
	v._init("div")
	return v
}

type _Span struct {
	_Element
}

func Span() *_Span{
	v := new(_Span)
	v._init("span")
	return v
}

type _Hr struct {
	_Element
}

func Hr() *_Hr{
	v := new(_Hr)
	v._init("hr")
	return v
}

type _Strong struct {
	_Element
}

func Strong() *_Strong{
	v := new(_Strong)
	v._init("strong")
	return v
}

type _Input struct {
	_Element
}

func Input() *_Input{
	v := new(_Input)
	v._init("input")
	v.attributes["type"] = "text"
	return v
}

func (this *_Input) Type(v string) *_Input {
	this.attributes["type"] = v
	return this
}

type _TextArea struct {
	_Element
}

func TextArea() *_TextArea{
	v := new(_TextArea)
	v._init("textarea")
	return v
}

func (this *_TextArea) Rows(i int) *_TextArea {
	this.attributes["rows"] = fmt.Sprintf("%v", i)
	return this
}

func (this *_TextArea) Columns(i int) *_TextArea {
	this.attributes["columns"] = fmt.Sprintf("%v", i)
	return this
}

type _Option struct {
	_Element	
}

func Option(val string, text string) *_Option{
	v := new(_Option) 
	v._init("option")
	v.Val(val)
	v.Text(text)
	return v
}

type _Select struct {
	_Element
}

func Select() *_Select{
	v := new(_Select)
	v._init("select")
	return v
}

func (this *_Select) Options(vs ...*_Option) *_Select {
	for _, v := range vs {
		this.childs = append(this.childs, v)
	}
	return this
}

type _Li struct {
	_Element
}

func Li() *_Li{
	v := new(_Li)
	v._init("li")
	return v
}

type _Ul struct {
	_Element
}

func Ul() *_Ul{
	v := new(_Ul)
	v._init("ul")
	return v
}

type _Img struct {
	_Element
	//src string
}

func Img() *_Img{
	v := new(_Img)
	v._init("img")
	return v
}

func (this *_Img) Src(src string) *_Img {
	this.attributes["src"] = src
	return this
}

type _A struct {
	_Element
}

func A() *_A{
	v := new(_A)
	v._init("a")
	return v
}

func (this *_A) Href(src string) *_A {
	this.attributes["href"] = src
	return this
}

type _P struct {
	_Element
}

func P() *_P{
	v := new(_P)
	v._init("p")
	return v
}

type _Button struct {
	_Element
}

func Button() *_Button{
	v := new(_Button)
	v._init("button")
	v.attributes["type"] = "button"
	return v
}

type _Table struct {
	_Element
}

func Table() *_Table{
	v := new(_Table)
	v._init("table")
	return v
}

type _THead struct {
	_Element
}

func THead() *_THead{
	v := new(_THead)
	v._init("thead")
	return v
}

type _TBody struct {
	_Element
}

func TBody() *_TBody{
	v := new(_TBody)
	v._init("thead")
	return v
}

type _TFoot struct {
	_Element
}

func TFoot() *_TFoot{
	v := new(_TFoot)
	v._init("tfoot")
	return v
}

type _Tr struct {
	_Element
}

func Tr() *_Tr{
	v := new(_Tr)
	v._init("tr")
	return v
}

type _Th struct {
	_Element
}

func Th() *_Th{
	v := new(_Th)
	v._init("td")
	return v
}

type _Td struct {
	_Element
}

func Td() *_Td{
	v := new(_Td)
	v._init("td")
	return v
}