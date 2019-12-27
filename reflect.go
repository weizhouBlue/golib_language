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
// 输入的每个切片，其中的 成员 可以是任意数据类型，并且，所有切片的、所有成员的 数据类型 可以各不相同
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




