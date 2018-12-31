#include "GoHttpsClient.h"
#include <stdlib.h>
#include <stdio.h>

__declspec(dllexport) GoInt CreateClient() {
    return GoCreateClient();
}

__declspec(dllexport) GoUint8 SetClientTimeout(GoInt clientId, GoInt timeoutInSeconds) {
    return GoSetClientTimeout(clientId, timeoutInSeconds);
}

__declspec(dllexport) GoInt CreateRequest(char* method, char* url, void* body) {
    return GoCreateRequest(method, url, body);
}

__declspec(dllexport)  GoUint8 SetRequestHeader(GoInt requestID, char* key, char* value){
    return GoSetRequestHeader(requestID, key, value);
}

__declspec(dllexport)  GoInt PerformRequest(GoInt clientID, GoInt requestID){
    return GoPerformRequest(clientID, requestID);
}

__declspec(dllexport)  char* GetResponseStatus(GoInt responseID){
    return GoGetResponseStatus(responseID);
}

__declspec(dllexport)  GoInt GetResponseStatusCode(GoInt responseID){
    return GoGetResponseStatusCode(responseID);
}

__declspec(dllexport)  char** GetResponseHeaderKeys(GoInt responseID){
    return GoGetResponseHeaderKeys(responseID);
}

__declspec(dllexport)  char** GetResponseHeaderValue(GoInt responseID, char* key){
    return GoGetResponseHeaderValue(responseID, key);
}

__declspec(dllexport)  void* GetResponseBody(GoInt responseID){
    return GoGetResponseBody(responseID);
}

__declspec(dllexport)  char* GetError(GoInt errorID){
    return GoGetError(errorID);
}

__declspec(dllexport)  void ReleaseObject(GoInt objectID){
    return GoReleaseObject(objectID);
}

__declspec(dllexport)  void Free(void* ptr){
    free(ptr);
}