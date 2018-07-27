package embeded

// #cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/
// #cgo CFLAGS: -I/usr/lib/jvm/java-8-oracle/include/linux/
// #cgo LDFLAGS: -L/usr/lib/jvm/java-8-oracle/jre/lib/amd64/server/ -ljvm
// #include <stdio.h>
// #include <jni.h>
//
// JNIEnv* create_vm(JavaVM **jvm)
// {
//     JNIEnv* env;
//     JavaVMInitArgs args;
//     JavaVMOption options;
//     args.version = JNI_VERSION_1_6;
//     args.nOptions = 1;
//     options.optionString = "-Djava.class.path=./";
//     args.options = &options;
//     args.ignoreUnrecognized = 0;
//     int rv;
//     rv = JNI_CreateJavaVM(jvm, (void**)&env, &args);
//     if (rv < 0 || !env)
//         printf("Unable to Launch JVM %d\n",rv);
//     else
//         printf("Launched JVM! :)\n");
//     return env;
// }
//
//
// JNIEnv *env;
// jclass hello_world_class;
// void init(){
//
// 	JavaVM *jvm;
// 	env = create_vm(&jvm);
//
//     hello_world_class = (*env)->FindClass(env, "helloWorld");
// }
//
// void mainMethod()
// {
//
// 	jmethodID main_method;
// 	main_method = (*env)->GetStaticMethodID(env, hello_world_class, "main", "([Ljava/lang/String;)V");
// 	(*env)->CallStaticVoidMethod(env, hello_world_class, main_method, NULL);
// }
// int squareMethod()
// {
//     jclass hello_world_class;
//     jmethodID square_method;
//     jint number=20;
//     hello_world_class = (*env)->FindClass(env, "helloWorld");
//     square_method = (*env)->GetStaticMethodID(env, hello_world_class, "square", "(I)I");
//
//     jint Resultat =  (*env)->CallStaticIntMethod(env, hello_world_class, square_method, number);
//
//     return Resultat;
// }
// void power()
// {
//     jclass hello_world_class;
//     jmethodID power_method;
//     jint exponent=3;
//     jint number=20;
//     hello_world_class = (*env)->FindClass(env, "helloWorld");
//     power_method = (*env)->GetStaticMethodID(env, hello_world_class, "power", "(II)I");
//     printf("%d raised to the %d power is %d\n", number, exponent,
//      (*env)->CallStaticIntMethod(env, hello_world_class, power_method, number, exponent));
// }
// void invoke_class(JNIEnv* env)
// {
//     jclass hello_world_class;
//     jmethodID main_method;
//     jmethodID square_method;
//     jmethodID power_method;
//     jint number=20;
//     jint exponent=3;
//     hello_world_class = (*env)->FindClass(env, "helloWorld");
//     main_method = (*env)->GetStaticMethodID(env, hello_world_class, "main", "([Ljava/lang/String;)V");
//     square_method = (*env)->GetStaticMethodID(env, hello_world_class, "square", "(I)I");
//     power_method = (*env)->GetStaticMethodID(env, hello_world_class, "power", "(II)I");
//     (*env)->CallStaticVoidMethod(env, hello_world_class, main_method, NULL);
//     printf("%d squared is %d\n", number,
//         (*env)->CallStaticIntMethod(env, hello_world_class, square_method, number));
//     printf("%d raised to the %d power is %d\n", number, exponent,
//         (*env)->CallStaticIntMethod(env, hello_world_class, power_method, number, exponent));
//
// }
// char* hello(char* s) {
// 	return s;
// }
//
// int x()
// {
//     JavaVM *jvm;
//     JNIEnv *env;
//     env = create_vm(&jvm);
//     if(env == NULL)
//         return 1;
//     invoke_class(env);
//     return 1;
// }
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
