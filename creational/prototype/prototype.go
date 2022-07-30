package prototype

import "fmt"

type INode interface {
	Print(string)
	Clone() INode
}

type File struct {
	Name string
}

func (f *File) Print(indentation string) {
	fmt.Println(indentation + " - " + f.Name)
}

func (f *File) Clone() INode {
	file := new(File)
	file.Name = f.Name + "_copy"

	return file
}

type Folder struct {
	Children []INode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + " - " + f.Name)

	for _, child := range f.Children {
		child.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() INode {
	newFolder := new(Folder)
	newFolder.Name = f.Name + "_copy"

	for _, child := range f.Children {
		newFolder.Children = append(newFolder.Children, child.Clone())
	}

	return newFolder
}
