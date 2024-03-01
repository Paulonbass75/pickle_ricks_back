SHELL=cmd
STRIPE_SECRET=sk_test_51ObramDQjpQrxGsCFStvC2wurnyTgLDWKwfnR8az6Shuoy9LdPN5iJOa0TnGG0J6M2iq1c1x4fYP562xzu8sUwuU00fCNs8LYq
STRIPE_KEY=pk_test_51ObramDQjpQrxGsC7KJuDzuCu08G6DD9XNpNpoFPd3AdrJiOoiIqkqP17vL9NF31fcU1PDHPgpua96jJB1EUcbBF00nymp22AS
Pickle_ricks_front_PORT=8080
Pickle_ricks_back_PORT=8081



## build: builds all binaries
build: clean build_front build_back
		@echo All binaries built!

## clean: cleans all binaries and runs go clean
clean:
		@echo Cleaning...
		@echo y | DEL /S dist
		@go clean
		@echo Cleaned and deleted binaries

## build_front: builds the front end
build_front:
		@echo Building front end...
		@go build -o dist/Pickle_ricks_front.exe ./cmd/web
		@echo Front end built!

## build_back: builds the back end
build_back:
		@echo Building back end...
		@go build -o dist/Pickle_ricks_back.exe ./cmd/api
		@echo Back end built!

## start: starts front and back end
start: start_front start_back

## start_front: starts the front end
start_front: build_front
		@echo Starting the front end...
		set STRIPE_KEY=${STRIPE_KEY}&& set STRIPE_SECRET=${STRIPE_SECRET}&& start /B .\dist\Pickle_ricks_front.exe -port=${Pickle_ricks_front_PORT}
		@echo Front end running!

## start_back: starts the back end
start_back: build_back
		@echo Starting the back end...
		set STRIPE_KEY=${STRIPE_KEY}&& set STRIPE_SECRET=${STRIPE_SECRET}&& start /B .\dist\Pickle_ricks_back.exe -port=${Pickle_ricks_back_PORT}
		@echo Back end running!

## stop: stops the front and back end
stop: stop_front stop_back
		@echo All applications stopped

## stop_front: stops the front end
stop_front:
		@echo Stopping the front end...
		@taskkill /IM Pickle_ricks_front.exe /F
		@echo Stopped front end

## stop_back: stops the back end
stop_back:
		@echo Stopping the back end...
		@taskkill /IM Pickle_ricks_back.exe /F
		@echo Stopped back end