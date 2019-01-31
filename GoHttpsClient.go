package main

import (
	"C"
	"bytes"
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"sync"
	"time"
)
import (
	"net/http/cookiejar"
	"unsafe"
)

var (
	mutex     = &sync.Mutex{}
	objectMap = make(map[int]interface{})

	errInvalidClient  = errors.New("Invalid Client Id")
	errInvalidRequest = errors.New("Invalid Request Id")
	errInvalidObject  = errors.New("Invalid Object Id")

	lastInput = ""
)

func main() {

}

//export GoCreateClient
func GoCreateClient() (id int) {
	id = allocateID()

	client := &http.Client{}
	client.Jar, _ = cookiejar.New(nil)
	setObject(id, client)
	return
}

//export GoSetClientTimeout
func GoSetClientTimeout(clientID, timeoutInSeconds int) (result bool) {
	client, exists := getClient(clientID)
	if exists == false {
		return false
	}

	client.Timeout = time.Second * time.Duration(timeoutInSeconds)
	return true
}

//export GoCreateRequest
func GoCreateRequest(cMethod, cUrl *C.char, cBody unsafe.Pointer) (id int) {
	id = allocateID()

	method := C.GoString(cMethod)
	url := C.GoString(cUrl)
	body := ptrToByteSlice(cBody)
	request, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		setObject(id, err)
		return
	}
	setObject(id, request)
	return
}

//export GoSetRequestHeader
func GoSetRequestHeader(requestID int, cKey, cValue *C.char) bool {
	request, exists := getRequest(requestID)
	if exists == false {
		return false
	}

	key := C.GoString(cKey)
	value := C.GoString(cValue)
	request.Header.Set(key, value)
	return true
}

//export GoPerformRequest
func GoPerformRequest(clientID, requestID int) (id int) {
	id = allocateID()

	client, exists := getClient(clientID)
	if exists == false {
		setObject(id, errInvalidClient)
		return
	}

	request, exists := getRequest(requestID)
	if exists == false {
		setObject(id, errInvalidRequest)
		return
	}

	result, err := client.Do(request)
	if err != nil {
		setObject(id, err)
		return
	}

	setObject(id, result)
	return
}

//export GoGetResponseStatus
func GoGetResponseStatus(responseID int) *C.char {
	response, exists := getResponse(responseID)
	if exists == false {
		return C.CString("")
	}

	return C.CString(response.Status)
}

//export GoGetResponseStatusCode
func GoGetResponseStatusCode(responseID int) int {
	response, exists := getResponse(responseID)
	if exists == false {
		return -1
	}

	return response.StatusCode
}

//export GoGetResponseHeaderKeys
func GoGetResponseHeaderKeys(responseID int) unsafe.Pointer {
	response, exists := getResponse(responseID)
	if exists == false {
		return nil
	}

	headers := make([]string, 0, len(response.Header))
	for key := range response.Header {
		headers = append(headers, key)
	}
	return stringSliceToPtr(headers)
}

//export GoGetResponseHeaderValue
func GoGetResponseHeaderValue(responseID int, cHeader *C.char) unsafe.Pointer {
	response, exists := getResponse(responseID)
	if exists == false {
		return nil
	}

	header := C.GoString(cHeader)
	return stringSliceToPtr(response.Header[header])
}

//export GoGetResponseBody
func GoGetResponseBody(responseID int) unsafe.Pointer {
	response, exists := getResponse(responseID)
	if exists == false {
		return nil
	}

	body, _ := ioutil.ReadAll(response.Body)
	return byteSliceToPtr(body)
}

//export GoGetError
func GoGetError(errorID int) *C.char {
	object, exists := getObject(errorID)
	if exists == false {
		return C.CString(errInvalidObject.Error())
	}

	err, ok := object.(error)
	if ok == false {
		return C.CString("")
	}

	return C.CString(err.Error())
}

//export GoReleaseObject
func GoReleaseObject(objectID int) {
	mutex.Lock()
	defer mutex.Unlock()

	delete(objectMap, objectID)
}

func getClient(id int) (*http.Client, bool) {
	object, exists := getObject(id)
	if exists == false {
		return nil, false
	}
	request, ok := object.(*http.Client)
	return request, ok
}

func getRequest(id int) (*http.Request, bool) {
	object, exists := getObject(id)
	if exists == false {
		return nil, false
	}
	request, ok := object.(*http.Request)
	return request, ok
}

func getResponse(id int) (*http.Response, bool) {
	object, exists := getObject(id)
	if exists == false {
		return nil, false
	}
	response, ok := object.(*http.Response)
	return response, ok
}

func allocateID() int {
	mutex.Lock()
	defer mutex.Unlock()

	for i := 1; i < math.MaxInt32; i++ {
		if _, exists := objectMap[i]; exists == false {
			objectMap[i] = nil
			return i
		}
	}

	return 0
}

func getObject(id int) (value interface{}, exists bool) {
	mutex.Lock()
	defer mutex.Unlock()

	value, exists = objectMap[id]
	return
}

func setObject(id int, value interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	objectMap[id] = value
}

func stringSliceToPtr(slice []string) unsafe.Pointer {
	var ptr unsafe.Pointer
	ptrSize := int(unsafe.Sizeof(ptr))
	arraySize := len(slice) + 1
	ptr = C.malloc(C.size_t(ptrSize * arraySize))
	cSlice := (*[1 << 28]*C.char)(ptr)

	for i := 0; i < len(slice); i++ {
		cSlice[i] = C.CString(slice[i])
	}
	cSlice[len(slice)] = nil

	return ptr
}

func byteSliceToPtr(slice []byte) unsafe.Pointer {
	var ptr unsafe.Pointer
	ptrSize := int(unsafe.Sizeof(ptr))
	ptr = C.malloc(C.size_t(ptrSize + len(slice)))

	size := (*int)(ptr)
	*size = len(slice)

	dataPtr := unsafe.Pointer(uintptr(ptr) + uintptr(ptrSize))
	cSlice := (*[1 << 28]byte)(dataPtr)
	copy(cSlice[0:len(slice)], slice)

	return ptr
}

func ptrToByteSlice(ptr unsafe.Pointer) []byte {
	ptrSize := int(unsafe.Sizeof(ptr))

	size := (*int)(ptr)

	dataPtr := unsafe.Pointer(uintptr(ptr) + uintptr(ptrSize))
	cSlice := (*[1 << 28]byte)(dataPtr)
	slice := make([]byte, *size)
	copy(slice, cSlice[0:len(slice)])

	return slice
}
