package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Println("hello world")
	fmt.Println(a + b)
}

// os包下的函数,使用前需要加os. ,Close函数除外
/*--------1
func Open
func Open(name string) (file *File, err error)
Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符
具有O_RDONLY模式。如果出错，错误底层类型是*PathError。
使用说明：
	a, err :=os.Open(src) //a需要备用则写,否则用_ 忽略
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
		注意:Open和OpenFile操作以后通常会用defer功能关闭,defer file_name.Close(),关闭,否则会出现内存占用/泄漏
*/

/*--------2
func OpenFile
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的
选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用
于I/O。如果出错，错误底层类型是*PathError。
使用说明：
	c, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}
		const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
		注意:Open和OpenFile操作以后通常会用defer功能关闭,defer file_name.Close(),关闭,否则会出现内存占用/泄漏
*/

/*--------3
func (*File) Close
func (f *File) Close() error
Close关闭文件f，使文件不能用于读写。它返回可能出现的错误。
*/

/*--------4
func ReadFile-----等价于原 ioutil.ReadFile方法,ioutil.ReadFile从go 1.16废弃了
func ReadFile(filename string) ([]byte, error)
ReadFile 从filename指定的文件中读取数据并返回文件的内容。成功的调用返回的err为nil而非EOF。
因为本函数定义为读取整个文件，它不会将读取返回的EOF视为应报告的错误。
使用说明:
	content, err := os.ReadFile(file) //隐式的打开和关闭文件,用于小文件
	if err != nil {
		fmt.Println("read file err=", err)
	}
*/

/*--------5
func WriteFile-----等价于原 ioutil.WriteFile方法,ioutil.WriteFile从go 1.16废弃了
func WriteFile(filename string, data []byte, perm os.FileMode) error
函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
使用说明:
//读取的数据data写到了新的文件file2中，全部覆盖写，file2原先的文件被清空，然后写入data文件中的数据
	err = os.WriteFile(file2, data, 0666) //隐式的打开和关闭文件
	if err != nil {
		fmt.Printf("write file err= %v\n", err)
	}
*/

// bufio包下的函数,使用前需要加bufio.
/*--------1
func NewReader
func NewReader(rd io.Reader) *Reader
NewReader创建一个具有默认大小缓冲、从r读取的*Reader。
使用说明：
reader := bufio.NewReader(a) //a为Open或者OpenFile返回的文件变量名,reader为接收变量
*/

/*--------2
func NewWriter
func NewWriter(w io.Writer) *Writer
NewWriter创建一个具有默认大小缓冲、写入w的*Writer。
使用说明：
writer := bufio.NewWriter(c)//c为Open或者OpenFile返回的文件变量名,writer为接收变量
*/

/*--------3
func (*Reader) ReadString
func (b *Reader) ReadString(delim byte) (line string, err error)
ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。如果
ReadString方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。
当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
使用说明:
	for {
		// \n读取一行的内容,直到换行为止
		str, err := reader.ReadString('\n') //调用reader为bufio.NewReader中的接收的文件变量名
		//读到一个换行符就结束
		if err == io.EOF { //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
*/

/*--------4
func (*Writer) WriteString
func (b *Writer) WriteString(s string) (int, error)
WriteString写入一个字符串。返回写入的字节数。如果返回值nn < len(s)，还会返回一个错误说明原因。
*/

// io包下的函数,使用前需要加io.
/*--------1
type Reader
type Reader interface {
    Read(p []byte) (n int, err error)
}
Reader接口用于包装基本的读取方法。

Read方法读取len(p)字节数据写入p。它返回写入的字节数和遇到的任何错误。即使Read方法返回值n < len(p)，
本方法在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，Read按惯例
会返回可以读取到的数据，而不是等待更多数据。

当Read在读取n > 0个字节后遭遇错误或者到达文件结尾时，会返回读取的字节数。它可能会在该次调用返回一个
非nil的错误，或者在下一次调用时返回0和该错误。一个常见的例子，Reader接口会在输入流的结尾返回非0的字
节数，返回值err == EOF或err == nil。但不管怎样，下一次Read调用必然返回(0, EOF)。调用者应该总是先处
理读取的n > 0字节再处理错误值。这么做可以正确的处理发生在读取部分数据后的I/O错误，也能正确处理EOF事件。

如果Read的某个实现返回0字节数和nil错误值，表示被阻碍；调用者应该将这种情况视为未进行操作。
*/

/*--------2
type Writer
type Writer interface {
    Write(p []byte) (n int, err error)
}
Writer接口用于包装基本的写入方法。

Write方法len(p) 字节数据从p写入底层的数据流。它会返回写入的字节数(0 <= n <= len(p))和遇到的任何导致
写入提取结束的错误。Write必须返回非nil的错误，如果它返回的 n < len(p)。Write不能修改切片p中的数据，
即使临时修改也不行。
*/

/*--------3
func Copy
func Copy(dst Writer, src Reader) (written int64, err error)
将src的数据拷贝到dst，直到在src上到达EOF或发生错误。返回拷贝的字节数和遇到的第一个错误。

对成功的调用，返回值err为nil而非EOF，因为Copy定义为从src读取直到EOF，它不会将读取到EOF视
为应报告的错误。如果src实现了WriterTo接口，本函数会调用src.WriteTo(dst)进行拷贝；否则如果
dst实现了ReaderFrom接口，本函数会调用dst.ReadFrom(src)进行拷贝。
使用说明:
reader := bufio.NewReader(a)
writer := bufio.NewWriter(c)
io.Copy(writer, reader)
*/
