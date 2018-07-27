all: run


helloWorld.class: helloWorld.java
	/usr/lib/jvm/java-8-oracle/bin/javac helloWorld.java

hello_world: hello_world.c
	gcc -c \
	-L /usr/lib/jvm/java-8-oracle/jre/lib/amd64/server/ \
	-ljvm \
	-I /usr/lib/jvm/java-8-oracle/include/ \
	-I /usr/lib/jvm/java-8-oracle/include/linux/ \
	hello_world.c -ljvm
	 ar rs libclibrary.a *.o

run: helloWorld.class hello_world
	export LD_LIBRARY_PATH="/usr/lib/jvm/java-8-oracle/jre/lib/amd64/server/"
	go test -v -timeout 30s external_test.go  -run ^TestALL$
# 	./hello_world

clean:
	rm -f helloWorld.class hello_world
