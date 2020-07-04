package LevelDB


//迭代器接口
type IteratorInter interface {
	//判断是否有下个值
	Next()bool
	//遍历键值
	Key()[]byte
	Value()[]byte
}

//键值对结构体
type Pair struct {
	Key []byte
	Value []byte
}


//迭代器的结构体
type Iterator struct {
	data []Pair
	index int
	length int
}

//创建默认的迭代器
func NewIterator(data map[string][]byte)*Iterator{
	newIterator:=new(Iterator)
	newIterator.index=-1
	newIterator.length=len(data)
	for k,v:=range data {
		tmp:=Pair{
			Key: []byte(k),
			Value: v,
		}
		newIterator.data=append(newIterator.data,tmp)
	}
	return newIterator
}

//判断是否存在下一个
func (i *Iterator)Next()bool{
	if i.index<len(i.data)-1{
		i.index++
		return true
	}
	return false
}

//获取key
func (i *Iterator)Key()[]byte{
	if i.index==-1||i.index>len(i.data){
		panic("error:index out of bound")
	}
	bytes := i.data[i.index].Key
	return bytes
}
//获取value
func (i *Iterator)Value()[]byte{
	if i.index>len(i.data) {
		panic("error:index out of bound")
	}
	bytes := i.data[i.index].Value
	return bytes
}
