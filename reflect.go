package golib_language
import(
	"fmt"
	"reflect"
	"strings"
	"runtime"
	"strconv"
)



var (
    EnableLog=false	
)



//====================log==========================

func getFileName( path string ) string {
    b:=strings.LastIndex(path,"/")
    if b>=0 {
        return path[b+1:]
    }else{
        return path
    }
}

func log( format string, a ...interface{} ) (n int, err error) {

    if EnableLog {

		prefix := ""
	    funcName,filepath ,line,ok := runtime.Caller(1)
	    if ok {
	    	file:=getFileName(filepath)
	    	funcname:=getFileName(runtime.FuncForPC(funcName).Name())
	    	prefix += "[" + file + " " + funcname + " " + strconv.Itoa(line) +  "]     "
	    }

        return fmt.Printf(prefix+format , a... )    
    }
    return  0,nil
}






//======================================


// 从各个切片中 提取出共同的元素，然后把这些共同的元素 组成新的切片 ，进行返回
// 输入的每个切片 可以是  []interface{} 
func SliceGetCommonElement( v ...interface{}  ) ( []interface{}  ,  error ){

	commonSlice:=[]interface{} {}

	if len(v)<1 {
		return nil , fmt.Errorf("slice number must be more than 2")
	}

	// find the slice with least element 
	var baseSlice interface{}
	leaseNum:=0
	for _ , v1 := range v {

		//check type
		if reflect.TypeOf(v1).Kind()!=reflect.Slice {
			return nil , fmt.Errorf("input is not slice : %v " , v1 )
		}

		v2:=reflect.ValueOf(v1)
		if v2.IsNil()==true  {
			return nil , fmt.Errorf("there is nil slice : %v " , v1 )
		}
		length:=v2.Len()
		// check empty slice
		if length==0  {
			return nil , fmt.Errorf("there is empty slice : %v " , v1 )
		}
		log("length=%v \n",length)

		if length <= leaseNum || leaseNum==0 {
			baseSlice=v1
			leaseNum=length
		}

	}
	log("the slice with lease element is %v \n" , baseSlice )


	// begin to find 
	t:=reflect.ValueOf(baseSlice)
NEXT_ELEMENT:
	for i:=0 ; i< t.Len() ; i++ {
		if t.Index(i).CanInterface()==false{
			return nil , fmt.Errorf("failed to interface{} : %v " , t.Index(i) )
		}
		v1:=t.Index(i).Interface()

		log("for %v\n" , v1 )

		//loop all elemet of all slice
NEXT_SLICE:		
		for _  , vSlice :=range v {
			p:=reflect.ValueOf(vSlice)
			for j:=0 ; j<p.Len() ; j++ {
				if p.Index(j).CanInterface()==false{
					return nil , fmt.Errorf("failed to interface{} : %v " , t.Index(j) )
				}
				v2:=p.Index(j).Interface()

				if reflect.TypeOf(v2).Kind() != reflect.TypeOf(v1).Kind() {
					log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
					continue 
				}
				if reflect.DeepEqual(v1, v2)==true{
					log("%v(%T) == %v(%T)" , v1 ,v1 ,v2, v2 )
					continue NEXT_SLICE 
				}else{
					log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
				}
			}
			log("no common element: %v  \n" , v1  )
			continue NEXT_ELEMENT
		}
		// find a common element
		commonSlice=append(commonSlice, v1 )
	}

	log("get common slice: %v \n" , commonSlice )

	return commonSlice, nil

}



// 从各个 map 中 提取出共同的元素，然后把这些共同的元素 组成新的 map ，进行返回
// 输入的每个map 可以是 map[string]interface{} 
func MapGetCommonElement( v ...interface{}  ) ( map[string]interface{}  ,  error ){

	commonMap:=map[string]interface{} {}

	if len(v)<1 {
		return nil , fmt.Errorf("map number must be more than 2")
	}

	// find the slice with least element 
	var baseMap interface{}
	leaseNum:=0
	for _ , v1 := range v {

		//check type
		if reflect.TypeOf(v1).Kind()!=reflect.Map {
			return nil , fmt.Errorf("input is not Map : %v " , v1 )
		}

		v2:=reflect.ValueOf(v1)
		if v2.IsNil()==true  {
			return nil , fmt.Errorf("there is nil Map : %v " , v1 )
		}
		length:=len(v2.MapKeys())
		// check empty Map
		if length==0  {
			return nil , fmt.Errorf("there is empty Map : %v " , v1 )
		}
		log("length=%v \n",length)

		if length <= leaseNum || leaseNum==0 {
			baseMap=v1
			leaseNum=length
		}

	}
	log("the map with lease element is %v \n" , baseMap )


	// begin to find 
	t:=reflect.ValueOf(baseMap)
NEXT_ELEMENT:
	for _ , v0 := range t.MapKeys() {
		if v0.Type().Kind() !=reflect.String {
			return nil , fmt.Errorf("key type is not string:  %v(%v) " , v0.Type() , v0 )
		}
		keyName:=v0.String()
		if t.MapIndex(v0).CanInterface()  ==false{
			return nil , fmt.Errorf("failed to interface{} : %v " , t.MapIndex(v0) )
		}
		v1:=t.MapIndex(v0).Interface()

		log("for %v\n" , v1 )

		//loop all elemet of all map
NEXT_SLICE:		
		for _  , vMap :=range v {
			p:=reflect.ValueOf(vMap)
			for _ , m := range p.MapKeys() {
				// check the key
				if m.Type().Kind()!=reflect.String {
					return nil , fmt.Errorf("key type is not string:  %v(%v) " , m.Type() , m )
				}
				vKeyName:=m.String()
				if keyName!=vKeyName {
					log("key(%v) != key(%v)" , keyName ,vKeyName )
					continue 
				}

				// check the vlaue
				if p.MapIndex(m).CanInterface()  ==false{
					return nil , fmt.Errorf("failed to interface{} : %v " , p.MapIndex(m) )
				} 
				v2:=p.MapIndex(m).Interface()
				if reflect.TypeOf(v2).Kind() != reflect.TypeOf(v1).Kind() {
					log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
					continue 
				}
				if reflect.DeepEqual(v1, v2)==true{
					log("%v(%T) == %v(%T)" , v1 ,v1 ,v2, v2 )
					continue NEXT_SLICE 
				}else{
					log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
				}
			}
			log("no common element: %v=%v  \n" ,keyName ,  v1  )
			continue NEXT_ELEMENT
		}
		// find a common element
		commonMap[keyName]=v1
	}

	log("get common map: %v \n" , commonMap )

	return commonMap, nil

}


//===================================

// subtrahend 切片 扮演 被减数 ， minuend切片 扮演 减数，实现集合的减法，在subtrahend中 去除所有minuend中出现的元素 ，最后返回新的切片
// subtrahend 和 minuend切片中的 所有成员的 可以是任何数据类型，且可以各不相同
func SliceMinus( subtrahend , minuend interface{}  ) (  output []interface{}  ,  err error ){

	//check type
	if reflect.TypeOf(subtrahend).Kind()!=reflect.Slice || reflect.TypeOf(minuend).Kind()!=reflect.Slice {
		return nil , fmt.Errorf("input is not slice "  )
	}

	if  minuend==nil || subtrahend==nil {
		return nil , fmt.Errorf("empty input " )
	}


	output=[]interface{} { }

	m:=reflect.ValueOf(subtrahend)
	n:=reflect.ValueOf(minuend)
NEXT_ELEMENT:	
	for i:=0 ; i<m.Len() ; i++ {
		if m.Index(i).CanInterface()==false{
			return nil , fmt.Errorf("failed to interface{} : %v " , m.Index(i) )
		}
		v1:=m.Index(i).Interface()

		for j:=0 ; j<n.Len() ; j++ {
			if n.Index(j).CanInterface()==false{
				return nil , fmt.Errorf("failed to interface{} : %v " , n.Index(j) )
			}
			v2:=n.Index(j).Interface()

			if reflect.TypeOf(v2).Kind() != reflect.TypeOf(v1).Kind() {
				log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
				continue NEXT_ELEMENT 
			}
			if reflect.DeepEqual(v1, v2)==true{
				log("%v(%T) == %v(%T)" , v1 ,v1 ,v2, v2 )
				continue NEXT_ELEMENT 
			}else{
				log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )				
			}
		}
		output=append(output , v1 )
	}

	return output, nil

}



// subtrahend 切片 扮演 被减数 ， minuend切片 扮演 减数，实现集合的减法，在subtrahend中 去除所有minuend中出现的元素 ，最后返回新的map
// subtrahend 和 minuend  可以是 map[string]interface{}
func MapMinus( subtrahend , minuend interface{}  ) (  map[string]interface{}  ,   error ){

	//check type
	if reflect.TypeOf(subtrahend).Kind()!=reflect.Map || reflect.TypeOf(minuend).Kind()!=reflect.Map {
		return nil , fmt.Errorf("input is not Map "  )
	}

	if  minuend==nil || subtrahend==nil {
		return nil , fmt.Errorf("empty input " )
	}


	output:=map[string]interface{} { }

	m:=reflect.ValueOf(subtrahend)
	n:=reflect.ValueOf(minuend)
NEXT_ELEMENT:	
	for _ , k1 := range m.MapKeys() {
		if k1.Type().Kind() !=reflect.String {
			return nil , fmt.Errorf("key type is not string:  %v(%v) " , k1.Type() , k1 )
		}
		keyName:=k1.String()
		if m.MapIndex(k1).CanInterface()  ==false{
			return nil , fmt.Errorf("failed to interface{} : %v " , m.MapIndex(k1) )
		}
		v1:=m.MapIndex(k1).Interface()


		for _ , k2 := range n.MapKeys() {
			//check key
			if k2.Type().Kind() !=reflect.String {
				return nil , fmt.Errorf("key type is not string:  %v(%v) " , k2.Type() , k2 )
			}
			vKeyName:=k2.String()
			if keyName!=vKeyName {
				log("key(%v) != key(%v)" , keyName ,vKeyName )
				continue 
			}

			// check the vlaue
			if n.MapIndex(k2).CanInterface()  ==false{
				return nil , fmt.Errorf("failed to interface{} : %v " , n.MapIndex(k2) )
			} 
			v2:=n.MapIndex(k2).Interface()
			if reflect.TypeOf(v2).Kind() != reflect.TypeOf(v1).Kind() {
				log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
				continue 
			}
			if reflect.DeepEqual(v1, v2)==true{
				log("%v(%T) == %v(%T)" , v1 ,v1 ,v2, v2 )
				continue NEXT_ELEMENT 
			}else{
				log("%v(%T) != %v(%T)" , v1 ,v1 ,v2, v2 )
			}
		}
		output[keyName]=v1 
	}

	return output, nil

}




//===================================

// chceckedSlice 可以是 []interface{} , checkedElement可以是任意类型的值
// 函数检查 chceckedSlice 中是否存在 和checkedElement值相同的元素
func SliceCheckElement( chceckedSlice interface{} , checkedElement interface{} ) (  exist bool  ,  err error ){

	exist=false
	err=nil

	if  chceckedSlice==nil || checkedElement== nil {
		err=fmt.Errorf("empty input  " )
		return 
	}

	//check type
	if reflect.TypeOf(chceckedSlice).Kind()!=reflect.Slice  {
		err=fmt.Errorf("input is not slice "  )
		return
	}


	m:=reflect.ValueOf(chceckedSlice)
NEXT_ELEMENT:	
	for i:=0 ; i<m.Len() ; i++ {
		if m.Index(i).CanInterface()==false{
			err=fmt.Errorf("failed to interface{} : %v " , m.Index(i) )
			return
		}
		v1:=m.Index(i).Interface()

		if reflect.TypeOf(checkedElement).Kind() != reflect.TypeOf(v1).Kind() {
			continue NEXT_ELEMENT 
		}

		if reflect.DeepEqual(v1, checkedElement)==true{
			exist=true
			return 
		}

	}

	return 

}





// chceckedMap 可以是 map[string]interface{} 
// 函数检查 chceckedMap 中是否存在  key为checkedKey、值为checkedValue 的元素
func MapCheckElement( chceckedMap interface{} , checkedKey string , checkedValue interface{} ) (  exist bool  ,  err error ){

	exist=false
	err=nil

	if  chceckedMap==nil || checkedValue== nil || len(checkedKey)==0 {
		err=fmt.Errorf("empty input  " )
		return 
	}

	//check type
	if reflect.TypeOf(chceckedMap).Kind()!=reflect.Map  {
		err=fmt.Errorf("input is not map "  )
		return
	}

	m:=reflect.ValueOf(chceckedMap)
NEXT_ELEMENT:	
	for _ , k1 := range m.MapKeys() {
		if k1.Type().Kind() !=reflect.String {
			err=fmt.Errorf("key type is not string:  %v(%v) " , k1.Type() , k1 )
			return
		}
		keyName:=k1.String()
		if m.MapIndex(k1).CanInterface()  ==false{
			err=fmt.Errorf("failed to interface{} : %v " , m.MapIndex(k1) )
			return
		}
		v1:=m.MapIndex(k1).Interface()

		if keyName!=checkedKey {
			continue NEXT_ELEMENT
		}

		if reflect.TypeOf(checkedValue).Kind() != reflect.TypeOf(v1).Kind() {
			continue NEXT_ELEMENT 
		}

		if reflect.DeepEqual(v1, checkedValue)==true{
			exist=true
			return 
		}

	}

	return 

}



