# 单例模式
确保某个对象在系统运行的过程中只被实例化一次。  
Go语言实现单例模式有四种方式，分别是懒汉模式、饿汉模式、双重检查、sync.Once。
## 懒汉模式（非线程安全）
缺点:多线程调用时，可能会多次实例化对象。
```go
var instance *People
var lock sync.Mutex

type People struct {
    Name string
}

func GetInstance(name string) *People {
lock.Lock()
defer lock.Unlock()
if instance == nil {
    instance = &People{Name: name}
}
return instance
}
```
## 懒汉模式（线程安全）
缺点:每次调用都进行加锁操作，效率不高。
```go
var instance *People
var lock sync.Mutex

type People struct {
    Name string
}

func GetInstance(name string) *People {
lock.Lock()
defer lock.Unlock()
if instance == nil {
    instance = &People{Name: name}
}
return instance
}
```
## 饿汉模式
在导入包的同时会创建对象，并持续占有内存。
```go
type People struct {
	Name string
}

func init(){
	instance = &People{Name: "Tom"}
}

func GetInstance() *People {
	return instance
}
```
## 双重检查
在懒汉模式(线程安全)的基础上进行优化，减少加锁的操作。
```go
var instance *People
var lock sync.Mutex

type People struct {
	Name string
}

func GetInstance() *People {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &People{Name: "Tom"}
		}
		lock.Unlock()
	}
	return instance
}
```
## sync.Once（推荐）
sync.Once本质上也是双重检查，但是写法会比自己写双重检查更简洁。
```go
var instance *People
var once sync.Once

type People struct {
    Name string
}

func GetInstance(name string) *People {
once.Do(func() {
    instance = &People{Name: name}
})
return instance
}
```