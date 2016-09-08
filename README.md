[ ![Codeship Status for fifthsegment/DslReports-CLI](https://codeship.com/projects/11a794a0-5188-0134-df5b-1ab2f5cbd726/status?branch=master)](#)

# DslReports-CLI

DslReports.com's commandline speed testing utility.

![Screenshot-of-the-test](http://i.imgur.com/VlsYj9h.png)




### How do I use this?

First, get the required dependencies:

1. `go get github.com/codegangsta/cli`
2. `go get github.com/fifthsegment/DslReports`

Then simply run : `go run main.go`

### Additional Flags

#### -d --debug

Use this to have useful debugging information output to the screen as well as saved in a folder called dslreports.com in the current working directory.


#### -o --output json|csv

Saves the speed test results in the format specified as the first argument


#### --down number


Specify the number of streams to be used for the download test.

#### --up number

Specify the number of streams to be used for the upload test.


