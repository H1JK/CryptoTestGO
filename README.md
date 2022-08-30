# CryptoTestGO

Yet another benchmark tool.  

## Run with one-click script

Script will use `-cpu 1 -benchmem` argument to launch the test, which means it will only test single core performance.  

Bash (cURL):  
```shell
curl -fsL https://raw.fastgit.org/H1JK/CryptoTestGO/master/run.sh | bash
```

Bash (Wget):
```shell
wget https://raw.fastgit.org/H1JK/CryptoTestGO/master/run.sh -O runCryptoTestGO.sh && chmod +x ./runCryptoTestGO.sh && ./runCryptoTestGO.sh
```

## Run with code
```shell
go test -cpu 1 -benchmem -bench=BenchmarkCrypto ./...
```

## License
This is licensed under MIT license and includes code from the Go standard library.  
