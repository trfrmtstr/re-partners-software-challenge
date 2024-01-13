# RE Partners take home assignment

## Dependencies:
* Docker and Docker-Compose

## How to run:

1. ensure nothing is running locally on port 80 (if it's not an option then change port 80 in `./docker-compose.yaml` to desired port)
2. in terminal run `docker-compose up`
3. go to `http://localhost:80` in browser

## Notes:

* `packing` package contains algorithm implementation
* `handlers` package contains single REST api route handler

## Possible improvements:
* `POST /find-optimal-packing` always takes in as params both pack sizes and number of items
  * this may not be ideal if pack sizes is huge array which would cause wasting network bandwidth with each call
  * in that case it would be better to assign IDs to various pack sizes and call API referring to that ID - request body would only contain ID and number of items
  * downside is that we have to persistently store pack sizes and their IDs
* `POST /find-optimal-packing` utilizes dynamic programming solution which memoizes previous subproblems - we could store solutions to previous subproblems between calls, so we don't recompute them with each call for same pack sizes array, to effectively lower CPU usage
* UI is done in vanilla JS/HTML because it was fastest to make, for bigger projects this approach doesn't make sense and either templating engine (server side rendering) or single page app (React, Angular) would be better option
