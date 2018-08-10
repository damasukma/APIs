package core

import(
	"os"
	"log"
)

func DEBUG(file string , data []byte)(string, error){
	var name string
	name = "DEBUG"
	if file != "" {
		name = file
	}
	fs, err := os.OpenFile("log/" + name, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666 )
	if err != nil{
		log.Fatalf("error opening file : %v", err)
		return "", err
	}

	defer fs.Close()
	fs.Write(data)

	return "Success", nil
	

	// log.SetOutput(fs)
	// log.Println("This is a test log entry")
}