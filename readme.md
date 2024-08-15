# SHOPS SCRAPING
 Get many shops product at the same time

## For non docker users
### Requirements
- Go 1.22.4
- NodeJS 22.x
- Internet connection
- A browser
### How to install
- install frontend dependencies: In the front folder run: `yarn`
- install server dependencies: In the root folder run: `go install`
- create a .env` file by  `.env.example` file
- fill the `.env`
- build front app: run `yarn build` in `front` folder
### How to launch app
- run go server: `go run main.go`
nb: the first launch is going to download chromium


## For dockers users
build and run dockerfile