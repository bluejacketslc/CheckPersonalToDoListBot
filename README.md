# Check Personal To Do List Bot
Created by Christopher Limawan
This project is for creating a LINE bot to save To Do Lists and remind users every 07.00.
You can use this by add "@319auaer" as friend in LINE. Instructions about usage of bot included in Bot's help. Enjoy the Bot :D
## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.
### Prerequisites
- Install Go (mine was 1.14.2)
- Install Line Bot Library for Go
```
go get github.com/line/line-bot-sdk-go/linebot
```
- Install Godotenv (For read .env file)
```
go get github.com/joho/godotenv
```
- Install Cron in Go (For scheduling system (in this project is remind every 07.00) )
```
go get github.com/robfig/cron
```
- Install Mysql Driver in Go (in this section, you can adapt based on what DBMS you use)
```
go get github.com/go-sql-driver/mysql
```
- Install UUID library for Go
```
go get github.com/google/uuid
```

### Deployment (in Linux Environment)
- Prepare a LINE Channel to be a bridge between this project and the bot itself. (You may refer how to create one in [https://medium.com/@ariaseta/membuat-bot-line-cuman-20-menit-part-1-105b9839cc1c]())
- Prepare a server or VPS (you can rent it in several provider, low performance specifications enough for testing)
- Install all prerequisites
- Install MySQL for database (because I am using MySQL to save several data, e.g. To Do List, Subscribers for Reminder itself)
```
apt install mysql-server
```
- Install Nginx as the Web Service, that will access Executable Project itself
```
apt install nginx
```
- Install Git to use this project's code
```
apt install git
```
- Clone the project's code
- Create a copy of ".env.example", named ".env"
- Fill .env with following details:
```
CHANNEL_SECRET=[channel secret you get from first step]
CHANNEL_ACCESS_TOKEN=[channel access token you get from first step]

APP_HOST=localhost
APP_PORT=8080

DB_DRIVER=mysql
DB_NAME=[database name]
DB_SERVER=[database server IP address]
DB_PORT=[database server's port]
DB_USER=[database's user]
DB_PASS=[database's password]
```
- Build the project with go build
```
go build
```
- Create SSL certificate for link with LINE bot callback (because LINE only accept HTTPS), how you generate HTTPS, can refer from [https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-18-04](https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-18-04)
- Create nginx site settings, with the file name, for example, "cptodoreminder"
```
nano /etc/nginx/sites-available/todoreminder
```
- Fill with the following details, after that, save and exit using Ctrl+X, then Y and ENTER.
```
server {
        listen 443 ssl;
        server_name [your domain];
        ssl on;
        ssl_certificate [saved fullchain file];
        ssl_certificate_key [saved private key file];

        location / {
                proxy_pass http://localhost:8080;
                proxy_http_version 1.1;
                proxy_set_header Upgrade $http_upgrade;
                proxy_set_header Connection "upgrade";
        }
}
```
- Verify Nginx settings using following command
```
nginx -t
```
- After your Nginx settings are validated, reload the settings
```
nginx -s reload
```
- Then deploy your Go project using following command
```
./[name of generated application from go build]
```
- Enjoy

### Built With
- [Go](https://golang.org/) - Main Language Used

### Contributing
Please read [CONTRIBUTING.md](https://github.com/christ0208/todoreminder/blob/master/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

### Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/christ0208/todoreminder/tags).

### Authors
- Christopher Limawan - Initial Work - [Check Personal To Do List Bot](https://github.com/christ0208/todoreminder/)

See also the list of [contributors](https://github.com/christ0208/todoreminder/graphs/contributors) who participated in this project.

### Acknowledgements
- [https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html]()
- [https://github.com/line/line-bot-sdk-go]()
- [https://godoc.org/github.com/robfig/cron]()
- [https://towardsdatascience.com/use-environment-variable-in-your-next-golang-project-39e17c3aaa66]()
- [https://gist.github.com/PurpleBooth/109311bb0361f32d87a2]()
