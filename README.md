# go-playcourt-drive
An unofficial playcourt drive golang sdk.

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