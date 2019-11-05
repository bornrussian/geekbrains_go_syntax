package main

import (
	"flag"
	"fmt"
)

func main() {
	isOverWrite := flag.Bool("w", false, "OverWrite?")

	flag.Parse()

	if *isOverWrite {
		fmt.Println("# isOverWrite = on")
	} else {
		fmt.Println("# isOverWrite = off")
	}

	if flag.NArg()!=2 {
		fmt.Println("Утилита для копирования файлов")
		fmt.Printf("Использование:\n\t filecopy [-w] src dst\n")
		fmt.Println("Опции:\n\t -w перезаписать dst если существует\n")
		fmt.Println("Пример:\n\t filecopy file1 file2\n\t filecopy -w file1 file2\n\t")
		return
	} else {
		srcFile:=flag.Arg(0)
		dstFile:=flag.Arg(1)
		fmt.Println(srcFile,dstFile)
	}
}