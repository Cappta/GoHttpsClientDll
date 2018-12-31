go build -buildmode=c-archive GoHttpsClient.go
gcc -shared -pthread GoHttpsClient.c GoHttpsClient.a -o GoHttpsClient.dll -lWinMM -lntdll -lWS2_32
del GoHttpsClient.a
xcopy C:\Users\sergio.fonseca\go\src\github.com\Cappta\GoHttpsClientDll\GoHttpsClient.dll C:\GIT\GoHttpsClientForCSharp\GoHttpsClientForCSharp\GoHttpsClient.dll /Y