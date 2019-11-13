# MaxIdleConnsPerHost

The http.defaultClient has very low `MaxIdleConnsPerHost` (2). If more than 2 simultaneous requests
are made, all the idle connections are occupied. This results creation of new temporary connections
that will be put into TIME_WAIT state after completion.
