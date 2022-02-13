
# Run the program in one terminal

```
go run main.go
```

# upload file from file path
 
```
curl --location --request GET http://127.0.0.1:8090/process \
--form file=@/Users/unknowntpo/repo/unknowntpo/sandbox/go/excelize/workbook/user_info.xlsx -O -J
```
