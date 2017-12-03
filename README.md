# BookTracker

[![Join the chat at https://gitter.im/BookTracker/Lobby](https://badges.gitter.im/BookTracker/Lobby.svg)](https://gitter.im/BookTracker/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
Keep track of books you've read. Team project for Software Engineering class. (Name tentative)

Server:

Create directory structure:
$GOPATH/src/github.com/mortonar
git glone https://github.com/mortonar/acs560_course_project.git
go get github.com/lib/pq
go get github.com/jinzhu/gorm
go get github.com/satori/go.uuid
from: $GOPATH/src/github.com/mortonar/acs560_course_project run: bootstrap.sh as root

Client Windows:
Map local ip address to domain.
Edit C:\Windows\System32\drivers\etc\hosts adding line below (replacing IP with local IP).
192.168.1.108 booktracker.com

