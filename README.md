# bcrypt
It is a package in Go that implements Provos and Mazières’s bcrypt adaptive hashing algorithm to calculate the hash.

## To Use
--> first intialize your repo [optional]
<br>
--> install bcrypt package

## Intialize

```go mod init```
or specified a name to your repository like this
```go mod init github.com/shubhammishra-1```

## Install uuid package

```bash
go get golang.org/x/crypto/bcrypt
```

## Salt
```A salt is random text added to the string to be hashed. ```
For example, you don't hash my_secret_password; you hash something like 1jfSLKe$*@SL$#)(Sslkfs$34:my_secret_password. The reason for this is that it makes it hard to set up a "rainbow table" to brute-force the passwords, even if the entire database of hashed passwords is stolen. If every password has a different salt, only the very weakest passwords (like "password" or "123456", which you should prohibit anyway) will be guessed.

## Cost

```A cost is a measure of how many times to run the hash```
how slow it is. You want it to be slow. Again, this is a redundant layer of security for if the hashed passwords are stolen. It makes it prohibitively expensive to brute-force anything.

```MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword```
```MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword```
```DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword```




## bcrypt.GenerateFromPassword(password []byte, cost int) ([]byte, error)

GenerateFromPassword returns the bcrypt hash of the password at the given cost. If the cost given is less than MinCost, the cost will be set to DefaultCost, instead. Use CompareHashAndPassword, as defined in this package, to compare the returned hashed password with its cleartext version. GenerateFromPassword does not accept passwords longer than 72 bytes, which is the longest password bcrypt will operate on. 

to convert data structures like string objects etc... into bytes ```[]byte(string)```



## bcrypt.CompareHashAndPassword(hashedPassword, password []byte) error
CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent. Returns nil on success, or an error on failure. 

both hashedPassword and original Password {password} must be in bytes 


## bcrypt.Cost(hashedPassword []byte) (int, error)

Cost returns the hashing cost used to create the given hashed password. When, in the future, the hashing cost of a password system needs to be increased in order to adjust for greater computational power, this function allows one to establish which passwords need to be updated. 






