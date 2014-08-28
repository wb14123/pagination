
Pagination middleware for martini.

It will return a handler function that parse `since_id`, `max_id` and `count` in the url query, then inject them to the `Pagination` struct.

The default value of `since_id` is 0, `max_id` is the max int number, `count` is 10.

The max value of `count` is 100.

You could specify what Json struct should be render to the client while error occurs in parsing.
