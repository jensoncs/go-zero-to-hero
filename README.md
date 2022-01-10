  # Golang Zero To Hero

  Opensource programming language developed by google.

  ## Why Golang

  Golang is a new language and its supports concurrency through creating maximum goroutins(light weight threads).
  Golang is a statically typed language.

  ## Type of languages

  * Compiled language -: Basically, compiled code can be executed directly by the computer's CPU. That is, the executable code is specified in the CPU's "native" language (assembly language).
     -  Statically typed programming languages do type checking at compile-time (e.g. c, c++, go).
     -  staically typed will inference the variable while intializing ( eg : var i = true, automatically identify its a boolenan.)

  ## Compile stage

  [Compiler Design - Phases of Compiler](https://www.tutorialspoint.com/compiler_design/compiler_design_phases_of_compiler.htm)



  * Interpreted languages -: however must be translated at run-time from any format to CPU machine instructions. This translation is done by an interpreter.
     -  Dynamically typed programming languages do type checking at run-time (e.g. ruby, python).

  # Why Golang?

  * Golang is open-source but backed up by a large corporation
  * Many big companies used it for their projects including: Google, Uber, Gojek
  * It’s fast: to learn, compile, deploy and run
  * modern language
  * The reason why many developers adopt Golang is its built-in concurrency, which enables you to carry out many processes at the same time. Concurrency is happening via channels and goroutines while garbage collector provides principal support for such execution. Also, in Go concurrency is happening with the static execution speed of C or C++. This is not possible in many back-end languages, in which building process is done in a sequence.
  * cross platform support

  ## What is multithreading?

  A technique by which a single set of code can be used by several processors at different stages of execution.

  ## Thread v/s process?

  The typical difference is that threads (of the same process) run in a shared memory space, while processes run in separate memory spaces.Threads are an operating environment feature, rather than a CPU feature (though the CPU typically has operations that make threads efficient).

  * An executing instance of a program is called a process.
  * A thread is a subset of the process.

  ## What is goroutines?

  Goroutines are functions or methods that run concurrently with other functions or methods. Goroutines can be thought of as light weight threads. The cost of creating a Goroutine is tiny when compared to a thread. Hence its common for Go applications to have thousands of Goroutines running concurrently.

  * Goroutines are extremely cheap when compared to threads. They are only a few kb in stack size and the stack can grow and shrink according to needs of the application whereas in the case of threads the stack size has to be specified and is fixed.
  * The Goroutines are multiplexed to fewer number of OS threads. There might be only one thread in a program with thousands of Goroutines
  * Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate. We will discuss channels in detail in the next tutorial.

  Eg : if you want to execute 100 jobs at a time. cron jobs. If you want to server multiple requests at a time. if you want ot do multiple task at a time.

  ## What is channels?

  Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

  ## What is race condition?

  A race condition occurs when two or more threads can access shared data and they try to change it at the same time. Because the thread scheduling algorithm can swap between threads at any time, you don't know the order in which the threads will attempt to access the shared data. Therefore, the result of the change in data is dependent on the thread scheduling algorithm, i.e. both threads are "racing" to access/change the data.


  # Install and start with GO

  ## Mac
  ```
  brew install go

  export GOPATH=/User/jenson/go

  GOROOT is for compiler/tools that comes from go installation.
  GOPATH is for your own go projects / 3rd party libraries (downloaded with "go get").
  ```

  ### Hello world program

  ```
  package main

  import "fmt"

  func main() {
    fmt.Println("Hello World!")
  }
  ```

  # Go-Tour

  ## Packages

  a. Userdefined packages
  b. systems packges


  Every Go program is made up of packages.
  Programm start running in package main.
  Used for resuability, readability and maintance

  ```
  package main

  import "fmt"
  import "math/rand"

  func main() {
    fmt.Println("my favourite number is :", rand.Intn(100))
  }

  ```
  ## Imports

  This code groups the imports into a parenthesized, "factored" import statement.

  ```
  import "fmt"
  import "math"

  OR

  import (
    "fmt"
    "math"
  )
  ```

  ## Naming conventions

  ```
  file name -: file-name.go

  Exposing variables - Pizza

  local variables - pizza

  E.g. To access exposed variables

  func main() {
    fmt.Println(math.Pi)
  }
  ```

  ## Commands

  ```
  go run main.go - run

  go build - To build the program

  go clean - to clean all unused files.

  go doc fmt - to see the doc page

  go help - to see available commands with help
  ```
  ## Function

  Function can accept zero or more parameters. 

  ```
  package main

  import (
    "fmt"
  )

  ## Will accept two values (integrer), output also will be as an integer

  func add(x, y int) int {
    return x + y
  }

  ## Call add function

  func main() {
    fmt.Println(add(3, 5))       
  }
  ```

  ```
  package main

  import (
    "fmt"
  )

  ## Accept two string and output two string

  func reverse(x, y string) (string, string) {
    return y, x
  }

  func main() {
    fmt.Println(reverse("jenson", "hello"))
  }

  ```
  ## Variables

  The var statement declares a list of variables; as in function argument lists, the type is last.A var statement can be at package or function level.

  ```
  package main

  import "fmt"

  ## declaration boolean variables

  var c, python, java bool

  func main() {
    ## declaration integrer variable

    var i int
    fmt.Println(i, c, python, java)
  }
  ```
  Inside a function you can use := for variable declaration

  ```
  package main

  import "fmt"

  func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
  }
  ```

  When ever assigning an variable it will save it in memory.

  ## Type inference

  It will inference data types by default.


  variable multiple declaration
  -----------------------------
  ```
  var (
  var i = 10
  var name = "jenson"
  )
  ```
  ## Types
vatr golang
  1. bool 
  2. string
  3.  int8, int16, 1nt32

  *Note*: demo - rune

  ## signed and unsigned integer


  polymorphisam
  ------------

  Many forms 

  same function with multiple signature.

  ```
  add(a,.b) {
  }
  add(a,b,c) {
  }
  ```
  static binding - while compile time only it will identify the constanst.
  ```
  func add(a,b int) int {     - Function signature

  xxxx                       - Func body
  xxxx
  }
  ```
  ## Call by reference and call by value

  Call by reference is in func it will refere to the memory address

  func add(&a,&b)

  Call by value is in func values will be passed.

  func add(2,3)

  If any func outputing two and you need only one :- You can use blank identifier(_) to skip any of the output or assign the value to some variable.

  In func You can directly return name is called named return

  ```
  func add(a,b int) (sum) {
   sum = a+b
   return
  }
  ```
  ## Exported names

  Variables starting with capital letters are exported variable in a package.

  Eg: math.Pi


  ## Init Function

  func init() {

  }
  initializing tasks are managed in this section

  ## Control statements

  * If-else
  ```
  if condition {

  }else {

  }
  ```
  Note: In golang if the else is in second line what will happen? ( By default golang will add ; after all the lines and it will break)


  ## Loop


  for initialization;condition;increment
  ```
  for i=0;i<10;i++ {

  }
  ```
  ## Break

  If you want to break the loop


  ## Continue

  Check some condition if the condition is passed it will continue to next increment.


  ## Notes

  i++ == i =i+1 == i+   =1

  Note: postingrement(i++) and preingrement(++i)

  ## Switch

  To cases
  ```
  finger := 2

  switch(finger) {

  case 1:
    print("x")
  case 2: 
    print("sd")
  default:
   print "dd"
  }
  ```

  ## Expression less switch

  ```
  num := 2

  switch {

  case num > 1 && num <=4 :

    fmt.println("hello")
  }
  ```
  ## Fallthrough

  If the case is passed and you want to execute next line then you can have fallthrough.
  ```
  case 1:
   print("ddd")
   fallthrough
  case 2:
  ```

  ## Array

  Collection of data

  var a[3] int

  3 - length

  a := [3]int{1,2,3} or a := [...]int{1,2,3}

  ... - will dynamically identify the length


  Arry = data type + size (int + 3)

  a := [3]int{1,2,3}

  b := [5]int

  b = a ( will fail because arry type is a combination of size + data type)


  Arry in go are value type not refence type

  ## Range :

  it return index and value. using for looping

  ```
  for _, v := range array {
                  fmt.Println(v)
                  sum += v
          }
  ```

  ## Slices

  #### Array disadvantages

  * Array extension
  * merging array

  Slice is very convinient, flexible.

  wrapper over array.

  slice is  a refeerence to the array.
  ```
  [1:4] - from 1st position to less than 4

  c := []int{1,2,4}
  ```
  length = length of slice
  capacity = number of elements in the under line array from the index where slice created.

  ```
  a[...]{1,2,3,4,5,6}

  b := [1:4]

  slice length is 3 {2,3,4}
  slice capacity is 5 {2,3,4,5,6}
  ```

  ## Make

  make is a function to create

  make([]int,5-length,5- capacity)

  ## Append

  appending numbers to slice

  ```
  a := append(a, "5")
  ```

  appaend - func append(s[]T, x....T) [] {}

  ## Memory management

  Garbage collection (mark and sweep in java)

  copy function will help you to copy slice to another slice and you can remove old unsed slice to save memory.


  ## Slice Internal

  Slice internally is a struct

  type slice struct {

  length int
  capacity int
  zeroth element byte
  }

  ## Variadic function

  internally(...) will converted as slice.

  In variadic func you can pass slice.

  Eg: 

  ```
  nums := []int{1,2,3,4}

  func find(1,nums...) {
  }
  ```
  **Note**: When your doing slice operation you dont need to output im the function, slice is refering to array and changes will directly happening to array.

 ## Slice-Important

 You can do slice on string. Because string is made up of runes, runes are made up of bytes, so string is made up of bytes, A string is a buch of bytes; a slice of bytes

## Signed and Unsigned Numbers

`-,+` are signed numbers

only negative numbers are unsigned numbers.

How signed and unsiged working
------------------------------

Eg:

0000  1000
0001  1001
0002  1002

If the starting bit we are considing as a signed bit(+,- indicator flag) then the 0001 will be -ve 1 and 10001 will be the +1. If numbers are noted with signed flag called signed numbers. Signed numbers are reducing the capacity by half but its cover lots of other features.

Unsigned numbers are only in positive and it supports more positive values. Eg : 0001 is 1 and 1000 will be 8.

## Map

Map is a key value store.

p := make(map[string]int)

string - Key type

int - value type


value,ok := map["key"] -  Return the value + the existance of the key

Values out to map while itterating will come in different order.


**Delete**

You can delete keys from map.

delete(map["key"])

Map is a reference type

**Equality In Map**

Map by default not providing any quality fuction. You can impliment map equality using `range` over each key and copare.

`reflect` - Package is providing map equality check.


**Map Implimentation**

* why hashtables
   * hash tables serach is fast, most widely used data structure.
   * How hash search is fast?
    * Example you have a dictionary, you need to find the meaning of `sea` in the noraml array you have to itterate throgh each and every word in the dictionary to get the word `sea`. In hash table hash function will create a pattern and put those pattern into small buckets. Now when your searching for `sea` it knows in which bucket they have to search. They will get it from those bucket very fast instead of itterating through all.
* map is build upon hashtables
* How hash tables work? Hash tables are build upon array.
* Hashfunction function(algoritham) will create hashes for your input and find a common pattern
* All the common pattern they will be putting into bucket(arry)



## String

* String in go is slice of bytes.
* Utf-8 encoded
* Unicode complained

Usually each character takes `1 byte`, just think about special chartacter if it takes `2 bytes`, it cant print. To solve this problem we have Rune.

* Strings are immutable

## Rune

rune := []rune()

## Pointers 

default value of pointers is nil.

`&a` = will refer to the memory address of a.

`*a` - Will reference the value of a
`*` - Reference

## Structs

structures

user defined data types

collection of fields.

```
type employee struct {
name string
empId int
}
```

### Annonymous struct

```
var employee struct {
name string
empId int
}
```
if your not specifing the type command is called annonymous struct.

### Compare structs 

structs are comparable if the fields are same.

each of their fields are comparable struct variable are not comparable if the struct fields are not comparable.

## Nested struct

```
type Person struct {
name string
address address
}

type address struct {
city,state string
}
```

## Methods 

Methods are also functions. method is a function with special reciever type that is between func keyword and method name.

```
func (special reciver type) methodName() (outout type) {
}
```

## Interface [kind of polymorphisam]

* Why interface

Interface specify what methods a type should have and the type woll decide how to impliment this method.

defines behaviour of an object.

set of method signature.

E.g : Animal is a interface, animals can walk, eat ...etc

But tiger is an animal, human is an animal, the walking and eating styleof those animals are different 


how to impliment those with interface.

```
type Animal interface {
  eat() string
  walk() string
}

type Tiger struct {
}

type Human struct {
}

func (t Tiger) walk() string {
}

func (h Human) walk() string {

}

func (t Tiger) eat() string {
}

func (h Human) eat() string {

}
```

## Concurrency

contex switching

E.g: if you have a single core processor. the processor will be allocated to one process at a time. it will switch to other process depends. the switching happening in nano seconds.

User[job1]--doing something[job2]--->continue[job1]
at 6.01                6.02             6.03

### parallelism

user1[job1] -----> user[job2]
6.00                6.00

Go concurrency is achived throgh goroutines.

go routines are function or method that run concurrently with other methods or functions. goroutines are light weight threads.
The coast of creating goroutines are cheap compare to thread.

### Advantages

* Its cheap compare to thread, they are few KB in size and the stack can grow and shrink according to the need. In threads the stack size is fixed.

* goroutines are multiplexed to fewer no of os threads.

go routines communicating through channel.

channels will prevent race condition.

func main() - main goroutine.

`go funcName()`

## Channels

Channels are pipes to carry some data. It needs data types.

var a chan int

a := make(chan,int)


## Sending and receiving data from channel

```
a := make(chan,int)

a <- 1 = Writing integer 1 into the channel

data := <- a   = Reading data from channel a

```
Sends and recives to a channel is blocking by default, means when data is sent to a channel the control os blocked until some of the goroutine reads from the channel. Similarly when data is read from a channel the data read is blocked until some goroutines write data to the channel.


## Deadlock
In an operating system, a deadlock occurs when a process or thread enters a waiting state because a requested system resource is held by another waiting process, which in turn is waiting for another resource held by another waiting process.

A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.

```
func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
```
The program will get stuck on the channel send operation waiting forever for someone to read the value. Go is able to detect situations like this at runtime.

## Buffer channel


## Waiting group

## Worker pool
