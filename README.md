# GoHttpsClientDll
Ferramenta para requisições HTTPS com suporte a TLS1.2 até em Windows XP

## Começando
Instale o MinGW e instale o GCC para Windows X86
Utilize a versão 1.10.5 para windows/386 de Golang, que é a versão mais atual compatível com Windows XP.  
Utilize o comando "go get github.com/Cappta/GoHttpsClientDll" para obter uma cópia compilavel do código fonte.  
Vá até a pasta C:\Users\<usuario>\go\src\github.com\Cappta\GoHttpsClientDll e rode o Build.bat
Será criado o GoHttpsClient.dll

Tanto o código fonte quanto o executavel estarão em:
C:\Users\<usuario>\go\src\github.com\Cappta\GoHttpsClientDll

## Uso do Software
Essa DLL deve ser encapsulada para uso facilitado em outra linguagem de programação, como o GoHttpsClientForCSharp

## Ferramentas
1. Golang

## Regras de colaboração
Deve estar em sincronia com a classe DownloadService do Cappta.Gp
 
 ## Colaboradores
 - Sérgio Fonseca - _Autor inicial_ - SammyROCK
