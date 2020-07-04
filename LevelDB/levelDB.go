package LevelDB

import "fmt"

/*
LevelDB:
1.就是将区块做持久化的工具；工作流程是先存储到磁盘中之后再存储到内存中，便于高效率的取数据
2.是单机版的k-v存储系统
3.key和value 都是任意的字节数组，支持内存和持久化存储
4.多操作可以当成一次原子操作，支持事务
5.包含基本的数据操作接口：Put(key value)、Get(key)、Delete(key)
局限：
1.非关系型数据库，不支持sql和索引
2.同一时间只支持单进程（多线程）访问db
3.不支持客户端-服务器模型，用户需要自己封装
*/

//go实现简版LevelDB
type DB struct {
	Path string
	Data map[string][]byte
}


//模拟开启连接
func New(path string)(*DB,error){
	return &DB{
		Path: path,
		Data: make(map[string][]byte),
	},nil
}

//模拟关闭连接
func (d *DB)Close()error{
	return nil
}

func (d *DB)Put(key ,value []byte)error{
	d.Data[string(key)]=value
	return nil
}

func (d *DB)Get(key []byte)(value []byte,err error){
	if v,ok:=d.Data[string(key)];ok {
		return v,nil
	}
	return nil,fmt.Errorf("the key not exist")
}

func (d *DB)Delete(key []byte)error{
	if _,ok:=d.Data[string(key)];ok {
		delete(d.Data,string(key))
		return nil
	}
	return fmt.Errorf("the key not exist")
}

//模拟遍历取值
func (d *DB) Iterator()*Iterator{
	return NewIterator(d.Data)
}


