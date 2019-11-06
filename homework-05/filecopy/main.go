//
//  Задача: Напишите утилиту для копирования файлов, используя пакет​ flag .
//

package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func main() {
	isVerbose := flag.Bool("v", false, "verbose mode")

	flag.Parse()

	if *isVerbose {
		fmt.Println("#isVerbose = on")
	}

	if flag.NArg()!=2 {
		fmt.Println("Утилита для копирования файлов")
		fmt.Printf("Использование:\n\t filecopy [-v] src dst\n")
		fmt.Println("Опции:\n\t -v описывать действия\n")
		fmt.Println("Пример:\n\t filecopy file1 file2\n\t filecopy -v file1 file2\n\t")
		return
	} else {
		srcFile:=flag.Arg(0)
		dstFile:=flag.Arg(1)
		bytes, err := copy (srcFile, dstFile)
		if err==nil {
			if *isVerbose {
				fmt.Printf("Файл '%v' успешно cкопирован в '%v' размером %v байт\n",srcFile,dstFile,bytes)
			}
		} else {
			fmt.Println(err)
		}
	}
}