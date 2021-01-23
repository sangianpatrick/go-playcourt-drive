![Go](https://github.com/sangianpatrick/go-playcourt-drive/workflows/Go/badge.svg?branch=master)
[![codecov](https://codecov.io/gh/sangianpatrick/go-playcourt-drive/branch/master/graph/badge.svg?token=sr4Nok6fwp)](https://codecov.io/gh/sangianpatrick/go-playcourt-drive)

# go-playcourt-drive
An unofficial playcourt drive golang sdk.

# Install
```
$ go get github.com/sangianpatrick/go-playcourt-drive
```

# How to use
```
func main() {
	c, err := drive.NewClient(&drive.Config{
		Host:            "http://drive.playcourt.test",
		Username:        "username",
		Password:        "password",
		MaxRetry:        3,
		BackoffInMillis: 50,
	})
	if err != nil {
		log.Fatal(err)
	}
	sessionID, err := c.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sessionID)
}
```
