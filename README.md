# Client for Google Drive API.

This client created for easily upload file to a Google Drive directory.

## How can I use this?

### Create a project, service account and generate credentials in the [Google Cloud console](https://console.cloud.google.com).

_Documentation from: 27-12-2023_

- add the project name, and click "OK";
- after this you will see the project dashboard;
- choose "APIs & Services", you will see a drop-down, choose "Enabled APIs & Services"
  in the bottom you will see the button "+ Enable APIs & Services", click;
- in fetched list, you need to find "Google Drive API", click;
- choose "Enable";
- in showed window choose "credentials";
- choose "+ create credentials";
- in drop-down choose "Service account";
- set name and click "Done";
- after this action you will be redirected to "APIs & Services" view,
  and you will see aew account list with "Service accounts", when will be your new account, click on his name.
- choose "Keys";
- choose "Add key";
- choose JSON;
- when you click "OK", the **file with credentials** will be downloaded to your computer.
  This file you will be used for authorize requests;
- on the left panel, choose "Service accounts"
- you will see **email address of a new service account**.
  You will need this address for add account to the Google Drive directory as editor;

After, you need:
- Add a email of the service to the Google Drive directory, when you will write files
- Add a file with credentials when you need. You need add this path in the client constructor.
- Copy a directory id, where you will be write files, and add this to parents in the client constructor.

```go
package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func readFile(filename string) ([]byte, error) {
  fileContent, err := os.ReadFile(filename)
  if err != nil {
    log.Fatal(err)
  }
  return fileContent, err
}

func main() {
  client, err := NewClient("credentials.json_file_path", []string{"directory_id"})
  if err != nil {
    log.Fatal(err)
  }

  filename := "example.txt"
  fileContent, err := readFile(filename)
  if err != nil {
    log.Fatal(err)
  }

  contentType := http.DetectContentType(fileContent)
  fileID, err := client.UploadFile(filename, contentType, fileContent)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%s\n", fileID)
}

```

