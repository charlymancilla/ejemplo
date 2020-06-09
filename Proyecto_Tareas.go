/* Writing web aplications*/
package main

/* We import "fmt" y "io/ioutil"*/
import (
	"fmt"
	"io/ioutil"
)

/* data structure, A wiki consists of a series of interconnected pages, each of which has a title and a body (the page content).*/
type Page struct {
	Title string
	Body  []byte
}

/*We create a SAVE method on page, to we can addres,  this is a persistent storage*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*The function loadPage constructs the file name from the title parameter, reads the file's contents into a new variable body, and returns a pointer to a Page literal constructed with the proper title and body values.*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	/*Callers of this function can now check the second parameter; if it is nil then it has successfully loaded a Page. If not, it will be an error that can be handled by the caller (see the language specification for details).*/
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*After compiling and executing this code, a file named TestPage.txt would be created, containing the contents of p1. The file would then be read into the struct p2, and its Body element printed to the screen.*/
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}




