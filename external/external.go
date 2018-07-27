package external

// #cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/
// #cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/linux/
// #cgo LDFLAGS: -L/usr/lib/jvm/java-8-oracle/jre/lib/amd64/server/ -ljvm
// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: ${SRCDIR}/libclibrary.a
// #include <stdio.h>
// #include <jni.h>
// #include <hello_world.h>
import "C"
import (
	"fmt"
)

func FillReport(filename, json string, parameter map[string]interface{}) uintptr {

	var t uintptr
	return t
}

func Export(t string, jasperPrints uintptr) []byte {

	return nil
}

func MainC() {
	C.x()
}

func Init() {
	C.init()
}

func MainMethodC() {
	C.mainMethod()
}
func SquareMethodC() {
	value := C.squareMethod()
	fmt.Println(value)
}
func PowerC() {
	C.power()
}

func HelloWorld(name string) string {
	cs := C.CString(name)
	return C.GoString(C.hello(cs))
}
