package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/Chaitanya-019/grpc-address-book-example-go/tutorial"
	"github.com/golang/protobuf/proto"
)

func writePerson(w io.Writer, p *pb.Person) {
	fmt.Fprintln(w, "Person ID : ", p.Id)
	fmt.Fprintln(w, "Name : ", p.Name)
	if p.Email != "" {
		fmt.Fprintln(w, "E-mail Address : ", p.Email)
	}

	for _, pn := range p.Phones {
		switch pn.Type {
		case pb.Person_HOME:
			fmt.Fprint(w, " Home Phone #: ")
		case pb.Person_MOBILE:
			fmt.Fprint(w, " Mobile Phone #: ")
		case pb.Person_WORK:
			fmt.Fprint(w, " Work Phone #: ")
		}
		fmt.Fprint(w, pn.Number)
	}
}

func listPeople(w io.Writer, book *pb.AddressBook) {
	for _, p := range book.People {
		writePerson(w, p)
	}
}

// Main reads the entire address book from a file and prints all the
// information inside.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// [START unmarshal_proto]
	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	// [END unmarshal_proto]

	listPeople(os.Stdout, book)
}
