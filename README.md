# smsapi

A REST API written in GoLang to send SMS to a Kannel server

[Run](#run)

[Test](#test)

[Setup](#setup)

# Run

Set the credentials to be used by the API in the environment variables SMSAPI_USER and SMSAPI_PASS

```
go run main.go
```

Access the API at http://localhost:8080/outbound/sms/

# Test

Unit tests:

```
go test -v smsapi/router smsapi/handlers smsapi/models
```

This outputs the names of the tests being run and the result, along with the time taken to run each test. It also outputs a summary result and the total time taken to run all the tests.

API tests:

Install newman using ``` npm install newman ```

Execute the tests using ``` newman run /smsapi/SMSApiTestSuite.postman_collection.json 	```

# Setup

REST API:

```
go install smsapi/
```

Run the executable in /smsapi/bin


Kannel:

Install the Kannel software by following the instructions in the [documentation.](http://kannel.org/download/1.4.4/userguide-1.4.4/userguide.html#AEN340)

Replace kannel.conf in /etc/kannel with this [kannel.conf.](https://github.com/shilps1583/smsapi/blob/master/conf/kannel.conf)

Modify /etc/default/kannel and uncomment the START_SMSBOX=1 line

Run Kannel as a service using the command

```
sudo service kannel start
```

Check the status of the Kannel gateway [here.](http://localhost:13000/status)

SMPPSIM:

Install the Kannel software by following the instructions in the [documentation.](http://www.seleniumsoftware.com/user-guide.htm#installation)

Replace conf/smppsim.props in the installation directory with this [smppsim.props.](https://github.com/shilps1583/smsapi/blob/master/conf/smppsim.props)

Run the following command

```
sh startsmppsim.sh
```

Confirm the setup using the admin tool [here.](http://localhost:88/)
