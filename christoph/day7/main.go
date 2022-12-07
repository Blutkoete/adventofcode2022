package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name   string
	Size   int
	Parent *Directory
}

func NewFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

type Directory struct {
	Name        string
	Files       []*File
	Directories []*Directory
	Parent      *Directory
}

func NewDirectory(name string) *Directory {
	return &Directory{
		Name:        name,
		Files:       make([]*File, 0),
		Directories: make([]*Directory, 0),
	}
}

func main() {
	round2()
}

func round1() {
	root := readInput("input.txt")
	directories := make([]*Directory, 0)
	directories = append(directories, root)
	totalSize := 0
	for len(directories) > 0 {
		currentDir := directories[0]
		directories = directories[1:]
		directories = append(directories, currentDir.Directories...)
		directorySize := getSize(currentDir)
		fmt.Printf("Processing %s = %d\n", currentDir.Name, directorySize)
		if directorySize <= 100000 {
			totalSize = totalSize + directorySize
		}
	}
	fmt.Printf("%d\n", totalSize)
}
func round2() {
	root := readInput("input.txt")
	totalSpace := 70000000
	neededSpace := 30000000
	rootSize := getSize(root)
	currentFreeSpace := totalSpace - rootSize
	tooDelete := neededSpace - currentFreeSpace
	fmt.Printf("tooDelete %d", tooDelete)
	candidates := make([]int, 0)
	directories := make([]*Directory, 0)
	directories = append(directories, root)
	for len(directories) > 0 {
		currentDir := directories[0]
		directories = directories[1:]
		directories = append(directories, currentDir.Directories...)
		directorySize := getSize(currentDir)
		if directorySize > tooDelete {
			fmt.Printf("Fit %s = %d\n", currentDir.Name, directorySize)
			candidates = append(candidates, directorySize)
		}
	}
	sort.Ints(candidates)
	for _, val := range candidates {
		fmt.Printf("%d\n", val)
	}

}

func getSize(dir *Directory) int {
	directories := make([]*Directory, 0)
	directories = append(directories, dir)
	totalSize := 0
	for len(directories) > 0 {
		currentDir := directories[0]
		directories = directories[1:]
		directories = append(directories, currentDir.Directories...)
		directorySize := 0
		for _, file := range currentDir.Files {
			directorySize = directorySize + file.Size
		}
		totalSize = totalSize + directorySize
	}
	return totalSize
}

func readInput(file string) *Directory {
	readFile, err := os.Open(file)
	checkError(err)
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var root *Directory = NewDirectory("/")
	var currentDir *Directory
	readingDir := false
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if readingDir && strings.Index(line, "$") != 0 {
			if strings.Index(line, "dir") == 0 {
				parts := strings.Split(line, " ")
				subdir := NewDirectory(parts[1])
				subdir.Parent = currentDir
				currentDir.Directories = append(currentDir.Directories, subdir)
			} else {
				parts := strings.Split(line, " ")
				size, err := strconv.Atoi(parts[0])
				checkError(err)
				currentDir.Files = append(currentDir.Files, NewFile(parts[1], size))
			}
		}
		// Is it a command
		if strings.Index(line, "$") == 0 {
			readingDir = false
			if line == "$ cd /" {
				currentDir = root
			} else if line == "$ cd .." {
				currentDir = currentDir.Parent
			} else if strings.Index(line, "$ cd ") == 0 {
				dirName := line[5:]
				currentDir = getDir(currentDir, dirName)
			} else if line == "$ ls" {
				readingDir = true
			}
		}
	}
	return root
}

func getDir(dir *Directory, name string) *Directory {
	for _, subdir := range dir.Directories {
		if subdir.Name == name {
			return subdir
		}
	}
	return nil
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
