# report
The Task
You need to implement the Report Exporting Service. The purpose of this service is to take abstract
reports and generate PDF documents and XML files out of them.
Services within the Suade system are stand-alone (flask) python applications with a thin layer of
(RESTful) web services as the interface.
You have two abstract reports made available to you on a PostgreSQL server located at:
postgres://candidate.suade.org/suade
Username: interview Password: LetMeIn
A wireframe of a potential PDF template is attached at the bottom of this document.

Setup GOLang environment:
```
apt install golang
add-apt-repository ppa:masterminds/glide -y && apt-get update && apt-get install glide

export GOPATH=~/go
cd ~/go/src
git clone <project>

glide up

go run reporter.go
```
