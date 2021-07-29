
# send-mail
This service will help you send mail to several people at the same time.

First, we need to create a config.json

```
{
  "addr": "127.0.0.1:8081"
  "mail": "example@gmail.com",
  "password": "Your Password",
  "host": "smtp.gmail.com",
  "port": "587"
}
```

If you have Yandex mail then

```
{
  "addr": "127.0.0.1:8081"
  "mail": "example@yandex.com",
  "password": "Your Password",
  "host": "smtp.yandex.ru",
  "port": "465"
}
```

Then we use the command "go run main.go"

If in the terminal "<--START-SERVER-->" then our service works in the port that you specified in config.json(addr)

Now we send form data.
```
{
  "mails": "["example@gmail.com", "example@yandex.com"]", //[]string
  "body": "your body", // string
  "title": "your title", // string
}
```
