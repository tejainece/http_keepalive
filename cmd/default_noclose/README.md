# Close response body

Response.Body should be closed even if the body is not required/read. Failing 
to close the body results in leaking OS sockets, which could result in quick
resource exhaustion.