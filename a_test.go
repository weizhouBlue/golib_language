package golib_language_test
import (
	golib "github.com/weizhouBlue/golib_language"
	"fmt"
	"testing"
)



func Test_0(t *testing.T){

	var ok bool

	old:=map[string]int{ "old":100 }
	a :=[]interface{}{	old }
	b := make( []interface{} , len(a) )

	v:=golib.DeepCopy(a)
	b , ok = v.( []interface{} ) 
	if !ok {
		fmt.Println("failed to convert type")
	}
	fmt.Println(b) // [map[old:100]]

	old["old"]=200
	fmt.Println(a)  // [map[old:200]]
	fmt.Println(b)  // [map[old:100]]




}



func Test_01(t *testing.T){

	a:=[]interface{} {
		"avc",
		111 ,
		map[string] string {
			"m":"44",
			"n": "55" , 
		} ,
	}

	if result , e:=golib.SliceDeepCopy(a) ; e!=nil{
		fmt.Printf("failed %v ", e)
	}else{
		fmt.Println(result) // 

	}


	b:=map[string] string  {
			"m":"44",
			"n": "55" , 
	}

	if result , e:=golib.MapDeepCopy(b) ; e!=nil{
		fmt.Printf("failed %v ", e)
	}else{
		fmt.Println(result) // 

	}


}


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
	e:=map[string]interface{} { "a": []int{1,2} , "b":100  , "c":"mm" }
	f:=map[string] []int { "a": []int{1,2}  }
	if  k , err := golib.MapMinus( e , f ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  k ) // map[b:100 c:mm]
	}


}




func Test_3(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:= map[string]int { "a":1, "b":2  }
	b:=map[string]int { "a":1, "c":1 }
	c:=map[string]int { "a":1, "d":22 }

	if  commonMap , err := golib.MapGetCommonElement( a, b  , c ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  commonMap ) // map[a:1]
	}


	// ----------------
	d:= map[string]interface{} { "a": []int{1,2}, "b":2  }
	e:=map[string]interface{} { "a": []int{1,2}, "b":222  }
	if  commonMap , err := golib.MapGetCommonElement( d , e   ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  commonMap ) // map[a:[1 2]]
	}

}


func Test_4(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:= map[string]int { "a":1, "b":2  }
	b:=map[string]int { "a":1, "c":1 }

	if  vMap , err := golib.MapMinus( a, b   ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  vMap ) // map[b:2]
	}



	// ----------------
	c:= map[string]interface{} { "a": []int{1,2 }, "b": "hellp"  }
	d:=map[string]interface{}  { "a": []int{1,2 }, "b":100  }

	if  vMap , err := golib.MapMinus( c , d   ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  vMap ) // map[b:hellp]
	}



}





func Test_5(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:=[]interface{} { 1, "2", []int{1,2}  }

	if  existed , err := golib.SliceCheckElement( a ,  []int{1,2} ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  existed ) //true
	}

	golib.EnableLog=false
	// ----------------
	v:=map[string]interface{} { 
		"a": 1 ,
		"b": "2" ,
		"c": []int{1,2} ,
	}

	if  existed , err := golib.MapCheckElement( v , "c" , []int{1,2} ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  existed )  // true
	}

}





func Test_6(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:=[]interface{} { 1, "2", []int{1,2} , "2" , 1  , "2" }
	if  result , err := golib.SliceRmRepeatedElem( a  ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result )  // [[1 2] 1 2]
	}


	b:=[]int { 1, 2, 3 ,2 ,1  , 2 }
	if  result , err := golib.SliceRmRepeatedElem( b ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result ) // [3 1 2]
	}


	// ----------------
	v:=map[string]interface{} { 
		"a": 1 ,
		"b": "2" ,
		"c": []int{1,2} ,
		"d": 1 ,
		"e": 1 ,
	}
	if  result , err := golib.MapRmRepeatedElem( v  ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result ) //map[a:1 b:2 c:[1 2]]
	}

}




func Test_7(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:=[]interface{} { 1, "mm", []int{1,2} , "mm" , 1  }
	b:=[]int { 1 ,2, 3 }
	c:=[]string{ "mm" , "nn" , "mm" }
	if  result , err := golib.SliceAdd( a , b ,c  ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result ) // [[1 2] 1 2 3 nn mm]
	}

	// ----------------
	d:=map[string]interface{} { "a":1, "b":"mm", "c": []int{1,2}  }
	e:=map[string]int { "d":1 , "e":2 }
	f:=map[string]string { "f": "a" , "g":"b" }
	if  result , err := golib.MapAdd(  d , e ,f  ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result ) // map[a:1 b:mm c:[1 2] d:1 e:2 f:a g:b]
	}

}







func Test_8(t *testing.T){

	golib.EnableLog=false

	// ----------------
	a:=[]interface{} {
		map[string] string {
			"b2": "101",
		} , 
		map[string] interface{} {
			"b1": "101",
		} , 
	}

	if  result , err := golib.SliceToSliceMapStringString( a  ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result ) // [[1 2] 1 2 3 nn mm]
	}


	a=[]interface{} {
		map[string] interface{} {
			"a1": 100 ,
			"a2": "103",
		} , 
		map[string] string {
			"b2": "101",
		} , 
	}

	if  result , err := golib.SliceToSliceMapStringString( a  ) ; err!=nil {
		fmt.Println(  err )
		t.FailNow()
	}else{
		fmt.Println(  result ) // [[1 2] 1 2 3 nn mm]
	}




}



