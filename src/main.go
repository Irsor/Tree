package main

import (
	"log"
	"os"
	"strconv"
)

func readDirectory(path string, symbols []bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	filesCount := len(files)
	for index, file := range files {
		var newSymbols []bool
		for _, data := range symbols {
			if data {
				print("│   ")
			} else {
				print("    ")
			}
		}

		if index == filesCount-1 {
			print("└───" + file.Name())
			newSymbols = append(symbols, false)
		} else {
			print("├───" + file.Name())
			newSymbols = append(symbols, true)
		}

		if !file.IsDir() {
			fileInfo, err := file.Info()
			if err != nil {
				log.Fatal(err)
			}

			if fileInfo.Size() == 0 {
				println(" (empty)")
			} else {
				println(" (" + strconv.Itoa(int(fileInfo.Size())) + "b)")
			}
		} else {
			println()
			readDirectory(path+"/"+file.Name(), newSymbols)
		}
	}
}

func main() {
	switch argsCount := len(os.Args); argsCount {
	case 1:
		readDirectory(".", make([]bool, 0))
		break
	case 2:
		readDirectory(os.Args[1], make([]bool, 0))
		break

		// Сделаю потом игнорирование файлов
	// case 3:
	// 	if os.Args[2] == "-f" {
	// 		readDirectory(os.Args[1], make([]bool, 0))
	// 	} else {
	// 		println("Invalid argument! Program shutting down...")
	// 	}
	// 	break
	default:
		println("Too many arguments! Program shutting down...")
		break
	}
}
