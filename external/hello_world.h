
#include <jni.h>
#ifndef CALLC_H
#define CALLC_H

JNIEnv* create_vm(JavaVM **jvm);
void init();
void mainMethod();
int squareMethod();
void power();
void invoke_class(JNIEnv* env);
char* hello(char* s) ;
int x();
#endif