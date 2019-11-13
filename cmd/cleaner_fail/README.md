# MaxConnsPerHost

New temporary connections will still be created if the number of simultaneous requests
exceeds the configured MaxIdleConnsPerHost.

This happens because by default MaxConnsPerHost is set to unlimited.