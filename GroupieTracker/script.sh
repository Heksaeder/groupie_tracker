cd /
cd 'C:\Program Files\MySQL\MySQL Server 8.0\bin'
./mysql --user="root" --database="groupes" -t --execute="SELECT name AS 'Band Name', creationDate AS 'Creation Date', firstAlbum AS 'First Album' FROM groupes;"
read
cd 'C:\wamp64\www\Groupie_Tracker'
go run main.go
read