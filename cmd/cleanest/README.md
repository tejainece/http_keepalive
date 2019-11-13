# MaxConnsPerHost

Setting MaxConnsPerHost to a reasonable value prevents creation of new temporary connections. This
will block the goroutinue until an idle connection (or free connection slot) is available for reuse.