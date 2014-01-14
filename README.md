# Go/Java WordSorter Comparison

The Java code in this project was originally an assignment from the university. The goal was to write a Java program that accepted a list of words, and wrote them out in sorted order, using threads for speed. I used this project to see how Google's new language, Go, compared against Java with emphasis on LOC and performance. I'm glad to say that Go easily beats Java in both these cases. Go is also dramatically easier on system memory than Java.

## Use

Go project can be run by entering the Go directory and issuing the following: `go run *.go 128 sowpods.txt out.txt`
128 is the number of goroutines, sowpods.txt is the input and out.txt is the output.

Java project can be run by entering Java directory and compile using `javac Sort.java` then run by issuing `java Sort 128 sowpods.txt out.txt`

## Results

On my Macbook Pro 13" late-2013 the results where the following:
Go: 460ms
Java: 620ms
