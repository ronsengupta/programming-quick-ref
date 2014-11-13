all:
	javac hello/lib/HelloLib.java
	javac hello/HelloMain.java
	jar cvf hello.jar hello
	java -cp hello.jar hello.HelloMain
clean:
	rm -f hello/*.class hello/lib/*.class hello.jar
