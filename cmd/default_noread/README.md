# Not reading body

Closing Response.Body is however not enough, all the data must be read before 
closing it (, eventhough the body is not required). Failing to do so will prevent
Go from reusing same connection for multiple requests using keep-alive.

Every new temporal connection made will result in TIME_WAIT socket state which will
lock up a socket for 4 minutes (configurable at kernel level). This could lead to
resource exhaustion.