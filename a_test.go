package golib_language_test
import (
	golib "github.com/weizhouBlue/golib_language"
	"fmt"
	"testing"
)




func Test_1(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:=[]int { 1, 2, 3 ,4 }
	b:=[]int { 5, 1 , 4}
	c:=[]int { 6, 1 , 4}
	if  commonSlice , err := golib.SliceGetCommonElement( a, b  , c ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  commonSlice ) // [1 , 4]
	}

	golib.EnableLog=false

	// ----------------
	e:=[]interface{} {  10 , "abc" , []int { 1,2 } ,  map[string]int { "b":200 }  }
	f:=[]interface{} { 300 , "abc" ,  []int { 1,2 } }
	if  commonSlice , err := golib.SliceGetCommonElement( e , f ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  commonSlice )  // [abc [1 2]]
	}

	// ----------------
	g:=[]interface{} {  map[string]int { "a":100 } ,  map[string]int { "b":200 }  }
	h:=[]interface{} { map[string]int { "c":100 } }
	if  commonSlice , err := golib.SliceGetCommonElement( g , h ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  commonSlice )  // [ ]
	}

}


func Test_2(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:=[]int { 1, 2, 3 ,4 }
	b:=[]int { 5, 1 , 4}
	if  k , err := golib.SliceMinus( a, b ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  k ) // [2 3]
	}


	// ----------------
	c:=[]interface{} { 1 , "abc" , []int{1,2} , map[string]int{"a":100} }
	d:=[]string { "abc"  }
	if  k , err := golib.SliceMinus( c, d ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  k ) // [1 [1 2] map[a:100]]
	}


	// ----------------
	e:=[]interface{} { 1 , "abc" , []int{1,2} , map[string]int{"a":100} }
	f:=[]interface{} { []int{1,2}  }
	if  k , err := golib.SliceMinus( e , f ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  k ) // [1 abc map[a:100]]
	}


}
