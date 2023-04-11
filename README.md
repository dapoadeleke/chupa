# Chupa
Chupa is a command-line tool written in Golang which makes HTTP requests and prints the address of the request along with the MD5 hash of the response. The tool is capable of performing requests in parallel to complete the task sooner. The order in which addresses are printed is not important. The tool is also able to limit the number of parallel requests to prevent exhausting local resources.

## Installation
1. Clone the Golang source code using the terminal: `git clonehttps://github.com/dapoadeleke/chupa.git`
2. Navigate to the root directory using command `cd chupa`
3. Build the project with command `go build cmd/chupa.go`

## Usage
The basic syntax for running MyHTTP is as follows:
`./chupa [flags] url1 url2 url3 ...`
The flags are as follows:
`-parallel`: The maximum number of parallel requests to make. Default value is 10.

## Sample Request
`./chupa -parallel 3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com`

## Sample Response
`http://google.com 8ff1c478ccca08cca025b028f68b352f` <br/>
`http://adjust.com 6b2560b9a5262571258cc173248b7492` <br/>
`http://yandex.com 4baab01ff9ff0f793bf423aeef539c9d` <br/>
`http://facebook.com ccae5ffa91c4936aef3efd5091a43f3e` <br/>
`http://twitter.com 857efe81a54c8b5c2241846ac4f08e66` <br/>
`http://reddit.com/r/funny ff3b2b7dcd9e716ca0adcbd208061c9a` <br/>
`http://reddit.com/r/notfunny ff3b2b7dcd9e716ca0adcbd208061c9a` <br/>
`http://yahoo.com e2d50a30b7bfbfda097d72e32578c6a6` <br/>
`http://baroquemusiclibrary.com 8e5138a0111364f08b10d37ed3371b11`

The output displays the URL followed by the MD5 hash of the response in the format of `URL MD5_HASH`.



