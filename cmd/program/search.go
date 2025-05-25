package program 

import(
	"fmt"
	"github.com/go-git/go-git/v5"
)

type Error struct  {
	Messafe string
}

func (e*Error) Error() string { 
	return fmt.Sprintf("Error: %s", e.Messafe)
}

func Search(args[]string) (string,error) {
	searchTerm := args[0]
	option := ["-un", "-nt", "-np"]
	if len(args) == 0 {
		return "", &Error{Messafe: "No search term provided"}
	}

	switch option { 
	case "-un":
		 repos,err := git.PlainOpen(".")
		 if err != nil { 
			return "", &Error{Messafe: "Failed to open git repository"}
		 }
	}
   if err:= &Error{Messafe: "Not implemented"}; err != nil {
		return "", err
   }
	
}